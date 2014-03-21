package examples

import (
	"../harvest"
	"fmt"
)

func ExercisePeople(apiClient *harvest.APIClient) {
	err, people := apiClient.People.List()
	if err != nil {
		fmt.Printf("\n%v\n", err)
	} else {
		fmt.Printf("%+v\n\n\n", people)
		err, person := apiClient.People.Find(people[0].Id)
		if err != nil {
			fmt.Printf("\n%v\n", err)
		} else {
			fmt.Printf("%v\n\n\n", person)
		}
	}
}
