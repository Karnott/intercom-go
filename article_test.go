package intercom

import "testing"

func TestListArticles(t *testing.T) {
	articleList, _ := (&ArticleService{Repository: TestArticleAPI{t: t}}).List(PageParams{})
	articles := articleList.Articles
	if articles[0].ID != "123456" {
		t.Errorf("Got article with ID %s, expected 123456", articles[0].ID)
	}
}

func TestFindArticle(t *testing.T) {
	article, _ := (&ArticleService{Repository: TestArticleAPI{t: t}}).Find("123456")
	if article.ID != "123456" {
		t.Errorf("Got article with ID %s, expected 123456", article.ID)
	}
}

type TestArticleAPI struct {
	t *testing.T
}

func (t TestArticleAPI) list(params articleListParams) (ArticleList, error) {
	return ArticleList{Articles: []Article{{ID: "123456", Title: "Getting Started", State: "published"}}}, nil
}

func (t TestArticleAPI) find(id string) (Article, error) {
	return Article{ID: id, Title: "Getting Started", State: "published"}, nil
}
