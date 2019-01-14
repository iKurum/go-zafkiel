package vm

import (
	"go-zafkiel/package/model"
)

//ExploreViewModel struct
type ExploreViewModel struct {
	BaseViewModel
	BasePageViewModel
	Posts []model.Post
}

//ExploreViewModel10p struct
type ExploreViewModel10p struct{}

//GetVM func
func (ExploreViewModel10p) GetVM(username string, page, limit int) ExploreViewModel {
	posts, total, _ := model.GetPostsByPageAndLimit(page, limit)
	v := ExploreViewModel{}
	v.SetTitle("Explore")
	v.Posts = *posts
	v.SetBasePageViewModel(total, page, limit)
	v.SetCurrentUser(username)
	return v
}
