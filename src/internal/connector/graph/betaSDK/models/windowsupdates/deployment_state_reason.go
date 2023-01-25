package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeploymentStateReason 
type DeploymentStateReason struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // The value property
    value *DeploymentStateReasonValue
}
// NewDeploymentStateReason instantiates a new deploymentStateReason and sets the default values.
func NewDeploymentStateReason()(*DeploymentStateReason) {
    m := &DeploymentStateReason{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeploymentStateReasonFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeploymentStateReasonFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeploymentStateReason(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeploymentStateReason) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeploymentStateReason) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeploymentStateReasonValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(*DeploymentStateReasonValue))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeploymentStateReason) GetOdataType()(*string) {
    return m.odataType
}
// GetValue gets the value property value. The value property
func (m *DeploymentStateReason) GetValue()(*DeploymentStateReasonValue) {
    return m.value
}
// Serialize serializes information the current object
func (m *DeploymentStateReason) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetValue() != nil {
        cast := (*m.GetValue()).String()
        err := writer.WriteStringValue("value", &cast)
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
func (m *DeploymentStateReason) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeploymentStateReason) SetOdataType(value *string)() {
    m.odataType = value
}
// SetValue sets the value property value. The value property
func (m *DeploymentStateReason) SetValue(value *DeploymentStateReasonValue)() {
    m.value = value
}
