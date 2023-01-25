package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerGeneralDeviceConfiguration 
type AndroidDeviceOwnerGeneralDeviceConfiguration struct {
    DeviceConfiguration
    // Indicates whether or not adding or removing accounts is disabled.
    accountsBlockModification *bool
    // Indicates whether or not the user is allowed to enable to unknown sources setting.
    appsAllowInstallFromUnknownSources *bool
    // Indicates the value of the app auto update policy. Possible values are: notConfigured, userChoice, never, wiFiOnly, always.
    appsAutoUpdatePolicy *AndroidDeviceOwnerAppAutoUpdatePolicyType
    // Indicates the permission policy for requests for runtime permissions if one is not defined for the app specifically. Possible values are: deviceDefault, prompt, autoGrant, autoDeny.
    appsDefaultPermissionPolicy *AndroidDeviceOwnerDefaultAppPermissionPolicyType
    // Whether or not to recommend all apps skip any first-time-use hints they may have added.
    appsRecommendSkippingFirstUseHints *bool
    // A list of managed apps that will have their data cleared during a global sign-out in AAD shared device mode. This collection can contain a maximum of 500 elements.
    azureAdSharedDeviceDataClearApps []AppListItemable
    // Indicates whether or not to block a user from configuring bluetooth.
    bluetoothBlockConfiguration *bool
    // Indicates whether or not to block a user from sharing contacts via bluetooth.
    bluetoothBlockContactSharing *bool
    // Indicates whether or not to disable the use of the camera.
    cameraBlocked *bool
    // Indicates whether or not to block Wi-Fi tethering.
    cellularBlockWiFiTethering *bool
    // Indicates whether or not to block users from any certificate credential configuration.
    certificateCredentialConfigurationDisabled *bool
    // Indicates whether or not text copied from one profile (personal or work) can be pasted in the other.
    crossProfilePoliciesAllowCopyPaste *bool
    // Indicates whether data from one profile (personal or work) can be shared with apps in the other profile. Possible values are: notConfigured, crossProfileDataSharingBlocked, dataSharingFromWorkToPersonalBlocked, crossProfileDataSharingAllowed, unkownFutureValue.
    crossProfilePoliciesAllowDataSharing *AndroidDeviceOwnerCrossProfileDataSharing
    // Indicates whether or not contacts stored in work profile are shown in personal profile contact searches/incoming calls.
    crossProfilePoliciesShowWorkContactsInPersonalProfile *bool
    // Indicates whether or not to block a user from data roaming.
    dataRoamingBlocked *bool
    // Indicates whether or not to block the user from manually changing the date or time on the device
    dateTimeConfigurationBlocked *bool
    // Represents the customized detailed help text provided to users when they attempt to modify managed settings on their device.
    detailedHelpText AndroidDeviceOwnerUserFacingMessageable
    // Represents the customized lock screen message provided to users when they attempt to modify managed settings on their device.
    deviceOwnerLockScreenMessage AndroidDeviceOwnerUserFacingMessageable
    // Android Device Owner Enrollment Profile types.
    enrollmentProfile *AndroidDeviceOwnerEnrollmentProfileType
    // Indicates whether or not the factory reset option in settings is disabled.
    factoryResetBlocked *bool
    // List of Google account emails that will be required to authenticate after a device is factory reset before it can be set up.
    factoryResetDeviceAdministratorEmails []string
    // Proxy is set up directly with host, port and excluded hosts.
    globalProxy AndroidDeviceOwnerGlobalProxyable
    // Indicates whether or not google accounts will be blocked.
    googleAccountsBlocked *bool
    // Indicates whether a user can access the device's Settings app while in Kiosk Mode.
    kioskCustomizationDeviceSettingsBlocked *bool
    // Whether the power menu is shown when a user long presses the Power button of a device in Kiosk Mode.
    kioskCustomizationPowerButtonActionsBlocked *bool
    // Indicates whether system info and notifications are disabled in Kiosk Mode. Possible values are: notConfigured, notificationsAndSystemInfoEnabled, systemInfoOnly.
    kioskCustomizationStatusBar *AndroidDeviceOwnerKioskCustomizationStatusBar
    // Indicates whether system error dialogs for crashed or unresponsive apps are shown in Kiosk Mode.
    kioskCustomizationSystemErrorWarnings *bool
    // Indicates which navigation features are enabled in Kiosk Mode. Possible values are: notConfigured, navigationEnabled, homeButtonOnly.
    kioskCustomizationSystemNavigation *AndroidDeviceOwnerKioskCustomizationSystemNavigation
    // Whether or not to enable app ordering in Kiosk Mode.
    kioskModeAppOrderEnabled *bool
    // The ordering of items on Kiosk Mode Managed Home Screen. This collection can contain a maximum of 500 elements.
    kioskModeAppPositions []AndroidDeviceOwnerKioskModeAppPositionItemable
    // A list of managed apps that will be shown when the device is in Kiosk Mode. This collection can contain a maximum of 500 elements.
    kioskModeApps []AppListItemable
    // Whether or not to alphabetize applications within a folder in Kiosk Mode.
    kioskModeAppsInFolderOrderedByName *bool
    // Whether or not to allow a user to configure Bluetooth settings in Kiosk Mode.
    kioskModeBluetoothConfigurationEnabled *bool
    // Whether or not to allow a user to easy access to the debug menu in Kiosk Mode.
    kioskModeDebugMenuEasyAccessEnabled *bool
    // Exit code to allow a user to escape from Kiosk Mode when the device is in Kiosk Mode.
    kioskModeExitCode *string
    // Whether or not to allow a user to use the flashlight in Kiosk Mode.
    kioskModeFlashlightConfigurationEnabled *bool
    // Folder icon configuration for managed home screen in Kiosk Mode. Possible values are: notConfigured, darkSquare, darkCircle, lightSquare, lightCircle.
    kioskModeFolderIcon *AndroidDeviceOwnerKioskModeFolderIcon
    // Number of rows for Managed Home Screen grid with app ordering enabled in Kiosk Mode. Valid values 1 to 9999999
    kioskModeGridHeight *int32
    // Number of columns for Managed Home Screen grid with app ordering enabled in Kiosk Mode. Valid values 1 to 9999999
    kioskModeGridWidth *int32
    // Icon size configuration for managed home screen in Kiosk Mode. Possible values are: notConfigured, smallest, small, regular, large, largest.
    kioskModeIconSize *AndroidDeviceOwnerKioskModeIconSize
    // Whether or not to lock home screen to the end user in Kiosk Mode.
    kioskModeLockHomeScreen *bool
    // A list of managed folders for a device in Kiosk Mode. This collection can contain a maximum of 500 elements.
    kioskModeManagedFolders []AndroidDeviceOwnerKioskModeManagedFolderable
    // Whether or not to automatically sign-out of MHS and Shared device mode applications after inactive for Managed Home Screen.
    kioskModeManagedHomeScreenAutoSignout *bool
    // Number of seconds to give user notice before automatically signing them out for Managed Home Screen. Valid values 0 to 9999999
    kioskModeManagedHomeScreenInactiveSignOutDelayInSeconds *int32
    // Number of seconds device is inactive before automatically signing user out for Managed Home Screen. Valid values 0 to 9999999
    kioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds *int32
    // Complexity of PIN for sign-in session for Managed Home Screen. Possible values are: notConfigured, simple, complex.
    kioskModeManagedHomeScreenPinComplexity *KioskModeManagedHomeScreenPinComplexity
    // Whether or not require user to set a PIN for sign-in session for Managed Home Screen.
    kioskModeManagedHomeScreenPinRequired *bool
    // Whether or not required user to enter session PIN if screensaver has appeared for Managed Home Screen.
    kioskModeManagedHomeScreenPinRequiredToResume *bool
    // Custom URL background for sign-in screen for Managed Home Screen.
    kioskModeManagedHomeScreenSignInBackground *string
    // Custom URL branding logo for sign-in screen and session pin page for Managed Home Screen.
    kioskModeManagedHomeScreenSignInBrandingLogo *string
    // Whether or not show sign-in screen for Managed Home Screen.
    kioskModeManagedHomeScreenSignInEnabled *bool
    // Whether or not to display the Managed Settings entry point on the managed home screen in Kiosk Mode.
    kioskModeManagedSettingsEntryDisabled *bool
    // Whether or not to allow a user to change the media volume in Kiosk Mode.
    kioskModeMediaVolumeConfigurationEnabled *bool
    // Screen orientation configuration for managed home screen in Kiosk Mode. Possible values are: notConfigured, portrait, landscape, autoRotate.
    kioskModeScreenOrientation *AndroidDeviceOwnerKioskModeScreenOrientation
    // Whether or not to enable screen saver mode or not in Kiosk Mode.
    kioskModeScreenSaverConfigurationEnabled *bool
    // Whether or not the device screen should show the screen saver if audio/video is playing in Kiosk Mode.
    kioskModeScreenSaverDetectMediaDisabled *bool
    // The number of seconds that the device will display the screen saver for in Kiosk Mode. Valid values 0 to 9999999
    kioskModeScreenSaverDisplayTimeInSeconds *int32
    // URL for an image that will be the device's screen saver in Kiosk Mode.
    kioskModeScreenSaverImageUrl *string
    // The number of seconds the device needs to be inactive for before the screen saver is shown in Kiosk Mode. Valid values 1 to 9999999
    kioskModeScreenSaverStartDelayInSeconds *int32
    // Whether or not to display application notification badges in Kiosk Mode.
    kioskModeShowAppNotificationBadge *bool
    // Whether or not to allow a user to access basic device information.
    kioskModeShowDeviceInfo *bool
    // Whether or not to use single app kiosk mode or multi-app kiosk mode. Possible values are: notConfigured, singleAppMode, multiAppMode.
    kioskModeUseManagedHomeScreenApp *KioskModeType
    // Whether or not to display a virtual home button when the device is in Kiosk Mode.
    kioskModeVirtualHomeButtonEnabled *bool
    // Indicates whether the virtual home button is a swipe up home button or a floating home button. Possible values are: notConfigured, swipeUp, floating.
    kioskModeVirtualHomeButtonType *AndroidDeviceOwnerVirtualHomeButtonType
    // URL to a publicly accessible image to use for the wallpaper when the device is in Kiosk Mode.
    kioskModeWallpaperUrl *string
    // The restricted set of WIFI SSIDs available for the user to configure in Kiosk Mode. This collection can contain a maximum of 500 elements.
    kioskModeWifiAllowedSsids []string
    // Whether or not to allow a user to configure Wi-Fi settings in Kiosk Mode.
    kioskModeWiFiConfigurationEnabled *bool
    // Indicates whether or not to block unmuting the microphone on the device.
    microphoneForceMute *bool
    // Indicates whether or not to you want configure Microsoft Launcher.
    microsoftLauncherConfigurationEnabled *bool
    // Indicates whether or not the user can modify the wallpaper to personalize their device.
    microsoftLauncherCustomWallpaperAllowUserModification *bool
    // Indicates whether or not to configure the wallpaper on the targeted devices.
    microsoftLauncherCustomWallpaperEnabled *bool
    // Indicates the URL for the image file to use as the wallpaper on the targeted devices.
    microsoftLauncherCustomWallpaperImageUrl *string
    // Indicates whether or not the user can modify the device dock configuration on the device.
    microsoftLauncherDockPresenceAllowUserModification *bool
    // Indicates whether or not you want to configure the device dock. Possible values are: notConfigured, show, hide, disabled.
    microsoftLauncherDockPresenceConfiguration *MicrosoftLauncherDockPresence
    // Indicates whether or not the user can modify the launcher feed on the device.
    microsoftLauncherFeedAllowUserModification *bool
    // Indicates whether or not you want to enable the launcher feed on the device.
    microsoftLauncherFeedEnabled *bool
    // Indicates the search bar placement configuration on the device. Possible values are: notConfigured, top, bottom, hide.
    microsoftLauncherSearchBarPlacementConfiguration *MicrosoftLauncherSearchBarPlacement
    // Indicates whether or not the device will allow connecting to a temporary network connection at boot time.
    networkEscapeHatchAllowed *bool
    // Indicates whether or not to block NFC outgoing beam.
    nfcBlockOutgoingBeam *bool
    // Indicates whether or not the keyguard is disabled.
    passwordBlockKeyguard *bool
    // List of device keyguard features to block. This collection can contain a maximum of 11 elements.
    passwordBlockKeyguardFeatures []AndroidKeyguardFeature
    // Indicates the amount of time that a password can be set for before it expires and a new password will be required. Valid values 1 to 365
    passwordExpirationDays *int32
    // Indicates the minimum length of the password required on the device. Valid values 4 to 16
    passwordMinimumLength *int32
    // Indicates the minimum number of letter characters required for device password. Valid values 1 to 16
    passwordMinimumLetterCharacters *int32
    // Indicates the minimum number of lower case characters required for device password. Valid values 1 to 16
    passwordMinimumLowerCaseCharacters *int32
    // Indicates the minimum number of non-letter characters required for device password. Valid values 1 to 16
    passwordMinimumNonLetterCharacters *int32
    // Indicates the minimum number of numeric characters required for device password. Valid values 1 to 16
    passwordMinimumNumericCharacters *int32
    // Indicates the minimum number of symbol characters required for device password. Valid values 1 to 16
    passwordMinimumSymbolCharacters *int32
    // Indicates the minimum number of upper case letter characters required for device password. Valid values 1 to 16
    passwordMinimumUpperCaseCharacters *int32
    // Minutes of inactivity before the screen times out.
    passwordMinutesOfInactivityBeforeScreenTimeout *int32
    // Indicates the length of password history, where the user will not be able to enter a new password that is the same as any password in the history. Valid values 0 to 24
    passwordPreviousPasswordCountToBlock *int32
    // Indicates the minimum password quality required on the device. Possible values are: deviceDefault, required, numeric, numericComplex, alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, customPassword.
    passwordRequiredType *AndroidDeviceOwnerRequiredPasswordType
    // Indicates the timeout period after which a device must be unlocked using a form of strong authentication. Possible values are: deviceDefault, daily, unkownFutureValue.
    passwordRequireUnlock *AndroidDeviceOwnerRequiredPasswordUnlock
    // Indicates the number of times a user can enter an incorrect password before the device is wiped. Valid values 4 to 11
    passwordSignInFailureCountBeforeFactoryReset *int32
    // Indicates whether the user can install apps from unknown sources on the personal profile.
    personalProfileAppsAllowInstallFromUnknownSources *bool
    // Indicates whether to disable the use of the camera on the personal profile.
    personalProfileCameraBlocked *bool
    // Policy applied to applications in the personal profile. This collection can contain a maximum of 500 elements.
    personalProfilePersonalApplications []AppListItemable
    // Used together with PersonalProfilePersonalApplications to control how apps in the personal profile are allowed or blocked. Possible values are: notConfigured, blockedApps, allowedApps.
    personalProfilePlayStoreMode *PersonalProfilePersonalPlayStoreMode
    // Indicates whether to disable the capability to take screenshots on the personal profile.
    personalProfileScreenCaptureBlocked *bool
    // Indicates the Play Store mode of the device. Possible values are: notConfigured, allowList, blockList.
    playStoreMode *AndroidDeviceOwnerPlayStoreMode
    // Indicates whether or not to disable the capability to take screenshots.
    screenCaptureBlocked *bool
    // Represents the security common criteria mode enabled provided to users when they attempt to modify managed settings on their device.
    securityCommonCriteriaModeEnabled *bool
    // Indicates whether or not the user is allowed to access developer settings like developer options and safe boot on the device.
    securityDeveloperSettingsEnabled *bool
    // Indicates whether or not verify apps is required.
    securityRequireVerifyApps *bool
    // Represents the customized short help text provided to users when they attempt to modify managed settings on their device.
    shortHelpText AndroidDeviceOwnerUserFacingMessageable
    // Indicates whether or the status bar is disabled, including notifications, quick settings and other screen overlays.
    statusBarBlocked *bool
    // List of modes in which the device's display will stay powered-on. This collection can contain a maximum of 4 elements.
    stayOnModes []AndroidDeviceOwnerBatteryPluggedMode
    // Indicates whether or not to allow USB mass storage.
    storageAllowUsb *bool
    // Indicates whether or not to block external media.
    storageBlockExternalMedia *bool
    // Indicates whether or not to block USB file transfer.
    storageBlockUsbFileTransfer *bool
    // Indicates the annually repeating time periods during which system updates are postponed. This collection can contain a maximum of 500 elements.
    systemUpdateFreezePeriods []AndroidDeviceOwnerSystemUpdateFreezePeriodable
    // The type of system update configuration. Possible values are: deviceDefault, postpone, windowed, automatic.
    systemUpdateInstallType *AndroidDeviceOwnerSystemUpdateInstallType
    // Indicates the number of minutes after midnight that the system update window ends. Valid values 0 to 1440
    systemUpdateWindowEndMinutesAfterMidnight *int32
    // Indicates the number of minutes after midnight that the system update window starts. Valid values 0 to 1440
    systemUpdateWindowStartMinutesAfterMidnight *int32
    // Whether or not to block Android system prompt windows, like toasts, phone activities, and system alerts.
    systemWindowsBlocked *bool
    // Indicates whether or not adding users and profiles is disabled.
    usersBlockAdd *bool
    // Indicates whether or not to disable removing other users from the device.
    usersBlockRemove *bool
    // Indicates whether or not adjusting the master volume is disabled.
    volumeBlockAdjustment *bool
    // If an always on VPN package name is specified, whether or not to lock network traffic when that VPN is disconnected.
    vpnAlwaysOnLockdownMode *bool
    // Android app package name for app that will handle an always-on VPN connection.
    vpnAlwaysOnPackageIdentifier *string
    // Indicates whether or not to block the user from editing the wifi connection settings.
    wifiBlockEditConfigurations *bool
    // Indicates whether or not to block the user from editing just the networks defined by the policy.
    wifiBlockEditPolicyDefinedConfigurations *bool
    // Indicates the number of days that a work profile password can be set before it expires and a new password will be required. Valid values 1 to 365
    workProfilePasswordExpirationDays *int32
    // Indicates the minimum length of the work profile password. Valid values 4 to 16
    workProfilePasswordMinimumLength *int32
    // Indicates the minimum number of letter characters required for the work profile password. Valid values 1 to 16
    workProfilePasswordMinimumLetterCharacters *int32
    // Indicates the minimum number of lower-case characters required for the work profile password. Valid values 1 to 16
    workProfilePasswordMinimumLowerCaseCharacters *int32
    // Indicates the minimum number of non-letter characters required for the work profile password. Valid values 1 to 16
    workProfilePasswordMinimumNonLetterCharacters *int32
    // Indicates the minimum number of numeric characters required for the work profile password. Valid values 1 to 16
    workProfilePasswordMinimumNumericCharacters *int32
    // Indicates the minimum number of symbol characters required for the work profile password. Valid values 1 to 16
    workProfilePasswordMinimumSymbolCharacters *int32
    // Indicates the minimum number of upper-case letter characters required for the work profile password. Valid values 1 to 16
    workProfilePasswordMinimumUpperCaseCharacters *int32
    // Indicates the length of the work profile password history, where the user will not be able to enter a new password that is the same as any password in the history. Valid values 0 to 24
    workProfilePasswordPreviousPasswordCountToBlock *int32
    // Indicates the minimum password quality required on the work profile password. Possible values are: deviceDefault, required, numeric, numericComplex, alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, customPassword.
    workProfilePasswordRequiredType *AndroidDeviceOwnerRequiredPasswordType
    // Indicates the timeout period after which a work profile must be unlocked using a form of strong authentication. Possible values are: deviceDefault, daily, unkownFutureValue.
    workProfilePasswordRequireUnlock *AndroidDeviceOwnerRequiredPasswordUnlock
    // Indicates the number of times a user can enter an incorrect work profile password before the device is wiped. Valid values 4 to 11
    workProfilePasswordSignInFailureCountBeforeFactoryReset *int32
}
// NewAndroidDeviceOwnerGeneralDeviceConfiguration instantiates a new AndroidDeviceOwnerGeneralDeviceConfiguration and sets the default values.
func NewAndroidDeviceOwnerGeneralDeviceConfiguration()(*AndroidDeviceOwnerGeneralDeviceConfiguration) {
    m := &AndroidDeviceOwnerGeneralDeviceConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidDeviceOwnerGeneralDeviceConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidDeviceOwnerGeneralDeviceConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerGeneralDeviceConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerGeneralDeviceConfiguration(), nil
}
// GetAccountsBlockModification gets the accountsBlockModification property value. Indicates whether or not adding or removing accounts is disabled.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetAccountsBlockModification()(*bool) {
    return m.accountsBlockModification
}
// GetAppsAllowInstallFromUnknownSources gets the appsAllowInstallFromUnknownSources property value. Indicates whether or not the user is allowed to enable to unknown sources setting.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetAppsAllowInstallFromUnknownSources()(*bool) {
    return m.appsAllowInstallFromUnknownSources
}
// GetAppsAutoUpdatePolicy gets the appsAutoUpdatePolicy property value. Indicates the value of the app auto update policy. Possible values are: notConfigured, userChoice, never, wiFiOnly, always.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetAppsAutoUpdatePolicy()(*AndroidDeviceOwnerAppAutoUpdatePolicyType) {
    return m.appsAutoUpdatePolicy
}
// GetAppsDefaultPermissionPolicy gets the appsDefaultPermissionPolicy property value. Indicates the permission policy for requests for runtime permissions if one is not defined for the app specifically. Possible values are: deviceDefault, prompt, autoGrant, autoDeny.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetAppsDefaultPermissionPolicy()(*AndroidDeviceOwnerDefaultAppPermissionPolicyType) {
    return m.appsDefaultPermissionPolicy
}
// GetAppsRecommendSkippingFirstUseHints gets the appsRecommendSkippingFirstUseHints property value. Whether or not to recommend all apps skip any first-time-use hints they may have added.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetAppsRecommendSkippingFirstUseHints()(*bool) {
    return m.appsRecommendSkippingFirstUseHints
}
// GetAzureAdSharedDeviceDataClearApps gets the azureAdSharedDeviceDataClearApps property value. A list of managed apps that will have their data cleared during a global sign-out in AAD shared device mode. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetAzureAdSharedDeviceDataClearApps()([]AppListItemable) {
    return m.azureAdSharedDeviceDataClearApps
}
// GetBluetoothBlockConfiguration gets the bluetoothBlockConfiguration property value. Indicates whether or not to block a user from configuring bluetooth.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetBluetoothBlockConfiguration()(*bool) {
    return m.bluetoothBlockConfiguration
}
// GetBluetoothBlockContactSharing gets the bluetoothBlockContactSharing property value. Indicates whether or not to block a user from sharing contacts via bluetooth.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetBluetoothBlockContactSharing()(*bool) {
    return m.bluetoothBlockContactSharing
}
// GetCameraBlocked gets the cameraBlocked property value. Indicates whether or not to disable the use of the camera.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetCameraBlocked()(*bool) {
    return m.cameraBlocked
}
// GetCellularBlockWiFiTethering gets the cellularBlockWiFiTethering property value. Indicates whether or not to block Wi-Fi tethering.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetCellularBlockWiFiTethering()(*bool) {
    return m.cellularBlockWiFiTethering
}
// GetCertificateCredentialConfigurationDisabled gets the certificateCredentialConfigurationDisabled property value. Indicates whether or not to block users from any certificate credential configuration.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetCertificateCredentialConfigurationDisabled()(*bool) {
    return m.certificateCredentialConfigurationDisabled
}
// GetCrossProfilePoliciesAllowCopyPaste gets the crossProfilePoliciesAllowCopyPaste property value. Indicates whether or not text copied from one profile (personal or work) can be pasted in the other.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetCrossProfilePoliciesAllowCopyPaste()(*bool) {
    return m.crossProfilePoliciesAllowCopyPaste
}
// GetCrossProfilePoliciesAllowDataSharing gets the crossProfilePoliciesAllowDataSharing property value. Indicates whether data from one profile (personal or work) can be shared with apps in the other profile. Possible values are: notConfigured, crossProfileDataSharingBlocked, dataSharingFromWorkToPersonalBlocked, crossProfileDataSharingAllowed, unkownFutureValue.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetCrossProfilePoliciesAllowDataSharing()(*AndroidDeviceOwnerCrossProfileDataSharing) {
    return m.crossProfilePoliciesAllowDataSharing
}
// GetCrossProfilePoliciesShowWorkContactsInPersonalProfile gets the crossProfilePoliciesShowWorkContactsInPersonalProfile property value. Indicates whether or not contacts stored in work profile are shown in personal profile contact searches/incoming calls.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetCrossProfilePoliciesShowWorkContactsInPersonalProfile()(*bool) {
    return m.crossProfilePoliciesShowWorkContactsInPersonalProfile
}
// GetDataRoamingBlocked gets the dataRoamingBlocked property value. Indicates whether or not to block a user from data roaming.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetDataRoamingBlocked()(*bool) {
    return m.dataRoamingBlocked
}
// GetDateTimeConfigurationBlocked gets the dateTimeConfigurationBlocked property value. Indicates whether or not to block the user from manually changing the date or time on the device
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetDateTimeConfigurationBlocked()(*bool) {
    return m.dateTimeConfigurationBlocked
}
// GetDetailedHelpText gets the detailedHelpText property value. Represents the customized detailed help text provided to users when they attempt to modify managed settings on their device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetDetailedHelpText()(AndroidDeviceOwnerUserFacingMessageable) {
    return m.detailedHelpText
}
// GetDeviceOwnerLockScreenMessage gets the deviceOwnerLockScreenMessage property value. Represents the customized lock screen message provided to users when they attempt to modify managed settings on their device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetDeviceOwnerLockScreenMessage()(AndroidDeviceOwnerUserFacingMessageable) {
    return m.deviceOwnerLockScreenMessage
}
// GetEnrollmentProfile gets the enrollmentProfile property value. Android Device Owner Enrollment Profile types.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetEnrollmentProfile()(*AndroidDeviceOwnerEnrollmentProfileType) {
    return m.enrollmentProfile
}
// GetFactoryResetBlocked gets the factoryResetBlocked property value. Indicates whether or not the factory reset option in settings is disabled.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetFactoryResetBlocked()(*bool) {
    return m.factoryResetBlocked
}
// GetFactoryResetDeviceAdministratorEmails gets the factoryResetDeviceAdministratorEmails property value. List of Google account emails that will be required to authenticate after a device is factory reset before it can be set up.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetFactoryResetDeviceAdministratorEmails()([]string) {
    return m.factoryResetDeviceAdministratorEmails
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["accountsBlockModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountsBlockModification(val)
        }
        return nil
    }
    res["appsAllowInstallFromUnknownSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppsAllowInstallFromUnknownSources(val)
        }
        return nil
    }
    res["appsAutoUpdatePolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerAppAutoUpdatePolicyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppsAutoUpdatePolicy(val.(*AndroidDeviceOwnerAppAutoUpdatePolicyType))
        }
        return nil
    }
    res["appsDefaultPermissionPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerDefaultAppPermissionPolicyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppsDefaultPermissionPolicy(val.(*AndroidDeviceOwnerDefaultAppPermissionPolicyType))
        }
        return nil
    }
    res["appsRecommendSkippingFirstUseHints"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppsRecommendSkippingFirstUseHints(val)
        }
        return nil
    }
    res["azureAdSharedDeviceDataClearApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppListItemable, len(val))
            for i, v := range val {
                res[i] = v.(AppListItemable)
            }
            m.SetAzureAdSharedDeviceDataClearApps(res)
        }
        return nil
    }
    res["bluetoothBlockConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBluetoothBlockConfiguration(val)
        }
        return nil
    }
    res["bluetoothBlockContactSharing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBluetoothBlockContactSharing(val)
        }
        return nil
    }
    res["cameraBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCameraBlocked(val)
        }
        return nil
    }
    res["cellularBlockWiFiTethering"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCellularBlockWiFiTethering(val)
        }
        return nil
    }
    res["certificateCredentialConfigurationDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateCredentialConfigurationDisabled(val)
        }
        return nil
    }
    res["crossProfilePoliciesAllowCopyPaste"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCrossProfilePoliciesAllowCopyPaste(val)
        }
        return nil
    }
    res["crossProfilePoliciesAllowDataSharing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerCrossProfileDataSharing)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCrossProfilePoliciesAllowDataSharing(val.(*AndroidDeviceOwnerCrossProfileDataSharing))
        }
        return nil
    }
    res["crossProfilePoliciesShowWorkContactsInPersonalProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCrossProfilePoliciesShowWorkContactsInPersonalProfile(val)
        }
        return nil
    }
    res["dataRoamingBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataRoamingBlocked(val)
        }
        return nil
    }
    res["dateTimeConfigurationBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateTimeConfigurationBlocked(val)
        }
        return nil
    }
    res["detailedHelpText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidDeviceOwnerUserFacingMessageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetailedHelpText(val.(AndroidDeviceOwnerUserFacingMessageable))
        }
        return nil
    }
    res["deviceOwnerLockScreenMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidDeviceOwnerUserFacingMessageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceOwnerLockScreenMessage(val.(AndroidDeviceOwnerUserFacingMessageable))
        }
        return nil
    }
    res["enrollmentProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerEnrollmentProfileType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentProfile(val.(*AndroidDeviceOwnerEnrollmentProfileType))
        }
        return nil
    }
    res["factoryResetBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFactoryResetBlocked(val)
        }
        return nil
    }
    res["factoryResetDeviceAdministratorEmails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetFactoryResetDeviceAdministratorEmails(res)
        }
        return nil
    }
    res["globalProxy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidDeviceOwnerGlobalProxyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGlobalProxy(val.(AndroidDeviceOwnerGlobalProxyable))
        }
        return nil
    }
    res["googleAccountsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGoogleAccountsBlocked(val)
        }
        return nil
    }
    res["kioskCustomizationDeviceSettingsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskCustomizationDeviceSettingsBlocked(val)
        }
        return nil
    }
    res["kioskCustomizationPowerButtonActionsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskCustomizationPowerButtonActionsBlocked(val)
        }
        return nil
    }
    res["kioskCustomizationStatusBar"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerKioskCustomizationStatusBar)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskCustomizationStatusBar(val.(*AndroidDeviceOwnerKioskCustomizationStatusBar))
        }
        return nil
    }
    res["kioskCustomizationSystemErrorWarnings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskCustomizationSystemErrorWarnings(val)
        }
        return nil
    }
    res["kioskCustomizationSystemNavigation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerKioskCustomizationSystemNavigation)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskCustomizationSystemNavigation(val.(*AndroidDeviceOwnerKioskCustomizationSystemNavigation))
        }
        return nil
    }
    res["kioskModeAppOrderEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAppOrderEnabled(val)
        }
        return nil
    }
    res["kioskModeAppPositions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidDeviceOwnerKioskModeAppPositionItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidDeviceOwnerKioskModeAppPositionItemable, len(val))
            for i, v := range val {
                res[i] = v.(AndroidDeviceOwnerKioskModeAppPositionItemable)
            }
            m.SetKioskModeAppPositions(res)
        }
        return nil
    }
    res["kioskModeApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppListItemable, len(val))
            for i, v := range val {
                res[i] = v.(AppListItemable)
            }
            m.SetKioskModeApps(res)
        }
        return nil
    }
    res["kioskModeAppsInFolderOrderedByName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAppsInFolderOrderedByName(val)
        }
        return nil
    }
    res["kioskModeBluetoothConfigurationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeBluetoothConfigurationEnabled(val)
        }
        return nil
    }
    res["kioskModeDebugMenuEasyAccessEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeDebugMenuEasyAccessEnabled(val)
        }
        return nil
    }
    res["kioskModeExitCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeExitCode(val)
        }
        return nil
    }
    res["kioskModeFlashlightConfigurationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeFlashlightConfigurationEnabled(val)
        }
        return nil
    }
    res["kioskModeFolderIcon"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerKioskModeFolderIcon)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeFolderIcon(val.(*AndroidDeviceOwnerKioskModeFolderIcon))
        }
        return nil
    }
    res["kioskModeGridHeight"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeGridHeight(val)
        }
        return nil
    }
    res["kioskModeGridWidth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeGridWidth(val)
        }
        return nil
    }
    res["kioskModeIconSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerKioskModeIconSize)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeIconSize(val.(*AndroidDeviceOwnerKioskModeIconSize))
        }
        return nil
    }
    res["kioskModeLockHomeScreen"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeLockHomeScreen(val)
        }
        return nil
    }
    res["kioskModeManagedFolders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidDeviceOwnerKioskModeManagedFolderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidDeviceOwnerKioskModeManagedFolderable, len(val))
            for i, v := range val {
                res[i] = v.(AndroidDeviceOwnerKioskModeManagedFolderable)
            }
            m.SetKioskModeManagedFolders(res)
        }
        return nil
    }
    res["kioskModeManagedHomeScreenAutoSignout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeManagedHomeScreenAutoSignout(val)
        }
        return nil
    }
    res["kioskModeManagedHomeScreenInactiveSignOutDelayInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeManagedHomeScreenInactiveSignOutDelayInSeconds(val)
        }
        return nil
    }
    res["kioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds(val)
        }
        return nil
    }
    res["kioskModeManagedHomeScreenPinComplexity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseKioskModeManagedHomeScreenPinComplexity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeManagedHomeScreenPinComplexity(val.(*KioskModeManagedHomeScreenPinComplexity))
        }
        return nil
    }
    res["kioskModeManagedHomeScreenPinRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeManagedHomeScreenPinRequired(val)
        }
        return nil
    }
    res["kioskModeManagedHomeScreenPinRequiredToResume"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeManagedHomeScreenPinRequiredToResume(val)
        }
        return nil
    }
    res["kioskModeManagedHomeScreenSignInBackground"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeManagedHomeScreenSignInBackground(val)
        }
        return nil
    }
    res["kioskModeManagedHomeScreenSignInBrandingLogo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeManagedHomeScreenSignInBrandingLogo(val)
        }
        return nil
    }
    res["kioskModeManagedHomeScreenSignInEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeManagedHomeScreenSignInEnabled(val)
        }
        return nil
    }
    res["kioskModeManagedSettingsEntryDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeManagedSettingsEntryDisabled(val)
        }
        return nil
    }
    res["kioskModeMediaVolumeConfigurationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeMediaVolumeConfigurationEnabled(val)
        }
        return nil
    }
    res["kioskModeScreenOrientation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerKioskModeScreenOrientation)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeScreenOrientation(val.(*AndroidDeviceOwnerKioskModeScreenOrientation))
        }
        return nil
    }
    res["kioskModeScreenSaverConfigurationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeScreenSaverConfigurationEnabled(val)
        }
        return nil
    }
    res["kioskModeScreenSaverDetectMediaDisabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeScreenSaverDetectMediaDisabled(val)
        }
        return nil
    }
    res["kioskModeScreenSaverDisplayTimeInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeScreenSaverDisplayTimeInSeconds(val)
        }
        return nil
    }
    res["kioskModeScreenSaverImageUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeScreenSaverImageUrl(val)
        }
        return nil
    }
    res["kioskModeScreenSaverStartDelayInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeScreenSaverStartDelayInSeconds(val)
        }
        return nil
    }
    res["kioskModeShowAppNotificationBadge"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeShowAppNotificationBadge(val)
        }
        return nil
    }
    res["kioskModeShowDeviceInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeShowDeviceInfo(val)
        }
        return nil
    }
    res["kioskModeUseManagedHomeScreenApp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseKioskModeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeUseManagedHomeScreenApp(val.(*KioskModeType))
        }
        return nil
    }
    res["kioskModeVirtualHomeButtonEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeVirtualHomeButtonEnabled(val)
        }
        return nil
    }
    res["kioskModeVirtualHomeButtonType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerVirtualHomeButtonType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeVirtualHomeButtonType(val.(*AndroidDeviceOwnerVirtualHomeButtonType))
        }
        return nil
    }
    res["kioskModeWallpaperUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeWallpaperUrl(val)
        }
        return nil
    }
    res["kioskModeWifiAllowedSsids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetKioskModeWifiAllowedSsids(res)
        }
        return nil
    }
    res["kioskModeWiFiConfigurationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeWiFiConfigurationEnabled(val)
        }
        return nil
    }
    res["microphoneForceMute"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrophoneForceMute(val)
        }
        return nil
    }
    res["microsoftLauncherConfigurationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftLauncherConfigurationEnabled(val)
        }
        return nil
    }
    res["microsoftLauncherCustomWallpaperAllowUserModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftLauncherCustomWallpaperAllowUserModification(val)
        }
        return nil
    }
    res["microsoftLauncherCustomWallpaperEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftLauncherCustomWallpaperEnabled(val)
        }
        return nil
    }
    res["microsoftLauncherCustomWallpaperImageUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftLauncherCustomWallpaperImageUrl(val)
        }
        return nil
    }
    res["microsoftLauncherDockPresenceAllowUserModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftLauncherDockPresenceAllowUserModification(val)
        }
        return nil
    }
    res["microsoftLauncherDockPresenceConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMicrosoftLauncherDockPresence)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftLauncherDockPresenceConfiguration(val.(*MicrosoftLauncherDockPresence))
        }
        return nil
    }
    res["microsoftLauncherFeedAllowUserModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftLauncherFeedAllowUserModification(val)
        }
        return nil
    }
    res["microsoftLauncherFeedEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftLauncherFeedEnabled(val)
        }
        return nil
    }
    res["microsoftLauncherSearchBarPlacementConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMicrosoftLauncherSearchBarPlacement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftLauncherSearchBarPlacementConfiguration(val.(*MicrosoftLauncherSearchBarPlacement))
        }
        return nil
    }
    res["networkEscapeHatchAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkEscapeHatchAllowed(val)
        }
        return nil
    }
    res["nfcBlockOutgoingBeam"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNfcBlockOutgoingBeam(val)
        }
        return nil
    }
    res["passwordBlockKeyguard"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordBlockKeyguard(val)
        }
        return nil
    }
    res["passwordBlockKeyguardFeatures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseAndroidKeyguardFeature)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidKeyguardFeature, len(val))
            for i, v := range val {
                res[i] = *(v.(*AndroidKeyguardFeature))
            }
            m.SetPasswordBlockKeyguardFeatures(res)
        }
        return nil
    }
    res["passwordExpirationDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordExpirationDays(val)
        }
        return nil
    }
    res["passwordMinimumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumLength(val)
        }
        return nil
    }
    res["passwordMinimumLetterCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumLetterCharacters(val)
        }
        return nil
    }
    res["passwordMinimumLowerCaseCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumLowerCaseCharacters(val)
        }
        return nil
    }
    res["passwordMinimumNonLetterCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumNonLetterCharacters(val)
        }
        return nil
    }
    res["passwordMinimumNumericCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumNumericCharacters(val)
        }
        return nil
    }
    res["passwordMinimumSymbolCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumSymbolCharacters(val)
        }
        return nil
    }
    res["passwordMinimumUpperCaseCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumUpperCaseCharacters(val)
        }
        return nil
    }
    res["passwordMinutesOfInactivityBeforeScreenTimeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinutesOfInactivityBeforeScreenTimeout(val)
        }
        return nil
    }
    res["passwordPreviousPasswordCountToBlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordPreviousPasswordCountToBlock(val)
        }
        return nil
    }
    res["passwordRequiredType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerRequiredPasswordType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRequiredType(val.(*AndroidDeviceOwnerRequiredPasswordType))
        }
        return nil
    }
    res["passwordRequireUnlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerRequiredPasswordUnlock)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRequireUnlock(val.(*AndroidDeviceOwnerRequiredPasswordUnlock))
        }
        return nil
    }
    res["passwordSignInFailureCountBeforeFactoryReset"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordSignInFailureCountBeforeFactoryReset(val)
        }
        return nil
    }
    res["personalProfileAppsAllowInstallFromUnknownSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersonalProfileAppsAllowInstallFromUnknownSources(val)
        }
        return nil
    }
    res["personalProfileCameraBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersonalProfileCameraBlocked(val)
        }
        return nil
    }
    res["personalProfilePersonalApplications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppListItemable, len(val))
            for i, v := range val {
                res[i] = v.(AppListItemable)
            }
            m.SetPersonalProfilePersonalApplications(res)
        }
        return nil
    }
    res["personalProfilePlayStoreMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePersonalProfilePersonalPlayStoreMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersonalProfilePlayStoreMode(val.(*PersonalProfilePersonalPlayStoreMode))
        }
        return nil
    }
    res["personalProfileScreenCaptureBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersonalProfileScreenCaptureBlocked(val)
        }
        return nil
    }
    res["playStoreMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerPlayStoreMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlayStoreMode(val.(*AndroidDeviceOwnerPlayStoreMode))
        }
        return nil
    }
    res["screenCaptureBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScreenCaptureBlocked(val)
        }
        return nil
    }
    res["securityCommonCriteriaModeEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityCommonCriteriaModeEnabled(val)
        }
        return nil
    }
    res["securityDeveloperSettingsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityDeveloperSettingsEnabled(val)
        }
        return nil
    }
    res["securityRequireVerifyApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityRequireVerifyApps(val)
        }
        return nil
    }
    res["shortHelpText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidDeviceOwnerUserFacingMessageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShortHelpText(val.(AndroidDeviceOwnerUserFacingMessageable))
        }
        return nil
    }
    res["statusBarBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusBarBlocked(val)
        }
        return nil
    }
    res["stayOnModes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseAndroidDeviceOwnerBatteryPluggedMode)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidDeviceOwnerBatteryPluggedMode, len(val))
            for i, v := range val {
                res[i] = *(v.(*AndroidDeviceOwnerBatteryPluggedMode))
            }
            m.SetStayOnModes(res)
        }
        return nil
    }
    res["storageAllowUsb"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageAllowUsb(val)
        }
        return nil
    }
    res["storageBlockExternalMedia"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageBlockExternalMedia(val)
        }
        return nil
    }
    res["storageBlockUsbFileTransfer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageBlockUsbFileTransfer(val)
        }
        return nil
    }
    res["systemUpdateFreezePeriods"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidDeviceOwnerSystemUpdateFreezePeriodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidDeviceOwnerSystemUpdateFreezePeriodable, len(val))
            for i, v := range val {
                res[i] = v.(AndroidDeviceOwnerSystemUpdateFreezePeriodable)
            }
            m.SetSystemUpdateFreezePeriods(res)
        }
        return nil
    }
    res["systemUpdateInstallType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerSystemUpdateInstallType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemUpdateInstallType(val.(*AndroidDeviceOwnerSystemUpdateInstallType))
        }
        return nil
    }
    res["systemUpdateWindowEndMinutesAfterMidnight"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemUpdateWindowEndMinutesAfterMidnight(val)
        }
        return nil
    }
    res["systemUpdateWindowStartMinutesAfterMidnight"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemUpdateWindowStartMinutesAfterMidnight(val)
        }
        return nil
    }
    res["systemWindowsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemWindowsBlocked(val)
        }
        return nil
    }
    res["usersBlockAdd"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsersBlockAdd(val)
        }
        return nil
    }
    res["usersBlockRemove"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsersBlockRemove(val)
        }
        return nil
    }
    res["volumeBlockAdjustment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVolumeBlockAdjustment(val)
        }
        return nil
    }
    res["vpnAlwaysOnLockdownMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVpnAlwaysOnLockdownMode(val)
        }
        return nil
    }
    res["vpnAlwaysOnPackageIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVpnAlwaysOnPackageIdentifier(val)
        }
        return nil
    }
    res["wifiBlockEditConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiBlockEditConfigurations(val)
        }
        return nil
    }
    res["wifiBlockEditPolicyDefinedConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiBlockEditPolicyDefinedConfigurations(val)
        }
        return nil
    }
    res["workProfilePasswordExpirationDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordExpirationDays(val)
        }
        return nil
    }
    res["workProfilePasswordMinimumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinimumLength(val)
        }
        return nil
    }
    res["workProfilePasswordMinimumLetterCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinimumLetterCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordMinimumLowerCaseCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinimumLowerCaseCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordMinimumNonLetterCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinimumNonLetterCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordMinimumNumericCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinimumNumericCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordMinimumSymbolCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinimumSymbolCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordMinimumUpperCaseCharacters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordMinimumUpperCaseCharacters(val)
        }
        return nil
    }
    res["workProfilePasswordPreviousPasswordCountToBlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordPreviousPasswordCountToBlock(val)
        }
        return nil
    }
    res["workProfilePasswordRequiredType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerRequiredPasswordType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordRequiredType(val.(*AndroidDeviceOwnerRequiredPasswordType))
        }
        return nil
    }
    res["workProfilePasswordRequireUnlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerRequiredPasswordUnlock)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordRequireUnlock(val.(*AndroidDeviceOwnerRequiredPasswordUnlock))
        }
        return nil
    }
    res["workProfilePasswordSignInFailureCountBeforeFactoryReset"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkProfilePasswordSignInFailureCountBeforeFactoryReset(val)
        }
        return nil
    }
    return res
}
// GetGlobalProxy gets the globalProxy property value. Proxy is set up directly with host, port and excluded hosts.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetGlobalProxy()(AndroidDeviceOwnerGlobalProxyable) {
    return m.globalProxy
}
// GetGoogleAccountsBlocked gets the googleAccountsBlocked property value. Indicates whether or not google accounts will be blocked.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetGoogleAccountsBlocked()(*bool) {
    return m.googleAccountsBlocked
}
// GetKioskCustomizationDeviceSettingsBlocked gets the kioskCustomizationDeviceSettingsBlocked property value. Indicates whether a user can access the device's Settings app while in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskCustomizationDeviceSettingsBlocked()(*bool) {
    return m.kioskCustomizationDeviceSettingsBlocked
}
// GetKioskCustomizationPowerButtonActionsBlocked gets the kioskCustomizationPowerButtonActionsBlocked property value. Whether the power menu is shown when a user long presses the Power button of a device in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskCustomizationPowerButtonActionsBlocked()(*bool) {
    return m.kioskCustomizationPowerButtonActionsBlocked
}
// GetKioskCustomizationStatusBar gets the kioskCustomizationStatusBar property value. Indicates whether system info and notifications are disabled in Kiosk Mode. Possible values are: notConfigured, notificationsAndSystemInfoEnabled, systemInfoOnly.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskCustomizationStatusBar()(*AndroidDeviceOwnerKioskCustomizationStatusBar) {
    return m.kioskCustomizationStatusBar
}
// GetKioskCustomizationSystemErrorWarnings gets the kioskCustomizationSystemErrorWarnings property value. Indicates whether system error dialogs for crashed or unresponsive apps are shown in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskCustomizationSystemErrorWarnings()(*bool) {
    return m.kioskCustomizationSystemErrorWarnings
}
// GetKioskCustomizationSystemNavigation gets the kioskCustomizationSystemNavigation property value. Indicates which navigation features are enabled in Kiosk Mode. Possible values are: notConfigured, navigationEnabled, homeButtonOnly.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskCustomizationSystemNavigation()(*AndroidDeviceOwnerKioskCustomizationSystemNavigation) {
    return m.kioskCustomizationSystemNavigation
}
// GetKioskModeAppOrderEnabled gets the kioskModeAppOrderEnabled property value. Whether or not to enable app ordering in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeAppOrderEnabled()(*bool) {
    return m.kioskModeAppOrderEnabled
}
// GetKioskModeAppPositions gets the kioskModeAppPositions property value. The ordering of items on Kiosk Mode Managed Home Screen. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeAppPositions()([]AndroidDeviceOwnerKioskModeAppPositionItemable) {
    return m.kioskModeAppPositions
}
// GetKioskModeApps gets the kioskModeApps property value. A list of managed apps that will be shown when the device is in Kiosk Mode. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeApps()([]AppListItemable) {
    return m.kioskModeApps
}
// GetKioskModeAppsInFolderOrderedByName gets the kioskModeAppsInFolderOrderedByName property value. Whether or not to alphabetize applications within a folder in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeAppsInFolderOrderedByName()(*bool) {
    return m.kioskModeAppsInFolderOrderedByName
}
// GetKioskModeBluetoothConfigurationEnabled gets the kioskModeBluetoothConfigurationEnabled property value. Whether or not to allow a user to configure Bluetooth settings in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeBluetoothConfigurationEnabled()(*bool) {
    return m.kioskModeBluetoothConfigurationEnabled
}
// GetKioskModeDebugMenuEasyAccessEnabled gets the kioskModeDebugMenuEasyAccessEnabled property value. Whether or not to allow a user to easy access to the debug menu in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeDebugMenuEasyAccessEnabled()(*bool) {
    return m.kioskModeDebugMenuEasyAccessEnabled
}
// GetKioskModeExitCode gets the kioskModeExitCode property value. Exit code to allow a user to escape from Kiosk Mode when the device is in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeExitCode()(*string) {
    return m.kioskModeExitCode
}
// GetKioskModeFlashlightConfigurationEnabled gets the kioskModeFlashlightConfigurationEnabled property value. Whether or not to allow a user to use the flashlight in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeFlashlightConfigurationEnabled()(*bool) {
    return m.kioskModeFlashlightConfigurationEnabled
}
// GetKioskModeFolderIcon gets the kioskModeFolderIcon property value. Folder icon configuration for managed home screen in Kiosk Mode. Possible values are: notConfigured, darkSquare, darkCircle, lightSquare, lightCircle.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeFolderIcon()(*AndroidDeviceOwnerKioskModeFolderIcon) {
    return m.kioskModeFolderIcon
}
// GetKioskModeGridHeight gets the kioskModeGridHeight property value. Number of rows for Managed Home Screen grid with app ordering enabled in Kiosk Mode. Valid values 1 to 9999999
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeGridHeight()(*int32) {
    return m.kioskModeGridHeight
}
// GetKioskModeGridWidth gets the kioskModeGridWidth property value. Number of columns for Managed Home Screen grid with app ordering enabled in Kiosk Mode. Valid values 1 to 9999999
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeGridWidth()(*int32) {
    return m.kioskModeGridWidth
}
// GetKioskModeIconSize gets the kioskModeIconSize property value. Icon size configuration for managed home screen in Kiosk Mode. Possible values are: notConfigured, smallest, small, regular, large, largest.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeIconSize()(*AndroidDeviceOwnerKioskModeIconSize) {
    return m.kioskModeIconSize
}
// GetKioskModeLockHomeScreen gets the kioskModeLockHomeScreen property value. Whether or not to lock home screen to the end user in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeLockHomeScreen()(*bool) {
    return m.kioskModeLockHomeScreen
}
// GetKioskModeManagedFolders gets the kioskModeManagedFolders property value. A list of managed folders for a device in Kiosk Mode. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeManagedFolders()([]AndroidDeviceOwnerKioskModeManagedFolderable) {
    return m.kioskModeManagedFolders
}
// GetKioskModeManagedHomeScreenAutoSignout gets the kioskModeManagedHomeScreenAutoSignout property value. Whether or not to automatically sign-out of MHS and Shared device mode applications after inactive for Managed Home Screen.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeManagedHomeScreenAutoSignout()(*bool) {
    return m.kioskModeManagedHomeScreenAutoSignout
}
// GetKioskModeManagedHomeScreenInactiveSignOutDelayInSeconds gets the kioskModeManagedHomeScreenInactiveSignOutDelayInSeconds property value. Number of seconds to give user notice before automatically signing them out for Managed Home Screen. Valid values 0 to 9999999
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeManagedHomeScreenInactiveSignOutDelayInSeconds()(*int32) {
    return m.kioskModeManagedHomeScreenInactiveSignOutDelayInSeconds
}
// GetKioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds gets the kioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds property value. Number of seconds device is inactive before automatically signing user out for Managed Home Screen. Valid values 0 to 9999999
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds()(*int32) {
    return m.kioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds
}
// GetKioskModeManagedHomeScreenPinComplexity gets the kioskModeManagedHomeScreenPinComplexity property value. Complexity of PIN for sign-in session for Managed Home Screen. Possible values are: notConfigured, simple, complex.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeManagedHomeScreenPinComplexity()(*KioskModeManagedHomeScreenPinComplexity) {
    return m.kioskModeManagedHomeScreenPinComplexity
}
// GetKioskModeManagedHomeScreenPinRequired gets the kioskModeManagedHomeScreenPinRequired property value. Whether or not require user to set a PIN for sign-in session for Managed Home Screen.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeManagedHomeScreenPinRequired()(*bool) {
    return m.kioskModeManagedHomeScreenPinRequired
}
// GetKioskModeManagedHomeScreenPinRequiredToResume gets the kioskModeManagedHomeScreenPinRequiredToResume property value. Whether or not required user to enter session PIN if screensaver has appeared for Managed Home Screen.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeManagedHomeScreenPinRequiredToResume()(*bool) {
    return m.kioskModeManagedHomeScreenPinRequiredToResume
}
// GetKioskModeManagedHomeScreenSignInBackground gets the kioskModeManagedHomeScreenSignInBackground property value. Custom URL background for sign-in screen for Managed Home Screen.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeManagedHomeScreenSignInBackground()(*string) {
    return m.kioskModeManagedHomeScreenSignInBackground
}
// GetKioskModeManagedHomeScreenSignInBrandingLogo gets the kioskModeManagedHomeScreenSignInBrandingLogo property value. Custom URL branding logo for sign-in screen and session pin page for Managed Home Screen.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeManagedHomeScreenSignInBrandingLogo()(*string) {
    return m.kioskModeManagedHomeScreenSignInBrandingLogo
}
// GetKioskModeManagedHomeScreenSignInEnabled gets the kioskModeManagedHomeScreenSignInEnabled property value. Whether or not show sign-in screen for Managed Home Screen.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeManagedHomeScreenSignInEnabled()(*bool) {
    return m.kioskModeManagedHomeScreenSignInEnabled
}
// GetKioskModeManagedSettingsEntryDisabled gets the kioskModeManagedSettingsEntryDisabled property value. Whether or not to display the Managed Settings entry point on the managed home screen in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeManagedSettingsEntryDisabled()(*bool) {
    return m.kioskModeManagedSettingsEntryDisabled
}
// GetKioskModeMediaVolumeConfigurationEnabled gets the kioskModeMediaVolumeConfigurationEnabled property value. Whether or not to allow a user to change the media volume in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeMediaVolumeConfigurationEnabled()(*bool) {
    return m.kioskModeMediaVolumeConfigurationEnabled
}
// GetKioskModeScreenOrientation gets the kioskModeScreenOrientation property value. Screen orientation configuration for managed home screen in Kiosk Mode. Possible values are: notConfigured, portrait, landscape, autoRotate.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeScreenOrientation()(*AndroidDeviceOwnerKioskModeScreenOrientation) {
    return m.kioskModeScreenOrientation
}
// GetKioskModeScreenSaverConfigurationEnabled gets the kioskModeScreenSaverConfigurationEnabled property value. Whether or not to enable screen saver mode or not in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeScreenSaverConfigurationEnabled()(*bool) {
    return m.kioskModeScreenSaverConfigurationEnabled
}
// GetKioskModeScreenSaverDetectMediaDisabled gets the kioskModeScreenSaverDetectMediaDisabled property value. Whether or not the device screen should show the screen saver if audio/video is playing in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeScreenSaverDetectMediaDisabled()(*bool) {
    return m.kioskModeScreenSaverDetectMediaDisabled
}
// GetKioskModeScreenSaverDisplayTimeInSeconds gets the kioskModeScreenSaverDisplayTimeInSeconds property value. The number of seconds that the device will display the screen saver for in Kiosk Mode. Valid values 0 to 9999999
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeScreenSaverDisplayTimeInSeconds()(*int32) {
    return m.kioskModeScreenSaverDisplayTimeInSeconds
}
// GetKioskModeScreenSaverImageUrl gets the kioskModeScreenSaverImageUrl property value. URL for an image that will be the device's screen saver in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeScreenSaverImageUrl()(*string) {
    return m.kioskModeScreenSaverImageUrl
}
// GetKioskModeScreenSaverStartDelayInSeconds gets the kioskModeScreenSaverStartDelayInSeconds property value. The number of seconds the device needs to be inactive for before the screen saver is shown in Kiosk Mode. Valid values 1 to 9999999
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeScreenSaverStartDelayInSeconds()(*int32) {
    return m.kioskModeScreenSaverStartDelayInSeconds
}
// GetKioskModeShowAppNotificationBadge gets the kioskModeShowAppNotificationBadge property value. Whether or not to display application notification badges in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeShowAppNotificationBadge()(*bool) {
    return m.kioskModeShowAppNotificationBadge
}
// GetKioskModeShowDeviceInfo gets the kioskModeShowDeviceInfo property value. Whether or not to allow a user to access basic device information.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeShowDeviceInfo()(*bool) {
    return m.kioskModeShowDeviceInfo
}
// GetKioskModeUseManagedHomeScreenApp gets the kioskModeUseManagedHomeScreenApp property value. Whether or not to use single app kiosk mode or multi-app kiosk mode. Possible values are: notConfigured, singleAppMode, multiAppMode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeUseManagedHomeScreenApp()(*KioskModeType) {
    return m.kioskModeUseManagedHomeScreenApp
}
// GetKioskModeVirtualHomeButtonEnabled gets the kioskModeVirtualHomeButtonEnabled property value. Whether or not to display a virtual home button when the device is in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeVirtualHomeButtonEnabled()(*bool) {
    return m.kioskModeVirtualHomeButtonEnabled
}
// GetKioskModeVirtualHomeButtonType gets the kioskModeVirtualHomeButtonType property value. Indicates whether the virtual home button is a swipe up home button or a floating home button. Possible values are: notConfigured, swipeUp, floating.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeVirtualHomeButtonType()(*AndroidDeviceOwnerVirtualHomeButtonType) {
    return m.kioskModeVirtualHomeButtonType
}
// GetKioskModeWallpaperUrl gets the kioskModeWallpaperUrl property value. URL to a publicly accessible image to use for the wallpaper when the device is in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeWallpaperUrl()(*string) {
    return m.kioskModeWallpaperUrl
}
// GetKioskModeWifiAllowedSsids gets the kioskModeWifiAllowedSsids property value. The restricted set of WIFI SSIDs available for the user to configure in Kiosk Mode. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeWifiAllowedSsids()([]string) {
    return m.kioskModeWifiAllowedSsids
}
// GetKioskModeWiFiConfigurationEnabled gets the kioskModeWiFiConfigurationEnabled property value. Whether or not to allow a user to configure Wi-Fi settings in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetKioskModeWiFiConfigurationEnabled()(*bool) {
    return m.kioskModeWiFiConfigurationEnabled
}
// GetMicrophoneForceMute gets the microphoneForceMute property value. Indicates whether or not to block unmuting the microphone on the device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetMicrophoneForceMute()(*bool) {
    return m.microphoneForceMute
}
// GetMicrosoftLauncherConfigurationEnabled gets the microsoftLauncherConfigurationEnabled property value. Indicates whether or not to you want configure Microsoft Launcher.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetMicrosoftLauncherConfigurationEnabled()(*bool) {
    return m.microsoftLauncherConfigurationEnabled
}
// GetMicrosoftLauncherCustomWallpaperAllowUserModification gets the microsoftLauncherCustomWallpaperAllowUserModification property value. Indicates whether or not the user can modify the wallpaper to personalize their device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetMicrosoftLauncherCustomWallpaperAllowUserModification()(*bool) {
    return m.microsoftLauncherCustomWallpaperAllowUserModification
}
// GetMicrosoftLauncherCustomWallpaperEnabled gets the microsoftLauncherCustomWallpaperEnabled property value. Indicates whether or not to configure the wallpaper on the targeted devices.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetMicrosoftLauncherCustomWallpaperEnabled()(*bool) {
    return m.microsoftLauncherCustomWallpaperEnabled
}
// GetMicrosoftLauncherCustomWallpaperImageUrl gets the microsoftLauncherCustomWallpaperImageUrl property value. Indicates the URL for the image file to use as the wallpaper on the targeted devices.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetMicrosoftLauncherCustomWallpaperImageUrl()(*string) {
    return m.microsoftLauncherCustomWallpaperImageUrl
}
// GetMicrosoftLauncherDockPresenceAllowUserModification gets the microsoftLauncherDockPresenceAllowUserModification property value. Indicates whether or not the user can modify the device dock configuration on the device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetMicrosoftLauncherDockPresenceAllowUserModification()(*bool) {
    return m.microsoftLauncherDockPresenceAllowUserModification
}
// GetMicrosoftLauncherDockPresenceConfiguration gets the microsoftLauncherDockPresenceConfiguration property value. Indicates whether or not you want to configure the device dock. Possible values are: notConfigured, show, hide, disabled.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetMicrosoftLauncherDockPresenceConfiguration()(*MicrosoftLauncherDockPresence) {
    return m.microsoftLauncherDockPresenceConfiguration
}
// GetMicrosoftLauncherFeedAllowUserModification gets the microsoftLauncherFeedAllowUserModification property value. Indicates whether or not the user can modify the launcher feed on the device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetMicrosoftLauncherFeedAllowUserModification()(*bool) {
    return m.microsoftLauncherFeedAllowUserModification
}
// GetMicrosoftLauncherFeedEnabled gets the microsoftLauncherFeedEnabled property value. Indicates whether or not you want to enable the launcher feed on the device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetMicrosoftLauncherFeedEnabled()(*bool) {
    return m.microsoftLauncherFeedEnabled
}
// GetMicrosoftLauncherSearchBarPlacementConfiguration gets the microsoftLauncherSearchBarPlacementConfiguration property value. Indicates the search bar placement configuration on the device. Possible values are: notConfigured, top, bottom, hide.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetMicrosoftLauncherSearchBarPlacementConfiguration()(*MicrosoftLauncherSearchBarPlacement) {
    return m.microsoftLauncherSearchBarPlacementConfiguration
}
// GetNetworkEscapeHatchAllowed gets the networkEscapeHatchAllowed property value. Indicates whether or not the device will allow connecting to a temporary network connection at boot time.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetNetworkEscapeHatchAllowed()(*bool) {
    return m.networkEscapeHatchAllowed
}
// GetNfcBlockOutgoingBeam gets the nfcBlockOutgoingBeam property value. Indicates whether or not to block NFC outgoing beam.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetNfcBlockOutgoingBeam()(*bool) {
    return m.nfcBlockOutgoingBeam
}
// GetPasswordBlockKeyguard gets the passwordBlockKeyguard property value. Indicates whether or not the keyguard is disabled.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordBlockKeyguard()(*bool) {
    return m.passwordBlockKeyguard
}
// GetPasswordBlockKeyguardFeatures gets the passwordBlockKeyguardFeatures property value. List of device keyguard features to block. This collection can contain a maximum of 11 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordBlockKeyguardFeatures()([]AndroidKeyguardFeature) {
    return m.passwordBlockKeyguardFeatures
}
// GetPasswordExpirationDays gets the passwordExpirationDays property value. Indicates the amount of time that a password can be set for before it expires and a new password will be required. Valid values 1 to 365
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordExpirationDays()(*int32) {
    return m.passwordExpirationDays
}
// GetPasswordMinimumLength gets the passwordMinimumLength property value. Indicates the minimum length of the password required on the device. Valid values 4 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordMinimumLength()(*int32) {
    return m.passwordMinimumLength
}
// GetPasswordMinimumLetterCharacters gets the passwordMinimumLetterCharacters property value. Indicates the minimum number of letter characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordMinimumLetterCharacters()(*int32) {
    return m.passwordMinimumLetterCharacters
}
// GetPasswordMinimumLowerCaseCharacters gets the passwordMinimumLowerCaseCharacters property value. Indicates the minimum number of lower case characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordMinimumLowerCaseCharacters()(*int32) {
    return m.passwordMinimumLowerCaseCharacters
}
// GetPasswordMinimumNonLetterCharacters gets the passwordMinimumNonLetterCharacters property value. Indicates the minimum number of non-letter characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordMinimumNonLetterCharacters()(*int32) {
    return m.passwordMinimumNonLetterCharacters
}
// GetPasswordMinimumNumericCharacters gets the passwordMinimumNumericCharacters property value. Indicates the minimum number of numeric characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordMinimumNumericCharacters()(*int32) {
    return m.passwordMinimumNumericCharacters
}
// GetPasswordMinimumSymbolCharacters gets the passwordMinimumSymbolCharacters property value. Indicates the minimum number of symbol characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordMinimumSymbolCharacters()(*int32) {
    return m.passwordMinimumSymbolCharacters
}
// GetPasswordMinimumUpperCaseCharacters gets the passwordMinimumUpperCaseCharacters property value. Indicates the minimum number of upper case letter characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordMinimumUpperCaseCharacters()(*int32) {
    return m.passwordMinimumUpperCaseCharacters
}
// GetPasswordMinutesOfInactivityBeforeScreenTimeout gets the passwordMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before the screen times out.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeout()(*int32) {
    return m.passwordMinutesOfInactivityBeforeScreenTimeout
}
// GetPasswordPreviousPasswordCountToBlock gets the passwordPreviousPasswordCountToBlock property value. Indicates the length of password history, where the user will not be able to enter a new password that is the same as any password in the history. Valid values 0 to 24
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordPreviousPasswordCountToBlock()(*int32) {
    return m.passwordPreviousPasswordCountToBlock
}
// GetPasswordRequiredType gets the passwordRequiredType property value. Indicates the minimum password quality required on the device. Possible values are: deviceDefault, required, numeric, numericComplex, alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, customPassword.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordRequiredType()(*AndroidDeviceOwnerRequiredPasswordType) {
    return m.passwordRequiredType
}
// GetPasswordRequireUnlock gets the passwordRequireUnlock property value. Indicates the timeout period after which a device must be unlocked using a form of strong authentication. Possible values are: deviceDefault, daily, unkownFutureValue.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordRequireUnlock()(*AndroidDeviceOwnerRequiredPasswordUnlock) {
    return m.passwordRequireUnlock
}
// GetPasswordSignInFailureCountBeforeFactoryReset gets the passwordSignInFailureCountBeforeFactoryReset property value. Indicates the number of times a user can enter an incorrect password before the device is wiped. Valid values 4 to 11
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPasswordSignInFailureCountBeforeFactoryReset()(*int32) {
    return m.passwordSignInFailureCountBeforeFactoryReset
}
// GetPersonalProfileAppsAllowInstallFromUnknownSources gets the personalProfileAppsAllowInstallFromUnknownSources property value. Indicates whether the user can install apps from unknown sources on the personal profile.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPersonalProfileAppsAllowInstallFromUnknownSources()(*bool) {
    return m.personalProfileAppsAllowInstallFromUnknownSources
}
// GetPersonalProfileCameraBlocked gets the personalProfileCameraBlocked property value. Indicates whether to disable the use of the camera on the personal profile.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPersonalProfileCameraBlocked()(*bool) {
    return m.personalProfileCameraBlocked
}
// GetPersonalProfilePersonalApplications gets the personalProfilePersonalApplications property value. Policy applied to applications in the personal profile. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPersonalProfilePersonalApplications()([]AppListItemable) {
    return m.personalProfilePersonalApplications
}
// GetPersonalProfilePlayStoreMode gets the personalProfilePlayStoreMode property value. Used together with PersonalProfilePersonalApplications to control how apps in the personal profile are allowed or blocked. Possible values are: notConfigured, blockedApps, allowedApps.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPersonalProfilePlayStoreMode()(*PersonalProfilePersonalPlayStoreMode) {
    return m.personalProfilePlayStoreMode
}
// GetPersonalProfileScreenCaptureBlocked gets the personalProfileScreenCaptureBlocked property value. Indicates whether to disable the capability to take screenshots on the personal profile.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPersonalProfileScreenCaptureBlocked()(*bool) {
    return m.personalProfileScreenCaptureBlocked
}
// GetPlayStoreMode gets the playStoreMode property value. Indicates the Play Store mode of the device. Possible values are: notConfigured, allowList, blockList.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetPlayStoreMode()(*AndroidDeviceOwnerPlayStoreMode) {
    return m.playStoreMode
}
// GetScreenCaptureBlocked gets the screenCaptureBlocked property value. Indicates whether or not to disable the capability to take screenshots.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetScreenCaptureBlocked()(*bool) {
    return m.screenCaptureBlocked
}
// GetSecurityCommonCriteriaModeEnabled gets the securityCommonCriteriaModeEnabled property value. Represents the security common criteria mode enabled provided to users when they attempt to modify managed settings on their device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetSecurityCommonCriteriaModeEnabled()(*bool) {
    return m.securityCommonCriteriaModeEnabled
}
// GetSecurityDeveloperSettingsEnabled gets the securityDeveloperSettingsEnabled property value. Indicates whether or not the user is allowed to access developer settings like developer options and safe boot on the device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetSecurityDeveloperSettingsEnabled()(*bool) {
    return m.securityDeveloperSettingsEnabled
}
// GetSecurityRequireVerifyApps gets the securityRequireVerifyApps property value. Indicates whether or not verify apps is required.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetSecurityRequireVerifyApps()(*bool) {
    return m.securityRequireVerifyApps
}
// GetShortHelpText gets the shortHelpText property value. Represents the customized short help text provided to users when they attempt to modify managed settings on their device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetShortHelpText()(AndroidDeviceOwnerUserFacingMessageable) {
    return m.shortHelpText
}
// GetStatusBarBlocked gets the statusBarBlocked property value. Indicates whether or the status bar is disabled, including notifications, quick settings and other screen overlays.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetStatusBarBlocked()(*bool) {
    return m.statusBarBlocked
}
// GetStayOnModes gets the stayOnModes property value. List of modes in which the device's display will stay powered-on. This collection can contain a maximum of 4 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetStayOnModes()([]AndroidDeviceOwnerBatteryPluggedMode) {
    return m.stayOnModes
}
// GetStorageAllowUsb gets the storageAllowUsb property value. Indicates whether or not to allow USB mass storage.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetStorageAllowUsb()(*bool) {
    return m.storageAllowUsb
}
// GetStorageBlockExternalMedia gets the storageBlockExternalMedia property value. Indicates whether or not to block external media.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetStorageBlockExternalMedia()(*bool) {
    return m.storageBlockExternalMedia
}
// GetStorageBlockUsbFileTransfer gets the storageBlockUsbFileTransfer property value. Indicates whether or not to block USB file transfer.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetStorageBlockUsbFileTransfer()(*bool) {
    return m.storageBlockUsbFileTransfer
}
// GetSystemUpdateFreezePeriods gets the systemUpdateFreezePeriods property value. Indicates the annually repeating time periods during which system updates are postponed. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetSystemUpdateFreezePeriods()([]AndroidDeviceOwnerSystemUpdateFreezePeriodable) {
    return m.systemUpdateFreezePeriods
}
// GetSystemUpdateInstallType gets the systemUpdateInstallType property value. The type of system update configuration. Possible values are: deviceDefault, postpone, windowed, automatic.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetSystemUpdateInstallType()(*AndroidDeviceOwnerSystemUpdateInstallType) {
    return m.systemUpdateInstallType
}
// GetSystemUpdateWindowEndMinutesAfterMidnight gets the systemUpdateWindowEndMinutesAfterMidnight property value. Indicates the number of minutes after midnight that the system update window ends. Valid values 0 to 1440
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetSystemUpdateWindowEndMinutesAfterMidnight()(*int32) {
    return m.systemUpdateWindowEndMinutesAfterMidnight
}
// GetSystemUpdateWindowStartMinutesAfterMidnight gets the systemUpdateWindowStartMinutesAfterMidnight property value. Indicates the number of minutes after midnight that the system update window starts. Valid values 0 to 1440
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetSystemUpdateWindowStartMinutesAfterMidnight()(*int32) {
    return m.systemUpdateWindowStartMinutesAfterMidnight
}
// GetSystemWindowsBlocked gets the systemWindowsBlocked property value. Whether or not to block Android system prompt windows, like toasts, phone activities, and system alerts.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetSystemWindowsBlocked()(*bool) {
    return m.systemWindowsBlocked
}
// GetUsersBlockAdd gets the usersBlockAdd property value. Indicates whether or not adding users and profiles is disabled.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetUsersBlockAdd()(*bool) {
    return m.usersBlockAdd
}
// GetUsersBlockRemove gets the usersBlockRemove property value. Indicates whether or not to disable removing other users from the device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetUsersBlockRemove()(*bool) {
    return m.usersBlockRemove
}
// GetVolumeBlockAdjustment gets the volumeBlockAdjustment property value. Indicates whether or not adjusting the master volume is disabled.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetVolumeBlockAdjustment()(*bool) {
    return m.volumeBlockAdjustment
}
// GetVpnAlwaysOnLockdownMode gets the vpnAlwaysOnLockdownMode property value. If an always on VPN package name is specified, whether or not to lock network traffic when that VPN is disconnected.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetVpnAlwaysOnLockdownMode()(*bool) {
    return m.vpnAlwaysOnLockdownMode
}
// GetVpnAlwaysOnPackageIdentifier gets the vpnAlwaysOnPackageIdentifier property value. Android app package name for app that will handle an always-on VPN connection.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetVpnAlwaysOnPackageIdentifier()(*string) {
    return m.vpnAlwaysOnPackageIdentifier
}
// GetWifiBlockEditConfigurations gets the wifiBlockEditConfigurations property value. Indicates whether or not to block the user from editing the wifi connection settings.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetWifiBlockEditConfigurations()(*bool) {
    return m.wifiBlockEditConfigurations
}
// GetWifiBlockEditPolicyDefinedConfigurations gets the wifiBlockEditPolicyDefinedConfigurations property value. Indicates whether or not to block the user from editing just the networks defined by the policy.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetWifiBlockEditPolicyDefinedConfigurations()(*bool) {
    return m.wifiBlockEditPolicyDefinedConfigurations
}
// GetWorkProfilePasswordExpirationDays gets the workProfilePasswordExpirationDays property value. Indicates the number of days that a work profile password can be set before it expires and a new password will be required. Valid values 1 to 365
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetWorkProfilePasswordExpirationDays()(*int32) {
    return m.workProfilePasswordExpirationDays
}
// GetWorkProfilePasswordMinimumLength gets the workProfilePasswordMinimumLength property value. Indicates the minimum length of the work profile password. Valid values 4 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetWorkProfilePasswordMinimumLength()(*int32) {
    return m.workProfilePasswordMinimumLength
}
// GetWorkProfilePasswordMinimumLetterCharacters gets the workProfilePasswordMinimumLetterCharacters property value. Indicates the minimum number of letter characters required for the work profile password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetWorkProfilePasswordMinimumLetterCharacters()(*int32) {
    return m.workProfilePasswordMinimumLetterCharacters
}
// GetWorkProfilePasswordMinimumLowerCaseCharacters gets the workProfilePasswordMinimumLowerCaseCharacters property value. Indicates the minimum number of lower-case characters required for the work profile password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetWorkProfilePasswordMinimumLowerCaseCharacters()(*int32) {
    return m.workProfilePasswordMinimumLowerCaseCharacters
}
// GetWorkProfilePasswordMinimumNonLetterCharacters gets the workProfilePasswordMinimumNonLetterCharacters property value. Indicates the minimum number of non-letter characters required for the work profile password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetWorkProfilePasswordMinimumNonLetterCharacters()(*int32) {
    return m.workProfilePasswordMinimumNonLetterCharacters
}
// GetWorkProfilePasswordMinimumNumericCharacters gets the workProfilePasswordMinimumNumericCharacters property value. Indicates the minimum number of numeric characters required for the work profile password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetWorkProfilePasswordMinimumNumericCharacters()(*int32) {
    return m.workProfilePasswordMinimumNumericCharacters
}
// GetWorkProfilePasswordMinimumSymbolCharacters gets the workProfilePasswordMinimumSymbolCharacters property value. Indicates the minimum number of symbol characters required for the work profile password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetWorkProfilePasswordMinimumSymbolCharacters()(*int32) {
    return m.workProfilePasswordMinimumSymbolCharacters
}
// GetWorkProfilePasswordMinimumUpperCaseCharacters gets the workProfilePasswordMinimumUpperCaseCharacters property value. Indicates the minimum number of upper-case letter characters required for the work profile password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetWorkProfilePasswordMinimumUpperCaseCharacters()(*int32) {
    return m.workProfilePasswordMinimumUpperCaseCharacters
}
// GetWorkProfilePasswordPreviousPasswordCountToBlock gets the workProfilePasswordPreviousPasswordCountToBlock property value. Indicates the length of the work profile password history, where the user will not be able to enter a new password that is the same as any password in the history. Valid values 0 to 24
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetWorkProfilePasswordPreviousPasswordCountToBlock()(*int32) {
    return m.workProfilePasswordPreviousPasswordCountToBlock
}
// GetWorkProfilePasswordRequiredType gets the workProfilePasswordRequiredType property value. Indicates the minimum password quality required on the work profile password. Possible values are: deviceDefault, required, numeric, numericComplex, alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, customPassword.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetWorkProfilePasswordRequiredType()(*AndroidDeviceOwnerRequiredPasswordType) {
    return m.workProfilePasswordRequiredType
}
// GetWorkProfilePasswordRequireUnlock gets the workProfilePasswordRequireUnlock property value. Indicates the timeout period after which a work profile must be unlocked using a form of strong authentication. Possible values are: deviceDefault, daily, unkownFutureValue.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetWorkProfilePasswordRequireUnlock()(*AndroidDeviceOwnerRequiredPasswordUnlock) {
    return m.workProfilePasswordRequireUnlock
}
// GetWorkProfilePasswordSignInFailureCountBeforeFactoryReset gets the workProfilePasswordSignInFailureCountBeforeFactoryReset property value. Indicates the number of times a user can enter an incorrect work profile password before the device is wiped. Valid values 4 to 11
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) GetWorkProfilePasswordSignInFailureCountBeforeFactoryReset()(*int32) {
    return m.workProfilePasswordSignInFailureCountBeforeFactoryReset
}
// Serialize serializes information the current object
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("accountsBlockModification", m.GetAccountsBlockModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("appsAllowInstallFromUnknownSources", m.GetAppsAllowInstallFromUnknownSources())
        if err != nil {
            return err
        }
    }
    if m.GetAppsAutoUpdatePolicy() != nil {
        cast := (*m.GetAppsAutoUpdatePolicy()).String()
        err = writer.WriteStringValue("appsAutoUpdatePolicy", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppsDefaultPermissionPolicy() != nil {
        cast := (*m.GetAppsDefaultPermissionPolicy()).String()
        err = writer.WriteStringValue("appsDefaultPermissionPolicy", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("appsRecommendSkippingFirstUseHints", m.GetAppsRecommendSkippingFirstUseHints())
        if err != nil {
            return err
        }
    }
    if m.GetAzureAdSharedDeviceDataClearApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAzureAdSharedDeviceDataClearApps()))
        for i, v := range m.GetAzureAdSharedDeviceDataClearApps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("azureAdSharedDeviceDataClearApps", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("bluetoothBlockConfiguration", m.GetBluetoothBlockConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("bluetoothBlockContactSharing", m.GetBluetoothBlockContactSharing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cameraBlocked", m.GetCameraBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cellularBlockWiFiTethering", m.GetCellularBlockWiFiTethering())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("certificateCredentialConfigurationDisabled", m.GetCertificateCredentialConfigurationDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("crossProfilePoliciesAllowCopyPaste", m.GetCrossProfilePoliciesAllowCopyPaste())
        if err != nil {
            return err
        }
    }
    if m.GetCrossProfilePoliciesAllowDataSharing() != nil {
        cast := (*m.GetCrossProfilePoliciesAllowDataSharing()).String()
        err = writer.WriteStringValue("crossProfilePoliciesAllowDataSharing", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("crossProfilePoliciesShowWorkContactsInPersonalProfile", m.GetCrossProfilePoliciesShowWorkContactsInPersonalProfile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("dataRoamingBlocked", m.GetDataRoamingBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("dateTimeConfigurationBlocked", m.GetDateTimeConfigurationBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("detailedHelpText", m.GetDetailedHelpText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceOwnerLockScreenMessage", m.GetDeviceOwnerLockScreenMessage())
        if err != nil {
            return err
        }
    }
    if m.GetEnrollmentProfile() != nil {
        cast := (*m.GetEnrollmentProfile()).String()
        err = writer.WriteStringValue("enrollmentProfile", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("factoryResetBlocked", m.GetFactoryResetBlocked())
        if err != nil {
            return err
        }
    }
    if m.GetFactoryResetDeviceAdministratorEmails() != nil {
        err = writer.WriteCollectionOfStringValues("factoryResetDeviceAdministratorEmails", m.GetFactoryResetDeviceAdministratorEmails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("globalProxy", m.GetGlobalProxy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("googleAccountsBlocked", m.GetGoogleAccountsBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskCustomizationDeviceSettingsBlocked", m.GetKioskCustomizationDeviceSettingsBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskCustomizationPowerButtonActionsBlocked", m.GetKioskCustomizationPowerButtonActionsBlocked())
        if err != nil {
            return err
        }
    }
    if m.GetKioskCustomizationStatusBar() != nil {
        cast := (*m.GetKioskCustomizationStatusBar()).String()
        err = writer.WriteStringValue("kioskCustomizationStatusBar", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskCustomizationSystemErrorWarnings", m.GetKioskCustomizationSystemErrorWarnings())
        if err != nil {
            return err
        }
    }
    if m.GetKioskCustomizationSystemNavigation() != nil {
        cast := (*m.GetKioskCustomizationSystemNavigation()).String()
        err = writer.WriteStringValue("kioskCustomizationSystemNavigation", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeAppOrderEnabled", m.GetKioskModeAppOrderEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetKioskModeAppPositions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetKioskModeAppPositions()))
        for i, v := range m.GetKioskModeAppPositions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("kioskModeAppPositions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetKioskModeApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetKioskModeApps()))
        for i, v := range m.GetKioskModeApps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("kioskModeApps", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeAppsInFolderOrderedByName", m.GetKioskModeAppsInFolderOrderedByName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeBluetoothConfigurationEnabled", m.GetKioskModeBluetoothConfigurationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeDebugMenuEasyAccessEnabled", m.GetKioskModeDebugMenuEasyAccessEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("kioskModeExitCode", m.GetKioskModeExitCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeFlashlightConfigurationEnabled", m.GetKioskModeFlashlightConfigurationEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetKioskModeFolderIcon() != nil {
        cast := (*m.GetKioskModeFolderIcon()).String()
        err = writer.WriteStringValue("kioskModeFolderIcon", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("kioskModeGridHeight", m.GetKioskModeGridHeight())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("kioskModeGridWidth", m.GetKioskModeGridWidth())
        if err != nil {
            return err
        }
    }
    if m.GetKioskModeIconSize() != nil {
        cast := (*m.GetKioskModeIconSize()).String()
        err = writer.WriteStringValue("kioskModeIconSize", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeLockHomeScreen", m.GetKioskModeLockHomeScreen())
        if err != nil {
            return err
        }
    }
    if m.GetKioskModeManagedFolders() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetKioskModeManagedFolders()))
        for i, v := range m.GetKioskModeManagedFolders() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("kioskModeManagedFolders", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeManagedHomeScreenAutoSignout", m.GetKioskModeManagedHomeScreenAutoSignout())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("kioskModeManagedHomeScreenInactiveSignOutDelayInSeconds", m.GetKioskModeManagedHomeScreenInactiveSignOutDelayInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("kioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds", m.GetKioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds())
        if err != nil {
            return err
        }
    }
    if m.GetKioskModeManagedHomeScreenPinComplexity() != nil {
        cast := (*m.GetKioskModeManagedHomeScreenPinComplexity()).String()
        err = writer.WriteStringValue("kioskModeManagedHomeScreenPinComplexity", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeManagedHomeScreenPinRequired", m.GetKioskModeManagedHomeScreenPinRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeManagedHomeScreenPinRequiredToResume", m.GetKioskModeManagedHomeScreenPinRequiredToResume())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("kioskModeManagedHomeScreenSignInBackground", m.GetKioskModeManagedHomeScreenSignInBackground())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("kioskModeManagedHomeScreenSignInBrandingLogo", m.GetKioskModeManagedHomeScreenSignInBrandingLogo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeManagedHomeScreenSignInEnabled", m.GetKioskModeManagedHomeScreenSignInEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeManagedSettingsEntryDisabled", m.GetKioskModeManagedSettingsEntryDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeMediaVolumeConfigurationEnabled", m.GetKioskModeMediaVolumeConfigurationEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetKioskModeScreenOrientation() != nil {
        cast := (*m.GetKioskModeScreenOrientation()).String()
        err = writer.WriteStringValue("kioskModeScreenOrientation", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeScreenSaverConfigurationEnabled", m.GetKioskModeScreenSaverConfigurationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeScreenSaverDetectMediaDisabled", m.GetKioskModeScreenSaverDetectMediaDisabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("kioskModeScreenSaverDisplayTimeInSeconds", m.GetKioskModeScreenSaverDisplayTimeInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("kioskModeScreenSaverImageUrl", m.GetKioskModeScreenSaverImageUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("kioskModeScreenSaverStartDelayInSeconds", m.GetKioskModeScreenSaverStartDelayInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeShowAppNotificationBadge", m.GetKioskModeShowAppNotificationBadge())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeShowDeviceInfo", m.GetKioskModeShowDeviceInfo())
        if err != nil {
            return err
        }
    }
    if m.GetKioskModeUseManagedHomeScreenApp() != nil {
        cast := (*m.GetKioskModeUseManagedHomeScreenApp()).String()
        err = writer.WriteStringValue("kioskModeUseManagedHomeScreenApp", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeVirtualHomeButtonEnabled", m.GetKioskModeVirtualHomeButtonEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetKioskModeVirtualHomeButtonType() != nil {
        cast := (*m.GetKioskModeVirtualHomeButtonType()).String()
        err = writer.WriteStringValue("kioskModeVirtualHomeButtonType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("kioskModeWallpaperUrl", m.GetKioskModeWallpaperUrl())
        if err != nil {
            return err
        }
    }
    if m.GetKioskModeWifiAllowedSsids() != nil {
        err = writer.WriteCollectionOfStringValues("kioskModeWifiAllowedSsids", m.GetKioskModeWifiAllowedSsids())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeWiFiConfigurationEnabled", m.GetKioskModeWiFiConfigurationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("microphoneForceMute", m.GetMicrophoneForceMute())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("microsoftLauncherConfigurationEnabled", m.GetMicrosoftLauncherConfigurationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("microsoftLauncherCustomWallpaperAllowUserModification", m.GetMicrosoftLauncherCustomWallpaperAllowUserModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("microsoftLauncherCustomWallpaperEnabled", m.GetMicrosoftLauncherCustomWallpaperEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("microsoftLauncherCustomWallpaperImageUrl", m.GetMicrosoftLauncherCustomWallpaperImageUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("microsoftLauncherDockPresenceAllowUserModification", m.GetMicrosoftLauncherDockPresenceAllowUserModification())
        if err != nil {
            return err
        }
    }
    if m.GetMicrosoftLauncherDockPresenceConfiguration() != nil {
        cast := (*m.GetMicrosoftLauncherDockPresenceConfiguration()).String()
        err = writer.WriteStringValue("microsoftLauncherDockPresenceConfiguration", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("microsoftLauncherFeedAllowUserModification", m.GetMicrosoftLauncherFeedAllowUserModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("microsoftLauncherFeedEnabled", m.GetMicrosoftLauncherFeedEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetMicrosoftLauncherSearchBarPlacementConfiguration() != nil {
        cast := (*m.GetMicrosoftLauncherSearchBarPlacementConfiguration()).String()
        err = writer.WriteStringValue("microsoftLauncherSearchBarPlacementConfiguration", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("networkEscapeHatchAllowed", m.GetNetworkEscapeHatchAllowed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("nfcBlockOutgoingBeam", m.GetNfcBlockOutgoingBeam())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordBlockKeyguard", m.GetPasswordBlockKeyguard())
        if err != nil {
            return err
        }
    }
    if m.GetPasswordBlockKeyguardFeatures() != nil {
        err = writer.WriteCollectionOfStringValues("passwordBlockKeyguardFeatures", SerializeAndroidKeyguardFeature(m.GetPasswordBlockKeyguardFeatures()))
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordExpirationDays", m.GetPasswordExpirationDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumLength", m.GetPasswordMinimumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumLetterCharacters", m.GetPasswordMinimumLetterCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumLowerCaseCharacters", m.GetPasswordMinimumLowerCaseCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumNonLetterCharacters", m.GetPasswordMinimumNonLetterCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumNumericCharacters", m.GetPasswordMinimumNumericCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumSymbolCharacters", m.GetPasswordMinimumSymbolCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumUpperCaseCharacters", m.GetPasswordMinimumUpperCaseCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinutesOfInactivityBeforeScreenTimeout", m.GetPasswordMinutesOfInactivityBeforeScreenTimeout())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordPreviousPasswordCountToBlock", m.GetPasswordPreviousPasswordCountToBlock())
        if err != nil {
            return err
        }
    }
    if m.GetPasswordRequiredType() != nil {
        cast := (*m.GetPasswordRequiredType()).String()
        err = writer.WriteStringValue("passwordRequiredType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPasswordRequireUnlock() != nil {
        cast := (*m.GetPasswordRequireUnlock()).String()
        err = writer.WriteStringValue("passwordRequireUnlock", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordSignInFailureCountBeforeFactoryReset", m.GetPasswordSignInFailureCountBeforeFactoryReset())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("personalProfileAppsAllowInstallFromUnknownSources", m.GetPersonalProfileAppsAllowInstallFromUnknownSources())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("personalProfileCameraBlocked", m.GetPersonalProfileCameraBlocked())
        if err != nil {
            return err
        }
    }
    if m.GetPersonalProfilePersonalApplications() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPersonalProfilePersonalApplications()))
        for i, v := range m.GetPersonalProfilePersonalApplications() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("personalProfilePersonalApplications", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPersonalProfilePlayStoreMode() != nil {
        cast := (*m.GetPersonalProfilePlayStoreMode()).String()
        err = writer.WriteStringValue("personalProfilePlayStoreMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("personalProfileScreenCaptureBlocked", m.GetPersonalProfileScreenCaptureBlocked())
        if err != nil {
            return err
        }
    }
    if m.GetPlayStoreMode() != nil {
        cast := (*m.GetPlayStoreMode()).String()
        err = writer.WriteStringValue("playStoreMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("screenCaptureBlocked", m.GetScreenCaptureBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityCommonCriteriaModeEnabled", m.GetSecurityCommonCriteriaModeEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityDeveloperSettingsEnabled", m.GetSecurityDeveloperSettingsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityRequireVerifyApps", m.GetSecurityRequireVerifyApps())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("shortHelpText", m.GetShortHelpText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("statusBarBlocked", m.GetStatusBarBlocked())
        if err != nil {
            return err
        }
    }
    if m.GetStayOnModes() != nil {
        err = writer.WriteCollectionOfStringValues("stayOnModes", SerializeAndroidDeviceOwnerBatteryPluggedMode(m.GetStayOnModes()))
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("storageAllowUsb", m.GetStorageAllowUsb())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("storageBlockExternalMedia", m.GetStorageBlockExternalMedia())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("storageBlockUsbFileTransfer", m.GetStorageBlockUsbFileTransfer())
        if err != nil {
            return err
        }
    }
    if m.GetSystemUpdateFreezePeriods() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSystemUpdateFreezePeriods()))
        for i, v := range m.GetSystemUpdateFreezePeriods() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("systemUpdateFreezePeriods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSystemUpdateInstallType() != nil {
        cast := (*m.GetSystemUpdateInstallType()).String()
        err = writer.WriteStringValue("systemUpdateInstallType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("systemUpdateWindowEndMinutesAfterMidnight", m.GetSystemUpdateWindowEndMinutesAfterMidnight())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("systemUpdateWindowStartMinutesAfterMidnight", m.GetSystemUpdateWindowStartMinutesAfterMidnight())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("systemWindowsBlocked", m.GetSystemWindowsBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usersBlockAdd", m.GetUsersBlockAdd())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usersBlockRemove", m.GetUsersBlockRemove())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("volumeBlockAdjustment", m.GetVolumeBlockAdjustment())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("vpnAlwaysOnLockdownMode", m.GetVpnAlwaysOnLockdownMode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("vpnAlwaysOnPackageIdentifier", m.GetVpnAlwaysOnPackageIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("wifiBlockEditConfigurations", m.GetWifiBlockEditConfigurations())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("wifiBlockEditPolicyDefinedConfigurations", m.GetWifiBlockEditPolicyDefinedConfigurations())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordExpirationDays", m.GetWorkProfilePasswordExpirationDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinimumLength", m.GetWorkProfilePasswordMinimumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinimumLetterCharacters", m.GetWorkProfilePasswordMinimumLetterCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinimumLowerCaseCharacters", m.GetWorkProfilePasswordMinimumLowerCaseCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinimumNonLetterCharacters", m.GetWorkProfilePasswordMinimumNonLetterCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinimumNumericCharacters", m.GetWorkProfilePasswordMinimumNumericCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinimumSymbolCharacters", m.GetWorkProfilePasswordMinimumSymbolCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordMinimumUpperCaseCharacters", m.GetWorkProfilePasswordMinimumUpperCaseCharacters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordPreviousPasswordCountToBlock", m.GetWorkProfilePasswordPreviousPasswordCountToBlock())
        if err != nil {
            return err
        }
    }
    if m.GetWorkProfilePasswordRequiredType() != nil {
        cast := (*m.GetWorkProfilePasswordRequiredType()).String()
        err = writer.WriteStringValue("workProfilePasswordRequiredType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetWorkProfilePasswordRequireUnlock() != nil {
        cast := (*m.GetWorkProfilePasswordRequireUnlock()).String()
        err = writer.WriteStringValue("workProfilePasswordRequireUnlock", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("workProfilePasswordSignInFailureCountBeforeFactoryReset", m.GetWorkProfilePasswordSignInFailureCountBeforeFactoryReset())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountsBlockModification sets the accountsBlockModification property value. Indicates whether or not adding or removing accounts is disabled.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetAccountsBlockModification(value *bool)() {
    m.accountsBlockModification = value
}
// SetAppsAllowInstallFromUnknownSources sets the appsAllowInstallFromUnknownSources property value. Indicates whether or not the user is allowed to enable to unknown sources setting.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetAppsAllowInstallFromUnknownSources(value *bool)() {
    m.appsAllowInstallFromUnknownSources = value
}
// SetAppsAutoUpdatePolicy sets the appsAutoUpdatePolicy property value. Indicates the value of the app auto update policy. Possible values are: notConfigured, userChoice, never, wiFiOnly, always.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetAppsAutoUpdatePolicy(value *AndroidDeviceOwnerAppAutoUpdatePolicyType)() {
    m.appsAutoUpdatePolicy = value
}
// SetAppsDefaultPermissionPolicy sets the appsDefaultPermissionPolicy property value. Indicates the permission policy for requests for runtime permissions if one is not defined for the app specifically. Possible values are: deviceDefault, prompt, autoGrant, autoDeny.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetAppsDefaultPermissionPolicy(value *AndroidDeviceOwnerDefaultAppPermissionPolicyType)() {
    m.appsDefaultPermissionPolicy = value
}
// SetAppsRecommendSkippingFirstUseHints sets the appsRecommendSkippingFirstUseHints property value. Whether or not to recommend all apps skip any first-time-use hints they may have added.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetAppsRecommendSkippingFirstUseHints(value *bool)() {
    m.appsRecommendSkippingFirstUseHints = value
}
// SetAzureAdSharedDeviceDataClearApps sets the azureAdSharedDeviceDataClearApps property value. A list of managed apps that will have their data cleared during a global sign-out in AAD shared device mode. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetAzureAdSharedDeviceDataClearApps(value []AppListItemable)() {
    m.azureAdSharedDeviceDataClearApps = value
}
// SetBluetoothBlockConfiguration sets the bluetoothBlockConfiguration property value. Indicates whether or not to block a user from configuring bluetooth.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetBluetoothBlockConfiguration(value *bool)() {
    m.bluetoothBlockConfiguration = value
}
// SetBluetoothBlockContactSharing sets the bluetoothBlockContactSharing property value. Indicates whether or not to block a user from sharing contacts via bluetooth.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetBluetoothBlockContactSharing(value *bool)() {
    m.bluetoothBlockContactSharing = value
}
// SetCameraBlocked sets the cameraBlocked property value. Indicates whether or not to disable the use of the camera.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetCameraBlocked(value *bool)() {
    m.cameraBlocked = value
}
// SetCellularBlockWiFiTethering sets the cellularBlockWiFiTethering property value. Indicates whether or not to block Wi-Fi tethering.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetCellularBlockWiFiTethering(value *bool)() {
    m.cellularBlockWiFiTethering = value
}
// SetCertificateCredentialConfigurationDisabled sets the certificateCredentialConfigurationDisabled property value. Indicates whether or not to block users from any certificate credential configuration.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetCertificateCredentialConfigurationDisabled(value *bool)() {
    m.certificateCredentialConfigurationDisabled = value
}
// SetCrossProfilePoliciesAllowCopyPaste sets the crossProfilePoliciesAllowCopyPaste property value. Indicates whether or not text copied from one profile (personal or work) can be pasted in the other.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetCrossProfilePoliciesAllowCopyPaste(value *bool)() {
    m.crossProfilePoliciesAllowCopyPaste = value
}
// SetCrossProfilePoliciesAllowDataSharing sets the crossProfilePoliciesAllowDataSharing property value. Indicates whether data from one profile (personal or work) can be shared with apps in the other profile. Possible values are: notConfigured, crossProfileDataSharingBlocked, dataSharingFromWorkToPersonalBlocked, crossProfileDataSharingAllowed, unkownFutureValue.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetCrossProfilePoliciesAllowDataSharing(value *AndroidDeviceOwnerCrossProfileDataSharing)() {
    m.crossProfilePoliciesAllowDataSharing = value
}
// SetCrossProfilePoliciesShowWorkContactsInPersonalProfile sets the crossProfilePoliciesShowWorkContactsInPersonalProfile property value. Indicates whether or not contacts stored in work profile are shown in personal profile contact searches/incoming calls.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetCrossProfilePoliciesShowWorkContactsInPersonalProfile(value *bool)() {
    m.crossProfilePoliciesShowWorkContactsInPersonalProfile = value
}
// SetDataRoamingBlocked sets the dataRoamingBlocked property value. Indicates whether or not to block a user from data roaming.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetDataRoamingBlocked(value *bool)() {
    m.dataRoamingBlocked = value
}
// SetDateTimeConfigurationBlocked sets the dateTimeConfigurationBlocked property value. Indicates whether or not to block the user from manually changing the date or time on the device
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetDateTimeConfigurationBlocked(value *bool)() {
    m.dateTimeConfigurationBlocked = value
}
// SetDetailedHelpText sets the detailedHelpText property value. Represents the customized detailed help text provided to users when they attempt to modify managed settings on their device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetDetailedHelpText(value AndroidDeviceOwnerUserFacingMessageable)() {
    m.detailedHelpText = value
}
// SetDeviceOwnerLockScreenMessage sets the deviceOwnerLockScreenMessage property value. Represents the customized lock screen message provided to users when they attempt to modify managed settings on their device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetDeviceOwnerLockScreenMessage(value AndroidDeviceOwnerUserFacingMessageable)() {
    m.deviceOwnerLockScreenMessage = value
}
// SetEnrollmentProfile sets the enrollmentProfile property value. Android Device Owner Enrollment Profile types.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetEnrollmentProfile(value *AndroidDeviceOwnerEnrollmentProfileType)() {
    m.enrollmentProfile = value
}
// SetFactoryResetBlocked sets the factoryResetBlocked property value. Indicates whether or not the factory reset option in settings is disabled.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetFactoryResetBlocked(value *bool)() {
    m.factoryResetBlocked = value
}
// SetFactoryResetDeviceAdministratorEmails sets the factoryResetDeviceAdministratorEmails property value. List of Google account emails that will be required to authenticate after a device is factory reset before it can be set up.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetFactoryResetDeviceAdministratorEmails(value []string)() {
    m.factoryResetDeviceAdministratorEmails = value
}
// SetGlobalProxy sets the globalProxy property value. Proxy is set up directly with host, port and excluded hosts.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetGlobalProxy(value AndroidDeviceOwnerGlobalProxyable)() {
    m.globalProxy = value
}
// SetGoogleAccountsBlocked sets the googleAccountsBlocked property value. Indicates whether or not google accounts will be blocked.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetGoogleAccountsBlocked(value *bool)() {
    m.googleAccountsBlocked = value
}
// SetKioskCustomizationDeviceSettingsBlocked sets the kioskCustomizationDeviceSettingsBlocked property value. Indicates whether a user can access the device's Settings app while in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskCustomizationDeviceSettingsBlocked(value *bool)() {
    m.kioskCustomizationDeviceSettingsBlocked = value
}
// SetKioskCustomizationPowerButtonActionsBlocked sets the kioskCustomizationPowerButtonActionsBlocked property value. Whether the power menu is shown when a user long presses the Power button of a device in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskCustomizationPowerButtonActionsBlocked(value *bool)() {
    m.kioskCustomizationPowerButtonActionsBlocked = value
}
// SetKioskCustomizationStatusBar sets the kioskCustomizationStatusBar property value. Indicates whether system info and notifications are disabled in Kiosk Mode. Possible values are: notConfigured, notificationsAndSystemInfoEnabled, systemInfoOnly.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskCustomizationStatusBar(value *AndroidDeviceOwnerKioskCustomizationStatusBar)() {
    m.kioskCustomizationStatusBar = value
}
// SetKioskCustomizationSystemErrorWarnings sets the kioskCustomizationSystemErrorWarnings property value. Indicates whether system error dialogs for crashed or unresponsive apps are shown in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskCustomizationSystemErrorWarnings(value *bool)() {
    m.kioskCustomizationSystemErrorWarnings = value
}
// SetKioskCustomizationSystemNavigation sets the kioskCustomizationSystemNavigation property value. Indicates which navigation features are enabled in Kiosk Mode. Possible values are: notConfigured, navigationEnabled, homeButtonOnly.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskCustomizationSystemNavigation(value *AndroidDeviceOwnerKioskCustomizationSystemNavigation)() {
    m.kioskCustomizationSystemNavigation = value
}
// SetKioskModeAppOrderEnabled sets the kioskModeAppOrderEnabled property value. Whether or not to enable app ordering in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeAppOrderEnabled(value *bool)() {
    m.kioskModeAppOrderEnabled = value
}
// SetKioskModeAppPositions sets the kioskModeAppPositions property value. The ordering of items on Kiosk Mode Managed Home Screen. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeAppPositions(value []AndroidDeviceOwnerKioskModeAppPositionItemable)() {
    m.kioskModeAppPositions = value
}
// SetKioskModeApps sets the kioskModeApps property value. A list of managed apps that will be shown when the device is in Kiosk Mode. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeApps(value []AppListItemable)() {
    m.kioskModeApps = value
}
// SetKioskModeAppsInFolderOrderedByName sets the kioskModeAppsInFolderOrderedByName property value. Whether or not to alphabetize applications within a folder in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeAppsInFolderOrderedByName(value *bool)() {
    m.kioskModeAppsInFolderOrderedByName = value
}
// SetKioskModeBluetoothConfigurationEnabled sets the kioskModeBluetoothConfigurationEnabled property value. Whether or not to allow a user to configure Bluetooth settings in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeBluetoothConfigurationEnabled(value *bool)() {
    m.kioskModeBluetoothConfigurationEnabled = value
}
// SetKioskModeDebugMenuEasyAccessEnabled sets the kioskModeDebugMenuEasyAccessEnabled property value. Whether or not to allow a user to easy access to the debug menu in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeDebugMenuEasyAccessEnabled(value *bool)() {
    m.kioskModeDebugMenuEasyAccessEnabled = value
}
// SetKioskModeExitCode sets the kioskModeExitCode property value. Exit code to allow a user to escape from Kiosk Mode when the device is in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeExitCode(value *string)() {
    m.kioskModeExitCode = value
}
// SetKioskModeFlashlightConfigurationEnabled sets the kioskModeFlashlightConfigurationEnabled property value. Whether or not to allow a user to use the flashlight in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeFlashlightConfigurationEnabled(value *bool)() {
    m.kioskModeFlashlightConfigurationEnabled = value
}
// SetKioskModeFolderIcon sets the kioskModeFolderIcon property value. Folder icon configuration for managed home screen in Kiosk Mode. Possible values are: notConfigured, darkSquare, darkCircle, lightSquare, lightCircle.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeFolderIcon(value *AndroidDeviceOwnerKioskModeFolderIcon)() {
    m.kioskModeFolderIcon = value
}
// SetKioskModeGridHeight sets the kioskModeGridHeight property value. Number of rows for Managed Home Screen grid with app ordering enabled in Kiosk Mode. Valid values 1 to 9999999
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeGridHeight(value *int32)() {
    m.kioskModeGridHeight = value
}
// SetKioskModeGridWidth sets the kioskModeGridWidth property value. Number of columns for Managed Home Screen grid with app ordering enabled in Kiosk Mode. Valid values 1 to 9999999
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeGridWidth(value *int32)() {
    m.kioskModeGridWidth = value
}
// SetKioskModeIconSize sets the kioskModeIconSize property value. Icon size configuration for managed home screen in Kiosk Mode. Possible values are: notConfigured, smallest, small, regular, large, largest.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeIconSize(value *AndroidDeviceOwnerKioskModeIconSize)() {
    m.kioskModeIconSize = value
}
// SetKioskModeLockHomeScreen sets the kioskModeLockHomeScreen property value. Whether or not to lock home screen to the end user in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeLockHomeScreen(value *bool)() {
    m.kioskModeLockHomeScreen = value
}
// SetKioskModeManagedFolders sets the kioskModeManagedFolders property value. A list of managed folders for a device in Kiosk Mode. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeManagedFolders(value []AndroidDeviceOwnerKioskModeManagedFolderable)() {
    m.kioskModeManagedFolders = value
}
// SetKioskModeManagedHomeScreenAutoSignout sets the kioskModeManagedHomeScreenAutoSignout property value. Whether or not to automatically sign-out of MHS and Shared device mode applications after inactive for Managed Home Screen.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeManagedHomeScreenAutoSignout(value *bool)() {
    m.kioskModeManagedHomeScreenAutoSignout = value
}
// SetKioskModeManagedHomeScreenInactiveSignOutDelayInSeconds sets the kioskModeManagedHomeScreenInactiveSignOutDelayInSeconds property value. Number of seconds to give user notice before automatically signing them out for Managed Home Screen. Valid values 0 to 9999999
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeManagedHomeScreenInactiveSignOutDelayInSeconds(value *int32)() {
    m.kioskModeManagedHomeScreenInactiveSignOutDelayInSeconds = value
}
// SetKioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds sets the kioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds property value. Number of seconds device is inactive before automatically signing user out for Managed Home Screen. Valid values 0 to 9999999
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds(value *int32)() {
    m.kioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds = value
}
// SetKioskModeManagedHomeScreenPinComplexity sets the kioskModeManagedHomeScreenPinComplexity property value. Complexity of PIN for sign-in session for Managed Home Screen. Possible values are: notConfigured, simple, complex.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeManagedHomeScreenPinComplexity(value *KioskModeManagedHomeScreenPinComplexity)() {
    m.kioskModeManagedHomeScreenPinComplexity = value
}
// SetKioskModeManagedHomeScreenPinRequired sets the kioskModeManagedHomeScreenPinRequired property value. Whether or not require user to set a PIN for sign-in session for Managed Home Screen.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeManagedHomeScreenPinRequired(value *bool)() {
    m.kioskModeManagedHomeScreenPinRequired = value
}
// SetKioskModeManagedHomeScreenPinRequiredToResume sets the kioskModeManagedHomeScreenPinRequiredToResume property value. Whether or not required user to enter session PIN if screensaver has appeared for Managed Home Screen.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeManagedHomeScreenPinRequiredToResume(value *bool)() {
    m.kioskModeManagedHomeScreenPinRequiredToResume = value
}
// SetKioskModeManagedHomeScreenSignInBackground sets the kioskModeManagedHomeScreenSignInBackground property value. Custom URL background for sign-in screen for Managed Home Screen.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeManagedHomeScreenSignInBackground(value *string)() {
    m.kioskModeManagedHomeScreenSignInBackground = value
}
// SetKioskModeManagedHomeScreenSignInBrandingLogo sets the kioskModeManagedHomeScreenSignInBrandingLogo property value. Custom URL branding logo for sign-in screen and session pin page for Managed Home Screen.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeManagedHomeScreenSignInBrandingLogo(value *string)() {
    m.kioskModeManagedHomeScreenSignInBrandingLogo = value
}
// SetKioskModeManagedHomeScreenSignInEnabled sets the kioskModeManagedHomeScreenSignInEnabled property value. Whether or not show sign-in screen for Managed Home Screen.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeManagedHomeScreenSignInEnabled(value *bool)() {
    m.kioskModeManagedHomeScreenSignInEnabled = value
}
// SetKioskModeManagedSettingsEntryDisabled sets the kioskModeManagedSettingsEntryDisabled property value. Whether or not to display the Managed Settings entry point on the managed home screen in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeManagedSettingsEntryDisabled(value *bool)() {
    m.kioskModeManagedSettingsEntryDisabled = value
}
// SetKioskModeMediaVolumeConfigurationEnabled sets the kioskModeMediaVolumeConfigurationEnabled property value. Whether or not to allow a user to change the media volume in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeMediaVolumeConfigurationEnabled(value *bool)() {
    m.kioskModeMediaVolumeConfigurationEnabled = value
}
// SetKioskModeScreenOrientation sets the kioskModeScreenOrientation property value. Screen orientation configuration for managed home screen in Kiosk Mode. Possible values are: notConfigured, portrait, landscape, autoRotate.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeScreenOrientation(value *AndroidDeviceOwnerKioskModeScreenOrientation)() {
    m.kioskModeScreenOrientation = value
}
// SetKioskModeScreenSaverConfigurationEnabled sets the kioskModeScreenSaverConfigurationEnabled property value. Whether or not to enable screen saver mode or not in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeScreenSaverConfigurationEnabled(value *bool)() {
    m.kioskModeScreenSaverConfigurationEnabled = value
}
// SetKioskModeScreenSaverDetectMediaDisabled sets the kioskModeScreenSaverDetectMediaDisabled property value. Whether or not the device screen should show the screen saver if audio/video is playing in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeScreenSaverDetectMediaDisabled(value *bool)() {
    m.kioskModeScreenSaverDetectMediaDisabled = value
}
// SetKioskModeScreenSaverDisplayTimeInSeconds sets the kioskModeScreenSaverDisplayTimeInSeconds property value. The number of seconds that the device will display the screen saver for in Kiosk Mode. Valid values 0 to 9999999
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeScreenSaverDisplayTimeInSeconds(value *int32)() {
    m.kioskModeScreenSaverDisplayTimeInSeconds = value
}
// SetKioskModeScreenSaverImageUrl sets the kioskModeScreenSaverImageUrl property value. URL for an image that will be the device's screen saver in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeScreenSaverImageUrl(value *string)() {
    m.kioskModeScreenSaverImageUrl = value
}
// SetKioskModeScreenSaverStartDelayInSeconds sets the kioskModeScreenSaverStartDelayInSeconds property value. The number of seconds the device needs to be inactive for before the screen saver is shown in Kiosk Mode. Valid values 1 to 9999999
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeScreenSaverStartDelayInSeconds(value *int32)() {
    m.kioskModeScreenSaverStartDelayInSeconds = value
}
// SetKioskModeShowAppNotificationBadge sets the kioskModeShowAppNotificationBadge property value. Whether or not to display application notification badges in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeShowAppNotificationBadge(value *bool)() {
    m.kioskModeShowAppNotificationBadge = value
}
// SetKioskModeShowDeviceInfo sets the kioskModeShowDeviceInfo property value. Whether or not to allow a user to access basic device information.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeShowDeviceInfo(value *bool)() {
    m.kioskModeShowDeviceInfo = value
}
// SetKioskModeUseManagedHomeScreenApp sets the kioskModeUseManagedHomeScreenApp property value. Whether or not to use single app kiosk mode or multi-app kiosk mode. Possible values are: notConfigured, singleAppMode, multiAppMode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeUseManagedHomeScreenApp(value *KioskModeType)() {
    m.kioskModeUseManagedHomeScreenApp = value
}
// SetKioskModeVirtualHomeButtonEnabled sets the kioskModeVirtualHomeButtonEnabled property value. Whether or not to display a virtual home button when the device is in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeVirtualHomeButtonEnabled(value *bool)() {
    m.kioskModeVirtualHomeButtonEnabled = value
}
// SetKioskModeVirtualHomeButtonType sets the kioskModeVirtualHomeButtonType property value. Indicates whether the virtual home button is a swipe up home button or a floating home button. Possible values are: notConfigured, swipeUp, floating.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeVirtualHomeButtonType(value *AndroidDeviceOwnerVirtualHomeButtonType)() {
    m.kioskModeVirtualHomeButtonType = value
}
// SetKioskModeWallpaperUrl sets the kioskModeWallpaperUrl property value. URL to a publicly accessible image to use for the wallpaper when the device is in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeWallpaperUrl(value *string)() {
    m.kioskModeWallpaperUrl = value
}
// SetKioskModeWifiAllowedSsids sets the kioskModeWifiAllowedSsids property value. The restricted set of WIFI SSIDs available for the user to configure in Kiosk Mode. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeWifiAllowedSsids(value []string)() {
    m.kioskModeWifiAllowedSsids = value
}
// SetKioskModeWiFiConfigurationEnabled sets the kioskModeWiFiConfigurationEnabled property value. Whether or not to allow a user to configure Wi-Fi settings in Kiosk Mode.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetKioskModeWiFiConfigurationEnabled(value *bool)() {
    m.kioskModeWiFiConfigurationEnabled = value
}
// SetMicrophoneForceMute sets the microphoneForceMute property value. Indicates whether or not to block unmuting the microphone on the device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetMicrophoneForceMute(value *bool)() {
    m.microphoneForceMute = value
}
// SetMicrosoftLauncherConfigurationEnabled sets the microsoftLauncherConfigurationEnabled property value. Indicates whether or not to you want configure Microsoft Launcher.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetMicrosoftLauncherConfigurationEnabled(value *bool)() {
    m.microsoftLauncherConfigurationEnabled = value
}
// SetMicrosoftLauncherCustomWallpaperAllowUserModification sets the microsoftLauncherCustomWallpaperAllowUserModification property value. Indicates whether or not the user can modify the wallpaper to personalize their device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetMicrosoftLauncherCustomWallpaperAllowUserModification(value *bool)() {
    m.microsoftLauncherCustomWallpaperAllowUserModification = value
}
// SetMicrosoftLauncherCustomWallpaperEnabled sets the microsoftLauncherCustomWallpaperEnabled property value. Indicates whether or not to configure the wallpaper on the targeted devices.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetMicrosoftLauncherCustomWallpaperEnabled(value *bool)() {
    m.microsoftLauncherCustomWallpaperEnabled = value
}
// SetMicrosoftLauncherCustomWallpaperImageUrl sets the microsoftLauncherCustomWallpaperImageUrl property value. Indicates the URL for the image file to use as the wallpaper on the targeted devices.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetMicrosoftLauncherCustomWallpaperImageUrl(value *string)() {
    m.microsoftLauncherCustomWallpaperImageUrl = value
}
// SetMicrosoftLauncherDockPresenceAllowUserModification sets the microsoftLauncherDockPresenceAllowUserModification property value. Indicates whether or not the user can modify the device dock configuration on the device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetMicrosoftLauncherDockPresenceAllowUserModification(value *bool)() {
    m.microsoftLauncherDockPresenceAllowUserModification = value
}
// SetMicrosoftLauncherDockPresenceConfiguration sets the microsoftLauncherDockPresenceConfiguration property value. Indicates whether or not you want to configure the device dock. Possible values are: notConfigured, show, hide, disabled.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetMicrosoftLauncherDockPresenceConfiguration(value *MicrosoftLauncherDockPresence)() {
    m.microsoftLauncherDockPresenceConfiguration = value
}
// SetMicrosoftLauncherFeedAllowUserModification sets the microsoftLauncherFeedAllowUserModification property value. Indicates whether or not the user can modify the launcher feed on the device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetMicrosoftLauncherFeedAllowUserModification(value *bool)() {
    m.microsoftLauncherFeedAllowUserModification = value
}
// SetMicrosoftLauncherFeedEnabled sets the microsoftLauncherFeedEnabled property value. Indicates whether or not you want to enable the launcher feed on the device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetMicrosoftLauncherFeedEnabled(value *bool)() {
    m.microsoftLauncherFeedEnabled = value
}
// SetMicrosoftLauncherSearchBarPlacementConfiguration sets the microsoftLauncherSearchBarPlacementConfiguration property value. Indicates the search bar placement configuration on the device. Possible values are: notConfigured, top, bottom, hide.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetMicrosoftLauncherSearchBarPlacementConfiguration(value *MicrosoftLauncherSearchBarPlacement)() {
    m.microsoftLauncherSearchBarPlacementConfiguration = value
}
// SetNetworkEscapeHatchAllowed sets the networkEscapeHatchAllowed property value. Indicates whether or not the device will allow connecting to a temporary network connection at boot time.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetNetworkEscapeHatchAllowed(value *bool)() {
    m.networkEscapeHatchAllowed = value
}
// SetNfcBlockOutgoingBeam sets the nfcBlockOutgoingBeam property value. Indicates whether or not to block NFC outgoing beam.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetNfcBlockOutgoingBeam(value *bool)() {
    m.nfcBlockOutgoingBeam = value
}
// SetPasswordBlockKeyguard sets the passwordBlockKeyguard property value. Indicates whether or not the keyguard is disabled.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordBlockKeyguard(value *bool)() {
    m.passwordBlockKeyguard = value
}
// SetPasswordBlockKeyguardFeatures sets the passwordBlockKeyguardFeatures property value. List of device keyguard features to block. This collection can contain a maximum of 11 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordBlockKeyguardFeatures(value []AndroidKeyguardFeature)() {
    m.passwordBlockKeyguardFeatures = value
}
// SetPasswordExpirationDays sets the passwordExpirationDays property value. Indicates the amount of time that a password can be set for before it expires and a new password will be required. Valid values 1 to 365
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordExpirationDays(value *int32)() {
    m.passwordExpirationDays = value
}
// SetPasswordMinimumLength sets the passwordMinimumLength property value. Indicates the minimum length of the password required on the device. Valid values 4 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordMinimumLength(value *int32)() {
    m.passwordMinimumLength = value
}
// SetPasswordMinimumLetterCharacters sets the passwordMinimumLetterCharacters property value. Indicates the minimum number of letter characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordMinimumLetterCharacters(value *int32)() {
    m.passwordMinimumLetterCharacters = value
}
// SetPasswordMinimumLowerCaseCharacters sets the passwordMinimumLowerCaseCharacters property value. Indicates the minimum number of lower case characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordMinimumLowerCaseCharacters(value *int32)() {
    m.passwordMinimumLowerCaseCharacters = value
}
// SetPasswordMinimumNonLetterCharacters sets the passwordMinimumNonLetterCharacters property value. Indicates the minimum number of non-letter characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordMinimumNonLetterCharacters(value *int32)() {
    m.passwordMinimumNonLetterCharacters = value
}
// SetPasswordMinimumNumericCharacters sets the passwordMinimumNumericCharacters property value. Indicates the minimum number of numeric characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordMinimumNumericCharacters(value *int32)() {
    m.passwordMinimumNumericCharacters = value
}
// SetPasswordMinimumSymbolCharacters sets the passwordMinimumSymbolCharacters property value. Indicates the minimum number of symbol characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordMinimumSymbolCharacters(value *int32)() {
    m.passwordMinimumSymbolCharacters = value
}
// SetPasswordMinimumUpperCaseCharacters sets the passwordMinimumUpperCaseCharacters property value. Indicates the minimum number of upper case letter characters required for device password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordMinimumUpperCaseCharacters(value *int32)() {
    m.passwordMinimumUpperCaseCharacters = value
}
// SetPasswordMinutesOfInactivityBeforeScreenTimeout sets the passwordMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before the screen times out.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeout(value *int32)() {
    m.passwordMinutesOfInactivityBeforeScreenTimeout = value
}
// SetPasswordPreviousPasswordCountToBlock sets the passwordPreviousPasswordCountToBlock property value. Indicates the length of password history, where the user will not be able to enter a new password that is the same as any password in the history. Valid values 0 to 24
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordPreviousPasswordCountToBlock(value *int32)() {
    m.passwordPreviousPasswordCountToBlock = value
}
// SetPasswordRequiredType sets the passwordRequiredType property value. Indicates the minimum password quality required on the device. Possible values are: deviceDefault, required, numeric, numericComplex, alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, customPassword.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordRequiredType(value *AndroidDeviceOwnerRequiredPasswordType)() {
    m.passwordRequiredType = value
}
// SetPasswordRequireUnlock sets the passwordRequireUnlock property value. Indicates the timeout period after which a device must be unlocked using a form of strong authentication. Possible values are: deviceDefault, daily, unkownFutureValue.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordRequireUnlock(value *AndroidDeviceOwnerRequiredPasswordUnlock)() {
    m.passwordRequireUnlock = value
}
// SetPasswordSignInFailureCountBeforeFactoryReset sets the passwordSignInFailureCountBeforeFactoryReset property value. Indicates the number of times a user can enter an incorrect password before the device is wiped. Valid values 4 to 11
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPasswordSignInFailureCountBeforeFactoryReset(value *int32)() {
    m.passwordSignInFailureCountBeforeFactoryReset = value
}
// SetPersonalProfileAppsAllowInstallFromUnknownSources sets the personalProfileAppsAllowInstallFromUnknownSources property value. Indicates whether the user can install apps from unknown sources on the personal profile.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPersonalProfileAppsAllowInstallFromUnknownSources(value *bool)() {
    m.personalProfileAppsAllowInstallFromUnknownSources = value
}
// SetPersonalProfileCameraBlocked sets the personalProfileCameraBlocked property value. Indicates whether to disable the use of the camera on the personal profile.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPersonalProfileCameraBlocked(value *bool)() {
    m.personalProfileCameraBlocked = value
}
// SetPersonalProfilePersonalApplications sets the personalProfilePersonalApplications property value. Policy applied to applications in the personal profile. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPersonalProfilePersonalApplications(value []AppListItemable)() {
    m.personalProfilePersonalApplications = value
}
// SetPersonalProfilePlayStoreMode sets the personalProfilePlayStoreMode property value. Used together with PersonalProfilePersonalApplications to control how apps in the personal profile are allowed or blocked. Possible values are: notConfigured, blockedApps, allowedApps.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPersonalProfilePlayStoreMode(value *PersonalProfilePersonalPlayStoreMode)() {
    m.personalProfilePlayStoreMode = value
}
// SetPersonalProfileScreenCaptureBlocked sets the personalProfileScreenCaptureBlocked property value. Indicates whether to disable the capability to take screenshots on the personal profile.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPersonalProfileScreenCaptureBlocked(value *bool)() {
    m.personalProfileScreenCaptureBlocked = value
}
// SetPlayStoreMode sets the playStoreMode property value. Indicates the Play Store mode of the device. Possible values are: notConfigured, allowList, blockList.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetPlayStoreMode(value *AndroidDeviceOwnerPlayStoreMode)() {
    m.playStoreMode = value
}
// SetScreenCaptureBlocked sets the screenCaptureBlocked property value. Indicates whether or not to disable the capability to take screenshots.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetScreenCaptureBlocked(value *bool)() {
    m.screenCaptureBlocked = value
}
// SetSecurityCommonCriteriaModeEnabled sets the securityCommonCriteriaModeEnabled property value. Represents the security common criteria mode enabled provided to users when they attempt to modify managed settings on their device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetSecurityCommonCriteriaModeEnabled(value *bool)() {
    m.securityCommonCriteriaModeEnabled = value
}
// SetSecurityDeveloperSettingsEnabled sets the securityDeveloperSettingsEnabled property value. Indicates whether or not the user is allowed to access developer settings like developer options and safe boot on the device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetSecurityDeveloperSettingsEnabled(value *bool)() {
    m.securityDeveloperSettingsEnabled = value
}
// SetSecurityRequireVerifyApps sets the securityRequireVerifyApps property value. Indicates whether or not verify apps is required.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetSecurityRequireVerifyApps(value *bool)() {
    m.securityRequireVerifyApps = value
}
// SetShortHelpText sets the shortHelpText property value. Represents the customized short help text provided to users when they attempt to modify managed settings on their device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetShortHelpText(value AndroidDeviceOwnerUserFacingMessageable)() {
    m.shortHelpText = value
}
// SetStatusBarBlocked sets the statusBarBlocked property value. Indicates whether or the status bar is disabled, including notifications, quick settings and other screen overlays.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetStatusBarBlocked(value *bool)() {
    m.statusBarBlocked = value
}
// SetStayOnModes sets the stayOnModes property value. List of modes in which the device's display will stay powered-on. This collection can contain a maximum of 4 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetStayOnModes(value []AndroidDeviceOwnerBatteryPluggedMode)() {
    m.stayOnModes = value
}
// SetStorageAllowUsb sets the storageAllowUsb property value. Indicates whether or not to allow USB mass storage.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetStorageAllowUsb(value *bool)() {
    m.storageAllowUsb = value
}
// SetStorageBlockExternalMedia sets the storageBlockExternalMedia property value. Indicates whether or not to block external media.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetStorageBlockExternalMedia(value *bool)() {
    m.storageBlockExternalMedia = value
}
// SetStorageBlockUsbFileTransfer sets the storageBlockUsbFileTransfer property value. Indicates whether or not to block USB file transfer.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetStorageBlockUsbFileTransfer(value *bool)() {
    m.storageBlockUsbFileTransfer = value
}
// SetSystemUpdateFreezePeriods sets the systemUpdateFreezePeriods property value. Indicates the annually repeating time periods during which system updates are postponed. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetSystemUpdateFreezePeriods(value []AndroidDeviceOwnerSystemUpdateFreezePeriodable)() {
    m.systemUpdateFreezePeriods = value
}
// SetSystemUpdateInstallType sets the systemUpdateInstallType property value. The type of system update configuration. Possible values are: deviceDefault, postpone, windowed, automatic.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetSystemUpdateInstallType(value *AndroidDeviceOwnerSystemUpdateInstallType)() {
    m.systemUpdateInstallType = value
}
// SetSystemUpdateWindowEndMinutesAfterMidnight sets the systemUpdateWindowEndMinutesAfterMidnight property value. Indicates the number of minutes after midnight that the system update window ends. Valid values 0 to 1440
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetSystemUpdateWindowEndMinutesAfterMidnight(value *int32)() {
    m.systemUpdateWindowEndMinutesAfterMidnight = value
}
// SetSystemUpdateWindowStartMinutesAfterMidnight sets the systemUpdateWindowStartMinutesAfterMidnight property value. Indicates the number of minutes after midnight that the system update window starts. Valid values 0 to 1440
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetSystemUpdateWindowStartMinutesAfterMidnight(value *int32)() {
    m.systemUpdateWindowStartMinutesAfterMidnight = value
}
// SetSystemWindowsBlocked sets the systemWindowsBlocked property value. Whether or not to block Android system prompt windows, like toasts, phone activities, and system alerts.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetSystemWindowsBlocked(value *bool)() {
    m.systemWindowsBlocked = value
}
// SetUsersBlockAdd sets the usersBlockAdd property value. Indicates whether or not adding users and profiles is disabled.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetUsersBlockAdd(value *bool)() {
    m.usersBlockAdd = value
}
// SetUsersBlockRemove sets the usersBlockRemove property value. Indicates whether or not to disable removing other users from the device.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetUsersBlockRemove(value *bool)() {
    m.usersBlockRemove = value
}
// SetVolumeBlockAdjustment sets the volumeBlockAdjustment property value. Indicates whether or not adjusting the master volume is disabled.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetVolumeBlockAdjustment(value *bool)() {
    m.volumeBlockAdjustment = value
}
// SetVpnAlwaysOnLockdownMode sets the vpnAlwaysOnLockdownMode property value. If an always on VPN package name is specified, whether or not to lock network traffic when that VPN is disconnected.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetVpnAlwaysOnLockdownMode(value *bool)() {
    m.vpnAlwaysOnLockdownMode = value
}
// SetVpnAlwaysOnPackageIdentifier sets the vpnAlwaysOnPackageIdentifier property value. Android app package name for app that will handle an always-on VPN connection.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetVpnAlwaysOnPackageIdentifier(value *string)() {
    m.vpnAlwaysOnPackageIdentifier = value
}
// SetWifiBlockEditConfigurations sets the wifiBlockEditConfigurations property value. Indicates whether or not to block the user from editing the wifi connection settings.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetWifiBlockEditConfigurations(value *bool)() {
    m.wifiBlockEditConfigurations = value
}
// SetWifiBlockEditPolicyDefinedConfigurations sets the wifiBlockEditPolicyDefinedConfigurations property value. Indicates whether or not to block the user from editing just the networks defined by the policy.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetWifiBlockEditPolicyDefinedConfigurations(value *bool)() {
    m.wifiBlockEditPolicyDefinedConfigurations = value
}
// SetWorkProfilePasswordExpirationDays sets the workProfilePasswordExpirationDays property value. Indicates the number of days that a work profile password can be set before it expires and a new password will be required. Valid values 1 to 365
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetWorkProfilePasswordExpirationDays(value *int32)() {
    m.workProfilePasswordExpirationDays = value
}
// SetWorkProfilePasswordMinimumLength sets the workProfilePasswordMinimumLength property value. Indicates the minimum length of the work profile password. Valid values 4 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetWorkProfilePasswordMinimumLength(value *int32)() {
    m.workProfilePasswordMinimumLength = value
}
// SetWorkProfilePasswordMinimumLetterCharacters sets the workProfilePasswordMinimumLetterCharacters property value. Indicates the minimum number of letter characters required for the work profile password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetWorkProfilePasswordMinimumLetterCharacters(value *int32)() {
    m.workProfilePasswordMinimumLetterCharacters = value
}
// SetWorkProfilePasswordMinimumLowerCaseCharacters sets the workProfilePasswordMinimumLowerCaseCharacters property value. Indicates the minimum number of lower-case characters required for the work profile password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetWorkProfilePasswordMinimumLowerCaseCharacters(value *int32)() {
    m.workProfilePasswordMinimumLowerCaseCharacters = value
}
// SetWorkProfilePasswordMinimumNonLetterCharacters sets the workProfilePasswordMinimumNonLetterCharacters property value. Indicates the minimum number of non-letter characters required for the work profile password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetWorkProfilePasswordMinimumNonLetterCharacters(value *int32)() {
    m.workProfilePasswordMinimumNonLetterCharacters = value
}
// SetWorkProfilePasswordMinimumNumericCharacters sets the workProfilePasswordMinimumNumericCharacters property value. Indicates the minimum number of numeric characters required for the work profile password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetWorkProfilePasswordMinimumNumericCharacters(value *int32)() {
    m.workProfilePasswordMinimumNumericCharacters = value
}
// SetWorkProfilePasswordMinimumSymbolCharacters sets the workProfilePasswordMinimumSymbolCharacters property value. Indicates the minimum number of symbol characters required for the work profile password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetWorkProfilePasswordMinimumSymbolCharacters(value *int32)() {
    m.workProfilePasswordMinimumSymbolCharacters = value
}
// SetWorkProfilePasswordMinimumUpperCaseCharacters sets the workProfilePasswordMinimumUpperCaseCharacters property value. Indicates the minimum number of upper-case letter characters required for the work profile password. Valid values 1 to 16
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetWorkProfilePasswordMinimumUpperCaseCharacters(value *int32)() {
    m.workProfilePasswordMinimumUpperCaseCharacters = value
}
// SetWorkProfilePasswordPreviousPasswordCountToBlock sets the workProfilePasswordPreviousPasswordCountToBlock property value. Indicates the length of the work profile password history, where the user will not be able to enter a new password that is the same as any password in the history. Valid values 0 to 24
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetWorkProfilePasswordPreviousPasswordCountToBlock(value *int32)() {
    m.workProfilePasswordPreviousPasswordCountToBlock = value
}
// SetWorkProfilePasswordRequiredType sets the workProfilePasswordRequiredType property value. Indicates the minimum password quality required on the work profile password. Possible values are: deviceDefault, required, numeric, numericComplex, alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, customPassword.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetWorkProfilePasswordRequiredType(value *AndroidDeviceOwnerRequiredPasswordType)() {
    m.workProfilePasswordRequiredType = value
}
// SetWorkProfilePasswordRequireUnlock sets the workProfilePasswordRequireUnlock property value. Indicates the timeout period after which a work profile must be unlocked using a form of strong authentication. Possible values are: deviceDefault, daily, unkownFutureValue.
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetWorkProfilePasswordRequireUnlock(value *AndroidDeviceOwnerRequiredPasswordUnlock)() {
    m.workProfilePasswordRequireUnlock = value
}
// SetWorkProfilePasswordSignInFailureCountBeforeFactoryReset sets the workProfilePasswordSignInFailureCountBeforeFactoryReset property value. Indicates the number of times a user can enter an incorrect work profile password before the device is wiped. Valid values 4 to 11
func (m *AndroidDeviceOwnerGeneralDeviceConfiguration) SetWorkProfilePasswordSignInFailureCountBeforeFactoryReset(value *int32)() {
    m.workProfilePasswordSignInFailureCountBeforeFactoryReset = value
}
