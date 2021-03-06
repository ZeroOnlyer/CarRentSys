package contronler

import (
	"CarSys/dao"
	"CarSys/model"
	"CarSys/utils"
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
	//创建Order
	order := &model.Order{
		OrderID:     orderID,
		CreateTime:  time.Now(),
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
			Name:    v.Car.Name,
			Price:   v.Car.Price,
			OrderID: orderID,
		}
		//保存每一项到数据库中
		dao.AddOrderItem(orderItem)
		//更新当前订单项的车辆数量
		car := v.Car
		car.Num = car.Num - int(v.Count)
		//更新车辆的信息
		dao.UpdateCar(car)
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
