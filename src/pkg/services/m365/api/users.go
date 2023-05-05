package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/alcionai/clues"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

// Variables
var (
	ErrMailBoxSettingsNotFound = clues.New("mailbox settings not found")
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
	ServicesEnabled map[path.ServiceType]struct{}
	Mailbox         MailboxInfo
}

type MailboxInfo struct {
	Purpose                    string
	ArchiveFolder              string
	DateFormat                 string
	TimeFormat                 string
	DelegateMeetMsgDeliveryOpt string
	Timezone                   string
	AutomaticRepliesSetting    AutomaticRepliesSettings
	Language                   Language
	WorkingHours               WorkingHours
	ErrGetMailBoxSetting       []error
	QuotaExceeded              bool
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
		ServicesEnabled: map[path.ServiceType]struct{}{
			path.ExchangeService: {},
			path.OneDriveService: {},
		},
	}
}

// ServiceEnabled returns true if the UserInfo has an entry for the
// service.  If no entry exists, the service is assumed to not be enabled.
func (ui *UserInfo) ServiceEnabled(service path.ServiceType) bool {
	if ui == nil || len(ui.ServicesEnabled) == 0 {
		return false
	}

	_, ok := ui.ServicesEnabled[service]

	return ok
}

// Returns if we can run delta queries on a mailbox. We cannot run
// them if the mailbox is full which is indicated by QuotaExceeded.
func (ui *UserInfo) CanMakeDeltaQueries() bool {
	return !ui.Mailbox.QuotaExceeded
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
	service, err := c.Service()
	if err != nil {
		return nil, err
	}

	var resp models.UserCollectionResponseable

	resp, err = service.Client().Users().Get(ctx, userOptions(&userFilterNoGuests))

	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting all users")
	}

	iter, err := msgraphgocore.NewPageIterator[models.Userable](
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

	iterator := func(item models.Userable) bool {
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

	resp, err = c.stable.Client().Users().ByUserId(identifier).Get(ctx, nil)

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

// GetAllIDsAndNames retrieves all users in the tenant and returns them in an idname.Cacher
func (c Users) GetAllIDsAndNames(ctx context.Context, errs *fault.Bus) (idname.Cacher, error) {
	all, err := c.GetAll(ctx, errs)
	if err != nil {
		return nil, clues.Wrap(err, "getting all users")
	}

	idToName := make(map[string]string, len(all))

	for _, u := range all {
		id := strings.ToLower(ptr.Val(u.GetId()))
		name := strings.ToLower(ptr.Val(u.GetUserPrincipalName()))

		idToName[id] = name
	}

	return idname.NewCache(idToName), nil
}

func (c Users) GetInfo(ctx context.Context, userID string) (*UserInfo, error) {
	// Assume all services are enabled
	// then filter down to only services the user has enabled
	userInfo := newUserInfo()

	requestParameters := users.ItemMailFoldersRequestBuilderGetQueryParameters{
		Select: []string{"id"},
		Top:    ptr.To[int32](1), // if we get any folders, then we have access.
	}

	options := users.ItemMailFoldersRequestBuilderGetRequestConfiguration{
		QueryParameters: &requestParameters,
	}

	mfs, err := c.GetMailFolders(ctx, userID, options)
	if err != nil {
		if graph.IsErrUserNotFound(err) {
			logger.CtxErr(ctx, err).Error("user not found")
			return nil, graph.Stack(ctx, clues.Stack(graph.ErrResourceOwnerNotFound, err))
		}

		if !graph.IsErrExchangeMailFolderNotFound(err) ||
			clues.HasLabel(err, graph.LabelStatus(http.StatusNotFound)) {
			logger.CtxErr(ctx, err).Error("getting user's mail folder")
			return nil, err
		}

		logger.Ctx(ctx).Info("resource owner does not have a mailbox enabled")
		delete(userInfo.ServicesEnabled, path.ExchangeService)
	}

	if _, err := c.GetDrives(ctx, userID); err != nil {
		if !clues.HasLabel(err, graph.LabelsMysiteNotFound) {
			logger.CtxErr(ctx, err).Error("getting user's drives")

			return nil, graph.Wrap(ctx, err, "getting user's drives")
		}

		logger.Ctx(ctx).Info("resource owner does not have a drive")

		delete(userInfo.ServicesEnabled, path.OneDriveService)
	}

	mbxInfo, err := c.getMailboxSettings(ctx, userID)
	if err != nil {
		return nil, err
	}

	userInfo.Mailbox = mbxInfo

	// TODO: This tries to determine if the user has hit their mailbox
	// limit by trying to fetch an item and seeing if we get the quota
	// exceeded error. Ideally(if available) we should convert this to
	// pull the user's usage via an api and compare if they have used
	// up their quota.
	if mfs != nil {
		mf := mfs.GetValue()[0] // we will always have one
		options := &users.ItemMailFoldersItemMessagesDeltaRequestBuilderGetRequestConfiguration{
			QueryParameters: &users.ItemMailFoldersItemMessagesDeltaRequestBuilderGetQueryParameters{
				Top: ptr.To[int32](1), // just one item is enough
			},
		}
		_, err = c.stable.Client().
			Users().ByUserId(userID).
			MailFolders().ByMailFolderId(ptr.Val(mf.GetId())).
			Messages().
			Delta().
			Get(ctx, options)

		if err != nil && !graph.IsErrQuotaExceeded(err) {
			return nil, err
		}

		userInfo.Mailbox.QuotaExceeded = graph.IsErrQuotaExceeded(err)
	}

	return userInfo, nil
}

// TODO: remove when exchange api goes into this package
func (c Users) GetMailFolders(
	ctx context.Context,
	userID string,
	options users.ItemMailFoldersRequestBuilderGetRequestConfiguration,
) (models.MailFolderCollectionResponseable, error) {
	mailFolders, err := c.stable.Client().Users().ByUserId(userID).MailFolders().Get(ctx, &options)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting MailFolders")
	}

	return mailFolders, nil
}

// TODO: remove when drive api goes into this package
func (c Users) GetDrives(ctx context.Context, userID string) (models.DriveCollectionResponseable, error) {
	drives, err := c.stable.Client().Users().ByUserId(userID).Drives().Get(ctx, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting drives")
	}

	return drives, nil
}

func (c Users) getMailboxSettings(
	ctx context.Context,
	userID string,
) (MailboxInfo, error) {
	var (
		rawURL  = fmt.Sprintf("https://graph.microsoft.com/v1.0/users/%s/mailboxSettings", userID)
		adapter = c.stable.Adapter()
		mi      = MailboxInfo{
			ErrGetMailBoxSetting: []error{},
		}
	)

	settings, err := users.NewUserItemRequestBuilder(rawURL, adapter).Get(ctx, nil)
	if err != nil && !(graph.IsErrAccessDenied(err) || graph.IsErrExchangeMailFolderNotFound(err)) {
		logger.CtxErr(ctx, err).Error("getting mailbox settings")
		return mi, graph.Wrap(ctx, err, "getting additional data")
	}

	if graph.IsErrAccessDenied(err) {
		logger.Ctx(ctx).Info("err getting additional data: access denied")

		mi.ErrGetMailBoxSetting = append(mi.ErrGetMailBoxSetting, clues.New("access denied"))

		return mi, nil
	}

	if graph.IsErrExchangeMailFolderNotFound(err) {
		logger.Ctx(ctx).Info("mailfolders not found")

		mi.ErrGetMailBoxSetting = append(mi.ErrGetMailBoxSetting, ErrMailBoxSettingsNotFound)

		return mi, nil
	}

	additionalData := settings.GetAdditionalData()

	mi.ArchiveFolder, err = toString(ctx, "archiveFolder", additionalData)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.Timezone, err = toString(ctx, "timeZone", additionalData)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.DateFormat, err = toString(ctx, "dateFormat", additionalData)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.TimeFormat, err = toString(ctx, "timeFormat", additionalData)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.Purpose, err = toString(ctx, "userPurpose", additionalData)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.DelegateMeetMsgDeliveryOpt, err = toString(ctx, "delegateMeetingMessageDeliveryOptions", additionalData)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	// decode automatic replies settings
	replySetting, err := toT[map[string]any](ctx, "automaticRepliesSetting", additionalData)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.AutomaticRepliesSetting.Status, err = toString(ctx, "status", replySetting)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.AutomaticRepliesSetting.ExternalAudience, err = toString(ctx, "externalAudience", replySetting)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.AutomaticRepliesSetting.ExternalReplyMessage, err = toString(ctx, "externalReplyMessage", replySetting)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.AutomaticRepliesSetting.InternalReplyMessage, err = toString(ctx, "internalReplyMessage", replySetting)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	// decode scheduledStartDateTime
	startDateTime, err := toT[map[string]any](ctx, "scheduledStartDateTime", replySetting)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.AutomaticRepliesSetting.ScheduledStartDateTime.DateTime, err = toString(ctx, "dateTime", startDateTime)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.AutomaticRepliesSetting.ScheduledStartDateTime.Timezone, err = toString(ctx, "timeZone", startDateTime)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	endDateTime, err := toT[map[string]any](ctx, "scheduledEndDateTime", replySetting)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.AutomaticRepliesSetting.ScheduledEndDateTime.DateTime, err = toString(ctx, "dateTime", endDateTime)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.AutomaticRepliesSetting.ScheduledEndDateTime.Timezone, err = toString(ctx, "timeZone", endDateTime)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	// Language decode
	language, err := toT[map[string]any](ctx, "language", additionalData)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.Language.DisplayName, err = toString(ctx, "displayName", language)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.Language.Locale, err = toString(ctx, "locale", language)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	// working hours
	workingHours, err := toT[map[string]any](ctx, "workingHours", additionalData)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.WorkingHours.StartTime, err = toString(ctx, "startTime", workingHours)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.WorkingHours.EndTime, err = toString(ctx, "endTime", workingHours)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	timeZone, err := toT[map[string]any](ctx, "timeZone", workingHours)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.WorkingHours.TimeZone.Name, err = toString(ctx, "name", timeZone)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	days, err := toT[[]any](ctx, "daysOfWeek", workingHours)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	for _, day := range days {
		s, err := anyToString(ctx, "dayOfTheWeek", day)
		mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)
		mi.WorkingHours.DaysOfWeek = append(mi.WorkingHours.DaysOfWeek, s)
	}

	return mi, nil
}

