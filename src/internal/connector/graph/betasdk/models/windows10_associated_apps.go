package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10AssociatedApps windows 10 Associated Application definition.
type Windows10AssociatedApps struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Windows 10 Application type.
    appType *Windows10AppType
    // Identifier.
    identifier *string
    // The OdataType property
    odataType *string
}
// NewWindows10AssociatedApps instantiates a new windows10AssociatedApps and sets the default values.
func NewWindows10AssociatedApps()(*Windows10AssociatedApps) {
    m := &Windows10AssociatedApps{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindows10AssociatedAppsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10AssociatedAppsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10AssociatedApps(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Windows10AssociatedApps) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAppType gets the appType property value. Windows 10 Application type.
func (m *Windows10AssociatedApps) GetAppType()(*Windows10AppType) {
    return m.appType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10AssociatedApps) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindows10AppType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppType(val.(*Windows10AppType))
        }
        return nil
    }
    res["identifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentifier(val)
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
// GetIdentifier gets the identifier property value. Identifier.
func (m *Windows10AssociatedApps) GetIdentifier()(*string) {
    return m.identifier
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Windows10AssociatedApps) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *Windows10AssociatedApps) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAppType() != nil {
        cast := (*m.GetAppType()).String()
        err := writer.WriteStringValue("appType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("identifier", m.GetIdentifier())
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
func (m *Windows10AssociatedApps) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAppType sets the appType property value. Windows 10 Application type.
func (m *Windows10AssociatedApps) SetAppType(value *Windows10AppType)() {
    m.appType = value
}
// SetIdentifier sets the identifier property value. Identifier.
func (m *Windows10AssociatedApps) SetIdentifier(value *string)() {
    m.identifier = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Windows10AssociatedApps) SetOdataType(value *string)() {
    m.odataType = value
}
