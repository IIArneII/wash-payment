package entity

type (
	WashServer struct {
		ID          string `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		GroupID     string `json:"groupId"`
		Deleted     bool   `json:"deleted"`
		Version     int64  `json:"version"`
	}
)
