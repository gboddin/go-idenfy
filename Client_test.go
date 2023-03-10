package idenfy

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestClient_DecodeReaderIdentityCallback(t *testing.T) {
	expectedSig := "e7386736c7f0e8ee293a7637a316d64f7193c3197c133576e3f72061100180e2"
	client, err := NewClient(WithCallbackSignatureKey("TestingKey"))
	assert.NoError(t, err, "Client creation failed")
	assert.NotNil(t, client, "Client is nil")
	webhook1, err := os.Open("testdata/webhook.1.json")
	assert.NoError(t, err, "Could not open test data")
	resp, err := client.DecodeReaderIdentityCallback(webhook1, expectedSig)
	assert.NoError(t, err)
	assert.Equal(t, "123", resp.ClientId)
	assert.Len(t, resp.AML, 1)
	assert.Len(t, resp.LID, 1)
	frontFile, found := resp.FileUrls["FRONT"]
	assert.True(t, found)
	assert.Equal(t, "https://s3.eu-west-1.amazonaws.com/production.users.storage/users_storage/users/<HASH>/FRONT.png?AWSAccessKeyId=<KEY>&Signature=<SIG>&Expires=<STAMP>", frontFile)
	assert.Equal(t, "ADDRESS EXAMPLE", resp.Data.Address)
	assert.Equal(t, "APPROVED", resp.Status.Overall)
	assert.Equal(t, int64(1554726960), resp.StartTime)
}
