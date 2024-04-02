package utils

// PaginatedResponse represents the structure of the paginated response
type PaginatedResponse struct {
	Data       []*Data `json:"data"`
	NextToken  string  `json:"next_token,omitempty"`
	TotalCount int     `json:"total_count"`
}
