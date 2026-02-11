package intercom

import (
	"encoding/json"
	"fmt"

	"github.com/karnott/intercom-go/interfaces"
)

// SectionRepository defines the interface for working with Sections through the API.
type SectionRepository interface {
	list(sectionListParams) (SectionList, error)
	find(id string) (Section, error)
}

// SectionAPI implements SectionRepository
type SectionAPI struct {
	httpClient interfaces.HTTPClient
}

type sectionListParams struct {
	PageParams
}

func (api SectionAPI) list(params sectionListParams) (SectionList, error) {
	sectionList := SectionList{}
	data, err := api.httpClient.Get("/help_center/sections", params)
	if err != nil {
		return sectionList, err
	}
	err = json.Unmarshal(data, &sectionList)
	return sectionList, err
}

func (api SectionAPI) find(id string) (Section, error) {
	section := Section{}
	data, err := api.httpClient.Get(fmt.Sprintf("/help_center/sections/%s", id), nil)
	if err != nil {
		return section, err
	}
	err = json.Unmarshal(data, &section)
	return section, err
}
