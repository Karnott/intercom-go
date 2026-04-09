package intercom

import (
	"encoding/json"
	"fmt"

	"github.com/karnott/intercom-go/v2/interfaces"
)

// CollectionRepository defines the interface for working with Collections through the API.
type CollectionRepository interface {
	list(collectionListParams) (CollectionList, error)
	find(id string) (Collection, error)
}

// CollectionAPI implements CollectionRepository
type CollectionAPI struct {
	httpClient interfaces.HTTPClient
}

type collectionListParams struct {
	PageParams
}

func (api CollectionAPI) list(params collectionListParams) (CollectionList, error) {
	collectionList := CollectionList{}
	data, err := api.httpClient.Get("/help_center/collections", params)
	if err != nil {
		return collectionList, err
	}
	err = json.Unmarshal(data, &collectionList)
	return collectionList, err
}

func (api CollectionAPI) find(id string) (Collection, error) {
	collection := Collection{}
	data, err := api.httpClient.Get(fmt.Sprintf("/help_center/collections/%s", id), nil)
	if err != nil {
		return collection, err
	}
	err = json.Unmarshal(data, &collection)
	return collection, err
}
