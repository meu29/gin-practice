package controller

import (
	"gin_app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var articleService = service.ArticleService{}

func GetArticles(c *gin.Context) {

	var articles service.GetArticlesResult = articleService.GetArticles(c.Query("keyword"), c.Query("page"))

	/* これらの値はhtmlが読み込んでいる別のhtml(今回は_header.html)も使用できる */
	c.HTML(http.StatusOK, "article.html", gin.H{
		"articles":   articles.Articles,
		"beforePage": articles.BeforePage,
		"nowPage":    articles.NowPage,
		"afterPage":  articles.AfterPage,
		"maxPage":    articles.MaxPage,
		"loginState": true,
	})

}
