package intercom

import "fmt"

// SectionService handles interactions with the API through a SectionRepository.
type SectionService struct {
	Repository SectionRepository
}

// Section represents a Help Center Section in Intercom (a subdivision of a Collection).
type Section struct {
	ID          string `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
	Name        string `json:"name,omitempty"`
	ParentID    int64  `json:"parent_id,omitempty"`
	CreatedAt   int64  `json:"created_at,omitempty"`
	UpdatedAt   int64  `json:"updated_at,omitempty"`
	URL         string `json:"url,omitempty"`
	Order       int64  `json:"order,omitempty"`
}

// SectionList holds a list of Sections and pagination info.
type SectionList struct {
	Pages    PageParams `json:"pages,omitempty"`
	TotalCount int64    `json:"total_count,omitempty"`
	Sections []Section  `json:"data,omitempty"`
}

// List all Sections for the App.
func (s *SectionService) List(params PageParams) (SectionList, error) {
	return s.Repository.list(sectionListParams{PageParams: params})
}

// Find a particular Section by ID.
func (s *SectionService) Find(id string) (Section, error) {
	return s.Repository.find(id)
}

func (s Section) String() string {
	return fmt.Sprintf("[intercom] section { id: %s, name: %s }", s.ID, s.Name)
}
