package dao

import (
	"CarSys/model"
	"CarSys/utils"
	"strconv"
)

//GetCars 获取数据库中所有的车辆
func GetCars() ([]*model.Car, error) {
	//
	sqlStr := "select ID,Name,Price,Num from cars"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var cars []*model.Car
	for rows.Next() {
		car := &model.Car{}
		rows.Scan(&car.ID, &car.Name, &car.Price, &car.Num)
		cars = append(cars, car)

	}
	return cars, nil
}

//AddCar 像数据库中添加一辆汽车
func AddCar(car *model.Car) error {
	sqlStr := "insert into cars(Name,Price,Num) values(?,?,?)"
	_, err := utils.Db.Exec(sqlStr, car.Name, car.Price, car.Num)
	if err != nil {
		return err
	}
	return nil
}

//DeleteCar 根据汽车的ID从数据库删除一辆汽车
func DeleteCar(carID string) error {
	sqlStr := "delete from cars where id = ?"
	_, err := utils.Db.Exec(sqlStr, carID)
	if err != nil {
		return err
	}
	return nil

}

//GetCarByID 根据汽车的ID从数据库中查询出这辆车
func GetCarByID(carID string) (*model.Car, error) {
	sqlStr := "select ID,Name,Price,Num from cars where id = ?"
	row := utils.Db.QueryRow(sqlStr, carID)
	car := &model.Car{}
	row.Scan(&car.ID, &car.Name, &car.Price, &car.Num)
	return car, nil
}

//UpdateCar 根据车辆的ID修改车辆信息
func UpdateCar(car *model.Car) error {
	sqlStr := "update cars set name=? ,price=?,num=? where id = ?"
	_, err := utils.Db.Exec(sqlStr, car.Name, car.Price, car.Num, car.ID)
	if err != nil {
		return err
	}
	return nil
}

//GetPageCars 获取带分页的车辆信息
func GetPageCars(pageNow string) (*model.Page, error) {
	ipageNow, _ := strconv.ParseInt(pageNow, 10, 64)
	//获取数据库中车辆的总数
	sqlStr := "select count(*)from cars"
	//定义一个变量接收总的车辆数量
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
	//获取当前页中的车辆
	sqlStr1 := "select ID,Name,Price,Num from cars limit ?,?"
	rows, err := utils.Db.Query(sqlStr1, (ipageNow-1)*pageLen, pageLen)
	if err != nil {
		return nil, err
	}
	var cars []*model.Car
	for rows.Next() {
		car := &model.Car{}
		rows.Scan(&car.ID, &car.Name, &car.Price, &car.Num)
		cars = append(cars, car)
	}

	//创建page
	page := &model.Page{
		Cars:      cars,
		PageNow:   ipageNow,
		PageLen:   pageLen,
		PageSum:   pageSum,
		PageFinal: pageFinal,
	}
	return page, nil

}

//GetPageCarsByPrice 获取带分页的根据价格查询到的车辆信息
func GetPageCarsByPrice(pageNow string, minPrice string, maxPrice string) (*model.Page, error) {
	ipageNow, _ := strconv.ParseInt(pageNow, 10, 64)
	//根据价格获取数据库中车辆的总数
	sqlStr := "select count(*)from cars where price between ? and ?"
	//定义一个变量接收总的车辆数量
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
	//获取当前页中的车辆
	sqlStr1 := "select ID,Name,Price,Num from cars where price between ? and ? limit ?,?"
	rows, err := utils.Db.Query(sqlStr1, minPrice, maxPrice, (ipageNow-1)*pageLen, pageLen)
	if err != nil {
		return nil, err
	}
	var cars []*model.Car
	for rows.Next() {
		car := &model.Car{}
		rows.Scan(&car.ID, &car.Name, &car.Price, &car.Num)
		cars = append(cars, car)
	}

	//创建page
	page := &model.Page{
		Cars:      cars,
		PageNow:   ipageNow,
		PageLen:   pageLen,
		PageSum:   pageSum,
		PageFinal: pageFinal,
	}
	return page, nil

}
