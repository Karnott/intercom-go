package intercom

import (
	"io/ioutil"
	"testing"
)

func TestAPIListSections(t *testing.T) {
	http := TestSectionHTTPClient{t: t, fixtureFilename: "fixtures/sections.json", expectedURI: "/help_center/sections"}
	api := SectionAPI{httpClient: &http}
	sectionList, err := api.list(sectionListParams{})
	if err != nil {
		t.Fatal(err)
	}
	if sectionList.Sections[0].ID != "10" {
		t.Errorf("Section list should start with section 10, but had %s", sectionList.Sections[0].ID)
	}
	if sectionList.Sections[0].ParentID != 1 {
		t.Errorf("Section should have parent_id 1, but had %d", sectionList.Sections[0].ParentID)
	}
}

func TestAPIFindSection(t *testing.T) {
	http := TestSectionHTTPClient{t: t, fixtureFilename: "fixtures/section.json", expectedURI: "/help_center/sections/10"}
	api := SectionAPI{httpClient: &http}
	section, err := api.find("10")
	if err != nil {
		t.Fatal(err)
	}
	if section.ID != "10" {
		t.Errorf("Section should have ID 10, but had %s", section.ID)
	}
}

type TestSectionHTTPClient struct {
	TestHTTPClient
	t               *testing.T
	fixtureFilename string
	expectedURI     string
}

func (t TestSectionHTTPClient) Get(uri string, params interface{}) ([]byte, error) {
	if uri != t.expectedURI {
		t.t.Errorf("URI was %s, expected %s", uri, t.expectedURI)
	}
	return ioutil.ReadFile(t.fixtureFilename)
}
