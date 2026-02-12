package intercom

import (
	"io/ioutil"
	"testing"
)

func TestAPIListArticles(t *testing.T) {
	http := TestArticleHTTPClient{t: t, fixtureFilename: "fixtures/articles.json", expectedURI: "/articles"}
	api := ArticleAPI{httpClient: &http}
	articleList, err := api.list(articleListParams{})
	if err != nil {
		t.Fatal(err)
	}
	if articleList.Articles[0].ID != "123456" {
		t.Errorf("Article list should start with article 123456, but had %s", articleList.Articles[0].ID)
	}
	if articleList.Pages.Page != 1 {
		t.Errorf("Article list should be on page 1, but was on page %d", articleList.Pages.Page)
	}
	if articleList.TotalCount != 1 {
		t.Errorf("Article list should have total_count 1, but had %d", articleList.TotalCount)
	}
}

func TestAPIFindArticle(t *testing.T) {
	http := TestArticleHTTPClient{t: t, fixtureFilename: "fixtures/article.json", expectedURI: "/articles/123456"}
	api := ArticleAPI{httpClient: &http}
	article, err := api.find("123456")
	if err != nil {
		t.Fatal(err)
	}
	if article.ID != "123456" {
		t.Errorf("Article should have ID 123456, but had %s", article.ID)
	}
	if article.Title != "Getting Started" {
		t.Errorf("Article should have title 'Getting Started', but had '%s'", article.Title)
	}
	if article.State != "published" {
		t.Errorf("Article should have state 'published', but had '%s'", article.State)
	}
	if article.Statistics.Views != 150 {
		t.Errorf("Article should have 150 views, but had %d", article.Statistics.Views)
	}
}

type TestArticleHTTPClient struct {
	TestHTTPClient
	t               *testing.T
	fixtureFilename string
	expectedURI     string
}

func (t TestArticleHTTPClient) Get(uri string, params interface{}) ([]byte, error) {
	if uri != t.expectedURI {
		t.t.Errorf("URI was %s, expected %s", uri, t.expectedURI)
	}
	return ioutil.ReadFile(t.fixtureFilename)
}
