package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsPrivacyDataAccessControlItemable 
type WindowsPrivacyDataAccessControlItemable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessLevel()(*WindowsPrivacyDataAccessLevel)
    GetAppDisplayName()(*string)
    GetAppPackageFamilyName()(*string)
    GetDataCategory()(*WindowsPrivacyDataCategory)
    SetAccessLevel(value *WindowsPrivacyDataAccessLevel)()
    SetAppDisplayName(value *string)()
    SetAppPackageFamilyName(value *string)()
    SetDataCategory(value *WindowsPrivacyDataCategory)()
}
