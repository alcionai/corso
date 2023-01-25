package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InformationProtectionLabel provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type InformationProtectionLabel struct {
    Entity
    // The color that the UI should display for the label, if configured.
    color *string
    // The admin-defined description for the label.
    description *string
    // Indicates whether the label is active or not. Active labels should be hidden or disabled in UI.
    isActive *bool
    // The plaintext name of the label.
    name *string
    // The parent label associated with a child label. Null if label has no parent.
    parent ParentLabelDetailsable
    // The sensitivity value of the label, where lower is less sensitive.
    sensitivity *int32
    // The tooltip that should be displayed for the label in a UI.
    tooltip *string
}
// NewInformationProtectionLabel instantiates a new informationProtectionLabel and sets the default values.
func NewInformationProtectionLabel()(*InformationProtectionLabel) {
    m := &InformationProtectionLabel{
        Entity: *NewEntity(),
    }
    return m
}
// CreateInformationProtectionLabelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInformationProtectionLabelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInformationProtectionLabel(), nil
}
// GetColor gets the color property value. The color that the UI should display for the label, if configured.
func (m *InformationProtectionLabel) GetColor()(*string) {
    return m.color
}
// GetDescription gets the description property value. The admin-defined description for the label.
func (m *InformationProtectionLabel) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InformationProtectionLabel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetObjectValue(CreateParentLabelDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParent(val.(ParentLabelDetailsable))
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
// GetIsActive gets the isActive property value. Indicates whether the label is active or not. Active labels should be hidden or disabled in UI.
func (m *InformationProtectionLabel) GetIsActive()(*bool) {
    return m.isActive
}
// GetName gets the name property value. The plaintext name of the label.
func (m *InformationProtectionLabel) GetName()(*string) {
    return m.name
}
// GetParent gets the parent property value. The parent label associated with a child label. Null if label has no parent.
func (m *InformationProtectionLabel) GetParent()(ParentLabelDetailsable) {
    return m.parent
}
// GetSensitivity gets the sensitivity property value. The sensitivity value of the label, where lower is less sensitive.
func (m *InformationProtectionLabel) GetSensitivity()(*int32) {
    return m.sensitivity
}
// GetTooltip gets the tooltip property value. The tooltip that should be displayed for the label in a UI.
func (m *InformationProtectionLabel) GetTooltip()(*string) {
    return m.tooltip
}
// Serialize serializes information the current object
func (m *InformationProtectionLabel) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
func (m *InformationProtectionLabel) SetColor(value *string)() {
    m.color = value
}
// SetDescription sets the description property value. The admin-defined description for the label.
func (m *InformationProtectionLabel) SetDescription(value *string)() {
    m.description = value
}
// SetIsActive sets the isActive property value. Indicates whether the label is active or not. Active labels should be hidden or disabled in UI.
func (m *InformationProtectionLabel) SetIsActive(value *bool)() {
    m.isActive = value
}
// SetName sets the name property value. The plaintext name of the label.
func (m *InformationProtectionLabel) SetName(value *string)() {
    m.name = value
}
// SetParent sets the parent property value. The parent label associated with a child label. Null if label has no parent.
func (m *InformationProtectionLabel) SetParent(value ParentLabelDetailsable)() {
    m.parent = value
}
// SetSensitivity sets the sensitivity property value. The sensitivity value of the label, where lower is less sensitive.
func (m *InformationProtectionLabel) SetSensitivity(value *int32)() {
    m.sensitivity = value
}
// SetTooltip sets the tooltip property value. The tooltip that should be displayed for the label in a UI.
func (m *InformationProtectionLabel) SetTooltip(value *string)() {
    m.tooltip = value
}
