package dao

import (
	"CarSys/model"
	"CarSys/utils"
)

//AddCartItem 向订单项表中插入订单项
func AddCartItem(cartItem *model.CartItem) error {
	sqlStr := "insert into cart_items(count,amount,car_id,cart_id) values(?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmout(), cartItem.Car.ID, cartItem.CartID)
	if err != nil {
		return err
	}
	return nil
}

//GetCartItemByCarIDAndCartID  根据车辆的id和订单的id来获取对应的订单项
func GetCartItemByCarIDAndCartID(carID string, cartID string) (*model.CartItem, error) {
	sqlStr := "select id,count,amount,cart_id from cart_items where car_ID=? and cart_id = ?"
	row := utils.Db.QueryRow(sqlStr, carID, cartID)
	cartItem := &model.CartItem{}
	err := row.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
	if err != nil {
		return nil, err
	}
	//根据车辆id查询车辆信息
	car, _ := GetCarByID(carID)
	//将车辆设置到订单项中
	cartItem.Car = car
	return cartItem, nil
}

//UpdateCarCount 根据购物项来更新订单中车辆的数量和金额
func UpdateCarCount(cartItem *model.CartItem) error {
	sqlStr := "update cart_items set count=? ,amount = ? where car_id = ? and cart_id=? "
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmout(), cartItem.Car.ID, cartItem.CartID)
	if err != nil {
		return err
	}
	return nil
}

//GetCartItemsByCartID 根据订单的id来获取订单中所有的订单项
func GetCartItemsByCartID(cartID string) ([]*model.CartItem, error) {
	sqlStr := "select id,count,amount,car_id,cart_id from cart_items where cart_ID=?"
	rows, err := utils.Db.Query(sqlStr, cartID)
	if err != nil {
		return nil, err
	}
	var cartItems []*model.CartItem
	for rows.Next() {
		//设置一个变量来接收car_id
		var carID string
		cartItem := &model.CartItem{}
		err2 := rows.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &carID, &cartItem.CartID)
		if err2 != nil {
			return nil, err
		}
		//根据car_id获取车辆信息
		car, _ := GetCarByID(carID)
		//将车辆信息设置到订单项中
		cartItem.Car = car
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}

//DeleteCartItemsByCartID 根据订单的id删除所有的订单项
func DeleteCartItemsByCartID(cartID string) error {
	sqlStr := "delete from cart_items where cart_id = ?"
	_, err := utils.Db.Exec(sqlStr, cartID)
	if err != nil {
		return err
	}
	return nil
}

//DeleteCartItemByID 根据订单项的id删除订单项
func DeleteCartItemByID(cartItemID string) error {
	sqlStr := "delete from cart_items where id = ?"
	_, err := utils.Db.Exec(sqlStr, cartItemID)
	if err != nil {
		return err
	}
	return nil
}
