package freshdesk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/flinkstech/go-freshdesk/querybuilder"
)

type UserManager interface {
	All() (UserSlice, error)
	Create(*User) (*User, error)
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
	ID               int                    `json:"id,omitempty"`
	Name             string                 `json:"name,omitempty"`
	Active           string                 `json:"active,omitempty"`
	Email            string                 `json:"email,omitempty"`
	JobTitle         string                 `json:"job_title,omitempty"`
	Language         string                 `json:"language,omitempty"`
	LastLoginAt      time.Time              `json:"last_login_at,omitempty"`
	Mobile           int                    `json:"mobile,omitempty"`
	Phone            int                    `json:"phone,omitempty"`
	TimeZone         string                 `json:"time_zone,omitempty"`
	CreatedAt        time.Time              `json:"created_at,omitempty"`
	UpdatedAt        time.Time              `json:"updated_at,omitempty"`
	Address          string                 `json:"address,omitempty"`
	Avatar           interface{}            `json:"avatar,omitempty"`
	CompanyID        int                    `json:"company_id,omitempty"`
	ViewAllTickets   bool                   `json:"view_all_tickets,omitempty"`
	CustomFields     map[string]interface{} `json:"custom_fields,omitempty"`
	Deleted          bool                   `json:"deleted,omitempty"`
	Description      string                 `json:"description,omitempty"`
	OtherEmails      []string               `json:"other_emails,omitempty"`
	Tags             []string               `json:"tags,omitempty"`
	TwitterID        string                 `json:"twitter_id,omitempty"`
	UniqueExternalID string                 `json:"unique_external_id,omitempty"`
	OtherCompanies   []string               `json:"other_companies,omitempty"`
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
		Slice UserSlice `json:"results,omitempty"`
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

func (manager userManager) Create(user *User) (*User, error) {
	output := &User{}
	jsonb, err := json.Marshal(user)
	if err != nil {
		return output, err
	}
	err = manager.client.postJSON(endpoints.contacts.create, jsonb, &output, http.StatusCreated)
	if err != nil {
		return &User{}, err
	}
	return output, nil
}

func (manager userManager) Update(id int, user *User) (*User, error) {
	output := &User{}
	jsonb, err := json.Marshal(user)
	if err != nil {
		return output, err
	}
	err = manager.client.put(endpoints.contacts.update(id), jsonb, &output, http.StatusOK)
	if err != nil {
		return &User{}, err
	}
	return output, nil
}
