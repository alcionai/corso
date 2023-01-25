package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedRoleSummary 
type PrivilegedRoleSummary struct {
    Entity
    // The number of users that have the role assigned and the role is activated.
    elevatedCount *int32
    // The number of users that have the role assigned but the role is deactivated.
    managedCount *int32
    // true if the role activation requires MFA. false if the role activation doesn't require MFA.
    mfaEnabled *bool
    // Possible values are: ok, bad. The value depends on the ratio of (managedCount / usersCount). If the ratio is less than a predefined threshold, ok is returned. Otherwise, bad is returned.
    status *RoleSummaryStatus
    // The number of users that are assigned with the role.
    usersCount *int32
}
// NewPrivilegedRoleSummary instantiates a new privilegedRoleSummary and sets the default values.
func NewPrivilegedRoleSummary()(*PrivilegedRoleSummary) {
    m := &PrivilegedRoleSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrivilegedRoleSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegedRoleSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedRoleSummary(), nil
}
// GetElevatedCount gets the elevatedCount property value. The number of users that have the role assigned and the role is activated.
func (m *PrivilegedRoleSummary) GetElevatedCount()(*int32) {
    return m.elevatedCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedRoleSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["elevatedCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetElevatedCount(val)
        }
        return nil
    }
    res["managedCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedCount(val)
        }
        return nil
    }
    res["mfaEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMfaEnabled(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRoleSummaryStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*RoleSummaryStatus))
        }
        return nil
    }
    res["usersCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsersCount(val)
        }
        return nil
    }
    return res
}
// GetManagedCount gets the managedCount property value. The number of users that have the role assigned but the role is deactivated.
func (m *PrivilegedRoleSummary) GetManagedCount()(*int32) {
    return m.managedCount
}
// GetMfaEnabled gets the mfaEnabled property value. true if the role activation requires MFA. false if the role activation doesn't require MFA.
func (m *PrivilegedRoleSummary) GetMfaEnabled()(*bool) {
    return m.mfaEnabled
}
// GetStatus gets the status property value. Possible values are: ok, bad. The value depends on the ratio of (managedCount / usersCount). If the ratio is less than a predefined threshold, ok is returned. Otherwise, bad is returned.
func (m *PrivilegedRoleSummary) GetStatus()(*RoleSummaryStatus) {
    return m.status
}
// GetUsersCount gets the usersCount property value. The number of users that are assigned with the role.
func (m *PrivilegedRoleSummary) GetUsersCount()(*int32) {
    return m.usersCount
}
// Serialize serializes information the current object
func (m *PrivilegedRoleSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("elevatedCount", m.GetElevatedCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("managedCount", m.GetManagedCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("mfaEnabled", m.GetMfaEnabled())
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
        err = writer.WriteInt32Value("usersCount", m.GetUsersCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetElevatedCount sets the elevatedCount property value. The number of users that have the role assigned and the role is activated.
func (m *PrivilegedRoleSummary) SetElevatedCount(value *int32)() {
    m.elevatedCount = value
}
// SetManagedCount sets the managedCount property value. The number of users that have the role assigned but the role is deactivated.
func (m *PrivilegedRoleSummary) SetManagedCount(value *int32)() {
    m.managedCount = value
}
// SetMfaEnabled sets the mfaEnabled property value. true if the role activation requires MFA. false if the role activation doesn't require MFA.
func (m *PrivilegedRoleSummary) SetMfaEnabled(value *bool)() {
    m.mfaEnabled = value
}
// SetStatus sets the status property value. Possible values are: ok, bad. The value depends on the ratio of (managedCount / usersCount). If the ratio is less than a predefined threshold, ok is returned. Otherwise, bad is returned.
func (m *PrivilegedRoleSummary) SetStatus(value *RoleSummaryStatus)() {
    m.status = value
}
// SetUsersCount sets the usersCount property value. The number of users that are assigned with the role.
func (m *PrivilegedRoleSummary) SetUsersCount(value *int32)() {
    m.usersCount = value
}
