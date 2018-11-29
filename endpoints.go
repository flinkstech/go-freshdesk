package freshdesk

type companyEndpoints struct {
	all string
}

type ticketEndpoints struct {
	all    string
	create string
}

var endpoints = struct {
	companies companyEndpoints
	tickets   ticketEndpoints
}{
	companies: companyEndpoints{
		all: "/api/v2/companies",
	},
	tickets: ticketEndpoints{
		all:    "/api/v2/tickets",
		create: "/api/v2/tickets",
	},
}
