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

	var logoutRouter = router.Group("/logout")
	{
		logoutRouter.POST("/", controller.Logout)
	}

	var signupRouter = router.Group("/signup")
	{
		signupRouter.GET("/", controller.SignUpForm)
		signupRouter.POST("/", controller.Signup)
	}

	var articleRouter = router.Group("/article")
	/* ルートルータ(ここでは変数router)でセッションチェックミドルウェアを読み込むと永久にリダイレクトし続けるので注意 */
	articleRouter.Use(middleware.SessionCheck())
	{
		articleRouter.GET("/", controller.GetArticles)
	}

	var likeRouter = router.Group("/like")
	likeRouter.Use((middleware.SessionCheck()))
	{
		likeRouter.POST("/", controller.AddLike)
	}

	router.Run(":8000")

}
