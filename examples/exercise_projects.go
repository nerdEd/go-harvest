package examples

import (
	"fmt"

	"github.com/nerded/go-harvest/harvest"
)

// ExerciseProjects hits the live API w/ a Projects request
func ExerciseProjects(apiClient *harvest.APIClient) {
	err, projects := apiClient.Project.List()
	if err != nil {
		fmt.Printf("\n%v\n", err)
	} else {
		fmt.Printf("%+v\n", projects)
		err, project := apiClient.Project.Find(projects[0].Id)
		if err != nil {
			fmt.Printf("\n%v\n", err)
		} else {
			fmt.Printf("%v\n", project)
		}
	}
}
