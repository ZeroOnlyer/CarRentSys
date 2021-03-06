package dao

import (
	"CarSys/model"
	"CarSys/utils"
)

//AddOrderItem 向数据库中插入付款单项
func AddOrderItem(orderItem *model.OrderItem) error {
	sqlStr := "insert into order_items(count,amount,name,price,order_id) values(?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, orderItem.Count, orderItem.Amount, orderItem.Name, orderItem.Price, orderItem.OrderID)
	if err != nil {
		return err
	}
	return nil
}
