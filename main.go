package main

import (
	"CarSys/contronler"
	"net/http"
)

func main() {
	//处理静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	//直接去html页面
	//http.HandleFunc("/main", contronler.IndexHandler)
	http.HandleFunc("/main", contronler.GetPageCarsByPrice)
	//去登录
	http.HandleFunc("/login", contronler.Login)
	//去注册
	http.HandleFunc("/regist", contronler.Regist)
	//去注销
	http.HandleFunc("/logout", contronler.Logout)
	//获取所有车辆
	//http.HandleFunc("/getCars", contronler.GetCars)
	//获取带分页的车辆信息
	http.HandleFunc("/getPageCars", contronler.GetPageCars)
	//获取根据价格查到的带分页的车辆信息
	http.HandleFunc("/getPageCarsByPrice", contronler.GetPageCarsByPrice)
	//添加车辆
	http.HandleFunc("/addCar", contronler.AddCar)
	//删除车辆
	http.HandleFunc("/deleteCar", contronler.DeleteCar)
	//去修改车辆的页面
	http.HandleFunc("/toUpdateCarPage", contronler.ToUpdateCarPage)
	//修改车辆
	http.HandleFunc("/updateCar", contronler.UpdateCar)
	//选择车辆租借并加载到订单中
	http.HandleFunc("/addCar2Cart", contronler.AddCar2Cart)
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

	http.ListenAndServe(":8080", nil)

}
