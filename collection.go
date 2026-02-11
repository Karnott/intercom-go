package intercom

import "fmt"

// CollectionService handles interactions with the API through a CollectionRepository.
type CollectionService struct {
	Repository CollectionRepository
}

// Collection represents a Help Center Collection in Intercom.
type Collection struct {
	ID          string `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
	WorkspaceID string `json:"workspace_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt   int64  `json:"created_at,omitempty"`
	UpdatedAt   int64  `json:"updated_at,omitempty"`
	URL         string `json:"url,omitempty"`
	Icon        string `json:"icon,omitempty"`
	Order       int64  `json:"order,omitempty"`
}

// CollectionList holds a list of Collections and pagination info.
type CollectionList struct {
	Pages       PageParams   `json:"pages,omitempty"`
	TotalCount  int64        `json:"total_count,omitempty"`
	Collections []Collection `json:"data,omitempty"`
}

// List all Collections for the App.
func (s *CollectionService) List(params PageParams) (CollectionList, error) {
	return s.Repository.list(collectionListParams{PageParams: params})
}

// Find a particular Collection by ID.
func (s *CollectionService) Find(id string) (Collection, error) {
	return s.Repository.find(id)
}

func (c Collection) String() string {
	return fmt.Sprintf("[intercom] collection { id: %s, name: %s }", c.ID, c.Name)
}
