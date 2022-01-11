package models

import (
	"time"

	"github.com/iris-contrib/middleware/jwt"
	"github.com/jameskeane/bcrypt"
	"gorm.io/gorm"
)

//User数据模型
type User struct {
	gorm.Model
	Name     string `gorm:"not null varchar(255)"`
	Username string `gorm:"unique; varchar(255)"`
	Password string `gorm:"not null varchar(255)"`
}

//根据username查询用户信息
func UserAdminCheckLogin(username string) *User {
	user := new(User)
	IsNotFound(Db.Where("username=?", username).First(user).Error)
	return user
}

//检查登录用户，并生成登录凭证token
func CheckLogin(username, password string) (response Token, status bool, msg string) {
	user := UserAdminCheckLogin(username)
	if user.ID == 0 {
		msg = "用户不存在"
		return
	}
	if ok := bcrypt.Match(password, user.Password); !ok {
		msg = "密码错误"
		return
	}
	expTime := time.Now().Add(time.Hour * time.Duration(1)).Unix() //1小时后token过期
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": expTime,
		"iat": time.Now().Unix(),
	})
	tokenString, _ := token.SignedString([]byte(MySecret))
	oauthToken := new(OauthToken)
	oauthToken.Token = tokenString
	oauthToken.UserId = user.ID
	oauthToken.Secret = "secret"
	oauthToken.Revoked = false
	oauthToken.ExpressIn = expTime
	oauthToken.CreatedAt = time.Now()
	response = oauthToken.OauthTokenCreate()
	status = true
	msg = "登录成功"
	return
}

//退出用户
func UserAdminLogout(userId uint) bool {
	ot := UpdateOauthTokenByUserId(userId)
	return ot.Revoked
}
