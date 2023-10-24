package entity

type (
	Organization struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		DisplayName string `json:"displayName"`
		Description string `json:"description"`
		Version     int64  `json:"version"`
		Deleted     bool   `json:"deleted"`
	}
)
