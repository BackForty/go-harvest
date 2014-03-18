package harvest

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type APIClient struct {
	username  string
	password  string
	token     string
	subdomain string
	Client    *http.Client
}

func newAPIClient(subdomain string) (c *APIClient) {
	c = new(APIClient)
	c.subdomain = subdomain
	c.Client = new(http.Client)
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

func (c *APIClient) GetJSON(path string) (jsonResponse []byte) {
	resourceURL := fmt.Sprintf("http://%v.harvestapp.com%v", c.subdomain, path)
	request, err := http.NewRequest("GET", resourceURL, nil)
	perror(err)

	request.SetBasicAuth(c.username, c.password)
	resp, err := c.Client.Do(request)
	perror(err)
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	perror(err)

	return contents
}

func (c *APIClient) GetClients() (clients []Client) {
	clients = GetClients(c)
	return
}

func (c *APIClient) GetClient(clientID int32) (client Client) {
	client = GetClient(clientID, c)
	return
}

func perror(err error) {
	if err != nil {
		panic(err)
	}
}
