package vm

import "go-zafkiel/package/model"

//ResetPasswordViewModel struct
type ResetPasswordViewModel struct {
	LoginViewModel
	Token string
}

//ResetPasswordViewModel10p struct
type ResetPasswordViewModel10p struct{}

//GetVM func
func (ResetPasswordViewModel10p) GetVM(token string) ResetPasswordViewModel {
	v := ResetPasswordViewModel{}
	v.SetTitle("Reset Password")
	v.Token = token
	return v
}

//CheckToken func
func CheckToken(tokenString string) (string, error) {
	return model.CheckToken(tokenString)
}

//ResetUserPassword func
func ResetUserPassword(username, password string) error {
	return model.UpdatePassword(username, password)
}
