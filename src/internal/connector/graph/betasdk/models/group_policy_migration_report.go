package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyMigrationReport the Group Policy migration report.
type GroupPolicyMigrationReport struct {
    Entity
    // The date and time at which the GroupPolicyMigrationReport was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The name of Group Policy Object from the GPO Xml Content
    displayName *string
    // The date and time at which the GroupPolicyMigrationReport was created.
    groupPolicyCreatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date and time at which the GroupPolicyMigrationReport was last modified.
    groupPolicyLastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The Group Policy Object GUID from GPO Xml content
    groupPolicyObjectId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // A list of group policy settings to MDM/Intune mappings.
    groupPolicySettingMappings []GroupPolicySettingMappingable
    // The date and time at which the GroupPolicyMigrationReport was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Indicates if the Group Policy Object file is covered and ready for Intune migration.
    migrationReadiness *GroupPolicyMigrationReadiness
    // The distinguished name of the OU.
    ouDistinguishedName *string
    // The list of scope tags for the configuration.
    roleScopeTagIds []string
    // The number of Group Policy Settings supported by Intune.
    supportedSettingsCount *int32
    // The Percentage of Group Policy Settings supported by Intune.
    supportedSettingsPercent *int32
    // The Targeted in AD property from GPO Xml Content
    targetedInActiveDirectory *bool
    // The total number of Group Policy Settings from GPO file.
    totalSettingsCount *int32
    // A list of unsupported group policy extensions inside the Group Policy Object.
    unsupportedGroupPolicyExtensions []UnsupportedGroupPolicyExtensionable
}
// NewGroupPolicyMigrationReport instantiates a new groupPolicyMigrationReport and sets the default values.
func NewGroupPolicyMigrationReport()(*GroupPolicyMigrationReport) {
    m := &GroupPolicyMigrationReport{
        Entity: *NewEntity(),
    }
    return m
}
// CreateGroupPolicyMigrationReportFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyMigrationReportFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyMigrationReport(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time at which the GroupPolicyMigrationReport was created.
func (m *GroupPolicyMigrationReport) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDisplayName gets the displayName property value. The name of Group Policy Object from the GPO Xml Content
func (m *GroupPolicyMigrationReport) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyMigrationReport) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
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
    res["groupPolicyCreatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupPolicyCreatedDateTime(val)
        }
        return nil
    }
    res["groupPolicyLastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupPolicyLastModifiedDateTime(val)
        }
        return nil
    }
    res["groupPolicyObjectId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupPolicyObjectId(val)
        }
        return nil
    }
    res["groupPolicySettingMappings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupPolicySettingMappingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupPolicySettingMappingable, len(val))
            for i, v := range val {
                res[i] = v.(GroupPolicySettingMappingable)
            }
            m.SetGroupPolicySettingMappings(res)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["migrationReadiness"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicyMigrationReadiness)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMigrationReadiness(val.(*GroupPolicyMigrationReadiness))
        }
        return nil
    }
    res["ouDistinguishedName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOuDistinguishedName(val)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["supportedSettingsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportedSettingsCount(val)
        }
        return nil
    }
    res["supportedSettingsPercent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportedSettingsPercent(val)
        }
        return nil
    }
    res["targetedInActiveDirectory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetedInActiveDirectory(val)
        }
        return nil
    }
    res["totalSettingsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalSettingsCount(val)
        }
        return nil
    }
    res["unsupportedGroupPolicyExtensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnsupportedGroupPolicyExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnsupportedGroupPolicyExtensionable, len(val))
            for i, v := range val {
                res[i] = v.(UnsupportedGroupPolicyExtensionable)
            }
            m.SetUnsupportedGroupPolicyExtensions(res)
        }
        return nil
    }
    return res
}
// GetGroupPolicyCreatedDateTime gets the groupPolicyCreatedDateTime property value. The date and time at which the GroupPolicyMigrationReport was created.
func (m *GroupPolicyMigrationReport) GetGroupPolicyCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.groupPolicyCreatedDateTime
}
// GetGroupPolicyLastModifiedDateTime gets the groupPolicyLastModifiedDateTime property value. The date and time at which the GroupPolicyMigrationReport was last modified.
func (m *GroupPolicyMigrationReport) GetGroupPolicyLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.groupPolicyLastModifiedDateTime
}
// GetGroupPolicyObjectId gets the groupPolicyObjectId property value. The Group Policy Object GUID from GPO Xml content
func (m *GroupPolicyMigrationReport) GetGroupPolicyObjectId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.groupPolicyObjectId
}
// GetGroupPolicySettingMappings gets the groupPolicySettingMappings property value. A list of group policy settings to MDM/Intune mappings.
func (m *GroupPolicyMigrationReport) GetGroupPolicySettingMappings()([]GroupPolicySettingMappingable) {
    return m.groupPolicySettingMappings
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time at which the GroupPolicyMigrationReport was last modified.
func (m *GroupPolicyMigrationReport) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetMigrationReadiness gets the migrationReadiness property value. Indicates if the Group Policy Object file is covered and ready for Intune migration.
func (m *GroupPolicyMigrationReport) GetMigrationReadiness()(*GroupPolicyMigrationReadiness) {
    return m.migrationReadiness
}
// GetOuDistinguishedName gets the ouDistinguishedName property value. The distinguished name of the OU.
func (m *GroupPolicyMigrationReport) GetOuDistinguishedName()(*string) {
    return m.ouDistinguishedName
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. The list of scope tags for the configuration.
func (m *GroupPolicyMigrationReport) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// GetSupportedSettingsCount gets the supportedSettingsCount property value. The number of Group Policy Settings supported by Intune.
func (m *GroupPolicyMigrationReport) GetSupportedSettingsCount()(*int32) {
    return m.supportedSettingsCount
}
// GetSupportedSettingsPercent gets the supportedSettingsPercent property value. The Percentage of Group Policy Settings supported by Intune.
func (m *GroupPolicyMigrationReport) GetSupportedSettingsPercent()(*int32) {
    return m.supportedSettingsPercent
}
// GetTargetedInActiveDirectory gets the targetedInActiveDirectory property value. The Targeted in AD property from GPO Xml Content
func (m *GroupPolicyMigrationReport) GetTargetedInActiveDirectory()(*bool) {
    return m.targetedInActiveDirectory
}
// GetTotalSettingsCount gets the totalSettingsCount property value. The total number of Group Policy Settings from GPO file.
func (m *GroupPolicyMigrationReport) GetTotalSettingsCount()(*int32) {
    return m.totalSettingsCount
}
// GetUnsupportedGroupPolicyExtensions gets the unsupportedGroupPolicyExtensions property value. A list of unsupported group policy extensions inside the Group Policy Object.
func (m *GroupPolicyMigrationReport) GetUnsupportedGroupPolicyExtensions()([]UnsupportedGroupPolicyExtensionable) {
    return m.unsupportedGroupPolicyExtensions
}
// Serialize serializes information the current object
func (m *GroupPolicyMigrationReport) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
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
        err = writer.WriteTimeValue("groupPolicyCreatedDateTime", m.GetGroupPolicyCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("groupPolicyLastModifiedDateTime", m.GetGroupPolicyLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("groupPolicyObjectId", m.GetGroupPolicyObjectId())
        if err != nil {
            return err
        }
    }
    if m.GetGroupPolicySettingMappings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroupPolicySettingMappings()))
        for i, v := range m.GetGroupPolicySettingMappings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("groupPolicySettingMappings", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetMigrationReadiness() != nil {
        cast := (*m.GetMigrationReadiness()).String()
        err = writer.WriteStringValue("migrationReadiness", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ouDistinguishedName", m.GetOuDistinguishedName())
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("supportedSettingsCount", m.GetSupportedSettingsCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("supportedSettingsPercent", m.GetSupportedSettingsPercent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("targetedInActiveDirectory", m.GetTargetedInActiveDirectory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalSettingsCount", m.GetTotalSettingsCount())
        if err != nil {
            return err
        }
    }
    if m.GetUnsupportedGroupPolicyExtensions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUnsupportedGroupPolicyExtensions()))
        for i, v := range m.GetUnsupportedGroupPolicyExtensions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("unsupportedGroupPolicyExtensions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time at which the GroupPolicyMigrationReport was created.
func (m *GroupPolicyMigrationReport) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDisplayName sets the displayName property value. The name of Group Policy Object from the GPO Xml Content
func (m *GroupPolicyMigrationReport) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGroupPolicyCreatedDateTime sets the groupPolicyCreatedDateTime property value. The date and time at which the GroupPolicyMigrationReport was created.
func (m *GroupPolicyMigrationReport) SetGroupPolicyCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.groupPolicyCreatedDateTime = value
}
// SetGroupPolicyLastModifiedDateTime sets the groupPolicyLastModifiedDateTime property value. The date and time at which the GroupPolicyMigrationReport was last modified.
func (m *GroupPolicyMigrationReport) SetGroupPolicyLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.groupPolicyLastModifiedDateTime = value
}
// SetGroupPolicyObjectId sets the groupPolicyObjectId property value. The Group Policy Object GUID from GPO Xml content
func (m *GroupPolicyMigrationReport) SetGroupPolicyObjectId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.groupPolicyObjectId = value
}
// SetGroupPolicySettingMappings sets the groupPolicySettingMappings property value. A list of group policy settings to MDM/Intune mappings.
func (m *GroupPolicyMigrationReport) SetGroupPolicySettingMappings(value []GroupPolicySettingMappingable)() {
    m.groupPolicySettingMappings = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time at which the GroupPolicyMigrationReport was last modified.
func (m *GroupPolicyMigrationReport) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetMigrationReadiness sets the migrationReadiness property value. Indicates if the Group Policy Object file is covered and ready for Intune migration.
func (m *GroupPolicyMigrationReport) SetMigrationReadiness(value *GroupPolicyMigrationReadiness)() {
    m.migrationReadiness = value
}
// SetOuDistinguishedName sets the ouDistinguishedName property value. The distinguished name of the OU.
func (m *GroupPolicyMigrationReport) SetOuDistinguishedName(value *string)() {
    m.ouDistinguishedName = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. The list of scope tags for the configuration.
func (m *GroupPolicyMigrationReport) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// SetSupportedSettingsCount sets the supportedSettingsCount property value. The number of Group Policy Settings supported by Intune.
func (m *GroupPolicyMigrationReport) SetSupportedSettingsCount(value *int32)() {
    m.supportedSettingsCount = value
}
// SetSupportedSettingsPercent sets the supportedSettingsPercent property value. The Percentage of Group Policy Settings supported by Intune.
func (m *GroupPolicyMigrationReport) SetSupportedSettingsPercent(value *int32)() {
    m.supportedSettingsPercent = value
}
// SetTargetedInActiveDirectory sets the targetedInActiveDirectory property value. The Targeted in AD property from GPO Xml Content
func (m *GroupPolicyMigrationReport) SetTargetedInActiveDirectory(value *bool)() {
    m.targetedInActiveDirectory = value
}
// SetTotalSettingsCount sets the totalSettingsCount property value. The total number of Group Policy Settings from GPO file.
func (m *GroupPolicyMigrationReport) SetTotalSettingsCount(value *int32)() {
    m.totalSettingsCount = value
}
// SetUnsupportedGroupPolicyExtensions sets the unsupportedGroupPolicyExtensions property value. A list of unsupported group policy extensions inside the Group Policy Object.
func (m *GroupPolicyMigrationReport) SetUnsupportedGroupPolicyExtensions(value []UnsupportedGroupPolicyExtensionable)() {
    m.unsupportedGroupPolicyExtensions = value
}
