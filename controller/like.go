package controller

import (
	"gin_app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var likeService = service.LikeService{}

func AddLike(c *gin.Context) {
	likeService.AddLike(c.PostForm("articleid"), c.PostForm("userid"))
	c.Redirect(http.StatusMovedPermanently, "/")
}
