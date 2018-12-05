package freshdesk

import "fmt"

type companyEndpoints struct {
	all string
}

type ticketEndpoints struct {
	all    string
	create string
}

type slaPolicyEndpoints struct {
	all    string
	update func(int) string
}

var endpoints = struct {
	companies   companyEndpoints
	slaPolicies slaPolicyEndpoints
	tickets     ticketEndpoints
}{
	companies: companyEndpoints{
		all: "/api/v2/companies",
	},
	slaPolicies: slaPolicyEndpoints{
		all:    "/api/v2/sla_policies",
		update: func(id int) string { return fmt.Sprintf("/api/v2/sla_policies/%d", id) },
	},
	tickets: ticketEndpoints{
		all:    "/api/v2/tickets",
		create: "/api/v2/tickets",
	},
}
