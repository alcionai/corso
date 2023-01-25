package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsUpdateState provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type WindowsUpdateState struct {
    Entity
    // Device display name.
    deviceDisplayName *string
    // The id of the device.
    deviceId *string
    // The current feature update version of the device.
    featureUpdateVersion *string
    // The date time that the Windows Update Agent did a successful scan.
    lastScanDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Last date time that the device sync with with Microsoft Intune.
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The Quality Update Version of the device.
    qualityUpdateVersion *string
    // Windows update for business configuration device states
    status *WindowsUpdateStatus
    // The id of the user.
    userId *string
    // User principal name.
    userPrincipalName *string
}
// NewWindowsUpdateState instantiates a new windowsUpdateState and sets the default values.
func NewWindowsUpdateState()(*WindowsUpdateState) {
    m := &WindowsUpdateState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsUpdateStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsUpdateStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsUpdateState(), nil
}
// GetDeviceDisplayName gets the deviceDisplayName property value. Device display name.
func (m *WindowsUpdateState) GetDeviceDisplayName()(*string) {
    return m.deviceDisplayName
}
// GetDeviceId gets the deviceId property value. The id of the device.
func (m *WindowsUpdateState) GetDeviceId()(*string) {
    return m.deviceId
}
// GetFeatureUpdateVersion gets the featureUpdateVersion property value. The current feature update version of the device.
func (m *WindowsUpdateState) GetFeatureUpdateVersion()(*string) {
    return m.featureUpdateVersion
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsUpdateState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceDisplayName(val)
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
    res["featureUpdateVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeatureUpdateVersion(val)
        }
        return nil
    }
    res["lastScanDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastScanDateTime(val)
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
    res["qualityUpdateVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQualityUpdateVersion(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsUpdateStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*WindowsUpdateStatus))
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
    return res
}
// GetLastScanDateTime gets the lastScanDateTime property value. The date time that the Windows Update Agent did a successful scan.
func (m *WindowsUpdateState) GetLastScanDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastScanDateTime
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. Last date time that the device sync with with Microsoft Intune.
func (m *WindowsUpdateState) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSyncDateTime
}
// GetQualityUpdateVersion gets the qualityUpdateVersion property value. The Quality Update Version of the device.
func (m *WindowsUpdateState) GetQualityUpdateVersion()(*string) {
    return m.qualityUpdateVersion
}
// GetStatus gets the status property value. Windows update for business configuration device states
func (m *WindowsUpdateState) GetStatus()(*WindowsUpdateStatus) {
    return m.status
}
// GetUserId gets the userId property value. The id of the user.
func (m *WindowsUpdateState) GetUserId()(*string) {
    return m.userId
}
// GetUserPrincipalName gets the userPrincipalName property value. User principal name.
func (m *WindowsUpdateState) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *WindowsUpdateState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("deviceDisplayName", m.GetDeviceDisplayName())
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
        err = writer.WriteStringValue("featureUpdateVersion", m.GetFeatureUpdateVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastScanDateTime", m.GetLastScanDateTime())
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
        err = writer.WriteStringValue("qualityUpdateVersion", m.GetQualityUpdateVersion())
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
        err = writer.WriteStringValue("userId", m.GetUserId())
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
// SetDeviceDisplayName sets the deviceDisplayName property value. Device display name.
func (m *WindowsUpdateState) SetDeviceDisplayName(value *string)() {
    m.deviceDisplayName = value
}
// SetDeviceId sets the deviceId property value. The id of the device.
func (m *WindowsUpdateState) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetFeatureUpdateVersion sets the featureUpdateVersion property value. The current feature update version of the device.
func (m *WindowsUpdateState) SetFeatureUpdateVersion(value *string)() {
    m.featureUpdateVersion = value
}
// SetLastScanDateTime sets the lastScanDateTime property value. The date time that the Windows Update Agent did a successful scan.
func (m *WindowsUpdateState) SetLastScanDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastScanDateTime = value
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. Last date time that the device sync with with Microsoft Intune.
func (m *WindowsUpdateState) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
// SetQualityUpdateVersion sets the qualityUpdateVersion property value. The Quality Update Version of the device.
func (m *WindowsUpdateState) SetQualityUpdateVersion(value *string)() {
    m.qualityUpdateVersion = value
}
// SetStatus sets the status property value. Windows update for business configuration device states
func (m *WindowsUpdateState) SetStatus(value *WindowsUpdateStatus)() {
    m.status = value
}
// SetUserId sets the userId property value. The id of the user.
func (m *WindowsUpdateState) SetUserId(value *string)() {
    m.userId = value
}
// SetUserPrincipalName sets the userPrincipalName property value. User principal name.
func (m *WindowsUpdateState) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
