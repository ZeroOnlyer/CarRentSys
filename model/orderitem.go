package model

//OrderItem 结构
type OrderItem struct {
	OrderItemID int64   //付款单每一项的id
	Count       int64   //付款单每一项的宠物服务数量
	Amount      float64 //付款单每一项的宠物服务租金和
	Name        string  //付款单宠物服务名称
	Price       float64 //付款单宠物服务价格
	OrderID     string  //付款单项所属的付款单
}
