package main

import ("../harvest"
        "fmt")

func main() {
  apiClient := harvest.NewAPIClientWithBasicAuth("YOUR_USERNAME", "YOUR_PASSWORD", "YOUR_SUBDOMAIN")
  clients := apiClient.GetClients()

  fmt.Printf("%+v\n\n\n", clients)

  client := apiClient.GetClient(clients[0].Id)
  fmt.Printf("%v\n\n\n", client)
}
