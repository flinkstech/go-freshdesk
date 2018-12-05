package main

import (
	"log"
	"os"

	freshdesk "github.com/flinkstech/go-freshdesk"
)

func main() {
	logger := log.New(os.Stdout, "[logger] ", 0)

	client := freshdesk.Init("domain", "apikey", &freshdesk.ClientOptions{Logger: logger}) // Or use freshdesk.EmptyOptions()

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
