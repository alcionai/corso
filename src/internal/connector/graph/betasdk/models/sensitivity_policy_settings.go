package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SensitivityPolicySettings 
type SensitivityPolicySettings struct {
    Entity
    // The applicableTo property
    applicableTo *SensitivityLabelTarget
    // The downgradeSensitivityRequiresJustification property
    downgradeSensitivityRequiresJustification *bool
    // The helpWebUrl property
    helpWebUrl *string
    // The isMandatory property
    isMandatory *bool
}
// NewSensitivityPolicySettings instantiates a new sensitivityPolicySettings and sets the default values.
func NewSensitivityPolicySettings()(*SensitivityPolicySettings) {
    m := &SensitivityPolicySettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSensitivityPolicySettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSensitivityPolicySettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSensitivityPolicySettings(), nil
}
// GetApplicableTo gets the applicableTo property value. The applicableTo property
func (m *SensitivityPolicySettings) GetApplicableTo()(*SensitivityLabelTarget) {
    return m.applicableTo
}
// GetDowngradeSensitivityRequiresJustification gets the downgradeSensitivityRequiresJustification property value. The downgradeSensitivityRequiresJustification property
func (m *SensitivityPolicySettings) GetDowngradeSensitivityRequiresJustification()(*bool) {
    return m.downgradeSensitivityRequiresJustification
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SensitivityPolicySettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicableTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensitivityLabelTarget)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicableTo(val.(*SensitivityLabelTarget))
        }
        return nil
    }
    res["downgradeSensitivityRequiresJustification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDowngradeSensitivityRequiresJustification(val)
        }
        return nil
    }
    res["helpWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHelpWebUrl(val)
        }
        return nil
    }
    res["isMandatory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMandatory(val)
        }
        return nil
    }
    return res
}
// GetHelpWebUrl gets the helpWebUrl property value. The helpWebUrl property
func (m *SensitivityPolicySettings) GetHelpWebUrl()(*string) {
    return m.helpWebUrl
}
// GetIsMandatory gets the isMandatory property value. The isMandatory property
func (m *SensitivityPolicySettings) GetIsMandatory()(*bool) {
    return m.isMandatory
}
// Serialize serializes information the current object
func (m *SensitivityPolicySettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApplicableTo() != nil {
        cast := (*m.GetApplicableTo()).String()
        err = writer.WriteStringValue("applicableTo", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("downgradeSensitivityRequiresJustification", m.GetDowngradeSensitivityRequiresJustification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("helpWebUrl", m.GetHelpWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMandatory", m.GetIsMandatory())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicableTo sets the applicableTo property value. The applicableTo property
func (m *SensitivityPolicySettings) SetApplicableTo(value *SensitivityLabelTarget)() {
    m.applicableTo = value
}
// SetDowngradeSensitivityRequiresJustification sets the downgradeSensitivityRequiresJustification property value. The downgradeSensitivityRequiresJustification property
func (m *SensitivityPolicySettings) SetDowngradeSensitivityRequiresJustification(value *bool)() {
    m.downgradeSensitivityRequiresJustification = value
}
// SetHelpWebUrl sets the helpWebUrl property value. The helpWebUrl property
func (m *SensitivityPolicySettings) SetHelpWebUrl(value *string)() {
    m.helpWebUrl = value
}
// SetIsMandatory sets the isMandatory property value. The isMandatory property
func (m *SensitivityPolicySettings) SetIsMandatory(value *bool)() {
    m.isMandatory = value
}
