package service

import (
	"gin_app/model"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type LoginService struct{}

func (l *LoginService) Login(userid string, password string) model.User {

	var db, err = gorm.Open("sqlite3", "sb.sqlite")
	if err != nil {
		panic(err)
	}

	var user model.User
	/* 変数userに代入 */
	db.First(&user, "userid = ? and password = ?", userid, password)

	return user

}
