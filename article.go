package intercom

import "fmt"

// ArticleService handles interactions with the API through an ArticleRepository.
type ArticleService struct {
	Repository ArticleRepository
}

// Article represents a Help Center Article in Intercom.
type Article struct {
	ID          string            `json:"id,omitempty"`
	Type        string            `json:"type,omitempty"`
	WorkspaceID string            `json:"workspace_id,omitempty"`
	Title       string            `json:"title,omitempty"`
	Description string            `json:"description,omitempty"`
	Body        string            `json:"body,omitempty"`
	AuthorID    int64             `json:"author_id,omitempty"`
	State       string            `json:"state,omitempty"`
	CreatedAt   int64             `json:"created_at,omitempty"`
	UpdatedAt   int64             `json:"updated_at,omitempty"`
	URL         string            `json:"url,omitempty"`
	ParentID    int64             `json:"parent_id,omitempty"`
	ParentType  string            `json:"parent_type,omitempty"`
	ParentIDs   []int64           `json:"parent_ids,omitempty"`
	Statistics  ArticleStatistics `json:"statistics,omitempty"`
}

// ArticleStatistics holds metrics for an Article.
type ArticleStatistics struct {
	Views                      int64   `json:"views,omitempty"`
	Conversations              int64   `json:"conversations,omitempty"`
	Reactions                  int64   `json:"reactions,omitempty"`
	HappyReactionPercentage    float64 `json:"happy_reaction_percentage,omitempty"`
	NeutralReactionPercentage  float64 `json:"neutral_reaction_percentage,omitempty"`
	SadReactionPercentage      float64 `json:"sad_reaction_percentage,omitempty"`
}

// ArticleList holds a list of Articles and pagination info.
type ArticleList struct {
	Pages      PageParams `json:"pages,omitempty"`
	TotalCount int64      `json:"total_count,omitempty"`
	Articles   []Article  `json:"data,omitempty"`
}

// List all Articles for the App.
func (s *ArticleService) List(params PageParams) (ArticleList, error) {
	return s.Repository.list(articleListParams{PageParams: params})
}

// Find a particular Article by ID.
func (s *ArticleService) Find(id string) (Article, error) {
	return s.Repository.find(id)
}

func (a Article) String() string {
	return fmt.Sprintf("[intercom] article { id: %s, title: %s, state: %s }", a.ID, a.Title, a.State)
}
