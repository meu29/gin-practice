package service

import (
	"encoding/json"
	"gin_app/model"
	"io/ioutil"
	"regexp"
)

type ArticleService struct{}

func (a *ArticleService) GetArticles(keyword string) []model.Article {

	var jsonData, err = ioutil.ReadFile("news_data.json")
	if err != nil {
		panic(err)
	}

	var articles []model.Article

	err = json.Unmarshal(jsonData, &articles)
	if err != nil {
		panic(err)
	}

	var filterd_articles []model.Article
	var regex = regexp.MustCompile(keyword)

	/* フィルター関数はない */
	for i := 0; i <= len(articles)-1; i++ {
		if regex.MatchString(articles[i].Description) {
			filterd_articles = append(filterd_articles, articles[i])
		}
	}

	return filterd_articles

}
