package harvest

import (
	"fmt"
	"time"
)

type PersonService struct {
	Service
}

type Person struct {
	Id                           int
	DefaultExpenseCategoryId     int
	DefaultExpenseProjectId      int
	DefaultTaskId                int
	DefaultTimeProjectId         int
	DefaultHourlyRate            float32
	FirstName                    string
	LastName                     string
	Email                        string
	IdentityUrl                  string
	OpensocialIdentifier         string
	Telephone                    string
	Timezone                     string
	CostRate                     string
	WeeklyDigestSentOn           string
	Department                   string
	IsContractor                 bool
	IsAdmin                      bool
	IsActive                     bool
	HasAccessToAllFutureProjects bool
	WantsNewsletter              bool
	WantsWeeklyDigest            bool
	CreatedAt                    time.Time
	UpdatedAt                    time.Time
}

type PersonResponse struct {
	Person Person `json:"user"`
}

func (c *PersonService) List() (err error, people []Person) {
	resourceURL := "/people.json"
	var personResponse []PersonResponse
	err = c.list(resourceURL, &personResponse)
	if err != nil {
		return
	}

	for _, element := range personResponse {
		people = append(people, element.Person)
	}
	return
}

func (c *PersonService) Find(personID int) (err error, person Person) {
	resourceURL := fmt.Sprintf("/people/%v.json", personID)
	var personResponse PersonResponse
	err = c.find(resourceURL, &personResponse)
	person = personResponse.Person
	return
}
