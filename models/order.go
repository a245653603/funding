package models

import (
	"funding/enums"
	"funding/forms"
	"funding/objects"
	"funding/resultModels"
)

// 订单
type Order struct {
	BaseModel
	UserId           uint64            // 买家 Id
	Name             string            // 收件人姓名
	Address          string            // 收件人地址
	Phone            string            // 收件人电话
	SellerId         uint64            // 卖家 Id
	ProductId        uint64            // 产品 Id
	ProductPackageId uint64            // 套餐 Id
	Nums             int               // 购买数量
	UnitPrice        float64           // 单价
	TotalPrice       float64           // 总价
	Status           enums.OrderStatus // 订单状态
	CheckingNumber   string            // 物流单号
}

// 根据订单的 ID 来获取订单
func FindOrderById(orderId uint64) (*Order, error) {
	var ret Order
	err := db.First(&ret, orderId).Error
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

// 根据用户 ID 来获取订单列表
func FindOrdersByUserId(userId uint64) ([]*Order, error) {
	var rets []*Order
	err := db.Find(&rets).Where("user_id = ?", userId).Error
	if err != nil {
		return nil, err
	}
	return rets, nil
}

// 根据卖家 ID 来获取订单列表
func FindOrdersBySellerId(seller uint64) ([]*Order, error) {
	var rets []*Order
	err := db.Find(&rets).Where("seller_id = ?", seller).Error
	if err != nil {
		return nil, err
	}
	return rets, nil
}

// 新增订单
func InsertOrder(order *Order) error {
	err := db.Create(order).Error
	return err
}

//删除订单 由于这里是软删除，所以只是把 delete_at 设置了一个值，实际上还存在数据库中,但并不能用 gorm 查到
func DeleteOrderById(id uint64) error {
	err := db.Delete(Order{}, "id = ?", id).Error
	return err
}

//根据 order.ID 来更新其他相应的字段
func UpdateOrder(order *Order) error {
	var rec Order
	err := db.First(&rec, "id = ?", order.ID).Error
	if err != nil {
		return err
	}
	err = db.Model(&rec).Update(order).Error
	return err
}

// 从表单信息新增订单
func NewOrderFromForm(userId uint64, form *forms.NewOrderForm) ([]uint64, error) {
	// 开始事务
	tx := db.Begin()
	// 地址信息不能为空
	if form.Name == "" || form.Address == "" || form.Phone == "" {
		return nil, &resultError.AddressInfoErr
	}
	orders := []Order{}
	orderIdList := []uint64{}
	// 多个订单循环插入
	for _, item := range form.OrderPkgList {
		// 检查产品 ID 是否存在
		if tx.Find(&Product{BaseModel: BaseModel{ID: item.ProductID}}).RecordNotFound() {
			tx.Rollback()
			return nil, &resultError.ProductNotFound
		}
		// 检查套餐 ID 是否存在
		if tx.Find(&ProductPackage{BaseModel: BaseModel{ID: item.ProductPackageID}}).RecordNotFound() {
			tx.Rollback()
			return nil, &resultError.ProductPkgNotFound
		}
		// 根据表单创建新订单
		newOrder := Order{
			UserId:           userId,
			Name:             form.Name,
			Address:          form.Address,
			Phone:            form.Phone,
			SellerId:         item.SellerID,
			ProductId:        item.ProductID,
			ProductPackageId: item.ProductPackageID,
			UnitPrice:        item.Price,
			Nums:             item.Nums,
			TotalPrice:       item.Price * float64(item.Nums),
			Status:           enums.OrderStatus_Ordered,
		}
		// 向数据库中插入新订单
		err := tx.Create(&newOrder).Error
		if err != nil {
			tx.Rollback()
			return nil, &resultError.OrderCreateErr
		}
		// 将添加的订单信息添加到 slice 中
		orders = append(orders, newOrder)
		orderIdList = append(orderIdList, newOrder.ID)
	}
	// 提交事务
	tx.Commit()

	// 这里返回订单列表是为了之后付款时修改状态用
	//return orders, nil

	// 还是返回订单 Id 列表，方便后续结算
	return orderIdList, nil
}

// 查询订单列表的 SQL 字段
const sqlOrderListField = `
SELECT
	o.id,o.user_id,p.user_id AS seller_id,su.nickname AS seller_nickname,
	o.product_package_id,o.nums,o.unit_price,pkg.product_id,pkg.freight,
	p.name AS product_name,pkg.price,pkg.stock,pkg.image_url,pkg.description,
	o.created_at,o.status AS order_status,o.total_price,
	o.name,o.address,o.phone
FROM
	orders o
JOIN
	product_packages pkg ON o.product_package_id = pkg.id
JOIN
	products p ON p.id = o.product_id
JOIN 
	users su ON su.id = p.user_id
`

// 根据页码和用户信息来获取订单列表
const sqlGetOrderList = sqlOrderListField + `
WHERE
	o.deleted_at IS NULL  AND
	p.deleted_at IS NULL  AND
	pkg.deleted_at IS NULL AND
	o.user_id = (?)
ORDER BY
	o.created_at DESC
LIMIT ? OFFSET ?
`

// 根据页码和用户信息获取订单
func GetOrderList(pageForm forms.PageForm, userId uint64) (*resultModels.OrderList, error) {
	result := resultModels.OrderList{}
	var list []*resultModels.OrderListItem

	page, pageSize := 1, 5
	// 如果页码和每页数量大于 0
	if pageForm.Page > 0 && pageForm.PageSize > 0 {
		page = pageForm.Page
		pageSize = pageForm.PageSize
	}

	// 统计总数
	err := db.Table("orders").Where("user_id = (?) AND deleted_at IS  NULL", userId).Count(&result.Total).Error
	if err != nil {
		return nil, err
	}
	// 根据 SQL 字符串拼接查询订单相关信息列表
	err = db.Raw(sqlGetOrderList, userId, pageSize, (page-1)*pageSize).Scan(&list).Error
	result.Page = page
	result.OrderList = list
	return &result, err
}

const sqlGetOrderListInOrderIds = sqlOrderListField + `
WHERE
	o.deleted_at IS NULL  AND
	p.deleted_at IS NULL  AND
	pkg.deleted_at IS NULL AND
	o.user_id = (?) AND
    o.id IN (?)
ORDER BY
	o.created_at DESC
`

// 根据订单id列表和用户 Id 查询订单信息
func GetOrderListByOrderIds(orderIds []uint64, userId uint64) ([]*resultModels.OrderListItem, error) {
	var list []*resultModels.OrderListItem
	// 根据 SQL 字符串拼接查询订单相关信息列表
	err := db.Raw(sqlGetOrderListInOrderIds, userId, orderIds).Scan(&list).Error
	return list, err
}

// 订单支付，支付后需要 更新订单状态、产品增加众筹金额和人数、相应套餐减少库存增加支持人数
// 需要在事务中处理，错了一步就全部回退并返回错误信息
func PayOrderByOrderId(orderId uint64) error {
	// 开始事务
	tx := db.Begin()

	tx.Commit()
	return nil
}

// 更新订单状态,订单状态的定义在 enums 里面 以 OrderStatus_ 开头
func UpdateOrderStatusByOrderId(orderId uint64) error {

	return nil
}
