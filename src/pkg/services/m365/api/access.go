package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/m365/graph"
)

// ---------------------------------------------------------------------------
// controller
// ---------------------------------------------------------------------------

func (c Client) Access() *Access {
	return &Access{c}
}

// Access is an interface-compliant provider of the client.
type Access struct {
	Client
}

// GetToken retrieves a m365 application auth token using client id and secret credentials.
// This token is not normally needed in order for corso to function, and is implemented
// primarily as a way to exercise the validity of those credentials without need of specific
// permissions.
func (c Access) GetToken(
	ctx context.Context,
) error {
	var (
		//nolint:lll
		// https://learn.microsoft.com/en-us/graph/connecting-external-content-connectors-api-postman#step-5-get-an-authentication-token
		rawURL = fmt.Sprintf(
			"https://login.microsoftonline.com/%s/oauth2/v2.0/token",
			c.Credentials.AzureTenantID)
		headers = map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
		}
		body = strings.NewReader(fmt.Sprintf(
			"client_id=%s"+
				"&client_secret=%s"+
				"&scope=https://graph.microsoft.com/.default"+
				"&grant_type=client_credentials",
			c.Credentials.AzureClientID,
			c.Credentials.AzureClientSecret))
	)

	resp, err := c.Post(ctx, rawURL, headers, body)
	if err != nil {
		return graph.Stack(ctx, err)
	}

	if resp.StatusCode == http.StatusBadRequest {
		return clues.New("incorrect tenant or application parameters")
	}

	if resp.StatusCode/100 == 4 || resp.StatusCode/100 == 5 {
		return clues.New("non-2xx response: " + resp.Status)
	}

	defer resp.Body.Close()

	return nil
}

type delegatedAccess struct {
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    string `json:"expires_in"`
	ExpiresOn    string `json:"expires_on"`
	NotBefore    string `json:"not_before"`
	Resource     string `json:"resource"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (da delegatedAccess) MinimumPrintable() any {
	return da
}

func (c *Access) GetDelegatedToken(
	ctx context.Context,
) (delegatedAccess, error) {
	var (
		//nolint:lll
		// https://dzone.com/articles/getting-access-token-for-microsoft-graph-using-oau
		rawURL = fmt.Sprintf(
			"https://login.microsoftonline.com/%s/oauth2/token",
			c.Credentials.AzureTenantID)
		headers = map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
		}
		body = strings.NewReader(fmt.Sprintf(
			"client_id=%s"+
				"&client_secret=%s"+
				"&resource=https://graph.microsoft.com"+
				"&grant_type=password"+
				"&username=%s"+
				"&password=%s",
			c.Credentials.AzureClientID,
			c.Credentials.AzureClientSecret,
			c.Credentials.AzureUsername,
			c.Credentials.AzureUserPassword))
	)

	resp, err := c.Post(ctx, rawURL, headers, body)
	if err != nil {
		return delegatedAccess{}, graph.Stack(ctx, err)
	}

	if resp.StatusCode == http.StatusBadRequest {
		return delegatedAccess{}, clues.New("incorrect tenant or credentials")
	}

	if resp.StatusCode/100 == 4 || resp.StatusCode/100 == 5 {
		return delegatedAccess{}, clues.New("non-2xx response: " + resp.Status)
	}

	defer resp.Body.Close()

	var da delegatedAccess
	err = json.NewDecoder(resp.Body).Decode(&da)

	return da, clues.Wrap(err, "undecodable body").WithClues(ctx).OrNil()
}
