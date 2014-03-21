package main

import (
	"./examples"
	"./harvest"
	"os"
)

func main() {
	apiClient := harvest.NewAPIClientWithBasicAuth(
		os.Getenv("HARVEST_USERNAME"),
		os.Getenv("HARVEST_PASSWORD"),
		os.Getenv("HARVEST_DOMAIN"))

	examples.ExerciseProjects(apiClient)
	examples.ExerciseClients(apiClient)
	examples.ExercisePeople(apiClient)
}
