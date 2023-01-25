package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ZebraFotaArtifactable 
type ZebraFotaArtifactable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBoardSupportPackageVersion()(*string)
    GetDescription()(*string)
    GetDeviceModel()(*string)
    GetOsVersion()(*string)
    GetPatchVersion()(*string)
    GetReleaseNotesUrl()(*string)
    SetBoardSupportPackageVersion(value *string)()
    SetDescription(value *string)()
    SetDeviceModel(value *string)()
    SetOsVersion(value *string)()
    SetPatchVersion(value *string)()
    SetReleaseNotesUrl(value *string)()
}
