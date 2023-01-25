package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedApproval 
type PrivilegedApproval struct {
    Entity
    // The approvalDuration property
    approvalDuration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // Possible values are: pending, approved, denied, aborted, canceled.
    approvalState *ApprovalState
    // The approvalType property
    approvalType *string
    // The approverReason property
    approverReason *string
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Read-only. The role assignment request for this approval object
    request PrivilegedRoleAssignmentRequestable
    // The requestorReason property
    requestorReason *string
    // The roleId property
    roleId *string
    // The roleInfo property
    roleInfo PrivilegedRoleable
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The userId property
    userId *string
}
// NewPrivilegedApproval instantiates a new PrivilegedApproval and sets the default values.
func NewPrivilegedApproval()(*PrivilegedApproval) {
    m := &PrivilegedApproval{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrivilegedApprovalFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegedApprovalFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedApproval(), nil
}
// GetApprovalDuration gets the approvalDuration property value. The approvalDuration property
func (m *PrivilegedApproval) GetApprovalDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.approvalDuration
}
// GetApprovalState gets the approvalState property value. Possible values are: pending, approved, denied, aborted, canceled.
func (m *PrivilegedApproval) GetApprovalState()(*ApprovalState) {
    return m.approvalState
}
// GetApprovalType gets the approvalType property value. The approvalType property
func (m *PrivilegedApproval) GetApprovalType()(*string) {
    return m.approvalType
}
// GetApproverReason gets the approverReason property value. The approverReason property
func (m *PrivilegedApproval) GetApproverReason()(*string) {
    return m.approverReason
}
// GetEndDateTime gets the endDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *PrivilegedApproval) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedApproval) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["approvalDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprovalDuration(val)
        }
        return nil
    }
    res["approvalState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseApprovalState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprovalState(val.(*ApprovalState))
        }
        return nil
    }
    res["approvalType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprovalType(val)
        }
        return nil
    }
    res["approverReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApproverReason(val)
        }
        return nil
    }
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["request"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrivilegedRoleAssignmentRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequest(val.(PrivilegedRoleAssignmentRequestable))
        }
        return nil
    }
    res["requestorReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestorReason(val)
        }
        return nil
    }
    res["roleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleId(val)
        }
        return nil
    }
    res["roleInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrivilegedRoleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleInfo(val.(PrivilegedRoleable))
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
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
    return res
}
// GetRequest gets the request property value. Read-only. The role assignment request for this approval object
func (m *PrivilegedApproval) GetRequest()(PrivilegedRoleAssignmentRequestable) {
    return m.request
}
// GetRequestorReason gets the requestorReason property value. The requestorReason property
func (m *PrivilegedApproval) GetRequestorReason()(*string) {
    return m.requestorReason
}
// GetRoleId gets the roleId property value. The roleId property
func (m *PrivilegedApproval) GetRoleId()(*string) {
    return m.roleId
}
// GetRoleInfo gets the roleInfo property value. The roleInfo property
func (m *PrivilegedApproval) GetRoleInfo()(PrivilegedRoleable) {
    return m.roleInfo
}
// GetStartDateTime gets the startDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *PrivilegedApproval) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// GetUserId gets the userId property value. The userId property
func (m *PrivilegedApproval) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *PrivilegedApproval) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteISODurationValue("approvalDuration", m.GetApprovalDuration())
        if err != nil {
            return err
        }
    }
    if m.GetApprovalState() != nil {
        cast := (*m.GetApprovalState()).String()
        err = writer.WriteStringValue("approvalState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("approvalType", m.GetApprovalType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("approverReason", m.GetApproverReason())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("request", m.GetRequest())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestorReason", m.GetRequestorReason())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleId", m.GetRoleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("roleInfo", m.GetRoleInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
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
    return nil
}
// SetApprovalDuration sets the approvalDuration property value. The approvalDuration property
func (m *PrivilegedApproval) SetApprovalDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.approvalDuration = value
}
// SetApprovalState sets the approvalState property value. Possible values are: pending, approved, denied, aborted, canceled.
func (m *PrivilegedApproval) SetApprovalState(value *ApprovalState)() {
    m.approvalState = value
}
// SetApprovalType sets the approvalType property value. The approvalType property
func (m *PrivilegedApproval) SetApprovalType(value *string)() {
    m.approvalType = value
}
// SetApproverReason sets the approverReason property value. The approverReason property
func (m *PrivilegedApproval) SetApproverReason(value *string)() {
    m.approverReason = value
}
// SetEndDateTime sets the endDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *PrivilegedApproval) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetRequest sets the request property value. Read-only. The role assignment request for this approval object
func (m *PrivilegedApproval) SetRequest(value PrivilegedRoleAssignmentRequestable)() {
    m.request = value
}
// SetRequestorReason sets the requestorReason property value. The requestorReason property
func (m *PrivilegedApproval) SetRequestorReason(value *string)() {
    m.requestorReason = value
}
// SetRoleId sets the roleId property value. The roleId property
func (m *PrivilegedApproval) SetRoleId(value *string)() {
    m.roleId = value
}
// SetRoleInfo sets the roleInfo property value. The roleInfo property
func (m *PrivilegedApproval) SetRoleInfo(value PrivilegedRoleable)() {
    m.roleInfo = value
}
// SetStartDateTime sets the startDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *PrivilegedApproval) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// SetUserId sets the userId property value. The userId property
func (m *PrivilegedApproval) SetUserId(value *string)() {
    m.userId = value
}
