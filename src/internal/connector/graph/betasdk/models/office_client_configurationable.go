package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OfficeClientConfigurationable 
type OfficeClientConfigurationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignments()([]OfficeClientConfigurationAssignmentable)
    GetCheckinStatuses()([]OfficeClientCheckinStatusable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetPolicyPayload()([]byte)
    GetPriority()(*int32)
    GetUserCheckinSummary()(OfficeUserCheckinSummaryable)
    GetUserPreferencePayload()([]byte)
    SetAssignments(value []OfficeClientConfigurationAssignmentable)()
    SetCheckinStatuses(value []OfficeClientCheckinStatusable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetPolicyPayload(value []byte)()
    SetPriority(value *int32)()
    SetUserCheckinSummary(value OfficeUserCheckinSummaryable)()
    SetUserPreferencePayload(value []byte)()
}
