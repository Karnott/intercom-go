package intercom

import "testing"

func TestListCollections(t *testing.T) {
	collectionList, _ := (&CollectionService{Repository: TestCollectionAPI{t: t}}).List(PageParams{})
	collections := collectionList.Collections
	if collections[0].ID != "1" {
		t.Errorf("Got collection with ID %s, expected 1", collections[0].ID)
	}
}

func TestFindCollection(t *testing.T) {
	collection, _ := (&CollectionService{Repository: TestCollectionAPI{t: t}}).Find("1")
	if collection.ID != "1" {
		t.Errorf("Got collection with ID %s, expected 1", collection.ID)
	}
}

type TestCollectionAPI struct {
	t *testing.T
}

func (t TestCollectionAPI) list(params collectionListParams) (CollectionList, error) {
	return CollectionList{Collections: []Collection{{ID: "1", Name: "Getting Started"}}}, nil
}

func (t TestCollectionAPI) find(id string) (Collection, error) {
	return Collection{ID: id, Name: "Getting Started"}, nil
}
