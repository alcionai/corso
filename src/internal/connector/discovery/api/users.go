package api

import (
	"context"

	absser "github.com/microsoft/kiota-abstractions-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/pkg/path"
)

// ---------------------------------------------------------------------------
// controller
// ---------------------------------------------------------------------------

func (c Client) Users() Users {
	return Users{c}
}

// Users is an interface-compliant provider of the client.
type Users struct {
	Client
}

// ---------------------------------------------------------------------------
// structs
// ---------------------------------------------------------------------------

type UserInfo struct {
	DiscoveredServices map[path.ServiceType]struct{}
}

func newUserInfo() *UserInfo {
	return &UserInfo{
		DiscoveredServices: map[path.ServiceType]struct{}{
			path.ExchangeService: {},
			path.OneDriveService: {},
		},
	}
}

// ---------------------------------------------------------------------------
// methods
// ---------------------------------------------------------------------------

const (
	userSelectID            = "id"
	userSelectPrincipalName = "userPrincipalName"
	userSelectDisplayName   = "displayName"
)

// Filter out both guest users, and (for on-prem installations) non-synced users.
// The latter filter makes an assumption that no on-prem users are guests; this might
// require more fine-tuned controls in the future.
// https://stackoverflow.com/questions/64044266/error-message-unsupported-or-invalid-query-filter-clause-specified-for-property
//
//nolint:lll
var userFilterNoGuests = "onPremisesSyncEnabled eq true OR userType ne 'Guest'"

// I can't believe I have to do this.
var t = true

func userOptions(fs *string) *users.UsersRequestBuilderGetRequestConfiguration {
	headers := absser.NewRequestHeaders()
	headers.Add("ConsistencyLevel", "eventual")

	return &users.UsersRequestBuilderGetRequestConfiguration{
		Headers: headers,
		QueryParameters: &users.UsersRequestBuilderGetQueryParameters{
			Select: []string{userSelectID, userSelectPrincipalName, userSelectDisplayName},
			Filter: fs,
			Count:  &t,
		},
	}
}

// GetAll retrieves all users.
func (c Users) GetAll(ctx context.Context) ([]models.Userable, error) {
	service, err := c.service()
	if err != nil {
		return nil, err
	}

	var resp models.UserCollectionResponseable

	err = graph.RunWithRetry(func() error {
		resp, err = service.Client().Users().Get(ctx, userOptions(&userFilterNoGuests))
		return err
	})

	if err != nil {
		return nil, support.ConnectorStackErrorTraceWrap(err, "getting all users")
	}

	iter, err := msgraphgocore.NewPageIterator(
		resp,
		service.Adapter(),
		models.CreateUserCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, support.ConnectorStackErrorTraceWrap(err, "constructing user iterator")
	}

	var (
		iterErrs error
		us       = make([]models.Userable, 0)
	)

	iterator := func(item any) bool {
		u, err := validateUser(item)
		if err != nil {
			iterErrs = support.WrapAndAppend("validating user", err, iterErrs)
		} else {
			us = append(us, u)
		}

		return true
	}

	if err := iter.Iterate(ctx, iterator); err != nil {
		return nil, support.ConnectorStackErrorTraceWrap(err, "iterating all users")
	}

	return us, iterErrs
}

func (c Users) GetByID(ctx context.Context, userID string) (models.Userable, error) {
	var (
		resp models.Userable
		err  error
	)

	err = graph.RunWithRetry(func() error {
		resp, err = c.stable.Client().UsersById(userID).Get(ctx, nil)
		return err
	})

	if err != nil {
		return nil, support.ConnectorStackErrorTraceWrap(err, "getting user by id")
	}

	return resp, err
}

func (c Users) GetInfo(ctx context.Context, userID string) (*UserInfo, error) {
	// Assume all services are enabled
	// then filter down to only services the user has enabled
	var (
		err      error
		userInfo = newUserInfo()
	)

	// TODO: OneDrive
	err = graph.RunWithRetry(func() error {
		_, err = c.stable.Client().UsersById(userID).MailFolders().Get(ctx, nil)
		return err
	})

	if err != nil {
		if !graph.IsErrExchangeMailFolderNotFound(err) {
			return nil, support.ConnectorStackErrorTraceWrap(err, "getting user's exchange mailfolders")
		}

		delete(userInfo.DiscoveredServices, path.ExchangeService)
	}

	return userInfo, nil
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

// validateUser ensures the item is a Userable, and contains the necessary
// identifiers that we handle with all users.
// returns the item as a Userable model.
func validateUser(item any) (models.Userable, error) {
	m, ok := item.(models.Userable)
	if !ok {
		return nil, errors.Errorf("expected Userable, got %T", item)
	}

	if m.GetId() == nil {
		return nil, errors.Errorf("missing ID")
	}

	if m.GetUserPrincipalName() == nil {
		return nil, errors.New("missing principalName")
	}

	return m, nil
}
