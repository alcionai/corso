package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppFileSystemDetectionable 
type Win32LobAppFileSystemDetectionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Win32LobAppDetectionable
    GetCheck32BitOn64System()(*bool)
    GetDetectionType()(*Win32LobAppFileSystemDetectionType)
    GetDetectionValue()(*string)
    GetFileOrFolderName()(*string)
    GetOperator()(*Win32LobAppDetectionOperator)
    GetPath()(*string)
    SetCheck32BitOn64System(value *bool)()
    SetDetectionType(value *Win32LobAppFileSystemDetectionType)()
    SetDetectionValue(value *string)()
    SetFileOrFolderName(value *string)()
    SetOperator(value *Win32LobAppDetectionOperator)()
    SetPath(value *string)()
}
