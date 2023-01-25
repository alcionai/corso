package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WinGetAppable 
type WinGetAppable interface {
    MobileAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInstallExperience()(WinGetAppInstallExperienceable)
    GetManifestHash()(*string)
    GetPackageIdentifier()(*string)
    SetInstallExperience(value WinGetAppInstallExperienceable)()
    SetManifestHash(value *string)()
    SetPackageIdentifier(value *string)()
}
