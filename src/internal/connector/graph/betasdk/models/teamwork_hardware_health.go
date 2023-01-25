package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkHardwareHealth 
type TeamworkHardwareHealth struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The system health details for a teamworkDevice.
    computeHealth TeamworkPeripheralHealthable
    // The health details about the HDMI ingest of a device.
    hdmiIngestHealth TeamworkPeripheralHealthable
    // The OdataType property
    odataType *string
}
// NewTeamworkHardwareHealth instantiates a new teamworkHardwareHealth and sets the default values.
func NewTeamworkHardwareHealth()(*TeamworkHardwareHealth) {
    m := &TeamworkHardwareHealth{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkHardwareHealthFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkHardwareHealthFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkHardwareHealth(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkHardwareHealth) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetComputeHealth gets the computeHealth property value. The system health details for a teamworkDevice.
func (m *TeamworkHardwareHealth) GetComputeHealth()(TeamworkPeripheralHealthable) {
    return m.computeHealth
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkHardwareHealth) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["computeHealth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralHealthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComputeHealth(val.(TeamworkPeripheralHealthable))
        }
        return nil
    }
    res["hdmiIngestHealth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralHealthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHdmiIngestHealth(val.(TeamworkPeripheralHealthable))
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
// GetHdmiIngestHealth gets the hdmiIngestHealth property value. The health details about the HDMI ingest of a device.
func (m *TeamworkHardwareHealth) GetHdmiIngestHealth()(TeamworkPeripheralHealthable) {
    return m.hdmiIngestHealth
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkHardwareHealth) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *TeamworkHardwareHealth) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("computeHealth", m.GetComputeHealth())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("hdmiIngestHealth", m.GetHdmiIngestHealth())
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
func (m *TeamworkHardwareHealth) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetComputeHealth sets the computeHealth property value. The system health details for a teamworkDevice.
func (m *TeamworkHardwareHealth) SetComputeHealth(value TeamworkPeripheralHealthable)() {
    m.computeHealth = value
}
// SetHdmiIngestHealth sets the hdmiIngestHealth property value. The health details about the HDMI ingest of a device.
func (m *TeamworkHardwareHealth) SetHdmiIngestHealth(value TeamworkPeripheralHealthable)() {
    m.hdmiIngestHealth = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkHardwareHealth) SetOdataType(value *string)() {
    m.odataType = value
}
