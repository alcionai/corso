package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcOrganizationSettingsable 
type CloudPcOrganizationSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnableMEMAutoEnroll()(*bool)
    GetEnableSingleSignOn()(*bool)
    GetOsVersion()(*CloudPcOperatingSystem)
    GetUserAccountType()(*CloudPcUserAccountType)
    GetWindowsSettings()(CloudPcWindowsSettingsable)
    SetEnableMEMAutoEnroll(value *bool)()
    SetEnableSingleSignOn(value *bool)()
    SetOsVersion(value *CloudPcOperatingSystem)()
    SetUserAccountType(value *CloudPcUserAccountType)()
    SetWindowsSettings(value CloudPcWindowsSettingsable)()
}
