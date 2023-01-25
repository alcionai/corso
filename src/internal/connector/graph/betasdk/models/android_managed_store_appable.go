package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidManagedStoreAppable 
type AndroidManagedStoreAppable interface {
    MobileAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppIdentifier()(*string)
    GetAppStoreUrl()(*string)
    GetAppTracks()([]AndroidManagedStoreAppTrackable)
    GetIsPrivate()(*bool)
    GetIsSystemApp()(*bool)
    GetPackageId()(*string)
    GetSupportsOemConfig()(*bool)
    GetTotalLicenseCount()(*int32)
    GetUsedLicenseCount()(*int32)
    SetAppIdentifier(value *string)()
    SetAppStoreUrl(value *string)()
    SetAppTracks(value []AndroidManagedStoreAppTrackable)()
    SetIsPrivate(value *bool)()
    SetIsSystemApp(value *bool)()
    SetPackageId(value *string)()
    SetSupportsOemConfig(value *bool)()
    SetTotalLicenseCount(value *int32)()
    SetUsedLicenseCount(value *int32)()
}
