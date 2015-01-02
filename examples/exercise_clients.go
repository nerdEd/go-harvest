package examples

import (
	"fmt"

	"github.com/nerded/go-harvest/harvest"
)

// ExerciseClients hits the live API w/ a Client request
func ExerciseClients(apiClient *harvest.APIClient) {
	err, clients := apiClient.Client.List()
	if err != nil {
		fmt.Printf("\n%v\n", err)
	} else {
		fmt.Printf("\n%+v\n", clients)
		err, client := apiClient.Client.Find(clients[0].Id)
		if err != nil {
			fmt.Printf("\n%v\n", err)
		} else {
			fmt.Printf("\n%v\n", client)
		}
	}
}
