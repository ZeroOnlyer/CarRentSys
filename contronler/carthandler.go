package contronler

import (
	"PetHome/dao"
	"PetHome/model"
	"PetHome/utils"
	"net/http"
	"strconv"
	"text/template"
)

//AddService2Cart 购买宠物服务并加载到订单中
func AddService2Cart(w http.ResponseWriter, r *http.Request) {
	//判断是否登录
	flag, session, _ := dao.IsLogin(r)
	if flag {
		//已经登录了
		//获取要购买的宠物服务的id
		serviceID := r.FormValue("serviceID")
		//根据宠物服务的id获取宠物服务信息
		service, _ := dao.GetServiceByID(serviceID)

		//获取用户的id
		userID := session.UserID
		//判断数据库中是否已经有当前用户的订单
		cart, _ := dao.GetCartByUserID(userID)
		if cart != nil {
			//当前用户已经有订单了,需要判断订单中是否已有要购买的这项宠物服务
			cartItem, _ := dao.GetCartItemByCarIDAndCartID(serviceID, cart.CartID)
			if cartItem != nil {
				//订单中已经有这项宠物服务，只需要把这项宠物服务所对应的订单项的数量加1，比如购买两次同样的宠物服务
				//1.获取订单中所有的订单项
				cartItems := cart.CartItems
				//2.遍历得到每一个订单项
				for _, v := range cartItems {
					//3.找到当前的订单项
					if cartItem.Service.ID == v.Service.ID {
						//将订单项中的宠物服务的数量加1
						v.Count = v.Count + 1
						//将数据库中的该订单项的数量和金额更新
						dao.UpdateCarCount(v)
					}
				}
			} else {
				//订单中没有这项宠物服务，就创建一个订单项并添加到数据库中
				//创建订单项
				cartItem := &model.CartItem{
					Service: service,
					Count:   1,
					CartID:  cart.CartID,
				}
				//将订单项添加到切片中
				cart.CartItems = append(cart.CartItems, cartItem)
				//将新创建的订单项添加到数据库中中
				dao.AddCartItem(cartItem)

			}
			//无论订单中是否有这项要购买的宠物服务，都要更新订单中的宠物服务的总数量和总租金
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
				Service: service,
				Count:   1,
				CartID:  cartID,
			}
			//3.将订单项添加到切片中
			cartItems = append(cartItems, cartItem)
			//4.将切片设置到订单中
			cart.CartItems = cartItems
			//5.将订单保存到数据库中
			dao.AddCart(cart)
		}
		//service.Num = service.Num - 1
		dao.UpdateService(service)
		w.Write([]byte("您刚刚添加了" + service.Name + "这项服务到购物车！"))
	} else {
		//没有登录
		w.Write([]byte("您还没有登录！不能购买哟！请先登录！"))
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
	//更新订单中的宠物服务的总数量和总租金
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
	//获取用户输入的宠物服务的数量
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
			//将当前订单项目的宠物服务数量设置为用户输入的值
			v.Count = iCarCount
			//更新数据库中该订单项的宠物服务的数量和总租金
			dao.UpdateCarCount(v)
		}
	}
	//更新订单中的订单项的总数量和总租金
	dao.UpdateCart(cart)
	//再次查询订单信息
	GetCartInFo(w, r)
}
