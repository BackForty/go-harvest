package examples

import (
	"../harvest"
	"fmt"
)

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
