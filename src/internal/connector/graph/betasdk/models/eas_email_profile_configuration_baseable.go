package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EasEmailProfileConfigurationBaseable 
type EasEmailProfileConfigurationBaseable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomDomainName()(*string)
    GetUserDomainNameSource()(*DomainNameSource)
    GetUsernameAADSource()(*UsernameSource)
    GetUsernameSource()(*UserEmailSource)
    SetCustomDomainName(value *string)()
    SetUserDomainNameSource(value *DomainNameSource)()
    SetUsernameAADSource(value *UsernameSource)()
    SetUsernameSource(value *UserEmailSource)()
}
