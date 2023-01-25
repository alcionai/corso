package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PayloadCoachmarkable 
type PayloadCoachmarkable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCoachmarkLocation()(CoachmarkLocationable)
    GetDescription()(*string)
    GetIndicator()(*string)
    GetIsValid()(*bool)
    GetLanguage()(*string)
    GetOdataType()(*string)
    GetOrder()(*string)
    SetCoachmarkLocation(value CoachmarkLocationable)()
    SetDescription(value *string)()
    SetIndicator(value *string)()
    SetIsValid(value *bool)()
    SetLanguage(value *string)()
    SetOdataType(value *string)()
    SetOrder(value *string)()
}
