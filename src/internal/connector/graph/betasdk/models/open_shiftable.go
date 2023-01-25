package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OpenShiftable 
type OpenShiftable interface {
    ChangeTrackedEntityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDraftOpenShift()(OpenShiftItemable)
    GetIsStagedForDeletion()(*bool)
    GetSchedulingGroupId()(*string)
    GetSharedOpenShift()(OpenShiftItemable)
    SetDraftOpenShift(value OpenShiftItemable)()
    SetIsStagedForDeletion(value *bool)()
    SetSchedulingGroupId(value *string)()
    SetSharedOpenShift(value OpenShiftItemable)()
}
