package vm

import (
	"zafkiel/package/config"
	"zafkiel/package/model"
)

//EmailViewModel struct
type EmailViewModel struct {
	Username string
	Token    string
	Server   string
}

//EmailViewModel10p struct
type EmailViewModel10p struct{}

//GetVM func
func (EmailViewModel10p) GetVM(email string) EmailViewModel {
	v := EmailViewModel{}
	u, _ := model.GetUserByEmail(email)
	v.Username = u.Username
	v.Token, _ = u.GenerateToken()
	v.Server = config.GetServerURL()
	return v
}
