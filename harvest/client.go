package harvest

import (
	"encoding/json"
	"fmt"
	"time"
)

type Client struct {
	Name                    string    `json:"name"`
	Currency                string    `json:"currency"`
	Active                  bool      `json:"active"`
	Id                      int       `json:"id"`
	HighriseId              int       `json:"highrise_id"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"created_at"`
	Details                 string    `json:"details"`
	DefaultInvoiceTimeframe string    `json:"default_invoice_timeframe"`
	LastInvoiceKind         string    `json:"last_invoice_kind"`
}

type ClientResponse struct {
	Client Client
}

func GetClients(apiClient *APIClient) (clients []Client) {
	contents := apiClient.GetJSON("/clients.json")

	var clientResponse []ClientResponse
	json.Unmarshal(contents, &clientResponse)
	for _, element := range clientResponse {
		clients = append(clients, element.Client)
	}
	return
}

func GetClient(clientID int, apiClient *APIClient) Client {
	resourceURL := fmt.Sprintf("/clients/%v.json", clientID)
	contents := apiClient.GetJSON(resourceURL)

	var clientResponse ClientResponse
	json.Unmarshal(contents, &clientResponse)
	return clientResponse.Client
}
