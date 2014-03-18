package harvest

import ("encoding/json"
        "fmt"
        "time")

type Client struct {
  Name string
  Currency string
  Active bool
  Id int32
  HighriseId int32
  CreatedAt time.Time
  UpdatedAt string
  Details string
  DefaultInvoiceTimeframe string
  LastInvoiceKind string
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

func GetClient(clientID int32, apiClient *APIClient) (Client) {
  resourceURL := fmt.Sprintf("/clients/%v.json", clientID)
  contents := apiClient.GetJSON(resourceURL)

  var clientResponse ClientResponse
  json.Unmarshal(contents, &clientResponse)
  return clientResponse.Client
}
