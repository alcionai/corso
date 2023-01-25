package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyPresentationComboBoxable 
type GroupPolicyPresentationComboBoxable interface {
    GroupPolicyUploadedPresentationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultValue()(*string)
    GetMaxLength()(*int64)
    GetRequired()(*bool)
    GetSuggestions()([]string)
    SetDefaultValue(value *string)()
    SetMaxLength(value *int64)()
    SetRequired(value *bool)()
    SetSuggestions(value []string)()
}
