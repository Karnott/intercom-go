package intercom

// PageParams determine paging information to and from the API
type PageParams struct {
	Page       int64 `json:"page" url:"page,omitempty"`
	PerPage    int64 `json:"per_page" url:"per_page,omitempty"`
	TotalPages int64 `json:"total_pages" url:"-"`
}

// CursorPages holds cursor-based pagination info returned by the API (v2.0+).
type CursorPages struct {
	Type       string              `json:"type,omitempty"`
	Page       int64               `json:"page,omitempty"`
	PerPage    int64               `json:"per_page,omitempty"`
	TotalPages int64               `json:"total_pages,omitempty"`
	Next       *StartingAfterParam `json:"next,omitempty"`
}

// StartingAfterParam holds the cursor for the next page.
type StartingAfterParam struct {
	PerPage       int64  `json:"per_page,omitempty"`
	StartingAfter string `json:"starting_after,omitempty"`
}
