package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows81WifiImportConfigurationable 
type Windows81WifiImportConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPayload()([]byte)
    GetPayloadFileName()(*string)
    GetProfileName()(*string)
    SetPayload(value []byte)()
    SetPayloadFileName(value *string)()
    SetProfileName(value *string)()
}
