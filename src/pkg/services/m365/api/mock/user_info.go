package mock

import (
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

func UserSettings() models.Userable {
	u := models.NewUser()

	u.SetAdditionalData(
		map[string]any{
			"archiveFolder":                         "archive",
			"timeZone":                              "UTC",
			"dateFormat":                            "MM/dd/yyyy",
			"timeFormat":                            "hh:mm tt",
			"userPurpose":                           "user",
			"delegateMeetingMessageDeliveryOptions": "test",
			"automaticRepliesSetting": map[string]any{
				"status":               "foo",
				"externalAudience":     "bar",
				"externalReplyMessage": "baz",
				"internalReplyMessage": "qux",
				"scheduledStartDateTime": map[string]any{
					"dateTime": "2020-01-01T00:00:00Z",
					"timeZone": "UTC",
				},
				"scheduledEndDateTime": map[string]any{
					"dateTime": "2020-01-01T00:00:00Z",
					"timeZone": "UTC",
				},
			},
			"language": map[string]any{
				"displayName": "en-US",
				"locale":      "US",
			},
			"workingHours": map[string]any{
				"daysOfWeek": []any{"monday"},
				"startTime":  "08:00:00.0000000",
				"endTime":    "17:00:00.0000000",
				"timeZone": map[string]any{
					"name": "UTC",
				},
			},
		})

	return u
}

func UserMailboxInfo() api.MailboxInfo {
	return api.MailboxInfo{
		Purpose:                    "user",
		ArchiveFolder:              "archive",
		DateFormat:                 "MM/dd/yyyy",
		TimeFormat:                 "hh:mm tt",
		DelegateMeetMsgDeliveryOpt: "test",
		Timezone:                   "UTC",
		AutomaticRepliesSetting: api.AutomaticRepliesSettings{
			Status:               "foo",
			ExternalAudience:     "bar",
			ExternalReplyMessage: "baz",
			InternalReplyMessage: "qux",
			ScheduledStartDateTime: api.TimeInfo{
				DateTime: "2020-01-01T00:00:00Z",
				Timezone: "UTC",
			},
			ScheduledEndDateTime: api.TimeInfo{
				DateTime: "2020-01-01T00:00:00Z",
				Timezone: "UTC",
			},
		},
		Language: api.Language{
			DisplayName: "en-US",
			Locale:      "US",
		},
		WorkingHours: api.WorkingHours{
			DaysOfWeek: []string{"monday"},
			StartTime:  "08:00:00.0000000",
			EndTime:    "17:00:00.0000000",
			TimeZone: struct {
				Name string
			}{
				Name: "UTC",
			},
		},

		ErrGetMailBoxSetting: []error{},
	}
}
