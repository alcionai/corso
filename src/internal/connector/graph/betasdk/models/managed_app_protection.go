package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedAppProtection 
type ManagedAppProtection struct {
    ManagedAppPolicy
    // Data storage locations where a user may store managed data.
    allowedDataIngestionLocations []ManagedAppDataIngestionLocation
    // Data storage locations where a user may store managed data.
    allowedDataStorageLocations []ManagedAppDataStorageLocation
    // Data can be transferred from/to these classes of apps
    allowedInboundDataTransferSources *ManagedAppDataTransferLevel
    // Specify the number of characters that may be cut or copied from Org data and accounts to any application. This setting overrides the AllowedOutboundClipboardSharingLevel restriction. Default value of '0' means no exception is allowed.
    allowedOutboundClipboardSharingExceptionLength *int32
    // Represents the level to which the device's clipboard may be shared between apps
    allowedOutboundClipboardSharingLevel *ManagedAppClipboardSharingLevel
    // Data can be transferred from/to these classes of apps
    allowedOutboundDataTransferDestinations *ManagedAppDataTransferLevel
    // An admin initiated action to be applied on a managed app.
    appActionIfDeviceComplianceRequired *ManagedAppRemediationAction
    // An admin initiated action to be applied on a managed app.
    appActionIfMaximumPinRetriesExceeded *ManagedAppRemediationAction
    // If set, it will specify what action to take in the case where the user is unable to checkin because their authentication token is invalid. This happens when the user is deleted or disabled in AAD. Possible values are: block, wipe, warn.
    appActionIfUnableToAuthenticateUser *ManagedAppRemediationAction
    // Indicates whether a user can bring data into org documents.
    blockDataIngestionIntoOrganizationDocuments *bool
    // Indicates whether contacts can be synced to the user's device.
    contactSyncBlocked *bool
    // Indicates whether the backup of a managed app's data is blocked.
    dataBackupBlocked *bool
    // Indicates whether device compliance is required.
    deviceComplianceRequired *bool
    // The classes of apps that are allowed to click-to-open a phone number, for making phone calls or sending text messages.
    dialerRestrictionLevel *ManagedAppPhoneNumberRedirectLevel
    // Indicates whether use of the app pin is required if the device pin is set.
    disableAppPinIfDevicePinIsSet *bool
    // Indicates whether use of the fingerprint reader is allowed in place of a pin if PinRequired is set to True.
    fingerprintBlocked *bool
    // A grace period before blocking app access during off clock hours.
    gracePeriodToBlockAppsDuringOffClockHours *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // Type of managed browser
    managedBrowser *ManagedBrowserType
    // Indicates whether internet links should be opened in the managed browser app, or any custom browser specified by CustomBrowserProtocol (for iOS) or CustomBrowserPackageId/CustomBrowserDisplayName (for Android)
    managedBrowserToOpenLinksRequired *bool
    // The maxium threat level allowed for an app to be compliant.
    maximumAllowedDeviceThreatLevel *ManagedAppDeviceThreatLevel
    // Maximum number of incorrect pin retry attempts before the managed app is either blocked or wiped.
    maximumPinRetries *int32
    // Versions bigger than the specified version will block the managed app from accessing company data.
    maximumRequiredOsVersion *string
    // Versions bigger than the specified version will block the managed app from accessing company data.
    maximumWarningOsVersion *string
    // Versions bigger than the specified version will block the managed app from accessing company data.
    maximumWipeOsVersion *string
    // Minimum pin length required for an app-level pin if PinRequired is set to True
    minimumPinLength *int32
    // Versions less than the specified version will block the managed app from accessing company data.
    minimumRequiredAppVersion *string
    // Versions less than the specified version will block the managed app from accessing company data.
    minimumRequiredOsVersion *string
    // Versions less than the specified version will result in warning message on the managed app.
    minimumWarningAppVersion *string
    // Versions less than the specified version will result in warning message on the managed app from accessing company data.
    minimumWarningOsVersion *string
    // Versions less than or equal to the specified version will wipe the managed app and the associated company data.
    minimumWipeAppVersion *string
    // Versions less than or equal to the specified version will wipe the managed app and the associated company data.
    minimumWipeOsVersion *string
    // Indicates how to prioritize which Mobile Threat Defense (MTD) partner is enabled for a given platform, when more than one is enabled. An app can only be actively using a single Mobile Threat Defense partner. When NULL, Microsoft Defender will be given preference. Otherwise setting the value to defenderOverThirdPartyPartner or thirdPartyPartnerOverDefender will make explicit which partner to prioritize. Possible values are: null, defenderOverThirdPartyPartner, thirdPartyPartnerOverDefender and unknownFutureValue. Default value is null. Possible values are: defenderOverThirdPartyPartner, thirdPartyPartnerOverDefender, unknownFutureValue.
    mobileThreatDefensePartnerPriority *MobileThreatDefensePartnerPriority
    // An admin initiated action to be applied on a managed app.
    mobileThreatDefenseRemediationAction *ManagedAppRemediationAction
    // Restrict managed app notification
    notificationRestriction *ManagedAppNotificationRestriction
    // Indicates whether organizational credentials are required for app use.
    organizationalCredentialsRequired *bool
    // TimePeriod before the all-level pin must be reset if PinRequired is set to True.
    periodBeforePinReset *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The period after which access is checked when the device is not connected to the internet.
    periodOfflineBeforeAccessCheck *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The amount of time an app is allowed to remain disconnected from the internet before all managed data it is wiped.
    periodOfflineBeforeWipeIsEnforced *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The period after which access is checked when the device is connected to the internet.
    periodOnlineBeforeAccessCheck *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // Character set which is to be used for a user's app PIN
    pinCharacterSet *ManagedAppPinCharacterSet
    // Indicates whether an app-level pin is required.
    pinRequired *bool
    // Timeout in minutes for an app pin instead of non biometrics passcode
    pinRequiredInsteadOfBiometricTimeout *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // Requires a pin to be unique from the number specified in this property.
    previousPinBlockCount *int32
    // Indicates whether printing is allowed from managed apps.
    printBlocked *bool
    // Indicates whether users may use the 'Save As' menu item to save a copy of protected files.
    saveAsBlocked *bool
    // Indicates whether simplePin is blocked.
    simplePinBlocked *bool
}
// NewManagedAppProtection instantiates a new ManagedAppProtection and sets the default values.
func NewManagedAppProtection()(*ManagedAppProtection) {
    m := &ManagedAppProtection{
        ManagedAppPolicy: *NewManagedAppPolicy(),
    }
    odataTypeValue := "#microsoft.graph.managedAppProtection";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateManagedAppProtectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedAppProtectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.androidManagedAppProtection":
                        return NewAndroidManagedAppProtection(), nil
                    case "#microsoft.graph.defaultManagedAppProtection":
                        return NewDefaultManagedAppProtection(), nil
                    case "#microsoft.graph.iosManagedAppProtection":
                        return NewIosManagedAppProtection(), nil
                    case "#microsoft.graph.targetedManagedAppProtection":
                        return NewTargetedManagedAppProtection(), nil
                }
            }
        }
    }
    return NewManagedAppProtection(), nil
}
// GetAllowedDataIngestionLocations gets the allowedDataIngestionLocations property value. Data storage locations where a user may store managed data.
func (m *ManagedAppProtection) GetAllowedDataIngestionLocations()([]ManagedAppDataIngestionLocation) {
    return m.allowedDataIngestionLocations
}
// GetAllowedDataStorageLocations gets the allowedDataStorageLocations property value. Data storage locations where a user may store managed data.
func (m *ManagedAppProtection) GetAllowedDataStorageLocations()([]ManagedAppDataStorageLocation) {
    return m.allowedDataStorageLocations
}
// GetAllowedInboundDataTransferSources gets the allowedInboundDataTransferSources property value. Data can be transferred from/to these classes of apps
func (m *ManagedAppProtection) GetAllowedInboundDataTransferSources()(*ManagedAppDataTransferLevel) {
    return m.allowedInboundDataTransferSources
}
// GetAllowedOutboundClipboardSharingExceptionLength gets the allowedOutboundClipboardSharingExceptionLength property value. Specify the number of characters that may be cut or copied from Org data and accounts to any application. This setting overrides the AllowedOutboundClipboardSharingLevel restriction. Default value of '0' means no exception is allowed.
func (m *ManagedAppProtection) GetAllowedOutboundClipboardSharingExceptionLength()(*int32) {
    return m.allowedOutboundClipboardSharingExceptionLength
}
// GetAllowedOutboundClipboardSharingLevel gets the allowedOutboundClipboardSharingLevel property value. Represents the level to which the device's clipboard may be shared between apps
func (m *ManagedAppProtection) GetAllowedOutboundClipboardSharingLevel()(*ManagedAppClipboardSharingLevel) {
    return m.allowedOutboundClipboardSharingLevel
}
// GetAllowedOutboundDataTransferDestinations gets the allowedOutboundDataTransferDestinations property value. Data can be transferred from/to these classes of apps
func (m *ManagedAppProtection) GetAllowedOutboundDataTransferDestinations()(*ManagedAppDataTransferLevel) {
    return m.allowedOutboundDataTransferDestinations
}
// GetAppActionIfDeviceComplianceRequired gets the appActionIfDeviceComplianceRequired property value. An admin initiated action to be applied on a managed app.
func (m *ManagedAppProtection) GetAppActionIfDeviceComplianceRequired()(*ManagedAppRemediationAction) {
    return m.appActionIfDeviceComplianceRequired
}
// GetAppActionIfMaximumPinRetriesExceeded gets the appActionIfMaximumPinRetriesExceeded property value. An admin initiated action to be applied on a managed app.
func (m *ManagedAppProtection) GetAppActionIfMaximumPinRetriesExceeded()(*ManagedAppRemediationAction) {
    return m.appActionIfMaximumPinRetriesExceeded
}
// GetAppActionIfUnableToAuthenticateUser gets the appActionIfUnableToAuthenticateUser property value. If set, it will specify what action to take in the case where the user is unable to checkin because their authentication token is invalid. This happens when the user is deleted or disabled in AAD. Possible values are: block, wipe, warn.
func (m *ManagedAppProtection) GetAppActionIfUnableToAuthenticateUser()(*ManagedAppRemediationAction) {
    return m.appActionIfUnableToAuthenticateUser
}
// GetBlockDataIngestionIntoOrganizationDocuments gets the blockDataIngestionIntoOrganizationDocuments property value. Indicates whether a user can bring data into org documents.
func (m *ManagedAppProtection) GetBlockDataIngestionIntoOrganizationDocuments()(*bool) {
    return m.blockDataIngestionIntoOrganizationDocuments
}
// GetContactSyncBlocked gets the contactSyncBlocked property value. Indicates whether contacts can be synced to the user's device.
func (m *ManagedAppProtection) GetContactSyncBlocked()(*bool) {
    return m.contactSyncBlocked
}
// GetDataBackupBlocked gets the dataBackupBlocked property value. Indicates whether the backup of a managed app's data is blocked.
func (m *ManagedAppProtection) GetDataBackupBlocked()(*bool) {
    return m.dataBackupBlocked
}
// GetDeviceComplianceRequired gets the deviceComplianceRequired property value. Indicates whether device compliance is required.
func (m *ManagedAppProtection) GetDeviceComplianceRequired()(*bool) {
    return m.deviceComplianceRequired
}
// GetDialerRestrictionLevel gets the dialerRestrictionLevel property value. The classes of apps that are allowed to click-to-open a phone number, for making phone calls or sending text messages.
func (m *ManagedAppProtection) GetDialerRestrictionLevel()(*ManagedAppPhoneNumberRedirectLevel) {
    return m.dialerRestrictionLevel
}
// GetDisableAppPinIfDevicePinIsSet gets the disableAppPinIfDevicePinIsSet property value. Indicates whether use of the app pin is required if the device pin is set.
func (m *ManagedAppProtection) GetDisableAppPinIfDevicePinIsSet()(*bool) {
    return m.disableAppPinIfDevicePinIsSet
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedAppProtection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ManagedAppPolicy.GetFieldDeserializers()
    res["allowedDataIngestionLocations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseManagedAppDataIngestionLocation)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppDataIngestionLocation, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedAppDataIngestionLocation))
            }
            m.SetAllowedDataIngestionLocations(res)
        }
        return nil
    }
    res["allowedDataStorageLocations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseManagedAppDataStorageLocation)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppDataStorageLocation, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedAppDataStorageLocation))
            }
            m.SetAllowedDataStorageLocations(res)
        }
        return nil
    }
    res["allowedInboundDataTransferSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppDataTransferLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedInboundDataTransferSources(val.(*ManagedAppDataTransferLevel))
        }
        return nil
    }
    res["allowedOutboundClipboardSharingExceptionLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedOutboundClipboardSharingExceptionLength(val)
        }
        return nil
    }
    res["allowedOutboundClipboardSharingLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppClipboardSharingLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedOutboundClipboardSharingLevel(val.(*ManagedAppClipboardSharingLevel))
        }
        return nil
    }
    res["allowedOutboundDataTransferDestinations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppDataTransferLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedOutboundDataTransferDestinations(val.(*ManagedAppDataTransferLevel))
        }
        return nil
    }
    res["appActionIfDeviceComplianceRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfDeviceComplianceRequired(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["appActionIfMaximumPinRetriesExceeded"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfMaximumPinRetriesExceeded(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["appActionIfUnableToAuthenticateUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfUnableToAuthenticateUser(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["blockDataIngestionIntoOrganizationDocuments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockDataIngestionIntoOrganizationDocuments(val)
        }
        return nil
    }
    res["contactSyncBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactSyncBlocked(val)
        }
        return nil
    }
    res["dataBackupBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataBackupBlocked(val)
        }
        return nil
    }
    res["deviceComplianceRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceComplianceRequired(val)
        }
        return nil
    }
    res["dialerRestrictionLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppPhoneNumberRedirectLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDialerRestrictionLevel(val.(*ManagedAppPhoneNumberRedirectLevel))
        }
        return nil
    }
    res["disableAppPinIfDevicePinIsSet"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableAppPinIfDevicePinIsSet(val)
        }
        return nil
    }
    res["fingerprintBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFingerprintBlocked(val)
        }
        return nil
    }
    res["gracePeriodToBlockAppsDuringOffClockHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGracePeriodToBlockAppsDuringOffClockHours(val)
        }
        return nil
    }
    res["managedBrowser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedBrowserType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedBrowser(val.(*ManagedBrowserType))
        }
        return nil
    }
    res["managedBrowserToOpenLinksRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedBrowserToOpenLinksRequired(val)
        }
        return nil
    }
    res["maximumAllowedDeviceThreatLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppDeviceThreatLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumAllowedDeviceThreatLevel(val.(*ManagedAppDeviceThreatLevel))
        }
        return nil
    }
    res["maximumPinRetries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumPinRetries(val)
        }
        return nil
    }
    res["maximumRequiredOsVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumRequiredOsVersion(val)
        }
        return nil
    }
    res["maximumWarningOsVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumWarningOsVersion(val)
        }
        return nil
    }
    res["maximumWipeOsVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumWipeOsVersion(val)
        }
        return nil
    }
    res["minimumPinLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumPinLength(val)
        }
        return nil
    }
    res["minimumRequiredAppVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumRequiredAppVersion(val)
        }
        return nil
    }
    res["minimumRequiredOsVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumRequiredOsVersion(val)
        }
        return nil
    }
    res["minimumWarningAppVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWarningAppVersion(val)
        }
        return nil
    }
    res["minimumWarningOsVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWarningOsVersion(val)
        }
        return nil
    }
    res["minimumWipeAppVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWipeAppVersion(val)
        }
        return nil
    }
    res["minimumWipeOsVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWipeOsVersion(val)
        }
        return nil
    }
    res["mobileThreatDefensePartnerPriority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMobileThreatDefensePartnerPriority)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobileThreatDefensePartnerPriority(val.(*MobileThreatDefensePartnerPriority))
        }
        return nil
    }
    res["mobileThreatDefenseRemediationAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobileThreatDefenseRemediationAction(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["notificationRestriction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppNotificationRestriction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationRestriction(val.(*ManagedAppNotificationRestriction))
        }
        return nil
    }
    res["organizationalCredentialsRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationalCredentialsRequired(val)
        }
        return nil
    }
    res["periodBeforePinReset"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeriodBeforePinReset(val)
        }
        return nil
    }
    res["periodOfflineBeforeAccessCheck"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeriodOfflineBeforeAccessCheck(val)
        }
        return nil
    }
    res["periodOfflineBeforeWipeIsEnforced"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeriodOfflineBeforeWipeIsEnforced(val)
        }
        return nil
    }
    res["periodOnlineBeforeAccessCheck"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeriodOnlineBeforeAccessCheck(val)
        }
        return nil
    }
    res["pinCharacterSet"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppPinCharacterSet)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinCharacterSet(val.(*ManagedAppPinCharacterSet))
        }
        return nil
    }
    res["pinRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinRequired(val)
        }
        return nil
    }
    res["pinRequiredInsteadOfBiometricTimeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinRequiredInsteadOfBiometricTimeout(val)
        }
        return nil
    }
    res["previousPinBlockCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviousPinBlockCount(val)
        }
        return nil
    }
    res["printBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrintBlocked(val)
        }
        return nil
    }
    res["saveAsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSaveAsBlocked(val)
        }
        return nil
    }
    res["simplePinBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSimplePinBlocked(val)
        }
        return nil
    }
    return res
}
// GetFingerprintBlocked gets the fingerprintBlocked property value. Indicates whether use of the fingerprint reader is allowed in place of a pin if PinRequired is set to True.
func (m *ManagedAppProtection) GetFingerprintBlocked()(*bool) {
    return m.fingerprintBlocked
}
// GetGracePeriodToBlockAppsDuringOffClockHours gets the gracePeriodToBlockAppsDuringOffClockHours property value. A grace period before blocking app access during off clock hours.
func (m *ManagedAppProtection) GetGracePeriodToBlockAppsDuringOffClockHours()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.gracePeriodToBlockAppsDuringOffClockHours
}
// GetManagedBrowser gets the managedBrowser property value. Type of managed browser
func (m *ManagedAppProtection) GetManagedBrowser()(*ManagedBrowserType) {
    return m.managedBrowser
}
// GetManagedBrowserToOpenLinksRequired gets the managedBrowserToOpenLinksRequired property value. Indicates whether internet links should be opened in the managed browser app, or any custom browser specified by CustomBrowserProtocol (for iOS) or CustomBrowserPackageId/CustomBrowserDisplayName (for Android)
func (m *ManagedAppProtection) GetManagedBrowserToOpenLinksRequired()(*bool) {
    return m.managedBrowserToOpenLinksRequired
}
// GetMaximumAllowedDeviceThreatLevel gets the maximumAllowedDeviceThreatLevel property value. The maxium threat level allowed for an app to be compliant.
func (m *ManagedAppProtection) GetMaximumAllowedDeviceThreatLevel()(*ManagedAppDeviceThreatLevel) {
    return m.maximumAllowedDeviceThreatLevel
}
// GetMaximumPinRetries gets the maximumPinRetries property value. Maximum number of incorrect pin retry attempts before the managed app is either blocked or wiped.
func (m *ManagedAppProtection) GetMaximumPinRetries()(*int32) {
    return m.maximumPinRetries
}
// GetMaximumRequiredOsVersion gets the maximumRequiredOsVersion property value. Versions bigger than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) GetMaximumRequiredOsVersion()(*string) {
    return m.maximumRequiredOsVersion
}
// GetMaximumWarningOsVersion gets the maximumWarningOsVersion property value. Versions bigger than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) GetMaximumWarningOsVersion()(*string) {
    return m.maximumWarningOsVersion
}
// GetMaximumWipeOsVersion gets the maximumWipeOsVersion property value. Versions bigger than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) GetMaximumWipeOsVersion()(*string) {
    return m.maximumWipeOsVersion
}
// GetMinimumPinLength gets the minimumPinLength property value. Minimum pin length required for an app-level pin if PinRequired is set to True
func (m *ManagedAppProtection) GetMinimumPinLength()(*int32) {
    return m.minimumPinLength
}
// GetMinimumRequiredAppVersion gets the minimumRequiredAppVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) GetMinimumRequiredAppVersion()(*string) {
    return m.minimumRequiredAppVersion
}
// GetMinimumRequiredOsVersion gets the minimumRequiredOsVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) GetMinimumRequiredOsVersion()(*string) {
    return m.minimumRequiredOsVersion
}
// GetMinimumWarningAppVersion gets the minimumWarningAppVersion property value. Versions less than the specified version will result in warning message on the managed app.
func (m *ManagedAppProtection) GetMinimumWarningAppVersion()(*string) {
    return m.minimumWarningAppVersion
}
// GetMinimumWarningOsVersion gets the minimumWarningOsVersion property value. Versions less than the specified version will result in warning message on the managed app from accessing company data.
func (m *ManagedAppProtection) GetMinimumWarningOsVersion()(*string) {
    return m.minimumWarningOsVersion
}
// GetMinimumWipeAppVersion gets the minimumWipeAppVersion property value. Versions less than or equal to the specified version will wipe the managed app and the associated company data.
func (m *ManagedAppProtection) GetMinimumWipeAppVersion()(*string) {
    return m.minimumWipeAppVersion
}
// GetMinimumWipeOsVersion gets the minimumWipeOsVersion property value. Versions less than or equal to the specified version will wipe the managed app and the associated company data.
func (m *ManagedAppProtection) GetMinimumWipeOsVersion()(*string) {
    return m.minimumWipeOsVersion
}
// GetMobileThreatDefensePartnerPriority gets the mobileThreatDefensePartnerPriority property value. Indicates how to prioritize which Mobile Threat Defense (MTD) partner is enabled for a given platform, when more than one is enabled. An app can only be actively using a single Mobile Threat Defense partner. When NULL, Microsoft Defender will be given preference. Otherwise setting the value to defenderOverThirdPartyPartner or thirdPartyPartnerOverDefender will make explicit which partner to prioritize. Possible values are: null, defenderOverThirdPartyPartner, thirdPartyPartnerOverDefender and unknownFutureValue. Default value is null. Possible values are: defenderOverThirdPartyPartner, thirdPartyPartnerOverDefender, unknownFutureValue.
func (m *ManagedAppProtection) GetMobileThreatDefensePartnerPriority()(*MobileThreatDefensePartnerPriority) {
    return m.mobileThreatDefensePartnerPriority
}
// GetMobileThreatDefenseRemediationAction gets the mobileThreatDefenseRemediationAction property value. An admin initiated action to be applied on a managed app.
func (m *ManagedAppProtection) GetMobileThreatDefenseRemediationAction()(*ManagedAppRemediationAction) {
    return m.mobileThreatDefenseRemediationAction
}
// GetNotificationRestriction gets the notificationRestriction property value. Restrict managed app notification
func (m *ManagedAppProtection) GetNotificationRestriction()(*ManagedAppNotificationRestriction) {
    return m.notificationRestriction
}
// GetOrganizationalCredentialsRequired gets the organizationalCredentialsRequired property value. Indicates whether organizational credentials are required for app use.
func (m *ManagedAppProtection) GetOrganizationalCredentialsRequired()(*bool) {
    return m.organizationalCredentialsRequired
}
// GetPeriodBeforePinReset gets the periodBeforePinReset property value. TimePeriod before the all-level pin must be reset if PinRequired is set to True.
func (m *ManagedAppProtection) GetPeriodBeforePinReset()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.periodBeforePinReset
}
// GetPeriodOfflineBeforeAccessCheck gets the periodOfflineBeforeAccessCheck property value. The period after which access is checked when the device is not connected to the internet.
func (m *ManagedAppProtection) GetPeriodOfflineBeforeAccessCheck()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.periodOfflineBeforeAccessCheck
}
// GetPeriodOfflineBeforeWipeIsEnforced gets the periodOfflineBeforeWipeIsEnforced property value. The amount of time an app is allowed to remain disconnected from the internet before all managed data it is wiped.
func (m *ManagedAppProtection) GetPeriodOfflineBeforeWipeIsEnforced()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.periodOfflineBeforeWipeIsEnforced
}
// GetPeriodOnlineBeforeAccessCheck gets the periodOnlineBeforeAccessCheck property value. The period after which access is checked when the device is connected to the internet.
func (m *ManagedAppProtection) GetPeriodOnlineBeforeAccessCheck()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.periodOnlineBeforeAccessCheck
}
// GetPinCharacterSet gets the pinCharacterSet property value. Character set which is to be used for a user's app PIN
func (m *ManagedAppProtection) GetPinCharacterSet()(*ManagedAppPinCharacterSet) {
    return m.pinCharacterSet
}
// GetPinRequired gets the pinRequired property value. Indicates whether an app-level pin is required.
func (m *ManagedAppProtection) GetPinRequired()(*bool) {
    return m.pinRequired
}
// GetPinRequiredInsteadOfBiometricTimeout gets the pinRequiredInsteadOfBiometricTimeout property value. Timeout in minutes for an app pin instead of non biometrics passcode
func (m *ManagedAppProtection) GetPinRequiredInsteadOfBiometricTimeout()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.pinRequiredInsteadOfBiometricTimeout
}
// GetPreviousPinBlockCount gets the previousPinBlockCount property value. Requires a pin to be unique from the number specified in this property.
func (m *ManagedAppProtection) GetPreviousPinBlockCount()(*int32) {
    return m.previousPinBlockCount
}
// GetPrintBlocked gets the printBlocked property value. Indicates whether printing is allowed from managed apps.
func (m *ManagedAppProtection) GetPrintBlocked()(*bool) {
    return m.printBlocked
}
// GetSaveAsBlocked gets the saveAsBlocked property value. Indicates whether users may use the 'Save As' menu item to save a copy of protected files.
func (m *ManagedAppProtection) GetSaveAsBlocked()(*bool) {
    return m.saveAsBlocked
}
// GetSimplePinBlocked gets the simplePinBlocked property value. Indicates whether simplePin is blocked.
func (m *ManagedAppProtection) GetSimplePinBlocked()(*bool) {
    return m.simplePinBlocked
}
// Serialize serializes information the current object
func (m *ManagedAppProtection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ManagedAppPolicy.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedDataIngestionLocations() != nil {
        err = writer.WriteCollectionOfStringValues("allowedDataIngestionLocations", SerializeManagedAppDataIngestionLocation(m.GetAllowedDataIngestionLocations()))
        if err != nil {
            return err
        }
    }
    if m.GetAllowedDataStorageLocations() != nil {
        err = writer.WriteCollectionOfStringValues("allowedDataStorageLocations", SerializeManagedAppDataStorageLocation(m.GetAllowedDataStorageLocations()))
        if err != nil {
            return err
        }
    }
    if m.GetAllowedInboundDataTransferSources() != nil {
        cast := (*m.GetAllowedInboundDataTransferSources()).String()
        err = writer.WriteStringValue("allowedInboundDataTransferSources", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("allowedOutboundClipboardSharingExceptionLength", m.GetAllowedOutboundClipboardSharingExceptionLength())
        if err != nil {
            return err
        }
    }
    if m.GetAllowedOutboundClipboardSharingLevel() != nil {
        cast := (*m.GetAllowedOutboundClipboardSharingLevel()).String()
        err = writer.WriteStringValue("allowedOutboundClipboardSharingLevel", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAllowedOutboundDataTransferDestinations() != nil {
        cast := (*m.GetAllowedOutboundDataTransferDestinations()).String()
        err = writer.WriteStringValue("allowedOutboundDataTransferDestinations", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppActionIfDeviceComplianceRequired() != nil {
        cast := (*m.GetAppActionIfDeviceComplianceRequired()).String()
        err = writer.WriteStringValue("appActionIfDeviceComplianceRequired", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppActionIfMaximumPinRetriesExceeded() != nil {
        cast := (*m.GetAppActionIfMaximumPinRetriesExceeded()).String()
        err = writer.WriteStringValue("appActionIfMaximumPinRetriesExceeded", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppActionIfUnableToAuthenticateUser() != nil {
        cast := (*m.GetAppActionIfUnableToAuthenticateUser()).String()
        err = writer.WriteStringValue("appActionIfUnableToAuthenticateUser", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("blockDataIngestionIntoOrganizationDocuments", m.GetBlockDataIngestionIntoOrganizationDocuments())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("contactSyncBlocked", m.GetContactSyncBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("dataBackupBlocked", m.GetDataBackupBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deviceComplianceRequired", m.GetDeviceComplianceRequired())
        if err != nil {
            return err
        }
    }
    if m.GetDialerRestrictionLevel() != nil {
        cast := (*m.GetDialerRestrictionLevel()).String()
        err = writer.WriteStringValue("dialerRestrictionLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disableAppPinIfDevicePinIsSet", m.GetDisableAppPinIfDevicePinIsSet())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("fingerprintBlocked", m.GetFingerprintBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("gracePeriodToBlockAppsDuringOffClockHours", m.GetGracePeriodToBlockAppsDuringOffClockHours())
        if err != nil {
            return err
        }
    }
    if m.GetManagedBrowser() != nil {
        cast := (*m.GetManagedBrowser()).String()
        err = writer.WriteStringValue("managedBrowser", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("managedBrowserToOpenLinksRequired", m.GetManagedBrowserToOpenLinksRequired())
        if err != nil {
            return err
        }
    }
    if m.GetMaximumAllowedDeviceThreatLevel() != nil {
        cast := (*m.GetMaximumAllowedDeviceThreatLevel()).String()
        err = writer.WriteStringValue("maximumAllowedDeviceThreatLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("maximumPinRetries", m.GetMaximumPinRetries())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("maximumRequiredOsVersion", m.GetMaximumRequiredOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("maximumWarningOsVersion", m.GetMaximumWarningOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("maximumWipeOsVersion", m.GetMaximumWipeOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumPinLength", m.GetMinimumPinLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumRequiredAppVersion", m.GetMinimumRequiredAppVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumRequiredOsVersion", m.GetMinimumRequiredOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumWarningAppVersion", m.GetMinimumWarningAppVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumWarningOsVersion", m.GetMinimumWarningOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumWipeAppVersion", m.GetMinimumWipeAppVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumWipeOsVersion", m.GetMinimumWipeOsVersion())
        if err != nil {
            return err
        }
    }
    if m.GetMobileThreatDefensePartnerPriority() != nil {
        cast := (*m.GetMobileThreatDefensePartnerPriority()).String()
        err = writer.WriteStringValue("mobileThreatDefensePartnerPriority", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileThreatDefenseRemediationAction() != nil {
        cast := (*m.GetMobileThreatDefenseRemediationAction()).String()
        err = writer.WriteStringValue("mobileThreatDefenseRemediationAction", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetNotificationRestriction() != nil {
        cast := (*m.GetNotificationRestriction()).String()
        err = writer.WriteStringValue("notificationRestriction", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("organizationalCredentialsRequired", m.GetOrganizationalCredentialsRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("periodBeforePinReset", m.GetPeriodBeforePinReset())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("periodOfflineBeforeAccessCheck", m.GetPeriodOfflineBeforeAccessCheck())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("periodOfflineBeforeWipeIsEnforced", m.GetPeriodOfflineBeforeWipeIsEnforced())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("periodOnlineBeforeAccessCheck", m.GetPeriodOnlineBeforeAccessCheck())
        if err != nil {
            return err
        }
    }
    if m.GetPinCharacterSet() != nil {
        cast := (*m.GetPinCharacterSet()).String()
        err = writer.WriteStringValue("pinCharacterSet", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("pinRequired", m.GetPinRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("pinRequiredInsteadOfBiometricTimeout", m.GetPinRequiredInsteadOfBiometricTimeout())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("previousPinBlockCount", m.GetPreviousPinBlockCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("printBlocked", m.GetPrintBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("saveAsBlocked", m.GetSaveAsBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("simplePinBlocked", m.GetSimplePinBlocked())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedDataIngestionLocations sets the allowedDataIngestionLocations property value. Data storage locations where a user may store managed data.
func (m *ManagedAppProtection) SetAllowedDataIngestionLocations(value []ManagedAppDataIngestionLocation)() {
    m.allowedDataIngestionLocations = value
}
// SetAllowedDataStorageLocations sets the allowedDataStorageLocations property value. Data storage locations where a user may store managed data.
func (m *ManagedAppProtection) SetAllowedDataStorageLocations(value []ManagedAppDataStorageLocation)() {
    m.allowedDataStorageLocations = value
}
// SetAllowedInboundDataTransferSources sets the allowedInboundDataTransferSources property value. Data can be transferred from/to these classes of apps
func (m *ManagedAppProtection) SetAllowedInboundDataTransferSources(value *ManagedAppDataTransferLevel)() {
    m.allowedInboundDataTransferSources = value
}
// SetAllowedOutboundClipboardSharingExceptionLength sets the allowedOutboundClipboardSharingExceptionLength property value. Specify the number of characters that may be cut or copied from Org data and accounts to any application. This setting overrides the AllowedOutboundClipboardSharingLevel restriction. Default value of '0' means no exception is allowed.
func (m *ManagedAppProtection) SetAllowedOutboundClipboardSharingExceptionLength(value *int32)() {
    m.allowedOutboundClipboardSharingExceptionLength = value
}
// SetAllowedOutboundClipboardSharingLevel sets the allowedOutboundClipboardSharingLevel property value. Represents the level to which the device's clipboard may be shared between apps
func (m *ManagedAppProtection) SetAllowedOutboundClipboardSharingLevel(value *ManagedAppClipboardSharingLevel)() {
    m.allowedOutboundClipboardSharingLevel = value
}
// SetAllowedOutboundDataTransferDestinations sets the allowedOutboundDataTransferDestinations property value. Data can be transferred from/to these classes of apps
func (m *ManagedAppProtection) SetAllowedOutboundDataTransferDestinations(value *ManagedAppDataTransferLevel)() {
    m.allowedOutboundDataTransferDestinations = value
}
// SetAppActionIfDeviceComplianceRequired sets the appActionIfDeviceComplianceRequired property value. An admin initiated action to be applied on a managed app.
func (m *ManagedAppProtection) SetAppActionIfDeviceComplianceRequired(value *ManagedAppRemediationAction)() {
    m.appActionIfDeviceComplianceRequired = value
}
// SetAppActionIfMaximumPinRetriesExceeded sets the appActionIfMaximumPinRetriesExceeded property value. An admin initiated action to be applied on a managed app.
func (m *ManagedAppProtection) SetAppActionIfMaximumPinRetriesExceeded(value *ManagedAppRemediationAction)() {
    m.appActionIfMaximumPinRetriesExceeded = value
}
// SetAppActionIfUnableToAuthenticateUser sets the appActionIfUnableToAuthenticateUser property value. If set, it will specify what action to take in the case where the user is unable to checkin because their authentication token is invalid. This happens when the user is deleted or disabled in AAD. Possible values are: block, wipe, warn.
func (m *ManagedAppProtection) SetAppActionIfUnableToAuthenticateUser(value *ManagedAppRemediationAction)() {
    m.appActionIfUnableToAuthenticateUser = value
}
// SetBlockDataIngestionIntoOrganizationDocuments sets the blockDataIngestionIntoOrganizationDocuments property value. Indicates whether a user can bring data into org documents.
func (m *ManagedAppProtection) SetBlockDataIngestionIntoOrganizationDocuments(value *bool)() {
    m.blockDataIngestionIntoOrganizationDocuments = value
}
// SetContactSyncBlocked sets the contactSyncBlocked property value. Indicates whether contacts can be synced to the user's device.
func (m *ManagedAppProtection) SetContactSyncBlocked(value *bool)() {
    m.contactSyncBlocked = value
}
// SetDataBackupBlocked sets the dataBackupBlocked property value. Indicates whether the backup of a managed app's data is blocked.
func (m *ManagedAppProtection) SetDataBackupBlocked(value *bool)() {
    m.dataBackupBlocked = value
}
// SetDeviceComplianceRequired sets the deviceComplianceRequired property value. Indicates whether device compliance is required.
func (m *ManagedAppProtection) SetDeviceComplianceRequired(value *bool)() {
    m.deviceComplianceRequired = value
}
// SetDialerRestrictionLevel sets the dialerRestrictionLevel property value. The classes of apps that are allowed to click-to-open a phone number, for making phone calls or sending text messages.
func (m *ManagedAppProtection) SetDialerRestrictionLevel(value *ManagedAppPhoneNumberRedirectLevel)() {
    m.dialerRestrictionLevel = value
}
// SetDisableAppPinIfDevicePinIsSet sets the disableAppPinIfDevicePinIsSet property value. Indicates whether use of the app pin is required if the device pin is set.
func (m *ManagedAppProtection) SetDisableAppPinIfDevicePinIsSet(value *bool)() {
    m.disableAppPinIfDevicePinIsSet = value
}
// SetFingerprintBlocked sets the fingerprintBlocked property value. Indicates whether use of the fingerprint reader is allowed in place of a pin if PinRequired is set to True.
func (m *ManagedAppProtection) SetFingerprintBlocked(value *bool)() {
    m.fingerprintBlocked = value
}
// SetGracePeriodToBlockAppsDuringOffClockHours sets the gracePeriodToBlockAppsDuringOffClockHours property value. A grace period before blocking app access during off clock hours.
func (m *ManagedAppProtection) SetGracePeriodToBlockAppsDuringOffClockHours(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.gracePeriodToBlockAppsDuringOffClockHours = value
}
// SetManagedBrowser sets the managedBrowser property value. Type of managed browser
func (m *ManagedAppProtection) SetManagedBrowser(value *ManagedBrowserType)() {
    m.managedBrowser = value
}
// SetManagedBrowserToOpenLinksRequired sets the managedBrowserToOpenLinksRequired property value. Indicates whether internet links should be opened in the managed browser app, or any custom browser specified by CustomBrowserProtocol (for iOS) or CustomBrowserPackageId/CustomBrowserDisplayName (for Android)
func (m *ManagedAppProtection) SetManagedBrowserToOpenLinksRequired(value *bool)() {
    m.managedBrowserToOpenLinksRequired = value
}
// SetMaximumAllowedDeviceThreatLevel sets the maximumAllowedDeviceThreatLevel property value. The maxium threat level allowed for an app to be compliant.
func (m *ManagedAppProtection) SetMaximumAllowedDeviceThreatLevel(value *ManagedAppDeviceThreatLevel)() {
    m.maximumAllowedDeviceThreatLevel = value
}
// SetMaximumPinRetries sets the maximumPinRetries property value. Maximum number of incorrect pin retry attempts before the managed app is either blocked or wiped.
func (m *ManagedAppProtection) SetMaximumPinRetries(value *int32)() {
    m.maximumPinRetries = value
}
// SetMaximumRequiredOsVersion sets the maximumRequiredOsVersion property value. Versions bigger than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) SetMaximumRequiredOsVersion(value *string)() {
    m.maximumRequiredOsVersion = value
}
// SetMaximumWarningOsVersion sets the maximumWarningOsVersion property value. Versions bigger than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) SetMaximumWarningOsVersion(value *string)() {
    m.maximumWarningOsVersion = value
}
// SetMaximumWipeOsVersion sets the maximumWipeOsVersion property value. Versions bigger than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) SetMaximumWipeOsVersion(value *string)() {
    m.maximumWipeOsVersion = value
}
// SetMinimumPinLength sets the minimumPinLength property value. Minimum pin length required for an app-level pin if PinRequired is set to True
func (m *ManagedAppProtection) SetMinimumPinLength(value *int32)() {
    m.minimumPinLength = value
}
// SetMinimumRequiredAppVersion sets the minimumRequiredAppVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) SetMinimumRequiredAppVersion(value *string)() {
    m.minimumRequiredAppVersion = value
}
// SetMinimumRequiredOsVersion sets the minimumRequiredOsVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) SetMinimumRequiredOsVersion(value *string)() {
    m.minimumRequiredOsVersion = value
}
// SetMinimumWarningAppVersion sets the minimumWarningAppVersion property value. Versions less than the specified version will result in warning message on the managed app.
func (m *ManagedAppProtection) SetMinimumWarningAppVersion(value *string)() {
    m.minimumWarningAppVersion = value
}
// SetMinimumWarningOsVersion sets the minimumWarningOsVersion property value. Versions less than the specified version will result in warning message on the managed app from accessing company data.
func (m *ManagedAppProtection) SetMinimumWarningOsVersion(value *string)() {
    m.minimumWarningOsVersion = value
}
// SetMinimumWipeAppVersion sets the minimumWipeAppVersion property value. Versions less than or equal to the specified version will wipe the managed app and the associated company data.
func (m *ManagedAppProtection) SetMinimumWipeAppVersion(value *string)() {
    m.minimumWipeAppVersion = value
}
// SetMinimumWipeOsVersion sets the minimumWipeOsVersion property value. Versions less than or equal to the specified version will wipe the managed app and the associated company data.
func (m *ManagedAppProtection) SetMinimumWipeOsVersion(value *string)() {
    m.minimumWipeOsVersion = value
}
// SetMobileThreatDefensePartnerPriority sets the mobileThreatDefensePartnerPriority property value. Indicates how to prioritize which Mobile Threat Defense (MTD) partner is enabled for a given platform, when more than one is enabled. An app can only be actively using a single Mobile Threat Defense partner. When NULL, Microsoft Defender will be given preference. Otherwise setting the value to defenderOverThirdPartyPartner or thirdPartyPartnerOverDefender will make explicit which partner to prioritize. Possible values are: null, defenderOverThirdPartyPartner, thirdPartyPartnerOverDefender and unknownFutureValue. Default value is null. Possible values are: defenderOverThirdPartyPartner, thirdPartyPartnerOverDefender, unknownFutureValue.
func (m *ManagedAppProtection) SetMobileThreatDefensePartnerPriority(value *MobileThreatDefensePartnerPriority)() {
    m.mobileThreatDefensePartnerPriority = value
}
// SetMobileThreatDefenseRemediationAction sets the mobileThreatDefenseRemediationAction property value. An admin initiated action to be applied on a managed app.
func (m *ManagedAppProtection) SetMobileThreatDefenseRemediationAction(value *ManagedAppRemediationAction)() {
    m.mobileThreatDefenseRemediationAction = value
}
// SetNotificationRestriction sets the notificationRestriction property value. Restrict managed app notification
func (m *ManagedAppProtection) SetNotificationRestriction(value *ManagedAppNotificationRestriction)() {
    m.notificationRestriction = value
}
// SetOrganizationalCredentialsRequired sets the organizationalCredentialsRequired property value. Indicates whether organizational credentials are required for app use.
func (m *ManagedAppProtection) SetOrganizationalCredentialsRequired(value *bool)() {
    m.organizationalCredentialsRequired = value
}
// SetPeriodBeforePinReset sets the periodBeforePinReset property value. TimePeriod before the all-level pin must be reset if PinRequired is set to True.
func (m *ManagedAppProtection) SetPeriodBeforePinReset(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.periodBeforePinReset = value
}
// SetPeriodOfflineBeforeAccessCheck sets the periodOfflineBeforeAccessCheck property value. The period after which access is checked when the device is not connected to the internet.
func (m *ManagedAppProtection) SetPeriodOfflineBeforeAccessCheck(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.periodOfflineBeforeAccessCheck = value
}
// SetPeriodOfflineBeforeWipeIsEnforced sets the periodOfflineBeforeWipeIsEnforced property value. The amount of time an app is allowed to remain disconnected from the internet before all managed data it is wiped.
func (m *ManagedAppProtection) SetPeriodOfflineBeforeWipeIsEnforced(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.periodOfflineBeforeWipeIsEnforced = value
}
// SetPeriodOnlineBeforeAccessCheck sets the periodOnlineBeforeAccessCheck property value. The period after which access is checked when the device is connected to the internet.
func (m *ManagedAppProtection) SetPeriodOnlineBeforeAccessCheck(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.periodOnlineBeforeAccessCheck = value
}
// SetPinCharacterSet sets the pinCharacterSet property value. Character set which is to be used for a user's app PIN
func (m *ManagedAppProtection) SetPinCharacterSet(value *ManagedAppPinCharacterSet)() {
    m.pinCharacterSet = value
}
// SetPinRequired sets the pinRequired property value. Indicates whether an app-level pin is required.
func (m *ManagedAppProtection) SetPinRequired(value *bool)() {
    m.pinRequired = value
}
// SetPinRequiredInsteadOfBiometricTimeout sets the pinRequiredInsteadOfBiometricTimeout property value. Timeout in minutes for an app pin instead of non biometrics passcode
func (m *ManagedAppProtection) SetPinRequiredInsteadOfBiometricTimeout(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.pinRequiredInsteadOfBiometricTimeout = value
}
// SetPreviousPinBlockCount sets the previousPinBlockCount property value. Requires a pin to be unique from the number specified in this property.
func (m *ManagedAppProtection) SetPreviousPinBlockCount(value *int32)() {
    m.previousPinBlockCount = value
}
// SetPrintBlocked sets the printBlocked property value. Indicates whether printing is allowed from managed apps.
func (m *ManagedAppProtection) SetPrintBlocked(value *bool)() {
    m.printBlocked = value
}
// SetSaveAsBlocked sets the saveAsBlocked property value. Indicates whether users may use the 'Save As' menu item to save a copy of protected files.
func (m *ManagedAppProtection) SetSaveAsBlocked(value *bool)() {
    m.saveAsBlocked = value
}
// SetSimplePinBlocked sets the simplePinBlocked property value. Indicates whether simplePin is blocked.
func (m *ManagedAppProtection) SetSimplePinBlocked(value *bool)() {
    m.simplePinBlocked = value
}
