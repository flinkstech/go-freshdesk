package freshdesk

import "fmt"

type companyEndpoints struct {
	all    string
	create string
}

type ticketEndpoints struct {
	all    string
	create string
}

type slaPolicyEndpoints struct {
	all    string
	update func(int) string
}

type categoryEndpoints struct {
	folders func(int) string
}

type folderEndpoints struct {
	articles func(int) string
}

type articleEndpoints struct {
	delete func(int) string
	get    func(int) string
}

type solutionEndpoints struct {
	categories string
	category   categoryEndpoints
	folder     folderEndpoints
	articles   articleEndpoints
}

var endpoints = struct {
	companies   companyEndpoints
	slaPolicies slaPolicyEndpoints
	solutions   solutionEndpoints
	tickets     ticketEndpoints
}{
	companies: companyEndpoints{
		all:    "/api/v2/companies",
		create: "/api/v2/companies",
	},
	slaPolicies: slaPolicyEndpoints{
		all:    "/api/v2/sla_policies",
		update: func(id int) string { return fmt.Sprintf("/api/v2/sla_policies/%d", id) },
	},
	solutions: solutionEndpoints{
		categories: "/api/v2/solutions/categories",
		category: categoryEndpoints{
			folders: func(id int) string { return fmt.Sprintf("/api/v2/solutions/categories/%d/folders", id) },
		},
		folder: folderEndpoints{
			articles: func(id int) string { return fmt.Sprintf("/api/v2/solutions/folders/%d/articles", id) },
		},
		articles: articleEndpoints{
			delete: func(id int) string { return fmt.Sprintf("/api/v2/solutions/articles/%d", id) },
			get:    func(id int) string { return fmt.Sprintf("/api/v2/solutions/articles/%d", id) }, // Not currently in use
		},
	},
	tickets: ticketEndpoints{
		all:    "/api/v2/tickets",
		create: "/api/v2/tickets",
	},
}
