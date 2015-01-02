package examples

import (
	"fmt"

	"github.com/nerded/go-harvest/harvest"
)

// ExerciseAccounts hits the live API w/ an Account request
func ExerciseAccounts(apiClient *harvest.APIClient) {
	account, err := apiClient.Account.Find()
	if err != nil {
		fmt.Printf("\n%v\n", err)
	} else {
		fmt.Printf("\n%v\n", account)
	}
}
