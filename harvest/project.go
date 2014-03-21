package harvest

import (
	"encoding/json"
	"fmt"
	"time"
)

type ProjectService struct {
	apiClient *APIClient
}

type Project struct {
	Id                               int       `json:"id"`
	ClientId                         int       `json:"client_id"`
	Name                             string    `json:"name"`
	Code                             string    `json:"code"`
	Notes                            string    `json:"notes"`
	BillBy                           string    `json:"bill_by"`
	BudgetBy                         string    `json:"budget_by"`
	Active                           bool      `json:"active"`
	CostBudgetIncludeExpenses        bool      `json:"cost_budget_include_expenses"`
	Billable                         bool      `json:"billable"`
	ShowBudgetToAll                  bool      `json:"show_budget_to_all"`
	CostBudget                       float32   `json:"cost_budget"`
	HourlyRate                       float32   `json:"hourly_rate"`
	Budget                           float32   `json:"budget"`
	NotifyWhenOverBudget             float32   `json:"notify_when_overbudget"`
	OverBudgetNotificationPercentage float32   `json:"over_budget_notification_percentage"`
	OverBudgetNotifiedAt             time.Time `json:"over_budget_notified_at"`
	CreatedAt                        time.Time `json:"created_at"`
	UpdateAt                         time.Time `json:"updated_at"`
	HintEarliestRecordAt             time.Time `json:"hint_earliest_record_at"`
	HintLatestRecordAt               time.Time `json:"hint_latest_record_at"`
}

type ProjectResponse struct {
	Project Project
}

func (c *ProjectService) List() (err error, projects []Project) {
	err, contents := c.apiClient.GetJSON("/projects.json")
	if err != nil {
		return
	}

	var projectResponse []ProjectResponse
	err = json.Unmarshal(contents, &projectResponse)
	if err != nil {
		return
	}

	for _, element := range projectResponse {
		projects = append(projects, element.Project)
	}
	return
}

func (c *ProjectService) Find(projectID int) (err error, project Project) {
	resourceURL := fmt.Sprintf("/projects/%v.json", projectID)
	err, contents := c.apiClient.GetJSON(resourceURL)
	if err != nil {
		return
	}

	var projectResponse ProjectResponse
	err = json.Unmarshal(contents, &projectResponse)
	if err != nil {
		return
	}

	project = projectResponse.Project
	return
}
