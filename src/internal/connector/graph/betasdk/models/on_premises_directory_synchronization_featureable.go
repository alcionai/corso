package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnPremisesDirectorySynchronizationFeatureable 
type OnPremisesDirectorySynchronizationFeatureable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBlockCloudObjectTakeoverThroughHardMatchEnabled()(*bool)
    GetBlockSoftMatchEnabled()(*bool)
    GetBypassDirSyncOverridesEnabled()(*bool)
    GetCloudPasswordPolicyForPasswordSyncedUsersEnabled()(*bool)
    GetConcurrentCredentialUpdateEnabled()(*bool)
    GetConcurrentOrgIdProvisioningEnabled()(*bool)
    GetDeviceWritebackEnabled()(*bool)
    GetDirectoryExtensionsEnabled()(*bool)
    GetFopeConflictResolutionEnabled()(*bool)
    GetGroupWriteBackEnabled()(*bool)
    GetOdataType()(*string)
    GetPasswordSyncEnabled()(*bool)
    GetPasswordWritebackEnabled()(*bool)
    GetQuarantineUponProxyAddressesConflictEnabled()(*bool)
    GetQuarantineUponUpnConflictEnabled()(*bool)
    GetSoftMatchOnUpnEnabled()(*bool)
    GetSynchronizeUpnForManagedUsersEnabled()(*bool)
    GetUnifiedGroupWritebackEnabled()(*bool)
    GetUserForcePasswordChangeOnLogonEnabled()(*bool)
    GetUserWritebackEnabled()(*bool)
    SetBlockCloudObjectTakeoverThroughHardMatchEnabled(value *bool)()
    SetBlockSoftMatchEnabled(value *bool)()
    SetBypassDirSyncOverridesEnabled(value *bool)()
    SetCloudPasswordPolicyForPasswordSyncedUsersEnabled(value *bool)()
    SetConcurrentCredentialUpdateEnabled(value *bool)()
    SetConcurrentOrgIdProvisioningEnabled(value *bool)()
    SetDeviceWritebackEnabled(value *bool)()
    SetDirectoryExtensionsEnabled(value *bool)()
    SetFopeConflictResolutionEnabled(value *bool)()
    SetGroupWriteBackEnabled(value *bool)()
    SetOdataType(value *string)()
    SetPasswordSyncEnabled(value *bool)()
    SetPasswordWritebackEnabled(value *bool)()
    SetQuarantineUponProxyAddressesConflictEnabled(value *bool)()
    SetQuarantineUponUpnConflictEnabled(value *bool)()
    SetSoftMatchOnUpnEnabled(value *bool)()
    SetSynchronizeUpnForManagedUsersEnabled(value *bool)()
    SetUnifiedGroupWritebackEnabled(value *bool)()
    SetUserForcePasswordChangeOnLogonEnabled(value *bool)()
    SetUserWritebackEnabled(value *bool)()
}
