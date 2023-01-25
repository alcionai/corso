package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SensitivityLabelable 
type SensitivityLabelable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicableTo()(*SensitivityLabelTarget)
    GetApplicationMode()(*ApplicationMode)
    GetAssignedPolicies()([]LabelPolicyable)
    GetAutoLabeling()(AutoLabelingable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIsDefault()(*bool)
    GetIsEndpointProtectionEnabled()(*bool)
    GetLabelActions()([]LabelActionBaseable)
    GetName()(*string)
    GetPriority()(*int32)
    GetSublabels()([]SensitivityLabelable)
    GetToolTip()(*string)
    SetApplicableTo(value *SensitivityLabelTarget)()
    SetApplicationMode(value *ApplicationMode)()
    SetAssignedPolicies(value []LabelPolicyable)()
    SetAutoLabeling(value AutoLabelingable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIsDefault(value *bool)()
    SetIsEndpointProtectionEnabled(value *bool)()
    SetLabelActions(value []LabelActionBaseable)()
    SetName(value *string)()
    SetPriority(value *int32)()
    SetSublabels(value []SensitivityLabelable)()
    SetToolTip(value *string)()
}
