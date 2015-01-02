package main

import (
	"os"

	"github.com/nerded/go-harvest/examples"
	"github.com/nerded/go-harvest/harvest"
)

func main() {
	apiClient := harvest.NewAPIClientWithBasicAuth(
		os.Getenv("HARVEST_USERNAME"),
		os.Getenv("HARVEST_PASSWORD"),
		os.Getenv("HARVEST_DOMAIN"))

	examples.ExerciseClients(apiClient)
	examples.ExerciseInvoices(apiClient)
	examples.ExercisePeople(apiClient)
	examples.ExerciseProjects(apiClient)
	examples.ExerciseAccounts(apiClient)
}
