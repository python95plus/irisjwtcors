package models

import "gorm.io/gorm"

type Token struct {
	Token string `json:"access_token"`
}

//token数据模型
type OauthToken struct {
	gorm.Model
	Token     string `gorm:"not null default '' comment('Token') varchar(255)"`
	UserId    uint   `gorm:"not null default 0 comment('UserId') bigint(20)"`
	Secret    string `gorm:"not null default '' comment('Secret') varchar(255)"`
	ExpressIn int64  `gorm:"not null default 0 comment('是否标准库') bigint(20)"`
	Revoked   bool   //退出账号时，值设置为true，且此时token作废
}

//创建token
func (ot *OauthToken) OauthTokenCreate() (response Token) {
	Db.Create(ot)
	response = Token{Token: ot.Token}
	return
}

//获取access_token信息
func GetOauthTokenByToken(token string) *OauthToken {
	ot := new(OauthToken)
	Db.Where("token=?", token).First(ot)
	return ot
}

//作废token
func UpdateOauthTokenByUserId(userId uint) (ot *OauthToken) {
	Db.Model(ot).Where("revoked=?", false).Where("user_id=?", userId).Updates(map[string]interface{}{"revoked": true})
	return
}
