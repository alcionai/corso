package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSDmgAppable 
type MacOSDmgAppable interface {
    MobileLobAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIgnoreVersionDetection()(*bool)
    GetIncludedApps()([]MacOSIncludedAppable)
    GetMinimumSupportedOperatingSystem()(MacOSMinimumOperatingSystemable)
    GetPrimaryBundleId()(*string)
    GetPrimaryBundleVersion()(*string)
    SetIgnoreVersionDetection(value *bool)()
    SetIncludedApps(value []MacOSIncludedAppable)()
    SetMinimumSupportedOperatingSystem(value MacOSMinimumOperatingSystemable)()
    SetPrimaryBundleId(value *string)()
    SetPrimaryBundleVersion(value *string)()
}
