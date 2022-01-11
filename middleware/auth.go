package middleware

import (
	"irisweb25/models"
	"time"

	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func AuthToken(ctx *context.Context) {
	value := ctx.Values().Get("jwt").(*jwt.Token)
	token := models.GetOauthTokenByToken(value.Raw)
	if token.Revoked || token.ExpressIn < time.Now().Unix() {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.StopExecution()
		return
	}
	ctx.Values().Set("auth_user_id", token.UserId)
	ctx.Next()
}

func JwtHandler() *jwt.Middleware {
	return jwt.New(jwt.Config{
		ValidationKeyGetter: func(*jwt.Token) (interface{}, error) {
			return []byte(models.MySecret), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
}
