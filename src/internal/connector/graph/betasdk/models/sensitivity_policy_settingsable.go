package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SensitivityPolicySettingsable 
type SensitivityPolicySettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicableTo()(*SensitivityLabelTarget)
    GetDowngradeSensitivityRequiresJustification()(*bool)
    GetHelpWebUrl()(*string)
    GetIsMandatory()(*bool)
    SetApplicableTo(value *SensitivityLabelTarget)()
    SetDowngradeSensitivityRequiresJustification(value *bool)()
    SetHelpWebUrl(value *string)()
    SetIsMandatory(value *bool)()
}
