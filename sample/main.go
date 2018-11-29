package main

import (
	freshdesk "github.com/flinkstech/go-freshdesk"
)

func main() {
	client := freshdesk.Init("domain", "apikey")

	companies, err := client.Companies.All()
	if err != nil {
		panic(err)
	}
	companies.Print()

	tickets, err := client.Tickets.All()
	if err != nil {
		panic(err)
	}
	tickets.Print()

	ticket, err := client.Tickets.Create(freshdesk.CreateTicket{
		Subject:     "Ticket Subject",
		Description: "Ticket description.",
		Email:       "identifier@domain.tld",
		Status:      freshdesk.StatusOpen.Value(),
		Priority:    freshdesk.PriorityLow.Value(),
	})
	if err != nil {
		panic(err)
	}
	ticket.Print()
}
