package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementTroubleshootingErrorResource object representing a link to troubleshooting information, the link could be to the Azure Portal or a Microsoft doc.
type DeviceManagementTroubleshootingErrorResource struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The link to the web resource. Can contain any of the following formatters: {{UPN}}, {{DeviceGUID}}, {{UserGUID}}
    link *string
    // The OdataType property
    odataType *string
    // Not yet documented
    text *string
}
// NewDeviceManagementTroubleshootingErrorResource instantiates a new deviceManagementTroubleshootingErrorResource and sets the default values.
func NewDeviceManagementTroubleshootingErrorResource()(*DeviceManagementTroubleshootingErrorResource) {
    m := &DeviceManagementTroubleshootingErrorResource{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementTroubleshootingErrorResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementTroubleshootingErrorResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementTroubleshootingErrorResource(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementTroubleshootingErrorResource) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementTroubleshootingErrorResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["link"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLink(val)
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
    res["text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val)
        }
        return nil
    }
    return res
}
// GetLink gets the link property value. The link to the web resource. Can contain any of the following formatters: {{UPN}}, {{DeviceGUID}}, {{UserGUID}}
func (m *DeviceManagementTroubleshootingErrorResource) GetLink()(*string) {
    return m.link
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementTroubleshootingErrorResource) GetOdataType()(*string) {
    return m.odataType
}
// GetText gets the text property value. Not yet documented
func (m *DeviceManagementTroubleshootingErrorResource) GetText()(*string) {
    return m.text
}
// Serialize serializes information the current object
func (m *DeviceManagementTroubleshootingErrorResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("link", m.GetLink())
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
        err := writer.WriteStringValue("text", m.GetText())
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
func (m *DeviceManagementTroubleshootingErrorResource) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetLink sets the link property value. The link to the web resource. Can contain any of the following formatters: {{UPN}}, {{DeviceGUID}}, {{UserGUID}}
func (m *DeviceManagementTroubleshootingErrorResource) SetLink(value *string)() {
    m.link = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementTroubleshootingErrorResource) SetOdataType(value *string)() {
    m.odataType = value
}
// SetText sets the text property value. Not yet documented
func (m *DeviceManagementTroubleshootingErrorResource) SetText(value *string)() {
    m.text = value
}
