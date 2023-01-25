package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityConfigurationTaskable 
type SecurityConfigurationTaskable interface {
    DeviceAppManagementTaskable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicablePlatform()(*EndpointSecurityConfigurationApplicablePlatform)
    GetEndpointSecurityPolicy()(*EndpointSecurityConfigurationType)
    GetEndpointSecurityPolicyProfile()(*EndpointSecurityConfigurationProfileType)
    GetInsights()(*string)
    GetIntendedSettings()([]KeyValuePairable)
    GetManagedDeviceCount()(*int32)
    GetManagedDevices()([]VulnerableManagedDeviceable)
    SetApplicablePlatform(value *EndpointSecurityConfigurationApplicablePlatform)()
    SetEndpointSecurityPolicy(value *EndpointSecurityConfigurationType)()
    SetEndpointSecurityPolicyProfile(value *EndpointSecurityConfigurationProfileType)()
    SetInsights(value *string)()
    SetIntendedSettings(value []KeyValuePairable)()
    SetManagedDeviceCount(value *int32)()
    SetManagedDevices(value []VulnerableManagedDeviceable)()
}
