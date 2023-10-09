package jwt

import (
	"time"

	"github.com/alcionai/clues"
	jwt "github.com/golang-jwt/jwt/v5"
)

// IsJWTExpired checks if the JWT token is past expiry by analyzing the
// "exp" claim present in the token. Token is considered expired if "exp"
// claim < current time. Missing "exp" claim is considered as non-expired.
// An error is returned if the supplied token is malformed.
func IsJWTExpired(
	rawToken string,
) (bool, error) {
	p := jwt.NewParser()

	// Note: Call to ParseUnverified is intentional since token verification is
	// not our objective. We only care about the embed claims in the token.
	// We assume the token signature is valid & verified by caller stack.
	token, _, err := p.ParseUnverified(rawToken, &jwt.RegisteredClaims{})
	if err != nil {
		return false, clues.Wrap(err, "invalid jwt")
	}

	t, err := token.Claims.GetExpirationTime()
	if err != nil {
		return false, clues.Wrap(err, "getting token expiry time")
	}

	if t == nil {
		return false, nil
	}

	expired := t.Before(time.Now())

	return expired, nil
}
