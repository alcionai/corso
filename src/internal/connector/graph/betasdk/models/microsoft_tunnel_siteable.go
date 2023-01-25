package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftTunnelSiteable 
type MicrosoftTunnelSiteable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetInternalNetworkProbeUrl()(*string)
    GetMicrosoftTunnelConfiguration()(MicrosoftTunnelConfigurationable)
    GetMicrosoftTunnelServers()([]MicrosoftTunnelServerable)
    GetPublicAddress()(*string)
    GetRoleScopeTagIds()([]string)
    GetUpgradeAutomatically()(*bool)
    GetUpgradeAvailable()(*bool)
    GetUpgradeWindowEndTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    GetUpgradeWindowStartTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    GetUpgradeWindowUtcOffsetInMinutes()(*int32)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetInternalNetworkProbeUrl(value *string)()
    SetMicrosoftTunnelConfiguration(value MicrosoftTunnelConfigurationable)()
    SetMicrosoftTunnelServers(value []MicrosoftTunnelServerable)()
    SetPublicAddress(value *string)()
    SetRoleScopeTagIds(value []string)()
    SetUpgradeAutomatically(value *bool)()
    SetUpgradeAvailable(value *bool)()
    SetUpgradeWindowEndTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)()
    SetUpgradeWindowStartTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)()
    SetUpgradeWindowUtcOffsetInMinutes(value *int32)()
}
