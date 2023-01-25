package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsManagedAppProtection 
type WindowsManagedAppProtection struct {
    ManagedAppPolicy
    // Data can be transferred from/to these classes of apps
    allowedInboundDataTransferSources *WindowsManagedAppDataTransferLevel
    // Represents the level to which the device's clipboard may be shared between apps
    allowedOutboundClipboardSharingLevel *WindowsManagedAppClipboardSharingLevel
    // Data can be transferred from/to these classes of apps
    allowedOutboundDataTransferDestinations *WindowsManagedAppDataTransferLevel
    // If set, it will specify what action to take in the case where the user is unable to checkin because their authentication token is invalid. This happens when the user is deleted or disabled in AAD. Some possible values are block or wipe. If this property is not set, no action will be taken. Possible values are: block, wipe, warn.
    appActionIfUnableToAuthenticateUser *ManagedAppRemediationAction
    // List of apps to which the policy is deployed.
    apps []ManagedMobileAppable
    // Navigation property to list of inclusion and exclusion groups to which the policy is deployed.
    assignments []TargetedManagedAppPolicyAssignmentable
    // Indicates the total number of applications for which the current policy is deployed.
    deployedAppCount *int32
    // When TRUE, indicates that the policy is deployed to some inclusion groups. When FALSE, indicates that the policy is not deployed to any inclusion groups. Default value is FALSE.
    isAssigned *bool
    // The maxium threat level allowed for an app to be compliant.
    maximumAllowedDeviceThreatLevel *ManagedAppDeviceThreatLevel
    // Versions bigger than the specified version will block the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
    maximumRequiredOsVersion *string
    // Versions bigger than the specified version will result in warning message on the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
    maximumWarningOsVersion *string
    // Versions bigger than the specified version will wipe the managed app and the associated company data. For example: '8.1.0' or '13.1.1'.
    maximumWipeOsVersion *string
    // Versions less than the specified version will block the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
    minimumRequiredAppVersion *string
    // Versions less than the specified version will block the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
    minimumRequiredOsVersion *string
    // Versions less than the specified version will block the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
    minimumRequiredSdkVersion *string
    // Versions less than the specified version will result in warning message on the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
    minimumWarningAppVersion *string
    // Versions less than the specified version will result in warning message on the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
    minimumWarningOsVersion *string
    // Versions less than the specified version will wipe the managed app and the associated company data. For example: '8.1.0' or '13.1.1'.
    minimumWipeAppVersion *string
    // Versions less than the specified version will wipe the managed app and the associated company data. For example: '8.1.0' or '13.1.1'.
    minimumWipeOsVersion *string
    // Versions less than the specified version will wipe the managed app and the associated company data. For example: '8.1.0' or '13.1.1'.
    minimumWipeSdkVersion *string
    // An admin initiated action to be applied on a managed app.
    mobileThreatDefenseRemediationAction *ManagedAppRemediationAction
    // The period after which access is checked when the device is not connected to the internet. For example, PT5M indicates that the interval is 5 minutes in duration. A timespan value of PT0S indicates that access will be blocked immediately when the device is not connected to the internet.
    periodOfflineBeforeAccessCheck *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The amount of time an app is allowed to remain disconnected from the internet before all managed data it is wiped. For example, P5D indicates that the interval is 5 days in duration. A timespan value of PT0S indicates that managed data will never be wiped when the device is not connected to the internet.
    periodOfflineBeforeWipeIsEnforced *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // When TRUE, indicates that printing is blocked from managed apps. When FALSE, indicates that printing is allowed from managed apps. Default value is FALSE.
    printBlocked *bool
}
// NewWindowsManagedAppProtection instantiates a new WindowsManagedAppProtection and sets the default values.
func NewWindowsManagedAppProtection()(*WindowsManagedAppProtection) {
    m := &WindowsManagedAppProtection{
        ManagedAppPolicy: *NewManagedAppPolicy(),
    }
    odataTypeValue := "#microsoft.graph.windowsManagedAppProtection";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsManagedAppProtectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsManagedAppProtectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsManagedAppProtection(), nil
}
// GetAllowedInboundDataTransferSources gets the allowedInboundDataTransferSources property value. Data can be transferred from/to these classes of apps
func (m *WindowsManagedAppProtection) GetAllowedInboundDataTransferSources()(*WindowsManagedAppDataTransferLevel) {
    return m.allowedInboundDataTransferSources
}
// GetAllowedOutboundClipboardSharingLevel gets the allowedOutboundClipboardSharingLevel property value. Represents the level to which the device's clipboard may be shared between apps
func (m *WindowsManagedAppProtection) GetAllowedOutboundClipboardSharingLevel()(*WindowsManagedAppClipboardSharingLevel) {
    return m.allowedOutboundClipboardSharingLevel
}
// GetAllowedOutboundDataTransferDestinations gets the allowedOutboundDataTransferDestinations property value. Data can be transferred from/to these classes of apps
func (m *WindowsManagedAppProtection) GetAllowedOutboundDataTransferDestinations()(*WindowsManagedAppDataTransferLevel) {
    return m.allowedOutboundDataTransferDestinations
}
// GetAppActionIfUnableToAuthenticateUser gets the appActionIfUnableToAuthenticateUser property value. If set, it will specify what action to take in the case where the user is unable to checkin because their authentication token is invalid. This happens when the user is deleted or disabled in AAD. Some possible values are block or wipe. If this property is not set, no action will be taken. Possible values are: block, wipe, warn.
func (m *WindowsManagedAppProtection) GetAppActionIfUnableToAuthenticateUser()(*ManagedAppRemediationAction) {
    return m.appActionIfUnableToAuthenticateUser
}
// GetApps gets the apps property value. List of apps to which the policy is deployed.
func (m *WindowsManagedAppProtection) GetApps()([]ManagedMobileAppable) {
    return m.apps
}
// GetAssignments gets the assignments property value. Navigation property to list of inclusion and exclusion groups to which the policy is deployed.
func (m *WindowsManagedAppProtection) GetAssignments()([]TargetedManagedAppPolicyAssignmentable) {
    return m.assignments
}
// GetDeployedAppCount gets the deployedAppCount property value. Indicates the total number of applications for which the current policy is deployed.
func (m *WindowsManagedAppProtection) GetDeployedAppCount()(*int32) {
    return m.deployedAppCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsManagedAppProtection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ManagedAppPolicy.GetFieldDeserializers()
    res["allowedInboundDataTransferSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsManagedAppDataTransferLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedInboundDataTransferSources(val.(*WindowsManagedAppDataTransferLevel))
        }
        return nil
    }
    res["allowedOutboundClipboardSharingLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsManagedAppClipboardSharingLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedOutboundClipboardSharingLevel(val.(*WindowsManagedAppClipboardSharingLevel))
        }
        return nil
    }
    res["allowedOutboundDataTransferDestinations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsManagedAppDataTransferLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedOutboundDataTransferDestinations(val.(*WindowsManagedAppDataTransferLevel))
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
    res["apps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedMobileAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedMobileAppable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedMobileAppable)
            }
            m.SetApps(res)
        }
        return nil
    }
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTargetedManagedAppPolicyAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TargetedManagedAppPolicyAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(TargetedManagedAppPolicyAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["deployedAppCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeployedAppCount(val)
        }
        return nil
    }
    res["isAssigned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAssigned(val)
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
    res["minimumRequiredSdkVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumRequiredSdkVersion(val)
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
    res["minimumWipeSdkVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWipeSdkVersion(val)
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
    return res
}
// GetIsAssigned gets the isAssigned property value. When TRUE, indicates that the policy is deployed to some inclusion groups. When FALSE, indicates that the policy is not deployed to any inclusion groups. Default value is FALSE.
func (m *WindowsManagedAppProtection) GetIsAssigned()(*bool) {
    return m.isAssigned
}
// GetMaximumAllowedDeviceThreatLevel gets the maximumAllowedDeviceThreatLevel property value. The maxium threat level allowed for an app to be compliant.
func (m *WindowsManagedAppProtection) GetMaximumAllowedDeviceThreatLevel()(*ManagedAppDeviceThreatLevel) {
    return m.maximumAllowedDeviceThreatLevel
}
// GetMaximumRequiredOsVersion gets the maximumRequiredOsVersion property value. Versions bigger than the specified version will block the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) GetMaximumRequiredOsVersion()(*string) {
    return m.maximumRequiredOsVersion
}
// GetMaximumWarningOsVersion gets the maximumWarningOsVersion property value. Versions bigger than the specified version will result in warning message on the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) GetMaximumWarningOsVersion()(*string) {
    return m.maximumWarningOsVersion
}
// GetMaximumWipeOsVersion gets the maximumWipeOsVersion property value. Versions bigger than the specified version will wipe the managed app and the associated company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) GetMaximumWipeOsVersion()(*string) {
    return m.maximumWipeOsVersion
}
// GetMinimumRequiredAppVersion gets the minimumRequiredAppVersion property value. Versions less than the specified version will block the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) GetMinimumRequiredAppVersion()(*string) {
    return m.minimumRequiredAppVersion
}
// GetMinimumRequiredOsVersion gets the minimumRequiredOsVersion property value. Versions less than the specified version will block the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) GetMinimumRequiredOsVersion()(*string) {
    return m.minimumRequiredOsVersion
}
// GetMinimumRequiredSdkVersion gets the minimumRequiredSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) GetMinimumRequiredSdkVersion()(*string) {
    return m.minimumRequiredSdkVersion
}
// GetMinimumWarningAppVersion gets the minimumWarningAppVersion property value. Versions less than the specified version will result in warning message on the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) GetMinimumWarningAppVersion()(*string) {
    return m.minimumWarningAppVersion
}
// GetMinimumWarningOsVersion gets the minimumWarningOsVersion property value. Versions less than the specified version will result in warning message on the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) GetMinimumWarningOsVersion()(*string) {
    return m.minimumWarningOsVersion
}
// GetMinimumWipeAppVersion gets the minimumWipeAppVersion property value. Versions less than the specified version will wipe the managed app and the associated company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) GetMinimumWipeAppVersion()(*string) {
    return m.minimumWipeAppVersion
}
// GetMinimumWipeOsVersion gets the minimumWipeOsVersion property value. Versions less than the specified version will wipe the managed app and the associated company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) GetMinimumWipeOsVersion()(*string) {
    return m.minimumWipeOsVersion
}
// GetMinimumWipeSdkVersion gets the minimumWipeSdkVersion property value. Versions less than the specified version will wipe the managed app and the associated company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) GetMinimumWipeSdkVersion()(*string) {
    return m.minimumWipeSdkVersion
}
// GetMobileThreatDefenseRemediationAction gets the mobileThreatDefenseRemediationAction property value. An admin initiated action to be applied on a managed app.
func (m *WindowsManagedAppProtection) GetMobileThreatDefenseRemediationAction()(*ManagedAppRemediationAction) {
    return m.mobileThreatDefenseRemediationAction
}
// GetPeriodOfflineBeforeAccessCheck gets the periodOfflineBeforeAccessCheck property value. The period after which access is checked when the device is not connected to the internet. For example, PT5M indicates that the interval is 5 minutes in duration. A timespan value of PT0S indicates that access will be blocked immediately when the device is not connected to the internet.
func (m *WindowsManagedAppProtection) GetPeriodOfflineBeforeAccessCheck()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.periodOfflineBeforeAccessCheck
}
// GetPeriodOfflineBeforeWipeIsEnforced gets the periodOfflineBeforeWipeIsEnforced property value. The amount of time an app is allowed to remain disconnected from the internet before all managed data it is wiped. For example, P5D indicates that the interval is 5 days in duration. A timespan value of PT0S indicates that managed data will never be wiped when the device is not connected to the internet.
func (m *WindowsManagedAppProtection) GetPeriodOfflineBeforeWipeIsEnforced()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.periodOfflineBeforeWipeIsEnforced
}
// GetPrintBlocked gets the printBlocked property value. When TRUE, indicates that printing is blocked from managed apps. When FALSE, indicates that printing is allowed from managed apps. Default value is FALSE.
func (m *WindowsManagedAppProtection) GetPrintBlocked()(*bool) {
    return m.printBlocked
}
// Serialize serializes information the current object
func (m *WindowsManagedAppProtection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ManagedAppPolicy.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedInboundDataTransferSources() != nil {
        cast := (*m.GetAllowedInboundDataTransferSources()).String()
        err = writer.WriteStringValue("allowedInboundDataTransferSources", &cast)
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
    if m.GetAppActionIfUnableToAuthenticateUser() != nil {
        cast := (*m.GetAppActionIfUnableToAuthenticateUser()).String()
        err = writer.WriteStringValue("appActionIfUnableToAuthenticateUser", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetApps()))
        for i, v := range m.GetApps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("apps", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deployedAppCount", m.GetDeployedAppCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAssigned", m.GetIsAssigned())
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
        err = writer.WriteStringValue("minimumRequiredSdkVersion", m.GetMinimumRequiredSdkVersion())
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
    {
        err = writer.WriteStringValue("minimumWipeSdkVersion", m.GetMinimumWipeSdkVersion())
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
        err = writer.WriteBoolValue("printBlocked", m.GetPrintBlocked())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedInboundDataTransferSources sets the allowedInboundDataTransferSources property value. Data can be transferred from/to these classes of apps
func (m *WindowsManagedAppProtection) SetAllowedInboundDataTransferSources(value *WindowsManagedAppDataTransferLevel)() {
    m.allowedInboundDataTransferSources = value
}
// SetAllowedOutboundClipboardSharingLevel sets the allowedOutboundClipboardSharingLevel property value. Represents the level to which the device's clipboard may be shared between apps
func (m *WindowsManagedAppProtection) SetAllowedOutboundClipboardSharingLevel(value *WindowsManagedAppClipboardSharingLevel)() {
    m.allowedOutboundClipboardSharingLevel = value
}
// SetAllowedOutboundDataTransferDestinations sets the allowedOutboundDataTransferDestinations property value. Data can be transferred from/to these classes of apps
func (m *WindowsManagedAppProtection) SetAllowedOutboundDataTransferDestinations(value *WindowsManagedAppDataTransferLevel)() {
    m.allowedOutboundDataTransferDestinations = value
}
// SetAppActionIfUnableToAuthenticateUser sets the appActionIfUnableToAuthenticateUser property value. If set, it will specify what action to take in the case where the user is unable to checkin because their authentication token is invalid. This happens when the user is deleted or disabled in AAD. Some possible values are block or wipe. If this property is not set, no action will be taken. Possible values are: block, wipe, warn.
func (m *WindowsManagedAppProtection) SetAppActionIfUnableToAuthenticateUser(value *ManagedAppRemediationAction)() {
    m.appActionIfUnableToAuthenticateUser = value
}
// SetApps sets the apps property value. List of apps to which the policy is deployed.
func (m *WindowsManagedAppProtection) SetApps(value []ManagedMobileAppable)() {
    m.apps = value
}
// SetAssignments sets the assignments property value. Navigation property to list of inclusion and exclusion groups to which the policy is deployed.
func (m *WindowsManagedAppProtection) SetAssignments(value []TargetedManagedAppPolicyAssignmentable)() {
    m.assignments = value
}
// SetDeployedAppCount sets the deployedAppCount property value. Indicates the total number of applications for which the current policy is deployed.
func (m *WindowsManagedAppProtection) SetDeployedAppCount(value *int32)() {
    m.deployedAppCount = value
}
// SetIsAssigned sets the isAssigned property value. When TRUE, indicates that the policy is deployed to some inclusion groups. When FALSE, indicates that the policy is not deployed to any inclusion groups. Default value is FALSE.
func (m *WindowsManagedAppProtection) SetIsAssigned(value *bool)() {
    m.isAssigned = value
}
// SetMaximumAllowedDeviceThreatLevel sets the maximumAllowedDeviceThreatLevel property value. The maxium threat level allowed for an app to be compliant.
func (m *WindowsManagedAppProtection) SetMaximumAllowedDeviceThreatLevel(value *ManagedAppDeviceThreatLevel)() {
    m.maximumAllowedDeviceThreatLevel = value
}
// SetMaximumRequiredOsVersion sets the maximumRequiredOsVersion property value. Versions bigger than the specified version will block the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) SetMaximumRequiredOsVersion(value *string)() {
    m.maximumRequiredOsVersion = value
}
// SetMaximumWarningOsVersion sets the maximumWarningOsVersion property value. Versions bigger than the specified version will result in warning message on the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) SetMaximumWarningOsVersion(value *string)() {
    m.maximumWarningOsVersion = value
}
// SetMaximumWipeOsVersion sets the maximumWipeOsVersion property value. Versions bigger than the specified version will wipe the managed app and the associated company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) SetMaximumWipeOsVersion(value *string)() {
    m.maximumWipeOsVersion = value
}
// SetMinimumRequiredAppVersion sets the minimumRequiredAppVersion property value. Versions less than the specified version will block the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) SetMinimumRequiredAppVersion(value *string)() {
    m.minimumRequiredAppVersion = value
}
// SetMinimumRequiredOsVersion sets the minimumRequiredOsVersion property value. Versions less than the specified version will block the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) SetMinimumRequiredOsVersion(value *string)() {
    m.minimumRequiredOsVersion = value
}
// SetMinimumRequiredSdkVersion sets the minimumRequiredSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) SetMinimumRequiredSdkVersion(value *string)() {
    m.minimumRequiredSdkVersion = value
}
// SetMinimumWarningAppVersion sets the minimumWarningAppVersion property value. Versions less than the specified version will result in warning message on the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) SetMinimumWarningAppVersion(value *string)() {
    m.minimumWarningAppVersion = value
}
// SetMinimumWarningOsVersion sets the minimumWarningOsVersion property value. Versions less than the specified version will result in warning message on the managed app from accessing company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) SetMinimumWarningOsVersion(value *string)() {
    m.minimumWarningOsVersion = value
}
// SetMinimumWipeAppVersion sets the minimumWipeAppVersion property value. Versions less than the specified version will wipe the managed app and the associated company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) SetMinimumWipeAppVersion(value *string)() {
    m.minimumWipeAppVersion = value
}
// SetMinimumWipeOsVersion sets the minimumWipeOsVersion property value. Versions less than the specified version will wipe the managed app and the associated company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) SetMinimumWipeOsVersion(value *string)() {
    m.minimumWipeOsVersion = value
}
// SetMinimumWipeSdkVersion sets the minimumWipeSdkVersion property value. Versions less than the specified version will wipe the managed app and the associated company data. For example: '8.1.0' or '13.1.1'.
func (m *WindowsManagedAppProtection) SetMinimumWipeSdkVersion(value *string)() {
    m.minimumWipeSdkVersion = value
}
// SetMobileThreatDefenseRemediationAction sets the mobileThreatDefenseRemediationAction property value. An admin initiated action to be applied on a managed app.
func (m *WindowsManagedAppProtection) SetMobileThreatDefenseRemediationAction(value *ManagedAppRemediationAction)() {
    m.mobileThreatDefenseRemediationAction = value
}
// SetPeriodOfflineBeforeAccessCheck sets the periodOfflineBeforeAccessCheck property value. The period after which access is checked when the device is not connected to the internet. For example, PT5M indicates that the interval is 5 minutes in duration. A timespan value of PT0S indicates that access will be blocked immediately when the device is not connected to the internet.
func (m *WindowsManagedAppProtection) SetPeriodOfflineBeforeAccessCheck(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.periodOfflineBeforeAccessCheck = value
}
// SetPeriodOfflineBeforeWipeIsEnforced sets the periodOfflineBeforeWipeIsEnforced property value. The amount of time an app is allowed to remain disconnected from the internet before all managed data it is wiped. For example, P5D indicates that the interval is 5 days in duration. A timespan value of PT0S indicates that managed data will never be wiped when the device is not connected to the internet.
func (m *WindowsManagedAppProtection) SetPeriodOfflineBeforeWipeIsEnforced(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.periodOfflineBeforeWipeIsEnforced = value
}
// SetPrintBlocked sets the printBlocked property value. When TRUE, indicates that printing is blocked from managed apps. When FALSE, indicates that printing is allowed from managed apps. Default value is FALSE.
func (m *WindowsManagedAppProtection) SetPrintBlocked(value *bool)() {
    m.printBlocked = value
}
