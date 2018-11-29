package freshdesk

import (
	"fmt"
	"time"
)

type TicketManager interface {
	All() (TicketSlice, error)
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
