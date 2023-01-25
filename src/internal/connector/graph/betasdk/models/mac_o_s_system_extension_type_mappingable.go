package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSSystemExtensionTypeMappingable 
type MacOSSystemExtensionTypeMappingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedTypes()(*MacOSSystemExtensionType)
    GetOdataType()(*string)
    GetTeamIdentifier()(*string)
    SetAllowedTypes(value *MacOSSystemExtensionType)()
    SetOdataType(value *string)()
    SetTeamIdentifier(value *string)()
}
