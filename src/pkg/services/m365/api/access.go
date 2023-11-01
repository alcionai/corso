package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
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

type delegatedResp struct {
	AccessToken     string `json:"access_token,omitempty"`
	Devicecode      string `json:"device_code,omitempty"`
	ExpiresIn       string `json:"expires_in,omitempty"`
	ExpiresOn       string `json:"expires_on,omitempty"`
	Interval        string `json:"interval,omitempty"`
	Message         string `json:"message,omitempty"`
	NotBefore       string `json:"not_before,omitempty"`
	RefreshToken    string `json:"refresh_token,omitempty"`
	Resource        string `json:"resource,omitempty"`
	Scope           string `json:"scope,omitempty"`
	TokenType       string `json:"token_type,omitempty"`
	UserCode        string `json:"user_code,omitempty"`
	VerificationURI string `json:"verification_uri,omitempty"`
}

func (da delegatedResp) MinimumPrintable() any {
	return da
}

type deviceResp struct {
	AccessToken     string `json:"access_token,omitempty"`
	Devicecode      string `json:"device_code,omitempty"`
	ExpiresIn       int    `json:"expires_in,omitempty"`
	ExpiresOn       string `json:"expires_on,omitempty"`
	IDToken         string `json:"id_token,omitempty"`
	Interval        int    `json:"interval,omitempty"`
	Message         string `json:"message,omitempty"`
	NotBefore       string `json:"not_before,omitempty"`
	RefreshToken    string `json:"refresh_token,omitempty"`
	Resource        string `json:"resource,omitempty"`
	Scope           string `json:"scope,omitempty"`
	TokenType       string `json:"token_type,omitempty"`
	UserCode        string `json:"user_code,omitempty"`
	VerificationURI string `json:"verification_uri,omitempty"`
}

func (da deviceResp) MinimumPrintable() any {
	return da
}

func (c *Access) GetDelegatedToken(
	ctx context.Context,
) (delegatedResp, error) {
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
		return delegatedResp{}, graph.Stack(ctx, err)
	}

	if resp.StatusCode == http.StatusBadRequest {
		return delegatedResp{}, clues.New("incorrect tenant or credentials")
	}

	if resp.StatusCode/100 == 4 || resp.StatusCode/100 == 5 {
		return delegatedResp{}, clues.New("non-2xx response: " + resp.Status)
	}

	defer resp.Body.Close()

	var ar delegatedResp
	err = json.NewDecoder(resp.Body).Decode(&ar)

	return ar, clues.Wrap(err, "undecodable resp body").WithClues(ctx).OrNil()
}

func (c *Access) RequestDeviceToken(
	ctx context.Context,
) (deviceResp, error) {
	var (
		//nolint:lll
		// https://dzone.com/articles/getting-access-token-for-microsoft-graph-using-oau
		rawURL = fmt.Sprintf(
			"https://login.microsoftonline.com/%s/oauth2/v2.0/devicecode",
			c.Credentials.AzureTenantID)
		headers = map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
		}
		body = strings.NewReader(fmt.Sprintf(
			"client_id=%s&client_secret=%s&scope=%s",
			c.Credentials.AzureClientID,
			c.Credentials.AzureClientSecret,
			"user.read openid profile offline_access"))
	)

	resp, err := c.Post(ctx, rawURL, headers, body)
	if err != nil {
		return deviceResp{}, graph.Stack(ctx, err)
	}

	if resp.StatusCode == http.StatusBadRequest {
		return deviceResp{}, clues.New("incorrect tenant or credentials")
	}

	if resp.StatusCode/100 == 4 || resp.StatusCode/100 == 5 {
		return deviceResp{}, clues.New("non-2xx response: " + resp.Status)
	}

	defer resp.Body.Close()

	var ar deviceResp
	err = json.NewDecoder(resp.Body).Decode(&ar)

	return ar, clues.Wrap(err, "undecodable resp body").WithClues(ctx).OrNil()
}

func (c *Access) GetDeviceToken(
	ctx context.Context,
	deviceCode string,
) (deviceResp, error) {
	var (
		//nolint:lll
		// https://dzone.com/articles/getting-access-token-for-microsoft-graph-using-oau
		rawURL = fmt.Sprintf(
			"https://login.microsoftonline.com/%s/oauth2/v2.0/token",
			c.Credentials.AzureTenantID)
		headers = map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
		}
		body = strings.NewReader(fmt.Sprintf(
			"grant_type=urn:ietf:params:oauth:grant-type:device_code"+
				"&client_id=%s"+
				"&client_secret=%s"+
				"&device_code=%s",
			c.Credentials.AzureClientID,
			c.Credentials.AzureClientSecret,
			deviceCode))
	)

	fmt.Printf("\n-----\ndc %q\n-----\n", deviceCode)

	resp, err := c.Post(ctx, rawURL, headers, body)
	if err != nil {
		err = graph.Stack(ctx, err)
		fmt.Printf("\n-----\nERROR %+v\n-----\n", clues.ToCore(err))

		return deviceResp{}, err
	}

	if resp.StatusCode == http.StatusBadRequest {
		respDump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return deviceResp{}, clues.Wrap(err, "dumping http response")
		}

		fmt.Printf("\n-----\nresp %+v\n-----\n", string(respDump))

		return deviceResp{}, clues.New("incorrect tenant or credentials")
	}

	if resp.StatusCode/100 == 4 || resp.StatusCode/100 == 5 {
		respDump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return deviceResp{}, clues.Wrap(err, "dumping http response")
		}

		fmt.Printf("\n-----\nresp %+v\n-----\n", string(respDump))

		return deviceResp{}, clues.New("non-2xx response: " + resp.Status)
	}

	defer resp.Body.Close()

	var ar deviceResp
	err = json.NewDecoder(resp.Body).Decode(&ar)

	return ar, clues.Wrap(err, "undecodable resp body").WithClues(ctx).OrNil()
}
