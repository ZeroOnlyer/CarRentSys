package model

import "time"

//Order 付款单
type Order struct {
	OrderID     string    //单号
	CreateTime  time.Time //下单时间
	TotalCount  int64     //最终付款的车辆数量
	TotalAmount float64   //最终付款
	State       int64     //1已出库   0未出库  2租借完成
	UserID      int64     //付款单单所属的用户
}
