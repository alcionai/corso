package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Credential 
type Credential struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The name of the field for this credential. e.g, username or password or phoneNumber. This is defined by the application. Must match what is in the html field on singleSignOnSettings/password object.
    fieldId *string
    // The OdataType property
    odataType *string
    // The type for this credential. Valid values: username, password, or other.
    type_escaped *string
    // The value for this credential. e.g, mysuperhiddenpassword. Note the value for passwords is write-only, the value can never be read back.
    value *string
}
// NewCredential instantiates a new credential and sets the default values.
func NewCredential()(*Credential) {
    m := &Credential{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCredentialFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCredentialFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCredential(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Credential) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Credential) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["fieldId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFieldId(val)
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetFieldId gets the fieldId property value. The name of the field for this credential. e.g, username or password or phoneNumber. This is defined by the application. Must match what is in the html field on singleSignOnSettings/password object.
func (m *Credential) GetFieldId()(*string) {
    return m.fieldId
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Credential) GetOdataType()(*string) {
    return m.odataType
}
// GetType gets the type property value. The type for this credential. Valid values: username, password, or other.
func (m *Credential) GetType()(*string) {
    return m.type_escaped
}
// GetValue gets the value property value. The value for this credential. e.g, mysuperhiddenpassword. Note the value for passwords is write-only, the value can never be read back.
func (m *Credential) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *Credential) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("fieldId", m.GetFieldId())
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
        err := writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *Credential) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetFieldId sets the fieldId property value. The name of the field for this credential. e.g, username or password or phoneNumber. This is defined by the application. Must match what is in the html field on singleSignOnSettings/password object.
func (m *Credential) SetFieldId(value *string)() {
    m.fieldId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Credential) SetOdataType(value *string)() {
    m.odataType = value
}
// SetType sets the type property value. The type for this credential. Valid values: username, password, or other.
func (m *Credential) SetType(value *string)() {
    m.type_escaped = value
}
// SetValue sets the value property value. The value for this credential. e.g, mysuperhiddenpassword. Note the value for passwords is write-only, the value can never be read back.
func (m *Credential) SetValue(value *string)() {
    m.value = value
}
