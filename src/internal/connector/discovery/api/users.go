package api

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
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
	HasMailBox         bool
	HasOneDrive        bool
	Mailbox            mailboxInfo
}

type mailboxInfo struct {
	Purpose                    string
	ArchiveFolder              string
	DateFormat                 string
	TimeFormat                 string
	DelegateMeetMsgDeliveryOpt string
	Timezone                   string
	AutomaticRepliesSetting    AutomaticRepliesSettings
	Language                   Language
	WorkingHours               WorkingHours
	ErrGetMailBoxSetting       clues.Err
}

type AutomaticRepliesSettings struct {
	ExternalAudience       string
	ExternalReplyMessage   string
	InternalReplyMessage   string
	ScheduledEndDateTime   timeInfo
	ScheduledStartDateTime timeInfo
	Status                 string
}

type timeInfo struct {
	DateTime string
	Timezone string
}

type Language struct {
	Locale      string
	DisplayName string
}

type WorkingHours struct {
	DaysOfWeek []string
	StartTime  string
	EndTime    string
	TimeZone   struct {
		Name string
	}
}

func newUserInfo() *UserInfo {
	return &UserInfo{
		DiscoveredServices: map[path.ServiceType]struct{}{
			path.ExchangeService: {},
			path.OneDriveService: {},
		},
	}
}

// ServiceEnabled returns true if the UserInfo has an entry for the
// service.  If no entry exists, the service is assumed to not be enabled.
func (ui *UserInfo) ServiceEnabled(service path.ServiceType) bool {
	if ui == nil || len(ui.DiscoveredServices) == 0 {
		return false
	}

	_, ok := ui.DiscoveredServices[service]

	return ok
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
// ne 'Guest' ensures we don't filter out users where userType = null, which can happen
// for user accounts created prior to 2014.  In order to use the `ne` comparator, we
// MUST include $count=true and the ConsistencyLevel: eventual header.
// https://stackoverflow.com/questions/49340485/how-to-filter-users-by-usertype-null
//
//nolint:lll
var userFilterNoGuests = "onPremisesSyncEnabled eq true OR userType ne 'Guest'"

// I can't believe I have to do this.
var t = true

func userOptions(fs *string) *users.UsersRequestBuilderGetRequestConfiguration {
	headers := abstractions.NewRequestHeaders()
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
func (c Users) GetAll(ctx context.Context, errs *fault.Bus) ([]models.Userable, error) {
	service, err := c.service()
	if err != nil {
		return nil, err
	}

	var resp models.UserCollectionResponseable

	resp, err = service.Client().Users().Get(ctx, userOptions(&userFilterNoGuests))

	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting all users")
	}

	iter, err := msgraphgocore.NewPageIterator(
		resp,
		service.Adapter(),
		models.CreateUserCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating users iterator")
	}

	var (
		us = make([]models.Userable, 0)
		el = errs.Local()
	)

	iterator := func(item any) bool {
		if el.Failure() != nil {
			return false
		}

		u, err := validateUser(item)
		if err != nil {
			el.AddRecoverable(graph.Wrap(ctx, err, "validating user"))
		} else {
			us = append(us, u)
		}

		return true
	}

	if err := iter.Iterate(ctx, iterator); err != nil {
		return nil, graph.Wrap(ctx, err, "iterating all users")
	}

	return us, el.Failure()
}

// GetByID looks up the user matching the given identifier.  The identifier can be either a
// canonical user id or a princpalName.
func (c Users) GetByID(ctx context.Context, identifier string) (models.Userable, error) {
	var (
		resp models.Userable
		err  error
	)

	resp, err = c.stable.Client().UsersById(identifier).Get(ctx, nil)

	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting user")
	}

	return resp, err
}

// GetIDAndName looks up the user matching the given ID, and returns
// its canonical ID and the PrincipalName as the name.
func (c Users) GetIDAndName(ctx context.Context, userID string) (string, string, error) {
	u, err := c.GetByID(ctx, userID)
	if err != nil {
		return "", "", err
	}

	return ptr.Val(u.GetId()), ptr.Val(u.GetUserPrincipalName()), nil
}

func (c Users) GetInfo(ctx context.Context, userID string) (*UserInfo, error) {
	// Assume all services are enabled
	// then filter down to only services the user has enabled
	var (
		err      error
		userInfo = newUserInfo()
	)

	err = c.allowsExchange(ctx, userID, userInfo)
	if err != nil {
		return nil, err
	}

	err = c.allowsOnedrive(ctx, userID, userInfo)
	if err != nil {
		return nil, err
	}

	err = c.getAdditionalData(ctx, userID, userInfo)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}

