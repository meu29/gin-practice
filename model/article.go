package model

import (
	"time"
)

/* DBではなくjsonから読み込むのでgorm不要 */
type Article struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
	Url         string    `json:"url"`
	UrlToImage  string    `json:"urlToImage"`
	Date        time.Time `json:"publishdAt"`
}
