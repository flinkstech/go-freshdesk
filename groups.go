package freshdesk

import (
	"fmt"
	"time"
)

type GroupManager interface {
	All() (CompanySlice, error)
}

type groupManager struct {
	client *ApiClient
}

func newGroupManager(client *ApiClient) groupManager {
	return groupManager{
		client,
	}
}

type Group struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	BusinessHourID   string    `json:"business_hour_id"`
	AgentIDs         []int     `json:"agent_ids"`
	AutoTicketAssign bool      `json:"auto_ticket_assign"`
	EscalateTo       int       `json:"escalate_to"`
	UnassignedFor    string    `json:"unassigned_for"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type GroupSlice []Group

func (s GroupSlice) Len() int { return len(s) }

func (s GroupSlice) Less(i, j int) bool { return s[i].ID < s[j].ID }

func (s GroupSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s GroupSlice) Print() {
	for _, group := range s {
		fmt.Println(group.Name)
	}
}

func (manager groupManager) All() (GroupSlice, error) {
	output := GroupSlice{}
	_, err := manager.client.get(endpoints.groups.all, &output)
	if err != nil {
		return GroupSlice{}, err
	}
	return output, nil
}
