package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeleteAction 
type DeleteAction struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The name of the item that was deleted.
    name *string
    // File or Folder, depending on the type of the deleted item.
    objectType *string
    // The OdataType property
    odataType *string
}
// NewDeleteAction instantiates a new deleteAction and sets the default values.
func NewDeleteAction()(*DeleteAction) {
    m := &DeleteAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeleteActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeleteActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeleteAction(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeleteAction) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeleteAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["objectType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObjectType(val)
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
// GetName gets the name property value. The name of the item that was deleted.
func (m *DeleteAction) GetName()(*string) {
    return m.name
}
// GetObjectType gets the objectType property value. File or Folder, depending on the type of the deleted item.
func (m *DeleteAction) GetObjectType()(*string) {
    return m.objectType
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeleteAction) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *DeleteAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("objectType", m.GetObjectType())
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
func (m *DeleteAction) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetName sets the name property value. The name of the item that was deleted.
func (m *DeleteAction) SetName(value *string)() {
    m.name = value
}
// SetObjectType sets the objectType property value. File or Folder, depending on the type of the deleted item.
func (m *DeleteAction) SetObjectType(value *string)() {
    m.objectType = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeleteAction) SetOdataType(value *string)() {
    m.odataType = value
}
