package dao

import (
	"PetHome/model"
	"PetHome/utils"
	"strconv"
)

//GetServices 获取数据库中所有的宠物服务
func GetServices() ([]*model.Service, error) {
	//
	sqlStr := "select ID,Name,Price,Num from services"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var services []*model.Service
	for rows.Next() {
		service := &model.Service{}
		rows.Scan(&service.ID, &service.Name, &service.Price, &service.Num)
		services = append(services, service)

	}
	return services, nil
}

//AddService 像数据库中添加一项宠物服务
func AddService(service *model.Service) error {
	sqlStr := "insert into services(Name,Price,Num,ImgPath) values(?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, service.Name, service.Price, service.Num, service.ImgPath)
	if err != nil {
		return err
	}
	return nil
}

//DeleteService 根据宠物服务的ID从数据库删除一项宠物服务
func DeleteService(serviceID string) error {
	sqlStr := "delete from services where id = ?"
	_, err := utils.Db.Exec(sqlStr, serviceID)
	if err != nil {
		return err
	}
	return nil

}

//GetServiceByID 根据宠物服务的ID从数据库中查询出这项宠物服务
func GetServiceByID(serviceID string) (*model.Service, error) {
	sqlStr := "select ID,Name,Price,Num from services where id = ?"
	row := utils.Db.QueryRow(sqlStr, serviceID)
	service := &model.Service{}
	row.Scan(&service.ID, &service.Name, &service.Price, &service.Num)
	return service, nil
}

//UpdateService 根据宠物服务的ID修改宠物服务信息
func UpdateService(service *model.Service) error {
	sqlStr := "update services set name=? ,price=?,num=? where id = ?"
	_, err := utils.Db.Exec(sqlStr, service.Name, service.Price, service.Num, service.ID)
	if err != nil {
		return err
	}
	return nil
}

//GetPageCars 获取带分页的宠物服务信息
func GetPageServices(pageNow string) (*model.Page, error) {
	ipageNow, _ := strconv.ParseInt(pageNow, 10, 64)
	//获取数据库中宠物服务的总数
	sqlStr := "select count(*)from services"
	//定义一个变量接收总的宠物服务数量
	var pageFinal int64
	row := utils.Db.QueryRow(sqlStr)
	row.Scan(&pageFinal)
	//设置每页只显示3条
	var pageLen int64 = 3
	//获取总的页数
	//定义一个变量接收总的页数
	var pageSum int64
	if pageFinal%pageLen == 0 {
		pageSum = pageFinal / pageLen
	} else {
		pageSum = pageFinal/pageLen + 1
	}
	//获取当前页中的宠物服务
	sqlStr1 := "select ID,Name,Price,Num,ImgPath from services limit ?,?"
	rows, err := utils.Db.Query(sqlStr1, (ipageNow-1)*pageLen, pageLen)
	if err != nil {
		return nil, err
	}
	var services []*model.Service
	for rows.Next() {
		service := &model.Service{}
		rows.Scan(&service.ID, &service.Name, &service.Price, &service.Num, &service.ImgPath)
		services = append(services, service)
	}

	//创建page
	page := &model.Page{
		Services:  services,
		PageNow:   ipageNow,
		PageLen:   pageLen,
		PageSum:   pageSum,
		PageFinal: pageFinal,
	}
	return page, nil

}

//GetPageServicesByPrice 获取带分页的根据价格查询到的宠物服务信息
func GetPageServicesByPrice(pageNow string, minPrice string, maxPrice string) (*model.Page, error) {
	ipageNow, _ := strconv.ParseInt(pageNow, 10, 64)
	//根据价格获取数据库中宠物服务的总数
	sqlStr := "select count(*)from services where price between ? and ?"
	//定义一个变量接收总的宠物服务数量
	var pageFinal int64
	row := utils.Db.QueryRow(sqlStr, minPrice, maxPrice)
	row.Scan(&pageFinal)
	//设置每页只显示3条
	var pageLen int64 = 3
	//获取总的页数
	//定义一个变量接收总的页数
	var pageSum int64
	if pageFinal%pageLen == 0 {
		pageSum = pageFinal / pageLen
	} else {
		pageSum = pageFinal/pageLen + 1
	}
	//获取当前页中的宠物服务
	sqlStr1 := "select ID,Name,Price,Num,ImgPath from services where price between ? and ? limit ?,?"
	rows, err := utils.Db.Query(sqlStr1, minPrice, maxPrice, (ipageNow-1)*pageLen, pageLen)
	if err != nil {
		return nil, err
	}
	var services []*model.Service
	for rows.Next() {
		service := &model.Service{}
		rows.Scan(&service.ID, &service.Name, &service.Price, &service.Num, &service.ImgPath)
		services = append(services, service)
	}

	//创建page
	page := &model.Page{
		Services:  services,
		PageNow:   ipageNow,
		PageLen:   pageLen,
		PageSum:   pageSum,
		PageFinal: pageFinal,
	}
	return page, nil

}
