package freshdesk

import (
	"log"
)

type apiClient struct {
	domain      string
	apiKey      string
	logger      *log.Logger
	Companies   CompanyManager
	Tickets     TicketManager
	SLAPolicies SLAPolicyManager
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
	client.Tickets = newTicketManager(&client)
	client.SLAPolicies = newSLAPolicyManager(&client)
	return client
}
