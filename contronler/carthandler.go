package contronler

import (
	"CarSys/dao"
	"CarSys/model"
	"CarSys/utils"
	"net/http"
	"strconv"
	"text/template"
)

//AddCar2Cart 租借车辆并加载到订单中
func AddCar2Cart(w http.ResponseWriter, r *http.Request) {
	//判断是否登录
	flag, session, _ := dao.IsLogin(r)
	if flag {
		//已经登录了
		//获取要租借的车辆的id
		carID := r.FormValue("carID")
		//根据车辆的id获取车辆信息
		car, _ := dao.GetCarByID(carID)

		//获取用户的id
		userID := session.UserID
		//判断数据库中是否已经有当前用户的订单
		cart, _ := dao.GetCartByUserID(userID)
		if cart != nil {
			//当前用户已经有订单了,需要判断订单中是否已有要租借的这辆汽车
			cartItem, _ := dao.GetCartItemByCarIDAndCartID(carID, cart.CartID)
			if cartItem != nil {
				//订单中已经有这辆车，只需要把这辆车所对应的订单项的数量加1，比如租借两次同样的汽车
				//1.获取订单中所有的订单项
				cartItems := cart.CartItems
				//2.遍历得到每一个订单项
				for _, v := range cartItems {
					//3.找到当前的订单项
					if cartItem.Car.ID == v.Car.ID {
						//将订单项中的车辆的数量加1
						v.Count = v.Count + 1
						//将数据库中的该订单项的数量和金额更新
						dao.UpdateCarCount(v)
					}
				}
			} else {
				//订单中没有这辆车，就创建一个订单项并添加到数据库中
				//创建订单项
				cartItem := &model.CartItem{
					Car:    car,
					Count:  1,
					CartID: cart.CartID,
				}
				//将订单项添加到切片中
				cart.CartItems = append(cart.CartItems, cartItem)
				//将新创建的订单项添加到数据库中中
				dao.AddCartItem(cartItem)

			}
			//无论订单中是否有这辆要租借的车，都要更新订单中的车辆的总数量和总租金
			dao.UpdateCart(cart)
		} else {
			//当前用户还没有订单，就创建一个订单并添加到数据库中
			//1.创建订单
			//生成订单的id
			cartID := utils.CreateUUID()
			cart := &model.Cart{
				CartID: cartID,
				UserID: userID,
			}
			//2.创建订单项
			var cartItems []*model.CartItem
			cartItem := &model.CartItem{
				Car:    car,
				Count:  1,
				CartID: cartID,
			}
			//3.将订单项添加到切片中
			cartItems = append(cartItems, cartItem)
			//4.将切片设置到订单中
			cart.CartItems = cartItems
			//5.将订单保存到数据库中
			dao.AddCart(cart)
		}
		w.Write([]byte("您刚刚租借了" + car.Name + "这辆车！"))
	} else {
		//没有登录
		w.Write([]byte("您还没有登录！不能租借！请先登录！"))
	}
}

//GetCartInFo 根据用户id获取订单信息
func GetCartInFo(w http.ResponseWriter, r *http.Request) {
	_, session, _ := dao.IsLogin(r)
	//获取用户id
	userID := session.UserID
	//根据用户的id从数据库中获取对应的订单
	cart, _ := dao.GetCartByUserID(userID)
	if cart != nil {
		//设置用户的姓名
		session.Cart = cart
		cart.Name = session.Name
		//解析模板文件
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		t.Execute(w, session)
	} else {
		//该用户还没有订单
		//解析模板文件
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		t.Execute(w, session)
	}

}

//DeleteCart 清空订单
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	//获取要删除的订单的id
	cartID := r.FormValue("cartId")
	//清空订单
	dao.DeleteCartByCartID(cartID)
	//再次查询订单信息
	GetCartInFo(w, r)
}

//DeleteCartItem 删除订单项
func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	//获取要删除的订单项的id
	cartItemID := r.FormValue("cartItemId")
	//将订单项的id转换为int64
	icartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	//获取session
	_, session, _ := dao.IsLogin(r)
	//获取用户的id
	userID := session.UserID
	//获取该用户的订单
	cart, _ := dao.GetCartByUserID(userID)
	//获取购物车中的订单项
	cartItems := cart.CartItems
	//遍历得到每一个订单项
	for k, v := range cartItems {
		//寻找要删除的订单项
		if v.CartItemID == icartItemID {
			//找到要删除的订单项
			//将当前订单项从切片中移除
			cartItems = append(cartItems[:k], cartItems[k+1:]...)
			//将删除订单项之后的切片再次赋给订单中的切片
			cart.CartItems = cartItems
			//将订单项从数据库中删除
			dao.DeleteCartItemByID(cartItemID)
		}
	}
	//更新订单中的车辆的总数量和总租金
	dao.UpdateCart(cart)
	//再次查询订单信息
	GetCartInFo(w, r)
}

//UpdateCartItem 更新订单项
func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	//获取要更新的订单项的id
	cartItemID := r.FormValue("cartItemId")
	//将订单项的id转换为int64
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	//获取用户输入的车辆的数量
	carCount := r.FormValue("carCount")
	//将数量转换为int64
	iCarCount, _ := strconv.ParseInt(carCount, 10, 64)
	//获取session
	_, session, _ := dao.IsLogin(r)
	//获取用户的id
	userID := session.UserID
	//获取该用户的订单
	cart, _ := dao.GetCartByUserID(userID)
	//获取订单项中所有的订单项
	cartItems := cart.CartItems
	//遍历得到每一个订单项
	for _, v := range cartItems {
		//寻找要更新的订单项
		if v.CartItemID == iCartItemID {
			//找到要更新的订单项
			//将当前订单项目的车辆数量设置为用户输入的值
			v.Count = iCarCount
			//更新数据库中该订单项的车辆的数量和总租金
			dao.UpdateCarCount(v)
		}
	}
	//更新订单中的订单项的总数量和总租金
	dao.UpdateCart(cart)
	//再次查询订单信息
	GetCartInFo(w, r)
}
