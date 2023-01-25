package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ActionUrl 
type ActionUrl struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The displayName property
    displayName *string
    // The OdataType property
    odataType *string
    // The url property
    url *string
}
// NewActionUrl instantiates a new actionUrl and sets the default values.
func NewActionUrl()(*ActionUrl) {
    m := &ActionUrl{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateActionUrlFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateActionUrlFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActionUrl(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ActionUrl) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *ActionUrl) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ActionUrl) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ActionUrl) GetOdataType()(*string) {
    return m.odataType
}
// GetUrl gets the url property value. The url property
func (m *ActionUrl) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *ActionUrl) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
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
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *ActionUrl) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *ActionUrl) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ActionUrl) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUrl sets the url property value. The url property
func (m *ActionUrl) SetUrl(value *string)() {
    m.url = value
}
