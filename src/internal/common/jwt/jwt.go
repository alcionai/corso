package jwt

import (
	"context"
	"time"

	"github.com/alcionai/clues"
	jwt "github.com/golang-jwt/jwt/v5"

	"github.com/alcionai/corso/src/pkg/logger"
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

// GetJWTLifetime returns the issued at(iat) and expiration time(exp) claims
// present in the JWT token. These are optional claims and may not be present
// in the token. Absence is not reported as an error.
//
// An error is returned if the supplied token is malformed. Times are returned
// in UTC to have parity with graph responses.
func GetJWTLifetime(
	ctx context.Context,
	rawToken string,
) (time.Time, time.Time, error) {
	var (
		issuedAt  time.Time
		expiresAt time.Time
	)

	p := jwt.NewParser()

	token, _, err := p.ParseUnverified(rawToken, &jwt.RegisteredClaims{})
	if err != nil {
		logger.CtxErr(ctx, err).Debug("parsing jwt token")
		return time.Time{}, time.Time{}, clues.Wrap(err, "invalid jwt")
	}

	exp, err := token.Claims.GetExpirationTime()
	if err != nil {
		logger.CtxErr(ctx, err).Debug("extracting exp claim")
		return time.Time{}, time.Time{}, clues.Wrap(err, "getting token expiry time")
	}

	iat, err := token.Claims.GetIssuedAt()
	if err != nil {
		logger.CtxErr(ctx, err).Debug("extracting iat claim")
		return time.Time{}, time.Time{}, clues.Wrap(err, "getting token issued at time")
	}

	// Absence of iat or exp claims is not reported as an error by jwt library as these
	// are optional as per spec.
	if iat != nil {
		issuedAt = iat.UTC()
	}

	if exp != nil {
		expiresAt = exp.UTC()
	}

	return issuedAt, expiresAt, nil
}
