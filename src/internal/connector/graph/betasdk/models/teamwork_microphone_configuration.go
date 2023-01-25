package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkMicrophoneConfiguration 
type TeamworkMicrophoneConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The defaultMicrophone property
    defaultMicrophone TeamworkPeripheralable
    // True if the configured microphone is optional. False if the microphone is not optional and the health state of the device should be computed.
    isMicrophoneOptional *bool
    // The microphones property
    microphones []TeamworkPeripheralable
    // The OdataType property
    odataType *string
}
// NewTeamworkMicrophoneConfiguration instantiates a new teamworkMicrophoneConfiguration and sets the default values.
func NewTeamworkMicrophoneConfiguration()(*TeamworkMicrophoneConfiguration) {
    m := &TeamworkMicrophoneConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkMicrophoneConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkMicrophoneConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkMicrophoneConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkMicrophoneConfiguration) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDefaultMicrophone gets the defaultMicrophone property value. The defaultMicrophone property
func (m *TeamworkMicrophoneConfiguration) GetDefaultMicrophone()(TeamworkPeripheralable) {
    return m.defaultMicrophone
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkMicrophoneConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["defaultMicrophone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultMicrophone(val.(TeamworkPeripheralable))
        }
        return nil
    }
    res["isMicrophoneOptional"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMicrophoneOptional(val)
        }
        return nil
    }
    res["microphones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTeamworkPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamworkPeripheralable, len(val))
            for i, v := range val {
                res[i] = v.(TeamworkPeripheralable)
            }
            m.SetMicrophones(res)
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
// GetIsMicrophoneOptional gets the isMicrophoneOptional property value. True if the configured microphone is optional. False if the microphone is not optional and the health state of the device should be computed.
func (m *TeamworkMicrophoneConfiguration) GetIsMicrophoneOptional()(*bool) {
    return m.isMicrophoneOptional
}
// GetMicrophones gets the microphones property value. The microphones property
func (m *TeamworkMicrophoneConfiguration) GetMicrophones()([]TeamworkPeripheralable) {
    return m.microphones
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkMicrophoneConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *TeamworkMicrophoneConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("defaultMicrophone", m.GetDefaultMicrophone())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isMicrophoneOptional", m.GetIsMicrophoneOptional())
        if err != nil {
            return err
        }
    }
    if m.GetMicrophones() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMicrophones()))
        for i, v := range m.GetMicrophones() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("microphones", cast)
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
func (m *TeamworkMicrophoneConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDefaultMicrophone sets the defaultMicrophone property value. The defaultMicrophone property
func (m *TeamworkMicrophoneConfiguration) SetDefaultMicrophone(value TeamworkPeripheralable)() {
    m.defaultMicrophone = value
}
// SetIsMicrophoneOptional sets the isMicrophoneOptional property value. True if the configured microphone is optional. False if the microphone is not optional and the health state of the device should be computed.
func (m *TeamworkMicrophoneConfiguration) SetIsMicrophoneOptional(value *bool)() {
    m.isMicrophoneOptional = value
}
// SetMicrophones sets the microphones property value. The microphones property
func (m *TeamworkMicrophoneConfiguration) SetMicrophones(value []TeamworkPeripheralable)() {
    m.microphones = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkMicrophoneConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
