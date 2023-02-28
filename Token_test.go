package idenfy

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenRequest_Validate(t *testing.T) {
	token := &TokenRequest{}
	assert.Error(t, token.Validate())
	token.ClientId = "123"
	assert.NoError(t, token.Validate())
	token.SuccessUrl = "ftp://8.8.8.8:333"
	assert.Error(t, token.Validate())
	token.SuccessUrl = "https://8.8.8.8"
	assert.NoError(t, token.Validate())
	token.DateOfBirth = &Date{
		Year:  1999,
		Month: 13,
		Day:   33,
	}
	_, err := json.Marshal(&token)
	assert.Error(t, err)
	token.DateOfBirth = &Date{
		Year:  1999,
		Month: 12,
		Day:   30,
	}
	_, err = json.Marshal(&token)
	assert.NoError(t, err)
	token.Sex = "male"
	assert.Error(t, token.Validate())
	token.Sex = "M"
	assert.NoError(t, token.Validate())
}
