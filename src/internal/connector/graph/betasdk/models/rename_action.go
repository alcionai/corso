package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RenameAction 
type RenameAction struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The new name of the item.
    newName *string
    // The OdataType property
    odataType *string
    // The previous name of the item.
    oldName *string
}
// NewRenameAction instantiates a new renameAction and sets the default values.
func NewRenameAction()(*RenameAction) {
    m := &RenameAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRenameActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRenameActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRenameAction(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RenameAction) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RenameAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["newName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewName(val)
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
    res["oldName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOldName(val)
        }
        return nil
    }
    return res
}
// GetNewName gets the newName property value. The new name of the item.
func (m *RenameAction) GetNewName()(*string) {
    return m.newName
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RenameAction) GetOdataType()(*string) {
    return m.odataType
}
// GetOldName gets the oldName property value. The previous name of the item.
func (m *RenameAction) GetOldName()(*string) {
    return m.oldName
}
// Serialize serializes information the current object
func (m *RenameAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("newName", m.GetNewName())
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
        err := writer.WriteStringValue("oldName", m.GetOldName())
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
func (m *RenameAction) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetNewName sets the newName property value. The new name of the item.
func (m *RenameAction) SetNewName(value *string)() {
    m.newName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RenameAction) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOldName sets the oldName property value. The previous name of the item.
func (m *RenameAction) SetOldName(value *string)() {
    m.oldName = value
}
