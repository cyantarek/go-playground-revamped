package domains

import "time"

type Code struct {
	Body      string    `json:"body"`
	ShortCode string    `json:"short_code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CodeService interface {
	SaveCode(shortCode string, c Code) error
	GetCode(shortCode string) *Code
}
