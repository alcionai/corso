package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PasswordSingleSignOnSettings 
type PasswordSingleSignOnSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The fields property
    fields []PasswordSingleSignOnFieldable
    // The OdataType property
    odataType *string
}
// NewPasswordSingleSignOnSettings instantiates a new passwordSingleSignOnSettings and sets the default values.
func NewPasswordSingleSignOnSettings()(*PasswordSingleSignOnSettings) {
    m := &PasswordSingleSignOnSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePasswordSingleSignOnSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePasswordSingleSignOnSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPasswordSingleSignOnSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PasswordSingleSignOnSettings) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PasswordSingleSignOnSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["fields"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePasswordSingleSignOnFieldFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PasswordSingleSignOnFieldable, len(val))
            for i, v := range val {
                res[i] = v.(PasswordSingleSignOnFieldable)
            }
            m.SetFields(res)
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
// GetFields gets the fields property value. The fields property
func (m *PasswordSingleSignOnSettings) GetFields()([]PasswordSingleSignOnFieldable) {
    return m.fields
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PasswordSingleSignOnSettings) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *PasswordSingleSignOnSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetFields() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFields()))
        for i, v := range m.GetFields() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("fields", cast)
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
func (m *PasswordSingleSignOnSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetFields sets the fields property value. The fields property
func (m *PasswordSingleSignOnSettings) SetFields(value []PasswordSingleSignOnFieldable)() {
    m.fields = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PasswordSingleSignOnSettings) SetOdataType(value *string)() {
    m.odataType = value
}
