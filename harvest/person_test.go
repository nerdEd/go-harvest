package harvest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FetchAllPeopleFromServerWithSuccessResponse(t *testing.T) {

	peopleResponse := ` [
                        {
                            "user": {
                                "id": 518561,
                                "email": "adam.v.duke@gmail.com",
                                "created_at": "2013-05-20T19:37:01Z",
                                "is_admin": false,
                                "first_name": "Adam",
                                "last_name": "Duke",
                                "timezone": "Eastern Time (US & Canada)",
                                "is_contractor": true,
                                "telephone": "",
                                "is_active": false,
                                "has_access_to_all_future_projects": false,
                                "default_hourly_rate": null,
                                "identity_url": null,
                                "department": "",
                                "wants_newsletter": false,
                                "updated_at": "2013-09-23T13:24:00Z",
                                "wants_weekly_digest": true,
                                "weekly_digest_sent_on": null,
                                "cost_rate": null,
                                "identity_account_id": 249801
                            }
                        },
                        {
                            "user": {
                                "id": 624830,
                                "email": "ali@inthebackforty.com",
                                "created_at": "2013-11-18T03:00:16Z",
                                "is_admin": false,
                                "first_name": "Ali",
                                "last_name": "Ibrahim",
                                "timezone": "Eastern Time (US & Canada)",
                                "is_contractor": false,
                                "telephone": "",
                                "is_active": false,
                                "has_access_to_all_future_projects": true,
                                "default_hourly_rate": 200,
                                "identity_url": null,
                                "department": "",
                                "wants_newsletter": true,
                                "updated_at": "2014-10-18T21:17:51Z",
                                "wants_weekly_digest": true,
                                "weekly_digest_sent_on": null,
                                "cost_rate": null,
                                "identity_account_id": 249801
                            }
                        }]`

	apiClient := NewAPIClientWithBasicAuth("username", "password", "subdomain")
	s, httpClient := mockHTTPClient(200, peopleResponse)
	defer s.Close()
	apiClient.HTTPClient = httpClient
	apiClient.disableSSL() // What a hack

	people, err := apiClient.People.List()

	assert.Nil(t, err)
	assert.NotEmpty(t, people)
	assert.Equal(t, 2, len(people))
}

func Test_FindPersonFromServerWithSuccessResponse(t *testing.T) {

	personJSONResponse := ` {
                            "user": {
                                "id": 518561,
                                "email": "adam.v.duke@gmail.com",
                                "created_at": "2013-05-20T19:37:01Z",
                                "is_admin": false,
                                "first_name": "Adam",
                                "last_name": "Duke",
                                "timezone": "Eastern Time (US & Canada)",
                                "is_contractor": true,
                                "telephone": "",
                                "is_active": false,
                                "has_access_to_all_future_projects": false,
                                "default_hourly_rate": null,
                                "identity_url": null,
                                "department": "",
                                "wants_newsletter": false,
                                "updated_at": "2013-09-23T13:24:00Z",
                                "wants_weekly_digest": true,
                                "weekly_digest_sent_on": null,
                                "cost_rate": null,
                                "identity_account_id": 249801
                            }
                        }`

	apiClient := NewAPIClientWithBasicAuth("username", "password", "subdomain")
	s, httpClient := mockHTTPClient(200, personJSONResponse)
	defer s.Close()
	apiClient.HTTPClient = httpClient
	apiClient.disableSSL() // What a hack

	person, err := apiClient.People.Find(518561)

	assert.Nil(t, err)
	assert.NotNil(t, person)
	assert.Equal(t, "adam.v.duke@gmail.com", person.Email)
	assert.Equal(t, false, person.IsAdmin)
	assert.Equal(t, false, person.IsActive)
}
