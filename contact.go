package freshdesk

import "time"

type Contact struct {
	Name        string    `json:"name"`
	Active      string    `json:"active"`
	Email       string    `json:"email"`
	JobTitle    string    `json:"job_title"`
	Language    string    `json:"language"`
	LastLoginAt time.Time `json:"last_login_at"`
	Mobile      int       `json:"mobile"`
	Phone       int       `json:"phone"`
	TimeZone    string    `json:"time_zone"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
