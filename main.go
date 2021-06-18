package main

import (
	"gin_app/controller"
	"gin_app/middleware"
	"gin_app/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	var initService = service.InitService{}
	initService.Init()

	var router = gin.Default()
	router.LoadHTMLGlob("template/*.html")

	var store = cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	var loginRouter = router.Group("/login")
	{
		loginRouter.GET("/", controller.LoginForm)
		loginRouter.POST("/", controller.Login)
	}

	var signupRouter = router.Group("/signup")
	{
		signupRouter.GET("/", controller.SignUpForm)
		signupRouter.POST("/", controller.Signup)
	}

	var homeRouter = router.Group("/home")
	/* ルートルータ(ここでは変数router)でセッションチェックミドルウェアを読み込むと永久にリダイレクトし続けるので注意 */
	homeRouter.Use(middleware.SessionCheck())
	{
		homeRouter.GET("/", controller.Home)
	}

	router.Run(":8000")

}
