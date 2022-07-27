package main

import (
	"PetHome/contronler"
	"net/http"
)

func main() {
	//处理静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	http.HandleFunc("/", contronler.GetPageServicesByPrice)
	//去登录
	http.HandleFunc("/login", contronler.Login)
	//去注册
	http.HandleFunc("/regist", contronler.Regist)
	//去注销
	http.HandleFunc("/logout", contronler.Logout)
	//获取带分页的宠物服务信息
	http.HandleFunc("/getPageServices", contronler.GetPageServices)
	//获取根据价格查到的带分页的宠物服务信息
	http.HandleFunc("/getPageServicesByPrice", contronler.GetPageServicesByPrice)
	//添加宠物服务
	http.HandleFunc("/addService", contronler.AddService)
	//删除宠物服务
	http.HandleFunc("/deleteService", contronler.DeleteService)
	//去修改宠物服务的页面
	http.HandleFunc("/toUpdateServicePage", contronler.ToUpdateServicePage)
	//修改宠物服务
	http.HandleFunc("/updateService", contronler.UpdateService)
	//选择宠物服务购买并加载到订单中
	http.HandleFunc("/addService2Cart", contronler.AddService2Cart)
	//获取订单信息
	http.HandleFunc("/getCartInFo", contronler.GetCartInFo)
	//清空订单
	http.HandleFunc("/deleteCart", contronler.DeleteCart)
	//删除订单项
	http.HandleFunc("/deleteCartItem", contronler.DeleteCartItem)
	//更新订单项
	http.HandleFunc("/updateCartItem", contronler.UpdateCartItem)
	//去付款
	http.HandleFunc("/checkout", contronler.Checkout)
	//获取所有订单
	http.HandleFunc("/getOrders", contronler.GetOrders)
	//获取我的订单
	http.HandleFunc("/getMyOrder", contronler.GetMyOrders)
	//获取订单详情，即订单所对应的所有的订单项
	http.HandleFunc("/getOrderInfo", contronler.GetOrderInfo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

}
