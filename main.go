package freshdesk

import (
	"log"
)

type ApiClient struct {
	domain      string
	apiKey      string
	logger      *log.Logger
	Companies   CompanyManager
	Groups      GroupManager
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
func Init(domain, apiKey string, options *ClientOptions) ApiClient {
	client := ApiClient{
		domain: domain,
		apiKey: apiKey,
	}
	if options != nil {
		client.logger = options.Logger
	}
	client.Companies = newCompanyManager(&client)
	client.Groups = newGroupManager(&client)
	client.SLAPolicies = newSLAPolicyManager(&client)
	client.Solutions = newSolutionManager(&client)
	client.Tickets = newTicketManager(&client)
	return client
}
