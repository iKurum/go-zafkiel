package vm

import (
	"log"
	"zafkiel/package/model"
)

//LoginViewModel struct
type LoginViewModel struct {
	BaseViewModel
	Errs []string
}

//LoginViewModel10p struct
type LoginViewModel10p struct{}

//GetVM func
func (LoginViewModel10p) GetVM() LoginViewModel {
	v := LoginViewModel{}
	v.SetTitle("Login")
	return v
}

//AddError func
func (v *LoginViewModel) AddError(errs ...string) {
	v.Errs = append(v.Errs, errs...)
}

//CheckLogin func
func CheckLogin(username, password string) bool {
	user, err := model.GetUserByUsername(username)
	if err != nil {
		log.Println("Can not find username:", username)
		log.Println("Error:", err)
		return false
	}
	return user.CheckPassword(password)
}