// verify mailbox enabled for user
func (c Users) allowsExchange(ctx context.Context, userID string, userInfo *UserInfo) error {
	userInfo.HasMailBox = true

	_, err := c.stable.Client().UsersById(userID).MailFolders().Get(ctx, nil)
	if err != nil {
		if !graph.IsErrExchangeMailFolderNotFound(err) {
			logger.Ctx(ctx).Errorf("err getting user's mail folder: %s", err)
			return graph.Wrap(ctx, err, "getting user's mail folder")
		}

		logger.Ctx(ctx).Infof("resource owner does not have a mailbox enabled")
		delete(userInfo.DiscoveredServices, path.ExchangeService)

		userInfo.HasMailBox = false
	}

	return nil
}

// verify onedrive enabled for user
func (c Users) allowsOnedrive(ctx context.Context, userID string, userInfo *UserInfo) error {
	userInfo.HasOneDrive = true

	_, err := c.stable.Client().UsersById(userID).Drives().Get(ctx, nil)
	if err != nil {
		err = graph.Stack(ctx, err)

		if !clues.HasLabel(err, graph.LabelsMysiteNotFound) {
			logger.Ctx(ctx).Errorf("err getting user's onedrive's data: %s", err)

			return graph.Wrap(ctx, err, "getting user's onedrive's data")
		}

		logger.Ctx(ctx).Infof("resource owner does not have a drive")

		// TODO: add delete onedrive serve
		// delete(userInfo.DiscoveredServices, path.OneDriveService)
		userInfo.HasOneDrive = false
	}

	return nil
}

