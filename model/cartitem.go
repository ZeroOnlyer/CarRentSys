package model

//CartItem 定义订单项的结构
type CartItem struct {
	CartItemID int64    //订单项的id
	Service    *Service //订单项里的宠物服务信息
	Count      int64    //订单项里的宠物服务数量
	Amount     float64  //订单项里的宠物服务总租金
	CartID     string   //当前订单项属于哪一个订单
}

//GetAmout 获取订单项里的租金总计
func (cartitem *CartItem) GetAmout() float64 {
	//获取当前订单项里的宠物服务的价格
	price := cartitem.Service.Price
	return float64(cartitem.Count) * price
}
