package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnsupportedDeviceConfigurationable 
type UnsupportedDeviceConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDetails()([]UnsupportedDeviceConfigurationDetailable)
    GetOriginalEntityTypeName()(*string)
    SetDetails(value []UnsupportedDeviceConfigurationDetailable)()
    SetOriginalEntityTypeName(value *string)()
}
