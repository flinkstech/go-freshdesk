package freshdesk

import (
	"fmt"
	"time"
)

type CompanyManager interface {
	All() (CompanySlice, error)
}

type companyManager struct {
	client *apiClient
}

func newCompanyManager(client *apiClient) companyManager {
	return companyManager{
		client,
	}
}

type Company struct {
	ID           int                    `json:"id"`
	Name         string                 `json:"name"`
	Description  string                 `json:"description"`
	Domains      []string               `json:"domains"`
	Note         string                 `json:"note"`
	HealthScore  string                 `json:"health_score"`
	AccountTier  string                 `json:"account_tier"`
	RenewalDate  time.Time              `json:"renewal_date"`
	Industry     string                 `json:"industry"`
	CreatedAt    time.Time              `json:"created_at"`
	UpdatedAt    time.Time              `json:"updated_at"`
	CustomFields map[string]interface{} `json:"custom_fields"`
}

type CompanySlice []Company

func (c CompanySlice) Len() int {
	return len(c)
}

func (c CompanySlice) Less(i, j int) bool {
	return c[i].ID < c[j].ID
}

func (c CompanySlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c CompanySlice) Print() {
	for _, company := range c {
		fmt.Println(company.Name)
	}
}

func (manager companyManager) All() (CompanySlice, error) {
	output := CompanySlice{}
	headers, err := manager.client.get(endpoints.companies.all, &output)
	if err != nil {
		return CompanySlice{}, err
	}
	for {
		if nextPage, ok := manager.client.getNextLink(headers); ok {
			nextSlice := CompanySlice{}
			headers, err = manager.client.get(nextPage, &nextSlice)
			if err != nil {
				return CompanySlice{}, err
			}
			output = append(output, nextSlice...)
			continue
		}
		break
	}
	return output, nil
}
