package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EntitlementManagementSettingsable 
type EntitlementManagementSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDaysUntilExternalUserDeletedAfterBlocked()(*int32)
    GetExternalUserLifecycleAction()(*string)
    SetDaysUntilExternalUserDeletedAfterBlocked(value *int32)()
    SetExternalUserLifecycleAction(value *string)()
}
