package api

import (
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/internal/common/tform"
	"github.com/alcionai/corso/src/pkg/path"
)

// ---------------------------------------------------------------------------
// User Info
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
	ScheduledEndDateTime   TimeInfo
	ScheduledStartDateTime TimeInfo
	Status                 string
}

type TimeInfo struct {
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

func ParseMailboxSettings(
	settings models.Userable,
	mi MailboxInfo,
) MailboxInfo {
	var (
		additionalData = settings.GetAdditionalData()
		err            error
	)

	mi.ArchiveFolder, err = str.AnyValueToString("archiveFolder", additionalData)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.Timezone, err = str.AnyValueToString("timeZone", additionalData)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.DateFormat, err = str.AnyValueToString("dateFormat", additionalData)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.TimeFormat, err = str.AnyValueToString("timeFormat", additionalData)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.Purpose, err = str.AnyValueToString("userPurpose", additionalData)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.DelegateMeetMsgDeliveryOpt, err = str.AnyValueToString("delegateMeetingMessageDeliveryOptions", additionalData)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	// decode automatic replies settings
	replySetting, err := tform.AnyValueToT[map[string]any]("automaticRepliesSetting", additionalData)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.AutomaticRepliesSetting.Status, err = str.AnyValueToString("status", replySetting)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.AutomaticRepliesSetting.ExternalAudience, err = str.AnyValueToString("externalAudience", replySetting)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.AutomaticRepliesSetting.ExternalReplyMessage, err = str.AnyValueToString("externalReplyMessage", replySetting)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.AutomaticRepliesSetting.InternalReplyMessage, err = str.AnyValueToString("internalReplyMessage", replySetting)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	// decode scheduledStartDateTime
	startDateTime, err := tform.AnyValueToT[map[string]any]("scheduledStartDateTime", replySetting)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.AutomaticRepliesSetting.ScheduledStartDateTime.DateTime, err = str.AnyValueToString("dateTime", startDateTime)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.AutomaticRepliesSetting.ScheduledStartDateTime.Timezone, err = str.AnyValueToString("timeZone", startDateTime)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	endDateTime, err := tform.AnyValueToT[map[string]any]("scheduledEndDateTime", replySetting)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.AutomaticRepliesSetting.ScheduledEndDateTime.DateTime, err = str.AnyValueToString("dateTime", endDateTime)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.AutomaticRepliesSetting.ScheduledEndDateTime.Timezone, err = str.AnyValueToString("timeZone", endDateTime)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	// Language decode
	language, err := tform.AnyValueToT[map[string]any]("language", additionalData)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.Language.DisplayName, err = str.AnyValueToString("displayName", language)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.Language.Locale, err = str.AnyValueToString("locale", language)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	// working hours
	workingHours, err := tform.AnyValueToT[map[string]any]("workingHours", additionalData)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.WorkingHours.StartTime, err = str.AnyValueToString("startTime", workingHours)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.WorkingHours.EndTime, err = str.AnyValueToString("endTime", workingHours)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	timeZone, err := tform.AnyValueToT[map[string]any]("timeZone", workingHours)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	mi.WorkingHours.TimeZone.Name, err = str.AnyValueToString("name", timeZone)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	days, err := tform.AnyValueToT[[]any]("daysOfWeek", workingHours)
	mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)

	for _, day := range days {
		s, err := str.AnyToString(day)
		mi.ErrGetMailBoxSetting = appendIfErr(mi.ErrGetMailBoxSetting, err)
		mi.WorkingHours.DaysOfWeek = append(mi.WorkingHours.DaysOfWeek, s)
	}

	return mi
}
