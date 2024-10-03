package entity

type (
	Group struct {
		ID             string `json:"id"`
		OrganizationID string `json:"organizationId"`
		Name           string `json:"name"`
		Description    string `json:"description"`
		Version        int    `json:"version"`
		Deleted        bool   `json:"deleted"`
	}
)
