package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GovernanceRoleSetting 
type GovernanceRoleSetting struct {
    Entity
    // The rule settings that are evaluated when an administrator tries to add an eligible role assignment.
    adminEligibleSettings []GovernanceRuleSettingable
    // The rule settings that are evaluated when an administrator tries to add a direct member role assignment.
    adminMemberSettings []GovernanceRuleSettingable
    // Read-only. Indicate if the roleSetting is a default roleSetting
    isDefault *bool
    // Read-only. The display name of the administrator who last updated the roleSetting.
    lastUpdatedBy *string
    // Read-only. The time when the role setting was last updated. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Read-only. The associated resource for this role setting.
    resource GovernanceResourceable
    // Required. The id of the resource that the role setting is associated with.
    resourceId *string
    // Read-only. The role definition that is enforced with this role setting.
    roleDefinition GovernanceRoleDefinitionable
    // Required. The id of the role definition that the role setting is associated with.
    roleDefinitionId *string
    // The rule settings that are evaluated when a user tries to add an eligible role assignment. The setting is not supported for now.
    userEligibleSettings []GovernanceRuleSettingable
    // The rule settings that are evaluated when a user tries to activate his role assignment.
    userMemberSettings []GovernanceRuleSettingable
}
// NewGovernanceRoleSetting instantiates a new governanceRoleSetting and sets the default values.
func NewGovernanceRoleSetting()(*GovernanceRoleSetting) {
    m := &GovernanceRoleSetting{
        Entity: *NewEntity(),
    }
    return m
}
// CreateGovernanceRoleSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGovernanceRoleSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGovernanceRoleSetting(), nil
}
// GetAdminEligibleSettings gets the adminEligibleSettings property value. The rule settings that are evaluated when an administrator tries to add an eligible role assignment.
func (m *GovernanceRoleSetting) GetAdminEligibleSettings()([]GovernanceRuleSettingable) {
    return m.adminEligibleSettings
}
// GetAdminMemberSettings gets the adminMemberSettings property value. The rule settings that are evaluated when an administrator tries to add a direct member role assignment.
func (m *GovernanceRoleSetting) GetAdminMemberSettings()([]GovernanceRuleSettingable) {
    return m.adminMemberSettings
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GovernanceRoleSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["adminEligibleSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGovernanceRuleSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRuleSettingable, len(val))
            for i, v := range val {
                res[i] = v.(GovernanceRuleSettingable)
            }
            m.SetAdminEligibleSettings(res)
        }
        return nil
    }
    res["adminMemberSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGovernanceRuleSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRuleSettingable, len(val))
            for i, v := range val {
                res[i] = v.(GovernanceRuleSettingable)
            }
            m.SetAdminMemberSettings(res)
        }
        return nil
    }
    res["isDefault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefault(val)
        }
        return nil
    }
    res["lastUpdatedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedBy(val)
        }
        return nil
    }
    res["lastUpdatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedDateTime(val)
        }
        return nil
    }
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGovernanceResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(GovernanceResourceable))
        }
        return nil
    }
    res["resourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
        }
        return nil
    }
    res["roleDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGovernanceRoleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleDefinition(val.(GovernanceRoleDefinitionable))
        }
        return nil
    }
    res["roleDefinitionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleDefinitionId(val)
        }
        return nil
    }
    res["userEligibleSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGovernanceRuleSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRuleSettingable, len(val))
            for i, v := range val {
                res[i] = v.(GovernanceRuleSettingable)
            }
            m.SetUserEligibleSettings(res)
        }
        return nil
    }
    res["userMemberSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGovernanceRuleSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernanceRuleSettingable, len(val))
            for i, v := range val {
                res[i] = v.(GovernanceRuleSettingable)
            }
            m.SetUserMemberSettings(res)
        }
        return nil
    }
    return res
}
// GetIsDefault gets the isDefault property value. Read-only. Indicate if the roleSetting is a default roleSetting
func (m *GovernanceRoleSetting) GetIsDefault()(*bool) {
    return m.isDefault
}
// GetLastUpdatedBy gets the lastUpdatedBy property value. Read-only. The display name of the administrator who last updated the roleSetting.
func (m *GovernanceRoleSetting) GetLastUpdatedBy()(*string) {
    return m.lastUpdatedBy
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. Read-only. The time when the role setting was last updated. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *GovernanceRoleSetting) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdatedDateTime
}
// GetResource gets the resource property value. Read-only. The associated resource for this role setting.
func (m *GovernanceRoleSetting) GetResource()(GovernanceResourceable) {
    return m.resource
}
// GetResourceId gets the resourceId property value. Required. The id of the resource that the role setting is associated with.
func (m *GovernanceRoleSetting) GetResourceId()(*string) {
    return m.resourceId
}
// GetRoleDefinition gets the roleDefinition property value. Read-only. The role definition that is enforced with this role setting.
func (m *GovernanceRoleSetting) GetRoleDefinition()(GovernanceRoleDefinitionable) {
    return m.roleDefinition
}
// GetRoleDefinitionId gets the roleDefinitionId property value. Required. The id of the role definition that the role setting is associated with.
func (m *GovernanceRoleSetting) GetRoleDefinitionId()(*string) {
    return m.roleDefinitionId
}
// GetUserEligibleSettings gets the userEligibleSettings property value. The rule settings that are evaluated when a user tries to add an eligible role assignment. The setting is not supported for now.
func (m *GovernanceRoleSetting) GetUserEligibleSettings()([]GovernanceRuleSettingable) {
    return m.userEligibleSettings
}
// GetUserMemberSettings gets the userMemberSettings property value. The rule settings that are evaluated when a user tries to activate his role assignment.
func (m *GovernanceRoleSetting) GetUserMemberSettings()([]GovernanceRuleSettingable) {
    return m.userMemberSettings
}
// Serialize serializes information the current object
func (m *GovernanceRoleSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdminEligibleSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAdminEligibleSettings()))
        for i, v := range m.GetAdminEligibleSettings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("adminEligibleSettings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAdminMemberSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAdminMemberSettings()))
        for i, v := range m.GetAdminMemberSettings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("adminMemberSettings", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastUpdatedBy", m.GetLastUpdatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceId", m.GetResourceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("roleDefinition", m.GetRoleDefinition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleDefinitionId", m.GetRoleDefinitionId())
        if err != nil {
            return err
        }
    }
    if m.GetUserEligibleSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserEligibleSettings()))
        for i, v := range m.GetUserEligibleSettings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userEligibleSettings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserMemberSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserMemberSettings()))
        for i, v := range m.GetUserMemberSettings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userMemberSettings", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdminEligibleSettings sets the adminEligibleSettings property value. The rule settings that are evaluated when an administrator tries to add an eligible role assignment.
