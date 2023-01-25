package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementTemplate entity that represents a defined collection of device settings
type DeviceManagementTemplate struct {
    Entity
    // Collection of setting categories within the template
    categories []DeviceManagementTemplateSettingCategoryable
    // The template's description
    description *string
    // The template's display name
    displayName *string
    // Number of Intents created from this template.
    intentCount *int32
    // The template is deprecated or not. Intents cannot be created from a deprecated template.
    isDeprecated *bool
    // Collection of templates this template can migrate to
    migratableTo []DeviceManagementTemplateable
    // Supported platform types for policies.
    platformType *PolicyPlatformType
    // When the template was published
    publishedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Collection of all settings this template has
    settings []DeviceManagementSettingInstanceable
    // Template subtype
    templateSubtype *DeviceManagementTemplateSubtype
    // Template type
    templateType *DeviceManagementTemplateType
    // The template's version information
    versionInfo *string
}
// NewDeviceManagementTemplate instantiates a new deviceManagementTemplate and sets the default values.
func NewDeviceManagementTemplate()(*DeviceManagementTemplate) {
    m := &DeviceManagementTemplate{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.securityBaselineTemplate":
                        return NewSecurityBaselineTemplate(), nil
                }
            }
        }
    }
    return NewDeviceManagementTemplate(), nil
}
// GetCategories gets the categories property value. Collection of setting categories within the template
func (m *DeviceManagementTemplate) GetCategories()([]DeviceManagementTemplateSettingCategoryable) {
    return m.categories
}
// GetDescription gets the description property value. The template's description
func (m *DeviceManagementTemplate) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The template's display name
func (m *DeviceManagementTemplate) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["categories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementTemplateSettingCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementTemplateSettingCategoryable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementTemplateSettingCategoryable)
            }
            m.SetCategories(res)
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
    res["intentCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntentCount(val)
        }
        return nil
    }
    res["isDeprecated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDeprecated(val)
        }
        return nil
    }
    res["migratableTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementTemplateable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementTemplateable)
            }
            m.SetMigratableTo(res)
        }
        return nil
    }
    res["platformType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePolicyPlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatformType(val.(*PolicyPlatformType))
        }
        return nil
    }
    res["publishedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishedDateTime(val)
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementSettingInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementSettingInstanceable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementSettingInstanceable)
            }
            m.SetSettings(res)
        }
        return nil
    }
    res["templateSubtype"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementTemplateSubtype)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateSubtype(val.(*DeviceManagementTemplateSubtype))
        }
        return nil
    }
    res["templateType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementTemplateType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateType(val.(*DeviceManagementTemplateType))
        }
        return nil
    }
    res["versionInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersionInfo(val)
        }
        return nil
    }
    return res
}
// GetIntentCount gets the intentCount property value. Number of Intents created from this template.
func (m *DeviceManagementTemplate) GetIntentCount()(*int32) {
    return m.intentCount
}
// GetIsDeprecated gets the isDeprecated property value. The template is deprecated or not. Intents cannot be created from a deprecated template.
func (m *DeviceManagementTemplate) GetIsDeprecated()(*bool) {
    return m.isDeprecated
}
// GetMigratableTo gets the migratableTo property value. Collection of templates this template can migrate to
func (m *DeviceManagementTemplate) GetMigratableTo()([]DeviceManagementTemplateable) {
    return m.migratableTo
}
// GetPlatformType gets the platformType property value. Supported platform types for policies.
func (m *DeviceManagementTemplate) GetPlatformType()(*PolicyPlatformType) {
    return m.platformType
}
// GetPublishedDateTime gets the publishedDateTime property value. When the template was published
func (m *DeviceManagementTemplate) GetPublishedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.publishedDateTime
}
// GetSettings gets the settings property value. Collection of all settings this template has
func (m *DeviceManagementTemplate) GetSettings()([]DeviceManagementSettingInstanceable) {
    return m.settings
}
// GetTemplateSubtype gets the templateSubtype property value. Template subtype
func (m *DeviceManagementTemplate) GetTemplateSubtype()(*DeviceManagementTemplateSubtype) {
    return m.templateSubtype
}
// GetTemplateType gets the templateType property value. Template type
func (m *DeviceManagementTemplate) GetTemplateType()(*DeviceManagementTemplateType) {
    return m.templateType
}
// GetVersionInfo gets the versionInfo property value. The template's version information
func (m *DeviceManagementTemplate) GetVersionInfo()(*string) {
    return m.versionInfo
}
// Serialize serializes information the current object
func (m *DeviceManagementTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCategories()))
        for i, v := range m.GetCategories() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("categories", cast)
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
        err = writer.WriteInt32Value("intentCount", m.GetIntentCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDeprecated", m.GetIsDeprecated())
        if err != nil {
            return err
        }
    }
    if m.GetMigratableTo() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMigratableTo()))
        for i, v := range m.GetMigratableTo() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("migratableTo", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPlatformType() != nil {
        cast := (*m.GetPlatformType()).String()
        err = writer.WriteStringValue("platformType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("publishedDateTime", m.GetPublishedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSettings()))
        for i, v := range m.GetSettings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("settings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTemplateSubtype() != nil {
        cast := (*m.GetTemplateSubtype()).String()
        err = writer.WriteStringValue("templateSubtype", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTemplateType() != nil {
        cast := (*m.GetTemplateType()).String()
        err = writer.WriteStringValue("templateType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("versionInfo", m.GetVersionInfo())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCategories sets the categories property value. Collection of setting categories within the template
func (m *DeviceManagementTemplate) SetCategories(value []DeviceManagementTemplateSettingCategoryable)() {
    m.categories = value
}
// SetDescription sets the description property value. The template's description
func (m *DeviceManagementTemplate) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The template's display name
func (m *DeviceManagementTemplate) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIntentCount sets the intentCount property value. Number of Intents created from this template.
func (m *DeviceManagementTemplate) SetIntentCount(value *int32)() {
    m.intentCount = value
}
// SetIsDeprecated sets the isDeprecated property value. The template is deprecated or not. Intents cannot be created from a deprecated template.
func (m *DeviceManagementTemplate) SetIsDeprecated(value *bool)() {
    m.isDeprecated = value
}
// SetMigratableTo sets the migratableTo property value. Collection of templates this template can migrate to
func (m *DeviceManagementTemplate) SetMigratableTo(value []DeviceManagementTemplateable)() {
    m.migratableTo = value
}
// SetPlatformType sets the platformType property value. Supported platform types for policies.
func (m *DeviceManagementTemplate) SetPlatformType(value *PolicyPlatformType)() {
    m.platformType = value
}
// SetPublishedDateTime sets the publishedDateTime property value. When the template was published
func (m *DeviceManagementTemplate) SetPublishedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.publishedDateTime = value
}
// SetSettings sets the settings property value. Collection of all settings this template has
func (m *DeviceManagementTemplate) SetSettings(value []DeviceManagementSettingInstanceable)() {
    m.settings = value
}
// SetTemplateSubtype sets the templateSubtype property value. Template subtype
func (m *DeviceManagementTemplate) SetTemplateSubtype(value *DeviceManagementTemplateSubtype)() {
    m.templateSubtype = value
}
// SetTemplateType sets the templateType property value. Template type
func (m *DeviceManagementTemplate) SetTemplateType(value *DeviceManagementTemplateType)() {
    m.templateType = value
}
// SetVersionInfo sets the versionInfo property value. The template's version information
func (m *DeviceManagementTemplate) SetVersionInfo(value *string)() {
    m.versionInfo = value
}
