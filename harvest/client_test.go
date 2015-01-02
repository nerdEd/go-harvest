package harvest

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FetchAllClientsFromServerWithSuccessResponse(t *testing.T) {
	testJSON := `{
                  "name": "Big Government",
                  "active": true
               }`

	b := []byte(testJSON)

	client := Client{}
	err := json.Unmarshal(b, &client)

	if err != nil {
		fmt.Println("There was an error again")
		t.Error(err)
	}

	assert.Equal(t, true, client.Active)
	assert.Equal(t, "Big Government", client.Name)
}

func Test_FetchAllClientsFromServerWithFailureResponse(t *testing.T) {
}
