package harvest

import (
	"fmt"
	"time"
)

type ClientService struct {
	Service
}

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

func (c *ClientService) List() (err error, clients []Client) {
	resourceURL := "/clients.json"
	var clientResponse []ClientResponse
	err = c.list(resourceURL, &clientResponse)
	if err != nil {
		return
	}
	for _, element := range clientResponse {
		clients = append(clients, element.Client)
	}
	return
}

func (c *ClientService) Find(clientID int) (err error, client Client) {
	resourceURL := fmt.Sprintf("/clients/%v.json", clientID)
	var clientResponse ClientResponse
	err = c.find(resourceURL, &clientResponse)

	return err, clientResponse.Client
}