func (c Users) getAdditionalData(ctx context.Context, userID string, userInfo *UserInfo) error {
	var (
		rawURL                     = fmt.Sprintf("https://graph.microsoft.com/v1.0/users/%s/mailboxSettings", userID)
		adapter                    = c.stable.Adapter()
		builder                    = users.NewUserItemRequestBuilder(rawURL, adapter)
		ErrMailBoxSettingsNotFound = clues.New("mailbox settings not found")
		mailBoundErr               clues.Err
	)

	newItem, err := builder.Get(ctx, nil)
	if err != nil && !(graph.IsErrAccessDenied(err) || graph.IsErrExchangeMailFolderNotFound(err)) {
		logger.Ctx(ctx).Errorf("err getting additional data: %s", err)

		return graph.Wrap(ctx, err, "getting additional data")
	}

	if graph.IsErrAccessDenied(err) {
		logger.Ctx(ctx).Infof("err getting additional data: access denied")

		userInfo.Mailbox.ErrGetMailBoxSetting = *clues.New("access denied")

		return nil
	}

	if graph.IsErrExchangeMailFolderNotFound(err) {
		logger.Ctx(ctx).Infof("err exchange mail folder not found")

		userInfo.Mailbox.ErrGetMailBoxSetting = *ErrMailBoxSettingsNotFound
		userInfo.HasMailBox = false

		return nil
	}

	additionalData := newItem.GetAdditionalData()

	userInfo.Mailbox.ArchiveFolder = toString(ctx, additionalData["archiveFolder"], &mailBoundErr)

	userInfo.Mailbox.Timezone = toString(ctx, additionalData["timeZone"], &mailBoundErr)

	userInfo.Mailbox.DateFormat = toString(ctx, additionalData["dateFormat"], &mailBoundErr)
	userInfo.Mailbox.TimeFormat = toString(ctx, additionalData["timeFormat"], &mailBoundErr)
	userInfo.Mailbox.Purpose = toString(ctx, additionalData["userPurpose"], &mailBoundErr)
	userInfo.Mailbox.DelegateMeetMsgDeliveryOpt = toString(
		ctx,
		additionalData["delegateMeetingMessageDeliveryOptions"],
		&mailBoundErr)

	// decode automatic replies settings
	replySetting := toMap(ctx, additionalData["automaticRepliesSetting"], &mailBoundErr)
	userInfo.Mailbox.AutomaticRepliesSetting.Status = toString(
		ctx,
		replySetting["status"],
		&mailBoundErr)
	userInfo.Mailbox.AutomaticRepliesSetting.ExternalAudience = toString(
		ctx,
		replySetting["externalAudience"],
		&mailBoundErr)
	userInfo.Mailbox.AutomaticRepliesSetting.ExternalReplyMessage = toString(
		ctx,
		replySetting["externalReplyMessage"],
		&mailBoundErr)
	userInfo.Mailbox.AutomaticRepliesSetting.InternalReplyMessage = toString(
		ctx,
		replySetting["internalReplyMessage"],
		&mailBoundErr)

	// decode scheduledStartDateTime
	startDateTime := toMap(ctx, replySetting["scheduledStartDateTime"], &mailBoundErr)
	userInfo.Mailbox.AutomaticRepliesSetting.ScheduledStartDateTime.DateTime = toString(
		ctx,
		startDateTime["dateTime"],
		&mailBoundErr)
	userInfo.Mailbox.AutomaticRepliesSetting.ScheduledStartDateTime.Timezone = toString(
		ctx,
		startDateTime["timeZone"],
		&mailBoundErr)

	endDateTime := toMap(ctx, replySetting["scheduledEndDateTime"], &mailBoundErr)
	userInfo.Mailbox.AutomaticRepliesSetting.ScheduledEndDateTime.DateTime = toString(
		ctx,
		endDateTime["dateTime"],
		&mailBoundErr)
	userInfo.Mailbox.AutomaticRepliesSetting.ScheduledEndDateTime.Timezone = toString(
		ctx,
		endDateTime["timeZone"],
		&mailBoundErr)

	// Language decode
	language := toMap(ctx, additionalData["language"], &mailBoundErr)
	userInfo.Mailbox.Language.DisplayName = toString(
		ctx,
		language["displayName"],
		&mailBoundErr)
	userInfo.Mailbox.Language.Locale = toString(ctx, language["locale"], &mailBoundErr)

	// working hours
	workingHours := toMap(ctx, additionalData["workingHours"], &mailBoundErr)
	userInfo.Mailbox.WorkingHours.StartTime = toString(
		ctx,
		workingHours["startTime"],
		&mailBoundErr)
	userInfo.Mailbox.WorkingHours.EndTime = toString(
		ctx,
		workingHours["endTime"],
		&mailBoundErr)

	timeZone := toMap(ctx, workingHours["timeZone"], &mailBoundErr)
	userInfo.Mailbox.WorkingHours.TimeZone.Name = toString(
		ctx,
		timeZone["name"],
		&mailBoundErr)

	days := toArray(ctx, workingHours["daysOfWeek"], &mailBoundErr)
	for _, day := range days {
		userInfo.Mailbox.WorkingHours.DaysOfWeek = append(userInfo.Mailbox.WorkingHours.DaysOfWeek,
			toString(ctx, day, &mailBoundErr))
	}

	if mailBoundErr.Error() != "" {
		userInfo.Mailbox.ErrGetMailBoxSetting = mailBoundErr
	}

	return nil
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
		return nil, clues.New(fmt.Sprintf("unexpected model: %T", item))
	}

	if m.GetId() == nil {
		return nil, clues.New("missing ID")
	}

	if m.GetUserPrincipalName() == nil {
		return nil, clues.New("missing principalName")
	}

	return m, nil
}

func toString(ctx context.Context, data any, mailBoxErr *clues.Err) string {
	ErrMailBoxSettingsNotFound := *clues.New("mailbox settings not found")

	dataPointer, ok := data.(*string)
	if !ok {
		logger.Ctx(ctx).Infof("error getting data from mailboxSettings")

		*mailBoxErr = ErrMailBoxSettingsNotFound

		return ""
	}

	value, ok := ptr.ValOK(dataPointer)
	if !ok {
		logger.Ctx(ctx).Infof("error getting value from pointer for mailboxSettings")

		*mailBoxErr = ErrMailBoxSettingsNotFound

		return ""
	}

	return value
}

func toMap(ctx context.Context, data any, mailBoxErr *clues.Err) map[string]interface{} {
	value, ok := data.(map[string]interface{})
	if !ok {
		logger.Ctx(ctx).Infof("error getting mailboxSettings")

		*mailBoxErr = *clues.New("mailbox settings not found")

		return value
	}

	return value
}

func toArray(ctx context.Context, data any, mailBoxErr *clues.Err) []interface{} {
	value, ok := data.([]interface{})
	if !ok {
		logger.Ctx(ctx).Infof("error getting mailboxSettings")

		*mailBoxErr = *clues.New("mailbox settings not found")

		return value
	}

	return value
}
