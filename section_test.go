package intercom

import "testing"

func TestListSections(t *testing.T) {
	sectionList, _ := (&SectionService{Repository: TestSectionAPI{t: t}}).List(PageParams{})
	sections := sectionList.Sections
	if sections[0].ID != "10" {
		t.Errorf("Got section with ID %s, expected 10", sections[0].ID)
	}
}

func TestFindSection(t *testing.T) {
	section, _ := (&SectionService{Repository: TestSectionAPI{t: t}}).Find("10")
	if section.ID != "10" {
		t.Errorf("Got section with ID %s, expected 10", section.ID)
	}
}

type TestSectionAPI struct {
	t *testing.T
}

func (t TestSectionAPI) list(params sectionListParams) (SectionList, error) {
	return SectionList{Sections: []Section{{ID: "10", Name: "Installation", ParentID: 1}}}, nil
}

func (t TestSectionAPI) find(id string) (Section, error) {
	return Section{ID: id, Name: "Installation", ParentID: 1}, nil
}
