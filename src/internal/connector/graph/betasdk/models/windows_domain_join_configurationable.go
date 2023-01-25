package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsDomainJoinConfigurationable 
type WindowsDomainJoinConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveDirectoryDomainName()(*string)
    GetComputerNameStaticPrefix()(*string)
    GetComputerNameSuffixRandomCharCount()(*int32)
    GetNetworkAccessConfigurations()([]DeviceConfigurationable)
    GetOrganizationalUnit()(*string)
    SetActiveDirectoryDomainName(value *string)()
    SetComputerNameStaticPrefix(value *string)()
    SetComputerNameSuffixRandomCharCount(value *int32)()
    SetNetworkAccessConfigurations(value []DeviceConfigurationable)()
    SetOrganizationalUnit(value *string)()
}
