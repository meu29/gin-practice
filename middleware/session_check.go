package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionCheck() gin.HandlerFunc {

	return func(c *gin.Context) {
		var session = sessions.Default(c)
		if session.Get("UserId") == nil {
			/* ログインページにリダイレクト */
			c.Redirect(http.StatusMovedPermanently, "/login")
		} else {
			/* コントローラーで定義した本来の処理に戻る(本来の目的のページに遷移) */
			c.Next()
		}
	}

}
