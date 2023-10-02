package jwt

import (
	"github.com/alcionai/clues"
	jwt "github.com/golang-jwt/jwt"
)

// IsJWTExpired checks if the JWT token is past expiry by analyzing the
// "exp" claim present in the token. Token is considered alive if :
// 1. time.now <= "exp" claim.
// 2. "exp" claim is missing.
// An error is returned if the supplied token is malformed.
func IsJWTExpired(
	rawToken string,
) (bool, error) {
	// Note: Call to ParseUnverified is intentional since token verification is
	// not our objective. We assume the token signature is valid & verified
	// by caller stack. We only care about the embed claims in the token.
	token, _, err := new(jwt.Parser).ParseUnverified(rawToken, jwt.MapClaims{})
	if err != nil {
		return false, clues.Wrap(err, "invalid jwt")
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	// If "exp" claim is missing, token is considered alive.
	expired := !claims.VerifyExpiresAt(jwt.TimeFunc().Unix(), false)

	return expired, nil
}
