package freshdesk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TicketManager interface {
	All() (TicketSlice, error)
	Create(CreateTicket) (Ticket, error)
}

type ticketManager struct {
	client *apiClient
}

func newTicketManager(client *apiClient) ticketManager {
	return ticketManager{
		client,
	}
}

type Ticket struct {
	ID                     int                    `json:"id"`
	Subject                string                 `json:"subject"`
	Type                   string                 `json:"type"`
	Description            string                 `json:"description"`
	Attachments            []interface{}          `json:"attachments"`
	CCEmails               []string               `json:"cc_emails"`
	CompanyID              int                    `json:"company_id"`
	Deleted                bool                   `json:"deleted"`
	DescriptionText        string                 `json:"description_text"`
	DueBy                  time.Time              `json:"due_by"`
	Email                  string                 `json:"email"`
	EmailConfigID          int                    `json:"email_config_id"`
	FacebookID             string                 `json:"facebook_id"`
	FirstResponseDueBy     time.Time              `json:"fr_due_by"`
	FirstResponseEscalated bool                   `json:"fr_escalated"`
	FwdEmails              []string               `json:"fwd_emails"`
	GroupID                int                    `json:"group_id"`
	IsEscalated            bool                   `json:"is_escalated"`
	Name                   string                 `json:"name"`
	Phone                  string                 `json:"phone"`
	Priority               int                    `json:"priority"`
	ProductID              int                    `json:"product_id"`
	ReplyCCEmails          []string               `json:"reply_cc_emails"`
	RequesterID            int                    `json:"requester_id"`
	ResponderID            int                    `json:"responder_id"`
	Source                 int                    `json:"source"`
	Spam                   bool                   `json:"spam"`
	Status                 int                    `json:"status"`
	Tags                   []string               `json:"tags"`
	ToEmails               []string               `json:"to_emails"`
	TwitterID              string                 `json:"twitter_id"`
	CreatedAt              time.Time              `json:"created_at"`
	UpdatedAt              time.Time              `json:"updated_at"`
	CustomFields           map[string]interface{} `json:"custom_fields"`
}

type CreateTicket struct {
	Name               string                 `json:"name,omitempty"`
	RequesterID        int                    `json:"requester_id,omitempty"`
	Email              string                 `json:"email,omitempty"`
	FacebookID         string                 `json:"facebook_id,omitempty"`
	Phone              string                 `json:"phone,omitempty"`
	TwitterID          string                 `json:"twitter_id,omitempty"`
	UniqueExternalID   string                 `json:"unique_external_id,omitempty"`
	Subject            string                 `json:"subject,omitempty"`
	Type               string                 `json:"type,omitempty"`
	Status             int                    `json:"status,omitempty"`
	Priority           int                    `json:"priority,omitempty"`
	Description        string                 `json:"description,omitempty"`
	ResponderID        int                    `json:"responder_id,omitempty"`
	Attachments        []interface{}          `json:"attachments,omitempty"`
	CCEmails           []string               `json:"cc_emails,omitempty"`
	CustomFields       map[string]interface{} `json:"custom_fields,omitempty"`
	DueBy              *time.Time             `json:"due_by,omitempty"`
	EmailConfigID      int                    `json:"email_config_id,omitempty"`
	FirstResponseDueBy *time.Time             `json:"fr_due_by,omitempty"`
	GroupID            int                    `json:"group_id,omitempty"`
	ProductID          int                    `json:"product_id,omitempty"`
	Source             int                    `json:"source,omitempty"`
	Tags               []string               `json:"tags,omitempty"`
	CompanyID          int                    `json:"company_id,omitempty"`
}

type Source int
type Status int
type Priority int

const (
	SourceEmail Source = 1 + iota
	SourcePortal
	SourcePhone
	_
	_
	_
	SourceChat
	SourceMobihelp
	SourceFeedbackWidget
	SourceOutboundEmail
)

const (
	StatusOpen Status = 2 + iota
	StatusPending
	StatusResolved
	StatusClosed
)

const (
	PriorityLow Priority = 1 + iota
	PriorityMedium
	PriorityHigh
	PriorityUrgent
)

func (s Source) Value() int {
	return int(s)
}

func (s Status) Value() int {
	return int(s)
}

func (p Priority) Value() int {
	return int(p)
}

func (t Ticket) Print() {
	jsonb, _ := json.MarshalIndent(t, "", "    ")
	fmt.Println(string(jsonb))
}

type TicketSlice []Ticket

func (s TicketSlice) Len() int { return len(s) }

func (s TicketSlice) Less(i, j int) bool { return s[i].ID < s[j].ID }

func (s TicketSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s TicketSlice) Print() {
	for _, ticket := range s {
		fmt.Println(ticket.Subject)
	}
}

func (manager ticketManager) All() (TicketSlice, error) {
	output := TicketSlice{}
	headers, err := manager.client.get(endpoints.tickets.all, &output)
	if err != nil {
		return TicketSlice{}, err
	}
	for {
		if nextPage, ok := manager.client.getNextLink(headers); ok {
			nextSlice := TicketSlice{}
			headers, err = manager.client.get(nextPage, &nextSlice)
			if err != nil {
				return TicketSlice{}, err
			}
			output = append(output, nextSlice...)
			continue
		}
		break
	}
	return output, nil
}

func (manager ticketManager) Create(ticket CreateTicket) (Ticket, error) {
	output := Ticket{}
	jsonb, err := json.Marshal(ticket)
	if err != nil {
		return output, nil
	}
	err = manager.client.postJSON(endpoints.tickets.create, jsonb, &output, http.StatusCreated)
	if err != nil {
		return Ticket{}, err
	}
	return output, nil
}