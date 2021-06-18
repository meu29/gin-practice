package service

import (
	"gin_app/model"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type InitService struct{}

/* レシーバ(このメソッドが構造体に属していることを示す) */
func (i *InitService) Init() {

	var db, err = gorm.Open("sqlite3", "db.sqlite")
	if err != nil {
		panic(err)
	}

	/* 自動マイグレーション(テーブルの作成) */
	/* db.sqliteファイルが存在しない場合も自動で作成 */
	db.AutoMigrate(&model.User{})
	db.Close()

}
