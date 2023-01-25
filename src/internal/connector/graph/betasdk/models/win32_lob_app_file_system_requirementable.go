package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppFileSystemRequirementable 
type Win32LobAppFileSystemRequirementable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Win32LobAppRequirementable
    GetCheck32BitOn64System()(*bool)
    GetDetectionType()(*Win32LobAppFileSystemDetectionType)
    GetFileOrFolderName()(*string)
    GetPath()(*string)
    SetCheck32BitOn64System(value *bool)()
    SetDetectionType(value *Win32LobAppFileSystemDetectionType)()
    SetFileOrFolderName(value *string)()
    SetPath(value *string)()
}
