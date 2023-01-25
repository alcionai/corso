package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceHealthScriptRunSummary 
type DeviceHealthScriptRunSummary struct {
    Entity
    // Number of devices on which the detection script execution encountered an error and did not complete
    detectionScriptErrorDeviceCount *int32
    // Number of devices for which the detection script was not applicable
    detectionScriptNotApplicableDeviceCount *int32
    // Number of devices which have not yet run the latest version of the device health script
    detectionScriptPendingDeviceCount *int32
    // Number of devices for which the detection script found an issue
    issueDetectedDeviceCount *int32
    // Number of devices that were remediated over the last 30 days
    issueRemediatedCumulativeDeviceCount *int32
    // Number of devices for which the remediation script was able to resolve the detected issue
    issueRemediatedDeviceCount *int32
    // Number of devices for which the remediation script executed successfully but failed to resolve the detected issue
    issueReoccurredDeviceCount *int32
    // Last run time for the script across all devices
    lastScriptRunDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Number of devices for which the detection script did not find an issue and the device is healthy
    noIssueDetectedDeviceCount *int32
    // Number of devices for which the remediation script execution encountered an error and did not complete
    remediationScriptErrorDeviceCount *int32
    // Number of devices for which remediation was skipped
    remediationSkippedDeviceCount *int32
}
// NewDeviceHealthScriptRunSummary instantiates a new deviceHealthScriptRunSummary and sets the default values.
func NewDeviceHealthScriptRunSummary()(*DeviceHealthScriptRunSummary) {
    m := &DeviceHealthScriptRunSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceHealthScriptRunSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceHealthScriptRunSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceHealthScriptRunSummary(), nil
}
// GetDetectionScriptErrorDeviceCount gets the detectionScriptErrorDeviceCount property value. Number of devices on which the detection script execution encountered an error and did not complete
func (m *DeviceHealthScriptRunSummary) GetDetectionScriptErrorDeviceCount()(*int32) {
    return m.detectionScriptErrorDeviceCount
}
// GetDetectionScriptNotApplicableDeviceCount gets the detectionScriptNotApplicableDeviceCount property value. Number of devices for which the detection script was not applicable
func (m *DeviceHealthScriptRunSummary) GetDetectionScriptNotApplicableDeviceCount()(*int32) {
    return m.detectionScriptNotApplicableDeviceCount
}
// GetDetectionScriptPendingDeviceCount gets the detectionScriptPendingDeviceCount property value. Number of devices which have not yet run the latest version of the device health script
func (m *DeviceHealthScriptRunSummary) GetDetectionScriptPendingDeviceCount()(*int32) {
    return m.detectionScriptPendingDeviceCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceHealthScriptRunSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["detectionScriptErrorDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionScriptErrorDeviceCount(val)
        }
        return nil
    }
    res["detectionScriptNotApplicableDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionScriptNotApplicableDeviceCount(val)
        }
        return nil
    }
    res["detectionScriptPendingDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionScriptPendingDeviceCount(val)
        }
        return nil
    }
    res["issueDetectedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueDetectedDeviceCount(val)
        }
        return nil
    }
    res["issueRemediatedCumulativeDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueRemediatedCumulativeDeviceCount(val)
        }
        return nil
    }
    res["issueRemediatedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueRemediatedDeviceCount(val)
        }
        return nil
    }
    res["issueReoccurredDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueReoccurredDeviceCount(val)
        }
        return nil
    }
    res["lastScriptRunDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastScriptRunDateTime(val)
        }
        return nil
    }
    res["noIssueDetectedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNoIssueDetectedDeviceCount(val)
        }
        return nil
    }
    res["remediationScriptErrorDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediationScriptErrorDeviceCount(val)
        }
        return nil
    }
    res["remediationSkippedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediationSkippedDeviceCount(val)
        }
        return nil
    }
    return res
}
// GetIssueDetectedDeviceCount gets the issueDetectedDeviceCount property value. Number of devices for which the detection script found an issue
func (m *DeviceHealthScriptRunSummary) GetIssueDetectedDeviceCount()(*int32) {
    return m.issueDetectedDeviceCount
}
// GetIssueRemediatedCumulativeDeviceCount gets the issueRemediatedCumulativeDeviceCount property value. Number of devices that were remediated over the last 30 days
func (m *DeviceHealthScriptRunSummary) GetIssueRemediatedCumulativeDeviceCount()(*int32) {
    return m.issueRemediatedCumulativeDeviceCount
}
// GetIssueRemediatedDeviceCount gets the issueRemediatedDeviceCount property value. Number of devices for which the remediation script was able to resolve the detected issue
func (m *DeviceHealthScriptRunSummary) GetIssueRemediatedDeviceCount()(*int32) {
    return m.issueRemediatedDeviceCount
}
// GetIssueReoccurredDeviceCount gets the issueReoccurredDeviceCount property value. Number of devices for which the remediation script executed successfully but failed to resolve the detected issue
func (m *DeviceHealthScriptRunSummary) GetIssueReoccurredDeviceCount()(*int32) {
    return m.issueReoccurredDeviceCount
}
// GetLastScriptRunDateTime gets the lastScriptRunDateTime property value. Last run time for the script across all devices
func (m *DeviceHealthScriptRunSummary) GetLastScriptRunDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastScriptRunDateTime
}
// GetNoIssueDetectedDeviceCount gets the noIssueDetectedDeviceCount property value. Number of devices for which the detection script did not find an issue and the device is healthy
func (m *DeviceHealthScriptRunSummary) GetNoIssueDetectedDeviceCount()(*int32) {
    return m.noIssueDetectedDeviceCount
}
// GetRemediationScriptErrorDeviceCount gets the remediationScriptErrorDeviceCount property value. Number of devices for which the remediation script execution encountered an error and did not complete
func (m *DeviceHealthScriptRunSummary) GetRemediationScriptErrorDeviceCount()(*int32) {
    return m.remediationScriptErrorDeviceCount
}
// GetRemediationSkippedDeviceCount gets the remediationSkippedDeviceCount property value. Number of devices for which remediation was skipped
func (m *DeviceHealthScriptRunSummary) GetRemediationSkippedDeviceCount()(*int32) {
    return m.remediationSkippedDeviceCount
}
// Serialize serializes information the current object
func (m *DeviceHealthScriptRunSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("detectionScriptErrorDeviceCount", m.GetDetectionScriptErrorDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("detectionScriptNotApplicableDeviceCount", m.GetDetectionScriptNotApplicableDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("detectionScriptPendingDeviceCount", m.GetDetectionScriptPendingDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("issueDetectedDeviceCount", m.GetIssueDetectedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("issueRemediatedCumulativeDeviceCount", m.GetIssueRemediatedCumulativeDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("issueRemediatedDeviceCount", m.GetIssueRemediatedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("issueReoccurredDeviceCount", m.GetIssueReoccurredDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastScriptRunDateTime", m.GetLastScriptRunDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("noIssueDetectedDeviceCount", m.GetNoIssueDetectedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("remediationScriptErrorDeviceCount", m.GetRemediationScriptErrorDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("remediationSkippedDeviceCount", m.GetRemediationSkippedDeviceCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDetectionScriptErrorDeviceCount sets the detectionScriptErrorDeviceCount property value. Number of devices on which the detection script execution encountered an error and did not complete
func (m *DeviceHealthScriptRunSummary) SetDetectionScriptErrorDeviceCount(value *int32)() {
    m.detectionScriptErrorDeviceCount = value
}
// SetDetectionScriptNotApplicableDeviceCount sets the detectionScriptNotApplicableDeviceCount property value. Number of devices for which the detection script was not applicable
func (m *DeviceHealthScriptRunSummary) SetDetectionScriptNotApplicableDeviceCount(value *int32)() {
    m.detectionScriptNotApplicableDeviceCount = value
}
// SetDetectionScriptPendingDeviceCount sets the detectionScriptPendingDeviceCount property value. Number of devices which have not yet run the latest version of the device health script
func (m *DeviceHealthScriptRunSummary) SetDetectionScriptPendingDeviceCount(value *int32)() {
    m.detectionScriptPendingDeviceCount = value
}
// SetIssueDetectedDeviceCount sets the issueDetectedDeviceCount property value. Number of devices for which the detection script found an issue
func (m *DeviceHealthScriptRunSummary) SetIssueDetectedDeviceCount(value *int32)() {
    m.issueDetectedDeviceCount = value
}
// SetIssueRemediatedCumulativeDeviceCount sets the issueRemediatedCumulativeDeviceCount property value. Number of devices that were remediated over the last 30 days
func (m *DeviceHealthScriptRunSummary) SetIssueRemediatedCumulativeDeviceCount(value *int32)() {
    m.issueRemediatedCumulativeDeviceCount = value
}
// SetIssueRemediatedDeviceCount sets the issueRemediatedDeviceCount property value. Number of devices for which the remediation script was able to resolve the detected issue
func (m *DeviceHealthScriptRunSummary) SetIssueRemediatedDeviceCount(value *int32)() {
    m.issueRemediatedDeviceCount = value
}
// SetIssueReoccurredDeviceCount sets the issueReoccurredDeviceCount property value. Number of devices for which the remediation script executed successfully but failed to resolve the detected issue
func (m *DeviceHealthScriptRunSummary) SetIssueReoccurredDeviceCount(value *int32)() {
    m.issueReoccurredDeviceCount = value
}
// SetLastScriptRunDateTime sets the lastScriptRunDateTime property value. Last run time for the script across all devices
func (m *DeviceHealthScriptRunSummary) SetLastScriptRunDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastScriptRunDateTime = value
}
// SetNoIssueDetectedDeviceCount sets the noIssueDetectedDeviceCount property value. Number of devices for which the detection script did not find an issue and the device is healthy
func (m *DeviceHealthScriptRunSummary) SetNoIssueDetectedDeviceCount(value *int32)() {
    m.noIssueDetectedDeviceCount = value
}
// SetRemediationScriptErrorDeviceCount sets the remediationScriptErrorDeviceCount property value. Number of devices for which the remediation script execution encountered an error and did not complete
func (m *DeviceHealthScriptRunSummary) SetRemediationScriptErrorDeviceCount(value *int32)() {
    m.remediationScriptErrorDeviceCount = value
}
// SetRemediationSkippedDeviceCount sets the remediationSkippedDeviceCount property value. Number of devices for which remediation was skipped
func (m *DeviceHealthScriptRunSummary) SetRemediationSkippedDeviceCount(value *int32)() {
    m.remediationSkippedDeviceCount = value
}
