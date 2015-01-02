package harvest

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"code.google.com/p/goauth2/oauth"
)

// APIClient entry point for interacting with the Harvest API
type APIClient struct {
	Username   string
	Password   string
	Subdomain  string
	HTTPClient *http.Client

	Client  *ClientService
	People  *PersonService
	Project *ProjectService
	Invoice *InvoiceService
	Account *AccountService

	protocol string
}

func newAPIClient(subdomain string, httpClient *http.Client) (c *APIClient) {
	c = new(APIClient)
	c.Subdomain = subdomain
	c.protocol = "https"

	if httpClient != nil {
		c.HTTPClient = httpClient
	} else {
		c.HTTPClient = new(http.Client)
	}

	c.Client = &ClientService{service{c}}
	c.People = &PersonService{service{c}}
	c.Project = &ProjectService{service{c}}
	c.Invoice = &InvoiceService{service{c}}
	c.Account = &AccountService{service{c}}
	return
}

// NewAPIClientWithBasicAuth creates a new APIClient that uses http basic auth
func NewAPIClientWithBasicAuth(username, password, subdomain string) (c *APIClient) {
	c = newAPIClient(subdomain, nil)
	c.Username = username
	c.Password = password
	return
}

// NewAPIClientWithAuthToken creates a new APIClient that uses an OAuth token
func NewAPIClientWithAuthToken(token, subdomain string) (c *APIClient) {
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: token},
	}

	c = newAPIClient(subdomain, t.Client())

	return
}

func (c *APIClient) disableSSL() {
	c.protocol = "http"
}

// GetJSON fetches JSON data from a given URL
func (c *APIClient) GetJSON(path string) (jsonResponse []byte, err error) {
	resourceURL := fmt.Sprintf("%v://%v.harvestapp.com%v", c.protocol, c.Subdomain, path)
	request, err := http.NewRequest("GET", resourceURL, nil)
	if err != nil {
		return
	}

	request.SetBasicAuth(c.Username, c.Password)
	resp, err := c.HTTPClient.Do(request)

	if err != nil {
		return
	}

	defer resp.Body.Close()
	jsonResponse, err = ioutil.ReadAll(resp.Body)
	return
}
