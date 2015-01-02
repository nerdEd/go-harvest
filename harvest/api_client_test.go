package harvest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertAllClientsInitialized(client APIClient, t *testing.T) {
	assert.NotNil(t, client.Client)
	assert.NotNil(t, client.People)
	assert.NotNil(t, client.Project)
	assert.NotNil(t, client.Invoice)
	assert.NotNil(t, client.Account)
}

func Test_NewAPIClientWithBasicAuth(t *testing.T) {

	client := NewAPIClientWithBasicAuth(
		"Some Username",
		"Some Password",
		"Some Domain")

	assert.Equal(t, "Some Username", client.Username)
	assert.Equal(t, "Some Password", client.Password)
	assert.Equal(t, "Some Domain", client.Subdomain)

	assertAllClientsInitialized(*client, t)
}

func Test_NewAPIClientWithAuthToken(t *testing.T) {

	client := NewAPIClientWithAuthToken(
		"SUPER AUTH TOKEN",
		"Some Domain")

	//TODO: Test that we're setting up the httpClient with a token

	assert.Equal(t, "Some Domain", client.Subdomain)
	assertAllClientsInitialized(*client, t)
}
