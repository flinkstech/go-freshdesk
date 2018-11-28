package freshdesk

import (
	"fmt"
	"time"
)

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

func (c client) GetCompanies() (CompanySlice, error) {
	output := CompanySlice{}
	headers, err := c.get(endpoints.Companies.All, &output)
	if err != nil {
		return CompanySlice{}, err
	}
	for {
		if nextPage, ok := getNextLink(headers); ok {
			nextSlice := CompanySlice{}
			headers, err = c.get(nextPage, &nextSlice)
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
