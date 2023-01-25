package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementAutopilotPolicyStatusDetail policy status detail item contained by an autopilot event.
type DeviceManagementAutopilotPolicyStatusDetail struct {
    Entity
    // The complianceStatus property
    complianceStatus *DeviceManagementAutopilotPolicyComplianceStatus
    // The friendly name of the policy.
    displayName *string
    // The errorode associated with the compliance or enforcement status of the policy. Error code for enforcement status takes precedence if it exists.
    errorCode *int32
    // Timestamp of the reported policy status
    lastReportedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The policyType property
    policyType *DeviceManagementAutopilotPolicyType
    // Indicates if this prolicy was tracked as part of the autopilot bootstrap enrollment sync session
    trackedOnEnrollmentStatus *bool
}
// NewDeviceManagementAutopilotPolicyStatusDetail instantiates a new deviceManagementAutopilotPolicyStatusDetail and sets the default values.
func NewDeviceManagementAutopilotPolicyStatusDetail()(*DeviceManagementAutopilotPolicyStatusDetail) {
    m := &DeviceManagementAutopilotPolicyStatusDetail{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementAutopilotPolicyStatusDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementAutopilotPolicyStatusDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementAutopilotPolicyStatusDetail(), nil
}
// GetComplianceStatus gets the complianceStatus property value. The complianceStatus property
func (m *DeviceManagementAutopilotPolicyStatusDetail) GetComplianceStatus()(*DeviceManagementAutopilotPolicyComplianceStatus) {
    return m.complianceStatus
}
// GetDisplayName gets the displayName property value. The friendly name of the policy.
func (m *DeviceManagementAutopilotPolicyStatusDetail) GetDisplayName()(*string) {
    return m.displayName
}
// GetErrorCode gets the errorCode property value. The errorode associated with the compliance or enforcement status of the policy. Error code for enforcement status takes precedence if it exists.
func (m *DeviceManagementAutopilotPolicyStatusDetail) GetErrorCode()(*int32) {
    return m.errorCode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementAutopilotPolicyStatusDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["complianceStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementAutopilotPolicyComplianceStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceStatus(val.(*DeviceManagementAutopilotPolicyComplianceStatus))
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
    res["errorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["lastReportedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastReportedDateTime(val)
        }
        return nil
    }
    res["policyType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementAutopilotPolicyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyType(val.(*DeviceManagementAutopilotPolicyType))
        }
        return nil
    }
    res["trackedOnEnrollmentStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrackedOnEnrollmentStatus(val)
        }
        return nil
    }
    return res
}
// GetLastReportedDateTime gets the lastReportedDateTime property value. Timestamp of the reported policy status
func (m *DeviceManagementAutopilotPolicyStatusDetail) GetLastReportedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastReportedDateTime
}
// GetPolicyType gets the policyType property value. The policyType property
func (m *DeviceManagementAutopilotPolicyStatusDetail) GetPolicyType()(*DeviceManagementAutopilotPolicyType) {
    return m.policyType
}
// GetTrackedOnEnrollmentStatus gets the trackedOnEnrollmentStatus property value. Indicates if this prolicy was tracked as part of the autopilot bootstrap enrollment sync session
func (m *DeviceManagementAutopilotPolicyStatusDetail) GetTrackedOnEnrollmentStatus()(*bool) {
    return m.trackedOnEnrollmentStatus
}
// Serialize serializes information the current object
func (m *DeviceManagementAutopilotPolicyStatusDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetComplianceStatus() != nil {
        cast := (*m.GetComplianceStatus()).String()
        err = writer.WriteStringValue("complianceStatus", &cast)
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
        err = writer.WriteInt32Value("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastReportedDateTime", m.GetLastReportedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPolicyType() != nil {
        cast := (*m.GetPolicyType()).String()
        err = writer.WriteStringValue("policyType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("trackedOnEnrollmentStatus", m.GetTrackedOnEnrollmentStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetComplianceStatus sets the complianceStatus property value. The complianceStatus property
func (m *DeviceManagementAutopilotPolicyStatusDetail) SetComplianceStatus(value *DeviceManagementAutopilotPolicyComplianceStatus)() {
    m.complianceStatus = value
}
// SetDisplayName sets the displayName property value. The friendly name of the policy.
func (m *DeviceManagementAutopilotPolicyStatusDetail) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetErrorCode sets the errorCode property value. The errorode associated with the compliance or enforcement status of the policy. Error code for enforcement status takes precedence if it exists.
func (m *DeviceManagementAutopilotPolicyStatusDetail) SetErrorCode(value *int32)() {
    m.errorCode = value
}
// SetLastReportedDateTime sets the lastReportedDateTime property value. Timestamp of the reported policy status
func (m *DeviceManagementAutopilotPolicyStatusDetail) SetLastReportedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastReportedDateTime = value
}
// SetPolicyType sets the policyType property value. The policyType property
func (m *DeviceManagementAutopilotPolicyStatusDetail) SetPolicyType(value *DeviceManagementAutopilotPolicyType)() {
    m.policyType = value
}
// SetTrackedOnEnrollmentStatus sets the trackedOnEnrollmentStatus property value. Indicates if this prolicy was tracked as part of the autopilot bootstrap enrollment sync session
func (m *DeviceManagementAutopilotPolicyStatusDetail) SetTrackedOnEnrollmentStatus(value *bool)() {
    m.trackedOnEnrollmentStatus = value
}
