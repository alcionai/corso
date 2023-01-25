package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppTroubleshootingAppTargetHistoryable 
type MobileAppTroubleshootingAppTargetHistoryable interface {
    MobileAppTroubleshootingHistoryItemable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetErrorCode()(*string)
    GetRunState()(*RunState)
    GetSecurityGroupId()(*string)
    SetErrorCode(value *string)()
    SetRunState(value *RunState)()
    SetSecurityGroupId(value *string)()
}
