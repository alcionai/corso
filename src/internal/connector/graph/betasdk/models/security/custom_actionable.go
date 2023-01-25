package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomActionable 
type CustomActionable interface {
    InformationProtectionActionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetProperties()([]KeyValuePairable)
    SetName(value *string)()
    SetProperties(value []KeyValuePairable)()
}
