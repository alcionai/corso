package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsPhone81AppXable 
type WindowsPhone81AppXable interface {
    MobileLobAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicableArchitectures()(*WindowsArchitecture)
    GetIdentityName()(*string)
    GetIdentityPublisherHash()(*string)
    GetIdentityResourceIdentifier()(*string)
    GetIdentityVersion()(*string)
    GetMinimumSupportedOperatingSystem()(WindowsMinimumOperatingSystemable)
    GetPhoneProductIdentifier()(*string)
    GetPhonePublisherId()(*string)
    SetApplicableArchitectures(value *WindowsArchitecture)()
    SetIdentityName(value *string)()
    SetIdentityPublisherHash(value *string)()
    SetIdentityResourceIdentifier(value *string)()
    SetIdentityVersion(value *string)()
    SetMinimumSupportedOperatingSystem(value WindowsMinimumOperatingSystemable)()
    SetPhoneProductIdentifier(value *string)()
    SetPhonePublisherId(value *string)()
}
