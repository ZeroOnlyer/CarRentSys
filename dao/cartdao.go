package dao

import (
	"PetHome/model"
	"PetHome/utils"
)

//AddCart 向订单表中插入订单
func AddCart(cart *model.Cart) error {
	sqlStr := "insert into carts(id,total_count,total_amount,user_id)values(?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, cart.CartID, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserID)
	if err != nil {
		return err
	}
	//获取订单中所有的订单项
	cartItems := cart.CartItems
	//遍历得到每一个订单项
	for _, cartItem := range cartItems {
		//将订单项插入到数据库中
		AddCartItem(cartItem)
	}
	return nil
}

//GetCartByUserID 根据用户id从数据库中查询对应的订单
func GetCartByUserID(userID int) (*model.Cart, error) {
	sqlStr := "select ID,total_count,total_amount,user_id from carts where user_id = ?"
	row := utils.Db.QueryRow(sqlStr, userID)
	cart := &model.Cart{}
	err := row.Scan(&cart.CartID, &cart.TotalCount, &cart.TotalAmount, &cart.UserID)
	if err != nil {
		return nil, err
	}
	//获取当前订单里所有的订单项
	cartItems, _ := GetCartItemsByCartID(cart.CartID)
	//把所有的订单项设置到查询到的订单中
	cart.CartItems = cartItems
	return cart, nil
}

//UpdateCart 更新订单中宠物服务的总数量和总金额
func UpdateCart(cart *model.Cart) error {
	sqlStr := "update carts set total_count=?,total_amount=? where id = ?"
	_, err := utils.Db.Exec(sqlStr, cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartID)
	if err != nil {
		return err
	}
	return nil
}

//DeleteCartByCartID 根据订单的id删除订单
func DeleteCartByCartID(cartID string) error {
	//删除订单前需要先删除所有对应的订单项
	err := DeleteCartItemsByCartID(cartID)
	if err != nil {
		return err
	}
	sqlStr := "delete from carts where id = ?"
	_, err2 := utils.Db.Exec(sqlStr, cartID)
	if err2 != nil {
		return err2
	}
	return nil
}
