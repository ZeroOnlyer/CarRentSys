package model

//CartItem 定义订单项的结构
type CartItem struct {
	CartItemID int64   //订单项的id
	Car        *Car    //订单项里的车辆信息
	Count      int64   //订单项里的车辆数量
	Amount     float64 //订单项里的车辆总租金
	CartID     string  //当前订单项属于哪一个订单
}

//GetAmout 获取订单项里的租金总计
func (cartitem *CartItem) GetAmout() float64 {
	//获取当前订单项里的车辆的价格
	price := cartitem.Car.Price
	return float64(cartitem.Count) * price
}
