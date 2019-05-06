package models

import (
	"errors"
	"funding/forms"
	"funding/resultModels"
	"time"
)

//产品
type Product struct {
	BaseModel
	Name            string           `json:"name"`                           //产品名
	BigImg          string           `json:"big_img"`                        //顶部大图
	SmallImg        string           `json:"small_img"`                      //列表小图
	UserId          string           `json:"user_id"`                        //发布者ID
	ProductType     int              `json:"product_type"`                   //产品类型
	CurrentPrice    float64          `json:"current_price"`                  //当前筹集金额
	TargetPrice     float64          `json:"target_price"`                   //目标筹集金额
	VerifyStatus    int              `json:"verify_status" gorm:"default:2"` //审核状态(0:未通过 1:已通过 2:待审核(默认))
	Backers         int              `json:"backers"`                        //支持人数
	EndTime         time.Time        `json:"end_time"`                       //截止时间
	DetailHtml      string           `json:"detail_html"`                    //介绍页详情 Html
	ProductPackages []ProductPackage `json:"product_packages"`               //商品套餐
}

//产品套餐
type ProductPackage struct {
	BaseModel
	ProductId   string  `json:"product_id"`   //对应产品 Id
	Description string  `json:"description"`  //套餐描述
	ImageUrl    string  `json:"image_url"`    //图片链接
	Price       float64 `json:"price"`        //套餐价格
	Stock       int64   `json:"stock"`        //剩余库存
	Total       int64   `json:"total"`        //套餐总数
	Backers     int     `json:"backers"`      //支持人数
	Freight     float64 `json:"freight"`      //运费
	DeliveryDay int64   `json:"delivery_day"` //发货时间 (众筹成功后多少天内)
}

func init() {
	//db.AutoMigrate(&Product{}, &ProductPackage{})
}

// 根据 分页 和 产品类型(0 为全部) 获取产品
func GetProductsByPageAndType(page int, pageSize int, productType int) ([]*Product, error) {
	if page <= 0 || pageSize <= 0 {
		return nil, errors.New("参数错误")
	}
	var results []*Product
	//分页限制
	pDb := db.Limit(pageSize).Offset((page - 1) * pageSize)
	//类型为 0 时不限制类型
	if productType != 0 {
		pDb = pDb.Where("product_type = ?", productType)
	}
	//倒序查询
	pDb = pDb.Order("id desc")

	err := pDb.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

// 获取全部产品
func GetAllProduct() ([]*Product, error) {
	var results []*Product
	err := db.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func GetProductWithPkg(productId uint64) (*Product, error) {
	var result Product
	err := db.Preload("ProductPackages").First(&result, productId).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetProductPackages(productId uint64) ([]*ProductPackage, error) {
	var result []*ProductPackage
	err := db.Where("product_id = ?", productId).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 根据页码等信息获取产品列表
func GetProductList(form forms.ProductListForm) (*resultModels.ProductList, error) {
	result := resultModels.ProductList{}
	var list []*resultModels.ProductContent
	cDb := db
	// 如果传入类型不为 0 （传入了 Type)，则根据 Type 查询
	if form.Type != 0 {
		cDb = cDb.Where("product_type = ?", form.Type)
	}
	page, pageSize := 1, 10
	// 如果页码和每页数量大于 0
	if form.Page > 0 && form.PageSize > 0 {
		page = form.Page
		pageSize = form.PageSize
	}

	// 只查询指定字段
	cDb = cDb.Select(resultModels.ProductContentField)
	// 未软删除的行
	cDb = cDb.Where("deleted_at IS NULL").Table("products")
	// 排序
	cDb = cDb.Order("created_at DESC,end_time DESC")
	// 统计总数
	err := cDb.Count(&result.Total).Error
	if err != nil {
		return nil, err
	}
	// 调整 Offset(偏移量，控制页数),和 Limit (数量限制，控制每页数量）
	cDb = cDb.Offset((page - 1) * pageSize).Limit(pageSize)
	err = cDb.Scan(&list).Error
	result.Page = form.Page
	result.ProductContents = list
	return &result, err
}
