package contronler

import (
	"CarSys/dao"
	"CarSys/model"
	"html/template"
	"net/http"
	"strconv"
)

//IndexHandler 去首页
/*func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//获取分页
	pageNow := r.FormValue("pageNow")
	if pageNow == "" {
		pageNow = "1"
	}
	//调用cardao中获取所有分页车辆的函数
	page, _ := dao.GetPageCars(pageNow)
	//解析模板
	t := template.Must(template.ParseFiles("views/index.html"))
	//执行
	t.Execute(w, page)
}*/

//GetPageCarsByPrice 获取分页的根据价格查询到的所有车辆
func GetPageCarsByPrice(w http.ResponseWriter, r *http.Request) {
	//获取分页
	pageNow := r.FormValue("pageNow")
	//获取价格
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")
	if pageNow == "" {
		pageNow = "1"
	}
	var page *model.Page
	if minPrice == "" && maxPrice == "" {
		//调用cardao中获取所有分页车辆的函数
		page, _ = dao.GetPageCars(pageNow)
	} else {
		//调用cardao中获取所有根据价格查询到的分页车辆的函数
		page, _ = dao.GetPageCarsByPrice(pageNow, minPrice, maxPrice)
		//将价格范围设置到page中
		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
	}
	//调用sessiondao中的IsLogin函数判断是否登录
	flag, session, superID := dao.IsLogin(r)
	if flag {
		//已经登录了，设置page结构体中的IsLogin和Name和SuperID
		page.IsLogin = true
		page.Name = session.Name
		if superID == 1 {
			page.SuperID = true
		} else {
			page.SuperID = false
		}

	}

	//解析模板
	t := template.Must(template.ParseFiles("views/index.html"))
	//执行
	t.Execute(w, page)
}

/*//GetCars 获取所有车辆
func GetCars(w http.ResponseWriter, r *http.Request) {
	//调用cardao中获取所有车辆的函数
	cars, _ := dao.GetCars()
	t := template.Must(template.ParseFiles("views/pages/manager/car_manager.html"))
	t.Execute(w, cars)
}*/

//GetPageCars 获取分页的所有车辆
func GetPageCars(w http.ResponseWriter, r *http.Request) {
	//获取分页
	pageNow := r.FormValue("pageNow")
	if pageNow == "" {
		pageNow = "1"
	}
	//调用cardao中获取所有分页车辆的函数
	page, _ := dao.GetPageCars(pageNow)
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/manager/car_manager.html"))
	//执行
	t.Execute(w, page)
}

//AddCar 添加一辆汽车
func AddCar(w http.ResponseWriter, r *http.Request) {
	//获取要添加的车辆信息
	name := r.PostFormValue("name")
	price := r.PostFormValue("price")
	num := r.PostFormValue("num")
	fprice, _ := strconv.ParseFloat(price, 64)
	inum, _ := strconv.ParseInt(num, 10, 0)
	//创建Car
	car := &model.Car{
		Name:  name,
		Price: fprice,
		Num:   int(inum),
	}
	//调用cardao中添加车辆的函数
	dao.AddCar(car)
	//调用处理器函数GetCars再查一次数据库看是否添加成功
	GetPageCars(w, r)
}

//DeleteCar 删除车辆
func DeleteCar(w http.ResponseWriter, r *http.Request) {
	//获取要删除的车辆的ID
	carID := r.FormValue("carID")
	//调用cardao中删除车辆的函数
	dao.DeleteCar(carID)
	//调用处理器函数GetCars再查一次数据库看是否删除成功
	GetPageCars(w, r)
}

//ToUpdateCarPage 去修改车辆的一个页面
func ToUpdateCarPage(w http.ResponseWriter, r *http.Request) {
	//获取要修改车辆的ID
	carID := r.FormValue("carID")
	//调用cardao中根据ID获取车辆的函数
	car, _ := dao.GetCarByID(carID)
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/manager/car_modify.html"))
	//执行
	t.Execute(w, car)

}

//UpdateCar 修改车辆
func UpdateCar(w http.ResponseWriter, r *http.Request) {
	//获取要修改的车辆信息
	carID := r.PostFormValue("carID")
	name := r.PostFormValue("name")
	price := r.PostFormValue("price")
	num := r.PostFormValue("num")
	icarID, _ := strconv.ParseInt(carID, 10, 0)
	fprice, _ := strconv.ParseFloat(price, 64)
	inum, _ := strconv.ParseInt(num, 10, 0)
	//创建Car
	car := &model.Car{
		ID:    int(icarID),
		Name:  name,
		Price: fprice,
		Num:   int(inum),
	}
	//调用cardao中修改车辆的函数
	dao.UpdateCar(car)
	//调用处理器函数GetCars再查一次数据库看是否修改成功
	GetPageCars(w, r)
}
