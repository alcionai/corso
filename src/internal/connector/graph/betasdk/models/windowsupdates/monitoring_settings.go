package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MonitoringSettings 
type MonitoringSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Specifies the rules through which monitoring signals can trigger actions on the deployment. Rules are combined using 'or'.
    monitoringRules []MonitoringRuleable
    // The OdataType property
    odataType *string
}
// NewMonitoringSettings instantiates a new monitoringSettings and sets the default values.
func NewMonitoringSettings()(*MonitoringSettings) {
    m := &MonitoringSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMonitoringSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMonitoringSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMonitoringSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MonitoringSettings) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MonitoringSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["monitoringRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMonitoringRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MonitoringRuleable, len(val))
            for i, v := range val {
                res[i] = v.(MonitoringRuleable)
            }
            m.SetMonitoringRules(res)
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
// GetMonitoringRules gets the monitoringRules property value. Specifies the rules through which monitoring signals can trigger actions on the deployment. Rules are combined using 'or'.
func (m *MonitoringSettings) GetMonitoringRules()([]MonitoringRuleable) {
    return m.monitoringRules
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MonitoringSettings) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *MonitoringSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetMonitoringRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMonitoringRules()))
        for i, v := range m.GetMonitoringRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("monitoringRules", cast)
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
func (m *MonitoringSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetMonitoringRules sets the monitoringRules property value. Specifies the rules through which monitoring signals can trigger actions on the deployment. Rules are combined using 'or'.
func (m *MonitoringSettings) SetMonitoringRules(value []MonitoringRuleable)() {
    m.monitoringRules = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MonitoringSettings) SetOdataType(value *string)() {
    m.odataType = value
}
