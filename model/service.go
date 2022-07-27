package model

//Service 定义宠物服务
type Service struct {
	ID      int     //ID
	Name    string  //名
	Price   float64 //价格
	Num     int     //库存
	ImgPath string  //图片
}
