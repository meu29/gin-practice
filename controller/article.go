package controller

import (
	"gin_app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var articleService = service.ArticleService{}

func GetArticles(c *gin.Context) {

	c.HTML(http.StatusOK, "article.html", gin.H{
		"articles": articleService.GetArticles(c.Query("keyword")),
	})

}
