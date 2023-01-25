package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkHardwareConfiguration 
type TeamworkHardwareConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The compute property
    compute TeamworkPeripheralable
    // The hdmiIngest property
    hdmiIngest TeamworkPeripheralable
    // The OdataType property
    odataType *string
    // The CPU model on the device.
    processorModel *string
}
// NewTeamworkHardwareConfiguration instantiates a new teamworkHardwareConfiguration and sets the default values.
func NewTeamworkHardwareConfiguration()(*TeamworkHardwareConfiguration) {
    m := &TeamworkHardwareConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkHardwareConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkHardwareConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkHardwareConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkHardwareConfiguration) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCompute gets the compute property value. The compute property
func (m *TeamworkHardwareConfiguration) GetCompute()(TeamworkPeripheralable) {
    return m.compute
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkHardwareConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["compute"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompute(val.(TeamworkPeripheralable))
        }
        return nil
    }
    res["hdmiIngest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHdmiIngest(val.(TeamworkPeripheralable))
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
    res["processorModel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessorModel(val)
        }
        return nil
    }
    return res
}
// GetHdmiIngest gets the hdmiIngest property value. The hdmiIngest property
func (m *TeamworkHardwareConfiguration) GetHdmiIngest()(TeamworkPeripheralable) {
    return m.hdmiIngest
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkHardwareConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetProcessorModel gets the processorModel property value. The CPU model on the device.
func (m *TeamworkHardwareConfiguration) GetProcessorModel()(*string) {
    return m.processorModel
}
// Serialize serializes information the current object
func (m *TeamworkHardwareConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("compute", m.GetCompute())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("hdmiIngest", m.GetHdmiIngest())
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
        err := writer.WriteStringValue("processorModel", m.GetProcessorModel())
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
func (m *TeamworkHardwareConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCompute sets the compute property value. The compute property
func (m *TeamworkHardwareConfiguration) SetCompute(value TeamworkPeripheralable)() {
    m.compute = value
}
// SetHdmiIngest sets the hdmiIngest property value. The hdmiIngest property
func (m *TeamworkHardwareConfiguration) SetHdmiIngest(value TeamworkPeripheralable)() {
    m.hdmiIngest = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkHardwareConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetProcessorModel sets the processorModel property value. The CPU model on the device.
func (m *TeamworkHardwareConfiguration) SetProcessorModel(value *string)() {
    m.processorModel = value
}
