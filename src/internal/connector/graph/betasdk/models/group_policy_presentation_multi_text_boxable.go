package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationMultiTextBoxable 
type GroupPolicyPresentationMultiTextBoxable interface {
    GroupPolicyUploadedPresentationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMaxLength()(*int64)
    GetMaxStrings()(*int64)
    GetRequired()(*bool)
    SetMaxLength(value *int64)()
    SetMaxStrings(value *int64)()
    SetRequired(value *bool)()
}
