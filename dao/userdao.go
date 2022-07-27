package dao

import (
	"PetHome/model"
	"PetHome/utils"
)

//CheckUserNameAndPassword 根据账号和密码从数据库查询一条记录
func CheckUserNameAndPassword(username string, password string) (*model.User, error) {
	//sql
	sqlStr := "select id,name,age,phone,username,password from users where username = ? and password = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, username, password)
	user := &model.User{}
	row.Scan(&user.ID, &user.Name, &user.Age, &user.Phone, &user.Username, &user.Password)
	return user, nil
}

//CheckUserName 判断账号是否存在，没有就可以注册
func CheckUserName(username string) (*model.User, error) {
	//sql
	sqlStr := "select id,name,age,phone,username,password from users where username = ? "
	//执行
	row := utils.Db.QueryRow(sqlStr, username)
	user := &model.User{}
	row.Scan(&user.ID, &user.Name, &user.Age, &user.Phone, &user.Username, &user.Password)
	return user, nil
}

//SaveUser 向数据库插入用户信息
func SaveUser(name string, age int, phone string, username string, password string) error {
	//sql
	sqlStr := "insert into users(name,age,phone,username,password)values(?,?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, name, age, phone, username, password)
	if err != nil {
		return err
	}
	return nil
}
