package intercom

import (
	"io/ioutil"
	"testing"
)

func TestAPIListCollections(t *testing.T) {
	http := TestCollectionHTTPClient{t: t, fixtureFilename: "fixtures/collections.json", expectedURI: "/help_center/collections"}
	api := CollectionAPI{httpClient: &http}
	collectionList, err := api.list(collectionListParams{})
	if err != nil {
		t.Fatal(err)
	}
	if collectionList.Collections[0].ID != "1" {
		t.Errorf("Collection list should start with collection 1, but had %s", collectionList.Collections[0].ID)
	}
	if collectionList.Collections[0].Name != "Getting Started" {
		t.Errorf("Collection should have name 'Getting Started', but had '%s'", collectionList.Collections[0].Name)
	}
}

func TestAPIFindCollection(t *testing.T) {
	http := TestCollectionHTTPClient{t: t, fixtureFilename: "fixtures/collection.json", expectedURI: "/help_center/collections/1"}
	api := CollectionAPI{httpClient: &http}
	collection, err := api.find("1")
	if err != nil {
		t.Fatal(err)
	}
	if collection.ID != "1" {
		t.Errorf("Collection should have ID 1, but had %s", collection.ID)
	}
}

type TestCollectionHTTPClient struct {
	TestHTTPClient
	t               *testing.T
	fixtureFilename string
	expectedURI     string
}

func (t TestCollectionHTTPClient) Get(uri string, params interface{}) ([]byte, error) {
	if uri != t.expectedURI {
		t.t.Errorf("URI was %s, expected %s", uri, t.expectedURI)
	}
	return ioutil.ReadFile(t.fixtureFilename)
}
