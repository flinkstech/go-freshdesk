package freshdesk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SLAPolicyManager interface {
	All() (SLAPolicySlice, error)
	Update(int, SLAPolicy) (SLAPolicy, error)
}

type slaPolicyManager struct {
	client *apiClient
}

func newSLAPolicyManager(client *apiClient) slaPolicyManager {
	return slaPolicyManager{
		client,
	}
}

type SLAPolicy struct {
	ID           int              `json:"id,omitempty"`
	Name         string           `json:"name,omitempty"`
	Description  string           `json:"description,omitempty"`
	ApplicableTo map[string][]int `json:"applicable_to,omitempty"`
	Active       bool             `json:"active,omitempty"`
	IsDefault    bool             `json:"is_default,omitempty"`
	Position     int              `json:"position,omitempty"`
	client       *apiClient
}

type SLAPolicyApplicableCompanyList struct {
	ApplicableTo ApplicableTo `json:"applicable_to,omitempty"`
}

type ApplicableTo struct {
	CompanyIDs []int `json:"company_ids"`
}

type SLAPolicySlice []SLAPolicy

func (slice SLAPolicySlice) Len() int {
	return len(slice)
}

func (slice SLAPolicySlice) Less(i, j int) bool {
	return slice[i].ID < slice[j].ID
}

func (slice SLAPolicySlice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (slice SLAPolicySlice) Print() {
	for _, policy := range slice {
		fmt.Println(policy.Name)
	}
}

func (manager slaPolicyManager) All() (SLAPolicySlice, error) {
	output := SLAPolicySlice{}
	_, err := manager.client.get(endpoints.slaPolicies.all, &output)
	if err != nil {
		return SLAPolicySlice{}, err
	}
	outputWithClient := SLAPolicySlice{}
	for _, policy := range output {
		policy.client = manager.client
		outputWithClient = append(outputWithClient, policy)
	}
	return outputWithClient, nil
}

func (manager slaPolicyManager) Update(id int, policy SLAPolicy) (SLAPolicy, error) {
	output := SLAPolicy{}
	jsonb, err := json.Marshal(policy)
	if err != nil {
		return output, nil
	}
	err = manager.client.put(endpoints.slaPolicies.update(id), jsonb, &output, http.StatusOK)
	if err != nil {
		return SLAPolicy{}, err
	}
	return output, nil
}

// EnsureCompanyPresent indempotently ensures an SLA policy is applied to a Company
func (policy SLAPolicy) EnsureCompanyPresent(companyID int) {
	newCompanyIDs := []int{}
	for key, values := range policy.ApplicableTo {
		if key == "company_ids" {
			for _, id := range values {
				if id == companyID {
					return
				}
				newCompanyIDs = append(newCompanyIDs, id)
			}
			break
		}
	}

	changes := &SLAPolicyApplicableCompanyList{}
	changes.ApplicableTo.CompanyIDs = append(newCompanyIDs, companyID)
	jsonb, err := json.Marshal(changes)
	if err != nil && policy.client.logger != nil {
		policy.client.logger.Println(err.Error())
	}
	output := SLAPolicy{}
	policy.client.put(endpoints.slaPolicies.update(policy.ID), jsonb, &output, http.StatusOK)
	jsonb, err = json.MarshalIndent(output, "", "\t")
	if err != nil {
		if policy.client.logger != nil {
			policy.client.logger.Println(err.Error())
		}
	} else {
		if policy.client.logger != nil {
			policy.client.logger.Println("hello", string(jsonb))
		}
	}
}

// EnsureCompanyAbsent indempotently ensures an SLA policy is applied to a Company
func (policy SLAPolicy) EnsureCompanyAbsent(companyID int) {
	// Check for the company_ids key and skip if it is not present
	mapContainsCompaniesFlag := false
	for key := range policy.ApplicableTo {
		if key == "company_ids" {
			mapContainsCompaniesFlag = true
		}
	}
	if !mapContainsCompaniesFlag {
		return
	}
	// Check for the company ID in question and skip if it is not present, collecting the other IDs along the way
	newCompanyIDs := []int{}
	companyIsPresentFlag := false
	for _, id := range policy.ApplicableTo["company_ids"] {
		if id == companyID {
			companyIsPresentFlag = true
			continue
		}
		newCompanyIDs = append(newCompanyIDs, id)
	}
	if !companyIsPresentFlag {
		return
	}

	changes := &SLAPolicyApplicableCompanyList{}
	changes.ApplicableTo.CompanyIDs = newCompanyIDs
	jsonb, err := json.Marshal(changes)
	if err != nil && policy.client.logger != nil {
		policy.client.logger.Println(err.Error())
	}
	output := SLAPolicy{}
	policy.client.put(endpoints.slaPolicies.update(policy.ID), jsonb, &output, http.StatusOK)
	jsonb, err = json.MarshalIndent(output, "", "\t")
	if err != nil {
		if policy.client.logger != nil {
			policy.client.logger.Println(err.Error())
		}
	} else {
		if policy.client.logger != nil {
			policy.client.logger.Println("hello", string(jsonb))
		}
	}
}
