package models

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db       *gorm.DB
	err      error
	MySecret = "HS2JDFKhu7Y1av7b"
)

func Register() {
	dsn := "root:@tcp(localhost:3306)/iris?charset=utf8&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		color.Red(fmt.Sprintf("gorm open 错误: %v", err))
	}
}

func IsNotFound(err error) {
	if ok := errors.Is(err, gorm.ErrRecordNotFound); !ok && err != nil {
		color.Red(fmt.Sprintf("error :%v \n ", err))
	}
}

type Response struct {
	Status bool        `json:"status"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"data"`
}
