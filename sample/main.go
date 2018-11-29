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
}
