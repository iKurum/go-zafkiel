package model

import (
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//User struct
type User struct {
	ID           int    `gorm:"primary_key"`
	Username     string `gorm:"type:varchar(64)"`
	Email        string `gorm:"type:varchar(120)"`
	PasswordHash string `gorm:"type:varchar(128)"`
	LastSeen     *time.Time
	AboutMe      string `gorm:"type:varchar(140)"`
	Avatar       string `gorm:"type:varchar(200)"`
	Posts        []Post
	Followers    []*User `gorm:"many2many:follower;association_jointable_foreignkey:follower_id"`
}

//SetPassword func
func (u *User) SetPassword(password string) {
	u.PasswordHash = GeneratePasswordHash(password)
}

//CheckPassword func
func (u *User) CheckPassword(password string) bool {
	return GeneratePasswordHash(password) == u.PasswordHash
}

//GetUserByUsername func
func GetUserByUsername(username string) (*User, error) {
	var user User
	if err := db.Where("username=?", username).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

//SetAvatar func : 设置头像
func (u *User) SetAvatar(email string) {
	u.Avatar = fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", MD5(email))
}

//AddUser func
func AddUser(username, password, email string) error {
	user := User{Username: username, Email: email}
	user.SetPassword(password)
	user.SetAvatar(email)
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return user.FollowSelf()
}

//UpdateUserByUsername func
func UpdateUserByUsername(username string, contents map[string]interface{}) error {
	item, err := GetUserByUsername(username)
	if err != nil {
		return err
	}
	return db.Model(item).Updates(contents).Error
}

//UpdateLastSeen func
func UpdateLastSeen(username string) error {
	c := map[string]interface{}{"last_seen": time.Now()}
	return UpdateUserByUsername(username, c)
}

//UpdateAboutMe func
func UpdateAboutMe(username, text string) error {
	c := map[string]interface{}{"about_me": text}
	return UpdateUserByUsername(username, c)
}

//UpdatePassword func
func UpdatePassword(username, password string) error {
	c := map[string]interface{}{"password_hash": MD5(password)}
	return UpdateUserByUsername(username, c)
}

//Follow func
func (u *User) Follow(username string) error {
	other, err := GetUserByUsername(username)
	if err != nil {
		return err
	}
	return db.Model(other).Association("Followers").Append(u).Error
}

//Unfollow func
func (u *User) Unfollow(username string) error {
	other, err := GetUserByUsername(username)
	if err != nil {
		return err
	}
	return db.Model(other).Association("Followers").Delete(u).Error
}

//FollowSelf func
func (u *User) FollowSelf() error {
	return db.Model(u).Association("Followers").Append(u).Error
}

//FollowersCount func
func (u *User) FollowersCount() int {
	return db.Model(u).Association("Followers").Count()
}

//FollowingIDs func
func (u *User) FollowingIDs() []int {
	var i []int
	rows, err := db.Table("follower").Where("follower_id = ?", u.ID).Select("user_id, follower_id").Rows()
	if err != nil {
		log.Println("Counting Following error:", err)
		return i
	}
	defer rows.Close()
	for rows.Next() {
		var id, followerID int
		rows.Scan(&id, &followerID)
		i = append(i, id)
	}
	return i
}

//FollowingCount func
func (u *User) FollowingCount() int {
	i := u.FollowingIDs()
	return len(i)
}

//FollowingPostsByPageAndLimit func
func (u *User) FollowingPostsByPageAndLimit(page, limit int) (*[]Post, int, error) {
	var total int
	var posts []Post
	offset := (page - 1) * limit
	i := u.FollowingIDs()
	if err := db.Preload("User").Order("timestamp desc").Where("user_id in (?)", i).Offset(offset).Limit(limit).Find(&posts).Error; err != nil {
		return nil, total, err
	}
	db.Model(&Post{}).Where("user_id in (?)", i).Count(&total)
	return &posts, total, nil
}

//IsFollowedByUser func
func (u *User) IsFollowedByUser(username string) bool {
	user, _ := GetUserByUsername(username)
	i := user.FollowingIDs()
	for _, id := range i {
		if u.ID == id {
			return true
		}
	}
	return false
}

//CreatePost func
func (u *User) CreatePost(body string) error {
	post := Post{Body: body, UserID: u.ID}
	return db.Create(&post).Error
}

//GetUserByEmail func
func GetUserByEmail(email string) (*User, error) {
	var user User
	if err := db.Where("email = ?", email).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

//GenerateToken func
func (u *User) GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.Username,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte("secret"))
}

//CheckToken func
func CheckToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["username"].(string), nil
	}
	return "", err
}

//FormattedLastSeen func : 时间格式
func (u *User) FormattedLastSeen() string {
	return u.LastSeen.Format("2006-01-02 15:04:05")
}
