package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosManagedAppProtection 
type IosManagedAppProtection struct {
    TargetedManagedAppProtection
    // Semicolon seperated list of device models allowed, as a string, for the managed app to work.
    allowedIosDeviceModels *string
    // An admin initiated action to be applied on a managed app.
    appActionIfIosDeviceModelNotAllowed *ManagedAppRemediationAction
    // Represents the level to which app data is encrypted for managed apps
    appDataEncryptionType *ManagedAppDataEncryptionType
    // List of apps to which the policy is deployed.
    apps []ManagedMobileAppable
    // A custom browser protocol to open weblink on iOS.
    customBrowserProtocol *string
    // Protocol of a custom dialer app to click-to-open a phone number on iOS, for example, skype:.
    customDialerAppProtocol *string
    // Count of apps to which the current policy is deployed.
    deployedAppCount *int32
    // Navigation property to deployment summary of the configuration.
    deploymentSummary ManagedAppPolicyDeploymentSummaryable
    // Disable protection of data transferred to other apps through IOS OpenIn option. This setting is only allowed to be True when AllowedOutboundDataTransferDestinations is set to ManagedApps.
    disableProtectionOfManagedOutboundOpenInData *bool
    // Apps in this list will be exempt from the policy and will be able to receive data from managed apps.
    exemptedAppProtocols []KeyValuePairable
    // A list of custom urls that are allowed to invocate an unmanaged app
    exemptedUniversalLinks []string
    // Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True.
    faceIdBlocked *bool
    // Defines if open-in operation is supported from the managed app to the filesharing locations selected. This setting only applies when AllowedOutboundDataTransferDestinations is set to ManagedApps and DisableProtectionOfManagedOutboundOpenInData is set to False.
    filterOpenInToOnlyManagedApps *bool
    // A list of custom urls that are allowed to invocate a managed app
    managedUniversalLinks []string
    // Versions less than the specified version will block the managed app from accessing company data.
    minimumRequiredSdkVersion *string
    // Versions less than the specified version will result in warning message on the managed app from accessing company data.
    minimumWarningSdkVersion *string
    // Versions less than the specified version will block the managed app from accessing company data.
    minimumWipeSdkVersion *string
    // Protect incoming data from unknown source. This setting is only allowed to be True when AllowedInboundDataTransferSources is set to AllApps.
    protectInboundDataFromUnknownSources *bool
    // Defines if third party keyboards are allowed while accessing a managed app
    thirdPartyKeyboardsBlocked *bool
}
// NewIosManagedAppProtection instantiates a new IosManagedAppProtection and sets the default values.
func NewIosManagedAppProtection()(*IosManagedAppProtection) {
    m := &IosManagedAppProtection{
        TargetedManagedAppProtection: *NewTargetedManagedAppProtection(),
    }
    odataTypeValue := "#microsoft.graph.iosManagedAppProtection";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateIosManagedAppProtectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosManagedAppProtectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosManagedAppProtection(), nil
}
// GetAllowedIosDeviceModels gets the allowedIosDeviceModels property value. Semicolon seperated list of device models allowed, as a string, for the managed app to work.
func (m *IosManagedAppProtection) GetAllowedIosDeviceModels()(*string) {
    return m.allowedIosDeviceModels
}
// GetAppActionIfIosDeviceModelNotAllowed gets the appActionIfIosDeviceModelNotAllowed property value. An admin initiated action to be applied on a managed app.
func (m *IosManagedAppProtection) GetAppActionIfIosDeviceModelNotAllowed()(*ManagedAppRemediationAction) {
    return m.appActionIfIosDeviceModelNotAllowed
}
// GetAppDataEncryptionType gets the appDataEncryptionType property value. Represents the level to which app data is encrypted for managed apps
func (m *IosManagedAppProtection) GetAppDataEncryptionType()(*ManagedAppDataEncryptionType) {
    return m.appDataEncryptionType
}
// GetApps gets the apps property value. List of apps to which the policy is deployed.
func (m *IosManagedAppProtection) GetApps()([]ManagedMobileAppable) {
    return m.apps
}
// GetCustomBrowserProtocol gets the customBrowserProtocol property value. A custom browser protocol to open weblink on iOS.
func (m *IosManagedAppProtection) GetCustomBrowserProtocol()(*string) {
    return m.customBrowserProtocol
}
// GetCustomDialerAppProtocol gets the customDialerAppProtocol property value. Protocol of a custom dialer app to click-to-open a phone number on iOS, for example, skype:.
func (m *IosManagedAppProtection) GetCustomDialerAppProtocol()(*string) {
    return m.customDialerAppProtocol
}
// GetDeployedAppCount gets the deployedAppCount property value. Count of apps to which the current policy is deployed.
func (m *IosManagedAppProtection) GetDeployedAppCount()(*int32) {
    return m.deployedAppCount
}
// GetDeploymentSummary gets the deploymentSummary property value. Navigation property to deployment summary of the configuration.
func (m *IosManagedAppProtection) GetDeploymentSummary()(ManagedAppPolicyDeploymentSummaryable) {
    return m.deploymentSummary
}
// GetDisableProtectionOfManagedOutboundOpenInData gets the disableProtectionOfManagedOutboundOpenInData property value. Disable protection of data transferred to other apps through IOS OpenIn option. This setting is only allowed to be True when AllowedOutboundDataTransferDestinations is set to ManagedApps.
func (m *IosManagedAppProtection) GetDisableProtectionOfManagedOutboundOpenInData()(*bool) {
    return m.disableProtectionOfManagedOutboundOpenInData
}
// GetExemptedAppProtocols gets the exemptedAppProtocols property value. Apps in this list will be exempt from the policy and will be able to receive data from managed apps.
func (m *IosManagedAppProtection) GetExemptedAppProtocols()([]KeyValuePairable) {
    return m.exemptedAppProtocols
}
// GetExemptedUniversalLinks gets the exemptedUniversalLinks property value. A list of custom urls that are allowed to invocate an unmanaged app
func (m *IosManagedAppProtection) GetExemptedUniversalLinks()([]string) {
    return m.exemptedUniversalLinks
}
// GetFaceIdBlocked gets the faceIdBlocked property value. Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True.
func (m *IosManagedAppProtection) GetFaceIdBlocked()(*bool) {
    return m.faceIdBlocked
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosManagedAppProtection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TargetedManagedAppProtection.GetFieldDeserializers()
    res["allowedIosDeviceModels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedIosDeviceModels(val)
        }
        return nil
    }
    res["appActionIfIosDeviceModelNotAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppRemediationAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppActionIfIosDeviceModelNotAllowed(val.(*ManagedAppRemediationAction))
        }
        return nil
    }
    res["appDataEncryptionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppDataEncryptionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDataEncryptionType(val.(*ManagedAppDataEncryptionType))
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
    res["customBrowserProtocol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomBrowserProtocol(val)
        }
        return nil
    }
    res["customDialerAppProtocol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomDialerAppProtocol(val)
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
    res["deploymentSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagedAppPolicyDeploymentSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentSummary(val.(ManagedAppPolicyDeploymentSummaryable))
        }
        return nil
    }
    res["disableProtectionOfManagedOutboundOpenInData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableProtectionOfManagedOutboundOpenInData(val)
        }
        return nil
    }
    res["exemptedAppProtocols"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(KeyValuePairable)
            }
            m.SetExemptedAppProtocols(res)
        }
        return nil
    }
    res["exemptedUniversalLinks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExemptedUniversalLinks(res)
        }
        return nil
    }
    res["faceIdBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFaceIdBlocked(val)
        }
        return nil
    }
    res["filterOpenInToOnlyManagedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilterOpenInToOnlyManagedApps(val)
        }
        return nil
    }
    res["managedUniversalLinks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetManagedUniversalLinks(res)
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
    res["minimumWarningSdkVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWarningSdkVersion(val)
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
    res["protectInboundDataFromUnknownSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtectInboundDataFromUnknownSources(val)
        }
        return nil
    }
    res["thirdPartyKeyboardsBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThirdPartyKeyboardsBlocked(val)
        }
        return nil
    }
    return res
}
// GetFilterOpenInToOnlyManagedApps gets the filterOpenInToOnlyManagedApps property value. Defines if open-in operation is supported from the managed app to the filesharing locations selected. This setting only applies when AllowedOutboundDataTransferDestinations is set to ManagedApps and DisableProtectionOfManagedOutboundOpenInData is set to False.
func (m *IosManagedAppProtection) GetFilterOpenInToOnlyManagedApps()(*bool) {
    return m.filterOpenInToOnlyManagedApps
}
// GetManagedUniversalLinks gets the managedUniversalLinks property value. A list of custom urls that are allowed to invocate a managed app
func (m *IosManagedAppProtection) GetManagedUniversalLinks()([]string) {
    return m.managedUniversalLinks
}
// GetMinimumRequiredSdkVersion gets the minimumRequiredSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *IosManagedAppProtection) GetMinimumRequiredSdkVersion()(*string) {
    return m.minimumRequiredSdkVersion
}
// GetMinimumWarningSdkVersion gets the minimumWarningSdkVersion property value. Versions less than the specified version will result in warning message on the managed app from accessing company data.
func (m *IosManagedAppProtection) GetMinimumWarningSdkVersion()(*string) {
    return m.minimumWarningSdkVersion
}
// GetMinimumWipeSdkVersion gets the minimumWipeSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *IosManagedAppProtection) GetMinimumWipeSdkVersion()(*string) {
    return m.minimumWipeSdkVersion
}
// GetProtectInboundDataFromUnknownSources gets the protectInboundDataFromUnknownSources property value. Protect incoming data from unknown source. This setting is only allowed to be True when AllowedInboundDataTransferSources is set to AllApps.
func (m *IosManagedAppProtection) GetProtectInboundDataFromUnknownSources()(*bool) {
    return m.protectInboundDataFromUnknownSources
}
// GetThirdPartyKeyboardsBlocked gets the thirdPartyKeyboardsBlocked property value. Defines if third party keyboards are allowed while accessing a managed app
func (m *IosManagedAppProtection) GetThirdPartyKeyboardsBlocked()(*bool) {
    return m.thirdPartyKeyboardsBlocked
}
// Serialize serializes information the current object
func (m *IosManagedAppProtection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TargetedManagedAppProtection.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("allowedIosDeviceModels", m.GetAllowedIosDeviceModels())
        if err != nil {
            return err
        }
    }
    if m.GetAppActionIfIosDeviceModelNotAllowed() != nil {
        cast := (*m.GetAppActionIfIosDeviceModelNotAllowed()).String()
        err = writer.WriteStringValue("appActionIfIosDeviceModelNotAllowed", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAppDataEncryptionType() != nil {
        cast := (*m.GetAppDataEncryptionType()).String()
        err = writer.WriteStringValue("appDataEncryptionType", &cast)
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
    {
        err = writer.WriteStringValue("customBrowserProtocol", m.GetCustomBrowserProtocol())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customDialerAppProtocol", m.GetCustomDialerAppProtocol())
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
        err = writer.WriteObjectValue("deploymentSummary", m.GetDeploymentSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disableProtectionOfManagedOutboundOpenInData", m.GetDisableProtectionOfManagedOutboundOpenInData())
        if err != nil {
            return err
        }
    }
    if m.GetExemptedAppProtocols() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExemptedAppProtocols()))
        for i, v := range m.GetExemptedAppProtocols() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("exemptedAppProtocols", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExemptedUniversalLinks() != nil {
        err = writer.WriteCollectionOfStringValues("exemptedUniversalLinks", m.GetExemptedUniversalLinks())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("faceIdBlocked", m.GetFaceIdBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("filterOpenInToOnlyManagedApps", m.GetFilterOpenInToOnlyManagedApps())
        if err != nil {
            return err
        }
    }
    if m.GetManagedUniversalLinks() != nil {
        err = writer.WriteCollectionOfStringValues("managedUniversalLinks", m.GetManagedUniversalLinks())
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
        err = writer.WriteStringValue("minimumWarningSdkVersion", m.GetMinimumWarningSdkVersion())
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
    {
        err = writer.WriteBoolValue("protectInboundDataFromUnknownSources", m.GetProtectInboundDataFromUnknownSources())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("thirdPartyKeyboardsBlocked", m.GetThirdPartyKeyboardsBlocked())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedIosDeviceModels sets the allowedIosDeviceModels property value. Semicolon seperated list of device models allowed, as a string, for the managed app to work.
func (m *IosManagedAppProtection) SetAllowedIosDeviceModels(value *string)() {
    m.allowedIosDeviceModels = value
}
// SetAppActionIfIosDeviceModelNotAllowed sets the appActionIfIosDeviceModelNotAllowed property value. An admin initiated action to be applied on a managed app.
func (m *IosManagedAppProtection) SetAppActionIfIosDeviceModelNotAllowed(value *ManagedAppRemediationAction)() {
    m.appActionIfIosDeviceModelNotAllowed = value
}
// SetAppDataEncryptionType sets the appDataEncryptionType property value. Represents the level to which app data is encrypted for managed apps
func (m *IosManagedAppProtection) SetAppDataEncryptionType(value *ManagedAppDataEncryptionType)() {
    m.appDataEncryptionType = value
}
// SetApps sets the apps property value. List of apps to which the policy is deployed.
func (m *IosManagedAppProtection) SetApps(value []ManagedMobileAppable)() {
    m.apps = value
}
// SetCustomBrowserProtocol sets the customBrowserProtocol property value. A custom browser protocol to open weblink on iOS.
func (m *IosManagedAppProtection) SetCustomBrowserProtocol(value *string)() {
    m.customBrowserProtocol = value
}
// SetCustomDialerAppProtocol sets the customDialerAppProtocol property value. Protocol of a custom dialer app to click-to-open a phone number on iOS, for example, skype:.
func (m *IosManagedAppProtection) SetCustomDialerAppProtocol(value *string)() {
    m.customDialerAppProtocol = value
}
// SetDeployedAppCount sets the deployedAppCount property value. Count of apps to which the current policy is deployed.
func (m *IosManagedAppProtection) SetDeployedAppCount(value *int32)() {
    m.deployedAppCount = value
}
// SetDeploymentSummary sets the deploymentSummary property value. Navigation property to deployment summary of the configuration.
func (m *IosManagedAppProtection) SetDeploymentSummary(value ManagedAppPolicyDeploymentSummaryable)() {
    m.deploymentSummary = value
}
// SetDisableProtectionOfManagedOutboundOpenInData sets the disableProtectionOfManagedOutboundOpenInData property value. Disable protection of data transferred to other apps through IOS OpenIn option. This setting is only allowed to be True when AllowedOutboundDataTransferDestinations is set to ManagedApps.
func (m *IosManagedAppProtection) SetDisableProtectionOfManagedOutboundOpenInData(value *bool)() {
    m.disableProtectionOfManagedOutboundOpenInData = value
}
// SetExemptedAppProtocols sets the exemptedAppProtocols property value. Apps in this list will be exempt from the policy and will be able to receive data from managed apps.
func (m *IosManagedAppProtection) SetExemptedAppProtocols(value []KeyValuePairable)() {
    m.exemptedAppProtocols = value
}
// SetExemptedUniversalLinks sets the exemptedUniversalLinks property value. A list of custom urls that are allowed to invocate an unmanaged app
func (m *IosManagedAppProtection) SetExemptedUniversalLinks(value []string)() {
    m.exemptedUniversalLinks = value
}
// SetFaceIdBlocked sets the faceIdBlocked property value. Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True.
func (m *IosManagedAppProtection) SetFaceIdBlocked(value *bool)() {
    m.faceIdBlocked = value
}
// SetFilterOpenInToOnlyManagedApps sets the filterOpenInToOnlyManagedApps property value. Defines if open-in operation is supported from the managed app to the filesharing locations selected. This setting only applies when AllowedOutboundDataTransferDestinations is set to ManagedApps and DisableProtectionOfManagedOutboundOpenInData is set to False.
func (m *IosManagedAppProtection) SetFilterOpenInToOnlyManagedApps(value *bool)() {
    m.filterOpenInToOnlyManagedApps = value
}
// SetManagedUniversalLinks sets the managedUniversalLinks property value. A list of custom urls that are allowed to invocate a managed app
func (m *IosManagedAppProtection) SetManagedUniversalLinks(value []string)() {
    m.managedUniversalLinks = value
}
// SetMinimumRequiredSdkVersion sets the minimumRequiredSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *IosManagedAppProtection) SetMinimumRequiredSdkVersion(value *string)() {
    m.minimumRequiredSdkVersion = value
}
// SetMinimumWarningSdkVersion sets the minimumWarningSdkVersion property value. Versions less than the specified version will result in warning message on the managed app from accessing company data.
func (m *IosManagedAppProtection) SetMinimumWarningSdkVersion(value *string)() {
    m.minimumWarningSdkVersion = value
}
// SetMinimumWipeSdkVersion sets the minimumWipeSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *IosManagedAppProtection) SetMinimumWipeSdkVersion(value *string)() {
    m.minimumWipeSdkVersion = value
}
// SetProtectInboundDataFromUnknownSources sets the protectInboundDataFromUnknownSources property value. Protect incoming data from unknown source. This setting is only allowed to be True when AllowedInboundDataTransferSources is set to AllApps.
func (m *IosManagedAppProtection) SetProtectInboundDataFromUnknownSources(value *bool)() {
    m.protectInboundDataFromUnknownSources = value
}
// SetThirdPartyKeyboardsBlocked sets the thirdPartyKeyboardsBlocked property value. Defines if third party keyboards are allowed while accessing a managed app
func (m *IosManagedAppProtection) SetThirdPartyKeyboardsBlocked(value *bool)() {
    m.thirdPartyKeyboardsBlocked = value
}
