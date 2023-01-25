package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PasswordSingleSignOnField 
type PasswordSingleSignOnField struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Title/label override for customization.
    customizedLabel *string
    // Label that would be used if no customizedLabel is provided. Read only.
    defaultLabel *string
    // Id used to identity the field type. This is an internal id and possible values are param_1, param_2, param_userName, param_password.
    fieldId *string
    // The OdataType property
    odataType *string
    // Type of the credential. The values can be text, password.
    type_escaped *string
}
// NewPasswordSingleSignOnField instantiates a new passwordSingleSignOnField and sets the default values.
func NewPasswordSingleSignOnField()(*PasswordSingleSignOnField) {
    m := &PasswordSingleSignOnField{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePasswordSingleSignOnFieldFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePasswordSingleSignOnFieldFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPasswordSingleSignOnField(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PasswordSingleSignOnField) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCustomizedLabel gets the customizedLabel property value. Title/label override for customization.
func (m *PasswordSingleSignOnField) GetCustomizedLabel()(*string) {
    return m.customizedLabel
}
// GetDefaultLabel gets the defaultLabel property value. Label that would be used if no customizedLabel is provided. Read only.
func (m *PasswordSingleSignOnField) GetDefaultLabel()(*string) {
    return m.defaultLabel
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PasswordSingleSignOnField) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["customizedLabel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomizedLabel(val)
        }
        return nil
    }
    res["defaultLabel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultLabel(val)
        }
        return nil
    }
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
    return res
}
// GetFieldId gets the fieldId property value. Id used to identity the field type. This is an internal id and possible values are param_1, param_2, param_userName, param_password.
func (m *PasswordSingleSignOnField) GetFieldId()(*string) {
    return m.fieldId
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PasswordSingleSignOnField) GetOdataType()(*string) {
    return m.odataType
}
// GetType gets the type property value. Type of the credential. The values can be text, password.
func (m *PasswordSingleSignOnField) GetType()(*string) {
    return m.type_escaped
}
// Serialize serializes information the current object
func (m *PasswordSingleSignOnField) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("customizedLabel", m.GetCustomizedLabel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("defaultLabel", m.GetDefaultLabel())
        if err != nil {
            return err
        }
    }
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PasswordSingleSignOnField) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCustomizedLabel sets the customizedLabel property value. Title/label override for customization.
func (m *PasswordSingleSignOnField) SetCustomizedLabel(value *string)() {
    m.customizedLabel = value
}
// SetDefaultLabel sets the defaultLabel property value. Label that would be used if no customizedLabel is provided. Read only.
func (m *PasswordSingleSignOnField) SetDefaultLabel(value *string)() {
    m.defaultLabel = value
}
// SetFieldId sets the fieldId property value. Id used to identity the field type. This is an internal id and possible values are param_1, param_2, param_userName, param_password.
func (m *PasswordSingleSignOnField) SetFieldId(value *string)() {
    m.fieldId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PasswordSingleSignOnField) SetOdataType(value *string)() {
    m.odataType = value
}
// SetType sets the type property value. Type of the credential. The values can be text, password.
func (m *PasswordSingleSignOnField) SetType(value *string)() {
    m.type_escaped = value
}
