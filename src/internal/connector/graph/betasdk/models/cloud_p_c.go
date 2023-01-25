package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPC provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type CloudPC struct {
    Entity
    // The Azure Active Directory (Azure AD) device ID of the Cloud PC.
    aadDeviceId *string
    // The connectivity health check result of a Cloud PC, including the updated timestamp and whether the Cloud PC is able to be connected or not.
    connectivityResult CloudPcConnectivityResultable
    // The diskEncryptionState property
    diskEncryptionState *CloudPcDiskEncryptionState
    // The display name of the Cloud PC.
    displayName *string
    // The date and time when the grace period ends and reprovisioning/deprovisioning happens. Required only if the status is inGracePeriod. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    gracePeriodEndDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Name of the OS image that's on the Cloud PC.
    imageDisplayName *string
    // The last login result of the Cloud PC. For example, { 'time': '2014-01-01T00:00:00Z'}.
    lastLoginResult CloudPcLoginResultable
    // The last modified date and time of the Cloud PC. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The last remote action result of the enterprise Cloud PCs. The supported remote actions are: Reboot, Rename, Reprovision, Restore, and Troubleshoot.
    lastRemoteActionResult CloudPcRemoteActionResultable
    // The Intune device ID of the Cloud PC.
    managedDeviceId *string
    // The Intune device name of the Cloud PC.
    managedDeviceName *string
    // The Azure network connection that is applied during the provisioning of Cloud PCs.
    onPremisesConnectionName *string
    // The version of the operating system (OS) to provision on Cloud PCs. Possible values are: windows10, windows11, and unknownFutureValue.
    osVersion *CloudPcOperatingSystem
    // The results of every partner agent's installation status on Cloud PC.
    partnerAgentInstallResults []CloudPcPartnerAgentInstallResultable
    // The provisioning policy ID of the Cloud PC.
    provisioningPolicyId *string
    // The provisioning policy that is applied during the provisioning of Cloud PCs.
    provisioningPolicyName *string
    // The provisioningType property
    provisioningType *CloudPcProvisioningType
    // The service plan ID of the Cloud PC.
    servicePlanId *string
    // The service plan name of the Cloud PC.
    servicePlanName *string
    // The service plan type of the Cloud PC.
    servicePlanType *CloudPcServicePlanType
    // The status property
    status *CloudPcStatus
    // The details of the Cloud PC status.
    statusDetails CloudPcStatusDetailsable
    // The account type of the user on provisioned Cloud PCs. Possible values are: standardUser, administrator, and unknownFutureValue.
    userAccountType *CloudPcUserAccountType
    // The user principal name (UPN) of the user assigned to the Cloud PC.
    userPrincipalName *string
}
// NewCloudPC instantiates a new cloudPC and sets the default values.
func NewCloudPC()(*CloudPC) {
    m := &CloudPC{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPCFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPCFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPC(), nil
}
// GetAadDeviceId gets the aadDeviceId property value. The Azure Active Directory (Azure AD) device ID of the Cloud PC.
func (m *CloudPC) GetAadDeviceId()(*string) {
    return m.aadDeviceId
}
// GetConnectivityResult gets the connectivityResult property value. The connectivity health check result of a Cloud PC, including the updated timestamp and whether the Cloud PC is able to be connected or not.
func (m *CloudPC) GetConnectivityResult()(CloudPcConnectivityResultable) {
    return m.connectivityResult
}
// GetDiskEncryptionState gets the diskEncryptionState property value. The diskEncryptionState property
func (m *CloudPC) GetDiskEncryptionState()(*CloudPcDiskEncryptionState) {
    return m.diskEncryptionState
}
// GetDisplayName gets the displayName property value. The display name of the Cloud PC.
func (m *CloudPC) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPC) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["aadDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAadDeviceId(val)
        }
        return nil
    }
    res["connectivityResult"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcConnectivityResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectivityResult(val.(CloudPcConnectivityResultable))
        }
        return nil
    }
    res["diskEncryptionState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcDiskEncryptionState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiskEncryptionState(val.(*CloudPcDiskEncryptionState))
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
    res["gracePeriodEndDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGracePeriodEndDateTime(val)
        }
        return nil
    }
    res["imageDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImageDisplayName(val)
        }
        return nil
    }
    res["lastLoginResult"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcLoginResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastLoginResult(val.(CloudPcLoginResultable))
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
    res["lastRemoteActionResult"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcRemoteActionResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRemoteActionResult(val.(CloudPcRemoteActionResultable))
        }
        return nil
    }
    res["managedDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceId(val)
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
    res["onPremisesConnectionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesConnectionName(val)
        }
        return nil
    }
    res["osVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcOperatingSystem)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val.(*CloudPcOperatingSystem))
        }
        return nil
    }
    res["partnerAgentInstallResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudPcPartnerAgentInstallResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudPcPartnerAgentInstallResultable, len(val))
            for i, v := range val {
                res[i] = v.(CloudPcPartnerAgentInstallResultable)
            }
            m.SetPartnerAgentInstallResults(res)
        }
        return nil
    }
    res["provisioningPolicyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningPolicyId(val)
        }
        return nil
    }
    res["provisioningPolicyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningPolicyName(val)
        }
        return nil
    }
    res["provisioningType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcProvisioningType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningType(val.(*CloudPcProvisioningType))
        }
        return nil
    }
    res["servicePlanId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePlanId(val)
        }
        return nil
    }
    res["servicePlanName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePlanName(val)
        }
        return nil
    }
    res["servicePlanType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcServicePlanType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePlanType(val.(*CloudPcServicePlanType))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*CloudPcStatus))
        }
        return nil
    }
    res["statusDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcStatusDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusDetails(val.(CloudPcStatusDetailsable))
        }
        return nil
    }
    res["userAccountType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcUserAccountType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserAccountType(val.(*CloudPcUserAccountType))
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
    return res
}
// GetGracePeriodEndDateTime gets the gracePeriodEndDateTime property value. The date and time when the grace period ends and reprovisioning/deprovisioning happens. Required only if the status is inGracePeriod. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *CloudPC) GetGracePeriodEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.gracePeriodEndDateTime
}
// GetImageDisplayName gets the imageDisplayName property value. Name of the OS image that's on the Cloud PC.
func (m *CloudPC) GetImageDisplayName()(*string) {
    return m.imageDisplayName
}
// GetLastLoginResult gets the lastLoginResult property value. The last login result of the Cloud PC. For example, { 'time': '2014-01-01T00:00:00Z'}.
func (m *CloudPC) GetLastLoginResult()(CloudPcLoginResultable) {
    return m.lastLoginResult
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The last modified date and time of the Cloud PC. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *CloudPC) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetLastRemoteActionResult gets the lastRemoteActionResult property value. The last remote action result of the enterprise Cloud PCs. The supported remote actions are: Reboot, Rename, Reprovision, Restore, and Troubleshoot.
func (m *CloudPC) GetLastRemoteActionResult()(CloudPcRemoteActionResultable) {
    return m.lastRemoteActionResult
}
// GetManagedDeviceId gets the managedDeviceId property value. The Intune device ID of the Cloud PC.
func (m *CloudPC) GetManagedDeviceId()(*string) {
    return m.managedDeviceId
}
// GetManagedDeviceName gets the managedDeviceName property value. The Intune device name of the Cloud PC.
func (m *CloudPC) GetManagedDeviceName()(*string) {
    return m.managedDeviceName
}
// GetOnPremisesConnectionName gets the onPremisesConnectionName property value. The Azure network connection that is applied during the provisioning of Cloud PCs.
func (m *CloudPC) GetOnPremisesConnectionName()(*string) {
    return m.onPremisesConnectionName
}
// GetOsVersion gets the osVersion property value. The version of the operating system (OS) to provision on Cloud PCs. Possible values are: windows10, windows11, and unknownFutureValue.
func (m *CloudPC) GetOsVersion()(*CloudPcOperatingSystem) {
    return m.osVersion
}
// GetPartnerAgentInstallResults gets the partnerAgentInstallResults property value. The results of every partner agent's installation status on Cloud PC.
func (m *CloudPC) GetPartnerAgentInstallResults()([]CloudPcPartnerAgentInstallResultable) {
    return m.partnerAgentInstallResults
}
// GetProvisioningPolicyId gets the provisioningPolicyId property value. The provisioning policy ID of the Cloud PC.
func (m *CloudPC) GetProvisioningPolicyId()(*string) {
    return m.provisioningPolicyId
}
// GetProvisioningPolicyName gets the provisioningPolicyName property value. The provisioning policy that is applied during the provisioning of Cloud PCs.
func (m *CloudPC) GetProvisioningPolicyName()(*string) {
    return m.provisioningPolicyName
}
// GetProvisioningType gets the provisioningType property value. The provisioningType property
func (m *CloudPC) GetProvisioningType()(*CloudPcProvisioningType) {
    return m.provisioningType
}
// GetServicePlanId gets the servicePlanId property value. The service plan ID of the Cloud PC.
func (m *CloudPC) GetServicePlanId()(*string) {
    return m.servicePlanId
}
// GetServicePlanName gets the servicePlanName property value. The service plan name of the Cloud PC.
func (m *CloudPC) GetServicePlanName()(*string) {
    return m.servicePlanName
}
// GetServicePlanType gets the servicePlanType property value. The service plan type of the Cloud PC.
func (m *CloudPC) GetServicePlanType()(*CloudPcServicePlanType) {
    return m.servicePlanType
}
// GetStatus gets the status property value. The status property
func (m *CloudPC) GetStatus()(*CloudPcStatus) {
    return m.status
}
// GetStatusDetails gets the statusDetails property value. The details of the Cloud PC status.
func (m *CloudPC) GetStatusDetails()(CloudPcStatusDetailsable) {
    return m.statusDetails
}
// GetUserAccountType gets the userAccountType property value. The account type of the user on provisioned Cloud PCs. Possible values are: standardUser, administrator, and unknownFutureValue.
func (m *CloudPC) GetUserAccountType()(*CloudPcUserAccountType) {
    return m.userAccountType
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name (UPN) of the user assigned to the Cloud PC.
func (m *CloudPC) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *CloudPC) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("aadDeviceId", m.GetAadDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("connectivityResult", m.GetConnectivityResult())
        if err != nil {
            return err
        }
    }
    if m.GetDiskEncryptionState() != nil {
        cast := (*m.GetDiskEncryptionState()).String()
        err = writer.WriteStringValue("diskEncryptionState", &cast)
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
    {
        err = writer.WriteTimeValue("gracePeriodEndDateTime", m.GetGracePeriodEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("imageDisplayName", m.GetImageDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastLoginResult", m.GetLastLoginResult())
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
    {
        err = writer.WriteObjectValue("lastRemoteActionResult", m.GetLastRemoteActionResult())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceId", m.GetManagedDeviceId())
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
    {
        err = writer.WriteStringValue("onPremisesConnectionName", m.GetOnPremisesConnectionName())
        if err != nil {
            return err
        }
    }
    if m.GetOsVersion() != nil {
        cast := (*m.GetOsVersion()).String()
        err = writer.WriteStringValue("osVersion", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPartnerAgentInstallResults() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPartnerAgentInstallResults()))
        for i, v := range m.GetPartnerAgentInstallResults() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("partnerAgentInstallResults", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("provisioningPolicyId", m.GetProvisioningPolicyId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("provisioningPolicyName", m.GetProvisioningPolicyName())
        if err != nil {
            return err
        }
    }
    if m.GetProvisioningType() != nil {
        cast := (*m.GetProvisioningType()).String()
        err = writer.WriteStringValue("provisioningType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePlanId", m.GetServicePlanId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePlanName", m.GetServicePlanName())
        if err != nil {
            return err
        }
    }
    if m.GetServicePlanType() != nil {
        cast := (*m.GetServicePlanType()).String()
        err = writer.WriteStringValue("servicePlanType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("statusDetails", m.GetStatusDetails())
        if err != nil {
            return err
        }
    }
    if m.GetUserAccountType() != nil {
        cast := (*m.GetUserAccountType()).String()
        err = writer.WriteStringValue("userAccountType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAadDeviceId sets the aadDeviceId property value. The Azure Active Directory (Azure AD) device ID of the Cloud PC.
func (m *CloudPC) SetAadDeviceId(value *string)() {
    m.aadDeviceId = value
}
// SetConnectivityResult sets the connectivityResult property value. The connectivity health check result of a Cloud PC, including the updated timestamp and whether the Cloud PC is able to be connected or not.
func (m *CloudPC) SetConnectivityResult(value CloudPcConnectivityResultable)() {
    m.connectivityResult = value
}
// SetDiskEncryptionState sets the diskEncryptionState property value. The diskEncryptionState property
func (m *CloudPC) SetDiskEncryptionState(value *CloudPcDiskEncryptionState)() {
    m.diskEncryptionState = value
}
// SetDisplayName sets the displayName property value. The display name of the Cloud PC.
func (m *CloudPC) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGracePeriodEndDateTime sets the gracePeriodEndDateTime property value. The date and time when the grace period ends and reprovisioning/deprovisioning happens. Required only if the status is inGracePeriod. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *CloudPC) SetGracePeriodEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.gracePeriodEndDateTime = value
}
// SetImageDisplayName sets the imageDisplayName property value. Name of the OS image that's on the Cloud PC.
func (m *CloudPC) SetImageDisplayName(value *string)() {
    m.imageDisplayName = value
}
// SetLastLoginResult sets the lastLoginResult property value. The last login result of the Cloud PC. For example, { 'time': '2014-01-01T00:00:00Z'}.
func (m *CloudPC) SetLastLoginResult(value CloudPcLoginResultable)() {
    m.lastLoginResult = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The last modified date and time of the Cloud PC. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *CloudPC) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetLastRemoteActionResult sets the lastRemoteActionResult property value. The last remote action result of the enterprise Cloud PCs. The supported remote actions are: Reboot, Rename, Reprovision, Restore, and Troubleshoot.
func (m *CloudPC) SetLastRemoteActionResult(value CloudPcRemoteActionResultable)() {
    m.lastRemoteActionResult = value
}
// SetManagedDeviceId sets the managedDeviceId property value. The Intune device ID of the Cloud PC.
func (m *CloudPC) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
// SetManagedDeviceName sets the managedDeviceName property value. The Intune device name of the Cloud PC.
func (m *CloudPC) SetManagedDeviceName(value *string)() {
    m.managedDeviceName = value
}
// SetOnPremisesConnectionName sets the onPremisesConnectionName property value. The Azure network connection that is applied during the provisioning of Cloud PCs.
func (m *CloudPC) SetOnPremisesConnectionName(value *string)() {
    m.onPremisesConnectionName = value
}
// SetOsVersion sets the osVersion property value. The version of the operating system (OS) to provision on Cloud PCs. Possible values are: windows10, windows11, and unknownFutureValue.
func (m *CloudPC) SetOsVersion(value *CloudPcOperatingSystem)() {
    m.osVersion = value
}
// SetPartnerAgentInstallResults sets the partnerAgentInstallResults property value. The results of every partner agent's installation status on Cloud PC.
func (m *CloudPC) SetPartnerAgentInstallResults(value []CloudPcPartnerAgentInstallResultable)() {
    m.partnerAgentInstallResults = value
}
// SetProvisioningPolicyId sets the provisioningPolicyId property value. The provisioning policy ID of the Cloud PC.
func (m *CloudPC) SetProvisioningPolicyId(value *string)() {
    m.provisioningPolicyId = value
}
// SetProvisioningPolicyName sets the provisioningPolicyName property value. The provisioning policy that is applied during the provisioning of Cloud PCs.
func (m *CloudPC) SetProvisioningPolicyName(value *string)() {
    m.provisioningPolicyName = value
}
// SetProvisioningType sets the provisioningType property value. The provisioningType property
func (m *CloudPC) SetProvisioningType(value *CloudPcProvisioningType)() {
    m.provisioningType = value
}
// SetServicePlanId sets the servicePlanId property value. The service plan ID of the Cloud PC.
func (m *CloudPC) SetServicePlanId(value *string)() {
    m.servicePlanId = value
}
// SetServicePlanName sets the servicePlanName property value. The service plan name of the Cloud PC.
func (m *CloudPC) SetServicePlanName(value *string)() {
    m.servicePlanName = value
}
// SetServicePlanType sets the servicePlanType property value. The service plan type of the Cloud PC.
func (m *CloudPC) SetServicePlanType(value *CloudPcServicePlanType)() {
    m.servicePlanType = value
}
// SetStatus sets the status property value. The status property
func (m *CloudPC) SetStatus(value *CloudPcStatus)() {
    m.status = value
}
// SetStatusDetails sets the statusDetails property value. The details of the Cloud PC status.
func (m *CloudPC) SetStatusDetails(value CloudPcStatusDetailsable)() {
    m.statusDetails = value
}
// SetUserAccountType sets the userAccountType property value. The account type of the user on provisioned Cloud PCs. Possible values are: standardUser, administrator, and unknownFutureValue.
func (m *CloudPC) SetUserAccountType(value *CloudPcUserAccountType)() {
    m.userAccountType = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name (UPN) of the user assigned to the Cloud PC.
func (m *CloudPC) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
