package model

//Cart 购物车
type Cart struct {
	CartID      string      //订单的id
	CartItems   []*CartItem //订单中所有的订单项
	TotalCount  int64       //订单中车辆的总数量，通过计算得到
	TotalAmount float64     //订单中车辆的总租金，通过计算得到
	UserID      int         //当前订单所属的用户
	Name        string      //当前订单所属用户的姓名
}

//GetTotalCount 获取订单中车辆的总数量
func (cart *Cart) GetTotalCount() int64 {
	var totalCount int64
	//遍历订单中的每一个订单项
	for _, v := range cart.CartItems {
		totalCount = totalCount + v.Count
	}
	return totalCount
}

//GetTotalAmount 获取订单中车辆的总租金
func (cart *Cart) GetTotalAmount() float64 {
	var totalAmount float64
	//遍历订单中的每一个订单项
	for _, v := range cart.CartItems {
		totalAmount = totalAmount + v.GetAmout()
	}
	return totalAmount
}
