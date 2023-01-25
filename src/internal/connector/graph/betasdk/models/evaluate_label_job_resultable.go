package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EvaluateLabelJobResultable 
type EvaluateLabelJobResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetResponsiblePolicy()(ResponsiblePolicyable)
    GetResponsibleSensitiveTypes()([]ResponsibleSensitiveTypeable)
    GetSensitivityLabel()(MatchingLabelable)
    SetOdataType(value *string)()
    SetResponsiblePolicy(value ResponsiblePolicyable)()
    SetResponsibleSensitiveTypes(value []ResponsibleSensitiveTypeable)()
    SetSensitivityLabel(value MatchingLabelable)()
}
