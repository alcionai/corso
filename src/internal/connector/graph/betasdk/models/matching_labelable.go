package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MatchingLabelable 
type MatchingLabelable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationMode()(*ApplicationMode)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetId()(*string)
    GetIsEndpointProtectionEnabled()(*bool)
    GetLabelActions()([]LabelActionBaseable)
    GetName()(*string)
    GetOdataType()(*string)
    GetPolicyTip()(*string)
    GetPriority()(*int32)
    GetToolTip()(*string)
    SetApplicationMode(value *ApplicationMode)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetId(value *string)()
    SetIsEndpointProtectionEnabled(value *bool)()
    SetLabelActions(value []LabelActionBaseable)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetPolicyTip(value *string)()
    SetPriority(value *int32)()
    SetToolTip(value *string)()
}
