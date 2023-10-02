package jwt

import (
	"testing"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type JWTUnitSuite struct {
	tester.Suite
}

func TestJWTUnitSuite(t *testing.T) {
	suite.Run(t, &JWTUnitSuite{Suite: tester.NewUnitSuite(t)})
}

// createJWTToken creates a JWT token with the specified expiration time.
func createJWTToken(
	expiration time.Time,
	claims jwt.MapClaims,
) (string, error) {
	// build claims from map
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(""))
}

func (suite *JWTUnitSuite) TestIsJWTExpired() {
	table := []struct {
		name      string
		expect    bool
		getToken  func() (string, error)
		expectErr assert.ErrorAssertionFunc
	}{
		{
			name: "alive token",
			getToken: func() (string, error) {
				return createJWTToken(
					time.Now().Add(time.Hour),
					jwt.MapClaims{
						"exp": time.Now().Add(time.Hour).Unix(),
					})
			},
			expect:    false,
			expectErr: assert.NoError,
		},
		{
			name: "expired token",
			getToken: func() (string, error) {
				return createJWTToken(
					time.Now().Add(time.Hour),
					jwt.MapClaims{
						"exp": time.Now().Add(-time.Hour).Unix(),
					})
			},
			expect:    true,
			expectErr: assert.NoError,
		},
		{
			name: "alive token, missing exp claim",
			getToken: func() (string, error) {
				return createJWTToken(time.Now().Add(time.Hour), jwt.MapClaims{})
			},
			expect:    false,
			expectErr: assert.NoError,
		},
		{
			name: "malformed token",
			getToken: func() (string, error) {
				return "header.claims.signature", nil
			},
			expect:    false,
			expectErr: assert.Error,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			_, flush := tester.NewContext(t)
			defer flush()

			token, err := test.getToken()
			require.NoError(t, err)

			expired, err := IsJWTExpired(token)
			test.expectErr(t, err)

			assert.Equal(t, test.expect, expired)
		})
	}
}
