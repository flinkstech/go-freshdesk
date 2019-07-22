package freshdesk

import "fmt"

type agentEndpoints struct {
	all string
	me  string
}
type articleEndpoints struct {
	delete func(int) string
	get    func(int) string
}

type categoryEndpoints struct {
	folders func(int) string
}

type companyEndpoints struct {
	all    string
	create string
	update func(int) string
}

type contactEndpoints struct {
	all    string
	create string
	update func(int) string
	search func(string) string
}

type folderEndpoints struct {
	articles func(int) string
}

type groupEndpoints struct {
	all string
}

type slaPolicyEndpoints struct {
	all    string
	update func(int) string
}

type solutionEndpoints struct {
	categories string
	category   categoryEndpoints
	folder     folderEndpoints
	articles   articleEndpoints
}

type ticketEndpoints struct {
	all    string
	create string
	view   func(int) string
	search func(string) string
}

var endpoints = struct {
	agents      agentEndpoints
	companies   companyEndpoints
	contacts    contactEndpoints
	groups      groupEndpoints
	slaPolicies slaPolicyEndpoints
	solutions   solutionEndpoints
	tickets     ticketEndpoints
}{
	agents: agentEndpoints{
		all: "/api/v2/agents",
		me:  "/api/v2/agents/me",
	},
	companies: companyEndpoints{
		all:    "/api/v2/companies",
		create: "/api/v2/companies",
		update: func(id int) string { return fmt.Sprintf("/api/v2/companies/%d", id) },
	},
	contacts: contactEndpoints{
		all:    "/api/v2/contacts",
		create: "/api/v2/contacts",
		update: func(id int) string { return fmt.Sprintf("/api/v2/contacts/%d", id) },
		search: func(query string) string { return fmt.Sprintf("/api/v2/search/contacts?%s", query) },
	},
	groups: groupEndpoints{
		all: "/api/v2/groups",
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
		view:   func(id int) string { return fmt.Sprintf("/api/v2/tickets/%d", id) },
		search: func(query string) string { return fmt.Sprintf("/api/v2/search/tickets?%s", query) },
	},
}
