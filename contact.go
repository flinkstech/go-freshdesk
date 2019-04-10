package freshdesk

import "time"

type Contact struct {
	Name        string     `json:"name,omitempty"`
	Active      string     `json:"active,omitempty"`
	Email       string     `json:"email,omitempty"`
	JobTitle    string     `json:"job_title,omitempty"`
	Language    string     `json:"language,omitempty"`
	LastLoginAt *time.Time `json:"last_login_at,omitempty"`
	Mobile      int        `json:"mobile,omitempty"`
	Phone       int        `json:"phone,omitempty"`
	TimeZone    string     `json:"time_zone,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}
