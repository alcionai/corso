package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// EmailThreatSubmissionPolicyable 
type EmailThreatSubmissionPolicyable interface {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomizedNotificationSenderEmailAddress()(*string)
    GetCustomizedReportRecipientEmailAddress()(*string)
    GetIsAlwaysReportEnabledForUsers()(*bool)
    GetIsAskMeEnabledForUsers()(*bool)
    GetIsCustomizedMessageEnabled()(*bool)
    GetIsCustomizedMessageEnabledForPhishing()(*bool)
    GetIsCustomizedNotificationSenderEnabled()(*bool)
    GetIsNeverReportEnabledForUsers()(*bool)
    GetIsOrganizationBrandingEnabled()(*bool)
    GetIsReportFromQuarantineEnabled()(*bool)
    GetIsReportToCustomizedEmailAddressEnabled()(*bool)
    GetIsReportToMicrosoftEnabled()(*bool)
    GetIsReviewEmailNotificationEnabled()(*bool)
    SetCustomizedNotificationSenderEmailAddress(value *string)()
    SetCustomizedReportRecipientEmailAddress(value *string)()
    SetIsAlwaysReportEnabledForUsers(value *bool)()
    SetIsAskMeEnabledForUsers(value *bool)()
    SetIsCustomizedMessageEnabled(value *bool)()
    SetIsCustomizedMessageEnabledForPhishing(value *bool)()
    SetIsCustomizedNotificationSenderEnabled(value *bool)()
    SetIsNeverReportEnabledForUsers(value *bool)()
    SetIsOrganizationBrandingEnabled(value *bool)()
    SetIsReportFromQuarantineEnabled(value *bool)()
    SetIsReportToCustomizedEmailAddressEnabled(value *bool)()
    SetIsReportToMicrosoftEnabled(value *bool)()
    SetIsReviewEmailNotificationEnabled(value *bool)()
}