func (m *GovernanceRoleSetting) SetAdminEligibleSettings(value []GovernanceRuleSettingable)() {
    m.adminEligibleSettings = value
}
// SetAdminMemberSettings sets the adminMemberSettings property value. The rule settings that are evaluated when an administrator tries to add a direct member role assignment.
func (m *GovernanceRoleSetting) SetAdminMemberSettings(value []GovernanceRuleSettingable)() {
    m.adminMemberSettings = value
}
// SetIsDefault sets the isDefault property value. Read-only. Indicate if the roleSetting is a default roleSetting
func (m *GovernanceRoleSetting) SetIsDefault(value *bool)() {
    m.isDefault = value
}
// SetLastUpdatedBy sets the lastUpdatedBy property value. Read-only. The display name of the administrator who last updated the roleSetting.
func (m *GovernanceRoleSetting) SetLastUpdatedBy(value *string)() {
    m.lastUpdatedBy = value
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. Read-only. The time when the role setting was last updated. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *GovernanceRoleSetting) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
// SetResource sets the resource property value. Read-only. The associated resource for this role setting.
func (m *GovernanceRoleSetting) SetResource(value GovernanceResourceable)() {
    m.resource = value
}
// SetResourceId sets the resourceId property value. Required. The id of the resource that the role setting is associated with.
func (m *GovernanceRoleSetting) SetResourceId(value *string)() {
    m.resourceId = value
}
// SetRoleDefinition sets the roleDefinition property value. Read-only. The role definition that is enforced with this role setting.
func (m *GovernanceRoleSetting) SetRoleDefinition(value GovernanceRoleDefinitionable)() {
    m.roleDefinition = value
}
// SetRoleDefinitionId sets the roleDefinitionId property value. Required. The id of the role definition that the role setting is associated with.
func (m *GovernanceRoleSetting) SetRoleDefinitionId(value *string)() {
    m.roleDefinitionId = value
}
// SetUserEligibleSettings sets the userEligibleSettings property value. The rule settings that are evaluated when a user tries to add an eligible role assignment. The setting is not supported for now.
func (m *GovernanceRoleSetting) SetUserEligibleSettings(value []GovernanceRuleSettingable)() {
    m.userEligibleSettings = value
}
// SetUserMemberSettings sets the userMemberSettings property value. The rule settings that are evaluated when a user tries to activate his role assignment.
func (m *GovernanceRoleSetting) SetUserMemberSettings(value []GovernanceRuleSettingable)() {
    m.userMemberSettings = value
}
