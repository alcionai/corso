package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementSettingDefinition entity representing the defintion for a given setting
type DeviceManagementSettingDefinition struct {
    Entity
    // Collection of constraints for the setting value
    constraints []DeviceManagementConstraintable
    // Collection of dependencies on other settings
    dependencies []DeviceManagementSettingDependencyable
    // The setting's description
    description *string
    // The setting's display name
    displayName *string
    // Url to setting documentation
    documentationUrl *string
    // subtitle of the setting header for more details about the category/section
    headerSubtitle *string
    // title of the setting header represents a category/section of a setting/settings
    headerTitle *string
    // If the setting is top level, it can be configured without the need to be wrapped in a collection or complex setting
    isTopLevel *bool
    // Keywords associated with the setting
    keywords []string
    // Placeholder text as an example of valid input
    placeholderText *string
    // The valueType property
    valueType *DeviceManangementIntentValueType
}
// NewDeviceManagementSettingDefinition instantiates a new deviceManagementSettingDefinition and sets the default values.
func NewDeviceManagementSettingDefinition()(*DeviceManagementSettingDefinition) {
    m := &DeviceManagementSettingDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementSettingDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementSettingDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.deviceManagementAbstractComplexSettingDefinition":
                        return NewDeviceManagementAbstractComplexSettingDefinition(), nil
                    case "#microsoft.graph.deviceManagementCollectionSettingDefinition":
                        return NewDeviceManagementCollectionSettingDefinition(), nil
                    case "#microsoft.graph.deviceManagementComplexSettingDefinition":
                        return NewDeviceManagementComplexSettingDefinition(), nil
                }
            }
        }
    }
    return NewDeviceManagementSettingDefinition(), nil
}
// GetConstraints gets the constraints property value. Collection of constraints for the setting value
func (m *DeviceManagementSettingDefinition) GetConstraints()([]DeviceManagementConstraintable) {
    return m.constraints
}
// GetDependencies gets the dependencies property value. Collection of dependencies on other settings
func (m *DeviceManagementSettingDefinition) GetDependencies()([]DeviceManagementSettingDependencyable) {
    return m.dependencies
}
// GetDescription gets the description property value. The setting's description
func (m *DeviceManagementSettingDefinition) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The setting's display name
func (m *DeviceManagementSettingDefinition) GetDisplayName()(*string) {
    return m.displayName
}
// GetDocumentationUrl gets the documentationUrl property value. Url to setting documentation
func (m *DeviceManagementSettingDefinition) GetDocumentationUrl()(*string) {
    return m.documentationUrl
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementSettingDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["constraints"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConstraintFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConstraintable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConstraintable)
            }
            m.SetConstraints(res)
        }
        return nil
    }
    res["dependencies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementSettingDependencyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementSettingDependencyable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementSettingDependencyable)
            }
            m.SetDependencies(res)
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
    res["documentationUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentationUrl(val)
        }
        return nil
    }
    res["headerSubtitle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeaderSubtitle(val)
        }
        return nil
    }
    res["headerTitle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeaderTitle(val)
        }
        return nil
    }
    res["isTopLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTopLevel(val)
        }
        return nil
    }
    res["keywords"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetKeywords(res)
        }
        return nil
    }
    res["placeholderText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlaceholderText(val)
        }
        return nil
    }
    res["valueType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManangementIntentValueType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueType(val.(*DeviceManangementIntentValueType))
        }
        return nil
    }
    return res
}
// GetHeaderSubtitle gets the headerSubtitle property value. subtitle of the setting header for more details about the category/section
func (m *DeviceManagementSettingDefinition) GetHeaderSubtitle()(*string) {
    return m.headerSubtitle
}
// GetHeaderTitle gets the headerTitle property value. title of the setting header represents a category/section of a setting/settings
func (m *DeviceManagementSettingDefinition) GetHeaderTitle()(*string) {
    return m.headerTitle
}
// GetIsTopLevel gets the isTopLevel property value. If the setting is top level, it can be configured without the need to be wrapped in a collection or complex setting
func (m *DeviceManagementSettingDefinition) GetIsTopLevel()(*bool) {
    return m.isTopLevel
}
// GetKeywords gets the keywords property value. Keywords associated with the setting
func (m *DeviceManagementSettingDefinition) GetKeywords()([]string) {
    return m.keywords
}
// GetPlaceholderText gets the placeholderText property value. Placeholder text as an example of valid input
func (m *DeviceManagementSettingDefinition) GetPlaceholderText()(*string) {
    return m.placeholderText
}
// GetValueType gets the valueType property value. The valueType property
func (m *DeviceManagementSettingDefinition) GetValueType()(*DeviceManangementIntentValueType) {
    return m.valueType
}
// Serialize serializes information the current object
func (m *DeviceManagementSettingDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetConstraints() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConstraints()))
        for i, v := range m.GetConstraints() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("constraints", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDependencies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDependencies()))
        for i, v := range m.GetDependencies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("dependencies", cast)
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
        err = writer.WriteStringValue("documentationUrl", m.GetDocumentationUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("headerSubtitle", m.GetHeaderSubtitle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("headerTitle", m.GetHeaderTitle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isTopLevel", m.GetIsTopLevel())
        if err != nil {
            return err
        }
    }
    if m.GetKeywords() != nil {
        err = writer.WriteCollectionOfStringValues("keywords", m.GetKeywords())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("placeholderText", m.GetPlaceholderText())
        if err != nil {
            return err
        }
    }
    if m.GetValueType() != nil {
        cast := (*m.GetValueType()).String()
        err = writer.WriteStringValue("valueType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConstraints sets the constraints property value. Collection of constraints for the setting value
func (m *DeviceManagementSettingDefinition) SetConstraints(value []DeviceManagementConstraintable)() {
    m.constraints = value
}
// SetDependencies sets the dependencies property value. Collection of dependencies on other settings
func (m *DeviceManagementSettingDefinition) SetDependencies(value []DeviceManagementSettingDependencyable)() {
    m.dependencies = value
}
// SetDescription sets the description property value. The setting's description
func (m *DeviceManagementSettingDefinition) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The setting's display name
func (m *DeviceManagementSettingDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDocumentationUrl sets the documentationUrl property value. Url to setting documentation
func (m *DeviceManagementSettingDefinition) SetDocumentationUrl(value *string)() {
    m.documentationUrl = value
}
// SetHeaderSubtitle sets the headerSubtitle property value. subtitle of the setting header for more details about the category/section
func (m *DeviceManagementSettingDefinition) SetHeaderSubtitle(value *string)() {
    m.headerSubtitle = value
}
// SetHeaderTitle sets the headerTitle property value. title of the setting header represents a category/section of a setting/settings
func (m *DeviceManagementSettingDefinition) SetHeaderTitle(value *string)() {
    m.headerTitle = value
}
// SetIsTopLevel sets the isTopLevel property value. If the setting is top level, it can be configured without the need to be wrapped in a collection or complex setting
func (m *DeviceManagementSettingDefinition) SetIsTopLevel(value *bool)() {
    m.isTopLevel = value
}
// SetKeywords sets the keywords property value. Keywords associated with the setting
func (m *DeviceManagementSettingDefinition) SetKeywords(value []string)() {
    m.keywords = value
}
// SetPlaceholderText sets the placeholderText property value. Placeholder text as an example of valid input
func (m *DeviceManagementSettingDefinition) SetPlaceholderText(value *string)() {
    m.placeholderText = value
}
// SetValueType sets the valueType property value. The valueType property
func (m *DeviceManagementSettingDefinition) SetValueType(value *DeviceManangementIntentValueType)() {
    m.valueType = value
}
