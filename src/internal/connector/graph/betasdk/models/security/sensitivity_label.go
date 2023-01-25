package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// SensitivityLabel provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type SensitivityLabel struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // The color that the UI should display for the label, if configured.
    color *string
    // Returns the supported content formats for the label.
    contentFormats []string
    // The admin-defined description for the label.
    description *string
    // Indicates whether the label has protection actions configured.
    hasProtection *bool
    // Indicates whether the label is active or not. Active labels should be hidden or disabled in the UI.
    isActive *bool
    // Indicates whether the label can be applied to content. False if the label is a parent with child labels.
    isAppliable *bool
    // The plaintext name of the label.
    name *string
    // The parent label associated with a child label. Null if the label has no parent.
    parent SensitivityLabelable
    // The sensitivity value of the label, where lower is less sensitive.
    sensitivity *int32
    // The tooltip that should be displayed for the label in a UI.
    tooltip *string
}
// NewSensitivityLabel instantiates a new sensitivityLabel and sets the default values.
func NewSensitivityLabel()(*SensitivityLabel) {
    m := &SensitivityLabel{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateSensitivityLabelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSensitivityLabelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSensitivityLabel(), nil
}
// GetColor gets the color property value. The color that the UI should display for the label, if configured.
func (m *SensitivityLabel) GetColor()(*string) {
    return m.color
}
// GetContentFormats gets the contentFormats property value. Returns the supported content formats for the label.
func (m *SensitivityLabel) GetContentFormats()([]string) {
    return m.contentFormats
}
// GetDescription gets the description property value. The admin-defined description for the label.
func (m *SensitivityLabel) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SensitivityLabel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["color"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val)
        }
        return nil
    }
    res["contentFormats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetContentFormats(res)
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
    res["hasProtection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasProtection(val)
        }
        return nil
    }
    res["isActive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsActive(val)
        }
        return nil
    }
    res["isAppliable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAppliable(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["parent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSensitivityLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParent(val.(SensitivityLabelable))
        }
        return nil
    }
    res["sensitivity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensitivity(val)
        }
        return nil
    }
    res["tooltip"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTooltip(val)
        }
        return nil
    }
    return res
}
// GetHasProtection gets the hasProtection property value. Indicates whether the label has protection actions configured.
func (m *SensitivityLabel) GetHasProtection()(*bool) {
    return m.hasProtection
}
// GetIsActive gets the isActive property value. Indicates whether the label is active or not. Active labels should be hidden or disabled in the UI.
func (m *SensitivityLabel) GetIsActive()(*bool) {
    return m.isActive
}
// GetIsAppliable gets the isAppliable property value. Indicates whether the label can be applied to content. False if the label is a parent with child labels.
func (m *SensitivityLabel) GetIsAppliable()(*bool) {
    return m.isAppliable
}
// GetName gets the name property value. The plaintext name of the label.
func (m *SensitivityLabel) GetName()(*string) {
    return m.name
}
// GetParent gets the parent property value. The parent label associated with a child label. Null if the label has no parent.
func (m *SensitivityLabel) GetParent()(SensitivityLabelable) {
    return m.parent
}
// GetSensitivity gets the sensitivity property value. The sensitivity value of the label, where lower is less sensitive.
func (m *SensitivityLabel) GetSensitivity()(*int32) {
    return m.sensitivity
}
// GetTooltip gets the tooltip property value. The tooltip that should be displayed for the label in a UI.
func (m *SensitivityLabel) GetTooltip()(*string) {
    return m.tooltip
}
// Serialize serializes information the current object
func (m *SensitivityLabel) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("color", m.GetColor())
        if err != nil {
            return err
        }
    }
    if m.GetContentFormats() != nil {
        err = writer.WriteCollectionOfStringValues("contentFormats", m.GetContentFormats())
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
        err = writer.WriteBoolValue("hasProtection", m.GetHasProtection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isActive", m.GetIsActive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAppliable", m.GetIsAppliable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("parent", m.GetParent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("sensitivity", m.GetSensitivity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tooltip", m.GetTooltip())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetColor sets the color property value. The color that the UI should display for the label, if configured.
func (m *SensitivityLabel) SetColor(value *string)() {
    m.color = value
}
// SetContentFormats sets the contentFormats property value. Returns the supported content formats for the label.
func (m *SensitivityLabel) SetContentFormats(value []string)() {
    m.contentFormats = value
}
// SetDescription sets the description property value. The admin-defined description for the label.
func (m *SensitivityLabel) SetDescription(value *string)() {
    m.description = value
}
// SetHasProtection sets the hasProtection property value. Indicates whether the label has protection actions configured.
func (m *SensitivityLabel) SetHasProtection(value *bool)() {
    m.hasProtection = value
}
// SetIsActive sets the isActive property value. Indicates whether the label is active or not. Active labels should be hidden or disabled in the UI.
func (m *SensitivityLabel) SetIsActive(value *bool)() {
    m.isActive = value
}
// SetIsAppliable sets the isAppliable property value. Indicates whether the label can be applied to content. False if the label is a parent with child labels.
func (m *SensitivityLabel) SetIsAppliable(value *bool)() {
    m.isAppliable = value
}
// SetName sets the name property value. The plaintext name of the label.
func (m *SensitivityLabel) SetName(value *string)() {
    m.name = value
}
// SetParent sets the parent property value. The parent label associated with a child label. Null if the label has no parent.
func (m *SensitivityLabel) SetParent(value SensitivityLabelable)() {
    m.parent = value
}
// SetSensitivity sets the sensitivity property value. The sensitivity value of the label, where lower is less sensitive.
func (m *SensitivityLabel) SetSensitivity(value *int32)() {
    m.sensitivity = value
}
// SetTooltip sets the tooltip property value. The tooltip that should be displayed for the label in a UI.
func (m *SensitivityLabel) SetTooltip(value *string)() {
    m.tooltip = value
}
