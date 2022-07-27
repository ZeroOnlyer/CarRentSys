package model

//Order 付款单
type Order struct {
	OrderID     string  //单号
	CreateTime  string  //下单时间
	TotalCount  int64   //最终付款的宠物服务数量
	TotalAmount float64 //最终付款
	State       int64   //订单状态
	UserID      int64   //付款单单所属的用户
}
