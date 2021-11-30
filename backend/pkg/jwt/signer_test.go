package jwt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testSign Signer

func TestMain(m *testing.M) {
	config := MustLoadConfig()
	var err error
	testSign, err = NewSigner(SignKey(config.Key), Iss(config.Iss), ExpireDuration(config.ExpiredDuration))
	if err == nil {
		m.Run()
	}
}

func TestSign(t *testing.T) {
	token, err := testSign.Sign(2, "admin2")

	assert.NoError(t, err)
	t.Log(token)
}

func TestParseSign(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiVXNlcm5hbWUiOiJhZG1pbjIiLCJpc3MiOiJibHVlLWRhc2hib2FyZCIsImV4cCI6MTYzNjE2NTI1NywiaWF0IjoxNjM2MDc4ODU3fQ.cS8hDN1yLxNOWUBWl4cMedxZMyQqIx0R_bt3nOBOGBE"
	claim, err := testSign.ParseClaims(tokenString)
	assert.NoError(t, err)
	t.Log(claim.ID, claim.Username)

}
