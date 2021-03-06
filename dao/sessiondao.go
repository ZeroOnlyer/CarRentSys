package dao

import (
	"CarSys/model"
	"CarSys/utils"
	"net/http"
)

//AddSession 向数据库中添加session
func AddSession(sess *model.Session) error {
	sqlStr := "insert into sessions values(?,?,?)"
	_, err := utils.Db.Exec(sqlStr, sess.SessionID, sess.Name, sess.UserID)
	if err != nil {
		return err
	}
	return nil
}

//DeleteSession 删除数据库中的Session
func DeleteSession(sessID string) error {
	sqlStr := "delete from sessions where SessionID = ?"
	_, err := utils.Db.Exec(sqlStr, sessID)
	if err != nil {
		return err
	}
	return nil
}

//GetSessionByID 根据session的ID值从数据库中查询session
func GetSessionByID(sessID string) (*model.Session, error) {
	sqlStr := "select sessionID ,Name,userID from sessions where sessionID = ?"
	//预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	//执行
	row := inStmt.QueryRow(sessID)
	//创建session
	sess := &model.Session{}
	//从数据库中的字段scan给session
	row.Scan(&sess.SessionID, &sess.Name, &sess.UserID)
	return sess, nil
}

//IsLogin 判断用户是否登录
func IsLogin(r *http.Request) (bool, *model.Session, int) {
	//根据cookie的name获取cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		//获取cookie的value
		cookieValue := cookie.Value
		//根据cookievalue去数据库中查询session
		session, _ := GetSessionByID(cookieValue)
		if session.UserID > 0 {
			//已经登录
			return true, session, session.UserID
		}
	}
	return false, nil, 0 //没有登录
}
