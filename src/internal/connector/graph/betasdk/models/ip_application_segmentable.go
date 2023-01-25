package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IpApplicationSegmentable 
type IpApplicationSegmentable interface {
    ApplicationSegmentable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDestinationHost()(*string)
    GetPort()(*int32)
    SetDestinationHost(value *string)()
    SetPort(value *int32)()
}
