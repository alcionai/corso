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

func (suite *JWTUnitSuite) TestGetJWTLifetime() {
	// Set of time values to be used in the tests.
	// Truncate to seconds for comparisons since jwt tokens have second
	// level precision.
	idToTime := map[string]time.Time{
		"T0": time.Now().UTC().Add(-time.Hour).Truncate(time.Second),
		"T1": time.Now().UTC().Truncate(time.Second),
		"T2": time.Now().UTC().Add(time.Hour).Truncate(time.Second),
	}

	table := []struct {
		name       string
		getToken   func() (string, error)
		expectFunc func(t *testing.T, iat time.Time, exp time.Time)
		expectErr  assert.ErrorAssertionFunc
	}{
		{
			name: "alive token",
			getToken: func() (string, error) {
				return createJWTToken(
					jwt.RegisteredClaims{
						IssuedAt:  jwt.NewNumericDate(idToTime["T0"]),
						ExpiresAt: jwt.NewNumericDate(idToTime["T1"]),
					})
			},
			expectFunc: func(t *testing.T, iat time.Time, exp time.Time) {
				assert.Equal(t, idToTime["T0"], iat)
				assert.Equal(t, idToTime["T1"], exp)
			},
			expectErr: assert.NoError,
		},
		// Test with a token which is not generated using the go-jwt lib.
		// This is a long lived token which is valid for 100 years.
		{
			name: "alive raw token with iat and exp claims",
			getToken: func() (string, error) {
				return rawToken, nil
			},
			expectFunc: func(t *testing.T, iat time.Time, exp time.Time) {
				assert.Less(t, iat, time.Now(), "iat should be in the past")
				assert.Greater(t, exp, time.Now(), "exp should be in the future")
			},
			expectErr: assert.NoError,
		},
		// Regardless of whether the token is expired or not, we should be able to
		// extract the iat and exp claims from it without error.
		{
			name: "expired token",
			getToken: func() (string, error) {
				return createJWTToken(
					jwt.RegisteredClaims{
						IssuedAt:  jwt.NewNumericDate(idToTime["T1"]),
						ExpiresAt: jwt.NewNumericDate(idToTime["T0"]),
					})
			},
			expectFunc: func(t *testing.T, iat time.Time, exp time.Time) {
				assert.Equal(t, idToTime["T1"], iat)
				assert.Equal(t, idToTime["T0"], exp)
			},
			expectErr: assert.NoError,
		},
		{
			name: "missing iat claim",
			getToken: func() (string, error) {
				return createJWTToken(
					jwt.RegisteredClaims{
						ExpiresAt: jwt.NewNumericDate(idToTime["T2"]),
					})
			},
			expectFunc: func(t *testing.T, iat time.Time, exp time.Time) {
				assert.Equal(t, time.Time{}, iat)
				assert.Equal(t, idToTime["T2"], exp)
			},
			expectErr: assert.NoError,
		},
		{
			name: "missing exp claim",
			getToken: func() (string, error) {
				return createJWTToken(
					jwt.RegisteredClaims{
						IssuedAt: jwt.NewNumericDate(idToTime["T0"]),
					})
			},
			expectFunc: func(t *testing.T, iat time.Time, exp time.Time) {
				assert.Equal(t, idToTime["T0"], iat)
				assert.Equal(t, time.Time{}, exp)
			},
			expectErr: assert.NoError,
		},
		{
			name: "both claims missing",
			getToken: func() (string, error) {
				return createJWTToken(jwt.RegisteredClaims{})
			},
			expectFunc: func(t *testing.T, iat time.Time, exp time.Time) {
				assert.Equal(t, time.Time{}, iat)
				assert.Equal(t, time.Time{}, exp)
			},
			expectErr: assert.NoError,
		},
		{
			name: "malformed token",
			getToken: func() (string, error) {
				return "header.claims.signature", nil
			},
			expectFunc: func(t *testing.T, iat time.Time, exp time.Time) {
				assert.Equal(t, time.Time{}, iat)
				assert.Equal(t, time.Time{}, exp)
			},
			expectErr: assert.Error,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			token, err := test.getToken()
			require.NoError(t, err)

			iat, exp, err := GetJWTLifetime(ctx, token)
			test.expectErr(t, err)

			test.expectFunc(t, iat, exp)
		})
	}
}
