package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DepIOSEnrollmentProfile 
type DepIOSEnrollmentProfile struct {
    DepEnrollmentBaseProfile
    // Indicates if Apperance screen is disabled
    appearanceScreenDisabled *bool
    // Indicates if the device will need to wait for configured confirmation
    awaitDeviceConfiguredConfirmation *bool
    // Carrier URL for activating device eSIM.
    carrierActivationUrl *string
    // If set, indicates which Vpp token should be used to deploy the Company Portal w/ device licensing. 'enableAuthenticationViaCompanyPortal' must be set in order for this property to be set.
    companyPortalVppTokenId *string
    // Indicates if Device To Device Migration is disabled
    deviceToDeviceMigrationDisabled *bool
    // This indicates whether the device is to be enrolled in a mode which enables multi user scenarios. Only applicable in shared iPads.
    enableSharedIPad *bool
    // Tells the device to enable single app mode and apply app-lock during enrollment. Default is false. 'enableAuthenticationViaCompanyPortal' and 'companyPortalVppTokenId' must be set for this property to be set.
    enableSingleAppEnrollmentMode *bool
    // Indicates if Express Language screen is disabled
    expressLanguageScreenDisabled *bool
    // Indicates if temporary sessions is enabled
    forceTemporarySession *bool
    // Indicates if home button sensitivity screen is disabled
    homeButtonScreenDisabled *bool
    // Indicates if iMessage and FaceTime screen is disabled
    iMessageAndFaceTimeScreenDisabled *bool
    // The iTunesPairingMode property
    iTunesPairingMode *ITunesPairingMode
    // Management certificates for Apple Configurator
    managementCertificates []ManagementCertificateWithThumbprintable
    // Indicates if onboarding setup screen is disabled
    onBoardingScreenDisabled *bool
    // Indicates if Passcode setup pane is disabled
    passCodeDisabled *bool
    // Indicates timeout before locked screen requires the user to enter the device passocde to unlock it
    passcodeLockGracePeriodInSeconds *int32
    // Indicates if Preferred language screen is disabled
    preferredLanguageScreenDisabled *bool
    // Indicates if Weclome screen is disabled
    restoreCompletedScreenDisabled *bool
    // Indicates if Restore from Android is disabled
    restoreFromAndroidDisabled *bool
    // This specifies the maximum number of users that can use a shared iPad. Only applicable in shared iPad mode.
    sharedIPadMaximumUserCount *int32
    // Indicates if the SIMSetup screen is disabled
    simSetupScreenDisabled *bool
    // Indicates if the mandatory sofware update screen is disabled
    softwareUpdateScreenDisabled *bool
    // Indicates timeout of temporary session
    temporarySessionTimeoutInSeconds *int32
    // Indicates if Weclome screen is disabled
    updateCompleteScreenDisabled *bool
    // Indicates that this apple device is designated to support 'shared device mode' scenarios. This is distinct from the 'shared iPad' scenario. See https://learn.microsoft.com/en-us/mem/intune/enrollment/device-enrollment-shared-ios
    userlessSharedAadModeEnabled *bool
    // Indicates timeout of temporary session
    userSessionTimeoutInSeconds *int32
    // Indicates if the watch migration screen is disabled
    watchMigrationScreenDisabled *bool
    // Indicates if Weclome screen is disabled
    welcomeScreenDisabled *bool
    // Indicates if zoom setup pane is disabled
    zoomDisabled *bool
}
// NewDepIOSEnrollmentProfile instantiates a new DepIOSEnrollmentProfile and sets the default values.
func NewDepIOSEnrollmentProfile()(*DepIOSEnrollmentProfile) {
    m := &DepIOSEnrollmentProfile{
        DepEnrollmentBaseProfile: *NewDepEnrollmentBaseProfile(),
    }
    odataTypeValue := "#microsoft.graph.depIOSEnrollmentProfile";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDepIOSEnrollmentProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDepIOSEnrollmentProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDepIOSEnrollmentProfile(), nil
}
// GetAppearanceScreenDisabled gets the appearanceScreenDisabled property value. Indicates if Apperance screen is disabled
func (m *DepIOSEnrollmentProfile) GetAppearanceScreenDisabled()(*bool) {
    return m.appearanceScreenDisabled
}
// GetAwaitDeviceConfiguredConfirmation gets the awaitDeviceConfiguredConfirmation property value. Indicates if the device will need to wait for configured confirmation
func (m *DepIOSEnrollmentProfile) GetAwaitDeviceConfiguredConfirmation()(*bool) {
    return m.awaitDeviceConfiguredConfirmation
}
// GetCarrierActivationUrl gets the carrierActivationUrl property value. Carrier URL for activating device eSIM.
func (m *DepIOSEnrollmentProfile) GetCarrierActivationUrl()(*string) {
    return m.carrierActivationUrl
}
// GetCompanyPortalVppTokenId gets the companyPortalVppTokenId property value. If set, indicates which Vpp token should be used to deploy the Company Portal w/ device licensing. 'enableAuthenticationViaCompanyPortal' must be set in order for this property to be set.
func (m *DepIOSEnrollmentProfile) GetCompanyPortalVppTokenId()(*string) {
    return m.companyPortalVppTokenId
}
// GetDeviceToDeviceMigrationDisabled gets the deviceToDeviceMigrationDisabled property value. Indicates if Device To Device Migration is disabled
func (m *DepIOSEnrollmentProfile) GetDeviceToDeviceMigrationDisabled()(*bool) {
    return m.deviceToDeviceMigrationDisabled
}
// GetEnableSharedIPad gets the enableSharedIPad property value. This indicates whether the device is to be enrolled in a mode which enables multi user scenarios. Only applicable in shared iPads.
func (m *DepIOSEnrollmentProfile) GetEnableSharedIPad()(*bool) {
    return m.enableSharedIPad
}
// GetEnableSingleAppEnrollmentMode gets the enableSingleAppEnrollmentMode property value. Tells the device to enable single app mode and apply app-lock during enrollment. Default is false. 'enableAuthenticationViaCompanyPortal' and 'companyPortalVppTokenId' must be set for this property to be set.
func (m *DepIOSEnrollmentProfile) GetEnableSingleAppEnrollmentMode()(*bool) {
    return m.enableSingleAppEnrollmentMode
}
// GetExpressLanguageScreenDisabled gets the expressLanguageScreenDisabled property value. Indicates if Express Language screen is disabled
func (m *DepIOSEnrollmentProfile) GetExpressLanguageScreenDisabled()(*bool) {
    return m.expressLanguageScreenDisabled
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DepIOSEnrollmentProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DepEnrollmentBaseProfile.GetFieldDeserializers()
    res["appearanceScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppearanceScreenDisabled(val)
        }
        return nil
    }
    res["awaitDeviceConfiguredConfirmation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAwaitDeviceConfiguredConfirmation(val)
        }
        return nil
    }
    res["carrierActivationUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCarrierActivationUrl(val)
        }
        return nil
    }
    res["companyPortalVppTokenId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompanyPortalVppTokenId(val)
        }
        return nil
    }
    res["deviceToDeviceMigrationDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceToDeviceMigrationDisabled(val)
        }
        return nil
    }
    res["enableSharedIPad"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableSharedIPad(val)
        }
        return nil
    }
    res["enableSingleAppEnrollmentMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableSingleAppEnrollmentMode(val)
        }
        return nil
    }
    res["expressLanguageScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpressLanguageScreenDisabled(val)
        }
        return nil
    }
    res["forceTemporarySession"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForceTemporarySession(val)
        }
        return nil
    }
    res["homeButtonScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomeButtonScreenDisabled(val)
        }
        return nil
    }
    res["iMessageAndFaceTimeScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIMessageAndFaceTimeScreenDisabled(val)
        }
        return nil
    }
    res["iTunesPairingMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseITunesPairingMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetITunesPairingMode(val.(*ITunesPairingMode))
        }
        return nil
    }
    res["managementCertificates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementCertificateWithThumbprintFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementCertificateWithThumbprintable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementCertificateWithThumbprintable)
            }
            m.SetManagementCertificates(res)
        }
        return nil
    }
    res["onBoardingScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnBoardingScreenDisabled(val)
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
    res["passcodeLockGracePeriodInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeLockGracePeriodInSeconds(val)
        }
        return nil
    }
    res["preferredLanguageScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferredLanguageScreenDisabled(val)
        }
        return nil
    }
    res["restoreCompletedScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestoreCompletedScreenDisabled(val)
        }
        return nil
    }
    res["restoreFromAndroidDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestoreFromAndroidDisabled(val)
        }
        return nil
    }
    res["sharedIPadMaximumUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedIPadMaximumUserCount(val)
        }
        return nil
    }
    res["simSetupScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSimSetupScreenDisabled(val)
        }
        return nil
    }
    res["softwareUpdateScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftwareUpdateScreenDisabled(val)
        }
        return nil
    }
    res["temporarySessionTimeoutInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemporarySessionTimeoutInSeconds(val)
        }
        return nil
    }
    res["updateCompleteScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateCompleteScreenDisabled(val)
        }
        return nil
    }
    res["userlessSharedAadModeEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserlessSharedAadModeEnabled(val)
        }
        return nil
    }
    res["userSessionTimeoutInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserSessionTimeoutInSeconds(val)
        }
        return nil
    }
    res["watchMigrationScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWatchMigrationScreenDisabled(val)
        }
        return nil
    }
    res["welcomeScreenDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWelcomeScreenDisabled(val)
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
// GetForceTemporarySession gets the forceTemporarySession property value. Indicates if temporary sessions is enabled
func (m *DepIOSEnrollmentProfile) GetForceTemporarySession()(*bool) {
    return m.forceTemporarySession
}
// GetHomeButtonScreenDisabled gets the homeButtonScreenDisabled property value. Indicates if home button sensitivity screen is disabled
func (m *DepIOSEnrollmentProfile) GetHomeButtonScreenDisabled()(*bool) {
    return m.homeButtonScreenDisabled
}
// GetIMessageAndFaceTimeScreenDisabled gets the iMessageAndFaceTimeScreenDisabled property value. Indicates if iMessage and FaceTime screen is disabled
func (m *DepIOSEnrollmentProfile) GetIMessageAndFaceTimeScreenDisabled()(*bool) {
    return m.iMessageAndFaceTimeScreenDisabled
}
// GetITunesPairingMode gets the iTunesPairingMode property value. The iTunesPairingMode property
func (m *DepIOSEnrollmentProfile) GetITunesPairingMode()(*ITunesPairingMode) {
    return m.iTunesPairingMode
}
// GetManagementCertificates gets the managementCertificates property value. Management certificates for Apple Configurator
func (m *DepIOSEnrollmentProfile) GetManagementCertificates()([]ManagementCertificateWithThumbprintable) {
    return m.managementCertificates
}
// GetOnBoardingScreenDisabled gets the onBoardingScreenDisabled property value. Indicates if onboarding setup screen is disabled
func (m *DepIOSEnrollmentProfile) GetOnBoardingScreenDisabled()(*bool) {
    return m.onBoardingScreenDisabled
}
// GetPassCodeDisabled gets the passCodeDisabled property value. Indicates if Passcode setup pane is disabled
func (m *DepIOSEnrollmentProfile) GetPassCodeDisabled()(*bool) {
    return m.passCodeDisabled
}
// GetPasscodeLockGracePeriodInSeconds gets the passcodeLockGracePeriodInSeconds property value. Indicates timeout before locked screen requires the user to enter the device passocde to unlock it
func (m *DepIOSEnrollmentProfile) GetPasscodeLockGracePeriodInSeconds()(*int32) {
    return m.passcodeLockGracePeriodInSeconds
}
// GetPreferredLanguageScreenDisabled gets the preferredLanguageScreenDisabled property value. Indicates if Preferred language screen is disabled
func (m *DepIOSEnrollmentProfile) GetPreferredLanguageScreenDisabled()(*bool) {
    return m.preferredLanguageScreenDisabled
}
// GetRestoreCompletedScreenDisabled gets the restoreCompletedScreenDisabled property value. Indicates if Weclome screen is disabled
func (m *DepIOSEnrollmentProfile) GetRestoreCompletedScreenDisabled()(*bool) {
    return m.restoreCompletedScreenDisabled
}
// GetRestoreFromAndroidDisabled gets the restoreFromAndroidDisabled property value. Indicates if Restore from Android is disabled
func (m *DepIOSEnrollmentProfile) GetRestoreFromAndroidDisabled()(*bool) {
    return m.restoreFromAndroidDisabled
}
// GetSharedIPadMaximumUserCount gets the sharedIPadMaximumUserCount property value. This specifies the maximum number of users that can use a shared iPad. Only applicable in shared iPad mode.
func (m *DepIOSEnrollmentProfile) GetSharedIPadMaximumUserCount()(*int32) {
    return m.sharedIPadMaximumUserCount
}
// GetSimSetupScreenDisabled gets the simSetupScreenDisabled property value. Indicates if the SIMSetup screen is disabled
func (m *DepIOSEnrollmentProfile) GetSimSetupScreenDisabled()(*bool) {
    return m.simSetupScreenDisabled
}
// GetSoftwareUpdateScreenDisabled gets the softwareUpdateScreenDisabled property value. Indicates if the mandatory sofware update screen is disabled
func (m *DepIOSEnrollmentProfile) GetSoftwareUpdateScreenDisabled()(*bool) {
    return m.softwareUpdateScreenDisabled
}
// GetTemporarySessionTimeoutInSeconds gets the temporarySessionTimeoutInSeconds property value. Indicates timeout of temporary session
func (m *DepIOSEnrollmentProfile) GetTemporarySessionTimeoutInSeconds()(*int32) {
    return m.temporarySessionTimeoutInSeconds
}
// GetUpdateCompleteScreenDisabled gets the updateCompleteScreenDisabled property value. Indicates if Weclome screen is disabled
func (m *DepIOSEnrollmentProfile) GetUpdateCompleteScreenDisabled()(*bool) {
    return m.updateCompleteScreenDisabled
}
// GetUserlessSharedAadModeEnabled gets the userlessSharedAadModeEnabled property value. Indicates that this apple device is designated to support 'shared device mode' scenarios. This is distinct from the 'shared iPad' scenario. See https://learn.microsoft.com/en-us/mem/intune/enrollment/device-enrollment-shared-ios
func (m *DepIOSEnrollmentProfile) GetUserlessSharedAadModeEnabled()(*bool) {
    return m.userlessSharedAadModeEnabled
}
// GetUserSessionTimeoutInSeconds gets the userSessionTimeoutInSeconds property value. Indicates timeout of temporary session
func (m *DepIOSEnrollmentProfile) GetUserSessionTimeoutInSeconds()(*int32) {
    return m.userSessionTimeoutInSeconds
}
// GetWatchMigrationScreenDisabled gets the watchMigrationScreenDisabled property value. Indicates if the watch migration screen is disabled
func (m *DepIOSEnrollmentProfile) GetWatchMigrationScreenDisabled()(*bool) {
    return m.watchMigrationScreenDisabled
}
// GetWelcomeScreenDisabled gets the welcomeScreenDisabled property value. Indicates if Weclome screen is disabled
func (m *DepIOSEnrollmentProfile) GetWelcomeScreenDisabled()(*bool) {
    return m.welcomeScreenDisabled
}
// GetZoomDisabled gets the zoomDisabled property value. Indicates if zoom setup pane is disabled
func (m *DepIOSEnrollmentProfile) GetZoomDisabled()(*bool) {
    return m.zoomDisabled
}
// Serialize serializes information the current object
func (m *DepIOSEnrollmentProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DepEnrollmentBaseProfile.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("appearanceScreenDisabled", m.GetAppearanceScreenDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("awaitDeviceConfiguredConfirmation", m.GetAwaitDeviceConfiguredConfirmation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("carrierActivationUrl", m.GetCarrierActivationUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("companyPortalVppTokenId", m.GetCompanyPortalVppTokenId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deviceToDeviceMigrationDisabled", m.GetDeviceToDeviceMigrationDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableSharedIPad", m.GetEnableSharedIPad())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableSingleAppEnrollmentMode", m.GetEnableSingleAppEnrollmentMode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("expressLanguageScreenDisabled", m.GetExpressLanguageScreenDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("forceTemporarySession", m.GetForceTemporarySession())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("homeButtonScreenDisabled", m.GetHomeButtonScreenDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iMessageAndFaceTimeScreenDisabled", m.GetIMessageAndFaceTimeScreenDisabled())
        if err != nil {
            return err
        }
    }
    if m.GetITunesPairingMode() != nil {
        cast := (*m.GetITunesPairingMode()).String()
        err = writer.WriteStringValue("iTunesPairingMode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementCertificates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagementCertificates()))
        for i, v := range m.GetManagementCertificates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managementCertificates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("onBoardingScreenDisabled", m.GetOnBoardingScreenDisabled())
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
        err = writer.WriteInt32Value("passcodeLockGracePeriodInSeconds", m.GetPasscodeLockGracePeriodInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("preferredLanguageScreenDisabled", m.GetPreferredLanguageScreenDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("restoreCompletedScreenDisabled", m.GetRestoreCompletedScreenDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("restoreFromAndroidDisabled", m.GetRestoreFromAndroidDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("sharedIPadMaximumUserCount", m.GetSharedIPadMaximumUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("simSetupScreenDisabled", m.GetSimSetupScreenDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("softwareUpdateScreenDisabled", m.GetSoftwareUpdateScreenDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("temporarySessionTimeoutInSeconds", m.GetTemporarySessionTimeoutInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("updateCompleteScreenDisabled", m.GetUpdateCompleteScreenDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("userlessSharedAadModeEnabled", m.GetUserlessSharedAadModeEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("userSessionTimeoutInSeconds", m.GetUserSessionTimeoutInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("watchMigrationScreenDisabled", m.GetWatchMigrationScreenDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("welcomeScreenDisabled", m.GetWelcomeScreenDisabled())
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
// SetAppearanceScreenDisabled sets the appearanceScreenDisabled property value. Indicates if Apperance screen is disabled
func (m *DepIOSEnrollmentProfile) SetAppearanceScreenDisabled(value *bool)() {
    m.appearanceScreenDisabled = value
}
// SetAwaitDeviceConfiguredConfirmation sets the awaitDeviceConfiguredConfirmation property value. Indicates if the device will need to wait for configured confirmation
func (m *DepIOSEnrollmentProfile) SetAwaitDeviceConfiguredConfirmation(value *bool)() {
    m.awaitDeviceConfiguredConfirmation = value
}
// SetCarrierActivationUrl sets the carrierActivationUrl property value. Carrier URL for activating device eSIM.
func (m *DepIOSEnrollmentProfile) SetCarrierActivationUrl(value *string)() {
    m.carrierActivationUrl = value
}
// SetCompanyPortalVppTokenId sets the companyPortalVppTokenId property value. If set, indicates which Vpp token should be used to deploy the Company Portal w/ device licensing. 'enableAuthenticationViaCompanyPortal' must be set in order for this property to be set.
func (m *DepIOSEnrollmentProfile) SetCompanyPortalVppTokenId(value *string)() {
    m.companyPortalVppTokenId = value
}
// SetDeviceToDeviceMigrationDisabled sets the deviceToDeviceMigrationDisabled property value. Indicates if Device To Device Migration is disabled
func (m *DepIOSEnrollmentProfile) SetDeviceToDeviceMigrationDisabled(value *bool)() {
    m.deviceToDeviceMigrationDisabled = value
}
// SetEnableSharedIPad sets the enableSharedIPad property value. This indicates whether the device is to be enrolled in a mode which enables multi user scenarios. Only applicable in shared iPads.
func (m *DepIOSEnrollmentProfile) SetEnableSharedIPad(value *bool)() {
    m.enableSharedIPad = value
}
// SetEnableSingleAppEnrollmentMode sets the enableSingleAppEnrollmentMode property value. Tells the device to enable single app mode and apply app-lock during enrollment. Default is false. 'enableAuthenticationViaCompanyPortal' and 'companyPortalVppTokenId' must be set for this property to be set.
func (m *DepIOSEnrollmentProfile) SetEnableSingleAppEnrollmentMode(value *bool)() {
    m.enableSingleAppEnrollmentMode = value
}
// SetExpressLanguageScreenDisabled sets the expressLanguageScreenDisabled property value. Indicates if Express Language screen is disabled
func (m *DepIOSEnrollmentProfile) SetExpressLanguageScreenDisabled(value *bool)() {
    m.expressLanguageScreenDisabled = value
}
// SetForceTemporarySession sets the forceTemporarySession property value. Indicates if temporary sessions is enabled
func (m *DepIOSEnrollmentProfile) SetForceTemporarySession(value *bool)() {
    m.forceTemporarySession = value
}
// SetHomeButtonScreenDisabled sets the homeButtonScreenDisabled property value. Indicates if home button sensitivity screen is disabled
func (m *DepIOSEnrollmentProfile) SetHomeButtonScreenDisabled(value *bool)() {
    m.homeButtonScreenDisabled = value
}
// SetIMessageAndFaceTimeScreenDisabled sets the iMessageAndFaceTimeScreenDisabled property value. Indicates if iMessage and FaceTime screen is disabled
func (m *DepIOSEnrollmentProfile) SetIMessageAndFaceTimeScreenDisabled(value *bool)() {
    m.iMessageAndFaceTimeScreenDisabled = value
}
// SetITunesPairingMode sets the iTunesPairingMode property value. The iTunesPairingMode property
func (m *DepIOSEnrollmentProfile) SetITunesPairingMode(value *ITunesPairingMode)() {
    m.iTunesPairingMode = value
}
// SetManagementCertificates sets the managementCertificates property value. Management certificates for Apple Configurator
func (m *DepIOSEnrollmentProfile) SetManagementCertificates(value []ManagementCertificateWithThumbprintable)() {
    m.managementCertificates = value
}
// SetOnBoardingScreenDisabled sets the onBoardingScreenDisabled property value. Indicates if onboarding setup screen is disabled
func (m *DepIOSEnrollmentProfile) SetOnBoardingScreenDisabled(value *bool)() {
    m.onBoardingScreenDisabled = value
}
// SetPassCodeDisabled sets the passCodeDisabled property value. Indicates if Passcode setup pane is disabled
func (m *DepIOSEnrollmentProfile) SetPassCodeDisabled(value *bool)() {
    m.passCodeDisabled = value
}
// SetPasscodeLockGracePeriodInSeconds sets the passcodeLockGracePeriodInSeconds property value. Indicates timeout before locked screen requires the user to enter the device passocde to unlock it
func (m *DepIOSEnrollmentProfile) SetPasscodeLockGracePeriodInSeconds(value *int32)() {
    m.passcodeLockGracePeriodInSeconds = value
}
// SetPreferredLanguageScreenDisabled sets the preferredLanguageScreenDisabled property value. Indicates if Preferred language screen is disabled
func (m *DepIOSEnrollmentProfile) SetPreferredLanguageScreenDisabled(value *bool)() {
    m.preferredLanguageScreenDisabled = value
}
// SetRestoreCompletedScreenDisabled sets the restoreCompletedScreenDisabled property value. Indicates if Weclome screen is disabled
func (m *DepIOSEnrollmentProfile) SetRestoreCompletedScreenDisabled(value *bool)() {
    m.restoreCompletedScreenDisabled = value
}
// SetRestoreFromAndroidDisabled sets the restoreFromAndroidDisabled property value. Indicates if Restore from Android is disabled
func (m *DepIOSEnrollmentProfile) SetRestoreFromAndroidDisabled(value *bool)() {
    m.restoreFromAndroidDisabled = value
}
// SetSharedIPadMaximumUserCount sets the sharedIPadMaximumUserCount property value. This specifies the maximum number of users that can use a shared iPad. Only applicable in shared iPad mode.
func (m *DepIOSEnrollmentProfile) SetSharedIPadMaximumUserCount(value *int32)() {
    m.sharedIPadMaximumUserCount = value
}
// SetSimSetupScreenDisabled sets the simSetupScreenDisabled property value. Indicates if the SIMSetup screen is disabled
func (m *DepIOSEnrollmentProfile) SetSimSetupScreenDisabled(value *bool)() {
    m.simSetupScreenDisabled = value
}
// SetSoftwareUpdateScreenDisabled sets the softwareUpdateScreenDisabled property value. Indicates if the mandatory sofware update screen is disabled
func (m *DepIOSEnrollmentProfile) SetSoftwareUpdateScreenDisabled(value *bool)() {
    m.softwareUpdateScreenDisabled = value
}
// SetTemporarySessionTimeoutInSeconds sets the temporarySessionTimeoutInSeconds property value. Indicates timeout of temporary session
func (m *DepIOSEnrollmentProfile) SetTemporarySessionTimeoutInSeconds(value *int32)() {
    m.temporarySessionTimeoutInSeconds = value
}
// SetUpdateCompleteScreenDisabled sets the updateCompleteScreenDisabled property value. Indicates if Weclome screen is disabled
func (m *DepIOSEnrollmentProfile) SetUpdateCompleteScreenDisabled(value *bool)() {
    m.updateCompleteScreenDisabled = value
}
// SetUserlessSharedAadModeEnabled sets the userlessSharedAadModeEnabled property value. Indicates that this apple device is designated to support 'shared device mode' scenarios. This is distinct from the 'shared iPad' scenario. See https://learn.microsoft.com/en-us/mem/intune/enrollment/device-enrollment-shared-ios
func (m *DepIOSEnrollmentProfile) SetUserlessSharedAadModeEnabled(value *bool)() {
    m.userlessSharedAadModeEnabled = value
}
// SetUserSessionTimeoutInSeconds sets the userSessionTimeoutInSeconds property value. Indicates timeout of temporary session
func (m *DepIOSEnrollmentProfile) SetUserSessionTimeoutInSeconds(value *int32)() {
    m.userSessionTimeoutInSeconds = value
}
// SetWatchMigrationScreenDisabled sets the watchMigrationScreenDisabled property value. Indicates if the watch migration screen is disabled
func (m *DepIOSEnrollmentProfile) SetWatchMigrationScreenDisabled(value *bool)() {
    m.watchMigrationScreenDisabled = value
}
// SetWelcomeScreenDisabled sets the welcomeScreenDisabled property value. Indicates if Weclome screen is disabled
func (m *DepIOSEnrollmentProfile) SetWelcomeScreenDisabled(value *bool)() {
    m.welcomeScreenDisabled = value
}
// SetZoomDisabled sets the zoomDisabled property value. Indicates if zoom setup pane is disabled
func (m *DepIOSEnrollmentProfile) SetZoomDisabled(value *bool)() {
    m.zoomDisabled = value
}
