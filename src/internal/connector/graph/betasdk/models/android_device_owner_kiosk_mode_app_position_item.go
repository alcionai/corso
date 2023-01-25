package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerKioskModeAppPositionItem an item in the list of app positions that sets the order of items on the Managed Home Screen
type AndroidDeviceOwnerKioskModeAppPositionItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Represents an item on the Android Device Owner Managed Home Screen (application, weblink or folder
    item AndroidDeviceOwnerKioskModeHomeScreenItemable
    // The OdataType property
    odataType *string
    // Position of the item on the grid. Valid values 0 to 9999999
    position *int32
}
// NewAndroidDeviceOwnerKioskModeAppPositionItem instantiates a new androidDeviceOwnerKioskModeAppPositionItem and sets the default values.
func NewAndroidDeviceOwnerKioskModeAppPositionItem()(*AndroidDeviceOwnerKioskModeAppPositionItem) {
    m := &AndroidDeviceOwnerKioskModeAppPositionItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAndroidDeviceOwnerKioskModeAppPositionItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerKioskModeAppPositionItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerKioskModeAppPositionItem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidDeviceOwnerKioskModeAppPositionItem) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerKioskModeAppPositionItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["item"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidDeviceOwnerKioskModeHomeScreenItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItem(val.(AndroidDeviceOwnerKioskModeHomeScreenItemable))
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
    res["position"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPosition(val)
        }
        return nil
    }
    return res
}
// GetItem gets the item property value. Represents an item on the Android Device Owner Managed Home Screen (application, weblink or folder
func (m *AndroidDeviceOwnerKioskModeAppPositionItem) GetItem()(AndroidDeviceOwnerKioskModeHomeScreenItemable) {
    return m.item
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AndroidDeviceOwnerKioskModeAppPositionItem) GetOdataType()(*string) {
    return m.odataType
}
// GetPosition gets the position property value. Position of the item on the grid. Valid values 0 to 9999999
func (m *AndroidDeviceOwnerKioskModeAppPositionItem) GetPosition()(*int32) {
    return m.position
}
// Serialize serializes information the current object
func (m *AndroidDeviceOwnerKioskModeAppPositionItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("item", m.GetItem())
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
        err := writer.WriteInt32Value("position", m.GetPosition())
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
func (m *AndroidDeviceOwnerKioskModeAppPositionItem) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetItem sets the item property value. Represents an item on the Android Device Owner Managed Home Screen (application, weblink or folder
func (m *AndroidDeviceOwnerKioskModeAppPositionItem) SetItem(value AndroidDeviceOwnerKioskModeHomeScreenItemable)() {
    m.item = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AndroidDeviceOwnerKioskModeAppPositionItem) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPosition sets the position property value. Position of the item on the grid. Valid values 0 to 9999999
func (m *AndroidDeviceOwnerKioskModeAppPositionItem) SetPosition(value *int32)() {
    m.position = value
}
