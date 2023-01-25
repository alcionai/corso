package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsPackageInformationable 
type WindowsPackageInformationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicableArchitecture()(*WindowsArchitecture)
    GetDisplayName()(*string)
    GetIdentityName()(*string)
    GetIdentityPublisher()(*string)
    GetIdentityResourceIdentifier()(*string)
    GetIdentityVersion()(*string)
    GetMinimumSupportedOperatingSystem()(WindowsMinimumOperatingSystemable)
    GetOdataType()(*string)
    SetApplicableArchitecture(value *WindowsArchitecture)()
    SetDisplayName(value *string)()
    SetIdentityName(value *string)()
    SetIdentityPublisher(value *string)()
    SetIdentityResourceIdentifier(value *string)()
    SetIdentityVersion(value *string)()
    SetMinimumSupportedOperatingSystem(value WindowsMinimumOperatingSystemable)()
    SetOdataType(value *string)()
}
