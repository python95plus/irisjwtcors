package controllers

import (
	"irisweb25/models"
	"strconv"

	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

//登录处理程序
func UserLogin(ctx iris.Context) {
	aul := new(models.User)
	if err := ctx.ReadJSON(&aul); err != nil {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(models.Response{
			Status: false,
			Msg:    nil,
			Data:   "请求参数错误",
		})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	msg, status, data := models.CheckLogin(aul.Username, aul.Password)
	ctx.JSON(models.Response{
		Status: status,
		Msg:    msg,
		Data:   data,
	})
}

//退出当前账号
func UserLogout(ctx iris.Context) {
	aui := ctx.Values().GetString("auth_user_id")
	id, _ := strconv.Atoi(aui)
	uid := uint(id)
	models.UserAdminLogout(uid)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(models.Response{Status: true, Msg: nil, Data: "退出"})
}

func UserLoginInfo(ctx iris.Context) {
	aui := ctx.Values().GetString("auth_user_id")
	value := ctx.Values().Get("jwt").(*jwt.Token)

	ctx.JSON(iris.Map{
		"aui":   aui,
		"value": value,
	})
}
