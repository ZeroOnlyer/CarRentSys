package model

//Page 定义分页
type Page struct {
	Cars      []*Car //每一页的所有车辆信息
	PageNow   int64  //当前页
	PageLen   int64  //每一页的显示数量
	PageSum   int64  //总页数
	PageFinal int64  //一共有多少辆车
	MinPrice  string //价格范围
	MaxPrice  string //价格范围
	IsLogin   bool   //判断是否登录
	Name      string //用户的姓名
	SuperID   bool   //判断是否是管理员
}

//IsHasPrev 判断是否有上一页
func (p *Page) IsHasPrev() bool {
	return p.PageNow > 1
}

//IsHasNext 判断是否有下一页
func (p *Page) IsHasNext() bool {
	return p.PageNow < p.PageSum
}

//GetPrevPage 获取上一页
func (p *Page) GetPrevPage() int64 {
	//判断是否有上一页
	if p.IsHasPrev() {
		return p.PageNow - 1
	}
	return 1

}

//GetNextPage 获取下一页
func (p *Page) GetNextPage() int64 {
	//判断是否有下一页
	if p.IsHasNext() {
		return p.PageNow + 1
	}
	return p.PageSum

}
