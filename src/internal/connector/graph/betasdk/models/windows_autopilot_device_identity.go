package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsAutopilotDeviceIdentity the windowsAutopilotDeviceIdentity resource represents a Windows Autopilot Device.
type WindowsAutopilotDeviceIdentity struct {
    Entity
    // Addressable user name.
    addressableUserName *string
    // AAD Device ID - to be deprecated
    azureActiveDirectoryDeviceId *string
    // AAD Device ID
    azureAdDeviceId *string
    // Deployment profile currently assigned to the Windows autopilot device.
    deploymentProfile WindowsAutopilotDeploymentProfileable
    // Profile set time of the Windows autopilot device.
    deploymentProfileAssignedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The deploymentProfileAssignmentDetailedStatus property
    deploymentProfileAssignmentDetailedStatus *WindowsAutopilotProfileAssignmentDetailedStatus
    // The deploymentProfileAssignmentStatus property
    deploymentProfileAssignmentStatus *WindowsAutopilotProfileAssignmentStatus
    // Surface Hub Device Account Password
    deviceAccountPassword *string
    // Surface Hub Device Account Upn
    deviceAccountUpn *string
    // Surface Hub Device Friendly Name
    deviceFriendlyName *string
    // Display Name
    displayName *string
    // The enrollmentState property
    enrollmentState *EnrollmentState
    // Group Tag of the Windows autopilot device.
    groupTag *string
    // Deployment profile intended to be assigned to the Windows autopilot device.
    intendedDeploymentProfile WindowsAutopilotDeploymentProfileable
    // Intune Last Contacted Date Time of the Windows autopilot device.
    lastContactedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Managed Device ID
    managedDeviceId *string
    // Oem manufacturer of the Windows autopilot device.
    manufacturer *string
    // Model name of the Windows autopilot device.
    model *string
    // Product Key of the Windows autopilot device.
    productKey *string
    // Purchase Order Identifier of the Windows autopilot device.
    purchaseOrderIdentifier *string
    // Device remediation status, indicating whether or not hardware has been changed for an Autopilot-registered device.
    remediationState *WindowsAutopilotDeviceRemediationState
    // RemediationState set time of Autopilot device.
    remediationStateLastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Resource Name.
    resourceName *string
    // Serial number of the Windows autopilot device.
    serialNumber *string
    // SKU Number
    skuNumber *string
    // System Family
    systemFamily *string
    // User Principal Name.
    userPrincipalName *string
}
// NewWindowsAutopilotDeviceIdentity instantiates a new windowsAutopilotDeviceIdentity and sets the default values.
func NewWindowsAutopilotDeviceIdentity()(*WindowsAutopilotDeviceIdentity) {
    m := &WindowsAutopilotDeviceIdentity{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsAutopilotDeviceIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsAutopilotDeviceIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsAutopilotDeviceIdentity(), nil
}
// GetAddressableUserName gets the addressableUserName property value. Addressable user name.
func (m *WindowsAutopilotDeviceIdentity) GetAddressableUserName()(*string) {
    return m.addressableUserName
}
// GetAzureActiveDirectoryDeviceId gets the azureActiveDirectoryDeviceId property value. AAD Device ID - to be deprecated
func (m *WindowsAutopilotDeviceIdentity) GetAzureActiveDirectoryDeviceId()(*string) {
    return m.azureActiveDirectoryDeviceId
}
// GetAzureAdDeviceId gets the azureAdDeviceId property value. AAD Device ID
func (m *WindowsAutopilotDeviceIdentity) GetAzureAdDeviceId()(*string) {
    return m.azureAdDeviceId
}
// GetDeploymentProfile gets the deploymentProfile property value. Deployment profile currently assigned to the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) GetDeploymentProfile()(WindowsAutopilotDeploymentProfileable) {
    return m.deploymentProfile
}
// GetDeploymentProfileAssignedDateTime gets the deploymentProfileAssignedDateTime property value. Profile set time of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) GetDeploymentProfileAssignedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.deploymentProfileAssignedDateTime
}
// GetDeploymentProfileAssignmentDetailedStatus gets the deploymentProfileAssignmentDetailedStatus property value. The deploymentProfileAssignmentDetailedStatus property
func (m *WindowsAutopilotDeviceIdentity) GetDeploymentProfileAssignmentDetailedStatus()(*WindowsAutopilotProfileAssignmentDetailedStatus) {
    return m.deploymentProfileAssignmentDetailedStatus
}
// GetDeploymentProfileAssignmentStatus gets the deploymentProfileAssignmentStatus property value. The deploymentProfileAssignmentStatus property
func (m *WindowsAutopilotDeviceIdentity) GetDeploymentProfileAssignmentStatus()(*WindowsAutopilotProfileAssignmentStatus) {
    return m.deploymentProfileAssignmentStatus
}
// GetDeviceAccountPassword gets the deviceAccountPassword property value. Surface Hub Device Account Password
func (m *WindowsAutopilotDeviceIdentity) GetDeviceAccountPassword()(*string) {
    return m.deviceAccountPassword
}
// GetDeviceAccountUpn gets the deviceAccountUpn property value. Surface Hub Device Account Upn
func (m *WindowsAutopilotDeviceIdentity) GetDeviceAccountUpn()(*string) {
    return m.deviceAccountUpn
}
// GetDeviceFriendlyName gets the deviceFriendlyName property value. Surface Hub Device Friendly Name
func (m *WindowsAutopilotDeviceIdentity) GetDeviceFriendlyName()(*string) {
    return m.deviceFriendlyName
}
// GetDisplayName gets the displayName property value. Display Name
func (m *WindowsAutopilotDeviceIdentity) GetDisplayName()(*string) {
    return m.displayName
}
// GetEnrollmentState gets the enrollmentState property value. The enrollmentState property
func (m *WindowsAutopilotDeviceIdentity) GetEnrollmentState()(*EnrollmentState) {
    return m.enrollmentState
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsAutopilotDeviceIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["addressableUserName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddressableUserName(val)
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
    res["azureAdDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureAdDeviceId(val)
        }
        return nil
    }
    res["deploymentProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsAutopilotDeploymentProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentProfile(val.(WindowsAutopilotDeploymentProfileable))
        }
        return nil
    }
    res["deploymentProfileAssignedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentProfileAssignedDateTime(val)
        }
        return nil
    }
    res["deploymentProfileAssignmentDetailedStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsAutopilotProfileAssignmentDetailedStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentProfileAssignmentDetailedStatus(val.(*WindowsAutopilotProfileAssignmentDetailedStatus))
        }
        return nil
    }
    res["deploymentProfileAssignmentStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsAutopilotProfileAssignmentStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentProfileAssignmentStatus(val.(*WindowsAutopilotProfileAssignmentStatus))
        }
        return nil
    }
    res["deviceAccountPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceAccountPassword(val)
        }
        return nil
    }
    res["deviceAccountUpn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceAccountUpn(val)
        }
        return nil
    }
    res["deviceFriendlyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceFriendlyName(val)
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
    res["enrollmentState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnrollmentState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentState(val.(*EnrollmentState))
        }
        return nil
    }
    res["groupTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupTag(val)
        }
        return nil
    }
    res["intendedDeploymentProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsAutopilotDeploymentProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntendedDeploymentProfile(val.(WindowsAutopilotDeploymentProfileable))
        }
        return nil
    }
    res["lastContactedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastContactedDateTime(val)
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
    res["productKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductKey(val)
        }
        return nil
    }
    res["purchaseOrderIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPurchaseOrderIdentifier(val)
        }
        return nil
    }
    res["remediationState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsAutopilotDeviceRemediationState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediationState(val.(*WindowsAutopilotDeviceRemediationState))
        }
        return nil
    }
    res["remediationStateLastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediationStateLastModifiedDateTime(val)
        }
        return nil
    }
    res["resourceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceName(val)
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
    res["skuNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkuNumber(val)
        }
        return nil
    }
    res["systemFamily"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemFamily(val)
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
// GetGroupTag gets the groupTag property value. Group Tag of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) GetGroupTag()(*string) {
    return m.groupTag
}
// GetIntendedDeploymentProfile gets the intendedDeploymentProfile property value. Deployment profile intended to be assigned to the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) GetIntendedDeploymentProfile()(WindowsAutopilotDeploymentProfileable) {
    return m.intendedDeploymentProfile
}
// GetLastContactedDateTime gets the lastContactedDateTime property value. Intune Last Contacted Date Time of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) GetLastContactedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastContactedDateTime
}
// GetManagedDeviceId gets the managedDeviceId property value. Managed Device ID
func (m *WindowsAutopilotDeviceIdentity) GetManagedDeviceId()(*string) {
    return m.managedDeviceId
}
// GetManufacturer gets the manufacturer property value. Oem manufacturer of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetModel gets the model property value. Model name of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) GetModel()(*string) {
    return m.model
}
// GetProductKey gets the productKey property value. Product Key of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) GetProductKey()(*string) {
    return m.productKey
}
// GetPurchaseOrderIdentifier gets the purchaseOrderIdentifier property value. Purchase Order Identifier of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) GetPurchaseOrderIdentifier()(*string) {
    return m.purchaseOrderIdentifier
}
// GetRemediationState gets the remediationState property value. Device remediation status, indicating whether or not hardware has been changed for an Autopilot-registered device.
func (m *WindowsAutopilotDeviceIdentity) GetRemediationState()(*WindowsAutopilotDeviceRemediationState) {
    return m.remediationState
}
// GetRemediationStateLastModifiedDateTime gets the remediationStateLastModifiedDateTime property value. RemediationState set time of Autopilot device.
func (m *WindowsAutopilotDeviceIdentity) GetRemediationStateLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.remediationStateLastModifiedDateTime
}
// GetResourceName gets the resourceName property value. Resource Name.
func (m *WindowsAutopilotDeviceIdentity) GetResourceName()(*string) {
    return m.resourceName
}
// GetSerialNumber gets the serialNumber property value. Serial number of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) GetSerialNumber()(*string) {
    return m.serialNumber
}
// GetSkuNumber gets the skuNumber property value. SKU Number
func (m *WindowsAutopilotDeviceIdentity) GetSkuNumber()(*string) {
    return m.skuNumber
}
// GetSystemFamily gets the systemFamily property value. System Family
func (m *WindowsAutopilotDeviceIdentity) GetSystemFamily()(*string) {
    return m.systemFamily
}
// GetUserPrincipalName gets the userPrincipalName property value. User Principal Name.
func (m *WindowsAutopilotDeviceIdentity) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *WindowsAutopilotDeviceIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("addressableUserName", m.GetAddressableUserName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureActiveDirectoryDeviceId", m.GetAzureActiveDirectoryDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureAdDeviceId", m.GetAzureAdDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deploymentProfile", m.GetDeploymentProfile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("deploymentProfileAssignedDateTime", m.GetDeploymentProfileAssignedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetDeploymentProfileAssignmentDetailedStatus() != nil {
        cast := (*m.GetDeploymentProfileAssignmentDetailedStatus()).String()
        err = writer.WriteStringValue("deploymentProfileAssignmentDetailedStatus", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeploymentProfileAssignmentStatus() != nil {
        cast := (*m.GetDeploymentProfileAssignmentStatus()).String()
        err = writer.WriteStringValue("deploymentProfileAssignmentStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceAccountPassword", m.GetDeviceAccountPassword())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceAccountUpn", m.GetDeviceAccountUpn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceFriendlyName", m.GetDeviceFriendlyName())
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
    if m.GetEnrollmentState() != nil {
        cast := (*m.GetEnrollmentState()).String()
        err = writer.WriteStringValue("enrollmentState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupTag", m.GetGroupTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("intendedDeploymentProfile", m.GetIntendedDeploymentProfile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastContactedDateTime", m.GetLastContactedDateTime())
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
        err = writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("productKey", m.GetProductKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("purchaseOrderIdentifier", m.GetPurchaseOrderIdentifier())
        if err != nil {
            return err
        }
    }
    if m.GetRemediationState() != nil {
        cast := (*m.GetRemediationState()).String()
        err = writer.WriteStringValue("remediationState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("remediationStateLastModifiedDateTime", m.GetRemediationStateLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceName", m.GetResourceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serialNumber", m.GetSerialNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("skuNumber", m.GetSkuNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("systemFamily", m.GetSystemFamily())
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
// SetAddressableUserName sets the addressableUserName property value. Addressable user name.
func (m *WindowsAutopilotDeviceIdentity) SetAddressableUserName(value *string)() {
    m.addressableUserName = value
}
// SetAzureActiveDirectoryDeviceId sets the azureActiveDirectoryDeviceId property value. AAD Device ID - to be deprecated
func (m *WindowsAutopilotDeviceIdentity) SetAzureActiveDirectoryDeviceId(value *string)() {
    m.azureActiveDirectoryDeviceId = value
}
// SetAzureAdDeviceId sets the azureAdDeviceId property value. AAD Device ID
func (m *WindowsAutopilotDeviceIdentity) SetAzureAdDeviceId(value *string)() {
    m.azureAdDeviceId = value
}
// SetDeploymentProfile sets the deploymentProfile property value. Deployment profile currently assigned to the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) SetDeploymentProfile(value WindowsAutopilotDeploymentProfileable)() {
    m.deploymentProfile = value
}
// SetDeploymentProfileAssignedDateTime sets the deploymentProfileAssignedDateTime property value. Profile set time of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) SetDeploymentProfileAssignedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.deploymentProfileAssignedDateTime = value
}
// SetDeploymentProfileAssignmentDetailedStatus sets the deploymentProfileAssignmentDetailedStatus property value. The deploymentProfileAssignmentDetailedStatus property
func (m *WindowsAutopilotDeviceIdentity) SetDeploymentProfileAssignmentDetailedStatus(value *WindowsAutopilotProfileAssignmentDetailedStatus)() {
    m.deploymentProfileAssignmentDetailedStatus = value
}
// SetDeploymentProfileAssignmentStatus sets the deploymentProfileAssignmentStatus property value. The deploymentProfileAssignmentStatus property
func (m *WindowsAutopilotDeviceIdentity) SetDeploymentProfileAssignmentStatus(value *WindowsAutopilotProfileAssignmentStatus)() {
    m.deploymentProfileAssignmentStatus = value
}
// SetDeviceAccountPassword sets the deviceAccountPassword property value. Surface Hub Device Account Password
func (m *WindowsAutopilotDeviceIdentity) SetDeviceAccountPassword(value *string)() {
    m.deviceAccountPassword = value
}
// SetDeviceAccountUpn sets the deviceAccountUpn property value. Surface Hub Device Account Upn
func (m *WindowsAutopilotDeviceIdentity) SetDeviceAccountUpn(value *string)() {
    m.deviceAccountUpn = value
}
// SetDeviceFriendlyName sets the deviceFriendlyName property value. Surface Hub Device Friendly Name
func (m *WindowsAutopilotDeviceIdentity) SetDeviceFriendlyName(value *string)() {
    m.deviceFriendlyName = value
}
// SetDisplayName sets the displayName property value. Display Name
func (m *WindowsAutopilotDeviceIdentity) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEnrollmentState sets the enrollmentState property value. The enrollmentState property
func (m *WindowsAutopilotDeviceIdentity) SetEnrollmentState(value *EnrollmentState)() {
    m.enrollmentState = value
}
// SetGroupTag sets the groupTag property value. Group Tag of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) SetGroupTag(value *string)() {
    m.groupTag = value
}
// SetIntendedDeploymentProfile sets the intendedDeploymentProfile property value. Deployment profile intended to be assigned to the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) SetIntendedDeploymentProfile(value WindowsAutopilotDeploymentProfileable)() {
    m.intendedDeploymentProfile = value
}
// SetLastContactedDateTime sets the lastContactedDateTime property value. Intune Last Contacted Date Time of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) SetLastContactedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastContactedDateTime = value
}
// SetManagedDeviceId sets the managedDeviceId property value. Managed Device ID
func (m *WindowsAutopilotDeviceIdentity) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
// SetManufacturer sets the manufacturer property value. Oem manufacturer of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetModel sets the model property value. Model name of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) SetModel(value *string)() {
    m.model = value
}
// SetProductKey sets the productKey property value. Product Key of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) SetProductKey(value *string)() {
    m.productKey = value
}
// SetPurchaseOrderIdentifier sets the purchaseOrderIdentifier property value. Purchase Order Identifier of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) SetPurchaseOrderIdentifier(value *string)() {
    m.purchaseOrderIdentifier = value
}
// SetRemediationState sets the remediationState property value. Device remediation status, indicating whether or not hardware has been changed for an Autopilot-registered device.
func (m *WindowsAutopilotDeviceIdentity) SetRemediationState(value *WindowsAutopilotDeviceRemediationState)() {
    m.remediationState = value
}
// SetRemediationStateLastModifiedDateTime sets the remediationStateLastModifiedDateTime property value. RemediationState set time of Autopilot device.
func (m *WindowsAutopilotDeviceIdentity) SetRemediationStateLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.remediationStateLastModifiedDateTime = value
}
// SetResourceName sets the resourceName property value. Resource Name.
func (m *WindowsAutopilotDeviceIdentity) SetResourceName(value *string)() {
    m.resourceName = value
}
// SetSerialNumber sets the serialNumber property value. Serial number of the Windows autopilot device.
func (m *WindowsAutopilotDeviceIdentity) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
// SetSkuNumber sets the skuNumber property value. SKU Number
func (m *WindowsAutopilotDeviceIdentity) SetSkuNumber(value *string)() {
    m.skuNumber = value
}
// SetSystemFamily sets the systemFamily property value. System Family
func (m *WindowsAutopilotDeviceIdentity) SetSystemFamily(value *string)() {
    m.systemFamily = value
}
// SetUserPrincipalName sets the userPrincipalName property value. User Principal Name.
func (m *WindowsAutopilotDeviceIdentity) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
