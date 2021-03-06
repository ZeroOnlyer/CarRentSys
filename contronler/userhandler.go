package contronler

import (
	"CarSys/dao"
	"CarSys/model"
	"CarSys/utils"
	"html/template"
	"net/http"
	"strconv"
)

//Logout 处理用户注销的函数
func Logout(w http.ResponseWriter, r *http.Request) {
	//获取cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		//获取cookie的value值
		cookieValue := cookie.Value
		//删除数据库中对应的session
		dao.DeleteSession(cookieValue)
		//设置cookie失效
		cookie.MaxAge = -1
		//将修改后的cookie发送给浏览器
		http.SetCookie(w, cookie)
	}
	//去首页
	GetPageCarsByPrice(w, r)

}

//Login 处理用户登录的函数
func Login(w http.ResponseWriter, r *http.Request) {
	//判断是否已经登录
	flag, _, _ := dao.IsLogin(r)
	if flag {
		//已经登录，直接去首页
		GetPageCarsByPrice(w, r)
	} else {
		//获取账号和密码clear
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		//调用customerdao中验证账号密码的方法
		user, _ := dao.CheckUserNameAndPassword(username, password)
		if user.ID > 0 {
			//账号密码正确
			//生成uuid作为session的id
			uuid := utils.CreateUUID()
			//创建一个session
			sess := &model.Session{
				SessionID: uuid,
				Name:      user.Name,
				UserID:    user.ID,
			}
			//将session保存到数据库中
			dao.AddSession(sess)
			//创建cookie，与session相关联
			cookie := http.Cookie{
				Name:     "user",
				Value:    uuid,
				HttpOnly: true,
			}
			//将cookie发送给浏览器
			http.SetCookie(w, &cookie)

			t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
			t.Execute(w, user)
		} else {
			//账号密码错误
			t := template.Must(template.ParseFiles("views/pages/user/login.html"))
			t.Execute(w, "账号或密码错误！")
		}
	}

}

//Regist 处理用户注册的函数
func Regist(w http.ResponseWriter, r *http.Request) {
	//获取新用户的信息
	name := r.PostFormValue("name")
	age, _ := strconv.Atoi(r.PostFormValue("age"))
	phone := r.PostFormValue("phone")
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	//调用customerdao中判断账号是否存在的方法
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		//账号已存在
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "账号已存在！")
	} else {
		//新账号不重复，账号可以注册
		//将新注册的用户保存到数据库中
		dao.SaveUser(name, age, phone, username, password)

		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "")
	}
}
