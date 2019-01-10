package vm

import "zafkiel/package/model"

//ProfileEditViewModel struct
type ProfileEditViewModel struct {
	LoginViewModel
	ProfileUser model.User
}

//ProfileEditViewModel10p struct
type ProfileEditViewModel10p struct{}

//GetVM func
func (ProfileEditViewModel10p) GetVM(username string) ProfileEditViewModel {
	v := ProfileEditViewModel{}
	u, _ := model.GetUserByUsername(username)
	v.SetTitle("Profile Edit")
	v.SetCurrentUser(username)
	v.ProfileUser = *u
	return v
}

//UpdateAboutMe func
func UpdateAboutMe(username, text string) error {
	return model.UpdateAboutMe(username, text)
}
