package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OemWarrantyInformationOnboardingable 
type OemWarrantyInformationOnboardingable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailable()(*bool)
    GetEnabled()(*bool)
    GetOemName()(*string)
    SetAvailable(value *bool)()
    SetEnabled(value *bool)()
    SetOemName(value *string)()
}
