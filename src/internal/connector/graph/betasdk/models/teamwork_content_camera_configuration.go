package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkContentCameraConfiguration 
type TeamworkContentCameraConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // True if the content camera is inverted.
    isContentCameraInverted *bool
    // True if the content camera is optional.
    isContentCameraOptional *bool
    // True if the content enhancement is enabled.
    isContentEnhancementEnabled *bool
    // The OdataType property
    odataType *string
}
// NewTeamworkContentCameraConfiguration instantiates a new teamworkContentCameraConfiguration and sets the default values.
func NewTeamworkContentCameraConfiguration()(*TeamworkContentCameraConfiguration) {
    m := &TeamworkContentCameraConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkContentCameraConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkContentCameraConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkContentCameraConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkContentCameraConfiguration) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkContentCameraConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isContentCameraInverted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsContentCameraInverted(val)
        }
        return nil
    }
    res["isContentCameraOptional"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsContentCameraOptional(val)
        }
        return nil
    }
    res["isContentEnhancementEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsContentEnhancementEnabled(val)
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
// GetIsContentCameraInverted gets the isContentCameraInverted property value. True if the content camera is inverted.
func (m *TeamworkContentCameraConfiguration) GetIsContentCameraInverted()(*bool) {
    return m.isContentCameraInverted
}
// GetIsContentCameraOptional gets the isContentCameraOptional property value. True if the content camera is optional.
func (m *TeamworkContentCameraConfiguration) GetIsContentCameraOptional()(*bool) {
    return m.isContentCameraOptional
}
// GetIsContentEnhancementEnabled gets the isContentEnhancementEnabled property value. True if the content enhancement is enabled.
func (m *TeamworkContentCameraConfiguration) GetIsContentEnhancementEnabled()(*bool) {
    return m.isContentEnhancementEnabled
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkContentCameraConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *TeamworkContentCameraConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isContentCameraInverted", m.GetIsContentCameraInverted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isContentCameraOptional", m.GetIsContentCameraOptional())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isContentEnhancementEnabled", m.GetIsContentEnhancementEnabled())
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
func (m *TeamworkContentCameraConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetIsContentCameraInverted sets the isContentCameraInverted property value. True if the content camera is inverted.
func (m *TeamworkContentCameraConfiguration) SetIsContentCameraInverted(value *bool)() {
    m.isContentCameraInverted = value
}
// SetIsContentCameraOptional sets the isContentCameraOptional property value. True if the content camera is optional.
func (m *TeamworkContentCameraConfiguration) SetIsContentCameraOptional(value *bool)() {
    m.isContentCameraOptional = value
}
// SetIsContentEnhancementEnabled sets the isContentEnhancementEnabled property value. True if the content enhancement is enabled.
func (m *TeamworkContentCameraConfiguration) SetIsContentEnhancementEnabled(value *bool)() {
    m.isContentEnhancementEnabled = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkContentCameraConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
