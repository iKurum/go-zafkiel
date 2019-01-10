package vm

import (
	"log"
	"zafkiel/package/model"
)

//ResetPasswordRequestViewModel struct
type ResetPasswordRequestViewModel struct {
	LoginViewModel
}

//ResetPasswordRequestViewModel10p struct
type ResetPasswordRequestViewModel10p struct{}

//GetVM func
func (ResetPasswordRequestViewModel10p) GetVM() ResetPasswordRequestViewModel {
	v := ResetPasswordRequestViewModel{}
	v.SetTitle("Forget Password")
	return v
}

//CheckEmailExist func
func CheckEmailExist(email string) bool {
	_, err := model.GetUserByEmail(email)
	if err != nil {
		log.Println("Can not find email:", email)
		return false
	}
	return true
}
