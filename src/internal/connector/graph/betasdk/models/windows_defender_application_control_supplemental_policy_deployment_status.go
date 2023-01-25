package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus contains properties for the deployment state of a WindowsDefenderApplicationControl supplemental policy for a device.
type WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus struct {
    Entity
    // Enum values for the various WindowsDefenderApplicationControl supplemental policy deployment statuses.
    deploymentStatus *WindowsDefenderApplicationControlSupplementalPolicyStatuses
    // Device ID.
    deviceId *string
    // Device name.
    deviceName *string
    // Last sync date time.
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Windows OS Version Description.
    osDescription *string
    // Windows OS Version.
    osVersion *string
    // The navigation link to the WindowsDefenderApplicationControl supplemental policy.
    policy WindowsDefenderApplicationControlSupplementalPolicyable
    // Human readable version of the WindowsDefenderApplicationControl supplemental policy.
    policyVersion *string
    // The name of the user of this device.
    userName *string
    // User Principal Name.
    userPrincipalName *string
}
// NewWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus instantiates a new windowsDefenderApplicationControlSupplementalPolicyDeploymentStatus and sets the default values.
func NewWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus()(*WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) {
    m := &WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus(), nil
}
// GetDeploymentStatus gets the deploymentStatus property value. Enum values for the various WindowsDefenderApplicationControl supplemental policy deployment statuses.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetDeploymentStatus()(*WindowsDefenderApplicationControlSupplementalPolicyStatuses) {
    return m.deploymentStatus
}
// GetDeviceId gets the deviceId property value. Device ID.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetDeviceId()(*string) {
    return m.deviceId
}
// GetDeviceName gets the deviceName property value. Device name.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetDeviceName()(*string) {
    return m.deviceName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deploymentStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsDefenderApplicationControlSupplementalPolicyStatuses)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentStatus(val.(*WindowsDefenderApplicationControlSupplementalPolicyStatuses))
        }
        return nil
    }
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
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
    res["osDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsDescription(val)
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
    res["policy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsDefenderApplicationControlSupplementalPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicy(val.(WindowsDefenderApplicationControlSupplementalPolicyable))
        }
        return nil
    }
    res["policyVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyVersion(val)
        }
        return nil
    }
    res["userName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserName(val)
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
// GetLastSyncDateTime gets the lastSyncDateTime property value. Last sync date time.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSyncDateTime
}
// GetOsDescription gets the osDescription property value. Windows OS Version Description.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetOsDescription()(*string) {
    return m.osDescription
}
// GetOsVersion gets the osVersion property value. Windows OS Version.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetOsVersion()(*string) {
    return m.osVersion
}
// GetPolicy gets the policy property value. The navigation link to the WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetPolicy()(WindowsDefenderApplicationControlSupplementalPolicyable) {
    return m.policy
}
// GetPolicyVersion gets the policyVersion property value. Human readable version of the WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetPolicyVersion()(*string) {
    return m.policyVersion
}
// GetUserName gets the userName property value. The name of the user of this device.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetUserName()(*string) {
    return m.userName
}
// GetUserPrincipalName gets the userPrincipalName property value. User Principal Name.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDeploymentStatus() != nil {
        cast := (*m.GetDeploymentStatus()).String()
        err = writer.WriteStringValue("deploymentStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSyncDateTime", m.GetLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osDescription", m.GetOsDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osVersion", m.GetOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("policy", m.GetPolicy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("policyVersion", m.GetPolicyVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userName", m.GetUserName())
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
// SetDeploymentStatus sets the deploymentStatus property value. Enum values for the various WindowsDefenderApplicationControl supplemental policy deployment statuses.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetDeploymentStatus(value *WindowsDefenderApplicationControlSupplementalPolicyStatuses)() {
    m.deploymentStatus = value
}
// SetDeviceId sets the deviceId property value. Device ID.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetDeviceName sets the deviceName property value. Device name.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. Last sync date time.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
// SetOsDescription sets the osDescription property value. Windows OS Version Description.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetOsDescription(value *string)() {
    m.osDescription = value
}
// SetOsVersion sets the osVersion property value. Windows OS Version.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetOsVersion(value *string)() {
    m.osVersion = value
}
// SetPolicy sets the policy property value. The navigation link to the WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetPolicy(value WindowsDefenderApplicationControlSupplementalPolicyable)() {
    m.policy = value
}
// SetPolicyVersion sets the policyVersion property value. Human readable version of the WindowsDefenderApplicationControl supplemental policy.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetPolicyVersion(value *string)() {
    m.policyVersion = value
}
// SetUserName sets the userName property value. The name of the user of this device.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetUserName(value *string)() {
    m.userName = value
}
// SetUserPrincipalName sets the userPrincipalName property value. User Principal Name.
func (m *WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
