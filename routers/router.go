package routers

import (
	"irisweb25/controllers"
	"irisweb25/middleware"

	"github.com/kataras/iris/v12"
)

func Register(app *iris.Application) {
	main := app.Party("/", middleware.CorsAuth()).AllowMethods(iris.MethodOptions)
	{
		v1 := main.Party("/v1")
		v1.Post("/admin/login", controllers.UserLogin)
		v1.PartyFunc("/admin", func(admin iris.Party) {
			admin.Use(middleware.JwtHandler().Serve, middleware.AuthToken) //JwtHandler().Serve会在中间件设置jwt(ContextKey的值)设置其token值， Authorization
			admin.Get("/", controllers.UserLoginInfo)                      // GET请求header中添加Authorization:Bearer eyJhbGci....token
			admin.Get("logout", controllers.UserLogout).Name = "退出"
		})
	}
}
