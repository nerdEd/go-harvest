package harvest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FetchAllClientsFromServerWithSuccessResponse(t *testing.T) {

	clientsJSONResponse := `[
                            {
                              "name": "Big Government",
                              "active": true
                            }
                          ]`

	apiClient := NewAPIClientWithBasicAuth("username", "password", "subdomain")
	s, httpClient := mockHTTPClient(200, clientsJSONResponse)
	defer s.Close()
	apiClient.HTTPClient = httpClient
	apiClient.disableSSL() // What a hack

	// Attemp to get clients
	err, clients := apiClient.Client.List()

	assert.Nil(t, err)
	assert.NotEmpty(t, clients)
}

func Test_FetchAllClientsFromServerWithFailureResponse(t *testing.T) {
}
