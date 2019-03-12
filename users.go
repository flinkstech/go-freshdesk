package freshdesk

import (
	"fmt"
	"time"

	"github.com/flinkstech/go-freshdesk/querybuilder"
)

type UserManager interface {
	All() (UserSlice, error)
	Search(querybuilder.Query) (UserResults, error)
}

type userManager struct {
	client *ApiClient
}

func newUserManager(client *ApiClient) userManager {
	return userManager{
		client,
	}
}

type UserResults struct {
	next    string
	Results UserSlice
	client  *ApiClient
}

type User struct {
	ID               int                    `json:"id"`
	Name             string                 `json:"name"`
	Active           string                 `json:"active"`
	Email            string                 `json:"email"`
	JobTitle         string                 `json:"job_title"`
	Language         string                 `json:"language"`
	LastLoginAt      time.Time              `json:"last_login_at"`
	Mobile           int                    `json:"mobile"`
	Phone            int                    `json:"phone"`
	TimeZone         string                 `json:"time_zone"`
	CreatedAt        time.Time              `json:"created_at"`
	UpdatedAt        time.Time              `json:"updated_at"`
	Address          string                 `json:"address"`
	Avatar           interface{}            `json:"avatar"`
	CompanyID        int                    `json:"company_id"`
	ViewAllTickets   bool                   `json:"view_all_tickets"`
	CustomFields     map[string]interface{} `json:"custom_fields"`
	Deleted          bool                   `json:"deleted"`
	Description      string                 `json:"description"`
	OtherEmails      []string               `json:"other_emails"`
	Tags             []string               `json:"tags"`
	TwitterID        string                 `json:"twitter_id"`
	UniqueExternalID string                 `json:"unique_external_id"`
	OtherCompanies   []string               `json:"other_companies"`
}

type UserSlice []User

func (s UserSlice) Len() int { return len(s) }

func (s UserSlice) Less(i, j int) bool { return s[i].ID < s[j].ID }

func (s UserSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s UserSlice) Print() {
	for _, user := range s {
		fmt.Println(user.Name)
	}
}

func (manager userManager) All() (UserSlice, error) {
	output := UserSlice{}
	headers, err := manager.client.get(endpoints.contacts.all, &output)
	if err != nil {
		return UserSlice{}, err
	}
	for {
		nextLink := manager.client.getNextLink(headers)
		if nextLink == "" {
			break
		}
		nextSlice := UserSlice{}
		headers, err = manager.client.get(nextLink, &nextSlice)
		if err != nil {
			return UserSlice{}, err
		}
		output = append(output, nextSlice...)
	}
	return output, nil
}

func (manager userManager) Search(query querybuilder.Query) (UserResults, error) {
	output := struct {
		Slice UserSlice `json:"results"`
	}{}
	headers, err := manager.client.get(endpoints.contacts.search(query.URLSafe()), &output)
	if err != nil {
		return UserResults{}, err
	}
	return UserResults{
		Results: output.Slice,
		client:  manager.client,
		next:    manager.client.getNextLink(headers),
	}, nil
}
