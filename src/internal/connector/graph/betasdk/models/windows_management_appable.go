package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsManagementAppable 
type WindowsManagementAppable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailableVersion()(*string)
    GetHealthStates()([]WindowsManagementAppHealthStateable)
    GetManagedInstaller()(*ManagedInstallerStatus)
    GetManagedInstallerConfiguredDateTime()(*string)
    SetAvailableVersion(value *string)()
    SetHealthStates(value []WindowsManagementAppHealthStateable)()
    SetManagedInstaller(value *ManagedInstallerStatus)()
    SetManagedInstallerConfiguredDateTime(value *string)()
}
