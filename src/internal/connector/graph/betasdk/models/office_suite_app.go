package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OfficeSuiteApp 
type OfficeSuiteApp struct {
    MobileApp
    // The value to accept the EULA automatically on the enduser's device.
    autoAcceptEula *bool
    // The property to represent the apps which are excluded from the selected Office365 Product Id.
    excludedApps ExcludedAppsable
    // The Enum to specify the level of display for the Installation Progress Setup UI on the Device.
    installProgressDisplayLevel *OfficeSuiteInstallProgressDisplayLevel
    // The property to represent the locales which are installed when the apps from Office365 is installed. It uses standard RFC 6033. Ref: https://technet.microsoft.com/library/cc179219(v=office.16).aspx
    localesToInstall []string
    // The property to represent the XML configuration file that can be specified for Office ProPlus Apps. Takes precedence over all other properties. When present, the XML configuration file will be used to create the app.
    officeConfigurationXml []byte
    // Contains properties for Windows architecture.
    officePlatformArchitecture *WindowsArchitecture
    // Describes the OfficeSuiteApp file format types that can be selected.
    officeSuiteAppDefaultFileFormat *OfficeSuiteDefaultFileFormatType
    // The Product Ids that represent the Office365 Suite SKU.
    productIds []OfficeProductId
    // The property to determine whether to uninstall existing Office MSI if an Office365 app suite is deployed to the device or not.
    shouldUninstallOlderVersionsOfOffice *bool
    // The property to represent the specific target version for the Office365 app suite that should be remained deployed on the devices.
    targetVersion *string
    // The Enum to specify the Office365 Updates Channel.
    updateChannel *OfficeUpdateChannel
    // The property to represent the update version in which the specific target version is available for the Office365 app suite.
    updateVersion *string
    // The property to represent that whether the shared computer activation is used not for Office365 app suite.
    useSharedComputerActivation *bool
}
// NewOfficeSuiteApp instantiates a new OfficeSuiteApp and sets the default values.
func NewOfficeSuiteApp()(*OfficeSuiteApp) {
    m := &OfficeSuiteApp{
        MobileApp: *NewMobileApp(),
    }
    odataTypeValue := "#microsoft.graph.officeSuiteApp";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOfficeSuiteAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOfficeSuiteAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOfficeSuiteApp(), nil
}
// GetAutoAcceptEula gets the autoAcceptEula property value. The value to accept the EULA automatically on the enduser's device.
func (m *OfficeSuiteApp) GetAutoAcceptEula()(*bool) {
    return m.autoAcceptEula
}
// GetExcludedApps gets the excludedApps property value. The property to represent the apps which are excluded from the selected Office365 Product Id.
func (m *OfficeSuiteApp) GetExcludedApps()(ExcludedAppsable) {
    return m.excludedApps
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OfficeSuiteApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileApp.GetFieldDeserializers()
    res["autoAcceptEula"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoAcceptEula(val)
        }
        return nil
    }
    res["excludedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateExcludedAppsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExcludedApps(val.(ExcludedAppsable))
        }
        return nil
    }
    res["installProgressDisplayLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOfficeSuiteInstallProgressDisplayLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallProgressDisplayLevel(val.(*OfficeSuiteInstallProgressDisplayLevel))
        }
        return nil
    }
    res["localesToInstall"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetLocalesToInstall(res)
        }
        return nil
    }
    res["officeConfigurationXml"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfficeConfigurationXml(val)
        }
        return nil
    }
    res["officePlatformArchitecture"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsArchitecture)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfficePlatformArchitecture(val.(*WindowsArchitecture))
        }
        return nil
    }
    res["officeSuiteAppDefaultFileFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOfficeSuiteDefaultFileFormatType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfficeSuiteAppDefaultFileFormat(val.(*OfficeSuiteDefaultFileFormatType))
        }
        return nil
    }
    res["productIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseOfficeProductId)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OfficeProductId, len(val))
            for i, v := range val {
                res[i] = *(v.(*OfficeProductId))
            }
            m.SetProductIds(res)
        }
        return nil
    }
    res["shouldUninstallOlderVersionsOfOffice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShouldUninstallOlderVersionsOfOffice(val)
        }
        return nil
    }
    res["targetVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetVersion(val)
        }
        return nil
    }
    res["updateChannel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOfficeUpdateChannel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateChannel(val.(*OfficeUpdateChannel))
        }
        return nil
    }
    res["updateVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateVersion(val)
        }
        return nil
    }
    res["useSharedComputerActivation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseSharedComputerActivation(val)
        }
        return nil
    }
    return res
}
// GetInstallProgressDisplayLevel gets the installProgressDisplayLevel property value. The Enum to specify the level of display for the Installation Progress Setup UI on the Device.
func (m *OfficeSuiteApp) GetInstallProgressDisplayLevel()(*OfficeSuiteInstallProgressDisplayLevel) {
    return m.installProgressDisplayLevel
}
// GetLocalesToInstall gets the localesToInstall property value. The property to represent the locales which are installed when the apps from Office365 is installed. It uses standard RFC 6033. Ref: https://technet.microsoft.com/library/cc179219(v=office.16).aspx
func (m *OfficeSuiteApp) GetLocalesToInstall()([]string) {
    return m.localesToInstall
}
// GetOfficeConfigurationXml gets the officeConfigurationXml property value. The property to represent the XML configuration file that can be specified for Office ProPlus Apps. Takes precedence over all other properties. When present, the XML configuration file will be used to create the app.
func (m *OfficeSuiteApp) GetOfficeConfigurationXml()([]byte) {
    return m.officeConfigurationXml
}
// GetOfficePlatformArchitecture gets the officePlatformArchitecture property value. Contains properties for Windows architecture.
func (m *OfficeSuiteApp) GetOfficePlatformArchitecture()(*WindowsArchitecture) {
    return m.officePlatformArchitecture
}
// GetOfficeSuiteAppDefaultFileFormat gets the officeSuiteAppDefaultFileFormat property value. Describes the OfficeSuiteApp file format types that can be selected.
func (m *OfficeSuiteApp) GetOfficeSuiteAppDefaultFileFormat()(*OfficeSuiteDefaultFileFormatType) {
    return m.officeSuiteAppDefaultFileFormat
}
// GetProductIds gets the productIds property value. The Product Ids that represent the Office365 Suite SKU.
func (m *OfficeSuiteApp) GetProductIds()([]OfficeProductId) {
    return m.productIds
}
// GetShouldUninstallOlderVersionsOfOffice gets the shouldUninstallOlderVersionsOfOffice property value. The property to determine whether to uninstall existing Office MSI if an Office365 app suite is deployed to the device or not.
func (m *OfficeSuiteApp) GetShouldUninstallOlderVersionsOfOffice()(*bool) {
    return m.shouldUninstallOlderVersionsOfOffice
}
// GetTargetVersion gets the targetVersion property value. The property to represent the specific target version for the Office365 app suite that should be remained deployed on the devices.
func (m *OfficeSuiteApp) GetTargetVersion()(*string) {
    return m.targetVersion
}
// GetUpdateChannel gets the updateChannel property value. The Enum to specify the Office365 Updates Channel.
func (m *OfficeSuiteApp) GetUpdateChannel()(*OfficeUpdateChannel) {
    return m.updateChannel
}
// GetUpdateVersion gets the updateVersion property value. The property to represent the update version in which the specific target version is available for the Office365 app suite.
func (m *OfficeSuiteApp) GetUpdateVersion()(*string) {
    return m.updateVersion
}
// GetUseSharedComputerActivation gets the useSharedComputerActivation property value. The property to represent that whether the shared computer activation is used not for Office365 app suite.
func (m *OfficeSuiteApp) GetUseSharedComputerActivation()(*bool) {
    return m.useSharedComputerActivation
}
// Serialize serializes information the current object
func (m *OfficeSuiteApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("autoAcceptEula", m.GetAutoAcceptEula())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("excludedApps", m.GetExcludedApps())
        if err != nil {
            return err
        }
    }
    if m.GetInstallProgressDisplayLevel() != nil {
        cast := (*m.GetInstallProgressDisplayLevel()).String()
        err = writer.WriteStringValue("installProgressDisplayLevel", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetLocalesToInstall() != nil {
        err = writer.WriteCollectionOfStringValues("localesToInstall", m.GetLocalesToInstall())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("officeConfigurationXml", m.GetOfficeConfigurationXml())
        if err != nil {
            return err
        }
    }
    if m.GetOfficePlatformArchitecture() != nil {
        cast := (*m.GetOfficePlatformArchitecture()).String()
        err = writer.WriteStringValue("officePlatformArchitecture", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOfficeSuiteAppDefaultFileFormat() != nil {
        cast := (*m.GetOfficeSuiteAppDefaultFileFormat()).String()
        err = writer.WriteStringValue("officeSuiteAppDefaultFileFormat", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetProductIds() != nil {
        err = writer.WriteCollectionOfStringValues("productIds", SerializeOfficeProductId(m.GetProductIds()))
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("shouldUninstallOlderVersionsOfOffice", m.GetShouldUninstallOlderVersionsOfOffice())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetVersion", m.GetTargetVersion())
        if err != nil {
            return err
        }
    }
    if m.GetUpdateChannel() != nil {
        cast := (*m.GetUpdateChannel()).String()
        err = writer.WriteStringValue("updateChannel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("updateVersion", m.GetUpdateVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("useSharedComputerActivation", m.GetUseSharedComputerActivation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAutoAcceptEula sets the autoAcceptEula property value. The value to accept the EULA automatically on the enduser's device.
func (m *OfficeSuiteApp) SetAutoAcceptEula(value *bool)() {
    m.autoAcceptEula = value
}
// SetExcludedApps sets the excludedApps property value. The property to represent the apps which are excluded from the selected Office365 Product Id.
func (m *OfficeSuiteApp) SetExcludedApps(value ExcludedAppsable)() {
    m.excludedApps = value
}
// SetInstallProgressDisplayLevel sets the installProgressDisplayLevel property value. The Enum to specify the level of display for the Installation Progress Setup UI on the Device.
func (m *OfficeSuiteApp) SetInstallProgressDisplayLevel(value *OfficeSuiteInstallProgressDisplayLevel)() {
    m.installProgressDisplayLevel = value
}
// SetLocalesToInstall sets the localesToInstall property value. The property to represent the locales which are installed when the apps from Office365 is installed. It uses standard RFC 6033. Ref: https://technet.microsoft.com/library/cc179219(v=office.16).aspx
func (m *OfficeSuiteApp) SetLocalesToInstall(value []string)() {
    m.localesToInstall = value
}
// SetOfficeConfigurationXml sets the officeConfigurationXml property value. The property to represent the XML configuration file that can be specified for Office ProPlus Apps. Takes precedence over all other properties. When present, the XML configuration file will be used to create the app.
func (m *OfficeSuiteApp) SetOfficeConfigurationXml(value []byte)() {
    m.officeConfigurationXml = value
}
// SetOfficePlatformArchitecture sets the officePlatformArchitecture property value. Contains properties for Windows architecture.
func (m *OfficeSuiteApp) SetOfficePlatformArchitecture(value *WindowsArchitecture)() {
    m.officePlatformArchitecture = value
}
// SetOfficeSuiteAppDefaultFileFormat sets the officeSuiteAppDefaultFileFormat property value. Describes the OfficeSuiteApp file format types that can be selected.
func (m *OfficeSuiteApp) SetOfficeSuiteAppDefaultFileFormat(value *OfficeSuiteDefaultFileFormatType)() {
    m.officeSuiteAppDefaultFileFormat = value
}
// SetProductIds sets the productIds property value. The Product Ids that represent the Office365 Suite SKU.
func (m *OfficeSuiteApp) SetProductIds(value []OfficeProductId)() {
    m.productIds = value
}
// SetShouldUninstallOlderVersionsOfOffice sets the shouldUninstallOlderVersionsOfOffice property value. The property to determine whether to uninstall existing Office MSI if an Office365 app suite is deployed to the device or not.
func (m *OfficeSuiteApp) SetShouldUninstallOlderVersionsOfOffice(value *bool)() {
    m.shouldUninstallOlderVersionsOfOffice = value
}
// SetTargetVersion sets the targetVersion property value. The property to represent the specific target version for the Office365 app suite that should be remained deployed on the devices.
func (m *OfficeSuiteApp) SetTargetVersion(value *string)() {
    m.targetVersion = value
}
// SetUpdateChannel sets the updateChannel property value. The Enum to specify the Office365 Updates Channel.
func (m *OfficeSuiteApp) SetUpdateChannel(value *OfficeUpdateChannel)() {
    m.updateChannel = value
}
// SetUpdateVersion sets the updateVersion property value. The property to represent the update version in which the specific target version is available for the Office365 app suite.
func (m *OfficeSuiteApp) SetUpdateVersion(value *string)() {
    m.updateVersion = value
}
// SetUseSharedComputerActivation sets the useSharedComputerActivation property value. The property to represent that whether the shared computer activation is used not for Office365 app suite.
func (m *OfficeSuiteApp) SetUseSharedComputerActivation(value *bool)() {
    m.useSharedComputerActivation = value
}
