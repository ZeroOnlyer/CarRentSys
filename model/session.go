package model

//Session 定义session结构
type Session struct {
	SessionID string //session的id
	Name      string //用户的姓名
	UserID    int    //用户的id
	Cart      *Cart  //用户的订单
	OrderID   string //付款单号
	Orders    []*Order
}
