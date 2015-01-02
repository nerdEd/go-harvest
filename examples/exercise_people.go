package examples

import (
	"fmt"

	"github.com/nerded/go-harvest/harvest"
)

// ExercisePeople hits the live API w/ a People request
func ExercisePeople(apiClient *harvest.APIClient) {
	people, err := apiClient.People.List()
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
