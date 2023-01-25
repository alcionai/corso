package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerGlobalProxyDirectable 
type AndroidDeviceOwnerGlobalProxyDirectable interface {
    AndroidDeviceOwnerGlobalProxyable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExcludedHosts()([]string)
    GetHost()(*string)
    GetPort()(*int32)
    SetExcludedHosts(value []string)()
    SetHost(value *string)()
    SetPort(value *int32)()
}
