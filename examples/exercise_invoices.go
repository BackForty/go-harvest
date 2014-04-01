package examples

import (
	"../harvest"
	"fmt"
)

func ExerciseInvoices(apiClient *harvest.APIClient) {
	err, invoices := apiClient.Invoice.List()
	if err != nil {
		fmt.Printf("\n%v\n", err)
	} else {
		fmt.Printf("\n%+v\n", invoices)
		err, invoice := apiClient.Invoice.Find(invoices[0].Id)
		if err != nil {
			fmt.Printf("\n%v\n", err)
		} else {
			fmt.Printf("\n%v\n", invoice)
		}
	}
}
