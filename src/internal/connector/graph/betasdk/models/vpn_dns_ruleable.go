package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VpnDnsRuleable 
type VpnDnsRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAutoTrigger()(*bool)
    GetName()(*string)
    GetOdataType()(*string)
    GetPersistent()(*bool)
    GetProxyServerUri()(*string)
    GetServers()([]string)
    SetAutoTrigger(value *bool)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetPersistent(value *bool)()
    SetProxyServerUri(value *string)()
    SetServers(value []string)()
}
