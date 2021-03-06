package dao

import (
	"CarSys/model"
	"fmt"
	"testing"
	"time"
)

/*func TestCustomer(t *testing.T) {
	//fmt.Println("测试userdao中的函数")
	//t.Run("验证账号和密码：", testLogin)
	//t.Run("验证账号：", testRegist)
	//t.Run("保存用户：", testSave)
}
*/
func testLogin(t *testing.T) {
	cu, _ := CheckUserNameAndPassword("admin", "admin")
	fmt.Println("获取的用户信息是：", cu)
}

func testRegist(t *testing.T) {
	cu, _ := CheckUserName("theshy")
	fmt.Println("获取用户信息是：", cu)
}

func testSave(t *testing.T) {
	SaveUser("李四", 48, "951531", "theshy", "8485")
}

/*func TestCar(t *testing.T) {
	fmt.Println("测试Cardao中的相关函数")
	//t.Run("测试获取所有车辆", testGetCars)
	//t.Run("测试添加一辆汽车", testAddCar)
	//t.Run("测试删除一辆汽车", testDeleteCar)
	//t.Run("测试查询一辆汽车", testGetCar)
	//t.Run("测试修改一辆汽车：", testUpdateCar)
	//t.Run("测试分页：", testGetPageCars)
	t.Run("测试分页根据价格查询到的车辆：", testGetPageCarsByPrice)
}
*/
func testGetCars(t *testing.T) {
	cars, _ := GetCars()
	for k, v := range cars {
		fmt.Printf("第%v辆汽车的信息是：%v\n", k+1, v)
	}
}

func testAddCar(t *testing.T) {
	car := &model.Car{
		Name:  "宾利",
		Price: 800,
		Num:   15,
	}
	//调用添加车辆的函数
	AddCar(car)
}

func testDeleteCar(t *testing.T) {
	//调用删除车辆的函数
	DeleteCar("10")
}

func testGetCar(t *testing.T) {
	//调用查询一辆车辆的函数
	car, _ := GetCarByID("7")
	fmt.Println("获取的车辆信息是：", car)
}

func testUpdateCar(t *testing.T) {
	car := &model.Car{
		ID:    6,
		Name:  "玛莎拉蒂",
		Price: 10000,
		Num:   2,
	}
	//调用修改车辆的函数
	UpdateCar(car)
}

func testGetPageCars(t *testing.T) {
	page, _ := GetPageCars("4")
	fmt.Println("当前页是：", page.PageNow)
	fmt.Println("总页数是：", page.PageSum)
	fmt.Println("总记录是：", page.PageFinal)
	fmt.Println("当前页的汽车有：")
	for _, v := range page.Cars {
		fmt.Println("车辆：", v)
	}
}

func testGetPageCarsByPrice(t *testing.T) {
	page, _ := GetPageCarsByPrice("1", "500", "1000")
	fmt.Println("当前页是：", page.PageNow)
	fmt.Println("总页数是：", page.PageSum)
	fmt.Println("总记录是：", page.PageFinal)
	fmt.Println("当前页的汽车有：")
	for _, v := range page.Cars {
		fmt.Println("车辆：", v)
	}
}

/*func TestSession(t *testing.T) {
	fmt.Println("测试session的相关函数")
	//t.Run("测试添加session", testAddSession)
	//t.Run("测试删除session", testDeleteSession)
	t.Run("测试获取session", testGetSession)
}*/

func testAddSession(t *testing.T) {
	sess := &model.Session{
		SessionID: "13838381438",
		Name:      "张三",
		UserID:    6,
	}
	AddSession(sess)

}

func testDeleteSession(t *testing.T) {
	DeleteSession("13838381438")
}

func testGetSession(t *testing.T) {
	sess, _ := GetSessionByID("9566c74d-1003-4c4d-7bbb-0407D1E2C649")
	fmt.Println("session的信息是：", sess)
}

/*func TestCart(t *testing.T) {
	fmt.Println("测试订单的相关函数")
	//t.Run("测试添加订单", testAddCart)
	//t.Run("测试根据车辆id获取购物项", testGetCartItemByCarID)
	//t.Run("测试根据订单的id获取所有的订单项", testGetCartItemsByCarID)
	//t.Run("测试根据用户的id获取对应的订单", testGetCartByUserID)
	//t.Run("测试根据车辆id和订单id和车辆数量来更新车辆的数量", testUpdateCarCount)
	//t.Run("测试根据订单的id删除订单和订单项", testDeleteCartByCartID)
	t.Run("测试删除订单项", testDeleteCartItemByID)
}*/

func testAddCart(t *testing.T) {
	//设置要租借的第一辆车
	car := &model.Car{
		ID:    1,
		Price: 1000,
	}
	//设置要租借的第二辆车
	car2 := &model.Car{
		ID:    2,
		Price: 3000,
	}
	//创建一个购物项切片
	var cartItems []*model.CartItem
	//创建这两个订单项
	cartItem := &model.CartItem{
		Car:    car,
		Count:  10,
		CartID: "66667777",
	}
	cartItems = append(cartItems, cartItem)
	cartItem2 := &model.CartItem{
		Car:    car2,
		Count:  2,
		CartID: "66667777",
	}
	cartItems = append(cartItems, cartItem2)

	//创建订单
	cart := &model.Cart{
		CartID:    "66667777",
		CartItems: cartItems,
		UserID:    5,
	}
	//将订单插入数据库中
	AddCart(cart)
}

func testGetCartItemByCarID(T *testing.T) {
	cartItem, _ := GetCartItemByCarIDAndCartID("2", "66667777")
	fmt.Println("车辆id为1的对应的订单项是：", cartItem)

}
func testGetCartItemsByCarID(T *testing.T) {
	cartItems, _ := GetCartItemsByCartID("66667777")
	for k, v := range cartItems {
		fmt.Printf("第%v个购物项是：%v\n", k+1, v)
	}
}

func testGetCartByUserID(t *testing.T) {
	cart, _ := GetCartByUserID(5)
	fmt.Println("id为5的用户的订单是：", cart)
}

func testUpdateCarCount(t *testing.T) {
	//UpdateCarCount(5, 1, "66667777")
}

func testDeleteCartByCartID(t *testing.T) {
	DeleteCartByCartID("52fdfc07-2182-454f-563f-5F0F9A621D72")
}

func testDeleteCartItemByID(t *testing.T) {
	DeleteCartItemByID("37")
}

func TestOrder(t *testing.T) {
	fmt.Println("测试相关函数")
	t.Run("测试添加付款单和付款单项", testAddOrder)
}

func testAddOrder(t *testing.T) {

	//创建付款单
	order := &model.Order{
		OrderID:     "1384544545",
		CreateTime:  time.Now(),
		TotalCount:  2,
		TotalAmount: 400,
		State:       0,
		UserID:      1,
	}
	//创建付款单项目
	orderItem := &model.OrderItem{
		Count:   1,
		Amount:  300,
		Name:    "长安汽车",
		Price:   300,
		OrderID: "1384544545",
	}
	orderItem2 := &model.OrderItem{
		Count:   1,
		Amount:  100,
		Name:    "大众",
		Price:   100,
		OrderID: "1384544545",
	}
	//保存付款单
	AddOrder(order)
	AddOrderItem(orderItem)
	AddOrderItem(orderItem2)
}
