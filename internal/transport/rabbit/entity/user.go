package entity

type (
	User struct {
		ID             string  `json:"id"`
		Email          string  `json:"email"`
		Name           string  `json:"name"`
		Role           string  `json:"role"`
		OrganizationID *string `json:"organizationId"`
		Version        int64   `json:"version"`
	}
)
