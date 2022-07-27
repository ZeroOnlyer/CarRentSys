package contronler

import (
	"PetHome/dao"
	"PetHome/model"
	"html/template"
	"net/http"
	"strconv"
)

//GetPageServicesByPrice 获取分页的根据价格查询到的所有宠物服务
func GetPageServicesByPrice(w http.ResponseWriter, r *http.Request) {
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
		//调用servicedao中获取所有分页宠物服务的函数
		page, _ = dao.GetPageServices(pageNow)
	} else {
		//调用servicedao中获取所有根据价格查询到的分页宠物服务的函数
		page, _ = dao.GetPageServicesByPrice(pageNow, minPrice, maxPrice)
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
	err := t.Execute(w, page)
	if err != nil {
		return
	}
}

//GetPageServices 获取分页的所有宠物服务
func GetPageServices(w http.ResponseWriter, r *http.Request) {
	//获取分页
	pageNow := r.FormValue("pageNow")
	if pageNow == "" {
		pageNow = "1"
	}
	//调用servicedao中获取所有分页宠物服务的函数
	page, _ := dao.GetPageServices(pageNow)
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/manager/service_manager.html"))
	//执行
	t.Execute(w, page)
}

//AddService 添加一项宠物服务
func AddService(w http.ResponseWriter, r *http.Request) {
	//获取要添加的宠物服务信息
	name := r.PostFormValue("name")
	price := r.PostFormValue("price")
	num := r.PostFormValue("num")
	imgPath := r.PostFormValue("imgPath")
	fprice, _ := strconv.ParseFloat(price, 64)
	inum, _ := strconv.ParseInt(num, 10, 0)
	//创建Service
	service := &model.Service{
		Name:    name,
		Price:   fprice,
		Num:     int(inum),
		ImgPath: imgPath,
	}
	//调用servicedao中添加宠物服务的函数
	dao.AddService(service)
	//调用处理器函数GetCars再查一次数据库看是否添加成功
	GetPageServices(w, r)
}

//DeleteService 删除宠物服务
func DeleteService(w http.ResponseWriter, r *http.Request) {
	//获取要删除的宠物服务的ID
	serviceID := r.FormValue("serviceID")
	//调用cardao中删除宠物服务的函数
	dao.DeleteService(serviceID)
	//调用处理器函数GetCars再查一次数据库看是否删除成功
	GetPageServices(w, r)
}

//ToUpdateServicePage 去修改宠物服务的一个页面
func ToUpdateServicePage(w http.ResponseWriter, r *http.Request) {
	//获取要修改宠物服务的ID
	serviceID := r.FormValue("serviceID")
	//调用cardao中根据ID获取宠物服务的函数
	service, _ := dao.GetServiceByID(serviceID)
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/manager/service_modify.html"))
	//执行
	t.Execute(w, service)

}

//UpdateService 修改宠物服务
func UpdateService(w http.ResponseWriter, r *http.Request) {
	//获取要修改的宠物服务信息
	serviceID := r.PostFormValue("serviceID")
	name := r.PostFormValue("name")
	price := r.PostFormValue("price")
	num := r.PostFormValue("num")
	iServiceID, _ := strconv.ParseInt(serviceID, 10, 0)
	fprice, _ := strconv.ParseFloat(price, 64)
	inum, _ := strconv.ParseInt(num, 10, 0)
	//创建Car
	service := &model.Service{
		ID:    int(iServiceID),
		Name:  name,
		Price: fprice,
		Num:   int(inum),
	}
	//调用servicedao中修改宠物服务的函数
	dao.UpdateService(service)
	//调用处理器函数GetCars再查一次数据库看是否修改成功
	GetPageServices(w, r)
}
