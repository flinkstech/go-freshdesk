package main

import (
	"log"
	"os"

	freshdesk "github.com/flinkstech/go-freshdesk"
)

func main() {
	logger := log.New(os.Stdout, "[logger] ", 0)

	// domain refresh to what precedes .freshdesk.com
	// eg: For test.freshdesk.com, the "domain" would be "test"
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

	note, err := client.Tickets.AddNote(ticket.ID, freshdesk.AddNote{Body: "This is a note"})
	if err != nil {
		panic(err)
	}
	note.Print()

	reply, err := client.Tickets.AddReply(ticket.ID, freshdesk.AddReply{Body: "This is a reply"})
	if err != nil {
		panic(err)
	}
	reply.Print()
}
