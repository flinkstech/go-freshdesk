package freshdesk

import (
	"fmt"
	"time"
)

type AgentManager interface {
	All() (AgentSlice, error)
}

type agentManager struct {
	client *ApiClient
}

func newAgentManager(client *ApiClient) agentManager {
	return agentManager{
		client,
	}
}

type Agent struct {
	Available      bool      `json:"available"`
	AvailableSince time.Time `json:"available_since"`
	ID             int       `json:"id"`
	Occasional     bool      `json:"occasional"`
	Signature      string    `json:"signature"`
	TicketScope    int       `json:"ticket_scope"`
	GroupIDs       []int     `json:"group_ids"`
	RoleIDs        []int     `json:"role_ids"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Contact        Contact   `json:"contact"`
}

type AgentSlice []Agent

func (s AgentSlice) Len() int { return len(s) }

func (s AgentSlice) Less(i, j int) bool { return s[i].ID < s[j].ID }

func (s AgentSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s AgentSlice) Print() {
	for _, agent := range s {
		fmt.Println(agent.Contact.Name)
	}
}

func (manager agentManager) All() (AgentSlice, error) {
	output := AgentSlice{}
	headers, err := manager.client.get(endpoints.agents.all, &output)
	if err != nil {
		return AgentSlice{}, err
	}
	nextSlice := AgentSlice{}
	for {
		nextLink := manager.client.getNextLink(headers)
		if nextLink == "" {
			break
		}
		headers, err = manager.client.get(nextLink, &nextSlice)
		if err != nil {
			return AgentSlice{}, err
		}
		output = append(output, nextSlice...)
	}
	return output, nil
}
