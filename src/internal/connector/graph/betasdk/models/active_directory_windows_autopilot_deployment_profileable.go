package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ActiveDirectoryWindowsAutopilotDeploymentProfileable 
type ActiveDirectoryWindowsAutopilotDeploymentProfileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WindowsAutopilotDeploymentProfileable
    GetDomainJoinConfiguration()(WindowsDomainJoinConfigurationable)
    GetHybridAzureADJoinSkipConnectivityCheck()(*bool)
    SetDomainJoinConfiguration(value WindowsDomainJoinConfigurationable)()
    SetHybridAzureADJoinSkipConnectivityCheck(value *bool)()
}
