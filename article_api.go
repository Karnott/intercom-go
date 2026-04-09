package intercom

import (
	"encoding/json"
	"fmt"

	"github.com/karnott/intercom-go/v2/interfaces"
)

// ArticleRepository defines the interface for working with Articles through the API.
type ArticleRepository interface {
	list(articleListParams) (ArticleList, error)
	find(id string) (Article, error)
}

// ArticleAPI implements ArticleRepository
type ArticleAPI struct {
	httpClient interfaces.HTTPClient
}

type articleListParams struct {
	PageParams
}

func (api ArticleAPI) list(params articleListParams) (ArticleList, error) {
	articleList := ArticleList{}
	data, err := api.httpClient.Get("/articles", params)
	if err != nil {
		return articleList, err
	}
	err = json.Unmarshal(data, &articleList)
	return articleList, err
}

func (api ArticleAPI) find(id string) (Article, error) {
	article := Article{}
	data, err := api.httpClient.Get(fmt.Sprintf("/articles/%s", id), nil)
	if err != nil {
		return article, err
	}
	err = json.Unmarshal(data, &article)
	return article, err
}
