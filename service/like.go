package service

import (
	"gin_app/model"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type LikeService struct{}

func (l *LikeService) AddLike(articeId string, userId string) {

	var db, err = gorm.Open("sqlite3", "db.sqlite")
	if err != nil {
		panic(err)
	}

	db.Create(&model.Like{ArticleId: articeId, UserId: userId})

}
