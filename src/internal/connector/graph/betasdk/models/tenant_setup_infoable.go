package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TenantSetupInfoable 
type TenantSetupInfoable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultRolesSettings()(PrivilegedRoleSettingsable)
    GetFirstTimeSetup()(*bool)
    GetRelevantRolesSettings()([]string)
    GetSetupStatus()(*SetupStatus)
    GetSkipSetup()(*bool)
    GetUserRolesActions()(*string)
    SetDefaultRolesSettings(value PrivilegedRoleSettingsable)()
    SetFirstTimeSetup(value *bool)()
    SetRelevantRolesSettings(value []string)()
    SetSetupStatus(value *SetupStatus)()
    SetSkipSetup(value *bool)()
    SetUserRolesActions(value *string)()
}
