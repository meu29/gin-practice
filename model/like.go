package model

import (
	"github.com/jinzhu/gorm"
)

type Like struct {
	gorm.Model
	ArticleId string /* 記事urlをidにする */
	UserId    string
}
