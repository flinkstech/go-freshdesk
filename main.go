package freshdesk

type apiClient struct {
	domain    string
	apiKey    string
	Companies CompanyManager
	Tickets   TicketManager
}

// Init initializes the package
func Init(domain, apiKey string) apiClient {
	client := apiClient{
		domain: domain,
		apiKey: apiKey,
	}
	client.Companies = newCompanyManager(&client)
	client.Tickets = newTicketManager(&client)
	return client
}
