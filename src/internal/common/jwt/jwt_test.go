package jwt

import (
	"testing"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/canario/src/internal/tester"
)

type JWTUnitSuite struct {
	tester.Suite
}

func TestJWTUnitSuite(t *testing.T) {
	suite.Run(t, &JWTUnitSuite{Suite: tester.NewUnitSuite(t)})
}

// createJWTToken creates a JWT token with the specified expiration time.
func createJWTToken(
	claims jwt.RegisteredClaims,
) (string, error) {
	// build claims from map
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(""))
}

const (
	// Raw test token valid for 100 years.
	rawToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9." +
		"eyJuYmYiOiIxNjkxODE5NTc5IiwiZXhwIjoiMzk0NTUyOTE3OSIsImVuZHBvaW50dXJsTGVuZ3RoIjoiMTYw" +
		"IiwiaXNsb29wYmFjayI6IlRydWUiLCJ2ZXIiOiJoYXNoZWRwcm9vZnRva2VuIiwicm9sZXMiOiJhbGxmaWxl" +
		"cy53cml0ZSBhbGxzaXRlcy5mdWxsY29udHJvbCBhbGxwcm9maWxlcy5yZWFkIiwidHQiOiIxIiwiYWxnIjoi" +
		"SFMyNTYifQ" +
		".signature"
)

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
					jwt.RegisteredClaims{
						ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
					})
			},
			expect:    false,
			expectErr: assert.NoError,
		},
		{
			name: "expired token",
			getToken: func() (string, error) {
				return createJWTToken(
					jwt.RegisteredClaims{
						ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Hour)),
					})
			},
			expect:    true,
			expectErr: assert.NoError,
		},
		// Test with a raw token which is not generated with go-jwt lib.
		{
			name: "alive raw token",
			getToken: func() (string, error) {
				return rawToken, nil
			},
			expect:    false,
			expectErr: assert.NoError,
		},
		{
			name: "alive token, missing exp claim",
			getToken: func() (string, error) {
				return createJWTToken(jwt.RegisteredClaims{})
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
