package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSLaunchItem represents an app in the list of macOS launch items
type MacOSLaunchItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Whether or not to hide the item from the Users and Groups List.
    hide *bool
    // The OdataType property
    odataType *string
    // Path to the launch item.
    path *string
}
// NewMacOSLaunchItem instantiates a new macOSLaunchItem and sets the default values.
func NewMacOSLaunchItem()(*MacOSLaunchItem) {
    m := &MacOSLaunchItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMacOSLaunchItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSLaunchItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSLaunchItem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MacOSLaunchItem) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSLaunchItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["hide"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHide(val)
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
    res["path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPath(val)
        }
        return nil
    }
    return res
}
// GetHide gets the hide property value. Whether or not to hide the item from the Users and Groups List.
func (m *MacOSLaunchItem) GetHide()(*bool) {
    return m.hide
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MacOSLaunchItem) GetOdataType()(*string) {
    return m.odataType
}
// GetPath gets the path property value. Path to the launch item.
func (m *MacOSLaunchItem) GetPath()(*string) {
    return m.path
}
// Serialize serializes information the current object
func (m *MacOSLaunchItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("hide", m.GetHide())
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
        err := writer.WriteStringValue("path", m.GetPath())
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
func (m *MacOSLaunchItem) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetHide sets the hide property value. Whether or not to hide the item from the Users and Groups List.
func (m *MacOSLaunchItem) SetHide(value *bool)() {
    m.hide = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MacOSLaunchItem) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPath sets the path property value. Path to the launch item.
func (m *MacOSLaunchItem) SetPath(value *string)() {
    m.path = value
}
