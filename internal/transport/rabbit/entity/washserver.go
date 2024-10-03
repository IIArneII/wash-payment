package entity

type (
	WashServer struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		GroupID     string `json:"groupId"`
		OwnerID     string `json:"ownerId"`
		Deleted     bool   `json:"deleted"`
		Version     int    `json:"version"`
	}
)
