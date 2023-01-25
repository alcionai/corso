package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationDecimalTextBoxable 
type GroupPolicyPresentationDecimalTextBoxable interface {
    GroupPolicyUploadedPresentationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultValue()(*int64)
    GetMaxValue()(*int64)
    GetMinValue()(*int64)
    GetRequired()(*bool)
    GetSpin()(*bool)
    GetSpinStep()(*int64)
    SetDefaultValue(value *int64)()
    SetMaxValue(value *int64)()
    SetMinValue(value *int64)()
    SetRequired(value *bool)()
    SetSpin(value *bool)()
    SetSpinStep(value *int64)()
}
