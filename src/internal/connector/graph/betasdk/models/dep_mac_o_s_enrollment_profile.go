package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DepMacOSEnrollmentProfile 
type DepMacOSEnrollmentProfile struct {
    DepEnrollmentBaseProfile
    // Indicates if Accessibility screen is disabled
    accessibilityScreenDisabled *bool
    // Indicates if UnlockWithWatch screen is disabled
    autoUnlockWithWatchDisabled *bool
    // Indicates if iCloud Documents and Desktop screen is disabled
    chooseYourLockScreenDisabled *bool
    // Indicates whether Setup Assistant will auto populate the primary account information
    dontAutoPopulatePrimaryAccountInfo *bool
    // Indicates whether the user will enable blockediting
    enableRestrictEditing *bool
    // Indicates if file vault is disabled
    fileVaultDisabled *bool
    // Indicates if iCloud Analytics screen is disabled
    iCloudDiagnosticsDisabled *bool
    // Indicates if iCloud Documents and Desktop screen is disabled
    iCloudStorageDisabled *bool
    // Indicates whether the profile is a local account
    isLocalPrimaryAccount *bool
    // Indicates whether the profile is a primary user
    isPrimaryUser *bool
    // Indicates whether the primary account information will be locked
    lockPrimaryAccountInfo *bool
    // Indicates whether or not this is the short name of the local account to manage
    managedLocalUserShortName *bool
    // Indicates if Passcode setup pane is disabled
    passCodeDisabled *bool
    // Indicates whether the user will prefill their account info
    prefillAccountInfo *bool
    // Indicates what the full name for the primary account is
    primaryAccountFullName *string
    // Indicates what the account name for the primary account is
    primaryAccountUserName *string
    // Indicates who the primary user of the profile is
    primaryUser *string
    // Indicates who the primary user of the profile is
    primaryUserFullName *string
    // Indicates if registration is disabled
    registrationDisabled *bool
    // Indicates if the device is network-tethered to run the command
    requestRequiresNetworkTether *bool
    // Indicates whether Setup Assistant will set the account as a regular user
    setPrimarySetupAccountAsRegularUser *bool
    // Indicates whether Setup Assistant will skip the user interface for primary account setup
    skipPrimarySetupAccountCreation *bool
    // Indicates if zoom setup pane is disabled
    zoomDisabled *bool
}
// NewDepMacOSEnrollmentProfile instantiates a new DepMacOSEnrollmentProfile and sets the default values.
func NewDepMacOSEnrollmentProfile()(*DepMacOSEnrollmentProfile) {
    m := &DepMacOSEnrollmentProfile{
        DepEnrollmentBaseProfile: *NewDepEnrollmentBaseProfile(),
    }
    odataTypeValue := "#microsoft.graph.depMacOSEnrollmentProfile";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDepMacOSEnrollmentProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDepMacOSEnrollmentProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDepMacOSEnrollmentProfile(), nil
}
// GetAccessibilityScreenDisabled gets the accessibilityScreenDisabled property value. Indicates if Accessibility screen is disabled
func (m *DepMacOSEnrollmentProfile) GetAccessibilityScreenDisabled()(*bool) {
    return m.accessibilityScreenDisabled
}
// GetAutoUnlockWithWatchDisabled gets the autoUnlockWithWatchDisabled property value. Indicates if UnlockWithWatch screen is disabled
func (m *DepMacOSEnrollmentProfile) GetAutoUnlockWithWatchDisabled()(*bool) {
    return m.autoUnlockWithWatchDisabled
}
// GetChooseYourLockScreenDisabled gets the chooseYourLockScreenDisabled property value. Indicates if iCloud Documents and Desktop screen is disabled
func (m *DepMacOSEnrollmentProfile) GetChooseYourLockScreenDisabled()(*bool) {
    return m.chooseYourLockScreenDisabled
}
// GetDontAutoPopulatePrimaryAccountInfo gets the dontAutoPopulatePrimaryAccountInfo property value. Indicates whether Setup Assistant will auto populate the primary account information
func (m *DepMacOSEnrollmentProfile) GetDontAutoPopulatePrimaryAccountInfo()(*bool) {
    return m.dontAutoPopulatePrimaryAccountInfo
}
// GetEnableRestrictEditing gets the enableRestrictEditing property value. Indicates whether the user will enable blockediting
func (m *DepMacOSEnrollmentProfile) GetEnableRestrictEditing()(*bool) {
    return m.enableRestrictEditing
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DepMacOSEnrollmentProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DepEnrollmentBaseProfile.GetFieldDeserializers()
    res["accessibilityScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessibilityScreenDisabled(val)
        }
        return nil
    }
    res["autoUnlockWithWatchDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoUnlockWithWatchDisabled(val)
        }
        return nil
    }
    res["chooseYourLockScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChooseYourLockScreenDisabled(val)
        }
        return nil
    }
    res["dontAutoPopulatePrimaryAccountInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDontAutoPopulatePrimaryAccountInfo(val)
        }
        return nil
    }
    res["enableRestrictEditing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableRestrictEditing(val)
        }
        return nil
    }
    res["fileVaultDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileVaultDisabled(val)
        }
        return nil
    }
    res["iCloudDiagnosticsDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudDiagnosticsDisabled(val)
        }
        return nil
    }
    res["iCloudStorageDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudStorageDisabled(val)
        }
        return nil
    }
    res["isLocalPrimaryAccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsLocalPrimaryAccount(val)
        }
        return nil
    }
    res["isPrimaryUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPrimaryUser(val)
        }
        return nil
    }
    res["lockPrimaryAccountInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLockPrimaryAccountInfo(val)
        }
        return nil
    }
    res["managedLocalUserShortName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedLocalUserShortName(val)
        }
        return nil
    }
    res["passCodeDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassCodeDisabled(val)
        }
        return nil
    }
    res["prefillAccountInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrefillAccountInfo(val)
        }
        return nil
    }
    res["primaryAccountFullName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrimaryAccountFullName(val)
        }
        return nil
    }
    res["primaryAccountUserName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrimaryAccountUserName(val)
        }
        return nil
    }
    res["primaryUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrimaryUser(val)
        }
        return nil
    }
    res["primaryUserFullName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrimaryUserFullName(val)
        }
        return nil
    }
    res["registrationDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistrationDisabled(val)
        }
        return nil
    }
    res["requestRequiresNetworkTether"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestRequiresNetworkTether(val)
        }
        return nil
    }
    res["setPrimarySetupAccountAsRegularUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSetPrimarySetupAccountAsRegularUser(val)
        }
        return nil
    }
    res["skipPrimarySetupAccountCreation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkipPrimarySetupAccountCreation(val)
        }
        return nil
    }
    res["zoomDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetZoomDisabled(val)
        }
        return nil
    }
    return res
}
// GetFileVaultDisabled gets the fileVaultDisabled property value. Indicates if file vault is disabled
func (m *DepMacOSEnrollmentProfile) GetFileVaultDisabled()(*bool) {
    return m.fileVaultDisabled
}
// GetICloudDiagnosticsDisabled gets the iCloudDiagnosticsDisabled property value. Indicates if iCloud Analytics screen is disabled
func (m *DepMacOSEnrollmentProfile) GetICloudDiagnosticsDisabled()(*bool) {
    return m.iCloudDiagnosticsDisabled
}
// GetICloudStorageDisabled gets the iCloudStorageDisabled property value. Indicates if iCloud Documents and Desktop screen is disabled
func (m *DepMacOSEnrollmentProfile) GetICloudStorageDisabled()(*bool) {
    return m.iCloudStorageDisabled
}
// GetIsLocalPrimaryAccount gets the isLocalPrimaryAccount property value. Indicates whether the profile is a local account
func (m *DepMacOSEnrollmentProfile) GetIsLocalPrimaryAccount()(*bool) {
    return m.isLocalPrimaryAccount
}
// GetIsPrimaryUser gets the isPrimaryUser property value. Indicates whether the profile is a primary user
func (m *DepMacOSEnrollmentProfile) GetIsPrimaryUser()(*bool) {
    return m.isPrimaryUser
}
// GetLockPrimaryAccountInfo gets the lockPrimaryAccountInfo property value. Indicates whether the primary account information will be locked
func (m *DepMacOSEnrollmentProfile) GetLockPrimaryAccountInfo()(*bool) {
    return m.lockPrimaryAccountInfo
}
// GetManagedLocalUserShortName gets the managedLocalUserShortName property value. Indicates whether or not this is the short name of the local account to manage
func (m *DepMacOSEnrollmentProfile) GetManagedLocalUserShortName()(*bool) {
    return m.managedLocalUserShortName
}
// GetPassCodeDisabled gets the passCodeDisabled property value. Indicates if Passcode setup pane is disabled
func (m *DepMacOSEnrollmentProfile) GetPassCodeDisabled()(*bool) {
    return m.passCodeDisabled
}
// GetPrefillAccountInfo gets the prefillAccountInfo property value. Indicates whether the user will prefill their account info
func (m *DepMacOSEnrollmentProfile) GetPrefillAccountInfo()(*bool) {
    return m.prefillAccountInfo
}
// GetPrimaryAccountFullName gets the primaryAccountFullName property value. Indicates what the full name for the primary account is
func (m *DepMacOSEnrollmentProfile) GetPrimaryAccountFullName()(*string) {
    return m.primaryAccountFullName
}
// GetPrimaryAccountUserName gets the primaryAccountUserName property value. Indicates what the account name for the primary account is
func (m *DepMacOSEnrollmentProfile) GetPrimaryAccountUserName()(*string) {
    return m.primaryAccountUserName
}
// GetPrimaryUser gets the primaryUser property value. Indicates who the primary user of the profile is
func (m *DepMacOSEnrollmentProfile) GetPrimaryUser()(*string) {
    return m.primaryUser
}
// GetPrimaryUserFullName gets the primaryUserFullName property value. Indicates who the primary user of the profile is
func (m *DepMacOSEnrollmentProfile) GetPrimaryUserFullName()(*string) {
    return m.primaryUserFullName
}
// GetRegistrationDisabled gets the registrationDisabled property value. Indicates if registration is disabled
func (m *DepMacOSEnrollmentProfile) GetRegistrationDisabled()(*bool) {
    return m.registrationDisabled
}
// GetRequestRequiresNetworkTether gets the requestRequiresNetworkTether property value. Indicates if the device is network-tethered to run the command
func (m *DepMacOSEnrollmentProfile) GetRequestRequiresNetworkTether()(*bool) {
    return m.requestRequiresNetworkTether
}
// GetSetPrimarySetupAccountAsRegularUser gets the setPrimarySetupAccountAsRegularUser property value. Indicates whether Setup Assistant will set the account as a regular user
func (m *DepMacOSEnrollmentProfile) GetSetPrimarySetupAccountAsRegularUser()(*bool) {
    return m.setPrimarySetupAccountAsRegularUser
}
// GetSkipPrimarySetupAccountCreation gets the skipPrimarySetupAccountCreation property value. Indicates whether Setup Assistant will skip the user interface for primary account setup
func (m *DepMacOSEnrollmentProfile) GetSkipPrimarySetupAccountCreation()(*bool) {
    return m.skipPrimarySetupAccountCreation
}
// GetZoomDisabled gets the zoomDisabled property value. Indicates if zoom setup pane is disabled
func (m *DepMacOSEnrollmentProfile) GetZoomDisabled()(*bool) {
    return m.zoomDisabled
}
// Serialize serializes information the current object
func (m *DepMacOSEnrollmentProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DepEnrollmentBaseProfile.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("accessibilityScreenDisabled", m.GetAccessibilityScreenDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("autoUnlockWithWatchDisabled", m.GetAutoUnlockWithWatchDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("chooseYourLockScreenDisabled", m.GetChooseYourLockScreenDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("dontAutoPopulatePrimaryAccountInfo", m.GetDontAutoPopulatePrimaryAccountInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableRestrictEditing", m.GetEnableRestrictEditing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("fileVaultDisabled", m.GetFileVaultDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudDiagnosticsDisabled", m.GetICloudDiagnosticsDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudStorageDisabled", m.GetICloudStorageDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isLocalPrimaryAccount", m.GetIsLocalPrimaryAccount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isPrimaryUser", m.GetIsPrimaryUser())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("lockPrimaryAccountInfo", m.GetLockPrimaryAccountInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("managedLocalUserShortName", m.GetManagedLocalUserShortName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passCodeDisabled", m.GetPassCodeDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("prefillAccountInfo", m.GetPrefillAccountInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("primaryAccountFullName", m.GetPrimaryAccountFullName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("primaryAccountUserName", m.GetPrimaryAccountUserName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("primaryUser", m.GetPrimaryUser())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("primaryUserFullName", m.GetPrimaryUserFullName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("registrationDisabled", m.GetRegistrationDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("requestRequiresNetworkTether", m.GetRequestRequiresNetworkTether())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("setPrimarySetupAccountAsRegularUser", m.GetSetPrimarySetupAccountAsRegularUser())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("skipPrimarySetupAccountCreation", m.GetSkipPrimarySetupAccountCreation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("zoomDisabled", m.GetZoomDisabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessibilityScreenDisabled sets the accessibilityScreenDisabled property value. Indicates if Accessibility screen is disabled
func (m *DepMacOSEnrollmentProfile) SetAccessibilityScreenDisabled(value *bool)() {
    m.accessibilityScreenDisabled = value
}
// SetAutoUnlockWithWatchDisabled sets the autoUnlockWithWatchDisabled property value. Indicates if UnlockWithWatch screen is disabled
func (m *DepMacOSEnrollmentProfile) SetAutoUnlockWithWatchDisabled(value *bool)() {
    m.autoUnlockWithWatchDisabled = value
}
// SetChooseYourLockScreenDisabled sets the chooseYourLockScreenDisabled property value. Indicates if iCloud Documents and Desktop screen is disabled
func (m *DepMacOSEnrollmentProfile) SetChooseYourLockScreenDisabled(value *bool)() {
    m.chooseYourLockScreenDisabled = value
}
// SetDontAutoPopulatePrimaryAccountInfo sets the dontAutoPopulatePrimaryAccountInfo property value. Indicates whether Setup Assistant will auto populate the primary account information
func (m *DepMacOSEnrollmentProfile) SetDontAutoPopulatePrimaryAccountInfo(value *bool)() {
    m.dontAutoPopulatePrimaryAccountInfo = value
}
// SetEnableRestrictEditing sets the enableRestrictEditing property value. Indicates whether the user will enable blockediting
func (m *DepMacOSEnrollmentProfile) SetEnableRestrictEditing(value *bool)() {
    m.enableRestrictEditing = value
}
// SetFileVaultDisabled sets the fileVaultDisabled property value. Indicates if file vault is disabled
func (m *DepMacOSEnrollmentProfile) SetFileVaultDisabled(value *bool)() {
    m.fileVaultDisabled = value
}
// SetICloudDiagnosticsDisabled sets the iCloudDiagnosticsDisabled property value. Indicates if iCloud Analytics screen is disabled
func (m *DepMacOSEnrollmentProfile) SetICloudDiagnosticsDisabled(value *bool)() {
    m.iCloudDiagnosticsDisabled = value
}
// SetICloudStorageDisabled sets the iCloudStorageDisabled property value. Indicates if iCloud Documents and Desktop screen is disabled
func (m *DepMacOSEnrollmentProfile) SetICloudStorageDisabled(value *bool)() {
    m.iCloudStorageDisabled = value
}
// SetIsLocalPrimaryAccount sets the isLocalPrimaryAccount property value. Indicates whether the profile is a local account
func (m *DepMacOSEnrollmentProfile) SetIsLocalPrimaryAccount(value *bool)() {
    m.isLocalPrimaryAccount = value
}
// SetIsPrimaryUser sets the isPrimaryUser property value. Indicates whether the profile is a primary user
func (m *DepMacOSEnrollmentProfile) SetIsPrimaryUser(value *bool)() {
    m.isPrimaryUser = value
}
// SetLockPrimaryAccountInfo sets the lockPrimaryAccountInfo property value. Indicates whether the primary account information will be locked
func (m *DepMacOSEnrollmentProfile) SetLockPrimaryAccountInfo(value *bool)() {
    m.lockPrimaryAccountInfo = value
}
// SetManagedLocalUserShortName sets the managedLocalUserShortName property value. Indicates whether or not this is the short name of the local account to manage
func (m *DepMacOSEnrollmentProfile) SetManagedLocalUserShortName(value *bool)() {
    m.managedLocalUserShortName = value
}
// SetPassCodeDisabled sets the passCodeDisabled property value. Indicates if Passcode setup pane is disabled
func (m *DepMacOSEnrollmentProfile) SetPassCodeDisabled(value *bool)() {
    m.passCodeDisabled = value
}
// SetPrefillAccountInfo sets the prefillAccountInfo property value. Indicates whether the user will prefill their account info
func (m *DepMacOSEnrollmentProfile) SetPrefillAccountInfo(value *bool)() {
    m.prefillAccountInfo = value
}
// SetPrimaryAccountFullName sets the primaryAccountFullName property value. Indicates what the full name for the primary account is
func (m *DepMacOSEnrollmentProfile) SetPrimaryAccountFullName(value *string)() {
    m.primaryAccountFullName = value
}
// SetPrimaryAccountUserName sets the primaryAccountUserName property value. Indicates what the account name for the primary account is
func (m *DepMacOSEnrollmentProfile) SetPrimaryAccountUserName(value *string)() {
    m.primaryAccountUserName = value
}
// SetPrimaryUser sets the primaryUser property value. Indicates who the primary user of the profile is
func (m *DepMacOSEnrollmentProfile) SetPrimaryUser(value *string)() {
    m.primaryUser = value
}
// SetPrimaryUserFullName sets the primaryUserFullName property value. Indicates who the primary user of the profile is
func (m *DepMacOSEnrollmentProfile) SetPrimaryUserFullName(value *string)() {
    m.primaryUserFullName = value
}
// SetRegistrationDisabled sets the registrationDisabled property value. Indicates if registration is disabled
func (m *DepMacOSEnrollmentProfile) SetRegistrationDisabled(value *bool)() {
    m.registrationDisabled = value
}
// SetRequestRequiresNetworkTether sets the requestRequiresNetworkTether property value. Indicates if the device is network-tethered to run the command
func (m *DepMacOSEnrollmentProfile) SetRequestRequiresNetworkTether(value *bool)() {
    m.requestRequiresNetworkTether = value
}
// SetSetPrimarySetupAccountAsRegularUser sets the setPrimarySetupAccountAsRegularUser property value. Indicates whether Setup Assistant will set the account as a regular user
func (m *DepMacOSEnrollmentProfile) SetSetPrimarySetupAccountAsRegularUser(value *bool)() {
    m.setPrimarySetupAccountAsRegularUser = value
}
// SetSkipPrimarySetupAccountCreation sets the skipPrimarySetupAccountCreation property value. Indicates whether Setup Assistant will skip the user interface for primary account setup
func (m *DepMacOSEnrollmentProfile) SetSkipPrimarySetupAccountCreation(value *bool)() {
    m.skipPrimarySetupAccountCreation = value
}
// SetZoomDisabled sets the zoomDisabled property value. Indicates if zoom setup pane is disabled
func (m *DepMacOSEnrollmentProfile) SetZoomDisabled(value *bool)() {
    m.zoomDisabled = value
}
