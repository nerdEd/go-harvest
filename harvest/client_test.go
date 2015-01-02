package harvest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FetchAllClientsFromServerWithSuccessResponse(t *testing.T) {

	clientsJSONResponse := ` [
                              {
                                  "client": {
                                      "id": 2039096,
                                      "name": "Alchemy",
                                      "active": false,
                                      "currency": "United States Dollars - USD",
                                      "highrise_id": null,
                                      "cache_version": 3836234653,
                                      "updated_at": "2014-09-26T15:11:05Z",
                                      "created_at": "2013-11-04T14:08:04Z",
                                      "currency_symbol": "$",
                                      "details": "",
                                      "default_invoice_timeframe": "Custom",
                                      "last_invoice_kind": "project"
                                  }
                              },
                              {
                                  "client": {
                                      "id": 2034799,
                                      "name": "Applications Online",
                                      "active": false,
                                      "currency": "United States Dollars - USD",
                                      "highrise_id": null,
                                      "cache_version": 2734976250,
                                      "updated_at": "2014-05-08T14:06:38Z",
                                      "created_at": "2013-11-01T18:31:57Z",
                                      "currency_symbol": "$",
                                      "details": "",
                                      "default_invoice_timeframe": "Custom",
                                      "last_invoice_kind": "task"
                                  }
                              }
                          ]`

	apiClient := NewAPIClientWithBasicAuth("username", "password", "subdomain")
	s, httpClient := mockHTTPClient(200, clientsJSONResponse)
	defer s.Close()
	apiClient.HTTPClient = httpClient
	apiClient.disableSSL() // What a hack

	// Attemp to get clients
	clients, err := apiClient.Client.List()

	assert.Nil(t, err)
	assert.NotEmpty(t, clients)
}

func Test_FindClientFromServerWithSuccessResponse(t *testing.T) {

	clientJSONResponse := ` {
                            "client": {
                                "id": 2039096,
                                "name": "Alchemy",
                                "active": false,
                                "currency": "United States Dollars - USD",
                                "highrise_id": null,
                                "cache_version": 3836234653,
                                "updated_at": "2014-09-26T15:11:05Z",
                                "created_at": "2013-11-04T14:08:04Z",
                                "currency_symbol": "$",
                                "details": "",
                                "default_invoice_timeframe": "Custom",
                                "last_invoice_kind": "project"
                            }
                        }`

	apiClient := NewAPIClientWithBasicAuth("username", "password", "subdomain")
	s, httpClient := mockHTTPClient(200, clientJSONResponse)
	defer s.Close()
	apiClient.HTTPClient = httpClient
	apiClient.disableSSL() // What a hack

	// Attemp to get clients
	client, err := apiClient.Client.Find(2039096)

	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "Alchemy", client.Name)
	assert.Equal(t, false, client.Active)
	assert.Equal(t, "United States Dollars - USD", client.Currency)
}

func Test_FindClientFromServerWithNotFoundResponse(t *testing.T) {

	clientJSONResponse := ` {
                              "status": "404",
                              "error": "Not Found"
                          }`

	apiClient := NewAPIClientWithBasicAuth("username", "password", "subdomain")
	s, httpClient := mockHTTPClient(404, clientJSONResponse)
	defer s.Close()
	apiClient.HTTPClient = httpClient
	apiClient.disableSSL() // What a hack

	// TODO: Make all this work
	// Attemp to get clients
	//client, err := apiClient.Client.Find(2039096)

	//assert.Nil(t, err)
	//assert.Nil(t, client)
}
