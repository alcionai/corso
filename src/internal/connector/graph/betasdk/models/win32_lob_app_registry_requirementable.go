package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppRegistryRequirementable 
type Win32LobAppRegistryRequirementable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Win32LobAppRequirementable
    GetCheck32BitOn64System()(*bool)
    GetDetectionType()(*Win32LobAppRegistryDetectionType)
    GetKeyPath()(*string)
    GetValueName()(*string)
    SetCheck32BitOn64System(value *bool)()
    SetDetectionType(value *Win32LobAppRegistryDetectionType)()
    SetKeyPath(value *string)()
    SetValueName(value *string)()
}
