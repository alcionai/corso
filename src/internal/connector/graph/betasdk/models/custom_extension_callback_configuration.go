package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomExtensionCallbackConfiguration 
type CustomExtensionCallbackConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // Callback time out in ISO 8601 time duration. Accepted time durations are between five minutes to three hours. For example, PT5M for five minutes and PT3H for three hours.
    timeoutDuration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
}
// NewCustomExtensionCallbackConfiguration instantiates a new customExtensionCallbackConfiguration and sets the default values.
func NewCustomExtensionCallbackConfiguration()(*CustomExtensionCallbackConfiguration) {
    m := &CustomExtensionCallbackConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCustomExtensionCallbackConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomExtensionCallbackConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomExtensionCallbackConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CustomExtensionCallbackConfiguration) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomExtensionCallbackConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["timeoutDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeoutDuration(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CustomExtensionCallbackConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetTimeoutDuration gets the timeoutDuration property value. Callback time out in ISO 8601 time duration. Accepted time durations are between five minutes to three hours. For example, PT5M for five minutes and PT3H for three hours.
func (m *CustomExtensionCallbackConfiguration) GetTimeoutDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.timeoutDuration
}
// Serialize serializes information the current object
func (m *CustomExtensionCallbackConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("timeoutDuration", m.GetTimeoutDuration())
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
func (m *CustomExtensionCallbackConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CustomExtensionCallbackConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTimeoutDuration sets the timeoutDuration property value. Callback time out in ISO 8601 time duration. Accepted time durations are between five minutes to three hours. For example, PT5M for five minutes and PT3H for three hours.
func (m *CustomExtensionCallbackConfiguration) SetTimeoutDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.timeoutDuration = value
}
