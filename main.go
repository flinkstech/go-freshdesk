package freshdesk

import (
	"log"
)

type apiClient struct {
	domain      string
	apiKey      string
	logger      *log.Logger
	Companies   CompanyManager
	SLAPolicies SLAPolicyManager
	Solutions   SolutionManager
	Tickets     TicketManager
}

type ClientOptions struct {
	Logger *log.Logger
}

func EmptyOptions() *ClientOptions {
	return nil
}

// Init initializes the package
func Init(domain, apiKey string, options *ClientOptions) apiClient {
	client := apiClient{
		domain: domain,
		apiKey: apiKey,
	}
	if options != nil {
		client.logger = options.Logger
	}
	client.Companies = newCompanyManager(&client)
	client.SLAPolicies = newSLAPolicyManager(&client)
	client.Solutions = newSolutionManager(&client)
	client.Tickets = newTicketManager(&client)
	return client
}