func appendIfErr(errs []error, err error) []error {
	if err == nil {
		return errs
	}

	return append(errs, err)
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

// validateUser ensures the item is a Userable, and contains the necessary
// identifiers that we handle with all users.
// returns the item as a Userable model.
// TODO(meain): no need to return item
func validateUser(item models.Userable) (models.Userable, error) {
	if item.GetId() == nil {
		return nil, clues.New("missing ID")
	}

	if item.GetUserPrincipalName() == nil {
		return nil, clues.New("missing principalName")
	}

	return item, nil
}

func toString(ctx context.Context, key string, data map[string]any) (string, error) {
	ctx = clues.Add(ctx, "setting_name", key)

	if len(data) == 0 {
		logger.Ctx(ctx).Info("not found: ", key)
		return "", ErrMailBoxSettingsNotFound
	}

	return anyToString(ctx, key, data[key])
}

func anyToString(ctx context.Context, key string, val any) (string, error) {
	if val == nil {
		logger.Ctx(ctx).Info("nil value: ", key)
		return "", ErrMailBoxSettingsNotFound
	}

	sp, ok := val.(*string)
	if !ok {
		logger.Ctx(ctx).Info("value is not a *string: ", key)
		return "", ErrMailBoxSettingsNotFound
	}

	return ptr.Val(sp), nil
}

func toT[T any](ctx context.Context, key string, data map[string]any) (T, error) {
	ctx = clues.Add(ctx, "setting_name", key)

	if len(data) == 0 {
		logger.Ctx(ctx).Info("not found: ", key)
		return *new(T), ErrMailBoxSettingsNotFound
	}

	val := data[key]

	if data == nil {
		logger.Ctx(ctx).Info("nil value: ", key)
		return *new(T), ErrMailBoxSettingsNotFound
	}

	value, ok := val.(T)
	if !ok {
		logger.Ctx(ctx).Info(fmt.Sprintf("unexpected type for %s: %T", key, val))
		return *new(T), ErrMailBoxSettingsNotFound
	}

	return value, nil
}
