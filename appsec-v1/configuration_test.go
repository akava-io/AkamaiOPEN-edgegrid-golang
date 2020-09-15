package appsec

import (
	"fmt"
	"testing"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

var (
	config = edgegrid.Config{
		Host:         "akaa-baseurl-xxxxxxxxxxx-xxxxxxxxxxxxx.luna.akamaiapis.net/",
		AccessToken:  "akab-access-token-xxx-xxxxxxxxxxxxxxxx",
		ClientToken:  "akab-client-token-xxx-xxxxxxxxxxxxxxxx",
		ClientSecret: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx=",
		MaxBody:      2048,
		Debug:        false,
	}
)

func TestGetConfigurations(t *testing.T) {

	defer gock.Off()

	mock := gock.New("https://akaa-baseurl-xxxxxxxxxxx-xxxxxxxxxxxxx.luna.akamaiapis.net/appsec/v1/configs")
	mock.
		Get("/appsec/v1/configs").
		HeaderPresent("Authorization").
		Reply(200).
		SetHeader("Content-Type", "application/json;charset=UTF-8").
		BodyString(fmt.Sprintf(`{
			"configurations": [
			  {
				"id": 22330,
				"latestVersion": 5,
				"name": "CaroTestTransition2Versioning",
				"description": "(user notes)"
			  },
			  {
				"id": 7180,
				"latestVersion": 9,
				"name": "Corporate Sites WAF",
				"productionVersion": 1,
				"stagingVersion": 2,
				"productionHostnames": [
				  "example.com",
				  "www.example.net",
				  "m.example.com"
				]
			  }
			]
		  }
		  `))

	Init(config)
	configuration := NewConfigurationResponse()

	err := configuration.GetConfiguration("TEST")

	//configurations := NewConfigurationsResponse()

	//err := configurations.GetConfigurations("TEST")
	assert.NoError(t, err)
	//assert.Equal(t, assert.IsType(t, &ZoneNamesResponse{}, nameList), true)
	assert.Equal(t, assert.IsType(t, &ConfigurationResponse{}, configuration), true)
	//assert.Equal(t, assert.IsType(t, &ConfigurationsResponse{}, configurations), true)
	//assert.Equal(t, configurations.Configurations[0].ID, 22330)
	//assert.Equal(t, configurations.Configurations[1].ID, 7180)

}
