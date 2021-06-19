package model

import (
	"time"
)

/* DBではなくjsonから読み込むのでgorm不要 */
type Article struct {
	Title      string    `json:"title"`
	Url        string    `json:"url"`
	UrlToImage string    `json:"urlToImage"`
	Author     string    `json:"author"`
	Date       time.Time `json:"publishdAt"`
}
