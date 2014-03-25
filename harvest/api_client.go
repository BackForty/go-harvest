package harvest

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type APIClient struct {
	username   string
	password   string
	token      string
	subdomain  string
	httpClient *http.Client

	Client  *ClientService
	People  *PersonService
	Project *ProjectService
}

func newAPIClient(subdomain string) (c *APIClient) {
	c = new(APIClient)
	c.subdomain = subdomain
	c.httpClient = new(http.Client)

	c.Client = &ClientService{Service{c}}
	c.People = &PersonService{Service{c}}
	c.Project = &ProjectService{Service{c}}
	return
}

func NewAPIClientWithBasicAuth(username, password, subdomain string) (c *APIClient) {
	c = newAPIClient(subdomain)
	c.username = username
	c.password = password
	return
}

func NewAPIClientWithAuthToken(token, subdomain string) (c *APIClient) {
	c = newAPIClient(subdomain)
	c.token = token
	return
}

func (c *APIClient) GetJSON(path string) (err error, jsonResponse []byte) {
	resourceURL := fmt.Sprintf("http://%v.harvestapp.com%v", c.subdomain, path)
	request, err := http.NewRequest("GET", resourceURL, nil)
	if err != nil {
		return
	}

	request.SetBasicAuth(c.username, c.password)
	resp, err := c.httpClient.Do(request)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	jsonResponse, err = ioutil.ReadAll(resp.Body)
	return
}
