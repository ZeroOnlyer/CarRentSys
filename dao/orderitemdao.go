package dao

import (
	"PetHome/model"
	"PetHome/utils"
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

//GetOrderItemsByOrderID 根据订单号获取该订单的所有订单项
func GetOrderItemsByOrderID(orderID string) ([]*model.OrderItem, error) {
	//写sql语句
	sql := "select id,count,amount,name,price,order_id from order_items where order_id = ?"
	//执行
	rows, err := utils.Db.Query(sql, orderID)
	if err != nil {
		return nil, err
	}
	var orderItems []*model.OrderItem
	for rows.Next() {
		orderItem := &model.OrderItem{}
		rows.Scan(&orderItem.OrderItemID, &orderItem.Count, &orderItem.Amount, &orderItem.Name, &orderItem.Price, &orderItem.OrderID)
		//添加到切片中
		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil
}
