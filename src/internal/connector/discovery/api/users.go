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
	DiscoveredServices         map[path.ServiceType]struct{}
	Purpose                    string
	ArchiveFolder              string
	DateFormat                 string
	TimeFormat                 string
	DelegateMeetMsgDeliveryOpt string
	Timezone                   string
	HasMailBox                 bool
	HasOneDrive                bool
	AutomaticRepliesSetting    AutomaticRepliesSettings
	Language                   Language
	WorkingHours               WorkingHours
	ErrGetMailBoxSetting       string
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
		rawURL  = fmt.Sprintf("https://graph.microsoft.com/v1.0/users/%s/mailboxSettings", userID)
		adapter = c.stable.Adapter()
		builder = users.NewUserItemRequestBuilder(rawURL, adapter)

		additionalData, replySetting, startDateTime, endDateTime,
		language, timeZone, workingHours map[string]interface{}

		days []interface{}
	)

	newItem, err := builder.Get(ctx, nil)
	if err != nil && !(graph.IsErrAccessDenied(err) || graph.IsErrExchangeMailFolderNotFound(err)) {
		logger.Ctx(ctx).Errorf("err getting additional data: %s", err)

		return graph.Wrap(ctx, err, "getting additional data")
	}

	if graph.IsErrAccessDenied(err) {
		logger.Ctx(ctx).Infof("err getting additional data: access denied")

		userInfo.ErrGetMailBoxSetting = "access denied"

		return nil
	}

	if graph.IsErrExchangeMailFolderNotFound(err) {
		logger.Ctx(ctx).Infof("err exchange mail folder not found")

		userInfo.ErrGetMailBoxSetting = "not found"
		userInfo.HasMailBox = false

		return nil
	}

	additionalData = newItem.GetAdditionalData()

	userInfo.ArchiveFolder = convertInterfaceToString(ctx, additionalData["archiveFolder"], userInfo)
	userInfo.Timezone = convertInterfaceToString(ctx, additionalData["timeZone"], userInfo)
	userInfo.DateFormat = convertInterfaceToString(ctx, additionalData["dateFormat"], userInfo)
	userInfo.TimeFormat = convertInterfaceToString(ctx, additionalData["timeFormat"], userInfo)
	userInfo.Purpose = convertInterfaceToString(ctx, additionalData["userPurpose"], userInfo)
	userInfo.DelegateMeetMsgDeliveryOpt = convertInterfaceToString(
		ctx,
		additionalData["delegateMeetingMessageDeliveryOptions"],
		userInfo)

	// decode automatic replies settings
	replySetting = convertInterfaceToMap(ctx, additionalData["automaticRepliesSetting"], userInfo)
	userInfo.AutomaticRepliesSetting.Status = convertInterfaceToString(ctx, replySetting["status"], userInfo)
	userInfo.AutomaticRepliesSetting.ExternalAudience = convertInterfaceToString(
		ctx,
		replySetting["externalAudience"],
		userInfo)
	userInfo.AutomaticRepliesSetting.ExternalReplyMessage = convertInterfaceToString(
		ctx,
		replySetting["externalReplyMessage"],
		userInfo)
	userInfo.AutomaticRepliesSetting.InternalReplyMessage = convertInterfaceToString(
		ctx,
		replySetting["internalReplyMessage"],
		userInfo)

	// decode scheduledStartDateTime
	startDateTime = convertInterfaceToMap(ctx, replySetting["scheduledStartDateTime"], userInfo)
	userInfo.AutomaticRepliesSetting.ScheduledStartDateTime.DateTime = convertInterfaceToString(
		ctx,
		startDateTime["dateTime"],
		userInfo)
	userInfo.AutomaticRepliesSetting.ScheduledStartDateTime.Timezone = convertInterfaceToString(
		ctx,
		startDateTime["timeZone"],
		userInfo)

	endDateTime = convertInterfaceToMap(ctx, replySetting["scheduledEndDateTime"], userInfo)
	userInfo.AutomaticRepliesSetting.ScheduledEndDateTime.DateTime = convertInterfaceToString(
		ctx,
		endDateTime["dateTime"],
		userInfo)
	userInfo.AutomaticRepliesSetting.ScheduledEndDateTime.Timezone = convertInterfaceToString(
		ctx,
		endDateTime["timeZone"],
		userInfo)

	// Language decode
	language = convertInterfaceToMap(ctx, additionalData["language"], userInfo)
	userInfo.Language.DisplayName = convertInterfaceToString(ctx, language["displayName"], userInfo)
	userInfo.Language.Locale = convertInterfaceToString(ctx, language["locale"], userInfo)

	// working hours
	workingHours = convertInterfaceToMap(ctx, additionalData["workingHours"], userInfo)
	userInfo.WorkingHours.StartTime = convertInterfaceToString(ctx, workingHours["startTime"], userInfo)
	userInfo.WorkingHours.EndTime = convertInterfaceToString(ctx, workingHours["endTime"], userInfo)
	timeZone = convertInterfaceToMap(ctx, workingHours["timeZone"], userInfo)
	userInfo.WorkingHours.TimeZone.Name = convertInterfaceToString(ctx, timeZone["name"], userInfo)

	days = convertInterfaceToArray(ctx, workingHours["daysOfWeek"], userInfo)
	for _, day := range days {
		userInfo.WorkingHours.DaysOfWeek = append(userInfo.WorkingHours.DaysOfWeek,
			convertInterfaceToString(ctx, day, userInfo))
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

func convertInterfaceToString(ctx context.Context, data interface{}, userInfo *UserInfo) string {
	var (
		ok          bool
		dataPointer *string
		value       string
	)

	dataPointer, ok = data.(*string)
	if !ok {
		logger.Ctx(ctx).Infof("error getting mailboxSettings")

		userInfo.ErrGetMailBoxSetting = "not found"

		return ""
	}

	value, ok = ptr.ValOK(dataPointer)
	if !ok {
		logger.Ctx(ctx).Infof("error getting mailboxSettings")

		userInfo.ErrGetMailBoxSetting = "not found"

		return ""
	}

	return value
}

func convertInterfaceToMap(ctx context.Context, data interface{}, userInfo *UserInfo) map[string]interface{} {
	var (
		ok    bool
		value map[string]interface{}
	)

	value, ok = data.(map[string]interface{})
	if !ok {
		logger.Ctx(ctx).Infof("error getting mailboxSettings")

		userInfo.ErrGetMailBoxSetting = "not found"

		return value
	}

	return value
}

func convertInterfaceToArray(ctx context.Context, data interface{}, userInfo *UserInfo) []interface{} {
	var (
		ok    bool
		value []interface{}
	)

	value, ok = data.([]interface{})
	if !ok {
		logger.Ctx(ctx).Infof("error getting mailboxSettings")

		userInfo.ErrGetMailBoxSetting = "not found"

		return value
	}

	return value
}
