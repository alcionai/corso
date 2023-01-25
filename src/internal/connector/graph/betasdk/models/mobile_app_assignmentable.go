package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppAssignmentable 
type MobileAppAssignmentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIntent()(*InstallIntent)
    GetSettings()(MobileAppAssignmentSettingsable)
    GetSource()(*DeviceAndAppManagementAssignmentSource)
    GetSourceId()(*string)
    GetTarget()(DeviceAndAppManagementAssignmentTargetable)
    SetIntent(value *InstallIntent)()
    SetSettings(value MobileAppAssignmentSettingsable)()
    SetSource(value *DeviceAndAppManagementAssignmentSource)()
    SetSourceId(value *string)()
    SetTarget(value DeviceAndAppManagementAssignmentTargetable)()
}
