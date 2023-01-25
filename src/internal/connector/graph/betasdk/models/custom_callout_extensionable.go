package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomCalloutExtensionable 
type CustomCalloutExtensionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationConfiguration()(CustomExtensionAuthenticationConfigurationable)
    GetClientConfiguration()(CustomExtensionClientConfigurationable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetEndpointConfiguration()(CustomExtensionEndpointConfigurationable)
    SetAuthenticationConfiguration(value CustomExtensionAuthenticationConfigurationable)()
    SetClientConfiguration(value CustomExtensionClientConfigurationable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetEndpointConfiguration(value CustomExtensionEndpointConfigurationable)()
}
