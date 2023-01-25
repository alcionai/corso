package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnPremisesDirectorySynchronizationFeature 
type OnPremisesDirectorySynchronizationFeature struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Used to block cloud object takeover via source anchor hard match if enabled.
    blockCloudObjectTakeoverThroughHardMatchEnabled *bool
    // Use to block soft match for all objects if enabled for the  tenant. Customers are encouraged to enable this feature and keep it enabled until soft matching is required again for their tenancy. This flag should be enabled again after any soft matching has been completed and is no longer needed.
    blockSoftMatchEnabled *bool
    // When true, persists the values of Mobile and OtherMobile in on-premises AD during sync cycles instead of values of MobilePhone or AlternateMobilePhones in Azure AD.
    bypassDirSyncOverridesEnabled *bool
    // Used to indicate that cloud password policy applies to users whose passwords are synchronized from on-premises.
    cloudPasswordPolicyForPasswordSyncedUsersEnabled *bool
    // Used to enable concurrent user credentials update in OrgId.
    concurrentCredentialUpdateEnabled *bool
    // Used to enable concurrent user creation in OrgId.
    concurrentOrgIdProvisioningEnabled *bool
    // Used to indicate that device write-back is enabled.
    deviceWritebackEnabled *bool
    // Used to indicate that directory extensions are being synced from on-premises AD to Azure AD.
    directoryExtensionsEnabled *bool
    // Used to indicate that for a Microsoft Forefront Online Protection for Exchange (FOPE) migrated tenant, the conflicting proxy address should be migrated over.
    fopeConflictResolutionEnabled *bool
    // Used to enable object-level group writeback feature for additional group types.
    groupWriteBackEnabled *bool
    // The OdataType property
    odataType *string
    // Used to indicate on-premise password synchronization is enabled.
    passwordSyncEnabled *bool
    // Used to indicate that writeback of password resets from Azure AD to on-premises AD is enabled.
    passwordWritebackEnabled *bool
    // Used to indicate that we should quarantine objects with conflicting proxy address.
    quarantineUponProxyAddressesConflictEnabled *bool
    // Used to indicate that we should quarantine objects conflicting with duplicate userPrincipalName.
    quarantineUponUpnConflictEnabled *bool
    // Used to indicate that we should soft match objects based on userPrincipalName.
    softMatchOnUpnEnabled *bool
    // Used to indicate that we should synchronize userPrincipalName objects for managed users with licenses.
    synchronizeUpnForManagedUsersEnabled *bool
    // Used to indicate that Microsoft 365 Group write-back is enabled.
    unifiedGroupWritebackEnabled *bool
    // Used to indicate that feature to force password change for a user on logon is enabled while synchronizing on-premise credentials.
    userForcePasswordChangeOnLogonEnabled *bool
    // Used to indicate that user writeback is enabled.
    userWritebackEnabled *bool
}
// NewOnPremisesDirectorySynchronizationFeature instantiates a new onPremisesDirectorySynchronizationFeature and sets the default values.
func NewOnPremisesDirectorySynchronizationFeature()(*OnPremisesDirectorySynchronizationFeature) {
    m := &OnPremisesDirectorySynchronizationFeature{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOnPremisesDirectorySynchronizationFeatureFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnPremisesDirectorySynchronizationFeatureFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnPremisesDirectorySynchronizationFeature(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnPremisesDirectorySynchronizationFeature) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetBlockCloudObjectTakeoverThroughHardMatchEnabled gets the blockCloudObjectTakeoverThroughHardMatchEnabled property value. Used to block cloud object takeover via source anchor hard match if enabled.
func (m *OnPremisesDirectorySynchronizationFeature) GetBlockCloudObjectTakeoverThroughHardMatchEnabled()(*bool) {
    return m.blockCloudObjectTakeoverThroughHardMatchEnabled
}
// GetBlockSoftMatchEnabled gets the blockSoftMatchEnabled property value. Use to block soft match for all objects if enabled for the  tenant. Customers are encouraged to enable this feature and keep it enabled until soft matching is required again for their tenancy. This flag should be enabled again after any soft matching has been completed and is no longer needed.
func (m *OnPremisesDirectorySynchronizationFeature) GetBlockSoftMatchEnabled()(*bool) {
    return m.blockSoftMatchEnabled
}
// GetBypassDirSyncOverridesEnabled gets the bypassDirSyncOverridesEnabled property value. When true, persists the values of Mobile and OtherMobile in on-premises AD during sync cycles instead of values of MobilePhone or AlternateMobilePhones in Azure AD.
func (m *OnPremisesDirectorySynchronizationFeature) GetBypassDirSyncOverridesEnabled()(*bool) {
    return m.bypassDirSyncOverridesEnabled
}
// GetCloudPasswordPolicyForPasswordSyncedUsersEnabled gets the cloudPasswordPolicyForPasswordSyncedUsersEnabled property value. Used to indicate that cloud password policy applies to users whose passwords are synchronized from on-premises.
func (m *OnPremisesDirectorySynchronizationFeature) GetCloudPasswordPolicyForPasswordSyncedUsersEnabled()(*bool) {
    return m.cloudPasswordPolicyForPasswordSyncedUsersEnabled
}
// GetConcurrentCredentialUpdateEnabled gets the concurrentCredentialUpdateEnabled property value. Used to enable concurrent user credentials update in OrgId.
func (m *OnPremisesDirectorySynchronizationFeature) GetConcurrentCredentialUpdateEnabled()(*bool) {
    return m.concurrentCredentialUpdateEnabled
}
// GetConcurrentOrgIdProvisioningEnabled gets the concurrentOrgIdProvisioningEnabled property value. Used to enable concurrent user creation in OrgId.
func (m *OnPremisesDirectorySynchronizationFeature) GetConcurrentOrgIdProvisioningEnabled()(*bool) {
    return m.concurrentOrgIdProvisioningEnabled
}
// GetDeviceWritebackEnabled gets the deviceWritebackEnabled property value. Used to indicate that device write-back is enabled.
func (m *OnPremisesDirectorySynchronizationFeature) GetDeviceWritebackEnabled()(*bool) {
    return m.deviceWritebackEnabled
}
// GetDirectoryExtensionsEnabled gets the directoryExtensionsEnabled property value. Used to indicate that directory extensions are being synced from on-premises AD to Azure AD.
func (m *OnPremisesDirectorySynchronizationFeature) GetDirectoryExtensionsEnabled()(*bool) {
    return m.directoryExtensionsEnabled
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnPremisesDirectorySynchronizationFeature) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["blockCloudObjectTakeoverThroughHardMatchEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockCloudObjectTakeoverThroughHardMatchEnabled(val)
        }
        return nil
    }
    res["blockSoftMatchEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockSoftMatchEnabled(val)
        }
        return nil
    }
    res["bypassDirSyncOverridesEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBypassDirSyncOverridesEnabled(val)
        }
        return nil
    }
    res["cloudPasswordPolicyForPasswordSyncedUsersEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPasswordPolicyForPasswordSyncedUsersEnabled(val)
        }
        return nil
    }
    res["concurrentCredentialUpdateEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConcurrentCredentialUpdateEnabled(val)
        }
        return nil
    }
    res["concurrentOrgIdProvisioningEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConcurrentOrgIdProvisioningEnabled(val)
        }
        return nil
    }
    res["deviceWritebackEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceWritebackEnabled(val)
        }
        return nil
    }
    res["directoryExtensionsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectoryExtensionsEnabled(val)
        }
        return nil
    }
    res["fopeConflictResolutionEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFopeConflictResolutionEnabled(val)
        }
        return nil
    }
    res["groupWriteBackEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupWriteBackEnabled(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["passwordSyncEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordSyncEnabled(val)
        }
        return nil
    }
    res["passwordWritebackEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordWritebackEnabled(val)
        }
        return nil
    }
    res["quarantineUponProxyAddressesConflictEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuarantineUponProxyAddressesConflictEnabled(val)
        }
        return nil
    }
    res["quarantineUponUpnConflictEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuarantineUponUpnConflictEnabled(val)
        }
        return nil
    }
    res["softMatchOnUpnEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftMatchOnUpnEnabled(val)
        }
        return nil
    }
    res["synchronizeUpnForManagedUsersEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSynchronizeUpnForManagedUsersEnabled(val)
        }
        return nil
    }
    res["unifiedGroupWritebackEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnifiedGroupWritebackEnabled(val)
        }
        return nil
    }
    res["userForcePasswordChangeOnLogonEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserForcePasswordChangeOnLogonEnabled(val)
        }
        return nil
    }
    res["userWritebackEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserWritebackEnabled(val)
        }
        return nil
    }
    return res
}
// GetFopeConflictResolutionEnabled gets the fopeConflictResolutionEnabled property value. Used to indicate that for a Microsoft Forefront Online Protection for Exchange (FOPE) migrated tenant, the conflicting proxy address should be migrated over.
func (m *OnPremisesDirectorySynchronizationFeature) GetFopeConflictResolutionEnabled()(*bool) {
    return m.fopeConflictResolutionEnabled
}
// GetGroupWriteBackEnabled gets the groupWriteBackEnabled property value. Used to enable object-level group writeback feature for additional group types.
func (m *OnPremisesDirectorySynchronizationFeature) GetGroupWriteBackEnabled()(*bool) {
    return m.groupWriteBackEnabled
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OnPremisesDirectorySynchronizationFeature) GetOdataType()(*string) {
    return m.odataType
}
// GetPasswordSyncEnabled gets the passwordSyncEnabled property value. Used to indicate on-premise password synchronization is enabled.
func (m *OnPremisesDirectorySynchronizationFeature) GetPasswordSyncEnabled()(*bool) {
    return m.passwordSyncEnabled
}
// GetPasswordWritebackEnabled gets the passwordWritebackEnabled property value. Used to indicate that writeback of password resets from Azure AD to on-premises AD is enabled.
func (m *OnPremisesDirectorySynchronizationFeature) GetPasswordWritebackEnabled()(*bool) {
    return m.passwordWritebackEnabled
}
// GetQuarantineUponProxyAddressesConflictEnabled gets the quarantineUponProxyAddressesConflictEnabled property value. Used to indicate that we should quarantine objects with conflicting proxy address.
func (m *OnPremisesDirectorySynchronizationFeature) GetQuarantineUponProxyAddressesConflictEnabled()(*bool) {
    return m.quarantineUponProxyAddressesConflictEnabled
}
// GetQuarantineUponUpnConflictEnabled gets the quarantineUponUpnConflictEnabled property value. Used to indicate that we should quarantine objects conflicting with duplicate userPrincipalName.
func (m *OnPremisesDirectorySynchronizationFeature) GetQuarantineUponUpnConflictEnabled()(*bool) {
    return m.quarantineUponUpnConflictEnabled
}
// GetSoftMatchOnUpnEnabled gets the softMatchOnUpnEnabled property value. Used to indicate that we should soft match objects based on userPrincipalName.
func (m *OnPremisesDirectorySynchronizationFeature) GetSoftMatchOnUpnEnabled()(*bool) {
    return m.softMatchOnUpnEnabled
}
// GetSynchronizeUpnForManagedUsersEnabled gets the synchronizeUpnForManagedUsersEnabled property value. Used to indicate that we should synchronize userPrincipalName objects for managed users with licenses.
func (m *OnPremisesDirectorySynchronizationFeature) GetSynchronizeUpnForManagedUsersEnabled()(*bool) {
    return m.synchronizeUpnForManagedUsersEnabled
}
// GetUnifiedGroupWritebackEnabled gets the unifiedGroupWritebackEnabled property value. Used to indicate that Microsoft 365 Group write-back is enabled.
func (m *OnPremisesDirectorySynchronizationFeature) GetUnifiedGroupWritebackEnabled()(*bool) {
    return m.unifiedGroupWritebackEnabled
}
// GetUserForcePasswordChangeOnLogonEnabled gets the userForcePasswordChangeOnLogonEnabled property value. Used to indicate that feature to force password change for a user on logon is enabled while synchronizing on-premise credentials.
func (m *OnPremisesDirectorySynchronizationFeature) GetUserForcePasswordChangeOnLogonEnabled()(*bool) {
    return m.userForcePasswordChangeOnLogonEnabled
}
// GetUserWritebackEnabled gets the userWritebackEnabled property value. Used to indicate that user writeback is enabled.
func (m *OnPremisesDirectorySynchronizationFeature) GetUserWritebackEnabled()(*bool) {
    return m.userWritebackEnabled
}
// Serialize serializes information the current object
func (m *OnPremisesDirectorySynchronizationFeature) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("blockCloudObjectTakeoverThroughHardMatchEnabled", m.GetBlockCloudObjectTakeoverThroughHardMatchEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("blockSoftMatchEnabled", m.GetBlockSoftMatchEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("bypassDirSyncOverridesEnabled", m.GetBypassDirSyncOverridesEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("cloudPasswordPolicyForPasswordSyncedUsersEnabled", m.GetCloudPasswordPolicyForPasswordSyncedUsersEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("concurrentCredentialUpdateEnabled", m.GetConcurrentCredentialUpdateEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("concurrentOrgIdProvisioningEnabled", m.GetConcurrentOrgIdProvisioningEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("deviceWritebackEnabled", m.GetDeviceWritebackEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("directoryExtensionsEnabled", m.GetDirectoryExtensionsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("fopeConflictResolutionEnabled", m.GetFopeConflictResolutionEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("groupWriteBackEnabled", m.GetGroupWriteBackEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("passwordSyncEnabled", m.GetPasswordSyncEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("passwordWritebackEnabled", m.GetPasswordWritebackEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("quarantineUponProxyAddressesConflictEnabled", m.GetQuarantineUponProxyAddressesConflictEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("quarantineUponUpnConflictEnabled", m.GetQuarantineUponUpnConflictEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("softMatchOnUpnEnabled", m.GetSoftMatchOnUpnEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("synchronizeUpnForManagedUsersEnabled", m.GetSynchronizeUpnForManagedUsersEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("unifiedGroupWritebackEnabled", m.GetUnifiedGroupWritebackEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("userForcePasswordChangeOnLogonEnabled", m.GetUserForcePasswordChangeOnLogonEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("userWritebackEnabled", m.GetUserWritebackEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnPremisesDirectorySynchronizationFeature) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetBlockCloudObjectTakeoverThroughHardMatchEnabled sets the blockCloudObjectTakeoverThroughHardMatchEnabled property value. Used to block cloud object takeover via source anchor hard match if enabled.
func (m *OnPremisesDirectorySynchronizationFeature) SetBlockCloudObjectTakeoverThroughHardMatchEnabled(value *bool)() {
    m.blockCloudObjectTakeoverThroughHardMatchEnabled = value
}
// SetBlockSoftMatchEnabled sets the blockSoftMatchEnabled property value. Use to block soft match for all objects if enabled for the  tenant. Customers are encouraged to enable this feature and keep it enabled until soft matching is required again for their tenancy. This flag should be enabled again after any soft matching has been completed and is no longer needed.
func (m *OnPremisesDirectorySynchronizationFeature) SetBlockSoftMatchEnabled(value *bool)() {
    m.blockSoftMatchEnabled = value
}
// SetBypassDirSyncOverridesEnabled sets the bypassDirSyncOverridesEnabled property value. When true, persists the values of Mobile and OtherMobile in on-premises AD during sync cycles instead of values of MobilePhone or AlternateMobilePhones in Azure AD.
func (m *OnPremisesDirectorySynchronizationFeature) SetBypassDirSyncOverridesEnabled(value *bool)() {
    m.bypassDirSyncOverridesEnabled = value
}
// SetCloudPasswordPolicyForPasswordSyncedUsersEnabled sets the cloudPasswordPolicyForPasswordSyncedUsersEnabled property value. Used to indicate that cloud password policy applies to users whose passwords are synchronized from on-premises.
func (m *OnPremisesDirectorySynchronizationFeature) SetCloudPasswordPolicyForPasswordSyncedUsersEnabled(value *bool)() {
    m.cloudPasswordPolicyForPasswordSyncedUsersEnabled = value
}
// SetConcurrentCredentialUpdateEnabled sets the concurrentCredentialUpdateEnabled property value. Used to enable concurrent user credentials update in OrgId.
func (m *OnPremisesDirectorySynchronizationFeature) SetConcurrentCredentialUpdateEnabled(value *bool)() {
    m.concurrentCredentialUpdateEnabled = value
}
// SetConcurrentOrgIdProvisioningEnabled sets the concurrentOrgIdProvisioningEnabled property value. Used to enable concurrent user creation in OrgId.
func (m *OnPremisesDirectorySynchronizationFeature) SetConcurrentOrgIdProvisioningEnabled(value *bool)() {
    m.concurrentOrgIdProvisioningEnabled = value
}
// SetDeviceWritebackEnabled sets the deviceWritebackEnabled property value. Used to indicate that device write-back is enabled.
func (m *OnPremisesDirectorySynchronizationFeature) SetDeviceWritebackEnabled(value *bool)() {
    m.deviceWritebackEnabled = value
}
// SetDirectoryExtensionsEnabled sets the directoryExtensionsEnabled property value. Used to indicate that directory extensions are being synced from on-premises AD to Azure AD.
func (m *OnPremisesDirectorySynchronizationFeature) SetDirectoryExtensionsEnabled(value *bool)() {
    m.directoryExtensionsEnabled = value
}
// SetFopeConflictResolutionEnabled sets the fopeConflictResolutionEnabled property value. Used to indicate that for a Microsoft Forefront Online Protection for Exchange (FOPE) migrated tenant, the conflicting proxy address should be migrated over.
func (m *OnPremisesDirectorySynchronizationFeature) SetFopeConflictResolutionEnabled(value *bool)() {
    m.fopeConflictResolutionEnabled = value
}
// SetGroupWriteBackEnabled sets the groupWriteBackEnabled property value. Used to enable object-level group writeback feature for additional group types.
func (m *OnPremisesDirectorySynchronizationFeature) SetGroupWriteBackEnabled(value *bool)() {
    m.groupWriteBackEnabled = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OnPremisesDirectorySynchronizationFeature) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPasswordSyncEnabled sets the passwordSyncEnabled property value. Used to indicate on-premise password synchronization is enabled.
func (m *OnPremisesDirectorySynchronizationFeature) SetPasswordSyncEnabled(value *bool)() {
    m.passwordSyncEnabled = value
}
// SetPasswordWritebackEnabled sets the passwordWritebackEnabled property value. Used to indicate that writeback of password resets from Azure AD to on-premises AD is enabled.
func (m *OnPremisesDirectorySynchronizationFeature) SetPasswordWritebackEnabled(value *bool)() {
    m.passwordWritebackEnabled = value
}
// SetQuarantineUponProxyAddressesConflictEnabled sets the quarantineUponProxyAddressesConflictEnabled property value. Used to indicate that we should quarantine objects with conflicting proxy address.
func (m *OnPremisesDirectorySynchronizationFeature) SetQuarantineUponProxyAddressesConflictEnabled(value *bool)() {
    m.quarantineUponProxyAddressesConflictEnabled = value
}
// SetQuarantineUponUpnConflictEnabled sets the quarantineUponUpnConflictEnabled property value. Used to indicate that we should quarantine objects conflicting with duplicate userPrincipalName.
func (m *OnPremisesDirectorySynchronizationFeature) SetQuarantineUponUpnConflictEnabled(value *bool)() {
    m.quarantineUponUpnConflictEnabled = value
}
// SetSoftMatchOnUpnEnabled sets the softMatchOnUpnEnabled property value. Used to indicate that we should soft match objects based on userPrincipalName.
func (m *OnPremisesDirectorySynchronizationFeature) SetSoftMatchOnUpnEnabled(value *bool)() {
    m.softMatchOnUpnEnabled = value
}
// SetSynchronizeUpnForManagedUsersEnabled sets the synchronizeUpnForManagedUsersEnabled property value. Used to indicate that we should synchronize userPrincipalName objects for managed users with licenses.
func (m *OnPremisesDirectorySynchronizationFeature) SetSynchronizeUpnForManagedUsersEnabled(value *bool)() {
    m.synchronizeUpnForManagedUsersEnabled = value
}
// SetUnifiedGroupWritebackEnabled sets the unifiedGroupWritebackEnabled property value. Used to indicate that Microsoft 365 Group write-back is enabled.
func (m *OnPremisesDirectorySynchronizationFeature) SetUnifiedGroupWritebackEnabled(value *bool)() {
    m.unifiedGroupWritebackEnabled = value
}
// SetUserForcePasswordChangeOnLogonEnabled sets the userForcePasswordChangeOnLogonEnabled property value. Used to indicate that feature to force password change for a user on logon is enabled while synchronizing on-premise credentials.
func (m *OnPremisesDirectorySynchronizationFeature) SetUserForcePasswordChangeOnLogonEnabled(value *bool)() {
    m.userForcePasswordChangeOnLogonEnabled = value
}
// SetUserWritebackEnabled sets the userWritebackEnabled property value. Used to indicate that user writeback is enabled.
func (m *OnPremisesDirectorySynchronizationFeature) SetUserWritebackEnabled(value *bool)() {
    m.userWritebackEnabled = value
}
