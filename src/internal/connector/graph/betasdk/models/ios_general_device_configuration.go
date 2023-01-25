package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosGeneralDeviceConfiguration 
type IosGeneralDeviceConfiguration struct {
    DeviceConfiguration
    // Indicates whether or not to allow account modification when the device is in supervised mode.
    accountBlockModification *bool
    // Indicates whether or not to allow activation lock when the device is in the supervised mode.
    activationLockAllowWhenSupervised *bool
    // Indicates whether or not to allow AirDrop when the device is in supervised mode.
    airDropBlocked *bool
    // Indicates whether or not to cause AirDrop to be considered an unmanaged drop target (iOS 9.0 and later).
    airDropForceUnmanagedDropTarget *bool
    // Indicates whether or not to enforce all devices receiving AirPlay requests from this device to use a pairing password.
    airPlayForcePairingPasswordForOutgoingRequests *bool
    // Indicates whether or not keychain storage of username and password for Airprint is blocked (iOS 11.0 and later).
    airPrintBlockCredentialsStorage *bool
    // Indicates whether or not AirPrint is blocked (iOS 11.0 and later).
    airPrintBlocked *bool
    // Indicates whether or not iBeacon discovery of AirPrint printers is blocked. This prevents spurious AirPrint Bluetooth beacons from phishing for network traffic (iOS 11.0 and later).
    airPrintBlockiBeaconDiscovery *bool
    // Indicates if trusted certificates are required for TLS printing communication (iOS 11.0 and later).
    airPrintForceTrustedTLS *bool
    // Prevents a user from adding any App Clips and removes any existing App Clips on the device.
    appClipsBlocked *bool
    // Indicates whether or not to block the user from using News when the device is in supervised mode (iOS 9.0 and later).
    appleNewsBlocked *bool
    // Limits Apple personalized advertising when true. Available in iOS 14 and later.
    applePersonalizedAdsBlocked *bool
    // Indicates whether or not to allow Apple Watch pairing when the device is in supervised mode (iOS 9.0 and later).
    appleWatchBlockPairing *bool
    // Indicates whether or not to force a paired Apple Watch to use Wrist Detection (iOS 8.2 and later).
    appleWatchForceWristDetection *bool
    // Indicates if the removal of apps is allowed.
    appRemovalBlocked *bool
    // Gets or sets the list of iOS apps allowed to autonomously enter Single App Mode. Supervised only. iOS 7.0 and later. This collection can contain a maximum of 500 elements.
    appsSingleAppModeList []AppListItemable
    // Indicates whether or not to block the automatic downloading of apps purchased on other devices when the device is in supervised mode (iOS 9.0 and later).
    appStoreBlockAutomaticDownloads *bool
    // Indicates whether or not to block the user from using the App Store. Requires a supervised device for iOS 13 and later.
    appStoreBlocked *bool
    // Indicates whether or not to block the user from making in app purchases.
    appStoreBlockInAppPurchases *bool
    // Indicates whether or not to block the App Store app, not restricting installation through Host apps. Applies to supervised mode only (iOS 9.0 and later).
    appStoreBlockUIAppInstallation *bool
    // Indicates whether or not to require a password when using the app store.
    appStoreRequirePassword *bool
    // List of apps in the visibility list (either visible/launchable apps list or hidden/unlaunchable apps list, controlled by AppsVisibilityListType) (iOS 9.3 and later). This collection can contain a maximum of 10000 elements.
    appsVisibilityList []AppListItemable
    // Possible values of the compliance app list.
    appsVisibilityListType *AppListType
    // Indicates whether or not to force user authentication before autofilling passwords and credit card information in Safari and other apps on supervised devices.
    autoFillForceAuthentication *bool
    // Blocks users from unlocking their device with Apple Watch. Available for devices running iOS and iPadOS versions 14.5 and later.
    autoUnlockBlocked *bool
    // Indicates whether or not the removal of system apps from the device is blocked on a supervised device (iOS 11.0 and later).
    blockSystemAppRemoval *bool
    // Indicates whether or not to allow modification of Bluetooth settings when the device is in supervised mode (iOS 10.0 and later).
    bluetoothBlockModification *bool
    // Indicates whether or not to block the user from accessing the camera of the device. Requires a supervised device for iOS 13 and later.
    cameraBlocked *bool
    // Indicates whether or not to block data roaming.
    cellularBlockDataRoaming *bool
    // Indicates whether or not to block global background fetch while roaming.
    cellularBlockGlobalBackgroundFetchWhileRoaming *bool
    // Indicates whether or not to allow changes to cellular app data usage settings when the device is in supervised mode.
    cellularBlockPerAppDataModification *bool
    // Indicates whether or not to block Personal Hotspot.
    cellularBlockPersonalHotspot *bool
    // Indicates whether or not to block the user from modifying the personal hotspot setting (iOS 12.2 or later).
    cellularBlockPersonalHotspotModification *bool
    // Indicates whether or not to allow users to change the settings of the cellular plan on a supervised device.
    cellularBlockPlanModification *bool
    // Indicates whether or not to block voice roaming.
    cellularBlockVoiceRoaming *bool
    // Indicates whether or not to block untrusted TLS certificates.
    certificatesBlockUntrustedTlsCertificates *bool
    // Indicates whether or not to allow remote screen observation by Classroom app when the device is in supervised mode (iOS 9.3 and later).
    classroomAppBlockRemoteScreenObservation *bool
    // Indicates whether or not to automatically give permission to the teacher of a managed course on the Classroom app to view a student's screen without prompting when the device is in supervised mode.
    classroomAppForceUnpromptedScreenObservation *bool
    // Indicates whether or not to automatically give permission to the teacher's requests, without prompting the student, when the device is in supervised mode.
    classroomForceAutomaticallyJoinClasses *bool
    // Indicates whether a student enrolled in an unmanaged course via Classroom will request permission from the teacher when attempting to leave the course (iOS 11.3 and later).
    classroomForceRequestPermissionToLeaveClasses *bool
    // Indicates whether or not to allow the teacher to lock apps or the device without prompting the student. Supervised only.
    classroomForceUnpromptedAppAndDeviceLock *bool
    // Possible values of the compliance app list.
    compliantAppListType *AppListType
    // List of apps in the compliance (either allow list or block list, controlled by CompliantAppListType). This collection can contain a maximum of 10000 elements.
    compliantAppsList []AppListItemable
    // Indicates whether or not to block the user from installing configuration profiles and certificates interactively when the device is in supervised mode.
    configurationProfileBlockChanges *bool
    // Indicates whether or not managed apps can write contacts to unmanaged contacts accounts (iOS 12.0 and later).
    contactsAllowManagedToUnmanagedWrite *bool
    // Indicates whether or not unmanaged apps can read from managed contacts accounts (iOS 12.0 or later).
    contactsAllowUnmanagedToManagedRead *bool
    // Indicates whether or not to block the continuous path keyboard when the device is supervised (iOS 13 or later).
    continuousPathKeyboardBlocked *bool
    // Indicates whether or not the Date and Time 'Set Automatically' feature is enabled and cannot be turned off by the user (iOS 12.0 and later).
    dateAndTimeForceSetAutomatically *bool
    // Indicates whether or not to block definition lookup when the device is in supervised mode (iOS 8.1.3 and later ).
    definitionLookupBlocked *bool
    // Indicates whether or not to allow the user to enables restrictions in the device settings when the device is in supervised mode.
    deviceBlockEnableRestrictions *bool
    // Indicates whether or not to allow the use of the 'Erase all content and settings' option on the device when the device is in supervised mode.
    deviceBlockEraseContentAndSettings *bool
    // Indicates whether or not to allow device name modification when the device is in supervised mode (iOS 9.0 and later).
    deviceBlockNameModification *bool
    // Indicates whether or not to block diagnostic data submission.
    diagnosticDataBlockSubmission *bool
    // Indicates whether or not to allow diagnostics submission settings modification when the device is in supervised mode (iOS 9.3.2 and later).
    diagnosticDataBlockSubmissionModification *bool
    // Indicates whether or not to block the user from viewing managed documents in unmanaged apps.
    documentsBlockManagedDocumentsInUnmanagedApps *bool
    // Indicates whether or not to block the user from viewing unmanaged documents in managed apps.
    documentsBlockUnmanagedDocumentsInManagedApps *bool
    // An email address lacking a suffix that matches any of these strings will be considered out-of-domain.
    emailInDomainSuffixes []string
    // Indicates whether or not to block the user from trusting an enterprise app.
    enterpriseAppBlockTrust *bool
    // [Deprecated] Configuring this setting and setting the value to 'true' has no effect on the device.
    enterpriseAppBlockTrustModification *bool
    // Indicates whether or not Enterprise book back up is blocked.
    enterpriseBookBlockBackup *bool
    // Indicates whether or not Enterprise book notes and highlights sync is blocked.
    enterpriseBookBlockMetadataSync *bool
    // Indicates whether or not to allow the addition or removal of cellular plans on the eSIM of a supervised device.
    esimBlockModification *bool
    // Indicates whether or not to block the user from using FaceTime. Requires a supervised device for iOS 13 and later.
    faceTimeBlocked *bool
    // Indicates if devices can access files or other resources on a network server using the Server Message Block (SMB) protocol. Available for devices running iOS and iPadOS, versions 13.0 and later.
    filesNetworkDriveAccessBlocked *bool
    // Indicates if sevices with access can connect to and open files on a USB drive. Available for devices running iOS and iPadOS, versions 13.0 and later.
    filesUsbDriveAccessBlocked *bool
    // Indicates whether or not to block Find My Device when the device is supervised (iOS 13 or later).
    findMyDeviceInFindMyAppBlocked *bool
    // Indicates whether or not to block changes to Find My Friends when the device is in supervised mode.
    findMyFriendsBlocked *bool
    // Indicates whether or not to block Find My Friends when the device is supervised (iOS 13 or later).
    findMyFriendsInFindMyAppBlocked *bool
    // Indicates whether or not to block the user from using Game Center when the device is in supervised mode.
    gameCenterBlocked *bool
    // Indicates whether or not to block the user from having friends in Game Center. Requires a supervised device for iOS 13 and later.
    gamingBlockGameCenterFriends *bool
    // Indicates whether or not to block the user from using multiplayer gaming. Requires a supervised device for iOS 13 and later.
    gamingBlockMultiplayer *bool
    // indicates whether or not to allow host pairing to control the devices an iOS device can pair with when the iOS device is in supervised mode.
    hostPairingBlocked *bool
    // Indicates whether or not to block the user from using the iBooks Store when the device is in supervised mode.
    iBooksStoreBlocked *bool
    // Indicates whether or not to block the user from downloading media from the iBookstore that has been tagged as erotica.
    iBooksStoreBlockErotica *bool
    // Indicates whether or not to block the user from continuing work they started on iOS device to another iOS or macOS device.
    iCloudBlockActivityContinuation *bool
    // Indicates whether or not to block iCloud backup. Requires a supervised device for iOS 13 and later.
    iCloudBlockBackup *bool
    // Indicates whether or not to block iCloud document sync. Requires a supervised device for iOS 13 and later.
    iCloudBlockDocumentSync *bool
    // Indicates whether or not to block Managed Apps Cloud Sync.
    iCloudBlockManagedAppsSync *bool
    // Indicates whether or not to block iCloud Photo Library.
    iCloudBlockPhotoLibrary *bool
    // Indicates whether or not to block iCloud Photo Stream Sync.
    iCloudBlockPhotoStreamSync *bool
    // Indicates whether or not to block Shared Photo Stream.
    iCloudBlockSharedPhotoStream *bool
    // iCloud private relay is an iCloud+ service that prevents networks and servers from monitoring a person's activity across the internet. By blocking iCloud private relay, Apple will not encrypt the traffic leaving the device. Available for devices running iOS 15 and later.
    iCloudPrivateRelayBlocked *bool
    // Indicates whether or not to require backups to iCloud be encrypted.
    iCloudRequireEncryptedBackup *bool
    // Indicates whether or not to block the iTunes app. Requires a supervised device for iOS 13 and later.
    iTunesBlocked *bool
    // Indicates whether or not to block the user from accessing explicit content in iTunes and the App Store. Requires a supervised device for iOS 13 and later.
    iTunesBlockExplicitContent *bool
    // Indicates whether or not to block Music service and revert Music app to classic mode when the device is in supervised mode (iOS 9.3 and later and macOS 10.12 and later).
    iTunesBlockMusicService *bool
    // Indicates whether or not to block the user from using iTunes Radio when the device is in supervised mode (iOS 9.3 and later).
    iTunesBlockRadio *bool
    // Indicates whether or not to block keyboard auto-correction when the device is in supervised mode (iOS 8.1.3 and later).
    keyboardBlockAutoCorrect *bool
    // Indicates whether or not to block the user from using dictation input when the device is in supervised mode.
    keyboardBlockDictation *bool
    // Indicates whether or not to block predictive keyboards when device is in supervised mode (iOS 8.1.3 and later).
    keyboardBlockPredictive *bool
    // Indicates whether or not to block keyboard shortcuts when the device is in supervised mode (iOS 9.0 and later).
    keyboardBlockShortcuts *bool
    // Indicates whether or not to block keyboard spell-checking when the device is in supervised mode (iOS 8.1.3 and later).
    keyboardBlockSpellCheck *bool
    // Indicates whether or not iCloud keychain synchronization is blocked. Requires a supervised device for iOS 13 and later.
    keychainBlockCloudSync *bool
    // Indicates whether or not to allow assistive speak while in kiosk mode.
    kioskModeAllowAssistiveSpeak *bool
    // Indicates whether or not to allow access to the Assistive Touch Settings while in kiosk mode.
    kioskModeAllowAssistiveTouchSettings *bool
    // Indicates whether or not to allow device auto lock while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockAutoLock instead.
    kioskModeAllowAutoLock *bool
    // Indicates whether or not to allow access to the Color Inversion Settings while in kiosk mode.
    kioskModeAllowColorInversionSettings *bool
    // Indicates whether or not to allow use of the ringer switch while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockRingerSwitch instead.
    kioskModeAllowRingerSwitch *bool
    // Indicates whether or not to allow screen rotation while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockScreenRotation instead.
    kioskModeAllowScreenRotation *bool
    // Indicates whether or not to allow use of the sleep button while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockSleepButton instead.
    kioskModeAllowSleepButton *bool
    // Indicates whether or not to allow use of the touchscreen while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockTouchscreen instead.
    kioskModeAllowTouchscreen *bool
    // Indicates whether or not to allow the user to toggle voice control in kiosk mode.
    kioskModeAllowVoiceControlModification *bool
    // Indicates whether or not to allow access to the voice over settings while in kiosk mode.
    kioskModeAllowVoiceOverSettings *bool
    // Indicates whether or not to allow use of the volume buttons while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockVolumeButtons instead.
    kioskModeAllowVolumeButtons *bool
    // Indicates whether or not to allow access to the zoom settings while in kiosk mode.
    kioskModeAllowZoomSettings *bool
    // URL in the app store to the app to use for kiosk mode. Use if KioskModeManagedAppId is not known.
    kioskModeAppStoreUrl *string
    // App source options for iOS kiosk mode.
    kioskModeAppType *IosKioskModeAppType
    // Indicates whether or not to block device auto lock while in kiosk mode.
    kioskModeBlockAutoLock *bool
    // Indicates whether or not to block use of the ringer switch while in kiosk mode.
    kioskModeBlockRingerSwitch *bool
    // Indicates whether or not to block screen rotation while in kiosk mode.
    kioskModeBlockScreenRotation *bool
    // Indicates whether or not to block use of the sleep button while in kiosk mode.
    kioskModeBlockSleepButton *bool
    // Indicates whether or not to block use of the touchscreen while in kiosk mode.
    kioskModeBlockTouchscreen *bool
    // Indicates whether or not to block the volume buttons while in Kiosk Mode.
    kioskModeBlockVolumeButtons *bool
    // ID for built-in apps to use for kiosk mode. Used when KioskModeManagedAppId and KioskModeAppStoreUrl are not set.
    kioskModeBuiltInAppId *string
    // Indicates whether or not to enable voice control in kiosk mode.
    kioskModeEnableVoiceControl *bool
    // Managed app id of the app to use for kiosk mode. If KioskModeManagedAppId is specified then KioskModeAppStoreUrl will be ignored.
    kioskModeManagedAppId *string
    // Indicates whether or not to require assistive touch while in kiosk mode.
    kioskModeRequireAssistiveTouch *bool
    // Indicates whether or not to require color inversion while in kiosk mode.
    kioskModeRequireColorInversion *bool
    // Indicates whether or not to require mono audio while in kiosk mode.
    kioskModeRequireMonoAudio *bool
    // Indicates whether or not to require voice over while in kiosk mode.
    kioskModeRequireVoiceOver *bool
    // Indicates whether or not to require zoom while in kiosk mode.
    kioskModeRequireZoom *bool
    // Indicates whether or not to block the user from using control center on the lock screen.
    lockScreenBlockControlCenter *bool
    // Indicates whether or not to block the user from using the notification view on the lock screen.
    lockScreenBlockNotificationView *bool
    // Indicates whether or not to block the user from using passbook when the device is locked.
    lockScreenBlockPassbook *bool
    // Indicates whether or not to block the user from using the Today View on the lock screen.
    lockScreenBlockTodayView *bool
    // Open-in management controls how people share data between unmanaged and managed apps. Setting this to true enforces copy/paste restrictions based on how you configured Block viewing corporate documents in unmanaged apps  and  Block viewing non-corporate documents in corporate apps.
    managedPasteboardRequired *bool
    // Apps rating as in media content
    mediaContentRatingApps *RatingAppsType
    // Media content rating settings for Australia
    mediaContentRatingAustralia MediaContentRatingAustraliaable
    // Media content rating settings for Canada
    mediaContentRatingCanada MediaContentRatingCanadaable
    // Media content rating settings for France
    mediaContentRatingFrance MediaContentRatingFranceable
    // Media content rating settings for Germany
    mediaContentRatingGermany MediaContentRatingGermanyable
    // Media content rating settings for Ireland
    mediaContentRatingIreland MediaContentRatingIrelandable
    // Media content rating settings for Japan
    mediaContentRatingJapan MediaContentRatingJapanable
    // Media content rating settings for New Zealand
    mediaContentRatingNewZealand MediaContentRatingNewZealandable
    // Media content rating settings for United Kingdom
    mediaContentRatingUnitedKingdom MediaContentRatingUnitedKingdomable
    // Media content rating settings for United States
    mediaContentRatingUnitedStates MediaContentRatingUnitedStatesable
    // Indicates whether or not to block the user from using the Messages app on the supervised device.
    messagesBlocked *bool
    // List of managed apps and the network rules that applies to them. This collection can contain a maximum of 1000 elements.
    networkUsageRules []IosNetworkUsageRuleable
    // Disable NFC to prevent devices from pairing with other NFC-enabled devices. Available for iOS/iPadOS devices running 14.2 and later.
    nfcBlocked *bool
    // Indicates whether or not to allow notifications settings modification (iOS 9.3 and later).
    notificationsBlockSettingsModification *bool
    // Disables connections to Siri servers so that users can’t use Siri to dictate text. Available for devices running iOS and iPadOS versions 14.5 and later.
    onDeviceOnlyDictationForced *bool
    // When set to TRUE, the setting disables connections to Siri servers so that users can’t use Siri to translate text. When set to FALSE, the setting allows connections to to Siri servers to users can use Siri to translate text. Available for devices running iOS and iPadOS versions 15.0 and later.
    onDeviceOnlyTranslationForced *bool
    // Block modification of registered Touch ID fingerprints when in supervised mode.
    passcodeBlockFingerprintModification *bool
    // Indicates whether or not to block fingerprint unlock.
    passcodeBlockFingerprintUnlock *bool
    // Indicates whether or not to allow passcode modification on the supervised device (iOS 9.0 and later).
    passcodeBlockModification *bool
    // Indicates whether or not to block simple passcodes.
    passcodeBlockSimple *bool
    // Number of days before the passcode expires. Valid values 1 to 65535
    passcodeExpirationDays *int32
    // Number of character sets a passcode must contain. Valid values 0 to 4
    passcodeMinimumCharacterSetCount *int32
    // Minimum length of passcode. Valid values 4 to 14
    passcodeMinimumLength *int32
    // Minutes of inactivity before a passcode is required.
    passcodeMinutesOfInactivityBeforeLock *int32
    // Minutes of inactivity before the screen times out.
    passcodeMinutesOfInactivityBeforeScreenTimeout *int32
    // Number of previous passcodes to block. Valid values 1 to 24
    passcodePreviousPasscodeBlockCount *int32
    // Indicates whether or not to require a passcode.
    passcodeRequired *bool
    // Possible values of required passwords.
    passcodeRequiredType *RequiredPasswordType
    // Number of sign in failures allowed before wiping the device. Valid values 2 to 11
    passcodeSignInFailureCountBeforeWipe *int32
    // Indicates whether or not to block sharing passwords with the AirDrop passwords feature iOS 12.0 and later).
    passwordBlockAirDropSharing *bool
    // Indicates if the AutoFill passwords feature is allowed (iOS 12.0 and later).
    passwordBlockAutoFill *bool
    // Indicates whether or not to block requesting passwords from nearby devices (iOS 12.0 and later).
    passwordBlockProximityRequests *bool
    // Indicates whether or not over-the-air PKI updates are blocked. Setting this restriction to false does not disable CRL and OCSP checks (iOS 7.0 and later).
    pkiBlockOTAUpdates *bool
    // Indicates whether or not to block the user from using podcasts on the supervised device (iOS 8.0 and later).
    podcastsBlocked *bool
    // Indicates if ad tracking is limited.(iOS 7.0 and later).
    privacyForceLimitAdTracking *bool
    // Indicates whether or not to enable the prompt to setup nearby devices with a supervised device.
    proximityBlockSetupToNewDevice *bool
    // Indicates whether or not to block the user from using Auto fill in Safari. Requires a supervised device for iOS 13 and later.
    safariBlockAutofill *bool
    // Indicates whether or not to block the user from using Safari. Requires a supervised device for iOS 13 and later.
    safariBlocked *bool
    // Indicates whether or not to block JavaScript in Safari.
    safariBlockJavaScript *bool
    // Indicates whether or not to block popups in Safari.
    safariBlockPopups *bool
    // Web Browser Cookie Settings.
    safariCookieSettings *WebBrowserCookieSettings
    // URLs matching the patterns listed here will be considered managed.
    safariManagedDomains []string
    // Users can save passwords in Safari only from URLs matching the patterns listed here. Applies to devices in supervised mode (iOS 9.3 and later).
    safariPasswordAutoFillDomains []string
    // Indicates whether or not to require fraud warning in Safari.
    safariRequireFraudWarning *bool
    // Indicates whether or not to block the user from taking Screenshots.
    screenCaptureBlocked *bool
    // Indicates whether or not to block temporary sessions on Shared iPads (iOS 13.4 or later).
    sharedDeviceBlockTemporarySessions *bool
    // Indicates whether or not to block the user from using Siri.
    siriBlocked *bool
    // Indicates whether or not to block the user from using Siri when locked.
    siriBlockedWhenLocked *bool
    // Indicates whether or not to block Siri from querying user-generated content when used on a supervised device.
    siriBlockUserGeneratedContent *bool
    // Indicates whether or not to prevent Siri from dictating, or speaking profane language on supervised device.
    siriRequireProfanityFilter *bool
    // Sets how many days a software update will be delyed for a supervised device. Valid values 0 to 90
    softwareUpdatesEnforcedDelayInDays *int32
    // Indicates whether or not to delay user visibility of software updates when the device is in supervised mode.
    softwareUpdatesForceDelayed *bool
    // Indicates whether or not to block Spotlight search from returning internet results on supervised device.
    spotlightBlockInternetResults *bool
    // Allow users to boot devices into recovery mode with unpaired devices. Available for devices running iOS and iPadOS versions 14.5 and later.
    unpairedExternalBootToRecoveryAllowed *bool
    // Indicates if connecting to USB accessories while the device is locked is allowed (iOS 11.4.1 and later).
    usbRestrictedModeBlocked *bool
    // Indicates whether or not to block voice dialing.
    voiceDialingBlocked *bool
    // Indicates whether or not the creation of VPN configurations is blocked (iOS 11.0 and later).
    vpnBlockCreation *bool
    // Indicates whether or not to allow wallpaper modification on supervised device (iOS 9.0 and later) .
    wallpaperBlockModification *bool
    // Indicates whether or not to force the device to use only Wi-Fi networks from configuration profiles when the device is in supervised mode. Available for devices running iOS and iPadOS versions 14.4 and earlier. Devices running 14.5+ should use the setting, 'WiFiConnectToAllowedNetworksOnlyForced.
    wiFiConnectOnlyToConfiguredNetworks *bool
    // Require devices to use Wi-Fi networks set up via configuration profiles. Available for devices running iOS and iPadOS versions 14.5 and later.
    wiFiConnectToAllowedNetworksOnlyForced *bool
    // Indicates whether or not Wi-Fi remains on, even when device is in airplane mode. Available for devices running iOS and iPadOS, versions 13.0 and later.
    wifiPowerOnForced *bool
}
// NewIosGeneralDeviceConfiguration instantiates a new IosGeneralDeviceConfiguration and sets the default values.
func NewIosGeneralDeviceConfiguration()(*IosGeneralDeviceConfiguration) {
    m := &IosGeneralDeviceConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.iosGeneralDeviceConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateIosGeneralDeviceConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosGeneralDeviceConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosGeneralDeviceConfiguration(), nil
}
// GetAccountBlockModification gets the accountBlockModification property value. Indicates whether or not to allow account modification when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) GetAccountBlockModification()(*bool) {
    return m.accountBlockModification
}
// GetActivationLockAllowWhenSupervised gets the activationLockAllowWhenSupervised property value. Indicates whether or not to allow activation lock when the device is in the supervised mode.
func (m *IosGeneralDeviceConfiguration) GetActivationLockAllowWhenSupervised()(*bool) {
    return m.activationLockAllowWhenSupervised
}
// GetAirDropBlocked gets the airDropBlocked property value. Indicates whether or not to allow AirDrop when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) GetAirDropBlocked()(*bool) {
    return m.airDropBlocked
}
// GetAirDropForceUnmanagedDropTarget gets the airDropForceUnmanagedDropTarget property value. Indicates whether or not to cause AirDrop to be considered an unmanaged drop target (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) GetAirDropForceUnmanagedDropTarget()(*bool) {
    return m.airDropForceUnmanagedDropTarget
}
// GetAirPlayForcePairingPasswordForOutgoingRequests gets the airPlayForcePairingPasswordForOutgoingRequests property value. Indicates whether or not to enforce all devices receiving AirPlay requests from this device to use a pairing password.
func (m *IosGeneralDeviceConfiguration) GetAirPlayForcePairingPasswordForOutgoingRequests()(*bool) {
    return m.airPlayForcePairingPasswordForOutgoingRequests
}
// GetAirPrintBlockCredentialsStorage gets the airPrintBlockCredentialsStorage property value. Indicates whether or not keychain storage of username and password for Airprint is blocked (iOS 11.0 and later).
func (m *IosGeneralDeviceConfiguration) GetAirPrintBlockCredentialsStorage()(*bool) {
    return m.airPrintBlockCredentialsStorage
}
// GetAirPrintBlocked gets the airPrintBlocked property value. Indicates whether or not AirPrint is blocked (iOS 11.0 and later).
func (m *IosGeneralDeviceConfiguration) GetAirPrintBlocked()(*bool) {
    return m.airPrintBlocked
}
// GetAirPrintBlockiBeaconDiscovery gets the airPrintBlockiBeaconDiscovery property value. Indicates whether or not iBeacon discovery of AirPrint printers is blocked. This prevents spurious AirPrint Bluetooth beacons from phishing for network traffic (iOS 11.0 and later).
func (m *IosGeneralDeviceConfiguration) GetAirPrintBlockiBeaconDiscovery()(*bool) {
    return m.airPrintBlockiBeaconDiscovery
}
// GetAirPrintForceTrustedTLS gets the airPrintForceTrustedTLS property value. Indicates if trusted certificates are required for TLS printing communication (iOS 11.0 and later).
func (m *IosGeneralDeviceConfiguration) GetAirPrintForceTrustedTLS()(*bool) {
    return m.airPrintForceTrustedTLS
}
// GetAppClipsBlocked gets the appClipsBlocked property value. Prevents a user from adding any App Clips and removes any existing App Clips on the device.
func (m *IosGeneralDeviceConfiguration) GetAppClipsBlocked()(*bool) {
    return m.appClipsBlocked
}
// GetAppleNewsBlocked gets the appleNewsBlocked property value. Indicates whether or not to block the user from using News when the device is in supervised mode (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) GetAppleNewsBlocked()(*bool) {
    return m.appleNewsBlocked
}
// GetApplePersonalizedAdsBlocked gets the applePersonalizedAdsBlocked property value. Limits Apple personalized advertising when true. Available in iOS 14 and later.
func (m *IosGeneralDeviceConfiguration) GetApplePersonalizedAdsBlocked()(*bool) {
    return m.applePersonalizedAdsBlocked
}
// GetAppleWatchBlockPairing gets the appleWatchBlockPairing property value. Indicates whether or not to allow Apple Watch pairing when the device is in supervised mode (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) GetAppleWatchBlockPairing()(*bool) {
    return m.appleWatchBlockPairing
}
// GetAppleWatchForceWristDetection gets the appleWatchForceWristDetection property value. Indicates whether or not to force a paired Apple Watch to use Wrist Detection (iOS 8.2 and later).
func (m *IosGeneralDeviceConfiguration) GetAppleWatchForceWristDetection()(*bool) {
    return m.appleWatchForceWristDetection
}
// GetAppRemovalBlocked gets the appRemovalBlocked property value. Indicates if the removal of apps is allowed.
func (m *IosGeneralDeviceConfiguration) GetAppRemovalBlocked()(*bool) {
    return m.appRemovalBlocked
}
// GetAppsSingleAppModeList gets the appsSingleAppModeList property value. Gets or sets the list of iOS apps allowed to autonomously enter Single App Mode. Supervised only. iOS 7.0 and later. This collection can contain a maximum of 500 elements.
func (m *IosGeneralDeviceConfiguration) GetAppsSingleAppModeList()([]AppListItemable) {
    return m.appsSingleAppModeList
}
// GetAppStoreBlockAutomaticDownloads gets the appStoreBlockAutomaticDownloads property value. Indicates whether or not to block the automatic downloading of apps purchased on other devices when the device is in supervised mode (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) GetAppStoreBlockAutomaticDownloads()(*bool) {
    return m.appStoreBlockAutomaticDownloads
}
// GetAppStoreBlocked gets the appStoreBlocked property value. Indicates whether or not to block the user from using the App Store. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) GetAppStoreBlocked()(*bool) {
    return m.appStoreBlocked
}
// GetAppStoreBlockInAppPurchases gets the appStoreBlockInAppPurchases property value. Indicates whether or not to block the user from making in app purchases.
func (m *IosGeneralDeviceConfiguration) GetAppStoreBlockInAppPurchases()(*bool) {
    return m.appStoreBlockInAppPurchases
}
// GetAppStoreBlockUIAppInstallation gets the appStoreBlockUIAppInstallation property value. Indicates whether or not to block the App Store app, not restricting installation through Host apps. Applies to supervised mode only (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) GetAppStoreBlockUIAppInstallation()(*bool) {
    return m.appStoreBlockUIAppInstallation
}
// GetAppStoreRequirePassword gets the appStoreRequirePassword property value. Indicates whether or not to require a password when using the app store.
func (m *IosGeneralDeviceConfiguration) GetAppStoreRequirePassword()(*bool) {
    return m.appStoreRequirePassword
}
// GetAppsVisibilityList gets the appsVisibilityList property value. List of apps in the visibility list (either visible/launchable apps list or hidden/unlaunchable apps list, controlled by AppsVisibilityListType) (iOS 9.3 and later). This collection can contain a maximum of 10000 elements.
func (m *IosGeneralDeviceConfiguration) GetAppsVisibilityList()([]AppListItemable) {
    return m.appsVisibilityList
}
// GetAppsVisibilityListType gets the appsVisibilityListType property value. Possible values of the compliance app list.
func (m *IosGeneralDeviceConfiguration) GetAppsVisibilityListType()(*AppListType) {
    return m.appsVisibilityListType
}
// GetAutoFillForceAuthentication gets the autoFillForceAuthentication property value. Indicates whether or not to force user authentication before autofilling passwords and credit card information in Safari and other apps on supervised devices.
func (m *IosGeneralDeviceConfiguration) GetAutoFillForceAuthentication()(*bool) {
    return m.autoFillForceAuthentication
}
// GetAutoUnlockBlocked gets the autoUnlockBlocked property value. Blocks users from unlocking their device with Apple Watch. Available for devices running iOS and iPadOS versions 14.5 and later.
func (m *IosGeneralDeviceConfiguration) GetAutoUnlockBlocked()(*bool) {
    return m.autoUnlockBlocked
}
// GetBlockSystemAppRemoval gets the blockSystemAppRemoval property value. Indicates whether or not the removal of system apps from the device is blocked on a supervised device (iOS 11.0 and later).
func (m *IosGeneralDeviceConfiguration) GetBlockSystemAppRemoval()(*bool) {
    return m.blockSystemAppRemoval
}
// GetBluetoothBlockModification gets the bluetoothBlockModification property value. Indicates whether or not to allow modification of Bluetooth settings when the device is in supervised mode (iOS 10.0 and later).
func (m *IosGeneralDeviceConfiguration) GetBluetoothBlockModification()(*bool) {
    return m.bluetoothBlockModification
}
// GetCameraBlocked gets the cameraBlocked property value. Indicates whether or not to block the user from accessing the camera of the device. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) GetCameraBlocked()(*bool) {
    return m.cameraBlocked
}
// GetCellularBlockDataRoaming gets the cellularBlockDataRoaming property value. Indicates whether or not to block data roaming.
func (m *IosGeneralDeviceConfiguration) GetCellularBlockDataRoaming()(*bool) {
    return m.cellularBlockDataRoaming
}
// GetCellularBlockGlobalBackgroundFetchWhileRoaming gets the cellularBlockGlobalBackgroundFetchWhileRoaming property value. Indicates whether or not to block global background fetch while roaming.
func (m *IosGeneralDeviceConfiguration) GetCellularBlockGlobalBackgroundFetchWhileRoaming()(*bool) {
    return m.cellularBlockGlobalBackgroundFetchWhileRoaming
}
// GetCellularBlockPerAppDataModification gets the cellularBlockPerAppDataModification property value. Indicates whether or not to allow changes to cellular app data usage settings when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) GetCellularBlockPerAppDataModification()(*bool) {
    return m.cellularBlockPerAppDataModification
}
// GetCellularBlockPersonalHotspot gets the cellularBlockPersonalHotspot property value. Indicates whether or not to block Personal Hotspot.
func (m *IosGeneralDeviceConfiguration) GetCellularBlockPersonalHotspot()(*bool) {
    return m.cellularBlockPersonalHotspot
}
// GetCellularBlockPersonalHotspotModification gets the cellularBlockPersonalHotspotModification property value. Indicates whether or not to block the user from modifying the personal hotspot setting (iOS 12.2 or later).
func (m *IosGeneralDeviceConfiguration) GetCellularBlockPersonalHotspotModification()(*bool) {
    return m.cellularBlockPersonalHotspotModification
}
// GetCellularBlockPlanModification gets the cellularBlockPlanModification property value. Indicates whether or not to allow users to change the settings of the cellular plan on a supervised device.
func (m *IosGeneralDeviceConfiguration) GetCellularBlockPlanModification()(*bool) {
    return m.cellularBlockPlanModification
}
// GetCellularBlockVoiceRoaming gets the cellularBlockVoiceRoaming property value. Indicates whether or not to block voice roaming.
func (m *IosGeneralDeviceConfiguration) GetCellularBlockVoiceRoaming()(*bool) {
    return m.cellularBlockVoiceRoaming
}
// GetCertificatesBlockUntrustedTlsCertificates gets the certificatesBlockUntrustedTlsCertificates property value. Indicates whether or not to block untrusted TLS certificates.
func (m *IosGeneralDeviceConfiguration) GetCertificatesBlockUntrustedTlsCertificates()(*bool) {
    return m.certificatesBlockUntrustedTlsCertificates
}
// GetClassroomAppBlockRemoteScreenObservation gets the classroomAppBlockRemoteScreenObservation property value. Indicates whether or not to allow remote screen observation by Classroom app when the device is in supervised mode (iOS 9.3 and later).
func (m *IosGeneralDeviceConfiguration) GetClassroomAppBlockRemoteScreenObservation()(*bool) {
    return m.classroomAppBlockRemoteScreenObservation
}
// GetClassroomAppForceUnpromptedScreenObservation gets the classroomAppForceUnpromptedScreenObservation property value. Indicates whether or not to automatically give permission to the teacher of a managed course on the Classroom app to view a student's screen without prompting when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) GetClassroomAppForceUnpromptedScreenObservation()(*bool) {
    return m.classroomAppForceUnpromptedScreenObservation
}
// GetClassroomForceAutomaticallyJoinClasses gets the classroomForceAutomaticallyJoinClasses property value. Indicates whether or not to automatically give permission to the teacher's requests, without prompting the student, when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) GetClassroomForceAutomaticallyJoinClasses()(*bool) {
    return m.classroomForceAutomaticallyJoinClasses
}
// GetClassroomForceRequestPermissionToLeaveClasses gets the classroomForceRequestPermissionToLeaveClasses property value. Indicates whether a student enrolled in an unmanaged course via Classroom will request permission from the teacher when attempting to leave the course (iOS 11.3 and later).
func (m *IosGeneralDeviceConfiguration) GetClassroomForceRequestPermissionToLeaveClasses()(*bool) {
    return m.classroomForceRequestPermissionToLeaveClasses
}
// GetClassroomForceUnpromptedAppAndDeviceLock gets the classroomForceUnpromptedAppAndDeviceLock property value. Indicates whether or not to allow the teacher to lock apps or the device without prompting the student. Supervised only.
func (m *IosGeneralDeviceConfiguration) GetClassroomForceUnpromptedAppAndDeviceLock()(*bool) {
    return m.classroomForceUnpromptedAppAndDeviceLock
}
// GetCompliantAppListType gets the compliantAppListType property value. Possible values of the compliance app list.
func (m *IosGeneralDeviceConfiguration) GetCompliantAppListType()(*AppListType) {
    return m.compliantAppListType
}
// GetCompliantAppsList gets the compliantAppsList property value. List of apps in the compliance (either allow list or block list, controlled by CompliantAppListType). This collection can contain a maximum of 10000 elements.
func (m *IosGeneralDeviceConfiguration) GetCompliantAppsList()([]AppListItemable) {
    return m.compliantAppsList
}
// GetConfigurationProfileBlockChanges gets the configurationProfileBlockChanges property value. Indicates whether or not to block the user from installing configuration profiles and certificates interactively when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) GetConfigurationProfileBlockChanges()(*bool) {
    return m.configurationProfileBlockChanges
}
// GetContactsAllowManagedToUnmanagedWrite gets the contactsAllowManagedToUnmanagedWrite property value. Indicates whether or not managed apps can write contacts to unmanaged contacts accounts (iOS 12.0 and later).
func (m *IosGeneralDeviceConfiguration) GetContactsAllowManagedToUnmanagedWrite()(*bool) {
    return m.contactsAllowManagedToUnmanagedWrite
}
// GetContactsAllowUnmanagedToManagedRead gets the contactsAllowUnmanagedToManagedRead property value. Indicates whether or not unmanaged apps can read from managed contacts accounts (iOS 12.0 or later).
func (m *IosGeneralDeviceConfiguration) GetContactsAllowUnmanagedToManagedRead()(*bool) {
    return m.contactsAllowUnmanagedToManagedRead
}
// GetContinuousPathKeyboardBlocked gets the continuousPathKeyboardBlocked property value. Indicates whether or not to block the continuous path keyboard when the device is supervised (iOS 13 or later).
func (m *IosGeneralDeviceConfiguration) GetContinuousPathKeyboardBlocked()(*bool) {
    return m.continuousPathKeyboardBlocked
}
// GetDateAndTimeForceSetAutomatically gets the dateAndTimeForceSetAutomatically property value. Indicates whether or not the Date and Time 'Set Automatically' feature is enabled and cannot be turned off by the user (iOS 12.0 and later).
func (m *IosGeneralDeviceConfiguration) GetDateAndTimeForceSetAutomatically()(*bool) {
    return m.dateAndTimeForceSetAutomatically
}
// GetDefinitionLookupBlocked gets the definitionLookupBlocked property value. Indicates whether or not to block definition lookup when the device is in supervised mode (iOS 8.1.3 and later ).
func (m *IosGeneralDeviceConfiguration) GetDefinitionLookupBlocked()(*bool) {
    return m.definitionLookupBlocked
}
// GetDeviceBlockEnableRestrictions gets the deviceBlockEnableRestrictions property value. Indicates whether or not to allow the user to enables restrictions in the device settings when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) GetDeviceBlockEnableRestrictions()(*bool) {
    return m.deviceBlockEnableRestrictions
}
// GetDeviceBlockEraseContentAndSettings gets the deviceBlockEraseContentAndSettings property value. Indicates whether or not to allow the use of the 'Erase all content and settings' option on the device when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) GetDeviceBlockEraseContentAndSettings()(*bool) {
    return m.deviceBlockEraseContentAndSettings
}
// GetDeviceBlockNameModification gets the deviceBlockNameModification property value. Indicates whether or not to allow device name modification when the device is in supervised mode (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) GetDeviceBlockNameModification()(*bool) {
    return m.deviceBlockNameModification
}
// GetDiagnosticDataBlockSubmission gets the diagnosticDataBlockSubmission property value. Indicates whether or not to block diagnostic data submission.
func (m *IosGeneralDeviceConfiguration) GetDiagnosticDataBlockSubmission()(*bool) {
    return m.diagnosticDataBlockSubmission
}
// GetDiagnosticDataBlockSubmissionModification gets the diagnosticDataBlockSubmissionModification property value. Indicates whether or not to allow diagnostics submission settings modification when the device is in supervised mode (iOS 9.3.2 and later).
func (m *IosGeneralDeviceConfiguration) GetDiagnosticDataBlockSubmissionModification()(*bool) {
    return m.diagnosticDataBlockSubmissionModification
}
// GetDocumentsBlockManagedDocumentsInUnmanagedApps gets the documentsBlockManagedDocumentsInUnmanagedApps property value. Indicates whether or not to block the user from viewing managed documents in unmanaged apps.
func (m *IosGeneralDeviceConfiguration) GetDocumentsBlockManagedDocumentsInUnmanagedApps()(*bool) {
    return m.documentsBlockManagedDocumentsInUnmanagedApps
}
// GetDocumentsBlockUnmanagedDocumentsInManagedApps gets the documentsBlockUnmanagedDocumentsInManagedApps property value. Indicates whether or not to block the user from viewing unmanaged documents in managed apps.
func (m *IosGeneralDeviceConfiguration) GetDocumentsBlockUnmanagedDocumentsInManagedApps()(*bool) {
    return m.documentsBlockUnmanagedDocumentsInManagedApps
}
// GetEmailInDomainSuffixes gets the emailInDomainSuffixes property value. An email address lacking a suffix that matches any of these strings will be considered out-of-domain.
func (m *IosGeneralDeviceConfiguration) GetEmailInDomainSuffixes()([]string) {
    return m.emailInDomainSuffixes
}
// GetEnterpriseAppBlockTrust gets the enterpriseAppBlockTrust property value. Indicates whether or not to block the user from trusting an enterprise app.
func (m *IosGeneralDeviceConfiguration) GetEnterpriseAppBlockTrust()(*bool) {
    return m.enterpriseAppBlockTrust
}
// GetEnterpriseAppBlockTrustModification gets the enterpriseAppBlockTrustModification property value. [Deprecated] Configuring this setting and setting the value to 'true' has no effect on the device.
func (m *IosGeneralDeviceConfiguration) GetEnterpriseAppBlockTrustModification()(*bool) {
    return m.enterpriseAppBlockTrustModification
}
// GetEnterpriseBookBlockBackup gets the enterpriseBookBlockBackup property value. Indicates whether or not Enterprise book back up is blocked.
func (m *IosGeneralDeviceConfiguration) GetEnterpriseBookBlockBackup()(*bool) {
    return m.enterpriseBookBlockBackup
}
// GetEnterpriseBookBlockMetadataSync gets the enterpriseBookBlockMetadataSync property value. Indicates whether or not Enterprise book notes and highlights sync is blocked.
func (m *IosGeneralDeviceConfiguration) GetEnterpriseBookBlockMetadataSync()(*bool) {
    return m.enterpriseBookBlockMetadataSync
}
// GetEsimBlockModification gets the esimBlockModification property value. Indicates whether or not to allow the addition or removal of cellular plans on the eSIM of a supervised device.
func (m *IosGeneralDeviceConfiguration) GetEsimBlockModification()(*bool) {
    return m.esimBlockModification
}
// GetFaceTimeBlocked gets the faceTimeBlocked property value. Indicates whether or not to block the user from using FaceTime. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) GetFaceTimeBlocked()(*bool) {
    return m.faceTimeBlocked
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosGeneralDeviceConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["accountBlockModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountBlockModification(val)
        }
        return nil
    }
    res["activationLockAllowWhenSupervised"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivationLockAllowWhenSupervised(val)
        }
        return nil
    }
    res["airDropBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAirDropBlocked(val)
        }
        return nil
    }
    res["airDropForceUnmanagedDropTarget"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAirDropForceUnmanagedDropTarget(val)
        }
        return nil
    }
    res["airPlayForcePairingPasswordForOutgoingRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAirPlayForcePairingPasswordForOutgoingRequests(val)
        }
        return nil
    }
    res["airPrintBlockCredentialsStorage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAirPrintBlockCredentialsStorage(val)
        }
        return nil
    }
    res["airPrintBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAirPrintBlocked(val)
        }
        return nil
    }
    res["airPrintBlockiBeaconDiscovery"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAirPrintBlockiBeaconDiscovery(val)
        }
        return nil
    }
    res["airPrintForceTrustedTLS"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAirPrintForceTrustedTLS(val)
        }
        return nil
    }
    res["appClipsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppClipsBlocked(val)
        }
        return nil
    }
    res["appleNewsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppleNewsBlocked(val)
        }
        return nil
    }
    res["applePersonalizedAdsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplePersonalizedAdsBlocked(val)
        }
        return nil
    }
    res["appleWatchBlockPairing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppleWatchBlockPairing(val)
        }
        return nil
    }
    res["appleWatchForceWristDetection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppleWatchForceWristDetection(val)
        }
        return nil
    }
    res["appRemovalBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppRemovalBlocked(val)
        }
        return nil
    }
    res["appsSingleAppModeList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppListItemable, len(val))
            for i, v := range val {
                res[i] = v.(AppListItemable)
            }
            m.SetAppsSingleAppModeList(res)
        }
        return nil
    }
    res["appStoreBlockAutomaticDownloads"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppStoreBlockAutomaticDownloads(val)
        }
        return nil
    }
    res["appStoreBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppStoreBlocked(val)
        }
        return nil
    }
    res["appStoreBlockInAppPurchases"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppStoreBlockInAppPurchases(val)
        }
        return nil
    }
    res["appStoreBlockUIAppInstallation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppStoreBlockUIAppInstallation(val)
        }
        return nil
    }
    res["appStoreRequirePassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppStoreRequirePassword(val)
        }
        return nil
    }
    res["appsVisibilityList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppListItemable, len(val))
            for i, v := range val {
                res[i] = v.(AppListItemable)
            }
            m.SetAppsVisibilityList(res)
        }
        return nil
    }
    res["appsVisibilityListType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppListType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppsVisibilityListType(val.(*AppListType))
        }
        return nil
    }
    res["autoFillForceAuthentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoFillForceAuthentication(val)
        }
        return nil
    }
    res["autoUnlockBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoUnlockBlocked(val)
        }
        return nil
    }
    res["blockSystemAppRemoval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockSystemAppRemoval(val)
        }
        return nil
    }
    res["bluetoothBlockModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBluetoothBlockModification(val)
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
    res["cellularBlockDataRoaming"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCellularBlockDataRoaming(val)
        }
        return nil
    }
    res["cellularBlockGlobalBackgroundFetchWhileRoaming"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCellularBlockGlobalBackgroundFetchWhileRoaming(val)
        }
        return nil
    }
    res["cellularBlockPerAppDataModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCellularBlockPerAppDataModification(val)
        }
        return nil
    }
    res["cellularBlockPersonalHotspot"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCellularBlockPersonalHotspot(val)
        }
        return nil
    }
    res["cellularBlockPersonalHotspotModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCellularBlockPersonalHotspotModification(val)
        }
        return nil
    }
    res["cellularBlockPlanModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCellularBlockPlanModification(val)
        }
        return nil
    }
    res["cellularBlockVoiceRoaming"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCellularBlockVoiceRoaming(val)
        }
        return nil
    }
    res["certificatesBlockUntrustedTlsCertificates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificatesBlockUntrustedTlsCertificates(val)
        }
        return nil
    }
    res["classroomAppBlockRemoteScreenObservation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassroomAppBlockRemoteScreenObservation(val)
        }
        return nil
    }
    res["classroomAppForceUnpromptedScreenObservation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassroomAppForceUnpromptedScreenObservation(val)
        }
        return nil
    }
    res["classroomForceAutomaticallyJoinClasses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassroomForceAutomaticallyJoinClasses(val)
        }
        return nil
    }
    res["classroomForceRequestPermissionToLeaveClasses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassroomForceRequestPermissionToLeaveClasses(val)
        }
        return nil
    }
    res["classroomForceUnpromptedAppAndDeviceLock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassroomForceUnpromptedAppAndDeviceLock(val)
        }
        return nil
    }
    res["compliantAppListType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppListType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliantAppListType(val.(*AppListType))
        }
        return nil
    }
    res["compliantAppsList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppListItemable, len(val))
            for i, v := range val {
                res[i] = v.(AppListItemable)
            }
            m.SetCompliantAppsList(res)
        }
        return nil
    }
    res["configurationProfileBlockChanges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationProfileBlockChanges(val)
        }
        return nil
    }
    res["contactsAllowManagedToUnmanagedWrite"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactsAllowManagedToUnmanagedWrite(val)
        }
        return nil
    }
    res["contactsAllowUnmanagedToManagedRead"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactsAllowUnmanagedToManagedRead(val)
        }
        return nil
    }
    res["continuousPathKeyboardBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContinuousPathKeyboardBlocked(val)
        }
        return nil
    }
    res["dateAndTimeForceSetAutomatically"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateAndTimeForceSetAutomatically(val)
        }
        return nil
    }
    res["definitionLookupBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefinitionLookupBlocked(val)
        }
        return nil
    }
    res["deviceBlockEnableRestrictions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceBlockEnableRestrictions(val)
        }
        return nil
    }
    res["deviceBlockEraseContentAndSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceBlockEraseContentAndSettings(val)
        }
        return nil
    }
    res["deviceBlockNameModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceBlockNameModification(val)
        }
        return nil
    }
    res["diagnosticDataBlockSubmission"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiagnosticDataBlockSubmission(val)
        }
        return nil
    }
    res["diagnosticDataBlockSubmissionModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiagnosticDataBlockSubmissionModification(val)
        }
        return nil
    }
    res["documentsBlockManagedDocumentsInUnmanagedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentsBlockManagedDocumentsInUnmanagedApps(val)
        }
        return nil
    }
    res["documentsBlockUnmanagedDocumentsInManagedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentsBlockUnmanagedDocumentsInManagedApps(val)
        }
        return nil
    }
    res["emailInDomainSuffixes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetEmailInDomainSuffixes(res)
        }
        return nil
    }
    res["enterpriseAppBlockTrust"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnterpriseAppBlockTrust(val)
        }
        return nil
    }
    res["enterpriseAppBlockTrustModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnterpriseAppBlockTrustModification(val)
        }
        return nil
    }
    res["enterpriseBookBlockBackup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnterpriseBookBlockBackup(val)
        }
        return nil
    }
    res["enterpriseBookBlockMetadataSync"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnterpriseBookBlockMetadataSync(val)
        }
        return nil
    }
    res["esimBlockModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEsimBlockModification(val)
        }
        return nil
    }
    res["faceTimeBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFaceTimeBlocked(val)
        }
        return nil
    }
    res["filesNetworkDriveAccessBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilesNetworkDriveAccessBlocked(val)
        }
        return nil
    }
    res["filesUsbDriveAccessBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilesUsbDriveAccessBlocked(val)
        }
        return nil
    }
    res["findMyDeviceInFindMyAppBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFindMyDeviceInFindMyAppBlocked(val)
        }
        return nil
    }
    res["findMyFriendsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFindMyFriendsBlocked(val)
        }
        return nil
    }
    res["findMyFriendsInFindMyAppBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFindMyFriendsInFindMyAppBlocked(val)
        }
        return nil
    }
    res["gameCenterBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGameCenterBlocked(val)
        }
        return nil
    }
    res["gamingBlockGameCenterFriends"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGamingBlockGameCenterFriends(val)
        }
        return nil
    }
    res["gamingBlockMultiplayer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGamingBlockMultiplayer(val)
        }
        return nil
    }
    res["hostPairingBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostPairingBlocked(val)
        }
        return nil
    }
    res["iBooksStoreBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIBooksStoreBlocked(val)
        }
        return nil
    }
    res["iBooksStoreBlockErotica"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIBooksStoreBlockErotica(val)
        }
        return nil
    }
    res["iCloudBlockActivityContinuation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudBlockActivityContinuation(val)
        }
        return nil
    }
    res["iCloudBlockBackup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudBlockBackup(val)
        }
        return nil
    }
    res["iCloudBlockDocumentSync"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudBlockDocumentSync(val)
        }
        return nil
    }
    res["iCloudBlockManagedAppsSync"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudBlockManagedAppsSync(val)
        }
        return nil
    }
    res["iCloudBlockPhotoLibrary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudBlockPhotoLibrary(val)
        }
        return nil
    }
    res["iCloudBlockPhotoStreamSync"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudBlockPhotoStreamSync(val)
        }
        return nil
    }
    res["iCloudBlockSharedPhotoStream"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudBlockSharedPhotoStream(val)
        }
        return nil
    }
    res["iCloudPrivateRelayBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudPrivateRelayBlocked(val)
        }
        return nil
    }
    res["iCloudRequireEncryptedBackup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICloudRequireEncryptedBackup(val)
        }
        return nil
    }
    res["iTunesBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetITunesBlocked(val)
        }
        return nil
    }
    res["iTunesBlockExplicitContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetITunesBlockExplicitContent(val)
        }
        return nil
    }
    res["iTunesBlockMusicService"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetITunesBlockMusicService(val)
        }
        return nil
    }
    res["iTunesBlockRadio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetITunesBlockRadio(val)
        }
        return nil
    }
    res["keyboardBlockAutoCorrect"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyboardBlockAutoCorrect(val)
        }
        return nil
    }
    res["keyboardBlockDictation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyboardBlockDictation(val)
        }
        return nil
    }
    res["keyboardBlockPredictive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyboardBlockPredictive(val)
        }
        return nil
    }
    res["keyboardBlockShortcuts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyboardBlockShortcuts(val)
        }
        return nil
    }
    res["keyboardBlockSpellCheck"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyboardBlockSpellCheck(val)
        }
        return nil
    }
    res["keychainBlockCloudSync"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeychainBlockCloudSync(val)
        }
        return nil
    }
    res["kioskModeAllowAssistiveSpeak"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAllowAssistiveSpeak(val)
        }
        return nil
    }
    res["kioskModeAllowAssistiveTouchSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAllowAssistiveTouchSettings(val)
        }
        return nil
    }
    res["kioskModeAllowAutoLock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAllowAutoLock(val)
        }
        return nil
    }
    res["kioskModeAllowColorInversionSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAllowColorInversionSettings(val)
        }
        return nil
    }
    res["kioskModeAllowRingerSwitch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAllowRingerSwitch(val)
        }
        return nil
    }
    res["kioskModeAllowScreenRotation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAllowScreenRotation(val)
        }
        return nil
    }
    res["kioskModeAllowSleepButton"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAllowSleepButton(val)
        }
        return nil
    }
    res["kioskModeAllowTouchscreen"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAllowTouchscreen(val)
        }
        return nil
    }
    res["kioskModeAllowVoiceControlModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAllowVoiceControlModification(val)
        }
        return nil
    }
    res["kioskModeAllowVoiceOverSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAllowVoiceOverSettings(val)
        }
        return nil
    }
    res["kioskModeAllowVolumeButtons"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAllowVolumeButtons(val)
        }
        return nil
    }
    res["kioskModeAllowZoomSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAllowZoomSettings(val)
        }
        return nil
    }
    res["kioskModeAppStoreUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAppStoreUrl(val)
        }
        return nil
    }
    res["kioskModeAppType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIosKioskModeAppType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeAppType(val.(*IosKioskModeAppType))
        }
        return nil
    }
    res["kioskModeBlockAutoLock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeBlockAutoLock(val)
        }
        return nil
    }
    res["kioskModeBlockRingerSwitch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeBlockRingerSwitch(val)
        }
        return nil
    }
    res["kioskModeBlockScreenRotation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeBlockScreenRotation(val)
        }
        return nil
    }
    res["kioskModeBlockSleepButton"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeBlockSleepButton(val)
        }
        return nil
    }
    res["kioskModeBlockTouchscreen"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeBlockTouchscreen(val)
        }
        return nil
    }
    res["kioskModeBlockVolumeButtons"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeBlockVolumeButtons(val)
        }
        return nil
    }
    res["kioskModeBuiltInAppId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeBuiltInAppId(val)
        }
        return nil
    }
    res["kioskModeEnableVoiceControl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeEnableVoiceControl(val)
        }
        return nil
    }
    res["kioskModeManagedAppId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeManagedAppId(val)
        }
        return nil
    }
    res["kioskModeRequireAssistiveTouch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeRequireAssistiveTouch(val)
        }
        return nil
    }
    res["kioskModeRequireColorInversion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeRequireColorInversion(val)
        }
        return nil
    }
    res["kioskModeRequireMonoAudio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeRequireMonoAudio(val)
        }
        return nil
    }
    res["kioskModeRequireVoiceOver"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeRequireVoiceOver(val)
        }
        return nil
    }
    res["kioskModeRequireZoom"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskModeRequireZoom(val)
        }
        return nil
    }
    res["lockScreenBlockControlCenter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLockScreenBlockControlCenter(val)
        }
        return nil
    }
    res["lockScreenBlockNotificationView"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLockScreenBlockNotificationView(val)
        }
        return nil
    }
    res["lockScreenBlockPassbook"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLockScreenBlockPassbook(val)
        }
        return nil
    }
    res["lockScreenBlockTodayView"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLockScreenBlockTodayView(val)
        }
        return nil
    }
    res["managedPasteboardRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedPasteboardRequired(val)
        }
        return nil
    }
    res["mediaContentRatingApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRatingAppsType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaContentRatingApps(val.(*RatingAppsType))
        }
        return nil
    }
    res["mediaContentRatingAustralia"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaContentRatingAustraliaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaContentRatingAustralia(val.(MediaContentRatingAustraliaable))
        }
        return nil
    }
    res["mediaContentRatingCanada"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaContentRatingCanadaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaContentRatingCanada(val.(MediaContentRatingCanadaable))
        }
        return nil
    }
    res["mediaContentRatingFrance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaContentRatingFranceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaContentRatingFrance(val.(MediaContentRatingFranceable))
        }
        return nil
    }
    res["mediaContentRatingGermany"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaContentRatingGermanyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaContentRatingGermany(val.(MediaContentRatingGermanyable))
        }
        return nil
    }
    res["mediaContentRatingIreland"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaContentRatingIrelandFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaContentRatingIreland(val.(MediaContentRatingIrelandable))
        }
        return nil
    }
    res["mediaContentRatingJapan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaContentRatingJapanFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaContentRatingJapan(val.(MediaContentRatingJapanable))
        }
        return nil
    }
    res["mediaContentRatingNewZealand"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaContentRatingNewZealandFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaContentRatingNewZealand(val.(MediaContentRatingNewZealandable))
        }
        return nil
    }
    res["mediaContentRatingUnitedKingdom"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaContentRatingUnitedKingdomFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaContentRatingUnitedKingdom(val.(MediaContentRatingUnitedKingdomable))
        }
        return nil
    }
    res["mediaContentRatingUnitedStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaContentRatingUnitedStatesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaContentRatingUnitedStates(val.(MediaContentRatingUnitedStatesable))
        }
        return nil
    }
    res["messagesBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessagesBlocked(val)
        }
        return nil
    }
    res["networkUsageRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIosNetworkUsageRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosNetworkUsageRuleable, len(val))
            for i, v := range val {
                res[i] = v.(IosNetworkUsageRuleable)
            }
            m.SetNetworkUsageRules(res)
        }
        return nil
    }
    res["nfcBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNfcBlocked(val)
        }
        return nil
    }
    res["notificationsBlockSettingsModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationsBlockSettingsModification(val)
        }
        return nil
    }
    res["onDeviceOnlyDictationForced"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnDeviceOnlyDictationForced(val)
        }
        return nil
    }
    res["onDeviceOnlyTranslationForced"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnDeviceOnlyTranslationForced(val)
        }
        return nil
    }
    res["passcodeBlockFingerprintModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeBlockFingerprintModification(val)
        }
        return nil
    }
    res["passcodeBlockFingerprintUnlock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeBlockFingerprintUnlock(val)
        }
        return nil
    }
    res["passcodeBlockModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeBlockModification(val)
        }
        return nil
    }
    res["passcodeBlockSimple"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeBlockSimple(val)
        }
        return nil
    }
    res["passcodeExpirationDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeExpirationDays(val)
        }
        return nil
    }
    res["passcodeMinimumCharacterSetCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeMinimumCharacterSetCount(val)
        }
        return nil
    }
    res["passcodeMinimumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeMinimumLength(val)
        }
        return nil
    }
    res["passcodeMinutesOfInactivityBeforeLock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeMinutesOfInactivityBeforeLock(val)
        }
        return nil
    }
    res["passcodeMinutesOfInactivityBeforeScreenTimeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeMinutesOfInactivityBeforeScreenTimeout(val)
        }
        return nil
    }
    res["passcodePreviousPasscodeBlockCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodePreviousPasscodeBlockCount(val)
        }
        return nil
    }
    res["passcodeRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeRequired(val)
        }
        return nil
    }
    res["passcodeRequiredType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRequiredPasswordType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeRequiredType(val.(*RequiredPasswordType))
        }
        return nil
    }
    res["passcodeSignInFailureCountBeforeWipe"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeSignInFailureCountBeforeWipe(val)
        }
        return nil
    }
    res["passwordBlockAirDropSharing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordBlockAirDropSharing(val)
        }
        return nil
    }
    res["passwordBlockAutoFill"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordBlockAutoFill(val)
        }
        return nil
    }
    res["passwordBlockProximityRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordBlockProximityRequests(val)
        }
        return nil
    }
    res["pkiBlockOTAUpdates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPkiBlockOTAUpdates(val)
        }
        return nil
    }
    res["podcastsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPodcastsBlocked(val)
        }
        return nil
    }
    res["privacyForceLimitAdTracking"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivacyForceLimitAdTracking(val)
        }
        return nil
    }
    res["proximityBlockSetupToNewDevice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProximityBlockSetupToNewDevice(val)
        }
        return nil
    }
    res["safariBlockAutofill"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSafariBlockAutofill(val)
        }
        return nil
    }
    res["safariBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSafariBlocked(val)
        }
        return nil
    }
    res["safariBlockJavaScript"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSafariBlockJavaScript(val)
        }
        return nil
    }
    res["safariBlockPopups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSafariBlockPopups(val)
        }
        return nil
    }
    res["safariCookieSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWebBrowserCookieSettings)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSafariCookieSettings(val.(*WebBrowserCookieSettings))
        }
        return nil
    }
    res["safariManagedDomains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSafariManagedDomains(res)
        }
        return nil
    }
    res["safariPasswordAutoFillDomains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSafariPasswordAutoFillDomains(res)
        }
        return nil
    }
    res["safariRequireFraudWarning"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSafariRequireFraudWarning(val)
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
    res["sharedDeviceBlockTemporarySessions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedDeviceBlockTemporarySessions(val)
        }
        return nil
    }
    res["siriBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiriBlocked(val)
        }
        return nil
    }
    res["siriBlockedWhenLocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiriBlockedWhenLocked(val)
        }
        return nil
    }
    res["siriBlockUserGeneratedContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiriBlockUserGeneratedContent(val)
        }
        return nil
    }
    res["siriRequireProfanityFilter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiriRequireProfanityFilter(val)
        }
        return nil
    }
    res["softwareUpdatesEnforcedDelayInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftwareUpdatesEnforcedDelayInDays(val)
        }
        return nil
    }
    res["softwareUpdatesForceDelayed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftwareUpdatesForceDelayed(val)
        }
        return nil
    }
    res["spotlightBlockInternetResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpotlightBlockInternetResults(val)
        }
        return nil
    }
    res["unpairedExternalBootToRecoveryAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnpairedExternalBootToRecoveryAllowed(val)
        }
        return nil
    }
    res["usbRestrictedModeBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsbRestrictedModeBlocked(val)
        }
        return nil
    }
    res["voiceDialingBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVoiceDialingBlocked(val)
        }
        return nil
    }
    res["vpnBlockCreation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVpnBlockCreation(val)
        }
        return nil
    }
    res["wallpaperBlockModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWallpaperBlockModification(val)
        }
        return nil
    }
    res["wiFiConnectOnlyToConfiguredNetworks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWiFiConnectOnlyToConfiguredNetworks(val)
        }
        return nil
    }
    res["wiFiConnectToAllowedNetworksOnlyForced"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWiFiConnectToAllowedNetworksOnlyForced(val)
        }
        return nil
    }
    res["wifiPowerOnForced"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiPowerOnForced(val)
        }
        return nil
    }
    return res
}
// GetFilesNetworkDriveAccessBlocked gets the filesNetworkDriveAccessBlocked property value. Indicates if devices can access files or other resources on a network server using the Server Message Block (SMB) protocol. Available for devices running iOS and iPadOS, versions 13.0 and later.
func (m *IosGeneralDeviceConfiguration) GetFilesNetworkDriveAccessBlocked()(*bool) {
    return m.filesNetworkDriveAccessBlocked
}
// GetFilesUsbDriveAccessBlocked gets the filesUsbDriveAccessBlocked property value. Indicates if sevices with access can connect to and open files on a USB drive. Available for devices running iOS and iPadOS, versions 13.0 and later.
func (m *IosGeneralDeviceConfiguration) GetFilesUsbDriveAccessBlocked()(*bool) {
    return m.filesUsbDriveAccessBlocked
}
// GetFindMyDeviceInFindMyAppBlocked gets the findMyDeviceInFindMyAppBlocked property value. Indicates whether or not to block Find My Device when the device is supervised (iOS 13 or later).
func (m *IosGeneralDeviceConfiguration) GetFindMyDeviceInFindMyAppBlocked()(*bool) {
    return m.findMyDeviceInFindMyAppBlocked
}
// GetFindMyFriendsBlocked gets the findMyFriendsBlocked property value. Indicates whether or not to block changes to Find My Friends when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) GetFindMyFriendsBlocked()(*bool) {
    return m.findMyFriendsBlocked
}
// GetFindMyFriendsInFindMyAppBlocked gets the findMyFriendsInFindMyAppBlocked property value. Indicates whether or not to block Find My Friends when the device is supervised (iOS 13 or later).
func (m *IosGeneralDeviceConfiguration) GetFindMyFriendsInFindMyAppBlocked()(*bool) {
    return m.findMyFriendsInFindMyAppBlocked
}
// GetGameCenterBlocked gets the gameCenterBlocked property value. Indicates whether or not to block the user from using Game Center when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) GetGameCenterBlocked()(*bool) {
    return m.gameCenterBlocked
}
// GetGamingBlockGameCenterFriends gets the gamingBlockGameCenterFriends property value. Indicates whether or not to block the user from having friends in Game Center. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) GetGamingBlockGameCenterFriends()(*bool) {
    return m.gamingBlockGameCenterFriends
}
// GetGamingBlockMultiplayer gets the gamingBlockMultiplayer property value. Indicates whether or not to block the user from using multiplayer gaming. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) GetGamingBlockMultiplayer()(*bool) {
    return m.gamingBlockMultiplayer
}
// GetHostPairingBlocked gets the hostPairingBlocked property value. indicates whether or not to allow host pairing to control the devices an iOS device can pair with when the iOS device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) GetHostPairingBlocked()(*bool) {
    return m.hostPairingBlocked
}
// GetIBooksStoreBlocked gets the iBooksStoreBlocked property value. Indicates whether or not to block the user from using the iBooks Store when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) GetIBooksStoreBlocked()(*bool) {
    return m.iBooksStoreBlocked
}
// GetIBooksStoreBlockErotica gets the iBooksStoreBlockErotica property value. Indicates whether or not to block the user from downloading media from the iBookstore that has been tagged as erotica.
func (m *IosGeneralDeviceConfiguration) GetIBooksStoreBlockErotica()(*bool) {
    return m.iBooksStoreBlockErotica
}
// GetICloudBlockActivityContinuation gets the iCloudBlockActivityContinuation property value. Indicates whether or not to block the user from continuing work they started on iOS device to another iOS or macOS device.
func (m *IosGeneralDeviceConfiguration) GetICloudBlockActivityContinuation()(*bool) {
    return m.iCloudBlockActivityContinuation
}
// GetICloudBlockBackup gets the iCloudBlockBackup property value. Indicates whether or not to block iCloud backup. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) GetICloudBlockBackup()(*bool) {
    return m.iCloudBlockBackup
}
// GetICloudBlockDocumentSync gets the iCloudBlockDocumentSync property value. Indicates whether or not to block iCloud document sync. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) GetICloudBlockDocumentSync()(*bool) {
    return m.iCloudBlockDocumentSync
}
// GetICloudBlockManagedAppsSync gets the iCloudBlockManagedAppsSync property value. Indicates whether or not to block Managed Apps Cloud Sync.
func (m *IosGeneralDeviceConfiguration) GetICloudBlockManagedAppsSync()(*bool) {
    return m.iCloudBlockManagedAppsSync
}
// GetICloudBlockPhotoLibrary gets the iCloudBlockPhotoLibrary property value. Indicates whether or not to block iCloud Photo Library.
func (m *IosGeneralDeviceConfiguration) GetICloudBlockPhotoLibrary()(*bool) {
    return m.iCloudBlockPhotoLibrary
}
// GetICloudBlockPhotoStreamSync gets the iCloudBlockPhotoStreamSync property value. Indicates whether or not to block iCloud Photo Stream Sync.
func (m *IosGeneralDeviceConfiguration) GetICloudBlockPhotoStreamSync()(*bool) {
    return m.iCloudBlockPhotoStreamSync
}
// GetICloudBlockSharedPhotoStream gets the iCloudBlockSharedPhotoStream property value. Indicates whether or not to block Shared Photo Stream.
func (m *IosGeneralDeviceConfiguration) GetICloudBlockSharedPhotoStream()(*bool) {
    return m.iCloudBlockSharedPhotoStream
}
// GetICloudPrivateRelayBlocked gets the iCloudPrivateRelayBlocked property value. iCloud private relay is an iCloud+ service that prevents networks and servers from monitoring a person's activity across the internet. By blocking iCloud private relay, Apple will not encrypt the traffic leaving the device. Available for devices running iOS 15 and later.
func (m *IosGeneralDeviceConfiguration) GetICloudPrivateRelayBlocked()(*bool) {
    return m.iCloudPrivateRelayBlocked
}
// GetICloudRequireEncryptedBackup gets the iCloudRequireEncryptedBackup property value. Indicates whether or not to require backups to iCloud be encrypted.
func (m *IosGeneralDeviceConfiguration) GetICloudRequireEncryptedBackup()(*bool) {
    return m.iCloudRequireEncryptedBackup
}
// GetITunesBlocked gets the iTunesBlocked property value. Indicates whether or not to block the iTunes app. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) GetITunesBlocked()(*bool) {
    return m.iTunesBlocked
}
// GetITunesBlockExplicitContent gets the iTunesBlockExplicitContent property value. Indicates whether or not to block the user from accessing explicit content in iTunes and the App Store. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) GetITunesBlockExplicitContent()(*bool) {
    return m.iTunesBlockExplicitContent
}
// GetITunesBlockMusicService gets the iTunesBlockMusicService property value. Indicates whether or not to block Music service and revert Music app to classic mode when the device is in supervised mode (iOS 9.3 and later and macOS 10.12 and later).
func (m *IosGeneralDeviceConfiguration) GetITunesBlockMusicService()(*bool) {
    return m.iTunesBlockMusicService
}
// GetITunesBlockRadio gets the iTunesBlockRadio property value. Indicates whether or not to block the user from using iTunes Radio when the device is in supervised mode (iOS 9.3 and later).
func (m *IosGeneralDeviceConfiguration) GetITunesBlockRadio()(*bool) {
    return m.iTunesBlockRadio
}
// GetKeyboardBlockAutoCorrect gets the keyboardBlockAutoCorrect property value. Indicates whether or not to block keyboard auto-correction when the device is in supervised mode (iOS 8.1.3 and later).
func (m *IosGeneralDeviceConfiguration) GetKeyboardBlockAutoCorrect()(*bool) {
    return m.keyboardBlockAutoCorrect
}
// GetKeyboardBlockDictation gets the keyboardBlockDictation property value. Indicates whether or not to block the user from using dictation input when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) GetKeyboardBlockDictation()(*bool) {
    return m.keyboardBlockDictation
}
// GetKeyboardBlockPredictive gets the keyboardBlockPredictive property value. Indicates whether or not to block predictive keyboards when device is in supervised mode (iOS 8.1.3 and later).
func (m *IosGeneralDeviceConfiguration) GetKeyboardBlockPredictive()(*bool) {
    return m.keyboardBlockPredictive
}
// GetKeyboardBlockShortcuts gets the keyboardBlockShortcuts property value. Indicates whether or not to block keyboard shortcuts when the device is in supervised mode (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) GetKeyboardBlockShortcuts()(*bool) {
    return m.keyboardBlockShortcuts
}
// GetKeyboardBlockSpellCheck gets the keyboardBlockSpellCheck property value. Indicates whether or not to block keyboard spell-checking when the device is in supervised mode (iOS 8.1.3 and later).
func (m *IosGeneralDeviceConfiguration) GetKeyboardBlockSpellCheck()(*bool) {
    return m.keyboardBlockSpellCheck
}
// GetKeychainBlockCloudSync gets the keychainBlockCloudSync property value. Indicates whether or not iCloud keychain synchronization is blocked. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) GetKeychainBlockCloudSync()(*bool) {
    return m.keychainBlockCloudSync
}
// GetKioskModeAllowAssistiveSpeak gets the kioskModeAllowAssistiveSpeak property value. Indicates whether or not to allow assistive speak while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) GetKioskModeAllowAssistiveSpeak()(*bool) {
    return m.kioskModeAllowAssistiveSpeak
}
// GetKioskModeAllowAssistiveTouchSettings gets the kioskModeAllowAssistiveTouchSettings property value. Indicates whether or not to allow access to the Assistive Touch Settings while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) GetKioskModeAllowAssistiveTouchSettings()(*bool) {
    return m.kioskModeAllowAssistiveTouchSettings
}
// GetKioskModeAllowAutoLock gets the kioskModeAllowAutoLock property value. Indicates whether or not to allow device auto lock while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockAutoLock instead.
func (m *IosGeneralDeviceConfiguration) GetKioskModeAllowAutoLock()(*bool) {
    return m.kioskModeAllowAutoLock
}
// GetKioskModeAllowColorInversionSettings gets the kioskModeAllowColorInversionSettings property value. Indicates whether or not to allow access to the Color Inversion Settings while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) GetKioskModeAllowColorInversionSettings()(*bool) {
    return m.kioskModeAllowColorInversionSettings
}
// GetKioskModeAllowRingerSwitch gets the kioskModeAllowRingerSwitch property value. Indicates whether or not to allow use of the ringer switch while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockRingerSwitch instead.
func (m *IosGeneralDeviceConfiguration) GetKioskModeAllowRingerSwitch()(*bool) {
    return m.kioskModeAllowRingerSwitch
}
// GetKioskModeAllowScreenRotation gets the kioskModeAllowScreenRotation property value. Indicates whether or not to allow screen rotation while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockScreenRotation instead.
func (m *IosGeneralDeviceConfiguration) GetKioskModeAllowScreenRotation()(*bool) {
    return m.kioskModeAllowScreenRotation
}
// GetKioskModeAllowSleepButton gets the kioskModeAllowSleepButton property value. Indicates whether or not to allow use of the sleep button while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockSleepButton instead.
func (m *IosGeneralDeviceConfiguration) GetKioskModeAllowSleepButton()(*bool) {
    return m.kioskModeAllowSleepButton
}
// GetKioskModeAllowTouchscreen gets the kioskModeAllowTouchscreen property value. Indicates whether or not to allow use of the touchscreen while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockTouchscreen instead.
func (m *IosGeneralDeviceConfiguration) GetKioskModeAllowTouchscreen()(*bool) {
    return m.kioskModeAllowTouchscreen
}
// GetKioskModeAllowVoiceControlModification gets the kioskModeAllowVoiceControlModification property value. Indicates whether or not to allow the user to toggle voice control in kiosk mode.
func (m *IosGeneralDeviceConfiguration) GetKioskModeAllowVoiceControlModification()(*bool) {
    return m.kioskModeAllowVoiceControlModification
}
// GetKioskModeAllowVoiceOverSettings gets the kioskModeAllowVoiceOverSettings property value. Indicates whether or not to allow access to the voice over settings while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) GetKioskModeAllowVoiceOverSettings()(*bool) {
    return m.kioskModeAllowVoiceOverSettings
}
// GetKioskModeAllowVolumeButtons gets the kioskModeAllowVolumeButtons property value. Indicates whether or not to allow use of the volume buttons while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockVolumeButtons instead.
func (m *IosGeneralDeviceConfiguration) GetKioskModeAllowVolumeButtons()(*bool) {
    return m.kioskModeAllowVolumeButtons
}
// GetKioskModeAllowZoomSettings gets the kioskModeAllowZoomSettings property value. Indicates whether or not to allow access to the zoom settings while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) GetKioskModeAllowZoomSettings()(*bool) {
    return m.kioskModeAllowZoomSettings
}
// GetKioskModeAppStoreUrl gets the kioskModeAppStoreUrl property value. URL in the app store to the app to use for kiosk mode. Use if KioskModeManagedAppId is not known.
func (m *IosGeneralDeviceConfiguration) GetKioskModeAppStoreUrl()(*string) {
    return m.kioskModeAppStoreUrl
}
// GetKioskModeAppType gets the kioskModeAppType property value. App source options for iOS kiosk mode.
func (m *IosGeneralDeviceConfiguration) GetKioskModeAppType()(*IosKioskModeAppType) {
    return m.kioskModeAppType
}
// GetKioskModeBlockAutoLock gets the kioskModeBlockAutoLock property value. Indicates whether or not to block device auto lock while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) GetKioskModeBlockAutoLock()(*bool) {
    return m.kioskModeBlockAutoLock
}
// GetKioskModeBlockRingerSwitch gets the kioskModeBlockRingerSwitch property value. Indicates whether or not to block use of the ringer switch while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) GetKioskModeBlockRingerSwitch()(*bool) {
    return m.kioskModeBlockRingerSwitch
}
// GetKioskModeBlockScreenRotation gets the kioskModeBlockScreenRotation property value. Indicates whether or not to block screen rotation while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) GetKioskModeBlockScreenRotation()(*bool) {
    return m.kioskModeBlockScreenRotation
}
// GetKioskModeBlockSleepButton gets the kioskModeBlockSleepButton property value. Indicates whether or not to block use of the sleep button while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) GetKioskModeBlockSleepButton()(*bool) {
    return m.kioskModeBlockSleepButton
}
// GetKioskModeBlockTouchscreen gets the kioskModeBlockTouchscreen property value. Indicates whether or not to block use of the touchscreen while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) GetKioskModeBlockTouchscreen()(*bool) {
    return m.kioskModeBlockTouchscreen
}
// GetKioskModeBlockVolumeButtons gets the kioskModeBlockVolumeButtons property value. Indicates whether or not to block the volume buttons while in Kiosk Mode.
func (m *IosGeneralDeviceConfiguration) GetKioskModeBlockVolumeButtons()(*bool) {
    return m.kioskModeBlockVolumeButtons
}
// GetKioskModeBuiltInAppId gets the kioskModeBuiltInAppId property value. ID for built-in apps to use for kiosk mode. Used when KioskModeManagedAppId and KioskModeAppStoreUrl are not set.
func (m *IosGeneralDeviceConfiguration) GetKioskModeBuiltInAppId()(*string) {
    return m.kioskModeBuiltInAppId
}
// GetKioskModeEnableVoiceControl gets the kioskModeEnableVoiceControl property value. Indicates whether or not to enable voice control in kiosk mode.
func (m *IosGeneralDeviceConfiguration) GetKioskModeEnableVoiceControl()(*bool) {
    return m.kioskModeEnableVoiceControl
}
// GetKioskModeManagedAppId gets the kioskModeManagedAppId property value. Managed app id of the app to use for kiosk mode. If KioskModeManagedAppId is specified then KioskModeAppStoreUrl will be ignored.
func (m *IosGeneralDeviceConfiguration) GetKioskModeManagedAppId()(*string) {
    return m.kioskModeManagedAppId
}
// GetKioskModeRequireAssistiveTouch gets the kioskModeRequireAssistiveTouch property value. Indicates whether or not to require assistive touch while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) GetKioskModeRequireAssistiveTouch()(*bool) {
    return m.kioskModeRequireAssistiveTouch
}
// GetKioskModeRequireColorInversion gets the kioskModeRequireColorInversion property value. Indicates whether or not to require color inversion while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) GetKioskModeRequireColorInversion()(*bool) {
    return m.kioskModeRequireColorInversion
}
// GetKioskModeRequireMonoAudio gets the kioskModeRequireMonoAudio property value. Indicates whether or not to require mono audio while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) GetKioskModeRequireMonoAudio()(*bool) {
    return m.kioskModeRequireMonoAudio
}
// GetKioskModeRequireVoiceOver gets the kioskModeRequireVoiceOver property value. Indicates whether or not to require voice over while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) GetKioskModeRequireVoiceOver()(*bool) {
    return m.kioskModeRequireVoiceOver
}
// GetKioskModeRequireZoom gets the kioskModeRequireZoom property value. Indicates whether or not to require zoom while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) GetKioskModeRequireZoom()(*bool) {
    return m.kioskModeRequireZoom
}
// GetLockScreenBlockControlCenter gets the lockScreenBlockControlCenter property value. Indicates whether or not to block the user from using control center on the lock screen.
func (m *IosGeneralDeviceConfiguration) GetLockScreenBlockControlCenter()(*bool) {
    return m.lockScreenBlockControlCenter
}
// GetLockScreenBlockNotificationView gets the lockScreenBlockNotificationView property value. Indicates whether or not to block the user from using the notification view on the lock screen.
func (m *IosGeneralDeviceConfiguration) GetLockScreenBlockNotificationView()(*bool) {
    return m.lockScreenBlockNotificationView
}
// GetLockScreenBlockPassbook gets the lockScreenBlockPassbook property value. Indicates whether or not to block the user from using passbook when the device is locked.
func (m *IosGeneralDeviceConfiguration) GetLockScreenBlockPassbook()(*bool) {
    return m.lockScreenBlockPassbook
}
// GetLockScreenBlockTodayView gets the lockScreenBlockTodayView property value. Indicates whether or not to block the user from using the Today View on the lock screen.
func (m *IosGeneralDeviceConfiguration) GetLockScreenBlockTodayView()(*bool) {
    return m.lockScreenBlockTodayView
}
// GetManagedPasteboardRequired gets the managedPasteboardRequired property value. Open-in management controls how people share data between unmanaged and managed apps. Setting this to true enforces copy/paste restrictions based on how you configured Block viewing corporate documents in unmanaged apps  and  Block viewing non-corporate documents in corporate apps.
func (m *IosGeneralDeviceConfiguration) GetManagedPasteboardRequired()(*bool) {
    return m.managedPasteboardRequired
}
// GetMediaContentRatingApps gets the mediaContentRatingApps property value. Apps rating as in media content
func (m *IosGeneralDeviceConfiguration) GetMediaContentRatingApps()(*RatingAppsType) {
    return m.mediaContentRatingApps
}
// GetMediaContentRatingAustralia gets the mediaContentRatingAustralia property value. Media content rating settings for Australia
func (m *IosGeneralDeviceConfiguration) GetMediaContentRatingAustralia()(MediaContentRatingAustraliaable) {
    return m.mediaContentRatingAustralia
}
// GetMediaContentRatingCanada gets the mediaContentRatingCanada property value. Media content rating settings for Canada
func (m *IosGeneralDeviceConfiguration) GetMediaContentRatingCanada()(MediaContentRatingCanadaable) {
    return m.mediaContentRatingCanada
}
// GetMediaContentRatingFrance gets the mediaContentRatingFrance property value. Media content rating settings for France
func (m *IosGeneralDeviceConfiguration) GetMediaContentRatingFrance()(MediaContentRatingFranceable) {
    return m.mediaContentRatingFrance
}
// GetMediaContentRatingGermany gets the mediaContentRatingGermany property value. Media content rating settings for Germany
func (m *IosGeneralDeviceConfiguration) GetMediaContentRatingGermany()(MediaContentRatingGermanyable) {
    return m.mediaContentRatingGermany
}
// GetMediaContentRatingIreland gets the mediaContentRatingIreland property value. Media content rating settings for Ireland
func (m *IosGeneralDeviceConfiguration) GetMediaContentRatingIreland()(MediaContentRatingIrelandable) {
    return m.mediaContentRatingIreland
}
// GetMediaContentRatingJapan gets the mediaContentRatingJapan property value. Media content rating settings for Japan
func (m *IosGeneralDeviceConfiguration) GetMediaContentRatingJapan()(MediaContentRatingJapanable) {
    return m.mediaContentRatingJapan
}
// GetMediaContentRatingNewZealand gets the mediaContentRatingNewZealand property value. Media content rating settings for New Zealand
func (m *IosGeneralDeviceConfiguration) GetMediaContentRatingNewZealand()(MediaContentRatingNewZealandable) {
    return m.mediaContentRatingNewZealand
}
// GetMediaContentRatingUnitedKingdom gets the mediaContentRatingUnitedKingdom property value. Media content rating settings for United Kingdom
func (m *IosGeneralDeviceConfiguration) GetMediaContentRatingUnitedKingdom()(MediaContentRatingUnitedKingdomable) {
    return m.mediaContentRatingUnitedKingdom
}
// GetMediaContentRatingUnitedStates gets the mediaContentRatingUnitedStates property value. Media content rating settings for United States
func (m *IosGeneralDeviceConfiguration) GetMediaContentRatingUnitedStates()(MediaContentRatingUnitedStatesable) {
    return m.mediaContentRatingUnitedStates
}
// GetMessagesBlocked gets the messagesBlocked property value. Indicates whether or not to block the user from using the Messages app on the supervised device.
func (m *IosGeneralDeviceConfiguration) GetMessagesBlocked()(*bool) {
    return m.messagesBlocked
}
// GetNetworkUsageRules gets the networkUsageRules property value. List of managed apps and the network rules that applies to them. This collection can contain a maximum of 1000 elements.
func (m *IosGeneralDeviceConfiguration) GetNetworkUsageRules()([]IosNetworkUsageRuleable) {
    return m.networkUsageRules
}
// GetNfcBlocked gets the nfcBlocked property value. Disable NFC to prevent devices from pairing with other NFC-enabled devices. Available for iOS/iPadOS devices running 14.2 and later.
func (m *IosGeneralDeviceConfiguration) GetNfcBlocked()(*bool) {
    return m.nfcBlocked
}
// GetNotificationsBlockSettingsModification gets the notificationsBlockSettingsModification property value. Indicates whether or not to allow notifications settings modification (iOS 9.3 and later).
func (m *IosGeneralDeviceConfiguration) GetNotificationsBlockSettingsModification()(*bool) {
    return m.notificationsBlockSettingsModification
}
// GetOnDeviceOnlyDictationForced gets the onDeviceOnlyDictationForced property value. Disables connections to Siri servers so that users can’t use Siri to dictate text. Available for devices running iOS and iPadOS versions 14.5 and later.
func (m *IosGeneralDeviceConfiguration) GetOnDeviceOnlyDictationForced()(*bool) {
    return m.onDeviceOnlyDictationForced
}
// GetOnDeviceOnlyTranslationForced gets the onDeviceOnlyTranslationForced property value. When set to TRUE, the setting disables connections to Siri servers so that users can’t use Siri to translate text. When set to FALSE, the setting allows connections to to Siri servers to users can use Siri to translate text. Available for devices running iOS and iPadOS versions 15.0 and later.
func (m *IosGeneralDeviceConfiguration) GetOnDeviceOnlyTranslationForced()(*bool) {
    return m.onDeviceOnlyTranslationForced
}
// GetPasscodeBlockFingerprintModification gets the passcodeBlockFingerprintModification property value. Block modification of registered Touch ID fingerprints when in supervised mode.
func (m *IosGeneralDeviceConfiguration) GetPasscodeBlockFingerprintModification()(*bool) {
    return m.passcodeBlockFingerprintModification
}
// GetPasscodeBlockFingerprintUnlock gets the passcodeBlockFingerprintUnlock property value. Indicates whether or not to block fingerprint unlock.
func (m *IosGeneralDeviceConfiguration) GetPasscodeBlockFingerprintUnlock()(*bool) {
    return m.passcodeBlockFingerprintUnlock
}
// GetPasscodeBlockModification gets the passcodeBlockModification property value. Indicates whether or not to allow passcode modification on the supervised device (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) GetPasscodeBlockModification()(*bool) {
    return m.passcodeBlockModification
}
// GetPasscodeBlockSimple gets the passcodeBlockSimple property value. Indicates whether or not to block simple passcodes.
func (m *IosGeneralDeviceConfiguration) GetPasscodeBlockSimple()(*bool) {
    return m.passcodeBlockSimple
}
// GetPasscodeExpirationDays gets the passcodeExpirationDays property value. Number of days before the passcode expires. Valid values 1 to 65535
func (m *IosGeneralDeviceConfiguration) GetPasscodeExpirationDays()(*int32) {
    return m.passcodeExpirationDays
}
// GetPasscodeMinimumCharacterSetCount gets the passcodeMinimumCharacterSetCount property value. Number of character sets a passcode must contain. Valid values 0 to 4
func (m *IosGeneralDeviceConfiguration) GetPasscodeMinimumCharacterSetCount()(*int32) {
    return m.passcodeMinimumCharacterSetCount
}
// GetPasscodeMinimumLength gets the passcodeMinimumLength property value. Minimum length of passcode. Valid values 4 to 14
func (m *IosGeneralDeviceConfiguration) GetPasscodeMinimumLength()(*int32) {
    return m.passcodeMinimumLength
}
// GetPasscodeMinutesOfInactivityBeforeLock gets the passcodeMinutesOfInactivityBeforeLock property value. Minutes of inactivity before a passcode is required.
func (m *IosGeneralDeviceConfiguration) GetPasscodeMinutesOfInactivityBeforeLock()(*int32) {
    return m.passcodeMinutesOfInactivityBeforeLock
}
// GetPasscodeMinutesOfInactivityBeforeScreenTimeout gets the passcodeMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before the screen times out.
func (m *IosGeneralDeviceConfiguration) GetPasscodeMinutesOfInactivityBeforeScreenTimeout()(*int32) {
    return m.passcodeMinutesOfInactivityBeforeScreenTimeout
}
// GetPasscodePreviousPasscodeBlockCount gets the passcodePreviousPasscodeBlockCount property value. Number of previous passcodes to block. Valid values 1 to 24
func (m *IosGeneralDeviceConfiguration) GetPasscodePreviousPasscodeBlockCount()(*int32) {
    return m.passcodePreviousPasscodeBlockCount
}
// GetPasscodeRequired gets the passcodeRequired property value. Indicates whether or not to require a passcode.
func (m *IosGeneralDeviceConfiguration) GetPasscodeRequired()(*bool) {
    return m.passcodeRequired
}
// GetPasscodeRequiredType gets the passcodeRequiredType property value. Possible values of required passwords.
func (m *IosGeneralDeviceConfiguration) GetPasscodeRequiredType()(*RequiredPasswordType) {
    return m.passcodeRequiredType
}
// GetPasscodeSignInFailureCountBeforeWipe gets the passcodeSignInFailureCountBeforeWipe property value. Number of sign in failures allowed before wiping the device. Valid values 2 to 11
func (m *IosGeneralDeviceConfiguration) GetPasscodeSignInFailureCountBeforeWipe()(*int32) {
    return m.passcodeSignInFailureCountBeforeWipe
}
// GetPasswordBlockAirDropSharing gets the passwordBlockAirDropSharing property value. Indicates whether or not to block sharing passwords with the AirDrop passwords feature iOS 12.0 and later).
func (m *IosGeneralDeviceConfiguration) GetPasswordBlockAirDropSharing()(*bool) {
    return m.passwordBlockAirDropSharing
}
// GetPasswordBlockAutoFill gets the passwordBlockAutoFill property value. Indicates if the AutoFill passwords feature is allowed (iOS 12.0 and later).
func (m *IosGeneralDeviceConfiguration) GetPasswordBlockAutoFill()(*bool) {
    return m.passwordBlockAutoFill
}
// GetPasswordBlockProximityRequests gets the passwordBlockProximityRequests property value. Indicates whether or not to block requesting passwords from nearby devices (iOS 12.0 and later).
func (m *IosGeneralDeviceConfiguration) GetPasswordBlockProximityRequests()(*bool) {
    return m.passwordBlockProximityRequests
}
// GetPkiBlockOTAUpdates gets the pkiBlockOTAUpdates property value. Indicates whether or not over-the-air PKI updates are blocked. Setting this restriction to false does not disable CRL and OCSP checks (iOS 7.0 and later).
func (m *IosGeneralDeviceConfiguration) GetPkiBlockOTAUpdates()(*bool) {
    return m.pkiBlockOTAUpdates
}
// GetPodcastsBlocked gets the podcastsBlocked property value. Indicates whether or not to block the user from using podcasts on the supervised device (iOS 8.0 and later).
func (m *IosGeneralDeviceConfiguration) GetPodcastsBlocked()(*bool) {
    return m.podcastsBlocked
}
// GetPrivacyForceLimitAdTracking gets the privacyForceLimitAdTracking property value. Indicates if ad tracking is limited.(iOS 7.0 and later).
func (m *IosGeneralDeviceConfiguration) GetPrivacyForceLimitAdTracking()(*bool) {
    return m.privacyForceLimitAdTracking
}
// GetProximityBlockSetupToNewDevice gets the proximityBlockSetupToNewDevice property value. Indicates whether or not to enable the prompt to setup nearby devices with a supervised device.
func (m *IosGeneralDeviceConfiguration) GetProximityBlockSetupToNewDevice()(*bool) {
    return m.proximityBlockSetupToNewDevice
}
// GetSafariBlockAutofill gets the safariBlockAutofill property value. Indicates whether or not to block the user from using Auto fill in Safari. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) GetSafariBlockAutofill()(*bool) {
    return m.safariBlockAutofill
}
// GetSafariBlocked gets the safariBlocked property value. Indicates whether or not to block the user from using Safari. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) GetSafariBlocked()(*bool) {
    return m.safariBlocked
}
// GetSafariBlockJavaScript gets the safariBlockJavaScript property value. Indicates whether or not to block JavaScript in Safari.
func (m *IosGeneralDeviceConfiguration) GetSafariBlockJavaScript()(*bool) {
    return m.safariBlockJavaScript
}
// GetSafariBlockPopups gets the safariBlockPopups property value. Indicates whether or not to block popups in Safari.
func (m *IosGeneralDeviceConfiguration) GetSafariBlockPopups()(*bool) {
    return m.safariBlockPopups
}
// GetSafariCookieSettings gets the safariCookieSettings property value. Web Browser Cookie Settings.
func (m *IosGeneralDeviceConfiguration) GetSafariCookieSettings()(*WebBrowserCookieSettings) {
    return m.safariCookieSettings
}
// GetSafariManagedDomains gets the safariManagedDomains property value. URLs matching the patterns listed here will be considered managed.
func (m *IosGeneralDeviceConfiguration) GetSafariManagedDomains()([]string) {
    return m.safariManagedDomains
}
// GetSafariPasswordAutoFillDomains gets the safariPasswordAutoFillDomains property value. Users can save passwords in Safari only from URLs matching the patterns listed here. Applies to devices in supervised mode (iOS 9.3 and later).
func (m *IosGeneralDeviceConfiguration) GetSafariPasswordAutoFillDomains()([]string) {
    return m.safariPasswordAutoFillDomains
}
// GetSafariRequireFraudWarning gets the safariRequireFraudWarning property value. Indicates whether or not to require fraud warning in Safari.
func (m *IosGeneralDeviceConfiguration) GetSafariRequireFraudWarning()(*bool) {
    return m.safariRequireFraudWarning
}
// GetScreenCaptureBlocked gets the screenCaptureBlocked property value. Indicates whether or not to block the user from taking Screenshots.
func (m *IosGeneralDeviceConfiguration) GetScreenCaptureBlocked()(*bool) {
    return m.screenCaptureBlocked
}
// GetSharedDeviceBlockTemporarySessions gets the sharedDeviceBlockTemporarySessions property value. Indicates whether or not to block temporary sessions on Shared iPads (iOS 13.4 or later).
func (m *IosGeneralDeviceConfiguration) GetSharedDeviceBlockTemporarySessions()(*bool) {
    return m.sharedDeviceBlockTemporarySessions
}
// GetSiriBlocked gets the siriBlocked property value. Indicates whether or not to block the user from using Siri.
func (m *IosGeneralDeviceConfiguration) GetSiriBlocked()(*bool) {
    return m.siriBlocked
}
// GetSiriBlockedWhenLocked gets the siriBlockedWhenLocked property value. Indicates whether or not to block the user from using Siri when locked.
func (m *IosGeneralDeviceConfiguration) GetSiriBlockedWhenLocked()(*bool) {
    return m.siriBlockedWhenLocked
}
// GetSiriBlockUserGeneratedContent gets the siriBlockUserGeneratedContent property value. Indicates whether or not to block Siri from querying user-generated content when used on a supervised device.
func (m *IosGeneralDeviceConfiguration) GetSiriBlockUserGeneratedContent()(*bool) {
    return m.siriBlockUserGeneratedContent
}
// GetSiriRequireProfanityFilter gets the siriRequireProfanityFilter property value. Indicates whether or not to prevent Siri from dictating, or speaking profane language on supervised device.
func (m *IosGeneralDeviceConfiguration) GetSiriRequireProfanityFilter()(*bool) {
    return m.siriRequireProfanityFilter
}
// GetSoftwareUpdatesEnforcedDelayInDays gets the softwareUpdatesEnforcedDelayInDays property value. Sets how many days a software update will be delyed for a supervised device. Valid values 0 to 90
func (m *IosGeneralDeviceConfiguration) GetSoftwareUpdatesEnforcedDelayInDays()(*int32) {
    return m.softwareUpdatesEnforcedDelayInDays
}
// GetSoftwareUpdatesForceDelayed gets the softwareUpdatesForceDelayed property value. Indicates whether or not to delay user visibility of software updates when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) GetSoftwareUpdatesForceDelayed()(*bool) {
    return m.softwareUpdatesForceDelayed
}
// GetSpotlightBlockInternetResults gets the spotlightBlockInternetResults property value. Indicates whether or not to block Spotlight search from returning internet results on supervised device.
func (m *IosGeneralDeviceConfiguration) GetSpotlightBlockInternetResults()(*bool) {
    return m.spotlightBlockInternetResults
}
// GetUnpairedExternalBootToRecoveryAllowed gets the unpairedExternalBootToRecoveryAllowed property value. Allow users to boot devices into recovery mode with unpaired devices. Available for devices running iOS and iPadOS versions 14.5 and later.
func (m *IosGeneralDeviceConfiguration) GetUnpairedExternalBootToRecoveryAllowed()(*bool) {
    return m.unpairedExternalBootToRecoveryAllowed
}
// GetUsbRestrictedModeBlocked gets the usbRestrictedModeBlocked property value. Indicates if connecting to USB accessories while the device is locked is allowed (iOS 11.4.1 and later).
func (m *IosGeneralDeviceConfiguration) GetUsbRestrictedModeBlocked()(*bool) {
    return m.usbRestrictedModeBlocked
}
// GetVoiceDialingBlocked gets the voiceDialingBlocked property value. Indicates whether or not to block voice dialing.
func (m *IosGeneralDeviceConfiguration) GetVoiceDialingBlocked()(*bool) {
    return m.voiceDialingBlocked
}
// GetVpnBlockCreation gets the vpnBlockCreation property value. Indicates whether or not the creation of VPN configurations is blocked (iOS 11.0 and later).
func (m *IosGeneralDeviceConfiguration) GetVpnBlockCreation()(*bool) {
    return m.vpnBlockCreation
}
// GetWallpaperBlockModification gets the wallpaperBlockModification property value. Indicates whether or not to allow wallpaper modification on supervised device (iOS 9.0 and later) .
func (m *IosGeneralDeviceConfiguration) GetWallpaperBlockModification()(*bool) {
    return m.wallpaperBlockModification
}
// GetWiFiConnectOnlyToConfiguredNetworks gets the wiFiConnectOnlyToConfiguredNetworks property value. Indicates whether or not to force the device to use only Wi-Fi networks from configuration profiles when the device is in supervised mode. Available for devices running iOS and iPadOS versions 14.4 and earlier. Devices running 14.5+ should use the setting, 'WiFiConnectToAllowedNetworksOnlyForced.
func (m *IosGeneralDeviceConfiguration) GetWiFiConnectOnlyToConfiguredNetworks()(*bool) {
    return m.wiFiConnectOnlyToConfiguredNetworks
}
// GetWiFiConnectToAllowedNetworksOnlyForced gets the wiFiConnectToAllowedNetworksOnlyForced property value. Require devices to use Wi-Fi networks set up via configuration profiles. Available for devices running iOS and iPadOS versions 14.5 and later.
func (m *IosGeneralDeviceConfiguration) GetWiFiConnectToAllowedNetworksOnlyForced()(*bool) {
    return m.wiFiConnectToAllowedNetworksOnlyForced
}
// GetWifiPowerOnForced gets the wifiPowerOnForced property value. Indicates whether or not Wi-Fi remains on, even when device is in airplane mode. Available for devices running iOS and iPadOS, versions 13.0 and later.
func (m *IosGeneralDeviceConfiguration) GetWifiPowerOnForced()(*bool) {
    return m.wifiPowerOnForced
}
// Serialize serializes information the current object
func (m *IosGeneralDeviceConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("accountBlockModification", m.GetAccountBlockModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("activationLockAllowWhenSupervised", m.GetActivationLockAllowWhenSupervised())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("airDropBlocked", m.GetAirDropBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("airDropForceUnmanagedDropTarget", m.GetAirDropForceUnmanagedDropTarget())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("airPlayForcePairingPasswordForOutgoingRequests", m.GetAirPlayForcePairingPasswordForOutgoingRequests())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("airPrintBlockCredentialsStorage", m.GetAirPrintBlockCredentialsStorage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("airPrintBlocked", m.GetAirPrintBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("airPrintBlockiBeaconDiscovery", m.GetAirPrintBlockiBeaconDiscovery())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("airPrintForceTrustedTLS", m.GetAirPrintForceTrustedTLS())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("appClipsBlocked", m.GetAppClipsBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("appleNewsBlocked", m.GetAppleNewsBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applePersonalizedAdsBlocked", m.GetApplePersonalizedAdsBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("appleWatchBlockPairing", m.GetAppleWatchBlockPairing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("appleWatchForceWristDetection", m.GetAppleWatchForceWristDetection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("appRemovalBlocked", m.GetAppRemovalBlocked())
        if err != nil {
            return err
        }
    }
    if m.GetAppsSingleAppModeList() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppsSingleAppModeList()))
        for i, v := range m.GetAppsSingleAppModeList() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("appsSingleAppModeList", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("appStoreBlockAutomaticDownloads", m.GetAppStoreBlockAutomaticDownloads())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("appStoreBlocked", m.GetAppStoreBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("appStoreBlockInAppPurchases", m.GetAppStoreBlockInAppPurchases())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("appStoreBlockUIAppInstallation", m.GetAppStoreBlockUIAppInstallation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("appStoreRequirePassword", m.GetAppStoreRequirePassword())
        if err != nil {
            return err
        }
    }
    if m.GetAppsVisibilityList() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppsVisibilityList()))
        for i, v := range m.GetAppsVisibilityList() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("appsVisibilityList", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppsVisibilityListType() != nil {
        cast := (*m.GetAppsVisibilityListType()).String()
        err = writer.WriteStringValue("appsVisibilityListType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("autoFillForceAuthentication", m.GetAutoFillForceAuthentication())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("autoUnlockBlocked", m.GetAutoUnlockBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("blockSystemAppRemoval", m.GetBlockSystemAppRemoval())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("bluetoothBlockModification", m.GetBluetoothBlockModification())
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
        err = writer.WriteBoolValue("cellularBlockDataRoaming", m.GetCellularBlockDataRoaming())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cellularBlockGlobalBackgroundFetchWhileRoaming", m.GetCellularBlockGlobalBackgroundFetchWhileRoaming())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cellularBlockPerAppDataModification", m.GetCellularBlockPerAppDataModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cellularBlockPersonalHotspot", m.GetCellularBlockPersonalHotspot())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cellularBlockPersonalHotspotModification", m.GetCellularBlockPersonalHotspotModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cellularBlockPlanModification", m.GetCellularBlockPlanModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cellularBlockVoiceRoaming", m.GetCellularBlockVoiceRoaming())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("certificatesBlockUntrustedTlsCertificates", m.GetCertificatesBlockUntrustedTlsCertificates())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("classroomAppBlockRemoteScreenObservation", m.GetClassroomAppBlockRemoteScreenObservation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("classroomAppForceUnpromptedScreenObservation", m.GetClassroomAppForceUnpromptedScreenObservation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("classroomForceAutomaticallyJoinClasses", m.GetClassroomForceAutomaticallyJoinClasses())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("classroomForceRequestPermissionToLeaveClasses", m.GetClassroomForceRequestPermissionToLeaveClasses())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("classroomForceUnpromptedAppAndDeviceLock", m.GetClassroomForceUnpromptedAppAndDeviceLock())
        if err != nil {
            return err
        }
    }
    if m.GetCompliantAppListType() != nil {
        cast := (*m.GetCompliantAppListType()).String()
        err = writer.WriteStringValue("compliantAppListType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCompliantAppsList() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCompliantAppsList()))
        for i, v := range m.GetCompliantAppsList() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("compliantAppsList", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("configurationProfileBlockChanges", m.GetConfigurationProfileBlockChanges())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("contactsAllowManagedToUnmanagedWrite", m.GetContactsAllowManagedToUnmanagedWrite())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("contactsAllowUnmanagedToManagedRead", m.GetContactsAllowUnmanagedToManagedRead())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("continuousPathKeyboardBlocked", m.GetContinuousPathKeyboardBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("dateAndTimeForceSetAutomatically", m.GetDateAndTimeForceSetAutomatically())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("definitionLookupBlocked", m.GetDefinitionLookupBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deviceBlockEnableRestrictions", m.GetDeviceBlockEnableRestrictions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deviceBlockEraseContentAndSettings", m.GetDeviceBlockEraseContentAndSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deviceBlockNameModification", m.GetDeviceBlockNameModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("diagnosticDataBlockSubmission", m.GetDiagnosticDataBlockSubmission())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("diagnosticDataBlockSubmissionModification", m.GetDiagnosticDataBlockSubmissionModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("documentsBlockManagedDocumentsInUnmanagedApps", m.GetDocumentsBlockManagedDocumentsInUnmanagedApps())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("documentsBlockUnmanagedDocumentsInManagedApps", m.GetDocumentsBlockUnmanagedDocumentsInManagedApps())
        if err != nil {
            return err
        }
    }
    if m.GetEmailInDomainSuffixes() != nil {
        err = writer.WriteCollectionOfStringValues("emailInDomainSuffixes", m.GetEmailInDomainSuffixes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enterpriseAppBlockTrust", m.GetEnterpriseAppBlockTrust())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enterpriseAppBlockTrustModification", m.GetEnterpriseAppBlockTrustModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enterpriseBookBlockBackup", m.GetEnterpriseBookBlockBackup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enterpriseBookBlockMetadataSync", m.GetEnterpriseBookBlockMetadataSync())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("esimBlockModification", m.GetEsimBlockModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("faceTimeBlocked", m.GetFaceTimeBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("filesNetworkDriveAccessBlocked", m.GetFilesNetworkDriveAccessBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("filesUsbDriveAccessBlocked", m.GetFilesUsbDriveAccessBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("findMyDeviceInFindMyAppBlocked", m.GetFindMyDeviceInFindMyAppBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("findMyFriendsBlocked", m.GetFindMyFriendsBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("findMyFriendsInFindMyAppBlocked", m.GetFindMyFriendsInFindMyAppBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("gameCenterBlocked", m.GetGameCenterBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("gamingBlockGameCenterFriends", m.GetGamingBlockGameCenterFriends())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("gamingBlockMultiplayer", m.GetGamingBlockMultiplayer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hostPairingBlocked", m.GetHostPairingBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iBooksStoreBlocked", m.GetIBooksStoreBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iBooksStoreBlockErotica", m.GetIBooksStoreBlockErotica())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudBlockActivityContinuation", m.GetICloudBlockActivityContinuation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudBlockBackup", m.GetICloudBlockBackup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudBlockDocumentSync", m.GetICloudBlockDocumentSync())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudBlockManagedAppsSync", m.GetICloudBlockManagedAppsSync())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudBlockPhotoLibrary", m.GetICloudBlockPhotoLibrary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudBlockPhotoStreamSync", m.GetICloudBlockPhotoStreamSync())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudBlockSharedPhotoStream", m.GetICloudBlockSharedPhotoStream())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudPrivateRelayBlocked", m.GetICloudPrivateRelayBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iCloudRequireEncryptedBackup", m.GetICloudRequireEncryptedBackup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iTunesBlocked", m.GetITunesBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iTunesBlockExplicitContent", m.GetITunesBlockExplicitContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iTunesBlockMusicService", m.GetITunesBlockMusicService())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iTunesBlockRadio", m.GetITunesBlockRadio())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("keyboardBlockAutoCorrect", m.GetKeyboardBlockAutoCorrect())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("keyboardBlockDictation", m.GetKeyboardBlockDictation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("keyboardBlockPredictive", m.GetKeyboardBlockPredictive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("keyboardBlockShortcuts", m.GetKeyboardBlockShortcuts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("keyboardBlockSpellCheck", m.GetKeyboardBlockSpellCheck())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("keychainBlockCloudSync", m.GetKeychainBlockCloudSync())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeAllowAssistiveSpeak", m.GetKioskModeAllowAssistiveSpeak())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeAllowAssistiveTouchSettings", m.GetKioskModeAllowAssistiveTouchSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeAllowAutoLock", m.GetKioskModeAllowAutoLock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeAllowColorInversionSettings", m.GetKioskModeAllowColorInversionSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeAllowRingerSwitch", m.GetKioskModeAllowRingerSwitch())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeAllowScreenRotation", m.GetKioskModeAllowScreenRotation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeAllowSleepButton", m.GetKioskModeAllowSleepButton())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeAllowTouchscreen", m.GetKioskModeAllowTouchscreen())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeAllowVoiceControlModification", m.GetKioskModeAllowVoiceControlModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeAllowVoiceOverSettings", m.GetKioskModeAllowVoiceOverSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeAllowVolumeButtons", m.GetKioskModeAllowVolumeButtons())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeAllowZoomSettings", m.GetKioskModeAllowZoomSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("kioskModeAppStoreUrl", m.GetKioskModeAppStoreUrl())
        if err != nil {
            return err
        }
    }
    if m.GetKioskModeAppType() != nil {
        cast := (*m.GetKioskModeAppType()).String()
        err = writer.WriteStringValue("kioskModeAppType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeBlockAutoLock", m.GetKioskModeBlockAutoLock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeBlockRingerSwitch", m.GetKioskModeBlockRingerSwitch())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeBlockScreenRotation", m.GetKioskModeBlockScreenRotation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeBlockSleepButton", m.GetKioskModeBlockSleepButton())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeBlockTouchscreen", m.GetKioskModeBlockTouchscreen())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeBlockVolumeButtons", m.GetKioskModeBlockVolumeButtons())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("kioskModeBuiltInAppId", m.GetKioskModeBuiltInAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeEnableVoiceControl", m.GetKioskModeEnableVoiceControl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("kioskModeManagedAppId", m.GetKioskModeManagedAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeRequireAssistiveTouch", m.GetKioskModeRequireAssistiveTouch())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeRequireColorInversion", m.GetKioskModeRequireColorInversion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeRequireMonoAudio", m.GetKioskModeRequireMonoAudio())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeRequireVoiceOver", m.GetKioskModeRequireVoiceOver())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kioskModeRequireZoom", m.GetKioskModeRequireZoom())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("lockScreenBlockControlCenter", m.GetLockScreenBlockControlCenter())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("lockScreenBlockNotificationView", m.GetLockScreenBlockNotificationView())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("lockScreenBlockPassbook", m.GetLockScreenBlockPassbook())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("lockScreenBlockTodayView", m.GetLockScreenBlockTodayView())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("managedPasteboardRequired", m.GetManagedPasteboardRequired())
        if err != nil {
            return err
        }
    }
    if m.GetMediaContentRatingApps() != nil {
        cast := (*m.GetMediaContentRatingApps()).String()
        err = writer.WriteStringValue("mediaContentRatingApps", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mediaContentRatingAustralia", m.GetMediaContentRatingAustralia())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mediaContentRatingCanada", m.GetMediaContentRatingCanada())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mediaContentRatingFrance", m.GetMediaContentRatingFrance())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mediaContentRatingGermany", m.GetMediaContentRatingGermany())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mediaContentRatingIreland", m.GetMediaContentRatingIreland())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mediaContentRatingJapan", m.GetMediaContentRatingJapan())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mediaContentRatingNewZealand", m.GetMediaContentRatingNewZealand())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mediaContentRatingUnitedKingdom", m.GetMediaContentRatingUnitedKingdom())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mediaContentRatingUnitedStates", m.GetMediaContentRatingUnitedStates())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("messagesBlocked", m.GetMessagesBlocked())
        if err != nil {
            return err
        }
    }
    if m.GetNetworkUsageRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNetworkUsageRules()))
        for i, v := range m.GetNetworkUsageRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("networkUsageRules", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("nfcBlocked", m.GetNfcBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("notificationsBlockSettingsModification", m.GetNotificationsBlockSettingsModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("onDeviceOnlyDictationForced", m.GetOnDeviceOnlyDictationForced())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("onDeviceOnlyTranslationForced", m.GetOnDeviceOnlyTranslationForced())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passcodeBlockFingerprintModification", m.GetPasscodeBlockFingerprintModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passcodeBlockFingerprintUnlock", m.GetPasscodeBlockFingerprintUnlock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passcodeBlockModification", m.GetPasscodeBlockModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passcodeBlockSimple", m.GetPasscodeBlockSimple())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passcodeExpirationDays", m.GetPasscodeExpirationDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passcodeMinimumCharacterSetCount", m.GetPasscodeMinimumCharacterSetCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passcodeMinimumLength", m.GetPasscodeMinimumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passcodeMinutesOfInactivityBeforeLock", m.GetPasscodeMinutesOfInactivityBeforeLock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passcodeMinutesOfInactivityBeforeScreenTimeout", m.GetPasscodeMinutesOfInactivityBeforeScreenTimeout())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passcodePreviousPasscodeBlockCount", m.GetPasscodePreviousPasscodeBlockCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passcodeRequired", m.GetPasscodeRequired())
        if err != nil {
            return err
        }
    }
    if m.GetPasscodeRequiredType() != nil {
        cast := (*m.GetPasscodeRequiredType()).String()
        err = writer.WriteStringValue("passcodeRequiredType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passcodeSignInFailureCountBeforeWipe", m.GetPasscodeSignInFailureCountBeforeWipe())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordBlockAirDropSharing", m.GetPasswordBlockAirDropSharing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordBlockAutoFill", m.GetPasswordBlockAutoFill())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordBlockProximityRequests", m.GetPasswordBlockProximityRequests())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("pkiBlockOTAUpdates", m.GetPkiBlockOTAUpdates())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("podcastsBlocked", m.GetPodcastsBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("privacyForceLimitAdTracking", m.GetPrivacyForceLimitAdTracking())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("proximityBlockSetupToNewDevice", m.GetProximityBlockSetupToNewDevice())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("safariBlockAutofill", m.GetSafariBlockAutofill())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("safariBlocked", m.GetSafariBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("safariBlockJavaScript", m.GetSafariBlockJavaScript())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("safariBlockPopups", m.GetSafariBlockPopups())
        if err != nil {
            return err
        }
    }
    if m.GetSafariCookieSettings() != nil {
        cast := (*m.GetSafariCookieSettings()).String()
        err = writer.WriteStringValue("safariCookieSettings", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSafariManagedDomains() != nil {
        err = writer.WriteCollectionOfStringValues("safariManagedDomains", m.GetSafariManagedDomains())
        if err != nil {
            return err
        }
    }
    if m.GetSafariPasswordAutoFillDomains() != nil {
        err = writer.WriteCollectionOfStringValues("safariPasswordAutoFillDomains", m.GetSafariPasswordAutoFillDomains())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("safariRequireFraudWarning", m.GetSafariRequireFraudWarning())
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
        err = writer.WriteBoolValue("sharedDeviceBlockTemporarySessions", m.GetSharedDeviceBlockTemporarySessions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("siriBlocked", m.GetSiriBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("siriBlockedWhenLocked", m.GetSiriBlockedWhenLocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("siriBlockUserGeneratedContent", m.GetSiriBlockUserGeneratedContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("siriRequireProfanityFilter", m.GetSiriRequireProfanityFilter())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("softwareUpdatesEnforcedDelayInDays", m.GetSoftwareUpdatesEnforcedDelayInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("softwareUpdatesForceDelayed", m.GetSoftwareUpdatesForceDelayed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("spotlightBlockInternetResults", m.GetSpotlightBlockInternetResults())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("unpairedExternalBootToRecoveryAllowed", m.GetUnpairedExternalBootToRecoveryAllowed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usbRestrictedModeBlocked", m.GetUsbRestrictedModeBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("voiceDialingBlocked", m.GetVoiceDialingBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("vpnBlockCreation", m.GetVpnBlockCreation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("wallpaperBlockModification", m.GetWallpaperBlockModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("wiFiConnectOnlyToConfiguredNetworks", m.GetWiFiConnectOnlyToConfiguredNetworks())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("wiFiConnectToAllowedNetworksOnlyForced", m.GetWiFiConnectToAllowedNetworksOnlyForced())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("wifiPowerOnForced", m.GetWifiPowerOnForced())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountBlockModification sets the accountBlockModification property value. Indicates whether or not to allow account modification when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) SetAccountBlockModification(value *bool)() {
    m.accountBlockModification = value
}
// SetActivationLockAllowWhenSupervised sets the activationLockAllowWhenSupervised property value. Indicates whether or not to allow activation lock when the device is in the supervised mode.
func (m *IosGeneralDeviceConfiguration) SetActivationLockAllowWhenSupervised(value *bool)() {
    m.activationLockAllowWhenSupervised = value
}
// SetAirDropBlocked sets the airDropBlocked property value. Indicates whether or not to allow AirDrop when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) SetAirDropBlocked(value *bool)() {
    m.airDropBlocked = value
}
// SetAirDropForceUnmanagedDropTarget sets the airDropForceUnmanagedDropTarget property value. Indicates whether or not to cause AirDrop to be considered an unmanaged drop target (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) SetAirDropForceUnmanagedDropTarget(value *bool)() {
    m.airDropForceUnmanagedDropTarget = value
}
// SetAirPlayForcePairingPasswordForOutgoingRequests sets the airPlayForcePairingPasswordForOutgoingRequests property value. Indicates whether or not to enforce all devices receiving AirPlay requests from this device to use a pairing password.
func (m *IosGeneralDeviceConfiguration) SetAirPlayForcePairingPasswordForOutgoingRequests(value *bool)() {
    m.airPlayForcePairingPasswordForOutgoingRequests = value
}
// SetAirPrintBlockCredentialsStorage sets the airPrintBlockCredentialsStorage property value. Indicates whether or not keychain storage of username and password for Airprint is blocked (iOS 11.0 and later).
func (m *IosGeneralDeviceConfiguration) SetAirPrintBlockCredentialsStorage(value *bool)() {
    m.airPrintBlockCredentialsStorage = value
}
// SetAirPrintBlocked sets the airPrintBlocked property value. Indicates whether or not AirPrint is blocked (iOS 11.0 and later).
func (m *IosGeneralDeviceConfiguration) SetAirPrintBlocked(value *bool)() {
    m.airPrintBlocked = value
}
// SetAirPrintBlockiBeaconDiscovery sets the airPrintBlockiBeaconDiscovery property value. Indicates whether or not iBeacon discovery of AirPrint printers is blocked. This prevents spurious AirPrint Bluetooth beacons from phishing for network traffic (iOS 11.0 and later).
func (m *IosGeneralDeviceConfiguration) SetAirPrintBlockiBeaconDiscovery(value *bool)() {
    m.airPrintBlockiBeaconDiscovery = value
}
// SetAirPrintForceTrustedTLS sets the airPrintForceTrustedTLS property value. Indicates if trusted certificates are required for TLS printing communication (iOS 11.0 and later).
func (m *IosGeneralDeviceConfiguration) SetAirPrintForceTrustedTLS(value *bool)() {
    m.airPrintForceTrustedTLS = value
}
// SetAppClipsBlocked sets the appClipsBlocked property value. Prevents a user from adding any App Clips and removes any existing App Clips on the device.
func (m *IosGeneralDeviceConfiguration) SetAppClipsBlocked(value *bool)() {
    m.appClipsBlocked = value
}
// SetAppleNewsBlocked sets the appleNewsBlocked property value. Indicates whether or not to block the user from using News when the device is in supervised mode (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) SetAppleNewsBlocked(value *bool)() {
    m.appleNewsBlocked = value
}
// SetApplePersonalizedAdsBlocked sets the applePersonalizedAdsBlocked property value. Limits Apple personalized advertising when true. Available in iOS 14 and later.
func (m *IosGeneralDeviceConfiguration) SetApplePersonalizedAdsBlocked(value *bool)() {
    m.applePersonalizedAdsBlocked = value
}
// SetAppleWatchBlockPairing sets the appleWatchBlockPairing property value. Indicates whether or not to allow Apple Watch pairing when the device is in supervised mode (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) SetAppleWatchBlockPairing(value *bool)() {
    m.appleWatchBlockPairing = value
}
// SetAppleWatchForceWristDetection sets the appleWatchForceWristDetection property value. Indicates whether or not to force a paired Apple Watch to use Wrist Detection (iOS 8.2 and later).
func (m *IosGeneralDeviceConfiguration) SetAppleWatchForceWristDetection(value *bool)() {
    m.appleWatchForceWristDetection = value
}
// SetAppRemovalBlocked sets the appRemovalBlocked property value. Indicates if the removal of apps is allowed.
func (m *IosGeneralDeviceConfiguration) SetAppRemovalBlocked(value *bool)() {
    m.appRemovalBlocked = value
}
// SetAppsSingleAppModeList sets the appsSingleAppModeList property value. Gets or sets the list of iOS apps allowed to autonomously enter Single App Mode. Supervised only. iOS 7.0 and later. This collection can contain a maximum of 500 elements.
func (m *IosGeneralDeviceConfiguration) SetAppsSingleAppModeList(value []AppListItemable)() {
    m.appsSingleAppModeList = value
}
// SetAppStoreBlockAutomaticDownloads sets the appStoreBlockAutomaticDownloads property value. Indicates whether or not to block the automatic downloading of apps purchased on other devices when the device is in supervised mode (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) SetAppStoreBlockAutomaticDownloads(value *bool)() {
    m.appStoreBlockAutomaticDownloads = value
}
// SetAppStoreBlocked sets the appStoreBlocked property value. Indicates whether or not to block the user from using the App Store. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) SetAppStoreBlocked(value *bool)() {
    m.appStoreBlocked = value
}
// SetAppStoreBlockInAppPurchases sets the appStoreBlockInAppPurchases property value. Indicates whether or not to block the user from making in app purchases.
func (m *IosGeneralDeviceConfiguration) SetAppStoreBlockInAppPurchases(value *bool)() {
    m.appStoreBlockInAppPurchases = value
}
// SetAppStoreBlockUIAppInstallation sets the appStoreBlockUIAppInstallation property value. Indicates whether or not to block the App Store app, not restricting installation through Host apps. Applies to supervised mode only (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) SetAppStoreBlockUIAppInstallation(value *bool)() {
    m.appStoreBlockUIAppInstallation = value
}
// SetAppStoreRequirePassword sets the appStoreRequirePassword property value. Indicates whether or not to require a password when using the app store.
func (m *IosGeneralDeviceConfiguration) SetAppStoreRequirePassword(value *bool)() {
    m.appStoreRequirePassword = value
}
// SetAppsVisibilityList sets the appsVisibilityList property value. List of apps in the visibility list (either visible/launchable apps list or hidden/unlaunchable apps list, controlled by AppsVisibilityListType) (iOS 9.3 and later). This collection can contain a maximum of 10000 elements.
func (m *IosGeneralDeviceConfiguration) SetAppsVisibilityList(value []AppListItemable)() {
    m.appsVisibilityList = value
}
// SetAppsVisibilityListType sets the appsVisibilityListType property value. Possible values of the compliance app list.
func (m *IosGeneralDeviceConfiguration) SetAppsVisibilityListType(value *AppListType)() {
    m.appsVisibilityListType = value
}
// SetAutoFillForceAuthentication sets the autoFillForceAuthentication property value. Indicates whether or not to force user authentication before autofilling passwords and credit card information in Safari and other apps on supervised devices.
func (m *IosGeneralDeviceConfiguration) SetAutoFillForceAuthentication(value *bool)() {
    m.autoFillForceAuthentication = value
}
// SetAutoUnlockBlocked sets the autoUnlockBlocked property value. Blocks users from unlocking their device with Apple Watch. Available for devices running iOS and iPadOS versions 14.5 and later.
func (m *IosGeneralDeviceConfiguration) SetAutoUnlockBlocked(value *bool)() {
    m.autoUnlockBlocked = value
}
// SetBlockSystemAppRemoval sets the blockSystemAppRemoval property value. Indicates whether or not the removal of system apps from the device is blocked on a supervised device (iOS 11.0 and later).
func (m *IosGeneralDeviceConfiguration) SetBlockSystemAppRemoval(value *bool)() {
    m.blockSystemAppRemoval = value
}
// SetBluetoothBlockModification sets the bluetoothBlockModification property value. Indicates whether or not to allow modification of Bluetooth settings when the device is in supervised mode (iOS 10.0 and later).
func (m *IosGeneralDeviceConfiguration) SetBluetoothBlockModification(value *bool)() {
    m.bluetoothBlockModification = value
}
// SetCameraBlocked sets the cameraBlocked property value. Indicates whether or not to block the user from accessing the camera of the device. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) SetCameraBlocked(value *bool)() {
    m.cameraBlocked = value
}
// SetCellularBlockDataRoaming sets the cellularBlockDataRoaming property value. Indicates whether or not to block data roaming.
func (m *IosGeneralDeviceConfiguration) SetCellularBlockDataRoaming(value *bool)() {
    m.cellularBlockDataRoaming = value
}
// SetCellularBlockGlobalBackgroundFetchWhileRoaming sets the cellularBlockGlobalBackgroundFetchWhileRoaming property value. Indicates whether or not to block global background fetch while roaming.
func (m *IosGeneralDeviceConfiguration) SetCellularBlockGlobalBackgroundFetchWhileRoaming(value *bool)() {
    m.cellularBlockGlobalBackgroundFetchWhileRoaming = value
}
// SetCellularBlockPerAppDataModification sets the cellularBlockPerAppDataModification property value. Indicates whether or not to allow changes to cellular app data usage settings when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) SetCellularBlockPerAppDataModification(value *bool)() {
    m.cellularBlockPerAppDataModification = value
}
// SetCellularBlockPersonalHotspot sets the cellularBlockPersonalHotspot property value. Indicates whether or not to block Personal Hotspot.
func (m *IosGeneralDeviceConfiguration) SetCellularBlockPersonalHotspot(value *bool)() {
    m.cellularBlockPersonalHotspot = value
}
// SetCellularBlockPersonalHotspotModification sets the cellularBlockPersonalHotspotModification property value. Indicates whether or not to block the user from modifying the personal hotspot setting (iOS 12.2 or later).
func (m *IosGeneralDeviceConfiguration) SetCellularBlockPersonalHotspotModification(value *bool)() {
    m.cellularBlockPersonalHotspotModification = value
}
// SetCellularBlockPlanModification sets the cellularBlockPlanModification property value. Indicates whether or not to allow users to change the settings of the cellular plan on a supervised device.
func (m *IosGeneralDeviceConfiguration) SetCellularBlockPlanModification(value *bool)() {
    m.cellularBlockPlanModification = value
}
// SetCellularBlockVoiceRoaming sets the cellularBlockVoiceRoaming property value. Indicates whether or not to block voice roaming.
func (m *IosGeneralDeviceConfiguration) SetCellularBlockVoiceRoaming(value *bool)() {
    m.cellularBlockVoiceRoaming = value
}
// SetCertificatesBlockUntrustedTlsCertificates sets the certificatesBlockUntrustedTlsCertificates property value. Indicates whether or not to block untrusted TLS certificates.
func (m *IosGeneralDeviceConfiguration) SetCertificatesBlockUntrustedTlsCertificates(value *bool)() {
    m.certificatesBlockUntrustedTlsCertificates = value
}
// SetClassroomAppBlockRemoteScreenObservation sets the classroomAppBlockRemoteScreenObservation property value. Indicates whether or not to allow remote screen observation by Classroom app when the device is in supervised mode (iOS 9.3 and later).
func (m *IosGeneralDeviceConfiguration) SetClassroomAppBlockRemoteScreenObservation(value *bool)() {
    m.classroomAppBlockRemoteScreenObservation = value
}
// SetClassroomAppForceUnpromptedScreenObservation sets the classroomAppForceUnpromptedScreenObservation property value. Indicates whether or not to automatically give permission to the teacher of a managed course on the Classroom app to view a student's screen without prompting when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) SetClassroomAppForceUnpromptedScreenObservation(value *bool)() {
    m.classroomAppForceUnpromptedScreenObservation = value
}
// SetClassroomForceAutomaticallyJoinClasses sets the classroomForceAutomaticallyJoinClasses property value. Indicates whether or not to automatically give permission to the teacher's requests, without prompting the student, when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) SetClassroomForceAutomaticallyJoinClasses(value *bool)() {
    m.classroomForceAutomaticallyJoinClasses = value
}
// SetClassroomForceRequestPermissionToLeaveClasses sets the classroomForceRequestPermissionToLeaveClasses property value. Indicates whether a student enrolled in an unmanaged course via Classroom will request permission from the teacher when attempting to leave the course (iOS 11.3 and later).
func (m *IosGeneralDeviceConfiguration) SetClassroomForceRequestPermissionToLeaveClasses(value *bool)() {
    m.classroomForceRequestPermissionToLeaveClasses = value
}
// SetClassroomForceUnpromptedAppAndDeviceLock sets the classroomForceUnpromptedAppAndDeviceLock property value. Indicates whether or not to allow the teacher to lock apps or the device without prompting the student. Supervised only.
func (m *IosGeneralDeviceConfiguration) SetClassroomForceUnpromptedAppAndDeviceLock(value *bool)() {
    m.classroomForceUnpromptedAppAndDeviceLock = value
}
// SetCompliantAppListType sets the compliantAppListType property value. Possible values of the compliance app list.
func (m *IosGeneralDeviceConfiguration) SetCompliantAppListType(value *AppListType)() {
    m.compliantAppListType = value
}
// SetCompliantAppsList sets the compliantAppsList property value. List of apps in the compliance (either allow list or block list, controlled by CompliantAppListType). This collection can contain a maximum of 10000 elements.
func (m *IosGeneralDeviceConfiguration) SetCompliantAppsList(value []AppListItemable)() {
    m.compliantAppsList = value
}
// SetConfigurationProfileBlockChanges sets the configurationProfileBlockChanges property value. Indicates whether or not to block the user from installing configuration profiles and certificates interactively when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) SetConfigurationProfileBlockChanges(value *bool)() {
    m.configurationProfileBlockChanges = value
}
// SetContactsAllowManagedToUnmanagedWrite sets the contactsAllowManagedToUnmanagedWrite property value. Indicates whether or not managed apps can write contacts to unmanaged contacts accounts (iOS 12.0 and later).
func (m *IosGeneralDeviceConfiguration) SetContactsAllowManagedToUnmanagedWrite(value *bool)() {
    m.contactsAllowManagedToUnmanagedWrite = value
}
// SetContactsAllowUnmanagedToManagedRead sets the contactsAllowUnmanagedToManagedRead property value. Indicates whether or not unmanaged apps can read from managed contacts accounts (iOS 12.0 or later).
func (m *IosGeneralDeviceConfiguration) SetContactsAllowUnmanagedToManagedRead(value *bool)() {
    m.contactsAllowUnmanagedToManagedRead = value
}
// SetContinuousPathKeyboardBlocked sets the continuousPathKeyboardBlocked property value. Indicates whether or not to block the continuous path keyboard when the device is supervised (iOS 13 or later).
func (m *IosGeneralDeviceConfiguration) SetContinuousPathKeyboardBlocked(value *bool)() {
    m.continuousPathKeyboardBlocked = value
}
// SetDateAndTimeForceSetAutomatically sets the dateAndTimeForceSetAutomatically property value. Indicates whether or not the Date and Time 'Set Automatically' feature is enabled and cannot be turned off by the user (iOS 12.0 and later).
func (m *IosGeneralDeviceConfiguration) SetDateAndTimeForceSetAutomatically(value *bool)() {
    m.dateAndTimeForceSetAutomatically = value
}
// SetDefinitionLookupBlocked sets the definitionLookupBlocked property value. Indicates whether or not to block definition lookup when the device is in supervised mode (iOS 8.1.3 and later ).
func (m *IosGeneralDeviceConfiguration) SetDefinitionLookupBlocked(value *bool)() {
    m.definitionLookupBlocked = value
}
// SetDeviceBlockEnableRestrictions sets the deviceBlockEnableRestrictions property value. Indicates whether or not to allow the user to enables restrictions in the device settings when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) SetDeviceBlockEnableRestrictions(value *bool)() {
    m.deviceBlockEnableRestrictions = value
}
// SetDeviceBlockEraseContentAndSettings sets the deviceBlockEraseContentAndSettings property value. Indicates whether or not to allow the use of the 'Erase all content and settings' option on the device when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) SetDeviceBlockEraseContentAndSettings(value *bool)() {
    m.deviceBlockEraseContentAndSettings = value
}
// SetDeviceBlockNameModification sets the deviceBlockNameModification property value. Indicates whether or not to allow device name modification when the device is in supervised mode (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) SetDeviceBlockNameModification(value *bool)() {
    m.deviceBlockNameModification = value
}
// SetDiagnosticDataBlockSubmission sets the diagnosticDataBlockSubmission property value. Indicates whether or not to block diagnostic data submission.
func (m *IosGeneralDeviceConfiguration) SetDiagnosticDataBlockSubmission(value *bool)() {
    m.diagnosticDataBlockSubmission = value
}
// SetDiagnosticDataBlockSubmissionModification sets the diagnosticDataBlockSubmissionModification property value. Indicates whether or not to allow diagnostics submission settings modification when the device is in supervised mode (iOS 9.3.2 and later).
func (m *IosGeneralDeviceConfiguration) SetDiagnosticDataBlockSubmissionModification(value *bool)() {
    m.diagnosticDataBlockSubmissionModification = value
}
// SetDocumentsBlockManagedDocumentsInUnmanagedApps sets the documentsBlockManagedDocumentsInUnmanagedApps property value. Indicates whether or not to block the user from viewing managed documents in unmanaged apps.
func (m *IosGeneralDeviceConfiguration) SetDocumentsBlockManagedDocumentsInUnmanagedApps(value *bool)() {
    m.documentsBlockManagedDocumentsInUnmanagedApps = value
}
// SetDocumentsBlockUnmanagedDocumentsInManagedApps sets the documentsBlockUnmanagedDocumentsInManagedApps property value. Indicates whether or not to block the user from viewing unmanaged documents in managed apps.
func (m *IosGeneralDeviceConfiguration) SetDocumentsBlockUnmanagedDocumentsInManagedApps(value *bool)() {
    m.documentsBlockUnmanagedDocumentsInManagedApps = value
}
// SetEmailInDomainSuffixes sets the emailInDomainSuffixes property value. An email address lacking a suffix that matches any of these strings will be considered out-of-domain.
func (m *IosGeneralDeviceConfiguration) SetEmailInDomainSuffixes(value []string)() {
    m.emailInDomainSuffixes = value
}
// SetEnterpriseAppBlockTrust sets the enterpriseAppBlockTrust property value. Indicates whether or not to block the user from trusting an enterprise app.
func (m *IosGeneralDeviceConfiguration) SetEnterpriseAppBlockTrust(value *bool)() {
    m.enterpriseAppBlockTrust = value
}
// SetEnterpriseAppBlockTrustModification sets the enterpriseAppBlockTrustModification property value. [Deprecated] Configuring this setting and setting the value to 'true' has no effect on the device.
func (m *IosGeneralDeviceConfiguration) SetEnterpriseAppBlockTrustModification(value *bool)() {
    m.enterpriseAppBlockTrustModification = value
}
// SetEnterpriseBookBlockBackup sets the enterpriseBookBlockBackup property value. Indicates whether or not Enterprise book back up is blocked.
func (m *IosGeneralDeviceConfiguration) SetEnterpriseBookBlockBackup(value *bool)() {
    m.enterpriseBookBlockBackup = value
}
// SetEnterpriseBookBlockMetadataSync sets the enterpriseBookBlockMetadataSync property value. Indicates whether or not Enterprise book notes and highlights sync is blocked.
func (m *IosGeneralDeviceConfiguration) SetEnterpriseBookBlockMetadataSync(value *bool)() {
    m.enterpriseBookBlockMetadataSync = value
}
// SetEsimBlockModification sets the esimBlockModification property value. Indicates whether or not to allow the addition or removal of cellular plans on the eSIM of a supervised device.
func (m *IosGeneralDeviceConfiguration) SetEsimBlockModification(value *bool)() {
    m.esimBlockModification = value
}
// SetFaceTimeBlocked sets the faceTimeBlocked property value. Indicates whether or not to block the user from using FaceTime. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) SetFaceTimeBlocked(value *bool)() {
    m.faceTimeBlocked = value
}
// SetFilesNetworkDriveAccessBlocked sets the filesNetworkDriveAccessBlocked property value. Indicates if devices can access files or other resources on a network server using the Server Message Block (SMB) protocol. Available for devices running iOS and iPadOS, versions 13.0 and later.
func (m *IosGeneralDeviceConfiguration) SetFilesNetworkDriveAccessBlocked(value *bool)() {
    m.filesNetworkDriveAccessBlocked = value
}
// SetFilesUsbDriveAccessBlocked sets the filesUsbDriveAccessBlocked property value. Indicates if sevices with access can connect to and open files on a USB drive. Available for devices running iOS and iPadOS, versions 13.0 and later.
func (m *IosGeneralDeviceConfiguration) SetFilesUsbDriveAccessBlocked(value *bool)() {
    m.filesUsbDriveAccessBlocked = value
}
// SetFindMyDeviceInFindMyAppBlocked sets the findMyDeviceInFindMyAppBlocked property value. Indicates whether or not to block Find My Device when the device is supervised (iOS 13 or later).
func (m *IosGeneralDeviceConfiguration) SetFindMyDeviceInFindMyAppBlocked(value *bool)() {
    m.findMyDeviceInFindMyAppBlocked = value
}
// SetFindMyFriendsBlocked sets the findMyFriendsBlocked property value. Indicates whether or not to block changes to Find My Friends when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) SetFindMyFriendsBlocked(value *bool)() {
    m.findMyFriendsBlocked = value
}
// SetFindMyFriendsInFindMyAppBlocked sets the findMyFriendsInFindMyAppBlocked property value. Indicates whether or not to block Find My Friends when the device is supervised (iOS 13 or later).
func (m *IosGeneralDeviceConfiguration) SetFindMyFriendsInFindMyAppBlocked(value *bool)() {
    m.findMyFriendsInFindMyAppBlocked = value
}
// SetGameCenterBlocked sets the gameCenterBlocked property value. Indicates whether or not to block the user from using Game Center when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) SetGameCenterBlocked(value *bool)() {
    m.gameCenterBlocked = value
}
// SetGamingBlockGameCenterFriends sets the gamingBlockGameCenterFriends property value. Indicates whether or not to block the user from having friends in Game Center. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) SetGamingBlockGameCenterFriends(value *bool)() {
    m.gamingBlockGameCenterFriends = value
}
// SetGamingBlockMultiplayer sets the gamingBlockMultiplayer property value. Indicates whether or not to block the user from using multiplayer gaming. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) SetGamingBlockMultiplayer(value *bool)() {
    m.gamingBlockMultiplayer = value
}
// SetHostPairingBlocked sets the hostPairingBlocked property value. indicates whether or not to allow host pairing to control the devices an iOS device can pair with when the iOS device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) SetHostPairingBlocked(value *bool)() {
    m.hostPairingBlocked = value
}
// SetIBooksStoreBlocked sets the iBooksStoreBlocked property value. Indicates whether or not to block the user from using the iBooks Store when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) SetIBooksStoreBlocked(value *bool)() {
    m.iBooksStoreBlocked = value
}
// SetIBooksStoreBlockErotica sets the iBooksStoreBlockErotica property value. Indicates whether or not to block the user from downloading media from the iBookstore that has been tagged as erotica.
func (m *IosGeneralDeviceConfiguration) SetIBooksStoreBlockErotica(value *bool)() {
    m.iBooksStoreBlockErotica = value
}
// SetICloudBlockActivityContinuation sets the iCloudBlockActivityContinuation property value. Indicates whether or not to block the user from continuing work they started on iOS device to another iOS or macOS device.
func (m *IosGeneralDeviceConfiguration) SetICloudBlockActivityContinuation(value *bool)() {
    m.iCloudBlockActivityContinuation = value
}
// SetICloudBlockBackup sets the iCloudBlockBackup property value. Indicates whether or not to block iCloud backup. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) SetICloudBlockBackup(value *bool)() {
    m.iCloudBlockBackup = value
}
// SetICloudBlockDocumentSync sets the iCloudBlockDocumentSync property value. Indicates whether or not to block iCloud document sync. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) SetICloudBlockDocumentSync(value *bool)() {
    m.iCloudBlockDocumentSync = value
}
// SetICloudBlockManagedAppsSync sets the iCloudBlockManagedAppsSync property value. Indicates whether or not to block Managed Apps Cloud Sync.
func (m *IosGeneralDeviceConfiguration) SetICloudBlockManagedAppsSync(value *bool)() {
    m.iCloudBlockManagedAppsSync = value
}
// SetICloudBlockPhotoLibrary sets the iCloudBlockPhotoLibrary property value. Indicates whether or not to block iCloud Photo Library.
func (m *IosGeneralDeviceConfiguration) SetICloudBlockPhotoLibrary(value *bool)() {
    m.iCloudBlockPhotoLibrary = value
}
// SetICloudBlockPhotoStreamSync sets the iCloudBlockPhotoStreamSync property value. Indicates whether or not to block iCloud Photo Stream Sync.
func (m *IosGeneralDeviceConfiguration) SetICloudBlockPhotoStreamSync(value *bool)() {
    m.iCloudBlockPhotoStreamSync = value
}
// SetICloudBlockSharedPhotoStream sets the iCloudBlockSharedPhotoStream property value. Indicates whether or not to block Shared Photo Stream.
func (m *IosGeneralDeviceConfiguration) SetICloudBlockSharedPhotoStream(value *bool)() {
    m.iCloudBlockSharedPhotoStream = value
}
// SetICloudPrivateRelayBlocked sets the iCloudPrivateRelayBlocked property value. iCloud private relay is an iCloud+ service that prevents networks and servers from monitoring a person's activity across the internet. By blocking iCloud private relay, Apple will not encrypt the traffic leaving the device. Available for devices running iOS 15 and later.
func (m *IosGeneralDeviceConfiguration) SetICloudPrivateRelayBlocked(value *bool)() {
    m.iCloudPrivateRelayBlocked = value
}
// SetICloudRequireEncryptedBackup sets the iCloudRequireEncryptedBackup property value. Indicates whether or not to require backups to iCloud be encrypted.
func (m *IosGeneralDeviceConfiguration) SetICloudRequireEncryptedBackup(value *bool)() {
    m.iCloudRequireEncryptedBackup = value
}
// SetITunesBlocked sets the iTunesBlocked property value. Indicates whether or not to block the iTunes app. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) SetITunesBlocked(value *bool)() {
    m.iTunesBlocked = value
}
// SetITunesBlockExplicitContent sets the iTunesBlockExplicitContent property value. Indicates whether or not to block the user from accessing explicit content in iTunes and the App Store. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) SetITunesBlockExplicitContent(value *bool)() {
    m.iTunesBlockExplicitContent = value
}
// SetITunesBlockMusicService sets the iTunesBlockMusicService property value. Indicates whether or not to block Music service and revert Music app to classic mode when the device is in supervised mode (iOS 9.3 and later and macOS 10.12 and later).
func (m *IosGeneralDeviceConfiguration) SetITunesBlockMusicService(value *bool)() {
    m.iTunesBlockMusicService = value
}
// SetITunesBlockRadio sets the iTunesBlockRadio property value. Indicates whether or not to block the user from using iTunes Radio when the device is in supervised mode (iOS 9.3 and later).
func (m *IosGeneralDeviceConfiguration) SetITunesBlockRadio(value *bool)() {
    m.iTunesBlockRadio = value
}
// SetKeyboardBlockAutoCorrect sets the keyboardBlockAutoCorrect property value. Indicates whether or not to block keyboard auto-correction when the device is in supervised mode (iOS 8.1.3 and later).
func (m *IosGeneralDeviceConfiguration) SetKeyboardBlockAutoCorrect(value *bool)() {
    m.keyboardBlockAutoCorrect = value
}
// SetKeyboardBlockDictation sets the keyboardBlockDictation property value. Indicates whether or not to block the user from using dictation input when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) SetKeyboardBlockDictation(value *bool)() {
    m.keyboardBlockDictation = value
}
// SetKeyboardBlockPredictive sets the keyboardBlockPredictive property value. Indicates whether or not to block predictive keyboards when device is in supervised mode (iOS 8.1.3 and later).
func (m *IosGeneralDeviceConfiguration) SetKeyboardBlockPredictive(value *bool)() {
    m.keyboardBlockPredictive = value
}
// SetKeyboardBlockShortcuts sets the keyboardBlockShortcuts property value. Indicates whether or not to block keyboard shortcuts when the device is in supervised mode (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) SetKeyboardBlockShortcuts(value *bool)() {
    m.keyboardBlockShortcuts = value
}
// SetKeyboardBlockSpellCheck sets the keyboardBlockSpellCheck property value. Indicates whether or not to block keyboard spell-checking when the device is in supervised mode (iOS 8.1.3 and later).
func (m *IosGeneralDeviceConfiguration) SetKeyboardBlockSpellCheck(value *bool)() {
    m.keyboardBlockSpellCheck = value
}
// SetKeychainBlockCloudSync sets the keychainBlockCloudSync property value. Indicates whether or not iCloud keychain synchronization is blocked. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) SetKeychainBlockCloudSync(value *bool)() {
    m.keychainBlockCloudSync = value
}
// SetKioskModeAllowAssistiveSpeak sets the kioskModeAllowAssistiveSpeak property value. Indicates whether or not to allow assistive speak while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) SetKioskModeAllowAssistiveSpeak(value *bool)() {
    m.kioskModeAllowAssistiveSpeak = value
}
// SetKioskModeAllowAssistiveTouchSettings sets the kioskModeAllowAssistiveTouchSettings property value. Indicates whether or not to allow access to the Assistive Touch Settings while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) SetKioskModeAllowAssistiveTouchSettings(value *bool)() {
    m.kioskModeAllowAssistiveTouchSettings = value
}
// SetKioskModeAllowAutoLock sets the kioskModeAllowAutoLock property value. Indicates whether or not to allow device auto lock while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockAutoLock instead.
func (m *IosGeneralDeviceConfiguration) SetKioskModeAllowAutoLock(value *bool)() {
    m.kioskModeAllowAutoLock = value
}
// SetKioskModeAllowColorInversionSettings sets the kioskModeAllowColorInversionSettings property value. Indicates whether or not to allow access to the Color Inversion Settings while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) SetKioskModeAllowColorInversionSettings(value *bool)() {
    m.kioskModeAllowColorInversionSettings = value
}
// SetKioskModeAllowRingerSwitch sets the kioskModeAllowRingerSwitch property value. Indicates whether or not to allow use of the ringer switch while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockRingerSwitch instead.
func (m *IosGeneralDeviceConfiguration) SetKioskModeAllowRingerSwitch(value *bool)() {
    m.kioskModeAllowRingerSwitch = value
}
// SetKioskModeAllowScreenRotation sets the kioskModeAllowScreenRotation property value. Indicates whether or not to allow screen rotation while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockScreenRotation instead.
func (m *IosGeneralDeviceConfiguration) SetKioskModeAllowScreenRotation(value *bool)() {
    m.kioskModeAllowScreenRotation = value
}
// SetKioskModeAllowSleepButton sets the kioskModeAllowSleepButton property value. Indicates whether or not to allow use of the sleep button while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockSleepButton instead.
func (m *IosGeneralDeviceConfiguration) SetKioskModeAllowSleepButton(value *bool)() {
    m.kioskModeAllowSleepButton = value
}
// SetKioskModeAllowTouchscreen sets the kioskModeAllowTouchscreen property value. Indicates whether or not to allow use of the touchscreen while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockTouchscreen instead.
func (m *IosGeneralDeviceConfiguration) SetKioskModeAllowTouchscreen(value *bool)() {
    m.kioskModeAllowTouchscreen = value
}
// SetKioskModeAllowVoiceControlModification sets the kioskModeAllowVoiceControlModification property value. Indicates whether or not to allow the user to toggle voice control in kiosk mode.
func (m *IosGeneralDeviceConfiguration) SetKioskModeAllowVoiceControlModification(value *bool)() {
    m.kioskModeAllowVoiceControlModification = value
}
// SetKioskModeAllowVoiceOverSettings sets the kioskModeAllowVoiceOverSettings property value. Indicates whether or not to allow access to the voice over settings while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) SetKioskModeAllowVoiceOverSettings(value *bool)() {
    m.kioskModeAllowVoiceOverSettings = value
}
// SetKioskModeAllowVolumeButtons sets the kioskModeAllowVolumeButtons property value. Indicates whether or not to allow use of the volume buttons while in kiosk mode. This property's functionality is redundant with the OS default and is deprecated. Use KioskModeBlockVolumeButtons instead.
func (m *IosGeneralDeviceConfiguration) SetKioskModeAllowVolumeButtons(value *bool)() {
    m.kioskModeAllowVolumeButtons = value
}
// SetKioskModeAllowZoomSettings sets the kioskModeAllowZoomSettings property value. Indicates whether or not to allow access to the zoom settings while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) SetKioskModeAllowZoomSettings(value *bool)() {
    m.kioskModeAllowZoomSettings = value
}
// SetKioskModeAppStoreUrl sets the kioskModeAppStoreUrl property value. URL in the app store to the app to use for kiosk mode. Use if KioskModeManagedAppId is not known.
func (m *IosGeneralDeviceConfiguration) SetKioskModeAppStoreUrl(value *string)() {
    m.kioskModeAppStoreUrl = value
}
// SetKioskModeAppType sets the kioskModeAppType property value. App source options for iOS kiosk mode.
func (m *IosGeneralDeviceConfiguration) SetKioskModeAppType(value *IosKioskModeAppType)() {
    m.kioskModeAppType = value
}
// SetKioskModeBlockAutoLock sets the kioskModeBlockAutoLock property value. Indicates whether or not to block device auto lock while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) SetKioskModeBlockAutoLock(value *bool)() {
    m.kioskModeBlockAutoLock = value
}
// SetKioskModeBlockRingerSwitch sets the kioskModeBlockRingerSwitch property value. Indicates whether or not to block use of the ringer switch while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) SetKioskModeBlockRingerSwitch(value *bool)() {
    m.kioskModeBlockRingerSwitch = value
}
// SetKioskModeBlockScreenRotation sets the kioskModeBlockScreenRotation property value. Indicates whether or not to block screen rotation while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) SetKioskModeBlockScreenRotation(value *bool)() {
    m.kioskModeBlockScreenRotation = value
}
// SetKioskModeBlockSleepButton sets the kioskModeBlockSleepButton property value. Indicates whether or not to block use of the sleep button while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) SetKioskModeBlockSleepButton(value *bool)() {
    m.kioskModeBlockSleepButton = value
}
// SetKioskModeBlockTouchscreen sets the kioskModeBlockTouchscreen property value. Indicates whether or not to block use of the touchscreen while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) SetKioskModeBlockTouchscreen(value *bool)() {
    m.kioskModeBlockTouchscreen = value
}
// SetKioskModeBlockVolumeButtons sets the kioskModeBlockVolumeButtons property value. Indicates whether or not to block the volume buttons while in Kiosk Mode.
func (m *IosGeneralDeviceConfiguration) SetKioskModeBlockVolumeButtons(value *bool)() {
    m.kioskModeBlockVolumeButtons = value
}
// SetKioskModeBuiltInAppId sets the kioskModeBuiltInAppId property value. ID for built-in apps to use for kiosk mode. Used when KioskModeManagedAppId and KioskModeAppStoreUrl are not set.
func (m *IosGeneralDeviceConfiguration) SetKioskModeBuiltInAppId(value *string)() {
    m.kioskModeBuiltInAppId = value
}
// SetKioskModeEnableVoiceControl sets the kioskModeEnableVoiceControl property value. Indicates whether or not to enable voice control in kiosk mode.
func (m *IosGeneralDeviceConfiguration) SetKioskModeEnableVoiceControl(value *bool)() {
    m.kioskModeEnableVoiceControl = value
}
// SetKioskModeManagedAppId sets the kioskModeManagedAppId property value. Managed app id of the app to use for kiosk mode. If KioskModeManagedAppId is specified then KioskModeAppStoreUrl will be ignored.
func (m *IosGeneralDeviceConfiguration) SetKioskModeManagedAppId(value *string)() {
    m.kioskModeManagedAppId = value
}
// SetKioskModeRequireAssistiveTouch sets the kioskModeRequireAssistiveTouch property value. Indicates whether or not to require assistive touch while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) SetKioskModeRequireAssistiveTouch(value *bool)() {
    m.kioskModeRequireAssistiveTouch = value
}
// SetKioskModeRequireColorInversion sets the kioskModeRequireColorInversion property value. Indicates whether or not to require color inversion while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) SetKioskModeRequireColorInversion(value *bool)() {
    m.kioskModeRequireColorInversion = value
}
// SetKioskModeRequireMonoAudio sets the kioskModeRequireMonoAudio property value. Indicates whether or not to require mono audio while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) SetKioskModeRequireMonoAudio(value *bool)() {
    m.kioskModeRequireMonoAudio = value
}
// SetKioskModeRequireVoiceOver sets the kioskModeRequireVoiceOver property value. Indicates whether or not to require voice over while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) SetKioskModeRequireVoiceOver(value *bool)() {
    m.kioskModeRequireVoiceOver = value
}
// SetKioskModeRequireZoom sets the kioskModeRequireZoom property value. Indicates whether or not to require zoom while in kiosk mode.
func (m *IosGeneralDeviceConfiguration) SetKioskModeRequireZoom(value *bool)() {
    m.kioskModeRequireZoom = value
}
// SetLockScreenBlockControlCenter sets the lockScreenBlockControlCenter property value. Indicates whether or not to block the user from using control center on the lock screen.
func (m *IosGeneralDeviceConfiguration) SetLockScreenBlockControlCenter(value *bool)() {
    m.lockScreenBlockControlCenter = value
}
// SetLockScreenBlockNotificationView sets the lockScreenBlockNotificationView property value. Indicates whether or not to block the user from using the notification view on the lock screen.
func (m *IosGeneralDeviceConfiguration) SetLockScreenBlockNotificationView(value *bool)() {
    m.lockScreenBlockNotificationView = value
}
// SetLockScreenBlockPassbook sets the lockScreenBlockPassbook property value. Indicates whether or not to block the user from using passbook when the device is locked.
func (m *IosGeneralDeviceConfiguration) SetLockScreenBlockPassbook(value *bool)() {
    m.lockScreenBlockPassbook = value
}
// SetLockScreenBlockTodayView sets the lockScreenBlockTodayView property value. Indicates whether or not to block the user from using the Today View on the lock screen.
func (m *IosGeneralDeviceConfiguration) SetLockScreenBlockTodayView(value *bool)() {
    m.lockScreenBlockTodayView = value
}
// SetManagedPasteboardRequired sets the managedPasteboardRequired property value. Open-in management controls how people share data between unmanaged and managed apps. Setting this to true enforces copy/paste restrictions based on how you configured Block viewing corporate documents in unmanaged apps  and  Block viewing non-corporate documents in corporate apps.
func (m *IosGeneralDeviceConfiguration) SetManagedPasteboardRequired(value *bool)() {
    m.managedPasteboardRequired = value
}
// SetMediaContentRatingApps sets the mediaContentRatingApps property value. Apps rating as in media content
func (m *IosGeneralDeviceConfiguration) SetMediaContentRatingApps(value *RatingAppsType)() {
    m.mediaContentRatingApps = value
}
// SetMediaContentRatingAustralia sets the mediaContentRatingAustralia property value. Media content rating settings for Australia
func (m *IosGeneralDeviceConfiguration) SetMediaContentRatingAustralia(value MediaContentRatingAustraliaable)() {
    m.mediaContentRatingAustralia = value
}
// SetMediaContentRatingCanada sets the mediaContentRatingCanada property value. Media content rating settings for Canada
func (m *IosGeneralDeviceConfiguration) SetMediaContentRatingCanada(value MediaContentRatingCanadaable)() {
    m.mediaContentRatingCanada = value
}
// SetMediaContentRatingFrance sets the mediaContentRatingFrance property value. Media content rating settings for France
func (m *IosGeneralDeviceConfiguration) SetMediaContentRatingFrance(value MediaContentRatingFranceable)() {
    m.mediaContentRatingFrance = value
}
// SetMediaContentRatingGermany sets the mediaContentRatingGermany property value. Media content rating settings for Germany
func (m *IosGeneralDeviceConfiguration) SetMediaContentRatingGermany(value MediaContentRatingGermanyable)() {
    m.mediaContentRatingGermany = value
}
// SetMediaContentRatingIreland sets the mediaContentRatingIreland property value. Media content rating settings for Ireland
func (m *IosGeneralDeviceConfiguration) SetMediaContentRatingIreland(value MediaContentRatingIrelandable)() {
    m.mediaContentRatingIreland = value
}
// SetMediaContentRatingJapan sets the mediaContentRatingJapan property value. Media content rating settings for Japan
func (m *IosGeneralDeviceConfiguration) SetMediaContentRatingJapan(value MediaContentRatingJapanable)() {
    m.mediaContentRatingJapan = value
}
// SetMediaContentRatingNewZealand sets the mediaContentRatingNewZealand property value. Media content rating settings for New Zealand
func (m *IosGeneralDeviceConfiguration) SetMediaContentRatingNewZealand(value MediaContentRatingNewZealandable)() {
    m.mediaContentRatingNewZealand = value
}
// SetMediaContentRatingUnitedKingdom sets the mediaContentRatingUnitedKingdom property value. Media content rating settings for United Kingdom
func (m *IosGeneralDeviceConfiguration) SetMediaContentRatingUnitedKingdom(value MediaContentRatingUnitedKingdomable)() {
    m.mediaContentRatingUnitedKingdom = value
}
// SetMediaContentRatingUnitedStates sets the mediaContentRatingUnitedStates property value. Media content rating settings for United States
func (m *IosGeneralDeviceConfiguration) SetMediaContentRatingUnitedStates(value MediaContentRatingUnitedStatesable)() {
    m.mediaContentRatingUnitedStates = value
}
// SetMessagesBlocked sets the messagesBlocked property value. Indicates whether or not to block the user from using the Messages app on the supervised device.
func (m *IosGeneralDeviceConfiguration) SetMessagesBlocked(value *bool)() {
    m.messagesBlocked = value
}
// SetNetworkUsageRules sets the networkUsageRules property value. List of managed apps and the network rules that applies to them. This collection can contain a maximum of 1000 elements.
func (m *IosGeneralDeviceConfiguration) SetNetworkUsageRules(value []IosNetworkUsageRuleable)() {
    m.networkUsageRules = value
}
// SetNfcBlocked sets the nfcBlocked property value. Disable NFC to prevent devices from pairing with other NFC-enabled devices. Available for iOS/iPadOS devices running 14.2 and later.
func (m *IosGeneralDeviceConfiguration) SetNfcBlocked(value *bool)() {
    m.nfcBlocked = value
}
// SetNotificationsBlockSettingsModification sets the notificationsBlockSettingsModification property value. Indicates whether or not to allow notifications settings modification (iOS 9.3 and later).
func (m *IosGeneralDeviceConfiguration) SetNotificationsBlockSettingsModification(value *bool)() {
    m.notificationsBlockSettingsModification = value
}
// SetOnDeviceOnlyDictationForced sets the onDeviceOnlyDictationForced property value. Disables connections to Siri servers so that users can’t use Siri to dictate text. Available for devices running iOS and iPadOS versions 14.5 and later.
func (m *IosGeneralDeviceConfiguration) SetOnDeviceOnlyDictationForced(value *bool)() {
    m.onDeviceOnlyDictationForced = value
}
// SetOnDeviceOnlyTranslationForced sets the onDeviceOnlyTranslationForced property value. When set to TRUE, the setting disables connections to Siri servers so that users can’t use Siri to translate text. When set to FALSE, the setting allows connections to to Siri servers to users can use Siri to translate text. Available for devices running iOS and iPadOS versions 15.0 and later.
func (m *IosGeneralDeviceConfiguration) SetOnDeviceOnlyTranslationForced(value *bool)() {
    m.onDeviceOnlyTranslationForced = value
}
// SetPasscodeBlockFingerprintModification sets the passcodeBlockFingerprintModification property value. Block modification of registered Touch ID fingerprints when in supervised mode.
func (m *IosGeneralDeviceConfiguration) SetPasscodeBlockFingerprintModification(value *bool)() {
    m.passcodeBlockFingerprintModification = value
}
// SetPasscodeBlockFingerprintUnlock sets the passcodeBlockFingerprintUnlock property value. Indicates whether or not to block fingerprint unlock.
func (m *IosGeneralDeviceConfiguration) SetPasscodeBlockFingerprintUnlock(value *bool)() {
    m.passcodeBlockFingerprintUnlock = value
}
// SetPasscodeBlockModification sets the passcodeBlockModification property value. Indicates whether or not to allow passcode modification on the supervised device (iOS 9.0 and later).
func (m *IosGeneralDeviceConfiguration) SetPasscodeBlockModification(value *bool)() {
    m.passcodeBlockModification = value
}
// SetPasscodeBlockSimple sets the passcodeBlockSimple property value. Indicates whether or not to block simple passcodes.
func (m *IosGeneralDeviceConfiguration) SetPasscodeBlockSimple(value *bool)() {
    m.passcodeBlockSimple = value
}
// SetPasscodeExpirationDays sets the passcodeExpirationDays property value. Number of days before the passcode expires. Valid values 1 to 65535
func (m *IosGeneralDeviceConfiguration) SetPasscodeExpirationDays(value *int32)() {
    m.passcodeExpirationDays = value
}
// SetPasscodeMinimumCharacterSetCount sets the passcodeMinimumCharacterSetCount property value. Number of character sets a passcode must contain. Valid values 0 to 4
func (m *IosGeneralDeviceConfiguration) SetPasscodeMinimumCharacterSetCount(value *int32)() {
    m.passcodeMinimumCharacterSetCount = value
}
// SetPasscodeMinimumLength sets the passcodeMinimumLength property value. Minimum length of passcode. Valid values 4 to 14
func (m *IosGeneralDeviceConfiguration) SetPasscodeMinimumLength(value *int32)() {
    m.passcodeMinimumLength = value
}
// SetPasscodeMinutesOfInactivityBeforeLock sets the passcodeMinutesOfInactivityBeforeLock property value. Minutes of inactivity before a passcode is required.
func (m *IosGeneralDeviceConfiguration) SetPasscodeMinutesOfInactivityBeforeLock(value *int32)() {
    m.passcodeMinutesOfInactivityBeforeLock = value
}
// SetPasscodeMinutesOfInactivityBeforeScreenTimeout sets the passcodeMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before the screen times out.
func (m *IosGeneralDeviceConfiguration) SetPasscodeMinutesOfInactivityBeforeScreenTimeout(value *int32)() {
    m.passcodeMinutesOfInactivityBeforeScreenTimeout = value
}
// SetPasscodePreviousPasscodeBlockCount sets the passcodePreviousPasscodeBlockCount property value. Number of previous passcodes to block. Valid values 1 to 24
func (m *IosGeneralDeviceConfiguration) SetPasscodePreviousPasscodeBlockCount(value *int32)() {
    m.passcodePreviousPasscodeBlockCount = value
}
// SetPasscodeRequired sets the passcodeRequired property value. Indicates whether or not to require a passcode.
func (m *IosGeneralDeviceConfiguration) SetPasscodeRequired(value *bool)() {
    m.passcodeRequired = value
}
// SetPasscodeRequiredType sets the passcodeRequiredType property value. Possible values of required passwords.
func (m *IosGeneralDeviceConfiguration) SetPasscodeRequiredType(value *RequiredPasswordType)() {
    m.passcodeRequiredType = value
}
// SetPasscodeSignInFailureCountBeforeWipe sets the passcodeSignInFailureCountBeforeWipe property value. Number of sign in failures allowed before wiping the device. Valid values 2 to 11
func (m *IosGeneralDeviceConfiguration) SetPasscodeSignInFailureCountBeforeWipe(value *int32)() {
    m.passcodeSignInFailureCountBeforeWipe = value
}
// SetPasswordBlockAirDropSharing sets the passwordBlockAirDropSharing property value. Indicates whether or not to block sharing passwords with the AirDrop passwords feature iOS 12.0 and later).
func (m *IosGeneralDeviceConfiguration) SetPasswordBlockAirDropSharing(value *bool)() {
    m.passwordBlockAirDropSharing = value
}
// SetPasswordBlockAutoFill sets the passwordBlockAutoFill property value. Indicates if the AutoFill passwords feature is allowed (iOS 12.0 and later).
func (m *IosGeneralDeviceConfiguration) SetPasswordBlockAutoFill(value *bool)() {
    m.passwordBlockAutoFill = value
}
// SetPasswordBlockProximityRequests sets the passwordBlockProximityRequests property value. Indicates whether or not to block requesting passwords from nearby devices (iOS 12.0 and later).
func (m *IosGeneralDeviceConfiguration) SetPasswordBlockProximityRequests(value *bool)() {
    m.passwordBlockProximityRequests = value
}
// SetPkiBlockOTAUpdates sets the pkiBlockOTAUpdates property value. Indicates whether or not over-the-air PKI updates are blocked. Setting this restriction to false does not disable CRL and OCSP checks (iOS 7.0 and later).
func (m *IosGeneralDeviceConfiguration) SetPkiBlockOTAUpdates(value *bool)() {
    m.pkiBlockOTAUpdates = value
}
// SetPodcastsBlocked sets the podcastsBlocked property value. Indicates whether or not to block the user from using podcasts on the supervised device (iOS 8.0 and later).
func (m *IosGeneralDeviceConfiguration) SetPodcastsBlocked(value *bool)() {
    m.podcastsBlocked = value
}
// SetPrivacyForceLimitAdTracking sets the privacyForceLimitAdTracking property value. Indicates if ad tracking is limited.(iOS 7.0 and later).
func (m *IosGeneralDeviceConfiguration) SetPrivacyForceLimitAdTracking(value *bool)() {
    m.privacyForceLimitAdTracking = value
}
// SetProximityBlockSetupToNewDevice sets the proximityBlockSetupToNewDevice property value. Indicates whether or not to enable the prompt to setup nearby devices with a supervised device.
func (m *IosGeneralDeviceConfiguration) SetProximityBlockSetupToNewDevice(value *bool)() {
    m.proximityBlockSetupToNewDevice = value
}
// SetSafariBlockAutofill sets the safariBlockAutofill property value. Indicates whether or not to block the user from using Auto fill in Safari. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) SetSafariBlockAutofill(value *bool)() {
    m.safariBlockAutofill = value
}
// SetSafariBlocked sets the safariBlocked property value. Indicates whether or not to block the user from using Safari. Requires a supervised device for iOS 13 and later.
func (m *IosGeneralDeviceConfiguration) SetSafariBlocked(value *bool)() {
    m.safariBlocked = value
}
// SetSafariBlockJavaScript sets the safariBlockJavaScript property value. Indicates whether or not to block JavaScript in Safari.
func (m *IosGeneralDeviceConfiguration) SetSafariBlockJavaScript(value *bool)() {
    m.safariBlockJavaScript = value
}
// SetSafariBlockPopups sets the safariBlockPopups property value. Indicates whether or not to block popups in Safari.
func (m *IosGeneralDeviceConfiguration) SetSafariBlockPopups(value *bool)() {
    m.safariBlockPopups = value
}
// SetSafariCookieSettings sets the safariCookieSettings property value. Web Browser Cookie Settings.
func (m *IosGeneralDeviceConfiguration) SetSafariCookieSettings(value *WebBrowserCookieSettings)() {
    m.safariCookieSettings = value
}
// SetSafariManagedDomains sets the safariManagedDomains property value. URLs matching the patterns listed here will be considered managed.
func (m *IosGeneralDeviceConfiguration) SetSafariManagedDomains(value []string)() {
    m.safariManagedDomains = value
}
// SetSafariPasswordAutoFillDomains sets the safariPasswordAutoFillDomains property value. Users can save passwords in Safari only from URLs matching the patterns listed here. Applies to devices in supervised mode (iOS 9.3 and later).
func (m *IosGeneralDeviceConfiguration) SetSafariPasswordAutoFillDomains(value []string)() {
    m.safariPasswordAutoFillDomains = value
}
// SetSafariRequireFraudWarning sets the safariRequireFraudWarning property value. Indicates whether or not to require fraud warning in Safari.
func (m *IosGeneralDeviceConfiguration) SetSafariRequireFraudWarning(value *bool)() {
    m.safariRequireFraudWarning = value
}
// SetScreenCaptureBlocked sets the screenCaptureBlocked property value. Indicates whether or not to block the user from taking Screenshots.
func (m *IosGeneralDeviceConfiguration) SetScreenCaptureBlocked(value *bool)() {
    m.screenCaptureBlocked = value
}
// SetSharedDeviceBlockTemporarySessions sets the sharedDeviceBlockTemporarySessions property value. Indicates whether or not to block temporary sessions on Shared iPads (iOS 13.4 or later).
func (m *IosGeneralDeviceConfiguration) SetSharedDeviceBlockTemporarySessions(value *bool)() {
    m.sharedDeviceBlockTemporarySessions = value
}
// SetSiriBlocked sets the siriBlocked property value. Indicates whether or not to block the user from using Siri.
func (m *IosGeneralDeviceConfiguration) SetSiriBlocked(value *bool)() {
    m.siriBlocked = value
}
// SetSiriBlockedWhenLocked sets the siriBlockedWhenLocked property value. Indicates whether or not to block the user from using Siri when locked.
func (m *IosGeneralDeviceConfiguration) SetSiriBlockedWhenLocked(value *bool)() {
    m.siriBlockedWhenLocked = value
}
// SetSiriBlockUserGeneratedContent sets the siriBlockUserGeneratedContent property value. Indicates whether or not to block Siri from querying user-generated content when used on a supervised device.
func (m *IosGeneralDeviceConfiguration) SetSiriBlockUserGeneratedContent(value *bool)() {
    m.siriBlockUserGeneratedContent = value
}
// SetSiriRequireProfanityFilter sets the siriRequireProfanityFilter property value. Indicates whether or not to prevent Siri from dictating, or speaking profane language on supervised device.
func (m *IosGeneralDeviceConfiguration) SetSiriRequireProfanityFilter(value *bool)() {
    m.siriRequireProfanityFilter = value
}
// SetSoftwareUpdatesEnforcedDelayInDays sets the softwareUpdatesEnforcedDelayInDays property value. Sets how many days a software update will be delyed for a supervised device. Valid values 0 to 90
func (m *IosGeneralDeviceConfiguration) SetSoftwareUpdatesEnforcedDelayInDays(value *int32)() {
    m.softwareUpdatesEnforcedDelayInDays = value
}
// SetSoftwareUpdatesForceDelayed sets the softwareUpdatesForceDelayed property value. Indicates whether or not to delay user visibility of software updates when the device is in supervised mode.
func (m *IosGeneralDeviceConfiguration) SetSoftwareUpdatesForceDelayed(value *bool)() {
    m.softwareUpdatesForceDelayed = value
}
// SetSpotlightBlockInternetResults sets the spotlightBlockInternetResults property value. Indicates whether or not to block Spotlight search from returning internet results on supervised device.
func (m *IosGeneralDeviceConfiguration) SetSpotlightBlockInternetResults(value *bool)() {
    m.spotlightBlockInternetResults = value
}
// SetUnpairedExternalBootToRecoveryAllowed sets the unpairedExternalBootToRecoveryAllowed property value. Allow users to boot devices into recovery mode with unpaired devices. Available for devices running iOS and iPadOS versions 14.5 and later.
func (m *IosGeneralDeviceConfiguration) SetUnpairedExternalBootToRecoveryAllowed(value *bool)() {
    m.unpairedExternalBootToRecoveryAllowed = value
}
// SetUsbRestrictedModeBlocked sets the usbRestrictedModeBlocked property value. Indicates if connecting to USB accessories while the device is locked is allowed (iOS 11.4.1 and later).
func (m *IosGeneralDeviceConfiguration) SetUsbRestrictedModeBlocked(value *bool)() {
    m.usbRestrictedModeBlocked = value
}
// SetVoiceDialingBlocked sets the voiceDialingBlocked property value. Indicates whether or not to block voice dialing.
func (m *IosGeneralDeviceConfiguration) SetVoiceDialingBlocked(value *bool)() {
    m.voiceDialingBlocked = value
}
// SetVpnBlockCreation sets the vpnBlockCreation property value. Indicates whether or not the creation of VPN configurations is blocked (iOS 11.0 and later).
func (m *IosGeneralDeviceConfiguration) SetVpnBlockCreation(value *bool)() {
    m.vpnBlockCreation = value
}
// SetWallpaperBlockModification sets the wallpaperBlockModification property value. Indicates whether or not to allow wallpaper modification on supervised device (iOS 9.0 and later) .
func (m *IosGeneralDeviceConfiguration) SetWallpaperBlockModification(value *bool)() {
    m.wallpaperBlockModification = value
}
// SetWiFiConnectOnlyToConfiguredNetworks sets the wiFiConnectOnlyToConfiguredNetworks property value. Indicates whether or not to force the device to use only Wi-Fi networks from configuration profiles when the device is in supervised mode. Available for devices running iOS and iPadOS versions 14.4 and earlier. Devices running 14.5+ should use the setting, 'WiFiConnectToAllowedNetworksOnlyForced.
func (m *IosGeneralDeviceConfiguration) SetWiFiConnectOnlyToConfiguredNetworks(value *bool)() {
    m.wiFiConnectOnlyToConfiguredNetworks = value
}
// SetWiFiConnectToAllowedNetworksOnlyForced sets the wiFiConnectToAllowedNetworksOnlyForced property value. Require devices to use Wi-Fi networks set up via configuration profiles. Available for devices running iOS and iPadOS versions 14.5 and later.
func (m *IosGeneralDeviceConfiguration) SetWiFiConnectToAllowedNetworksOnlyForced(value *bool)() {
    m.wiFiConnectToAllowedNetworksOnlyForced = value
}
// SetWifiPowerOnForced sets the wifiPowerOnForced property value. Indicates whether or not Wi-Fi remains on, even when device is in airplane mode. Available for devices running iOS and iPadOS, versions 13.0 and later.
func (m *IosGeneralDeviceConfiguration) SetWifiPowerOnForced(value *bool)() {
    m.wifiPowerOnForced = value
}
