package harvest

import (
	"fmt"
	"time"
)

// ClientService wraps interactions with Harvest Client API
type ClientService struct {
	service
}

// Client represents a Harvest Client
type Client struct {
	Name                    string    `json:"name"`
	Currency                string    `json:"currency"`
	Active                  bool      `json:"active"`
	ID                      int       `json:"id"`
	HighriseID              int       `json:"highrise_id"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"created_at"`
	Details                 string    `json:"details"`
	DefaultInvoiceTimeframe string    `json:"default_invoice_timeframe"`
	LastInvoiceKind         string    `json:"last_invoice_kind"`
}

type clientResponse struct {
	Client Client
}

// List fetch all clients
func (c *ClientService) List() (clients []Client, err error) {
	resourceURL := "/clients.json"
	var clientResponse []clientResponse
	err = c.list(resourceURL, &clientResponse)
	if err != nil {
		return
	}
	for _, element := range clientResponse {
		clients = append(clients, element.Client)
	}
	return
}

// Find gets single client based on ID
func (c *ClientService) Find(clientID int) (client Client, err error) {
	resourceURL := fmt.Sprintf("/clients/%v.json", clientID)
	var clientResponse clientResponse
	err = c.find(resourceURL, &clientResponse)

	return clientResponse.Client, err
}
