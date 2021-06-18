package controller

import (
	"gin_app/service"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"message": "",
	})
}

func Login(c *gin.Context) {

	var loginService = service.LoginService{}
	var user = loginService.Login(c.PostForm("userid"), c.PostForm("password"))

	if user.UserId == "" || user.Name == "" {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"message": "ログインに失敗しました",
		})
	} else {
		var session = sessions.Default(c)
		session.Set("UserId", user.UserId)
		session.Set("Name", user.Name)
		session.Save()
		c.Redirect(http.StatusMovedPermanently, "/article")
	}

}

func Logout(c *gin.Context) {

	var session = sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusMovedPermanently, "/login")

}
