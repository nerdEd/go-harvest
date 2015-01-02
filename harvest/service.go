package harvest

import (
	"encoding/json"
)

type service struct {
	apiClient *APIClient
}

func (service service) find(resourceURL string, unmarshalContainer interface{}) (err error) {
	contents, err := service.apiClient.GetJSON(resourceURL)
	if err != nil {
		return
	}

	err = json.Unmarshal(contents, unmarshalContainer)

	return
}

func (service service) list(resourceURL string, unmarshalContainer interface{}) (err error) {
	contents, err := service.apiClient.GetJSON(resourceURL)
	if err != nil {
		return
	}

	err = json.Unmarshal(contents, &unmarshalContainer)
	return
}
