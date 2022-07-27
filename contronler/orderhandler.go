package contronler

import (
	"PetHome/dao"
	"PetHome/model"
	"PetHome/utils"
	"html/template"
	"net/http"
	"time"
)

//Checkout 去付款
func Checkout(w http.ResponseWriter, r *http.Request) {
	//获取session
	_, session, _ := dao.IsLogin(r)
	//获取用户的id
	userID := session.UserID
	//获取订单
	cart, _ := dao.GetCartByUserID(userID)
	//生成付款单号
	orderID := utils.CreateUUID()
	//创建生成订单的时间
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	//创建Order
	order := &model.Order{
		OrderID:     orderID,
		CreateTime:  timeStr,
		TotalCount:  cart.TotalCount,
		TotalAmount: cart.GetTotalAmount(),
		State:       0,
		UserID:      int64(userID),
	}
	//保存到数据库
	dao.AddOrder(order)
	//保存付款单项
	//获取原来订单里所有的订单项
	cartItems := cart.CartItems
	//遍历得到每一个订单项
	for _, v := range cartItems {
		//创建OrderItem
		orderItem := &model.OrderItem{
			Count:   v.Count,
			Amount:  v.Amount,
			Name:    v.Service.Name,
			Price:   v.Service.Price,
			OrderID: orderID,
		}
		//保存每一项到数据库中
		dao.AddOrderItem(orderItem)
		//更新当前订单项的宠物服务数量
		service := v.Service
		service.Num = service.Num - int(v.Count)
		//更新宠物服务的信息
		dao.UpdateService(service)
	}

	//付款后清空原订单
	dao.DeleteCartByCartID(cart.CartID)
	//将付款单号设置到session中
	session.OrderID = orderID
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
	//执行
	t.Execute(w, session)
}

//GetOrders 获取所有订单
func GetOrders(w http.ResponseWriter, r *http.Request) {
	//调用dao中获取所有订单的函数
	orders, _ := dao.GetOrders()
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/order/order_manager.html"))
	//执行
	t.Execute(w, orders)
}

//GetMyOrders 获取我的订单
func GetMyOrders(w http.ResponseWriter, r *http.Request) {
	//获取session
	_, session, _ := dao.IsLogin(r)
	//获取用户的id
	userID := session.UserID
	//调用dao中获取用户的所有订单的函数
	orders, _ := dao.GetMyOrders(userID)
	//将订单设置到session中
	session.Orders = orders
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/order/order.html"))
	//执行
	t.Execute(w, session)
}

//GetOrderInfo 获取订单对应的订单项
func GetOrderInfo(w http.ResponseWriter, r *http.Request) {
	//获取订单号
	orderID := r.FormValue("orderId")
	//根据订单号调用dao中获取所有订单项的函数
	orderItems, _ := dao.GetOrderItemsByOrderID(orderID)
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/order/order_info.html"))
	//执行
	t.Execute(w, orderItems)
}
