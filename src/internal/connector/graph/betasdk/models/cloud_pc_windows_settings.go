package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcWindowsSettings 
type CloudPcWindowsSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The Windows language/region tag to use for language pack configuration and localization of the Cloud PC. The default value is en-US, which corresponds to English (United States).
    language *string
    // The OdataType property
    odataType *string
}
// NewCloudPcWindowsSettings instantiates a new cloudPcWindowsSettings and sets the default values.
func NewCloudPcWindowsSettings()(*CloudPcWindowsSettings) {
    m := &CloudPcWindowsSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCloudPcWindowsSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcWindowsSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcWindowsSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcWindowsSettings) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcWindowsSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["language"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguage(val)
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
// GetLanguage gets the language property value. The Windows language/region tag to use for language pack configuration and localization of the Cloud PC. The default value is en-US, which corresponds to English (United States).
func (m *CloudPcWindowsSettings) GetLanguage()(*string) {
    return m.language
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CloudPcWindowsSettings) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *CloudPcWindowsSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("language", m.GetLanguage())
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
func (m *CloudPcWindowsSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetLanguage sets the language property value. The Windows language/region tag to use for language pack configuration and localization of the Cloud PC. The default value is en-US, which corresponds to English (United States).
func (m *CloudPcWindowsSettings) SetLanguage(value *string)() {
    m.language = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcWindowsSettings) SetOdataType(value *string)() {
    m.odataType = value
}
