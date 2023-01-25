package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationPolicyTemplate device Management Configuration Policy Template
type DeviceManagementConfigurationPolicyTemplate struct {
    Entity
    // Allow unmanaged setting templates
    allowUnmanagedSettings *bool
    // Template base identifier
    baseId *string
    // Template description
    description *string
    // Template display name
    displayName *string
    // Description of template version
    displayVersion *string
    // Describes current lifecycle state of a template
    lifecycleState *DeviceManagementTemplateLifecycleState
    // Supported platform types.
    platforms *DeviceManagementConfigurationPlatforms
    // Number of setting templates. Valid values 0 to 2147483647. This property is read-only.
    settingTemplateCount *int32
    // Setting templates
    settingTemplates []DeviceManagementConfigurationSettingTemplateable
    // Describes which technology this setting can be deployed with
    technologies *DeviceManagementConfigurationTechnologies
    // Describes the TemplateFamily for the Template entity
    templateFamily *DeviceManagementConfigurationTemplateFamily
    // Template version. Valid values 1 to 2147483647. This property is read-only.
    version *int32
}
// NewDeviceManagementConfigurationPolicyTemplate instantiates a new deviceManagementConfigurationPolicyTemplate and sets the default values.
func NewDeviceManagementConfigurationPolicyTemplate()(*DeviceManagementConfigurationPolicyTemplate) {
    m := &DeviceManagementConfigurationPolicyTemplate{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementConfigurationPolicyTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationPolicyTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationPolicyTemplate(), nil
}
// GetAllowUnmanagedSettings gets the allowUnmanagedSettings property value. Allow unmanaged setting templates
func (m *DeviceManagementConfigurationPolicyTemplate) GetAllowUnmanagedSettings()(*bool) {
    return m.allowUnmanagedSettings
}
// GetBaseId gets the baseId property value. Template base identifier
func (m *DeviceManagementConfigurationPolicyTemplate) GetBaseId()(*string) {
    return m.baseId
}
// GetDescription gets the description property value. Template description
func (m *DeviceManagementConfigurationPolicyTemplate) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Template display name
func (m *DeviceManagementConfigurationPolicyTemplate) GetDisplayName()(*string) {
    return m.displayName
}
// GetDisplayVersion gets the displayVersion property value. Description of template version
func (m *DeviceManagementConfigurationPolicyTemplate) GetDisplayVersion()(*string) {
    return m.displayVersion
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationPolicyTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowUnmanagedSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowUnmanagedSettings(val)
        }
        return nil
    }
    res["baseId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBaseId(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["displayVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayVersion(val)
        }
        return nil
    }
    res["lifecycleState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementTemplateLifecycleState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLifecycleState(val.(*DeviceManagementTemplateLifecycleState))
        }
        return nil
    }
    res["platforms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationPlatforms)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatforms(val.(*DeviceManagementConfigurationPlatforms))
        }
        return nil
    }
    res["settingTemplateCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingTemplateCount(val)
        }
        return nil
    }
    res["settingTemplates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationSettingTemplateable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationSettingTemplateable)
            }
            m.SetSettingTemplates(res)
        }
        return nil
    }
    res["technologies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationTechnologies)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTechnologies(val.(*DeviceManagementConfigurationTechnologies))
        }
        return nil
    }
    res["templateFamily"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationTemplateFamily)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateFamily(val.(*DeviceManagementConfigurationTemplateFamily))
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetLifecycleState gets the lifecycleState property value. Describes current lifecycle state of a template
func (m *DeviceManagementConfigurationPolicyTemplate) GetLifecycleState()(*DeviceManagementTemplateLifecycleState) {
    return m.lifecycleState
}
// GetPlatforms gets the platforms property value. Supported platform types.
func (m *DeviceManagementConfigurationPolicyTemplate) GetPlatforms()(*DeviceManagementConfigurationPlatforms) {
    return m.platforms
}
// GetSettingTemplateCount gets the settingTemplateCount property value. Number of setting templates. Valid values 0 to 2147483647. This property is read-only.
func (m *DeviceManagementConfigurationPolicyTemplate) GetSettingTemplateCount()(*int32) {
    return m.settingTemplateCount
}
// GetSettingTemplates gets the settingTemplates property value. Setting templates
func (m *DeviceManagementConfigurationPolicyTemplate) GetSettingTemplates()([]DeviceManagementConfigurationSettingTemplateable) {
    return m.settingTemplates
}
// GetTechnologies gets the technologies property value. Describes which technology this setting can be deployed with
func (m *DeviceManagementConfigurationPolicyTemplate) GetTechnologies()(*DeviceManagementConfigurationTechnologies) {
    return m.technologies
}
// GetTemplateFamily gets the templateFamily property value. Describes the TemplateFamily for the Template entity
func (m *DeviceManagementConfigurationPolicyTemplate) GetTemplateFamily()(*DeviceManagementConfigurationTemplateFamily) {
    return m.templateFamily
}
// GetVersion gets the version property value. Template version. Valid values 1 to 2147483647. This property is read-only.
func (m *DeviceManagementConfigurationPolicyTemplate) GetVersion()(*int32) {
    return m.version
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationPolicyTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowUnmanagedSettings", m.GetAllowUnmanagedSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("baseId", m.GetBaseId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
        err = writer.WriteStringValue("displayVersion", m.GetDisplayVersion())
        if err != nil {
            return err
        }
    }
    if m.GetLifecycleState() != nil {
        cast := (*m.GetLifecycleState()).String()
        err = writer.WriteStringValue("lifecycleState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPlatforms() != nil {
        cast := (*m.GetPlatforms()).String()
        err = writer.WriteStringValue("platforms", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSettingTemplates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSettingTemplates()))
        for i, v := range m.GetSettingTemplates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("settingTemplates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTechnologies() != nil {
        cast := (*m.GetTechnologies()).String()
        err = writer.WriteStringValue("technologies", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTemplateFamily() != nil {
        cast := (*m.GetTemplateFamily()).String()
        err = writer.WriteStringValue("templateFamily", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowUnmanagedSettings sets the allowUnmanagedSettings property value. Allow unmanaged setting templates
func (m *DeviceManagementConfigurationPolicyTemplate) SetAllowUnmanagedSettings(value *bool)() {
    m.allowUnmanagedSettings = value
}
// SetBaseId sets the baseId property value. Template base identifier
func (m *DeviceManagementConfigurationPolicyTemplate) SetBaseId(value *string)() {
    m.baseId = value
}
// SetDescription sets the description property value. Template description
func (m *DeviceManagementConfigurationPolicyTemplate) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Template display name
func (m *DeviceManagementConfigurationPolicyTemplate) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDisplayVersion sets the displayVersion property value. Description of template version
func (m *DeviceManagementConfigurationPolicyTemplate) SetDisplayVersion(value *string)() {
    m.displayVersion = value
}
// SetLifecycleState sets the lifecycleState property value. Describes current lifecycle state of a template
func (m *DeviceManagementConfigurationPolicyTemplate) SetLifecycleState(value *DeviceManagementTemplateLifecycleState)() {
    m.lifecycleState = value
}
// SetPlatforms sets the platforms property value. Supported platform types.
func (m *DeviceManagementConfigurationPolicyTemplate) SetPlatforms(value *DeviceManagementConfigurationPlatforms)() {
    m.platforms = value
}
// SetSettingTemplateCount sets the settingTemplateCount property value. Number of setting templates. Valid values 0 to 2147483647. This property is read-only.
func (m *DeviceManagementConfigurationPolicyTemplate) SetSettingTemplateCount(value *int32)() {
    m.settingTemplateCount = value
}
// SetSettingTemplates sets the settingTemplates property value. Setting templates
func (m *DeviceManagementConfigurationPolicyTemplate) SetSettingTemplates(value []DeviceManagementConfigurationSettingTemplateable)() {
    m.settingTemplates = value
}
// SetTechnologies sets the technologies property value. Describes which technology this setting can be deployed with
func (m *DeviceManagementConfigurationPolicyTemplate) SetTechnologies(value *DeviceManagementConfigurationTechnologies)() {
    m.technologies = value
}
// SetTemplateFamily sets the templateFamily property value. Describes the TemplateFamily for the Template entity
func (m *DeviceManagementConfigurationPolicyTemplate) SetTemplateFamily(value *DeviceManagementConfigurationTemplateFamily)() {
    m.templateFamily = value
}
// SetVersion sets the version property value. Template version. Valid values 1 to 2147483647. This property is read-only.
func (m *DeviceManagementConfigurationPolicyTemplate) SetVersion(value *int32)() {
    m.version = value
}
