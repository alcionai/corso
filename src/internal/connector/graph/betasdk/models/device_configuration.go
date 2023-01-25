package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceConfiguration device Configuration.
type DeviceConfiguration struct {
    Entity
    // The list of assignments for the device configuration profile.
    assignments []DeviceConfigurationAssignmentable
    // DateTime the object was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Admin provided description of the Device Configuration.
    description *string
    // The device mode applicability rule for this Policy.
    deviceManagementApplicabilityRuleDeviceMode DeviceManagementApplicabilityRuleDeviceModeable
    // The OS edition applicability for this Policy.
    deviceManagementApplicabilityRuleOsEdition DeviceManagementApplicabilityRuleOsEditionable
    // The OS version applicability rule for this Policy.
    deviceManagementApplicabilityRuleOsVersion DeviceManagementApplicabilityRuleOsVersionable
    // Device Configuration Setting State Device Summary
    deviceSettingStateSummaries []SettingStateDeviceSummaryable
    // Device configuration installation status by device.
    deviceStatuses []DeviceConfigurationDeviceStatusable
    // Device Configuration devices status overview
    deviceStatusOverview DeviceConfigurationDeviceOverviewable
    // Admin provided name of the device configuration.
    displayName *string
    // The list of group assignments for the device configuration profile.
    groupAssignments []DeviceConfigurationGroupAssignmentable
    // DateTime the object was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // List of Scope Tags for this Entity instance.
    roleScopeTagIds []string
    // Indicates whether or not the underlying Device Configuration supports the assignment of scope tags. Assigning to the ScopeTags property is not allowed when this value is false and entities will not be visible to scoped users. This occurs for Legacy policies created in Silverlight and can be resolved by deleting and recreating the policy in the Azure Portal. This property is read-only.
    supportsScopeTags *bool
    // Device configuration installation status by user.
    userStatuses []DeviceConfigurationUserStatusable
    // Device Configuration users status overview
    userStatusOverview DeviceConfigurationUserOverviewable
    // Version of the device configuration.
    version *int32
}
// NewDeviceConfiguration instantiates a new deviceConfiguration and sets the default values.
func NewDeviceConfiguration()(*DeviceConfiguration) {
    m := &DeviceConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.androidCertificateProfileBase":
                        return NewAndroidCertificateProfileBase(), nil
                    case "#microsoft.graph.androidCustomConfiguration":
                        return NewAndroidCustomConfiguration(), nil
                    case "#microsoft.graph.androidDeviceOwnerCertificateProfileBase":
                        return NewAndroidDeviceOwnerCertificateProfileBase(), nil
                    case "#microsoft.graph.androidDeviceOwnerDerivedCredentialAuthenticationConfiguration":
                        return NewAndroidDeviceOwnerDerivedCredentialAuthenticationConfiguration(), nil
                    case "#microsoft.graph.androidDeviceOwnerEnterpriseWiFiConfiguration":
                        return NewAndroidDeviceOwnerEnterpriseWiFiConfiguration(), nil
                    case "#microsoft.graph.androidDeviceOwnerGeneralDeviceConfiguration":
                        return NewAndroidDeviceOwnerGeneralDeviceConfiguration(), nil
                    case "#microsoft.graph.androidDeviceOwnerImportedPFXCertificateProfile":
                        return NewAndroidDeviceOwnerImportedPFXCertificateProfile(), nil
                    case "#microsoft.graph.androidDeviceOwnerPkcsCertificateProfile":
                        return NewAndroidDeviceOwnerPkcsCertificateProfile(), nil
                    case "#microsoft.graph.androidDeviceOwnerScepCertificateProfile":
                        return NewAndroidDeviceOwnerScepCertificateProfile(), nil
                    case "#microsoft.graph.androidDeviceOwnerTrustedRootCertificate":
                        return NewAndroidDeviceOwnerTrustedRootCertificate(), nil
                    case "#microsoft.graph.androidDeviceOwnerVpnConfiguration":
                        return NewAndroidDeviceOwnerVpnConfiguration(), nil
                    case "#microsoft.graph.androidDeviceOwnerWiFiConfiguration":
                        return NewAndroidDeviceOwnerWiFiConfiguration(), nil
                    case "#microsoft.graph.androidEasEmailProfileConfiguration":
                        return NewAndroidEasEmailProfileConfiguration(), nil
                    case "#microsoft.graph.androidEnterpriseWiFiConfiguration":
                        return NewAndroidEnterpriseWiFiConfiguration(), nil
                    case "#microsoft.graph.androidForWorkCertificateProfileBase":
                        return NewAndroidForWorkCertificateProfileBase(), nil
                    case "#microsoft.graph.androidForWorkCustomConfiguration":
                        return NewAndroidForWorkCustomConfiguration(), nil
                    case "#microsoft.graph.androidForWorkEasEmailProfileBase":
                        return NewAndroidForWorkEasEmailProfileBase(), nil
                    case "#microsoft.graph.androidForWorkEnterpriseWiFiConfiguration":
                        return NewAndroidForWorkEnterpriseWiFiConfiguration(), nil
                    case "#microsoft.graph.androidForWorkGeneralDeviceConfiguration":
                        return NewAndroidForWorkGeneralDeviceConfiguration(), nil
                    case "#microsoft.graph.androidForWorkGmailEasConfiguration":
                        return NewAndroidForWorkGmailEasConfiguration(), nil
                    case "#microsoft.graph.androidForWorkImportedPFXCertificateProfile":
                        return NewAndroidForWorkImportedPFXCertificateProfile(), nil
                    case "#microsoft.graph.androidForWorkNineWorkEasConfiguration":
                        return NewAndroidForWorkNineWorkEasConfiguration(), nil
                    case "#microsoft.graph.androidForWorkPkcsCertificateProfile":
                        return NewAndroidForWorkPkcsCertificateProfile(), nil
                    case "#microsoft.graph.androidForWorkScepCertificateProfile":
                        return NewAndroidForWorkScepCertificateProfile(), nil
                    case "#microsoft.graph.androidForWorkTrustedRootCertificate":
                        return NewAndroidForWorkTrustedRootCertificate(), nil
                    case "#microsoft.graph.androidForWorkVpnConfiguration":
                        return NewAndroidForWorkVpnConfiguration(), nil
                    case "#microsoft.graph.androidForWorkWiFiConfiguration":
                        return NewAndroidForWorkWiFiConfiguration(), nil
                    case "#microsoft.graph.androidGeneralDeviceConfiguration":
                        return NewAndroidGeneralDeviceConfiguration(), nil
                    case "#microsoft.graph.androidImportedPFXCertificateProfile":
                        return NewAndroidImportedPFXCertificateProfile(), nil
                    case "#microsoft.graph.androidOmaCpConfiguration":
                        return NewAndroidOmaCpConfiguration(), nil
                    case "#microsoft.graph.androidPkcsCertificateProfile":
                        return NewAndroidPkcsCertificateProfile(), nil
                    case "#microsoft.graph.androidScepCertificateProfile":
                        return NewAndroidScepCertificateProfile(), nil
                    case "#microsoft.graph.androidTrustedRootCertificate":
                        return NewAndroidTrustedRootCertificate(), nil
                    case "#microsoft.graph.androidVpnConfiguration":
                        return NewAndroidVpnConfiguration(), nil
                    case "#microsoft.graph.androidWiFiConfiguration":
                        return NewAndroidWiFiConfiguration(), nil
                    case "#microsoft.graph.androidWorkProfileCertificateProfileBase":
                        return NewAndroidWorkProfileCertificateProfileBase(), nil
                    case "#microsoft.graph.androidWorkProfileCustomConfiguration":
                        return NewAndroidWorkProfileCustomConfiguration(), nil
                    case "#microsoft.graph.androidWorkProfileEasEmailProfileBase":
                        return NewAndroidWorkProfileEasEmailProfileBase(), nil
                    case "#microsoft.graph.androidWorkProfileEnterpriseWiFiConfiguration":
                        return NewAndroidWorkProfileEnterpriseWiFiConfiguration(), nil
                    case "#microsoft.graph.androidWorkProfileGeneralDeviceConfiguration":
                        return NewAndroidWorkProfileGeneralDeviceConfiguration(), nil
                    case "#microsoft.graph.androidWorkProfileGmailEasConfiguration":
                        return NewAndroidWorkProfileGmailEasConfiguration(), nil
                    case "#microsoft.graph.androidWorkProfileNineWorkEasConfiguration":
                        return NewAndroidWorkProfileNineWorkEasConfiguration(), nil
                    case "#microsoft.graph.androidWorkProfilePkcsCertificateProfile":
                        return NewAndroidWorkProfilePkcsCertificateProfile(), nil
                    case "#microsoft.graph.androidWorkProfileScepCertificateProfile":
                        return NewAndroidWorkProfileScepCertificateProfile(), nil
                    case "#microsoft.graph.androidWorkProfileTrustedRootCertificate":
                        return NewAndroidWorkProfileTrustedRootCertificate(), nil
                    case "#microsoft.graph.androidWorkProfileVpnConfiguration":
                        return NewAndroidWorkProfileVpnConfiguration(), nil
                    case "#microsoft.graph.androidWorkProfileWiFiConfiguration":
                        return NewAndroidWorkProfileWiFiConfiguration(), nil
                    case "#microsoft.graph.aospDeviceOwnerCertificateProfileBase":
                        return NewAospDeviceOwnerCertificateProfileBase(), nil
                    case "#microsoft.graph.aospDeviceOwnerDeviceConfiguration":
                        return NewAospDeviceOwnerDeviceConfiguration(), nil
                    case "#microsoft.graph.aospDeviceOwnerEnterpriseWiFiConfiguration":
                        return NewAospDeviceOwnerEnterpriseWiFiConfiguration(), nil
                    case "#microsoft.graph.aospDeviceOwnerPkcsCertificateProfile":
                        return NewAospDeviceOwnerPkcsCertificateProfile(), nil
                    case "#microsoft.graph.aospDeviceOwnerScepCertificateProfile":
                        return NewAospDeviceOwnerScepCertificateProfile(), nil
                    case "#microsoft.graph.aospDeviceOwnerTrustedRootCertificate":
                        return NewAospDeviceOwnerTrustedRootCertificate(), nil
                    case "#microsoft.graph.aospDeviceOwnerWiFiConfiguration":
                        return NewAospDeviceOwnerWiFiConfiguration(), nil
                    case "#microsoft.graph.appleDeviceFeaturesConfigurationBase":
                        return NewAppleDeviceFeaturesConfigurationBase(), nil
                    case "#microsoft.graph.appleExpeditedCheckinConfigurationBase":
                        return NewAppleExpeditedCheckinConfigurationBase(), nil
                    case "#microsoft.graph.appleVpnConfiguration":
                        return NewAppleVpnConfiguration(), nil
                    case "#microsoft.graph.easEmailProfileConfigurationBase":
                        return NewEasEmailProfileConfigurationBase(), nil
                    case "#microsoft.graph.editionUpgradeConfiguration":
                        return NewEditionUpgradeConfiguration(), nil
                    case "#microsoft.graph.iosCertificateProfile":
                        return NewIosCertificateProfile(), nil
                    case "#microsoft.graph.iosCertificateProfileBase":
                        return NewIosCertificateProfileBase(), nil
                    case "#microsoft.graph.iosCustomConfiguration":
                        return NewIosCustomConfiguration(), nil
                    case "#microsoft.graph.iosDerivedCredentialAuthenticationConfiguration":
                        return NewIosDerivedCredentialAuthenticationConfiguration(), nil
                    case "#microsoft.graph.iosDeviceFeaturesConfiguration":
                        return NewIosDeviceFeaturesConfiguration(), nil
                    case "#microsoft.graph.iosEasEmailProfileConfiguration":
                        return NewIosEasEmailProfileConfiguration(), nil
                    case "#microsoft.graph.iosEducationDeviceConfiguration":
                        return NewIosEducationDeviceConfiguration(), nil
                    case "#microsoft.graph.iosEduDeviceConfiguration":
                        return NewIosEduDeviceConfiguration(), nil
                    case "#microsoft.graph.iosEnterpriseWiFiConfiguration":
                        return NewIosEnterpriseWiFiConfiguration(), nil
                    case "#microsoft.graph.iosExpeditedCheckinConfiguration":
                        return NewIosExpeditedCheckinConfiguration(), nil
                    case "#microsoft.graph.iosGeneralDeviceConfiguration":
                        return NewIosGeneralDeviceConfiguration(), nil
                    case "#microsoft.graph.iosikEv2VpnConfiguration":
                        return NewIosikEv2VpnConfiguration(), nil
                    case "#microsoft.graph.iosImportedPFXCertificateProfile":
                        return NewIosImportedPFXCertificateProfile(), nil
                    case "#microsoft.graph.iosPkcsCertificateProfile":
                        return NewIosPkcsCertificateProfile(), nil
                    case "#microsoft.graph.iosScepCertificateProfile":
                        return NewIosScepCertificateProfile(), nil
                    case "#microsoft.graph.iosTrustedRootCertificate":
                        return NewIosTrustedRootCertificate(), nil
                    case "#microsoft.graph.iosUpdateConfiguration":
                        return NewIosUpdateConfiguration(), nil
                    case "#microsoft.graph.iosVpnConfiguration":
                        return NewIosVpnConfiguration(), nil
                    case "#microsoft.graph.iosWiFiConfiguration":
                        return NewIosWiFiConfiguration(), nil
                    case "#microsoft.graph.macOSCertificateProfileBase":
                        return NewMacOSCertificateProfileBase(), nil
                    case "#microsoft.graph.macOSCustomAppConfiguration":
                        return NewMacOSCustomAppConfiguration(), nil
                    case "#microsoft.graph.macOSCustomConfiguration":
                        return NewMacOSCustomConfiguration(), nil
                    case "#microsoft.graph.macOSDeviceFeaturesConfiguration":
                        return NewMacOSDeviceFeaturesConfiguration(), nil
                    case "#microsoft.graph.macOSEndpointProtectionConfiguration":
                        return NewMacOSEndpointProtectionConfiguration(), nil
                    case "#microsoft.graph.macOSEnterpriseWiFiConfiguration":
                        return NewMacOSEnterpriseWiFiConfiguration(), nil
                    case "#microsoft.graph.macOSExtensionsConfiguration":
                        return NewMacOSExtensionsConfiguration(), nil
                    case "#microsoft.graph.macOSGeneralDeviceConfiguration":
                        return NewMacOSGeneralDeviceConfiguration(), nil
                    case "#microsoft.graph.macOSImportedPFXCertificateProfile":
                        return NewMacOSImportedPFXCertificateProfile(), nil
                    case "#microsoft.graph.macOSPkcsCertificateProfile":
                        return NewMacOSPkcsCertificateProfile(), nil
                    case "#microsoft.graph.macOSScepCertificateProfile":
                        return NewMacOSScepCertificateProfile(), nil
                    case "#microsoft.graph.macOSSoftwareUpdateConfiguration":
                        return NewMacOSSoftwareUpdateConfiguration(), nil
                    case "#microsoft.graph.macOSTrustedRootCertificate":
                        return NewMacOSTrustedRootCertificate(), nil
                    case "#microsoft.graph.macOSVpnConfiguration":
                        return NewMacOSVpnConfiguration(), nil
                    case "#microsoft.graph.macOSWiFiConfiguration":
                        return NewMacOSWiFiConfiguration(), nil
                    case "#microsoft.graph.macOSWiredNetworkConfiguration":
                        return NewMacOSWiredNetworkConfiguration(), nil
                    case "#microsoft.graph.sharedPCConfiguration":
                        return NewSharedPCConfiguration(), nil
                    case "#microsoft.graph.unsupportedDeviceConfiguration":
                        return NewUnsupportedDeviceConfiguration(), nil
                    case "#microsoft.graph.vpnConfiguration":
                        return NewVpnConfiguration(), nil
                    case "#microsoft.graph.windows10CertificateProfileBase":
                        return NewWindows10CertificateProfileBase(), nil
                    case "#microsoft.graph.windows10CustomConfiguration":
                        return NewWindows10CustomConfiguration(), nil
                    case "#microsoft.graph.windows10DeviceFirmwareConfigurationInterface":
                        return NewWindows10DeviceFirmwareConfigurationInterface(), nil
                    case "#microsoft.graph.windows10EasEmailProfileConfiguration":
                        return NewWindows10EasEmailProfileConfiguration(), nil
                    case "#microsoft.graph.windows10EndpointProtectionConfiguration":
                        return NewWindows10EndpointProtectionConfiguration(), nil
                    case "#microsoft.graph.windows10EnterpriseModernAppManagementConfiguration":
                        return NewWindows10EnterpriseModernAppManagementConfiguration(), nil
                    case "#microsoft.graph.windows10GeneralConfiguration":
                        return NewWindows10GeneralConfiguration(), nil
                    case "#microsoft.graph.windows10ImportedPFXCertificateProfile":
                        return NewWindows10ImportedPFXCertificateProfile(), nil
                    case "#microsoft.graph.windows10NetworkBoundaryConfiguration":
                        return NewWindows10NetworkBoundaryConfiguration(), nil
                    case "#microsoft.graph.windows10PFXImportCertificateProfile":
                        return NewWindows10PFXImportCertificateProfile(), nil
                    case "#microsoft.graph.windows10PkcsCertificateProfile":
                        return NewWindows10PkcsCertificateProfile(), nil
                    case "#microsoft.graph.windows10SecureAssessmentConfiguration":
                        return NewWindows10SecureAssessmentConfiguration(), nil
                    case "#microsoft.graph.windows10TeamGeneralConfiguration":
                        return NewWindows10TeamGeneralConfiguration(), nil
                    case "#microsoft.graph.windows10VpnConfiguration":
                        return NewWindows10VpnConfiguration(), nil
                    case "#microsoft.graph.windows81CertificateProfileBase":
                        return NewWindows81CertificateProfileBase(), nil
                    case "#microsoft.graph.windows81GeneralConfiguration":
                        return NewWindows81GeneralConfiguration(), nil
                    case "#microsoft.graph.windows81SCEPCertificateProfile":
                        return NewWindows81SCEPCertificateProfile(), nil
                    case "#microsoft.graph.windows81TrustedRootCertificate":
                        return NewWindows81TrustedRootCertificate(), nil
                    case "#microsoft.graph.windows81VpnConfiguration":
                        return NewWindows81VpnConfiguration(), nil
                    case "#microsoft.graph.windows81WifiImportConfiguration":
                        return NewWindows81WifiImportConfiguration(), nil
                    case "#microsoft.graph.windowsCertificateProfileBase":
                        return NewWindowsCertificateProfileBase(), nil
                    case "#microsoft.graph.windowsDefenderAdvancedThreatProtectionConfiguration":
                        return NewWindowsDefenderAdvancedThreatProtectionConfiguration(), nil
                    case "#microsoft.graph.windowsDeliveryOptimizationConfiguration":
                        return NewWindowsDeliveryOptimizationConfiguration(), nil
                    case "#microsoft.graph.windowsDomainJoinConfiguration":
                        return NewWindowsDomainJoinConfiguration(), nil
                    case "#microsoft.graph.windowsHealthMonitoringConfiguration":
                        return NewWindowsHealthMonitoringConfiguration(), nil
                    case "#microsoft.graph.windowsIdentityProtectionConfiguration":
                        return NewWindowsIdentityProtectionConfiguration(), nil
                    case "#microsoft.graph.windowsKioskConfiguration":
                        return NewWindowsKioskConfiguration(), nil
                    case "#microsoft.graph.windowsPhone81CertificateProfileBase":
                        return NewWindowsPhone81CertificateProfileBase(), nil
                    case "#microsoft.graph.windowsPhone81CustomConfiguration":
                        return NewWindowsPhone81CustomConfiguration(), nil
                    case "#microsoft.graph.windowsPhone81GeneralConfiguration":
                        return NewWindowsPhone81GeneralConfiguration(), nil
                    case "#microsoft.graph.windowsPhone81ImportedPFXCertificateProfile":
                        return NewWindowsPhone81ImportedPFXCertificateProfile(), nil
                    case "#microsoft.graph.windowsPhone81SCEPCertificateProfile":
                        return NewWindowsPhone81SCEPCertificateProfile(), nil
                    case "#microsoft.graph.windowsPhone81TrustedRootCertificate":
                        return NewWindowsPhone81TrustedRootCertificate(), nil
                    case "#microsoft.graph.windowsPhone81VpnConfiguration":
                        return NewWindowsPhone81VpnConfiguration(), nil
                    case "#microsoft.graph.windowsPhoneEASEmailProfileConfiguration":
                        return NewWindowsPhoneEASEmailProfileConfiguration(), nil
                    case "#microsoft.graph.windowsUpdateForBusinessConfiguration":
                        return NewWindowsUpdateForBusinessConfiguration(), nil
                    case "#microsoft.graph.windowsVpnConfiguration":
                        return NewWindowsVpnConfiguration(), nil
                    case "#microsoft.graph.windowsWifiConfiguration":
                        return NewWindowsWifiConfiguration(), nil
                    case "#microsoft.graph.windowsWifiEnterpriseEAPConfiguration":
                        return NewWindowsWifiEnterpriseEAPConfiguration(), nil
                    case "#microsoft.graph.windowsWiredNetworkConfiguration":
                        return NewWindowsWiredNetworkConfiguration(), nil
                }
            }
        }
    }
    return NewDeviceConfiguration(), nil
}
// GetAssignments gets the assignments property value. The list of assignments for the device configuration profile.
func (m *DeviceConfiguration) GetAssignments()([]DeviceConfigurationAssignmentable) {
    return m.assignments
}
// GetCreatedDateTime gets the createdDateTime property value. DateTime the object was created.
func (m *DeviceConfiguration) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. Admin provided description of the Device Configuration.
func (m *DeviceConfiguration) GetDescription()(*string) {
    return m.description
}
// GetDeviceManagementApplicabilityRuleDeviceMode gets the deviceManagementApplicabilityRuleDeviceMode property value. The device mode applicability rule for this Policy.
func (m *DeviceConfiguration) GetDeviceManagementApplicabilityRuleDeviceMode()(DeviceManagementApplicabilityRuleDeviceModeable) {
    return m.deviceManagementApplicabilityRuleDeviceMode
}
// GetDeviceManagementApplicabilityRuleOsEdition gets the deviceManagementApplicabilityRuleOsEdition property value. The OS edition applicability for this Policy.
func (m *DeviceConfiguration) GetDeviceManagementApplicabilityRuleOsEdition()(DeviceManagementApplicabilityRuleOsEditionable) {
    return m.deviceManagementApplicabilityRuleOsEdition
}
// GetDeviceManagementApplicabilityRuleOsVersion gets the deviceManagementApplicabilityRuleOsVersion property value. The OS version applicability rule for this Policy.
func (m *DeviceConfiguration) GetDeviceManagementApplicabilityRuleOsVersion()(DeviceManagementApplicabilityRuleOsVersionable) {
    return m.deviceManagementApplicabilityRuleOsVersion
}
// GetDeviceSettingStateSummaries gets the deviceSettingStateSummaries property value. Device Configuration Setting State Device Summary
func (m *DeviceConfiguration) GetDeviceSettingStateSummaries()([]SettingStateDeviceSummaryable) {
    return m.deviceSettingStateSummaries
}
// GetDeviceStatuses gets the deviceStatuses property value. Device configuration installation status by device.
func (m *DeviceConfiguration) GetDeviceStatuses()([]DeviceConfigurationDeviceStatusable) {
    return m.deviceStatuses
}
// GetDeviceStatusOverview gets the deviceStatusOverview property value. Device Configuration devices status overview
func (m *DeviceConfiguration) GetDeviceStatusOverview()(DeviceConfigurationDeviceOverviewable) {
    return m.deviceStatusOverview
}
// GetDisplayName gets the displayName property value. Admin provided name of the device configuration.
func (m *DeviceConfiguration) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceConfigurationAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceConfigurationAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["deviceManagementApplicabilityRuleDeviceMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementApplicabilityRuleDeviceModeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceManagementApplicabilityRuleDeviceMode(val.(DeviceManagementApplicabilityRuleDeviceModeable))
        }
        return nil
    }
    res["deviceManagementApplicabilityRuleOsEdition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementApplicabilityRuleOsEditionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceManagementApplicabilityRuleOsEdition(val.(DeviceManagementApplicabilityRuleOsEditionable))
        }
        return nil
    }
    res["deviceManagementApplicabilityRuleOsVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementApplicabilityRuleOsVersionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceManagementApplicabilityRuleOsVersion(val.(DeviceManagementApplicabilityRuleOsVersionable))
        }
        return nil
    }
    res["deviceSettingStateSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSettingStateDeviceSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SettingStateDeviceSummaryable, len(val))
            for i, v := range val {
                res[i] = v.(SettingStateDeviceSummaryable)
            }
            m.SetDeviceSettingStateSummaries(res)
        }
        return nil
    }
    res["deviceStatuses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceConfigurationDeviceStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationDeviceStatusable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceConfigurationDeviceStatusable)
            }
            m.SetDeviceStatuses(res)
        }
        return nil
    }
    res["deviceStatusOverview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceConfigurationDeviceOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceStatusOverview(val.(DeviceConfigurationDeviceOverviewable))
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["groupAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceConfigurationGroupAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationGroupAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceConfigurationGroupAssignmentable)
            }
            m.SetGroupAssignments(res)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["supportsScopeTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportsScopeTags(val)
        }
        return nil
    }
    res["userStatuses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceConfigurationUserStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationUserStatusable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceConfigurationUserStatusable)
            }
            m.SetUserStatuses(res)
        }
        return nil
    }
    res["userStatusOverview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceConfigurationUserOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserStatusOverview(val.(DeviceConfigurationUserOverviewable))
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetGroupAssignments gets the groupAssignments property value. The list of group assignments for the device configuration profile.
func (m *DeviceConfiguration) GetGroupAssignments()([]DeviceConfigurationGroupAssignmentable) {
    return m.groupAssignments
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *DeviceConfiguration) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DeviceConfiguration) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// GetSupportsScopeTags gets the supportsScopeTags property value. Indicates whether or not the underlying Device Configuration supports the assignment of scope tags. Assigning to the ScopeTags property is not allowed when this value is false and entities will not be visible to scoped users. This occurs for Legacy policies created in Silverlight and can be resolved by deleting and recreating the policy in the Azure Portal. This property is read-only.
func (m *DeviceConfiguration) GetSupportsScopeTags()(*bool) {
    return m.supportsScopeTags
}
// GetUserStatuses gets the userStatuses property value. Device configuration installation status by user.
func (m *DeviceConfiguration) GetUserStatuses()([]DeviceConfigurationUserStatusable) {
    return m.userStatuses
}
// GetUserStatusOverview gets the userStatusOverview property value. Device Configuration users status overview
func (m *DeviceConfiguration) GetUserStatusOverview()(DeviceConfigurationUserOverviewable) {
    return m.userStatusOverview
}
// GetVersion gets the version property value. Version of the device configuration.
func (m *DeviceConfiguration) GetVersion()(*int32) {
    return m.version
}
// Serialize serializes information the current object
func (m *DeviceConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceManagementApplicabilityRuleDeviceMode", m.GetDeviceManagementApplicabilityRuleDeviceMode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceManagementApplicabilityRuleOsEdition", m.GetDeviceManagementApplicabilityRuleOsEdition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceManagementApplicabilityRuleOsVersion", m.GetDeviceManagementApplicabilityRuleOsVersion())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceSettingStateSummaries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceSettingStateSummaries()))
        for i, v := range m.GetDeviceSettingStateSummaries() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceSettingStateSummaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceStatuses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceStatuses()))
        for i, v := range m.GetDeviceStatuses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceStatusOverview", m.GetDeviceStatusOverview())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetGroupAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroupAssignments()))
        for i, v := range m.GetGroupAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("groupAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    if m.GetUserStatuses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserStatuses()))
        for i, v := range m.GetUserStatuses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userStatusOverview", m.GetUserStatusOverview())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The list of assignments for the device configuration profile.
func (m *DeviceConfiguration) SetAssignments(value []DeviceConfigurationAssignmentable)() {
    m.assignments = value
}
// SetCreatedDateTime sets the createdDateTime property value. DateTime the object was created.
func (m *DeviceConfiguration) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Admin provided description of the Device Configuration.
func (m *DeviceConfiguration) SetDescription(value *string)() {
    m.description = value
}
// SetDeviceManagementApplicabilityRuleDeviceMode sets the deviceManagementApplicabilityRuleDeviceMode property value. The device mode applicability rule for this Policy.
func (m *DeviceConfiguration) SetDeviceManagementApplicabilityRuleDeviceMode(value DeviceManagementApplicabilityRuleDeviceModeable)() {
    m.deviceManagementApplicabilityRuleDeviceMode = value
}
// SetDeviceManagementApplicabilityRuleOsEdition sets the deviceManagementApplicabilityRuleOsEdition property value. The OS edition applicability for this Policy.
func (m *DeviceConfiguration) SetDeviceManagementApplicabilityRuleOsEdition(value DeviceManagementApplicabilityRuleOsEditionable)() {
    m.deviceManagementApplicabilityRuleOsEdition = value
}
// SetDeviceManagementApplicabilityRuleOsVersion sets the deviceManagementApplicabilityRuleOsVersion property value. The OS version applicability rule for this Policy.
func (m *DeviceConfiguration) SetDeviceManagementApplicabilityRuleOsVersion(value DeviceManagementApplicabilityRuleOsVersionable)() {
    m.deviceManagementApplicabilityRuleOsVersion = value
}
// SetDeviceSettingStateSummaries sets the deviceSettingStateSummaries property value. Device Configuration Setting State Device Summary
func (m *DeviceConfiguration) SetDeviceSettingStateSummaries(value []SettingStateDeviceSummaryable)() {
    m.deviceSettingStateSummaries = value
}
// SetDeviceStatuses sets the deviceStatuses property value. Device configuration installation status by device.
func (m *DeviceConfiguration) SetDeviceStatuses(value []DeviceConfigurationDeviceStatusable)() {
    m.deviceStatuses = value
}
// SetDeviceStatusOverview sets the deviceStatusOverview property value. Device Configuration devices status overview
func (m *DeviceConfiguration) SetDeviceStatusOverview(value DeviceConfigurationDeviceOverviewable)() {
    m.deviceStatusOverview = value
}
// SetDisplayName sets the displayName property value. Admin provided name of the device configuration.
func (m *DeviceConfiguration) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGroupAssignments sets the groupAssignments property value. The list of group assignments for the device configuration profile.
func (m *DeviceConfiguration) SetGroupAssignments(value []DeviceConfigurationGroupAssignmentable)() {
    m.groupAssignments = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *DeviceConfiguration) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DeviceConfiguration) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// SetSupportsScopeTags sets the supportsScopeTags property value. Indicates whether or not the underlying Device Configuration supports the assignment of scope tags. Assigning to the ScopeTags property is not allowed when this value is false and entities will not be visible to scoped users. This occurs for Legacy policies created in Silverlight and can be resolved by deleting and recreating the policy in the Azure Portal. This property is read-only.
func (m *DeviceConfiguration) SetSupportsScopeTags(value *bool)() {
    m.supportsScopeTags = value
}
// SetUserStatuses sets the userStatuses property value. Device configuration installation status by user.
func (m *DeviceConfiguration) SetUserStatuses(value []DeviceConfigurationUserStatusable)() {
    m.userStatuses = value
}
// SetUserStatusOverview sets the userStatusOverview property value. Device Configuration users status overview
func (m *DeviceConfiguration) SetUserStatusOverview(value DeviceConfigurationUserOverviewable)() {
    m.userStatusOverview = value
}
// SetVersion sets the version property value. Version of the device configuration.
func (m *DeviceConfiguration) SetVersion(value *int32)() {
    m.version = value
}
