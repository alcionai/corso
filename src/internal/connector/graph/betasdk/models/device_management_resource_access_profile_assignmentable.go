package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementResourceAccessProfileAssignmentable 
type DeviceManagementResourceAccessProfileAssignmentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIntent()(*DeviceManagementResourceAccessProfileIntent)
    GetSourceId()(*string)
    GetTarget()(DeviceAndAppManagementAssignmentTargetable)
    SetIntent(value *DeviceManagementResourceAccessProfileIntent)()
    SetSourceId(value *string)()
    SetTarget(value DeviceAndAppManagementAssignmentTargetable)()
}
