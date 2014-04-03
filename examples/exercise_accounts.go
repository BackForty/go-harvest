package examples

import (
	"fmt"

	"../harvest"
)

func ExerciseAccounts(apiClient *harvest.APIClient) {
	err, account := apiClient.Account.Find()
	if err != nil {
		fmt.Printf("\n%v\n", err)
	} else {
		fmt.Printf("\n%v\n", account)
	}
}
