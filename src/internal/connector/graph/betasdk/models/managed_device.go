package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDevice devices that are managed or pre-enrolled through Intune
type ManagedDevice struct {
    Entity
    // Whether the device is Azure Active Directory registered. This property is read-only.
    aadRegistered *bool
    // The code that allows the Activation Lock on managed device to be bypassed. Default, is Null (Non-Default property) for this property when returned as part of managedDevice entity in LIST call. Individual GET call with select query options is needed to retrieve actual values. Supports: $select. $Search is not supported. Read-only. This property is read-only.
    activationLockBypassCode *string
    // Android security patch level. This property is read-only.
    androidSecurityPatchLevel *string
    // Managed device mobile app configuration states for this device.
    assignmentFilterEvaluationStatusDetails []AssignmentFilterEvaluationStatusDetailsable
    // Reports if the managed device is enrolled via auto-pilot. This property is read-only.
    autopilotEnrolled *bool
    // The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
    azureActiveDirectoryDeviceId *string
    // The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
    azureADDeviceId *string
    // Whether the device is Azure Active Directory registered. This property is read-only.
    azureADRegistered *bool
    // Reports if the managed device has an escrowed Bootstrap Token. This is only for macOS devices. To get, include BootstrapTokenEscrowed in the select clause and query with a device id. If FALSE, no bootstrap token is escrowed. If TRUE, the device has escrowed a bootstrap token with Intune. This property is read-only.
    bootstrapTokenEscrowed *bool
    // Chassis type.
    chassisType *ChassisType
    // List of properties of the ChromeOS Device.
    chromeOSDeviceInfo []ChromeOSDevicePropertyable
    // The cloudPcRemoteActionResults property
    cloudPcRemoteActionResults []CloudPcRemoteActionResultable
    // The DateTime when device compliance grace period expires. This property is read-only.
    complianceGracePeriodExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Compliance state.
    complianceState *ComplianceState
    // ConfigrMgr client enabled features. This property is read-only.
    configurationManagerClientEnabledFeatures ConfigurationManagerClientEnabledFeaturesable
    // Configuration manager client health state, valid only for devices managed by MDM/ConfigMgr Agent
    configurationManagerClientHealthState ConfigurationManagerClientHealthStateable
    // Configuration manager client information, valid only for devices managed, duel-managed or tri-managed by ConfigMgr Agent
    configurationManagerClientInformation ConfigurationManagerClientInformationable
    // All applications currently installed on the device
    detectedApps []DetectedAppable
    // List of ComplexType deviceActionResult objects. This property is read-only.
    deviceActionResults []DeviceActionResultable
    // Device category
    deviceCategory DeviceCategoryable
    // Device category display name. This property is read-only.
    deviceCategoryDisplayName *string
    // Device compliance policy states for this device.
    deviceCompliancePolicyStates []DeviceCompliancePolicyStateable
    // Device configuration states for this device.
    deviceConfigurationStates []DeviceConfigurationStateable
    // Possible ways of adding a mobile device to management.
    deviceEnrollmentType *DeviceEnrollmentType
    // Indicates whether the device is DFCI managed. When TRUE the device is DFCI managed. When FALSE, the device is not DFCI managed. The default value is FALSE.
    deviceFirmwareConfigurationInterfaceManaged *bool
    // The device health attestation state. This property is read-only.
    deviceHealthAttestationState DeviceHealthAttestationStateable
    // Name of the device. This property is read-only.
    deviceName *string
    // Device registration status.
    deviceRegistrationState *DeviceRegistrationState
    // Device type.
    deviceType *DeviceType
    // Whether the device is Exchange ActiveSync activated. This property is read-only.
    easActivated *bool
    // Exchange ActivationSync activation time of the device. This property is read-only.
    easActivationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Exchange ActiveSync Id of the device. This property is read-only.
    easDeviceId *string
    // Email(s) for the user associated with the device. This property is read-only.
    emailAddress *string
    // Enrollment time of the device. This property is read-only.
    enrolledDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Name of the enrollment profile assigned to the device. Default value is empty string, indicating no enrollment profile was assgined. This property is read-only.
    enrollmentProfileName *string
    // Ethernet MAC. Default, is Null (Non-Default property) for this property when returned as part of managedDevice entity. Individual get call with select query options is needed to retrieve actual values. Example: deviceManagement/managedDevices({managedDeviceId})?$select=ethernetMacAddress Supports: $select. $Search is not supported. Read-only. This property is read-only.
    ethernetMacAddress *string
    // Device Exchange Access State.
    exchangeAccessState *DeviceManagementExchangeAccessState
    // Device Exchange Access State Reason.
    exchangeAccessStateReason *DeviceManagementExchangeAccessStateReason
    // Last time the device contacted Exchange. This property is read-only.
    exchangeLastSuccessfulSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Free Storage in Bytes. Default value is 0. Read-only. This property is read-only.
    freeStorageSpaceInBytes *int64
    // The hardward details for the device.  Includes information such as storage space, manufacturer, serial number, etc. Return default value in LIST managedDevices. Real value only returned in singel device GET call with device id and included in select parameter. Supports: $select. $Search is not supported. Read-only. This property is read-only.
    hardwareInformation HardwareInformationable
    // Integrated Circuit Card Identifier, it is A SIM card's unique identification number. Return default value null in LIST managedDevices. Real value only returned in singel device GET call with device id and included in select parameter. Supports: $select. $Search is not supported. Read-only. This property is read-only.
    iccid *string
    // IMEI. This property is read-only.
    imei *string
    // Device encryption status. This property is read-only.
    isEncrypted *bool
    // Device supervised status. This property is read-only.
    isSupervised *bool
    // whether the device is jail broken or rooted. This property is read-only.
    jailBroken *string
    // Device enrollment join type.
    joinType *JoinType
    // The date and time that the device last completed a successful sync with Intune. This property is read-only.
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // List of log collection requests
    logCollectionRequests []DeviceLogCollectionResponseable
    // State of lost mode, indicating if lost mode is enabled or disabled
    lostModeState *LostModeState
    // Managed device mobile app configuration states for this device.
    managedDeviceMobileAppConfigurationStates []ManagedDeviceMobileAppConfigurationStateable
    // Automatically generated name to identify a device. Can be overwritten to a user friendly name.
    managedDeviceName *string
    // Owner type of device.
    managedDeviceOwnerType *ManagedDeviceOwnerType
    // Management agent type.
    managementAgent *ManagementAgentType
    // Reports device management certificate expiration date. This property is read-only.
    managementCertificateExpirationDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Device management features.
    managementFeatures *ManagedDeviceManagementFeatures
    // Management state of device in Microsoft Intune.
    managementState *ManagementState
    // Manufacturer of the device. This property is read-only.
    manufacturer *string
    // MEID. This property is read-only.
    meid *string
    // Model of the device. This property is read-only.
    model *string
    // Notes on the device created by IT Admin. Return default value null in LIST managedDevices. Real value only returned in singel device GET call with device id and included in select parameter. Supports: $select.  $Search is not supported.
    notes *string
    // Operating system of the device. Windows, iOS, etc. This property is read-only.
    operatingSystem *string
    // Operating system version of the device. This property is read-only.
    osVersion *string
    // Owner type of device.
    ownerType *OwnerType
    // Available health states for the Device Health API
    partnerReportedThreatState *ManagedDevicePartnerReportedHealthState
    // Phone number of the device. This property is read-only.
    phoneNumber *string
    // Total Memory in Bytes. Return default value 0 in LIST managedDevices. Real value only returned in singel device GET call with device id and included in select parameter. Supports: $select. Default value is 0. Read-only. This property is read-only.
    physicalMemoryInBytes *int64
    // Reports the DateTime the preferMdmOverGroupPolicy setting was set.  When set, the Intune MDM settings will override Group Policy settings if there is a conflict. Read Only. This property is read-only.
    preferMdmOverGroupPolicyAppliedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Processor architecture
    processorArchitecture *ManagedDeviceArchitecture
    // An error string that identifies issues when creating Remote Assistance session objects. This property is read-only.
    remoteAssistanceSessionErrorDetails *string
    // Url that allows a Remote Assistance session to be established with the device. This property is read-only.
    remoteAssistanceSessionUrl *string
    // Reports if the managed iOS device is user approval enrollment. This property is read-only.
    requireUserEnrollmentApproval *bool
    // Indicates the time after when a device will be auto retired because of scheduled action. This property is read-only.
    retireAfterDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // List of Scope Tag IDs for this Device instance.
    roleScopeTagIds []string
    // Security baseline states for this device.
    securityBaselineStates []SecurityBaselineStateable
    // SerialNumber. This property is read-only.
    serialNumber *string
    // Device sku family
    skuFamily *string
    // Device sku number, see also: https://learn.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-getproductinfo. Valid values 0 to 2147483647. This property is read-only.
    skuNumber *int32
    // Specification version. This property is read-only.
    specificationVersion *string
    // Subscriber Carrier. This property is read-only.
    subscriberCarrier *string
    // Total Storage in Bytes. This property is read-only.
    totalStorageSpaceInBytes *int64
    // Unique Device Identifier for iOS and macOS devices. Return default value null in LIST managedDevices. Real value only returned in singel device GET call with device id and included in select parameter. Supports: $select. $Search is not supported. Read-only. This property is read-only.
    udid *string
    // User display name. This property is read-only.
    userDisplayName *string
    // Unique Identifier for the user associated with the device. This property is read-only.
    userId *string
    // Device user principal name. This property is read-only.
    userPrincipalName *string
    // The primary users associated with the managed device.
    users []Userable
    // Indicates the last logged on users of a device. This property is read-only.
    usersLoggedOn []LoggedOnUserable
    // Wi-Fi MAC. This property is read-only.
    wiFiMacAddress *string
    // Count of active malware for this windows device. This property is read-only.
    windowsActiveMalwareCount *int32
    // The device protection status. This property is read-only.
    windowsProtectionState WindowsProtectionStateable
    // Count of remediated malware for this windows device. This property is read-only.
    windowsRemediatedMalwareCount *int32
}
// NewManagedDevice instantiates a new managedDevice and sets the default values.
func NewManagedDevice()(*ManagedDevice) {
    m := &ManagedDevice{
        Entity: *NewEntity(),
    }
    return m
}
// CreateManagedDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.windowsManagedDevice":
                        return NewWindowsManagedDevice(), nil
                }
            }
        }
    }
    return NewManagedDevice(), nil
}
// GetAadRegistered gets the aadRegistered property value. Whether the device is Azure Active Directory registered. This property is read-only.
func (m *ManagedDevice) GetAadRegistered()(*bool) {
    return m.aadRegistered
}
// GetActivationLockBypassCode gets the activationLockBypassCode property value. The code that allows the Activation Lock on managed device to be bypassed. Default, is Null (Non-Default property) for this property when returned as part of managedDevice entity in LIST call. Individual GET call with select query options is needed to retrieve actual values. Supports: $select. $Search is not supported. Read-only. This property is read-only.
func (m *ManagedDevice) GetActivationLockBypassCode()(*string) {
    return m.activationLockBypassCode
}
// GetAndroidSecurityPatchLevel gets the androidSecurityPatchLevel property value. Android security patch level. This property is read-only.
func (m *ManagedDevice) GetAndroidSecurityPatchLevel()(*string) {
    return m.androidSecurityPatchLevel
}
// GetAssignmentFilterEvaluationStatusDetails gets the assignmentFilterEvaluationStatusDetails property value. Managed device mobile app configuration states for this device.
func (m *ManagedDevice) GetAssignmentFilterEvaluationStatusDetails()([]AssignmentFilterEvaluationStatusDetailsable) {
    return m.assignmentFilterEvaluationStatusDetails
}
// GetAutopilotEnrolled gets the autopilotEnrolled property value. Reports if the managed device is enrolled via auto-pilot. This property is read-only.
func (m *ManagedDevice) GetAutopilotEnrolled()(*bool) {
    return m.autopilotEnrolled
}
// GetAzureActiveDirectoryDeviceId gets the azureActiveDirectoryDeviceId property value. The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
func (m *ManagedDevice) GetAzureActiveDirectoryDeviceId()(*string) {
    return m.azureActiveDirectoryDeviceId
}
// GetAzureADDeviceId gets the azureADDeviceId property value. The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
func (m *ManagedDevice) GetAzureADDeviceId()(*string) {
    return m.azureADDeviceId
}
// GetAzureADRegistered gets the azureADRegistered property value. Whether the device is Azure Active Directory registered. This property is read-only.
func (m *ManagedDevice) GetAzureADRegistered()(*bool) {
    return m.azureADRegistered
}
// GetBootstrapTokenEscrowed gets the bootstrapTokenEscrowed property value. Reports if the managed device has an escrowed Bootstrap Token. This is only for macOS devices. To get, include BootstrapTokenEscrowed in the select clause and query with a device id. If FALSE, no bootstrap token is escrowed. If TRUE, the device has escrowed a bootstrap token with Intune. This property is read-only.
func (m *ManagedDevice) GetBootstrapTokenEscrowed()(*bool) {
    return m.bootstrapTokenEscrowed
}
// GetChassisType gets the chassisType property value. Chassis type.
func (m *ManagedDevice) GetChassisType()(*ChassisType) {
    return m.chassisType
}
// GetChromeOSDeviceInfo gets the chromeOSDeviceInfo property value. List of properties of the ChromeOS Device.
func (m *ManagedDevice) GetChromeOSDeviceInfo()([]ChromeOSDevicePropertyable) {
    return m.chromeOSDeviceInfo
}
// GetCloudPcRemoteActionResults gets the cloudPcRemoteActionResults property value. The cloudPcRemoteActionResults property
func (m *ManagedDevice) GetCloudPcRemoteActionResults()([]CloudPcRemoteActionResultable) {
    return m.cloudPcRemoteActionResults
}
// GetComplianceGracePeriodExpirationDateTime gets the complianceGracePeriodExpirationDateTime property value. The DateTime when device compliance grace period expires. This property is read-only.
func (m *ManagedDevice) GetComplianceGracePeriodExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.complianceGracePeriodExpirationDateTime
}
// GetComplianceState gets the complianceState property value. Compliance state.
func (m *ManagedDevice) GetComplianceState()(*ComplianceState) {
    return m.complianceState
}
// GetConfigurationManagerClientEnabledFeatures gets the configurationManagerClientEnabledFeatures property value. ConfigrMgr client enabled features. This property is read-only.
func (m *ManagedDevice) GetConfigurationManagerClientEnabledFeatures()(ConfigurationManagerClientEnabledFeaturesable) {
    return m.configurationManagerClientEnabledFeatures
}
// GetConfigurationManagerClientHealthState gets the configurationManagerClientHealthState property value. Configuration manager client health state, valid only for devices managed by MDM/ConfigMgr Agent
func (m *ManagedDevice) GetConfigurationManagerClientHealthState()(ConfigurationManagerClientHealthStateable) {
    return m.configurationManagerClientHealthState
}
// GetConfigurationManagerClientInformation gets the configurationManagerClientInformation property value. Configuration manager client information, valid only for devices managed, duel-managed or tri-managed by ConfigMgr Agent
func (m *ManagedDevice) GetConfigurationManagerClientInformation()(ConfigurationManagerClientInformationable) {
    return m.configurationManagerClientInformation
}
// GetDetectedApps gets the detectedApps property value. All applications currently installed on the device
func (m *ManagedDevice) GetDetectedApps()([]DetectedAppable) {
    return m.detectedApps
}
// GetDeviceActionResults gets the deviceActionResults property value. List of ComplexType deviceActionResult objects. This property is read-only.
func (m *ManagedDevice) GetDeviceActionResults()([]DeviceActionResultable) {
    return m.deviceActionResults
}
// GetDeviceCategory gets the deviceCategory property value. Device category
func (m *ManagedDevice) GetDeviceCategory()(DeviceCategoryable) {
    return m.deviceCategory
}
// GetDeviceCategoryDisplayName gets the deviceCategoryDisplayName property value. Device category display name. This property is read-only.
func (m *ManagedDevice) GetDeviceCategoryDisplayName()(*string) {
    return m.deviceCategoryDisplayName
}
// GetDeviceCompliancePolicyStates gets the deviceCompliancePolicyStates property value. Device compliance policy states for this device.
func (m *ManagedDevice) GetDeviceCompliancePolicyStates()([]DeviceCompliancePolicyStateable) {
    return m.deviceCompliancePolicyStates
}
// GetDeviceConfigurationStates gets the deviceConfigurationStates property value. Device configuration states for this device.
func (m *ManagedDevice) GetDeviceConfigurationStates()([]DeviceConfigurationStateable) {
    return m.deviceConfigurationStates
}
// GetDeviceEnrollmentType gets the deviceEnrollmentType property value. Possible ways of adding a mobile device to management.
func (m *ManagedDevice) GetDeviceEnrollmentType()(*DeviceEnrollmentType) {
    return m.deviceEnrollmentType
}
// GetDeviceFirmwareConfigurationInterfaceManaged gets the deviceFirmwareConfigurationInterfaceManaged property value. Indicates whether the device is DFCI managed. When TRUE the device is DFCI managed. When FALSE, the device is not DFCI managed. The default value is FALSE.
func (m *ManagedDevice) GetDeviceFirmwareConfigurationInterfaceManaged()(*bool) {
    return m.deviceFirmwareConfigurationInterfaceManaged
}
// GetDeviceHealthAttestationState gets the deviceHealthAttestationState property value. The device health attestation state. This property is read-only.
func (m *ManagedDevice) GetDeviceHealthAttestationState()(DeviceHealthAttestationStateable) {
    return m.deviceHealthAttestationState
}
// GetDeviceName gets the deviceName property value. Name of the device. This property is read-only.
func (m *ManagedDevice) GetDeviceName()(*string) {
    return m.deviceName
}
// GetDeviceRegistrationState gets the deviceRegistrationState property value. Device registration status.
func (m *ManagedDevice) GetDeviceRegistrationState()(*DeviceRegistrationState) {
    return m.deviceRegistrationState
}
// GetDeviceType gets the deviceType property value. Device type.
func (m *ManagedDevice) GetDeviceType()(*DeviceType) {
    return m.deviceType
}
// GetEasActivated gets the easActivated property value. Whether the device is Exchange ActiveSync activated. This property is read-only.
func (m *ManagedDevice) GetEasActivated()(*bool) {
    return m.easActivated
}
// GetEasActivationDateTime gets the easActivationDateTime property value. Exchange ActivationSync activation time of the device. This property is read-only.
func (m *ManagedDevice) GetEasActivationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.easActivationDateTime
}
// GetEasDeviceId gets the easDeviceId property value. Exchange ActiveSync Id of the device. This property is read-only.
func (m *ManagedDevice) GetEasDeviceId()(*string) {
    return m.easDeviceId
}
// GetEmailAddress gets the emailAddress property value. Email(s) for the user associated with the device. This property is read-only.
func (m *ManagedDevice) GetEmailAddress()(*string) {
    return m.emailAddress
}
// GetEnrolledDateTime gets the enrolledDateTime property value. Enrollment time of the device. This property is read-only.
func (m *ManagedDevice) GetEnrolledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.enrolledDateTime
}
// GetEnrollmentProfileName gets the enrollmentProfileName property value. Name of the enrollment profile assigned to the device. Default value is empty string, indicating no enrollment profile was assgined. This property is read-only.
func (m *ManagedDevice) GetEnrollmentProfileName()(*string) {
    return m.enrollmentProfileName
}
// GetEthernetMacAddress gets the ethernetMacAddress property value. Ethernet MAC. Default, is Null (Non-Default property) for this property when returned as part of managedDevice entity. Individual get call with select query options is needed to retrieve actual values. Example: deviceManagement/managedDevices({managedDeviceId})?$select=ethernetMacAddress Supports: $select. $Search is not supported. Read-only. This property is read-only.
func (m *ManagedDevice) GetEthernetMacAddress()(*string) {
    return m.ethernetMacAddress
}
// GetExchangeAccessState gets the exchangeAccessState property value. Device Exchange Access State.
func (m *ManagedDevice) GetExchangeAccessState()(*DeviceManagementExchangeAccessState) {
    return m.exchangeAccessState
}
// GetExchangeAccessStateReason gets the exchangeAccessStateReason property value. Device Exchange Access State Reason.
func (m *ManagedDevice) GetExchangeAccessStateReason()(*DeviceManagementExchangeAccessStateReason) {
    return m.exchangeAccessStateReason
}
// GetExchangeLastSuccessfulSyncDateTime gets the exchangeLastSuccessfulSyncDateTime property value. Last time the device contacted Exchange. This property is read-only.
func (m *ManagedDevice) GetExchangeLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.exchangeLastSuccessfulSyncDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDevice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["aadRegistered"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAadRegistered(val)
        }
        return nil
    }
    res["activationLockBypassCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivationLockBypassCode(val)
        }
        return nil
    }
    res["androidSecurityPatchLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidSecurityPatchLevel(val)
        }
        return nil
    }
    res["assignmentFilterEvaluationStatusDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAssignmentFilterEvaluationStatusDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignmentFilterEvaluationStatusDetailsable, len(val))
            for i, v := range val {
                res[i] = v.(AssignmentFilterEvaluationStatusDetailsable)
            }
            m.SetAssignmentFilterEvaluationStatusDetails(res)
        }
        return nil
    }
    res["autopilotEnrolled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutopilotEnrolled(val)
        }
        return nil
    }
    res["azureActiveDirectoryDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureActiveDirectoryDeviceId(val)
        }
        return nil
    }
    res["azureADDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureADDeviceId(val)
        }
        return nil
    }
    res["azureADRegistered"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureADRegistered(val)
        }
        return nil
    }
    res["bootstrapTokenEscrowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBootstrapTokenEscrowed(val)
        }
        return nil
    }
    res["chassisType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseChassisType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChassisType(val.(*ChassisType))
        }
        return nil
    }
    res["chromeOSDeviceInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateChromeOSDevicePropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ChromeOSDevicePropertyable, len(val))
            for i, v := range val {
                res[i] = v.(ChromeOSDevicePropertyable)
            }
            m.SetChromeOSDeviceInfo(res)
        }
        return nil
    }
    res["cloudPcRemoteActionResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcRemoteActionResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcRemoteActionResultable, len(val))
            for i, v := range val {
                res[i] = v.(CloudPcRemoteActionResultable)
            }
            m.SetCloudPcRemoteActionResults(res)
        }
        return nil
    }
    res["complianceGracePeriodExpirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceGracePeriodExpirationDateTime(val)
        }
        return nil
    }
    res["complianceState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseComplianceState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceState(val.(*ComplianceState))
        }
        return nil
    }
    res["configurationManagerClientEnabledFeatures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConfigurationManagerClientEnabledFeaturesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationManagerClientEnabledFeatures(val.(ConfigurationManagerClientEnabledFeaturesable))
        }
        return nil
    }
    res["configurationManagerClientHealthState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConfigurationManagerClientHealthStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationManagerClientHealthState(val.(ConfigurationManagerClientHealthStateable))
        }
        return nil
    }
    res["configurationManagerClientInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConfigurationManagerClientInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationManagerClientInformation(val.(ConfigurationManagerClientInformationable))
        }
        return nil
    }
    res["detectedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDetectedAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DetectedAppable, len(val))
            for i, v := range val {
                res[i] = v.(DetectedAppable)
            }
            m.SetDetectedApps(res)
        }
        return nil
    }
    res["deviceActionResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceActionResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceActionResultable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceActionResultable)
            }
            m.SetDeviceActionResults(res)
        }
        return nil
    }
    res["deviceCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCategory(val.(DeviceCategoryable))
        }
        return nil
    }
    res["deviceCategoryDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCategoryDisplayName(val)
        }
        return nil
    }
    res["deviceCompliancePolicyStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceCompliancePolicyStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceCompliancePolicyStateable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceCompliancePolicyStateable)
            }
            m.SetDeviceCompliancePolicyStates(res)
        }
        return nil
    }
    res["deviceConfigurationStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceConfigurationStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationStateable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceConfigurationStateable)
            }
            m.SetDeviceConfigurationStates(res)
        }
        return nil
    }
    res["deviceEnrollmentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceEnrollmentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceEnrollmentType(val.(*DeviceEnrollmentType))
        }
        return nil
    }
    res["deviceFirmwareConfigurationInterfaceManaged"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceFirmwareConfigurationInterfaceManaged(val)
        }
        return nil
    }
    res["deviceHealthAttestationState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceHealthAttestationStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceHealthAttestationState(val.(DeviceHealthAttestationStateable))
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["deviceRegistrationState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceRegistrationState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceRegistrationState(val.(*DeviceRegistrationState))
        }
        return nil
    }
    res["deviceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceType(val.(*DeviceType))
        }
        return nil
    }
    res["easActivated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEasActivated(val)
        }
        return nil
    }
    res["easActivationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEasActivationDateTime(val)
        }
        return nil
    }
    res["easDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEasDeviceId(val)
        }
        return nil
    }
    res["emailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddress(val)
        }
        return nil
    }
    res["enrolledDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrolledDateTime(val)
        }
        return nil
    }
    res["enrollmentProfileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentProfileName(val)
        }
        return nil
    }
    res["ethernetMacAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEthernetMacAddress(val)
        }
        return nil
    }
    res["exchangeAccessState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeAccessState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeAccessState(val.(*DeviceManagementExchangeAccessState))
        }
        return nil
    }
    res["exchangeAccessStateReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementExchangeAccessStateReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeAccessStateReason(val.(*DeviceManagementExchangeAccessStateReason))
        }
        return nil
    }
    res["exchangeLastSuccessfulSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeLastSuccessfulSyncDateTime(val)
        }
        return nil
    }
    res["freeStorageSpaceInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFreeStorageSpaceInBytes(val)
        }
        return nil
    }
    res["hardwareInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateHardwareInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHardwareInformation(val.(HardwareInformationable))
        }
        return nil
    }
    res["iccid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIccid(val)
        }
        return nil
    }
    res["imei"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImei(val)
        }
        return nil
    }
    res["isEncrypted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEncrypted(val)
        }
        return nil
    }
    res["isSupervised"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSupervised(val)
        }
        return nil
    }
    res["jailBroken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJailBroken(val)
        }
        return nil
    }
    res["joinType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseJoinType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinType(val.(*JoinType))
        }
        return nil
    }
    res["lastSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncDateTime(val)
        }
        return nil
    }
    res["logCollectionRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceLogCollectionResponseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceLogCollectionResponseable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceLogCollectionResponseable)
            }
            m.SetLogCollectionRequests(res)
        }
        return nil
    }
    res["lostModeState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLostModeState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLostModeState(val.(*LostModeState))
        }
        return nil
    }
    res["managedDeviceMobileAppConfigurationStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceMobileAppConfigurationStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceMobileAppConfigurationStateable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedDeviceMobileAppConfigurationStateable)
            }
            m.SetManagedDeviceMobileAppConfigurationStates(res)
        }
        return nil
    }
    res["managedDeviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceName(val)
        }
        return nil
    }
    res["managedDeviceOwnerType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedDeviceOwnerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceOwnerType(val.(*ManagedDeviceOwnerType))
        }
        return nil
    }
    res["managementAgent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementAgentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementAgent(val.(*ManagementAgentType))
        }
        return nil
    }
    res["managementCertificateExpirationDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementCertificateExpirationDate(val)
        }
        return nil
    }
    res["managementFeatures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedDeviceManagementFeatures)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementFeatures(val.(*ManagedDeviceManagementFeatures))
        }
        return nil
    }
    res["managementState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementState(val.(*ManagementState))
        }
        return nil
    }
    res["manufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["meid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeid(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val)
        }
        return nil
    }
    res["operatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystem(val)
        }
        return nil
    }
    res["osVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    res["ownerType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOwnerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerType(val.(*OwnerType))
        }
        return nil
    }
    res["partnerReportedThreatState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedDevicePartnerReportedHealthState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerReportedThreatState(val.(*ManagedDevicePartnerReportedHealthState))
        }
        return nil
    }
    res["phoneNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneNumber(val)
        }
        return nil
    }
    res["physicalMemoryInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhysicalMemoryInBytes(val)
        }
        return nil
    }
    res["preferMdmOverGroupPolicyAppliedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferMdmOverGroupPolicyAppliedDateTime(val)
        }
        return nil
    }
    res["processorArchitecture"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedDeviceArchitecture)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessorArchitecture(val.(*ManagedDeviceArchitecture))
        }
        return nil
    }
    res["remoteAssistanceSessionErrorDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteAssistanceSessionErrorDetails(val)
        }
        return nil
    }
    res["remoteAssistanceSessionUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteAssistanceSessionUrl(val)
        }
        return nil
    }
    res["requireUserEnrollmentApproval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireUserEnrollmentApproval(val)
        }
        return nil
    }
    res["retireAfterDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRetireAfterDateTime(val)
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
    res["securityBaselineStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecurityBaselineStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecurityBaselineStateable, len(val))
            for i, v := range val {
                res[i] = v.(SecurityBaselineStateable)
            }
            m.SetSecurityBaselineStates(res)
        }
        return nil
    }
    res["serialNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSerialNumber(val)
        }
        return nil
    }
    res["skuFamily"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkuFamily(val)
        }
        return nil
    }
    res["skuNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkuNumber(val)
        }
        return nil
    }
    res["specificationVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpecificationVersion(val)
        }
        return nil
    }
    res["subscriberCarrier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriberCarrier(val)
        }
        return nil
    }
    res["totalStorageSpaceInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalStorageSpaceInBytes(val)
        }
        return nil
    }
    res["udid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUdid(val)
        }
        return nil
    }
    res["userDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDisplayName(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    res["users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Userable, len(val))
            for i, v := range val {
                res[i] = v.(Userable)
            }
            m.SetUsers(res)
        }
        return nil
    }
    res["usersLoggedOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLoggedOnUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LoggedOnUserable, len(val))
            for i, v := range val {
                res[i] = v.(LoggedOnUserable)
            }
            m.SetUsersLoggedOn(res)
        }
        return nil
    }
    res["wiFiMacAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWiFiMacAddress(val)
        }
        return nil
    }
    res["windowsActiveMalwareCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsActiveMalwareCount(val)
        }
        return nil
    }
    res["windowsProtectionState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsProtectionStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsProtectionState(val.(WindowsProtectionStateable))
        }
        return nil
    }
    res["windowsRemediatedMalwareCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsRemediatedMalwareCount(val)
        }
        return nil
    }
    return res
}
// GetFreeStorageSpaceInBytes gets the freeStorageSpaceInBytes property value. Free Storage in Bytes. Default value is 0. Read-only. This property is read-only.
func (m *ManagedDevice) GetFreeStorageSpaceInBytes()(*int64) {
    return m.freeStorageSpaceInBytes
}
// GetHardwareInformation gets the hardwareInformation property value. The hardward details for the device.  Includes information such as storage space, manufacturer, serial number, etc. Return default value in LIST managedDevices. Real value only returned in singel device GET call with device id and included in select parameter. Supports: $select. $Search is not supported. Read-only. This property is read-only.
func (m *ManagedDevice) GetHardwareInformation()(HardwareInformationable) {
    return m.hardwareInformation
}
// GetIccid gets the iccid property value. Integrated Circuit Card Identifier, it is A SIM card's unique identification number. Return default value null in LIST managedDevices. Real value only returned in singel device GET call with device id and included in select parameter. Supports: $select. $Search is not supported. Read-only. This property is read-only.
func (m *ManagedDevice) GetIccid()(*string) {
    return m.iccid
}
// GetImei gets the imei property value. IMEI. This property is read-only.
func (m *ManagedDevice) GetImei()(*string) {
    return m.imei
}
// GetIsEncrypted gets the isEncrypted property value. Device encryption status. This property is read-only.
func (m *ManagedDevice) GetIsEncrypted()(*bool) {
    return m.isEncrypted
}
// GetIsSupervised gets the isSupervised property value. Device supervised status. This property is read-only.
func (m *ManagedDevice) GetIsSupervised()(*bool) {
    return m.isSupervised
}
// GetJailBroken gets the jailBroken property value. whether the device is jail broken or rooted. This property is read-only.
func (m *ManagedDevice) GetJailBroken()(*string) {
    return m.jailBroken
}
// GetJoinType gets the joinType property value. Device enrollment join type.
func (m *ManagedDevice) GetJoinType()(*JoinType) {
    return m.joinType
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. The date and time that the device last completed a successful sync with Intune. This property is read-only.
func (m *ManagedDevice) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSyncDateTime
}
// GetLogCollectionRequests gets the logCollectionRequests property value. List of log collection requests
func (m *ManagedDevice) GetLogCollectionRequests()([]DeviceLogCollectionResponseable) {
    return m.logCollectionRequests
}
// GetLostModeState gets the lostModeState property value. State of lost mode, indicating if lost mode is enabled or disabled
func (m *ManagedDevice) GetLostModeState()(*LostModeState) {
    return m.lostModeState
}
// GetManagedDeviceMobileAppConfigurationStates gets the managedDeviceMobileAppConfigurationStates property value. Managed device mobile app configuration states for this device.
func (m *ManagedDevice) GetManagedDeviceMobileAppConfigurationStates()([]ManagedDeviceMobileAppConfigurationStateable) {
    return m.managedDeviceMobileAppConfigurationStates
}
// GetManagedDeviceName gets the managedDeviceName property value. Automatically generated name to identify a device. Can be overwritten to a user friendly name.
func (m *ManagedDevice) GetManagedDeviceName()(*string) {
    return m.managedDeviceName
}
// GetManagedDeviceOwnerType gets the managedDeviceOwnerType property value. Owner type of device.
func (m *ManagedDevice) GetManagedDeviceOwnerType()(*ManagedDeviceOwnerType) {
    return m.managedDeviceOwnerType
}
// GetManagementAgent gets the managementAgent property value. Management agent type.
func (m *ManagedDevice) GetManagementAgent()(*ManagementAgentType) {
    return m.managementAgent
}
// GetManagementCertificateExpirationDate gets the managementCertificateExpirationDate property value. Reports device management certificate expiration date. This property is read-only.
func (m *ManagedDevice) GetManagementCertificateExpirationDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.managementCertificateExpirationDate
}
// GetManagementFeatures gets the managementFeatures property value. Device management features.
func (m *ManagedDevice) GetManagementFeatures()(*ManagedDeviceManagementFeatures) {
    return m.managementFeatures
}
// GetManagementState gets the managementState property value. Management state of device in Microsoft Intune.
func (m *ManagedDevice) GetManagementState()(*ManagementState) {
    return m.managementState
}
// GetManufacturer gets the manufacturer property value. Manufacturer of the device. This property is read-only.
func (m *ManagedDevice) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetMeid gets the meid property value. MEID. This property is read-only.
func (m *ManagedDevice) GetMeid()(*string) {
    return m.meid
}
// GetModel gets the model property value. Model of the device. This property is read-only.
func (m *ManagedDevice) GetModel()(*string) {
    return m.model
}
// GetNotes gets the notes property value. Notes on the device created by IT Admin. Return default value null in LIST managedDevices. Real value only returned in singel device GET call with device id and included in select parameter. Supports: $select.  $Search is not supported.
func (m *ManagedDevice) GetNotes()(*string) {
    return m.notes
}
// GetOperatingSystem gets the operatingSystem property value. Operating system of the device. Windows, iOS, etc. This property is read-only.
func (m *ManagedDevice) GetOperatingSystem()(*string) {
    return m.operatingSystem
}
// GetOsVersion gets the osVersion property value. Operating system version of the device. This property is read-only.
func (m *ManagedDevice) GetOsVersion()(*string) {
    return m.osVersion
}
// GetOwnerType gets the ownerType property value. Owner type of device.
func (m *ManagedDevice) GetOwnerType()(*OwnerType) {
    return m.ownerType
}
// GetPartnerReportedThreatState gets the partnerReportedThreatState property value. Available health states for the Device Health API
func (m *ManagedDevice) GetPartnerReportedThreatState()(*ManagedDevicePartnerReportedHealthState) {
    return m.partnerReportedThreatState
}
// GetPhoneNumber gets the phoneNumber property value. Phone number of the device. This property is read-only.
func (m *ManagedDevice) GetPhoneNumber()(*string) {
    return m.phoneNumber
}
// GetPhysicalMemoryInBytes gets the physicalMemoryInBytes property value. Total Memory in Bytes. Return default value 0 in LIST managedDevices. Real value only returned in singel device GET call with device id and included in select parameter. Supports: $select. Default value is 0. Read-only. This property is read-only.
func (m *ManagedDevice) GetPhysicalMemoryInBytes()(*int64) {
    return m.physicalMemoryInBytes
}
// GetPreferMdmOverGroupPolicyAppliedDateTime gets the preferMdmOverGroupPolicyAppliedDateTime property value. Reports the DateTime the preferMdmOverGroupPolicy setting was set.  When set, the Intune MDM settings will override Group Policy settings if there is a conflict. Read Only. This property is read-only.
func (m *ManagedDevice) GetPreferMdmOverGroupPolicyAppliedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.preferMdmOverGroupPolicyAppliedDateTime
}
// GetProcessorArchitecture gets the processorArchitecture property value. Processor architecture
func (m *ManagedDevice) GetProcessorArchitecture()(*ManagedDeviceArchitecture) {
    return m.processorArchitecture
}
// GetRemoteAssistanceSessionErrorDetails gets the remoteAssistanceSessionErrorDetails property value. An error string that identifies issues when creating Remote Assistance session objects. This property is read-only.
func (m *ManagedDevice) GetRemoteAssistanceSessionErrorDetails()(*string) {
    return m.remoteAssistanceSessionErrorDetails
}
// GetRemoteAssistanceSessionUrl gets the remoteAssistanceSessionUrl property value. Url that allows a Remote Assistance session to be established with the device. This property is read-only.
func (m *ManagedDevice) GetRemoteAssistanceSessionUrl()(*string) {
    return m.remoteAssistanceSessionUrl
}
// GetRequireUserEnrollmentApproval gets the requireUserEnrollmentApproval property value. Reports if the managed iOS device is user approval enrollment. This property is read-only.
func (m *ManagedDevice) GetRequireUserEnrollmentApproval()(*bool) {
    return m.requireUserEnrollmentApproval
}
// GetRetireAfterDateTime gets the retireAfterDateTime property value. Indicates the time after when a device will be auto retired because of scheduled action. This property is read-only.
func (m *ManagedDevice) GetRetireAfterDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.retireAfterDateTime
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tag IDs for this Device instance.
func (m *ManagedDevice) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// GetSecurityBaselineStates gets the securityBaselineStates property value. Security baseline states for this device.
func (m *ManagedDevice) GetSecurityBaselineStates()([]SecurityBaselineStateable) {
    return m.securityBaselineStates
}
// GetSerialNumber gets the serialNumber property value. SerialNumber. This property is read-only.
func (m *ManagedDevice) GetSerialNumber()(*string) {
    return m.serialNumber
}
// GetSkuFamily gets the skuFamily property value. Device sku family
func (m *ManagedDevice) GetSkuFamily()(*string) {
    return m.skuFamily
}
// GetSkuNumber gets the skuNumber property value. Device sku number, see also: https://learn.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-getproductinfo. Valid values 0 to 2147483647. This property is read-only.
func (m *ManagedDevice) GetSkuNumber()(*int32) {
    return m.skuNumber
}
// GetSpecificationVersion gets the specificationVersion property value. Specification version. This property is read-only.
func (m *ManagedDevice) GetSpecificationVersion()(*string) {
    return m.specificationVersion
}
// GetSubscriberCarrier gets the subscriberCarrier property value. Subscriber Carrier. This property is read-only.
func (m *ManagedDevice) GetSubscriberCarrier()(*string) {
    return m.subscriberCarrier
}
// GetTotalStorageSpaceInBytes gets the totalStorageSpaceInBytes property value. Total Storage in Bytes. This property is read-only.
func (m *ManagedDevice) GetTotalStorageSpaceInBytes()(*int64) {
    return m.totalStorageSpaceInBytes
}
// GetUdid gets the udid property value. Unique Device Identifier for iOS and macOS devices. Return default value null in LIST managedDevices. Real value only returned in singel device GET call with device id and included in select parameter. Supports: $select. $Search is not supported. Read-only. This property is read-only.
func (m *ManagedDevice) GetUdid()(*string) {
    return m.udid
}
// GetUserDisplayName gets the userDisplayName property value. User display name. This property is read-only.
func (m *ManagedDevice) GetUserDisplayName()(*string) {
    return m.userDisplayName
}
// GetUserId gets the userId property value. Unique Identifier for the user associated with the device. This property is read-only.
func (m *ManagedDevice) GetUserId()(*string) {
    return m.userId
}
// GetUserPrincipalName gets the userPrincipalName property value. Device user principal name. This property is read-only.
func (m *ManagedDevice) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// GetUsers gets the users property value. The primary users associated with the managed device.
func (m *ManagedDevice) GetUsers()([]Userable) {
    return m.users
}
// GetUsersLoggedOn gets the usersLoggedOn property value. Indicates the last logged on users of a device. This property is read-only.
func (m *ManagedDevice) GetUsersLoggedOn()([]LoggedOnUserable) {
    return m.usersLoggedOn
}
// GetWiFiMacAddress gets the wiFiMacAddress property value. Wi-Fi MAC. This property is read-only.
func (m *ManagedDevice) GetWiFiMacAddress()(*string) {
    return m.wiFiMacAddress
}
// GetWindowsActiveMalwareCount gets the windowsActiveMalwareCount property value. Count of active malware for this windows device. This property is read-only.
func (m *ManagedDevice) GetWindowsActiveMalwareCount()(*int32) {
    return m.windowsActiveMalwareCount
}
// GetWindowsProtectionState gets the windowsProtectionState property value. The device protection status. This property is read-only.
func (m *ManagedDevice) GetWindowsProtectionState()(WindowsProtectionStateable) {
    return m.windowsProtectionState
}
// GetWindowsRemediatedMalwareCount gets the windowsRemediatedMalwareCount property value. Count of remediated malware for this windows device. This property is read-only.
func (m *ManagedDevice) GetWindowsRemediatedMalwareCount()(*int32) {
    return m.windowsRemediatedMalwareCount
}
// Serialize serializes information the current object
func (m *ManagedDevice) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignmentFilterEvaluationStatusDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignmentFilterEvaluationStatusDetails()))
        for i, v := range m.GetAssignmentFilterEvaluationStatusDetails() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignmentFilterEvaluationStatusDetails", cast)
        if err != nil {
            return err
        }
    }
    if m.GetChassisType() != nil {
        cast := (*m.GetChassisType()).String()
        err = writer.WriteStringValue("chassisType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetChromeOSDeviceInfo() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChromeOSDeviceInfo()))
        for i, v := range m.GetChromeOSDeviceInfo() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("chromeOSDeviceInfo", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCloudPcRemoteActionResults() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCloudPcRemoteActionResults()))
        for i, v := range m.GetCloudPcRemoteActionResults() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("cloudPcRemoteActionResults", cast)
        if err != nil {
            return err
        }
    }
    if m.GetComplianceState() != nil {
        cast := (*m.GetComplianceState()).String()
        err = writer.WriteStringValue("complianceState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("configurationManagerClientHealthState", m.GetConfigurationManagerClientHealthState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("configurationManagerClientInformation", m.GetConfigurationManagerClientInformation())
        if err != nil {
            return err
        }
    }
    if m.GetDetectedApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDetectedApps()))
        for i, v := range m.GetDetectedApps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("detectedApps", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceCategory", m.GetDeviceCategory())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceCompliancePolicyStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceCompliancePolicyStates()))
        for i, v := range m.GetDeviceCompliancePolicyStates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceCompliancePolicyStates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceConfigurationStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceConfigurationStates()))
        for i, v := range m.GetDeviceConfigurationStates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceConfigurationStates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceEnrollmentType() != nil {
        cast := (*m.GetDeviceEnrollmentType()).String()
        err = writer.WriteStringValue("deviceEnrollmentType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deviceFirmwareConfigurationInterfaceManaged", m.GetDeviceFirmwareConfigurationInterfaceManaged())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceRegistrationState() != nil {
        cast := (*m.GetDeviceRegistrationState()).String()
        err = writer.WriteStringValue("deviceRegistrationState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceType() != nil {
        cast := (*m.GetDeviceType()).String()
        err = writer.WriteStringValue("deviceType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetExchangeAccessState() != nil {
        cast := (*m.GetExchangeAccessState()).String()
        err = writer.WriteStringValue("exchangeAccessState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetExchangeAccessStateReason() != nil {
        cast := (*m.GetExchangeAccessStateReason()).String()
        err = writer.WriteStringValue("exchangeAccessStateReason", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetJoinType() != nil {
        cast := (*m.GetJoinType()).String()
        err = writer.WriteStringValue("joinType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetLogCollectionRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLogCollectionRequests()))
        for i, v := range m.GetLogCollectionRequests() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("logCollectionRequests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetLostModeState() != nil {
        cast := (*m.GetLostModeState()).String()
        err = writer.WriteStringValue("lostModeState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedDeviceMobileAppConfigurationStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedDeviceMobileAppConfigurationStates()))
        for i, v := range m.GetManagedDeviceMobileAppConfigurationStates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedDeviceMobileAppConfigurationStates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceName", m.GetManagedDeviceName())
        if err != nil {
            return err
        }
    }
    if m.GetManagedDeviceOwnerType() != nil {
        cast := (*m.GetManagedDeviceOwnerType()).String()
        err = writer.WriteStringValue("managedDeviceOwnerType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementAgent() != nil {
        cast := (*m.GetManagementAgent()).String()
        err = writer.WriteStringValue("managementAgent", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementFeatures() != nil {
        cast := (*m.GetManagementFeatures()).String()
        err = writer.WriteStringValue("managementFeatures", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementState() != nil {
        cast := (*m.GetManagementState()).String()
        err = writer.WriteStringValue("managementState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    if m.GetOwnerType() != nil {
        cast := (*m.GetOwnerType()).String()
        err = writer.WriteStringValue("ownerType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPartnerReportedThreatState() != nil {
        cast := (*m.GetPartnerReportedThreatState()).String()
        err = writer.WriteStringValue("partnerReportedThreatState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetProcessorArchitecture() != nil {
        cast := (*m.GetProcessorArchitecture()).String()
        err = writer.WriteStringValue("processorArchitecture", &cast)
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
    if m.GetSecurityBaselineStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSecurityBaselineStates()))
        for i, v := range m.GetSecurityBaselineStates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("securityBaselineStates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("skuFamily", m.GetSkuFamily())
        if err != nil {
            return err
        }
    }
    if m.GetUsers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUsers()))
        for i, v := range m.GetUsers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("users", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("windowsProtectionState", m.GetWindowsProtectionState())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAadRegistered sets the aadRegistered property value. Whether the device is Azure Active Directory registered. This property is read-only.
func (m *ManagedDevice) SetAadRegistered(value *bool)() {
    m.aadRegistered = value
}
// SetActivationLockBypassCode sets the activationLockBypassCode property value. The code that allows the Activation Lock on managed device to be bypassed. Default, is Null (Non-Default property) for this property when returned as part of managedDevice entity in LIST call. Individual GET call with select query options is needed to retrieve actual values. Supports: $select. $Search is not supported. Read-only. This property is read-only.
func (m *ManagedDevice) SetActivationLockBypassCode(value *string)() {
    m.activationLockBypassCode = value
}
// SetAndroidSecurityPatchLevel sets the androidSecurityPatchLevel property value. Android security patch level. This property is read-only.
func (m *ManagedDevice) SetAndroidSecurityPatchLevel(value *string)() {
    m.androidSecurityPatchLevel = value
}
// SetAssignmentFilterEvaluationStatusDetails sets the assignmentFilterEvaluationStatusDetails property value. Managed device mobile app configuration states for this device.
func (m *ManagedDevice) SetAssignmentFilterEvaluationStatusDetails(value []AssignmentFilterEvaluationStatusDetailsable)() {
    m.assignmentFilterEvaluationStatusDetails = value
}
// SetAutopilotEnrolled sets the autopilotEnrolled property value. Reports if the managed device is enrolled via auto-pilot. This property is read-only.
func (m *ManagedDevice) SetAutopilotEnrolled(value *bool)() {
    m.autopilotEnrolled = value
}
// SetAzureActiveDirectoryDeviceId sets the azureActiveDirectoryDeviceId property value. The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
func (m *ManagedDevice) SetAzureActiveDirectoryDeviceId(value *string)() {
    m.azureActiveDirectoryDeviceId = value
}
// SetAzureADDeviceId sets the azureADDeviceId property value. The unique identifier for the Azure Active Directory device. Read only. This property is read-only.
func (m *ManagedDevice) SetAzureADDeviceId(value *string)() {
    m.azureADDeviceId = value
}
// SetAzureADRegistered sets the azureADRegistered property value. Whether the device is Azure Active Directory registered. This property is read-only.
func (m *ManagedDevice) SetAzureADRegistered(value *bool)() {
    m.azureADRegistered = value
}
// SetBootstrapTokenEscrowed sets the bootstrapTokenEscrowed property value. Reports if the managed device has an escrowed Bootstrap Token. This is only for macOS devices. To get, include BootstrapTokenEscrowed in the select clause and query with a device id. If FALSE, no bootstrap token is escrowed. If TRUE, the device has escrowed a bootstrap token with Intune. This property is read-only.
func (m *ManagedDevice) SetBootstrapTokenEscrowed(value *bool)() {
    m.bootstrapTokenEscrowed = value
}
// SetChassisType sets the chassisType property value. Chassis type.
func (m *ManagedDevice) SetChassisType(value *ChassisType)() {
    m.chassisType = value
}
// SetChromeOSDeviceInfo sets the chromeOSDeviceInfo property value. List of properties of the ChromeOS Device.
func (m *ManagedDevice) SetChromeOSDeviceInfo(value []ChromeOSDevicePropertyable)() {
    m.chromeOSDeviceInfo = value
}
// SetCloudPcRemoteActionResults sets the cloudPcRemoteActionResults property value. The cloudPcRemoteActionResults property
func (m *ManagedDevice) SetCloudPcRemoteActionResults(value []CloudPcRemoteActionResultable)() {
    m.cloudPcRemoteActionResults = value
}
// SetComplianceGracePeriodExpirationDateTime sets the complianceGracePeriodExpirationDateTime property value. The DateTime when device compliance grace period expires. This property is read-only.
func (m *ManagedDevice) SetComplianceGracePeriodExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.complianceGracePeriodExpirationDateTime = value
}
// SetComplianceState sets the complianceState property value. Compliance state.
func (m *ManagedDevice) SetComplianceState(value *ComplianceState)() {
    m.complianceState = value
}
// SetConfigurationManagerClientEnabledFeatures sets the configurationManagerClientEnabledFeatures property value. ConfigrMgr client enabled features. This property is read-only.
func (m *ManagedDevice) SetConfigurationManagerClientEnabledFeatures(value ConfigurationManagerClientEnabledFeaturesable)() {
    m.configurationManagerClientEnabledFeatures = value
}
// SetConfigurationManagerClientHealthState sets the configurationManagerClientHealthState property value. Configuration manager client health state, valid only for devices managed by MDM/ConfigMgr Agent
func (m *ManagedDevice) SetConfigurationManagerClientHealthState(value ConfigurationManagerClientHealthStateable)() {
    m.configurationManagerClientHealthState = value
}
// SetConfigurationManagerClientInformation sets the configurationManagerClientInformation property value. Configuration manager client information, valid only for devices managed, duel-managed or tri-managed by ConfigMgr Agent
func (m *ManagedDevice) SetConfigurationManagerClientInformation(value ConfigurationManagerClientInformationable)() {
    m.configurationManagerClientInformation = value
}
// SetDetectedApps sets the detectedApps property value. All applications currently installed on the device
func (m *ManagedDevice) SetDetectedApps(value []DetectedAppable)() {
    m.detectedApps = value
}
// SetDeviceActionResults sets the deviceActionResults property value. List of ComplexType deviceActionResult objects. This property is read-only.
func (m *ManagedDevice) SetDeviceActionResults(value []DeviceActionResultable)() {
    m.deviceActionResults = value
}
// SetDeviceCategory sets the deviceCategory property value. Device category
func (m *ManagedDevice) SetDeviceCategory(value DeviceCategoryable)() {
    m.deviceCategory = value
}
// SetDeviceCategoryDisplayName sets the deviceCategoryDisplayName property value. Device category display name. This property is read-only.
func (m *ManagedDevice) SetDeviceCategoryDisplayName(value *string)() {
    m.deviceCategoryDisplayName = value
}
// SetDeviceCompliancePolicyStates sets the deviceCompliancePolicyStates property value. Device compliance policy states for this device.
func (m *ManagedDevice) SetDeviceCompliancePolicyStates(value []DeviceCompliancePolicyStateable)() {
    m.deviceCompliancePolicyStates = value
}
// SetDeviceConfigurationStates sets the deviceConfigurationStates property value. Device configuration states for this device.
func (m *ManagedDevice) SetDeviceConfigurationStates(value []DeviceConfigurationStateable)() {
    m.deviceConfigurationStates = value
}
// SetDeviceEnrollmentType sets the deviceEnrollmentType property value. Possible ways of adding a mobile device to management.
func (m *ManagedDevice) SetDeviceEnrollmentType(value *DeviceEnrollmentType)() {
    m.deviceEnrollmentType = value
}
// SetDeviceFirmwareConfigurationInterfaceManaged sets the deviceFirmwareConfigurationInterfaceManaged property value. Indicates whether the device is DFCI managed. When TRUE the device is DFCI managed. When FALSE, the device is not DFCI managed. The default value is FALSE.
func (m *ManagedDevice) SetDeviceFirmwareConfigurationInterfaceManaged(value *bool)() {
    m.deviceFirmwareConfigurationInterfaceManaged = value
}
// SetDeviceHealthAttestationState sets the deviceHealthAttestationState property value. The device health attestation state. This property is read-only.
func (m *ManagedDevice) SetDeviceHealthAttestationState(value DeviceHealthAttestationStateable)() {
    m.deviceHealthAttestationState = value
}
// SetDeviceName sets the deviceName property value. Name of the device. This property is read-only.
func (m *ManagedDevice) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetDeviceRegistrationState sets the deviceRegistrationState property value. Device registration status.
func (m *ManagedDevice) SetDeviceRegistrationState(value *DeviceRegistrationState)() {
    m.deviceRegistrationState = value
}
// SetDeviceType sets the deviceType property value. Device type.
func (m *ManagedDevice) SetDeviceType(value *DeviceType)() {
    m.deviceType = value
}
// SetEasActivated sets the easActivated property value. Whether the device is Exchange ActiveSync activated. This property is read-only.
func (m *ManagedDevice) SetEasActivated(value *bool)() {
    m.easActivated = value
}
// SetEasActivationDateTime sets the easActivationDateTime property value. Exchange ActivationSync activation time of the device. This property is read-only.
func (m *ManagedDevice) SetEasActivationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.easActivationDateTime = value
}
// SetEasDeviceId sets the easDeviceId property value. Exchange ActiveSync Id of the device. This property is read-only.
func (m *ManagedDevice) SetEasDeviceId(value *string)() {
    m.easDeviceId = value
}
// SetEmailAddress sets the emailAddress property value. Email(s) for the user associated with the device. This property is read-only.
func (m *ManagedDevice) SetEmailAddress(value *string)() {
    m.emailAddress = value
}
// SetEnrolledDateTime sets the enrolledDateTime property value. Enrollment time of the device. This property is read-only.
func (m *ManagedDevice) SetEnrolledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.enrolledDateTime = value
}
// SetEnrollmentProfileName sets the enrollmentProfileName property value. Name of the enrollment profile assigned to the device. Default value is empty string, indicating no enrollment profile was assgined. This property is read-only.
func (m *ManagedDevice) SetEnrollmentProfileName(value *string)() {
    m.enrollmentProfileName = value
}
// SetEthernetMacAddress sets the ethernetMacAddress property value. Ethernet MAC. Default, is Null (Non-Default property) for this property when returned as part of managedDevice entity. Individual get call with select query options is needed to retrieve actual values. Example: deviceManagement/managedDevices({managedDeviceId})?$select=ethernetMacAddress Supports: $select. $Search is not supported. Read-only. This property is read-only.
func (m *ManagedDevice) SetEthernetMacAddress(value *string)() {
    m.ethernetMacAddress = value
}
// SetExchangeAccessState sets the exchangeAccessState property value. Device Exchange Access State.
func (m *ManagedDevice) SetExchangeAccessState(value *DeviceManagementExchangeAccessState)() {
    m.exchangeAccessState = value
}
// SetExchangeAccessStateReason sets the exchangeAccessStateReason property value. Device Exchange Access State Reason.
func (m *ManagedDevice) SetExchangeAccessStateReason(value *DeviceManagementExchangeAccessStateReason)() {
    m.exchangeAccessStateReason = value
}
// SetExchangeLastSuccessfulSyncDateTime sets the exchangeLastSuccessfulSyncDateTime property value. Last time the device contacted Exchange. This property is read-only.
func (m *ManagedDevice) SetExchangeLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.exchangeLastSuccessfulSyncDateTime = value
}
// SetFreeStorageSpaceInBytes sets the freeStorageSpaceInBytes property value. Free Storage in Bytes. Default value is 0. Read-only. This property is read-only.
func (m *ManagedDevice) SetFreeStorageSpaceInBytes(value *int64)() {
    m.freeStorageSpaceInBytes = value
}
// SetHardwareInformation sets the hardwareInformation property value. The hardward details for the device.  Includes information such as storage space, manufacturer, serial number, etc. Return default value in LIST managedDevices. Real value only returned in singel device GET call with device id and included in select parameter. Supports: $select. $Search is not supported. Read-only. This property is read-only.
func (m *ManagedDevice) SetHardwareInformation(value HardwareInformationable)() {
    m.hardwareInformation = value
}
// SetIccid sets the iccid property value. Integrated Circuit Card Identifier, it is A SIM card's unique identification number. Return default value null in LIST managedDevices. Real value only returned in singel device GET call with device id and included in select parameter. Supports: $select. $Search is not supported. Read-only. This property is read-only.
func (m *ManagedDevice) SetIccid(value *string)() {
    m.iccid = value
}
// SetImei sets the imei property value. IMEI. This property is read-only.
func (m *ManagedDevice) SetImei(value *string)() {
    m.imei = value
}
// SetIsEncrypted sets the isEncrypted property value. Device encryption status. This property is read-only.
func (m *ManagedDevice) SetIsEncrypted(value *bool)() {
    m.isEncrypted = value
}
// SetIsSupervised sets the isSupervised property value. Device supervised status. This property is read-only.
func (m *ManagedDevice) SetIsSupervised(value *bool)() {
    m.isSupervised = value
}
// SetJailBroken sets the jailBroken property value. whether the device is jail broken or rooted. This property is read-only.
func (m *ManagedDevice) SetJailBroken(value *string)() {
    m.jailBroken = value
}
// SetJoinType sets the joinType property value. Device enrollment join type.
func (m *ManagedDevice) SetJoinType(value *JoinType)() {
    m.joinType = value
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. The date and time that the device last completed a successful sync with Intune. This property is read-only.
func (m *ManagedDevice) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
// SetLogCollectionRequests sets the logCollectionRequests property value. List of log collection requests
func (m *ManagedDevice) SetLogCollectionRequests(value []DeviceLogCollectionResponseable)() {
    m.logCollectionRequests = value
}
// SetLostModeState sets the lostModeState property value. State of lost mode, indicating if lost mode is enabled or disabled
func (m *ManagedDevice) SetLostModeState(value *LostModeState)() {
    m.lostModeState = value
}
// SetManagedDeviceMobileAppConfigurationStates sets the managedDeviceMobileAppConfigurationStates property value. Managed device mobile app configuration states for this device.
func (m *ManagedDevice) SetManagedDeviceMobileAppConfigurationStates(value []ManagedDeviceMobileAppConfigurationStateable)() {
    m.managedDeviceMobileAppConfigurationStates = value
}
// SetManagedDeviceName sets the managedDeviceName property value. Automatically generated name to identify a device. Can be overwritten to a user friendly name.
func (m *ManagedDevice) SetManagedDeviceName(value *string)() {
    m.managedDeviceName = value
}
// SetManagedDeviceOwnerType sets the managedDeviceOwnerType property value. Owner type of device.
func (m *ManagedDevice) SetManagedDeviceOwnerType(value *ManagedDeviceOwnerType)() {
    m.managedDeviceOwnerType = value
}
// SetManagementAgent sets the managementAgent property value. Management agent type.
func (m *ManagedDevice) SetManagementAgent(value *ManagementAgentType)() {
    m.managementAgent = value
}
// SetManagementCertificateExpirationDate sets the managementCertificateExpirationDate property value. Reports device management certificate expiration date. This property is read-only.
func (m *ManagedDevice) SetManagementCertificateExpirationDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.managementCertificateExpirationDate = value
}
// SetManagementFeatures sets the managementFeatures property value. Device management features.
func (m *ManagedDevice) SetManagementFeatures(value *ManagedDeviceManagementFeatures)() {
    m.managementFeatures = value
}
// SetManagementState sets the managementState property value. Management state of device in Microsoft Intune.
func (m *ManagedDevice) SetManagementState(value *ManagementState)() {
    m.managementState = value
}
// SetManufacturer sets the manufacturer property value. Manufacturer of the device. This property is read-only.
func (m *ManagedDevice) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetMeid sets the meid property value. MEID. This property is read-only.
func (m *ManagedDevice) SetMeid(value *string)() {
    m.meid = value
}
// SetModel sets the model property value. Model of the device. This property is read-only.
func (m *ManagedDevice) SetModel(value *string)() {
    m.model = value
}
// SetNotes sets the notes property value. Notes on the device created by IT Admin. Return default value null in LIST managedDevices. Real value only returned in singel device GET call with device id and included in select parameter. Supports: $select.  $Search is not supported.
func (m *ManagedDevice) SetNotes(value *string)() {
    m.notes = value
}
// SetOperatingSystem sets the operatingSystem property value. Operating system of the device. Windows, iOS, etc. This property is read-only.
func (m *ManagedDevice) SetOperatingSystem(value *string)() {
    m.operatingSystem = value
}
// SetOsVersion sets the osVersion property value. Operating system version of the device. This property is read-only.
func (m *ManagedDevice) SetOsVersion(value *string)() {
    m.osVersion = value
}
// SetOwnerType sets the ownerType property value. Owner type of device.
func (m *ManagedDevice) SetOwnerType(value *OwnerType)() {
    m.ownerType = value
}
// SetPartnerReportedThreatState sets the partnerReportedThreatState property value. Available health states for the Device Health API
func (m *ManagedDevice) SetPartnerReportedThreatState(value *ManagedDevicePartnerReportedHealthState)() {
    m.partnerReportedThreatState = value
}
// SetPhoneNumber sets the phoneNumber property value. Phone number of the device. This property is read-only.
func (m *ManagedDevice) SetPhoneNumber(value *string)() {
    m.phoneNumber = value
}
// SetPhysicalMemoryInBytes sets the physicalMemoryInBytes property value. Total Memory in Bytes. Return default value 0 in LIST managedDevices. Real value only returned in singel device GET call with device id and included in select parameter. Supports: $select. Default value is 0. Read-only. This property is read-only.
func (m *ManagedDevice) SetPhysicalMemoryInBytes(value *int64)() {
    m.physicalMemoryInBytes = value
}
// SetPreferMdmOverGroupPolicyAppliedDateTime sets the preferMdmOverGroupPolicyAppliedDateTime property value. Reports the DateTime the preferMdmOverGroupPolicy setting was set.  When set, the Intune MDM settings will override Group Policy settings if there is a conflict. Read Only. This property is read-only.
func (m *ManagedDevice) SetPreferMdmOverGroupPolicyAppliedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.preferMdmOverGroupPolicyAppliedDateTime = value
}
// SetProcessorArchitecture sets the processorArchitecture property value. Processor architecture
func (m *ManagedDevice) SetProcessorArchitecture(value *ManagedDeviceArchitecture)() {
    m.processorArchitecture = value
}
// SetRemoteAssistanceSessionErrorDetails sets the remoteAssistanceSessionErrorDetails property value. An error string that identifies issues when creating Remote Assistance session objects. This property is read-only.
func (m *ManagedDevice) SetRemoteAssistanceSessionErrorDetails(value *string)() {
    m.remoteAssistanceSessionErrorDetails = value
}
// SetRemoteAssistanceSessionUrl sets the remoteAssistanceSessionUrl property value. Url that allows a Remote Assistance session to be established with the device. This property is read-only.
func (m *ManagedDevice) SetRemoteAssistanceSessionUrl(value *string)() {
    m.remoteAssistanceSessionUrl = value
}
// SetRequireUserEnrollmentApproval sets the requireUserEnrollmentApproval property value. Reports if the managed iOS device is user approval enrollment. This property is read-only.
func (m *ManagedDevice) SetRequireUserEnrollmentApproval(value *bool)() {
    m.requireUserEnrollmentApproval = value
}
// SetRetireAfterDateTime sets the retireAfterDateTime property value. Indicates the time after when a device will be auto retired because of scheduled action. This property is read-only.
func (m *ManagedDevice) SetRetireAfterDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.retireAfterDateTime = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tag IDs for this Device instance.
func (m *ManagedDevice) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// SetSecurityBaselineStates sets the securityBaselineStates property value. Security baseline states for this device.
func (m *ManagedDevice) SetSecurityBaselineStates(value []SecurityBaselineStateable)() {
    m.securityBaselineStates = value
}
// SetSerialNumber sets the serialNumber property value. SerialNumber. This property is read-only.
func (m *ManagedDevice) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
// SetSkuFamily sets the skuFamily property value. Device sku family
func (m *ManagedDevice) SetSkuFamily(value *string)() {
    m.skuFamily = value
}
// SetSkuNumber sets the skuNumber property value. Device sku number, see also: https://learn.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-getproductinfo. Valid values 0 to 2147483647. This property is read-only.
func (m *ManagedDevice) SetSkuNumber(value *int32)() {
    m.skuNumber = value
}
// SetSpecificationVersion sets the specificationVersion property value. Specification version. This property is read-only.
func (m *ManagedDevice) SetSpecificationVersion(value *string)() {
    m.specificationVersion = value
}
// SetSubscriberCarrier sets the subscriberCarrier property value. Subscriber Carrier. This property is read-only.
func (m *ManagedDevice) SetSubscriberCarrier(value *string)() {
    m.subscriberCarrier = value
}
// SetTotalStorageSpaceInBytes sets the totalStorageSpaceInBytes property value. Total Storage in Bytes. This property is read-only.
func (m *ManagedDevice) SetTotalStorageSpaceInBytes(value *int64)() {
    m.totalStorageSpaceInBytes = value
}
// SetUdid sets the udid property value. Unique Device Identifier for iOS and macOS devices. Return default value null in LIST managedDevices. Real value only returned in singel device GET call with device id and included in select parameter. Supports: $select. $Search is not supported. Read-only. This property is read-only.
func (m *ManagedDevice) SetUdid(value *string)() {
    m.udid = value
}
// SetUserDisplayName sets the userDisplayName property value. User display name. This property is read-only.
func (m *ManagedDevice) SetUserDisplayName(value *string)() {
    m.userDisplayName = value
}
// SetUserId sets the userId property value. Unique Identifier for the user associated with the device. This property is read-only.
func (m *ManagedDevice) SetUserId(value *string)() {
    m.userId = value
}
// SetUserPrincipalName sets the userPrincipalName property value. Device user principal name. This property is read-only.
func (m *ManagedDevice) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// SetUsers sets the users property value. The primary users associated with the managed device.
func (m *ManagedDevice) SetUsers(value []Userable)() {
    m.users = value
}
// SetUsersLoggedOn sets the usersLoggedOn property value. Indicates the last logged on users of a device. This property is read-only.
func (m *ManagedDevice) SetUsersLoggedOn(value []LoggedOnUserable)() {
    m.usersLoggedOn = value
}
// SetWiFiMacAddress sets the wiFiMacAddress property value. Wi-Fi MAC. This property is read-only.
func (m *ManagedDevice) SetWiFiMacAddress(value *string)() {
    m.wiFiMacAddress = value
}
// SetWindowsActiveMalwareCount sets the windowsActiveMalwareCount property value. Count of active malware for this windows device. This property is read-only.
func (m *ManagedDevice) SetWindowsActiveMalwareCount(value *int32)() {
    m.windowsActiveMalwareCount = value
}
// SetWindowsProtectionState sets the windowsProtectionState property value. The device protection status. This property is read-only.
func (m *ManagedDevice) SetWindowsProtectionState(value WindowsProtectionStateable)() {
    m.windowsProtectionState = value
}
// SetWindowsRemediatedMalwareCount sets the windowsRemediatedMalwareCount property value. Count of remediated malware for this windows device. This property is read-only.
func (m *ManagedDevice) SetWindowsRemediatedMalwareCount(value *int32)() {
    m.windowsRemediatedMalwareCount = value
}
