package appsec

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetConfigurationClone(t *testing.T) {

	configId := 8277
	version := 2
	defer gock.Off()

	mock := gock.New(fmt.Sprintf("https://akaa-baseurl-xxxxxxxxxxx-xxxxxxxxxxxxx.luna.akamaiapis.net/appsec/v1/configs/%d/versions/%d", configId, version))
	mock.
		Get(fmt.Sprintf("/appsec/v1/configs/%d/versions/%d", configId, version)).
		HeaderPresent("Authorization").
		Reply(200).
		SetHeader("Content-Type", "application/json;charset=UTF-8").
		BodyString(fmt.Sprintf(`{
    "configId": 8277,
    "configName": "TestConfig",
    "version": 2,
    "versionNotes": "Membership Benefits",
    "createDate": "2013-10-07T17:58:52Z",
    "createdBy": "user1",
    "basedOn": 1,
    "production": {
        "status": "Active",
        "time": "2014-07-08T07:40:00Z"
    },
    "staging": {
        "status": "Inactive"
    }
}`))

	Init(config)
	configurationclone := NewConfigurationCloneResponse()
	configurationclone.ConfigID = 8277
	configurationclone.Version = 2
	err := configurationclone.GetConfigurationClone("TEST")

	//configurations := NewConfigurationsResponse()

	//err := configurations.GetConfigurations("TEST")
	assert.NoError(t, err)
	//assert.Equal(t, assert.IsType(t, &ZoneNamesResponse{}, nameList), true)
	assert.Equal(t, assert.IsType(t, &ConfigurationCloneResponse{}, configurationclone), true)
	//assert.Equal(t, assert.IsType(t, &ConfigurationsResponse{}, configurations), true)
	//assert.Equal(t, configurations.Configurations[0].ID, 22330)
	//assert.Equal(t, configurations.Configurations[1].ID, 7180)

}
