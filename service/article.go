package service

import (
	"encoding/json"
	"gin_app/model"
	"io/ioutil"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type ArticleService struct{}

type GetArticlesResult struct {
	Articles []model.Article
	Pages    int
}

func (a *ArticleService) GetArticles(keyword string, page string) GetArticlesResult {

	var jsonData, err = ioutil.ReadFile("news_data.json")
	if err != nil {
		panic(err)
	}

	var articles []model.Article

	err = json.Unmarshal(jsonData, &articles)
	if err != nil {
		panic(err)
	}

	/* 文字列から数値に変換 クエリが数値でも文字列として受け取る */
	/* エラーの場合は0になる */
	var pageNumber, _ = strconv.Atoi(page)

	if pageNumber < 1 {
		log.Println(err)
		pageNumber = 1
	}

	var filterd_articles []model.Article
	var keywordRegex = regexp.MustCompile(keyword)
	var emailRegex = regexp.MustCompile(`[\w\-\._]+@[\w\-\._]+\.[A-Za-z]+`)
	var authorRegex = regexp.MustCompile(`https?://[\w!\?/\+\-_~=;\.,\*&@#\$%\(\)'\[\]]+`)

	var startIndex int = (pageNumber - 1) * 9
	var goalIndex int = startIndex + 8

	if goalIndex > len(articles)-1 {
		goalIndex = len(articles) - 1
	}

	/* 1ページにつき9記事 */
	/* フィルター関数はない */
	for i := startIndex; i <= goalIndex; i++ {

		var splitedStrings []string

		if authorRegex.MatchString(articles[i].Author) {
			splitedStrings = strings.Split(articles[i].Author, "/")
			articles[i].Author = splitedStrings[len(splitedStrings)-1]
		} else if emailRegex.MatchString(articles[i].Author) {
			splitedStrings = strings.Split(articles[i].Author, "@")
			articles[i].Author = splitedStrings[len(splitedStrings)-1]
		}

		if keywordRegex.MatchString(articles[i].Title) {
			filterd_articles = append(filterd_articles, articles[i])
		}

	}

	return GetArticlesResult{Articles: filterd_articles, Pages: int(math.Ceil(float64(len(articles)) / 9.0))}

}
