package harvest

import (
	"encoding/json"
)

type Service struct {
	apiClient *APIClient
}

func (service Service) find(resourceURL string, unmarshalContainer interface{}) (err error) {
	err, contents := service.apiClient.GetJSON(resourceURL)
	if err != nil {
		return
	}

	err = json.Unmarshal(contents, unmarshalContainer)

	return
}

func (service Service) list(resourceURL string, unmarshalContainer interface{}) (err error) {
	err, contents := service.apiClient.GetJSON(resourceURL)
	if err != nil {
		return
	}

	err = json.Unmarshal(contents, &unmarshalContainer)
	return
}
