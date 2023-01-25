package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedRoleAssignment provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type PrivilegedRoleAssignment struct {
    Entity
    // The UTC DateTime when the temporary privileged role assignment will be expired. For permanent role assignment, the value is null.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // true if the role assignment is activated. false if the role assignment is deactivated.
    isElevated *bool
    // Result message set by the service.
    resultMessage *string
    // Role identifier. In GUID string format.
    roleId *string
    // Read-only. Nullable. The associated role information.
    roleInfo PrivilegedRoleable
    // User identifier. In GUID string format.
    userId *string
}
// NewPrivilegedRoleAssignment instantiates a new privilegedRoleAssignment and sets the default values.
func NewPrivilegedRoleAssignment()(*PrivilegedRoleAssignment) {
    m := &PrivilegedRoleAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrivilegedRoleAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegedRoleAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedRoleAssignment(), nil
}
// GetExpirationDateTime gets the expirationDateTime property value. The UTC DateTime when the temporary privileged role assignment will be expired. For permanent role assignment, the value is null.
func (m *PrivilegedRoleAssignment) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedRoleAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["isElevated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsElevated(val)
        }
        return nil
    }
    res["resultMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResultMessage(val)
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
// GetIsElevated gets the isElevated property value. true if the role assignment is activated. false if the role assignment is deactivated.
func (m *PrivilegedRoleAssignment) GetIsElevated()(*bool) {
    return m.isElevated
}
// GetResultMessage gets the resultMessage property value. Result message set by the service.
func (m *PrivilegedRoleAssignment) GetResultMessage()(*string) {
    return m.resultMessage
}
// GetRoleId gets the roleId property value. Role identifier. In GUID string format.
func (m *PrivilegedRoleAssignment) GetRoleId()(*string) {
    return m.roleId
}
// GetRoleInfo gets the roleInfo property value. Read-only. Nullable. The associated role information.
func (m *PrivilegedRoleAssignment) GetRoleInfo()(PrivilegedRoleable) {
    return m.roleInfo
}
// GetUserId gets the userId property value. User identifier. In GUID string format.
func (m *PrivilegedRoleAssignment) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *PrivilegedRoleAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isElevated", m.GetIsElevated())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resultMessage", m.GetResultMessage())
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
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExpirationDateTime sets the expirationDateTime property value. The UTC DateTime when the temporary privileged role assignment will be expired. For permanent role assignment, the value is null.
func (m *PrivilegedRoleAssignment) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetIsElevated sets the isElevated property value. true if the role assignment is activated. false if the role assignment is deactivated.
func (m *PrivilegedRoleAssignment) SetIsElevated(value *bool)() {
    m.isElevated = value
}
// SetResultMessage sets the resultMessage property value. Result message set by the service.
func (m *PrivilegedRoleAssignment) SetResultMessage(value *string)() {
    m.resultMessage = value
}
// SetRoleId sets the roleId property value. Role identifier. In GUID string format.
func (m *PrivilegedRoleAssignment) SetRoleId(value *string)() {
    m.roleId = value
}
// SetRoleInfo sets the roleInfo property value. Read-only. Nullable. The associated role information.
func (m *PrivilegedRoleAssignment) SetRoleInfo(value PrivilegedRoleable)() {
    m.roleInfo = value
}
// SetUserId sets the userId property value. User identifier. In GUID string format.
func (m *PrivilegedRoleAssignment) SetUserId(value *string)() {
    m.userId = value
}
