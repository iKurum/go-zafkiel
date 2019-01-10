package vm

//BaseViewModel struct
type BaseViewModel struct {
	Title       string
	CurrentUser string
}

//BasePageViewModel struct
// PrevPage: 上一页的页码
// NextPage: 下一页的页码
// Total: 总页数
// CurrentPage: 当前页码
// Limit: 每页显示项目数
type BasePageViewModel struct {
	PrevPage    int
	NextPage    int
	Total       int
	CurrentPage int
	Limit       int
}

//SetTitle func
func (v *BaseViewModel) SetTitle(title string) {
	v.Title = title
}

//SetCurrentUser func
func (v *BaseViewModel) SetCurrentUser(username string) {
	v.CurrentUser = username
}

//SetPrevAndNextPage func
func (v *BasePageViewModel) SetPrevAndNextPage() {
	if v.CurrentPage > 1 {
		v.PrevPage = v.CurrentPage - 1
	}
	if (v.Total-1)/v.Limit >= v.CurrentPage {
		v.NextPage = v.CurrentPage + 1
	}
}

//SetBasePageViewModel func
func (v *BasePageViewModel) SetBasePageViewModel(total, page, limit int) {
	v.Total = total
	v.CurrentPage = page
	v.Limit = limit
	v.SetPrevAndNextPage()
}
