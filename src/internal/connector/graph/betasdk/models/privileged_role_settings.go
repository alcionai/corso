package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedRoleSettings 
type PrivilegedRoleSettings struct {
    Entity
    // true if the approval is required when activate the role. false if the approval is not required when activate the role.
    approvalOnElevation *bool
    // List of Approval ids, if approval is required for activation.
    approverIds []string
    // The duration when the role is activated.
    elevationDuration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // true if mfaOnElevation is configurable. false if mfaOnElevation is not configurable.
    isMfaOnElevationConfigurable *bool
    // Internal used only.
    lastGlobalAdmin *bool
    // Maximal duration for the activated role.
    maxElavationDuration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // true if MFA is required to activate the role. false if MFA is not required to activate the role.
    mfaOnElevation *bool
    // Minimal duration for the activated role.
    minElevationDuration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // true if send notification to the end user when the role is activated. false if do not send notification when the role is activated.
    notificationToUserOnElevation *bool
    // true if the ticketing information is required when activate the role. false if the ticketing information is not required when activate the role.
    ticketingInfoOnElevation *bool
}
// NewPrivilegedRoleSettings instantiates a new privilegedRoleSettings and sets the default values.
func NewPrivilegedRoleSettings()(*PrivilegedRoleSettings) {
    m := &PrivilegedRoleSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrivilegedRoleSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegedRoleSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedRoleSettings(), nil
}
// GetApprovalOnElevation gets the approvalOnElevation property value. true if the approval is required when activate the role. false if the approval is not required when activate the role.
func (m *PrivilegedRoleSettings) GetApprovalOnElevation()(*bool) {
    return m.approvalOnElevation
}
// GetApproverIds gets the approverIds property value. List of Approval ids, if approval is required for activation.
func (m *PrivilegedRoleSettings) GetApproverIds()([]string) {
    return m.approverIds
}
// GetElevationDuration gets the elevationDuration property value. The duration when the role is activated.
func (m *PrivilegedRoleSettings) GetElevationDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.elevationDuration
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedRoleSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["approvalOnElevation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprovalOnElevation(val)
        }
        return nil
    }
    res["approverIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetApproverIds(res)
        }
        return nil
    }
    res["elevationDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetElevationDuration(val)
        }
        return nil
    }
    res["isMfaOnElevationConfigurable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMfaOnElevationConfigurable(val)
        }
        return nil
    }
    res["lastGlobalAdmin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastGlobalAdmin(val)
        }
        return nil
    }
    res["maxElavationDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxElavationDuration(val)
        }
        return nil
    }
    res["mfaOnElevation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMfaOnElevation(val)
        }
        return nil
    }
    res["minElevationDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinElevationDuration(val)
        }
        return nil
    }
    res["notificationToUserOnElevation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationToUserOnElevation(val)
        }
        return nil
    }
    res["ticketingInfoOnElevation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTicketingInfoOnElevation(val)
        }
        return nil
    }
    return res
}
// GetIsMfaOnElevationConfigurable gets the isMfaOnElevationConfigurable property value. true if mfaOnElevation is configurable. false if mfaOnElevation is not configurable.
func (m *PrivilegedRoleSettings) GetIsMfaOnElevationConfigurable()(*bool) {
    return m.isMfaOnElevationConfigurable
}
// GetLastGlobalAdmin gets the lastGlobalAdmin property value. Internal used only.
func (m *PrivilegedRoleSettings) GetLastGlobalAdmin()(*bool) {
    return m.lastGlobalAdmin
}
// GetMaxElavationDuration gets the maxElavationDuration property value. Maximal duration for the activated role.
func (m *PrivilegedRoleSettings) GetMaxElavationDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.maxElavationDuration
}
// GetMfaOnElevation gets the mfaOnElevation property value. true if MFA is required to activate the role. false if MFA is not required to activate the role.
func (m *PrivilegedRoleSettings) GetMfaOnElevation()(*bool) {
    return m.mfaOnElevation
}
// GetMinElevationDuration gets the minElevationDuration property value. Minimal duration for the activated role.
func (m *PrivilegedRoleSettings) GetMinElevationDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.minElevationDuration
}
// GetNotificationToUserOnElevation gets the notificationToUserOnElevation property value. true if send notification to the end user when the role is activated. false if do not send notification when the role is activated.
func (m *PrivilegedRoleSettings) GetNotificationToUserOnElevation()(*bool) {
    return m.notificationToUserOnElevation
}
// GetTicketingInfoOnElevation gets the ticketingInfoOnElevation property value. true if the ticketing information is required when activate the role. false if the ticketing information is not required when activate the role.
func (m *PrivilegedRoleSettings) GetTicketingInfoOnElevation()(*bool) {
    return m.ticketingInfoOnElevation
}
// Serialize serializes information the current object
func (m *PrivilegedRoleSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("approvalOnElevation", m.GetApprovalOnElevation())
        if err != nil {
            return err
        }
    }
    if m.GetApproverIds() != nil {
        err = writer.WriteCollectionOfStringValues("approverIds", m.GetApproverIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("elevationDuration", m.GetElevationDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMfaOnElevationConfigurable", m.GetIsMfaOnElevationConfigurable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("lastGlobalAdmin", m.GetLastGlobalAdmin())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("maxElavationDuration", m.GetMaxElavationDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("mfaOnElevation", m.GetMfaOnElevation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("minElevationDuration", m.GetMinElevationDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("notificationToUserOnElevation", m.GetNotificationToUserOnElevation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("ticketingInfoOnElevation", m.GetTicketingInfoOnElevation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApprovalOnElevation sets the approvalOnElevation property value. true if the approval is required when activate the role. false if the approval is not required when activate the role.
func (m *PrivilegedRoleSettings) SetApprovalOnElevation(value *bool)() {
    m.approvalOnElevation = value
}
// SetApproverIds sets the approverIds property value. List of Approval ids, if approval is required for activation.
func (m *PrivilegedRoleSettings) SetApproverIds(value []string)() {
    m.approverIds = value
}
// SetElevationDuration sets the elevationDuration property value. The duration when the role is activated.
func (m *PrivilegedRoleSettings) SetElevationDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.elevationDuration = value
}
// SetIsMfaOnElevationConfigurable sets the isMfaOnElevationConfigurable property value. true if mfaOnElevation is configurable. false if mfaOnElevation is not configurable.
func (m *PrivilegedRoleSettings) SetIsMfaOnElevationConfigurable(value *bool)() {
    m.isMfaOnElevationConfigurable = value
}
// SetLastGlobalAdmin sets the lastGlobalAdmin property value. Internal used only.
func (m *PrivilegedRoleSettings) SetLastGlobalAdmin(value *bool)() {
    m.lastGlobalAdmin = value
}
// SetMaxElavationDuration sets the maxElavationDuration property value. Maximal duration for the activated role.
func (m *PrivilegedRoleSettings) SetMaxElavationDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.maxElavationDuration = value
}
// SetMfaOnElevation sets the mfaOnElevation property value. true if MFA is required to activate the role. false if MFA is not required to activate the role.
func (m *PrivilegedRoleSettings) SetMfaOnElevation(value *bool)() {
    m.mfaOnElevation = value
}
// SetMinElevationDuration sets the minElevationDuration property value. Minimal duration for the activated role.
func (m *PrivilegedRoleSettings) SetMinElevationDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.minElevationDuration = value
}
// SetNotificationToUserOnElevation sets the notificationToUserOnElevation property value. true if send notification to the end user when the role is activated. false if do not send notification when the role is activated.
func (m *PrivilegedRoleSettings) SetNotificationToUserOnElevation(value *bool)() {
    m.notificationToUserOnElevation = value
}
// SetTicketingInfoOnElevation sets the ticketingInfoOnElevation property value. true if the ticketing information is required when activate the role. false if the ticketing information is not required when activate the role.
func (m *PrivilegedRoleSettings) SetTicketingInfoOnElevation(value *bool)() {
    m.ticketingInfoOnElevation = value
}
