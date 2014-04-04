package harvest

import "testing"
import "../harvest"

func testServiceInitialization(client *harvest.APIClient, t *testing.T) {
	if client.Client == nil {
		t.Error("Client service not initialized")
	}
	if client.People == nil {
		t.Error("People service not initialized")
	}
	if client.Project == nil {
		t.Error("Project service not initialized")
	}
	if client.Invoice == nil {
		t.Error("Invoice service not initialized")
	}
	if client.Account == nil {
		t.Error("Account service not initialized")
	}
}

func Test_NewAPIClientWithBasicAuth(t *testing.T) {
	client := harvest.NewAPIClientWithBasicAuth(
		"Some Username",
		"Some Password",
		"Some Domain")

	if client.Username != "Some Username" {
		t.Error("Username not set on client")
	}

	if client.Password != "Some Password" {
		t.Error("Password not set on client")
	}

	if client.Subdomain != "Some Domain" {
		t.Error("Domain not set on client")
	}

	testServiceInitialization(client, t)

}

func Test_NewAPIClientWithAuthToken(t *testing.T) {
	client := harvest.NewAPIClientWithAuthToken(
		"SUPER AUTH TOKEN",
		"Some Domain")

	//TODO: Test that we're setting up the httpClient with a token

	if client.Subdomain != "Some Domain" {
		t.Error("Domain not set on client")
	}

	testServiceInitialization(client, t)
}
