package resultModels

import (
	"time"
)

//SortOrder     int            `json:"sortOrder"`
//Position      int            `json:"position"`
//Status        int            `json:"status"`
//Remark        string         `json:"remark"`
type HomeResult struct {
	Name            string           `json:"name"`
	Type            int              `json:"type"`
	LimitNum        int              `json:"limit_num"`
	ProductContents []ProductContent `json:"product_contents"`
}

type ProductContent struct {
	ID uint64 `json:"product_id"`
	//产品名
	Name string `json:"name"`
	//大图
	BigImg string `json:"big_img"`
	//列表小图
	SmallImg string `json:"small_img"`
	//产品类型
	ProductType int `json:"product_type"`
	//当前筹集金额
	CurrentPrice float64 `json:"current_price"`
	//目标筹集金额
	TargetPrice float64 `json:"target_price"`
	//支持人数
	Backers int `json:"backers"`
	//当前时间
	CurrentTime time.Time `json:"current_time"`
	//截止时间
	EndTime time.Time `json:"end_time"`
}