package service

import (
	"crypto/rand"
	"encoding/base64"
	"gin_app/model"
	"math/big"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type SignupService struct{}

func (s *SignupService) Signup(name string, password string) string {

	var db, err = gorm.Open("sqlite3", "db.sqlite")
	if err != nil {
		panic(err)
	}

	var runes = make([]byte, 64)

	for i := 0; i < 64; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(255))
		runes[i] = byte(num.Int64())
	}

	var userid = base64.RawStdEncoding.EncodeToString(runes)

	db.Create(&model.User{UserId: userid, Name: name, Password: password})
	db.Close()

	return userid

}
