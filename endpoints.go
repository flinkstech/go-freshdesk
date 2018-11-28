package freshdesk

type companyEndpoints struct {
	All string
}

var endpoints = struct {
	Companies companyEndpoints
}{
	Companies: companyEndpoints{
		All: "/api/v2/companies",
	},
}
