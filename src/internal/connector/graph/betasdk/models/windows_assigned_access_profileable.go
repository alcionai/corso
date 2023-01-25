package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsAssignedAccessProfileable 
type WindowsAssignedAccessProfileable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppUserModelIds()([]string)
    GetDesktopAppPaths()([]string)
    GetProfileName()(*string)
    GetShowTaskBar()(*bool)
    GetStartMenuLayoutXml()([]byte)
    GetUserAccounts()([]string)
    SetAppUserModelIds(value []string)()
    SetDesktopAppPaths(value []string)()
    SetProfileName(value *string)()
    SetShowTaskBar(value *bool)()
    SetStartMenuLayoutXml(value []byte)()
    SetUserAccounts(value []string)()
}
