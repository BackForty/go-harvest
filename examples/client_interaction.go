package main

import (
	"../harvest"
	"fmt"
)

func main() {
	apiClient := harvest.NewAPIClientWithBasicAuth(
		"YOUR_USERNAME",
		"YOUR_PASSWORD",
		"YOUR_SUBDOMAIN")

	err, clients := apiClient.Client.List()
	if err != nil {
		fmt.Printf("error getting clients")
	}

	fmt.Printf("%+v\n\n\n", clients)

	err, client := apiClient.Client.Find(clients[0].Id)
	if err != nil {
		fmt.Printf("error getting individual client")
	}
	fmt.Printf("%v\n\n\n", client)

	err, people := apiClient.People.List()
	if err != nil {
		fmt.Printf("error getting people listing")
	}
	fmt.Printf("%+v\n\n\n", people)

	err, person := apiClient.People.Find(people[0].Id)
	if err != nil {
		fmt.Printf("error getting individual person")
	}
	fmt.Printf("%v\n\n\n", person)
}
