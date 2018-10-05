package xorm

type PageParams interface {
	Page() int       // 页码数
	PageSize() int   //页面数量
	Limit() []int    //分页
	OrderBy() string //排序字段
	Asc() bool       //排序顺序 默认为false,也就是默认倒序
}
