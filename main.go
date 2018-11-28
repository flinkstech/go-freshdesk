package freshdesk

type Client interface {
	GetCompanies() (CompanySlice, error)
}

func Init(domain, apiKey string) Client {
	return client{
		Domain: domain,
		APIKey: apiKey,
	}
}

type client struct {
	Domain string
	APIKey string
}
