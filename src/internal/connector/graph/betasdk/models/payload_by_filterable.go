package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PayloadByFilterable 
type PayloadByFilterable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignmentFilterType()(*DeviceAndAppManagementAssignmentFilterType)
    GetGroupId()(*string)
    GetOdataType()(*string)
    GetPayloadId()(*string)
    GetPayloadType()(*AssociatedAssignmentPayloadType)
    SetAssignmentFilterType(value *DeviceAndAppManagementAssignmentFilterType)()
    SetGroupId(value *string)()
    SetOdataType(value *string)()
    SetPayloadId(value *string)()
    SetPayloadType(value *AssociatedAssignmentPayloadType)()
}
