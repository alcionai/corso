package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/m365/graph"
)

// ---------------------------------------------------------------------------
// controller
// ---------------------------------------------------------------------------

func (c Client) Access() Access {
	return Access{c}
}

// Access is an interface-compliant provider of the client.
type Access struct {
	Client
}

func (c Access) GetToken(
	ctx context.Context,
) error {
	var (
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
