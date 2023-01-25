package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidPermissionAction mapping between an Android app permission and the action Android should take when that permission is requested.
type AndroidPermissionAction struct {
    // Android action taken when an app requests a dangerous permission.
    action *AndroidPermissionActionType
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // Android permission string, defined in the official Android documentation.  Example 'android.permission.READ_CONTACTS'.
    permission *string
}
// NewAndroidPermissionAction instantiates a new androidPermissionAction and sets the default values.
func NewAndroidPermissionAction()(*AndroidPermissionAction) {
    m := &AndroidPermissionAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAndroidPermissionActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidPermissionActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidPermissionAction(), nil
}
// GetAction gets the action property value. Android action taken when an app requests a dangerous permission.
func (m *AndroidPermissionAction) GetAction()(*AndroidPermissionActionType) {
    return m.action
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidPermissionAction) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidPermissionAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidPermissionActionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*AndroidPermissionActionType))
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
    res["permission"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermission(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AndroidPermissionAction) GetOdataType()(*string) {
    return m.odataType
}
// GetPermission gets the permission property value. Android permission string, defined in the official Android documentation.  Example 'android.permission.READ_CONTACTS'.
func (m *AndroidPermissionAction) GetPermission()(*string) {
    return m.permission
}
// Serialize serializes information the current object
func (m *AndroidPermissionAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAction() != nil {
        cast := (*m.GetAction()).String()
        err := writer.WriteStringValue("action", &cast)
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
        err := writer.WriteStringValue("permission", m.GetPermission())
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
// SetAction sets the action property value. Android action taken when an app requests a dangerous permission.
func (m *AndroidPermissionAction) SetAction(value *AndroidPermissionActionType)() {
    m.action = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidPermissionAction) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AndroidPermissionAction) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPermission sets the permission property value. Android permission string, defined in the official Android documentation.  Example 'android.permission.READ_CONTACTS'.
func (m *AndroidPermissionAction) SetPermission(value *string)() {
    m.permission = value
}
