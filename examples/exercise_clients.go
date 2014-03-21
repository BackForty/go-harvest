package examples

import (
	"../harvest"
	"fmt"
)

func ExerciseClients(apiClient *harvest.APIClient) {
	err, clients := apiClient.Client.List()
	if err != nil {
		fmt.Printf("\n%v\n", err)
	} else {
		fmt.Printf("%+v\n\n\n", clients)
		err, client := apiClient.Client.Find(clients[0].Id)
		if err != nil {
			fmt.Printf("\n%v\n", err)
		} else {
			fmt.Printf("%v\n\n\n", client)
		}
	}
}
