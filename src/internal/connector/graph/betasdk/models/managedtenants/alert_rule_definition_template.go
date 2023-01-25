package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AlertRuleDefinitionTemplate 
type AlertRuleDefinitionTemplate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The defaultSeverity property
    defaultSeverity *AlertSeverity
    // The OdataType property
    odataType *string
}
// NewAlertRuleDefinitionTemplate instantiates a new alertRuleDefinitionTemplate and sets the default values.
func NewAlertRuleDefinitionTemplate()(*AlertRuleDefinitionTemplate) {
    m := &AlertRuleDefinitionTemplate{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAlertRuleDefinitionTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAlertRuleDefinitionTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAlertRuleDefinitionTemplate(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AlertRuleDefinitionTemplate) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDefaultSeverity gets the defaultSeverity property value. The defaultSeverity property
func (m *AlertRuleDefinitionTemplate) GetDefaultSeverity()(*AlertSeverity) {
    return m.defaultSeverity
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AlertRuleDefinitionTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["defaultSeverity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAlertSeverity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultSeverity(val.(*AlertSeverity))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AlertRuleDefinitionTemplate) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AlertRuleDefinitionTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDefaultSeverity() != nil {
        cast := (*m.GetDefaultSeverity()).String()
        err := writer.WriteStringValue("defaultSeverity", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AlertRuleDefinitionTemplate) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDefaultSeverity sets the defaultSeverity property value. The defaultSeverity property
func (m *AlertRuleDefinitionTemplate) SetDefaultSeverity(value *AlertSeverity)() {
    m.defaultSeverity = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AlertRuleDefinitionTemplate) SetOdataType(value *string)() {
    m.odataType = value
}
