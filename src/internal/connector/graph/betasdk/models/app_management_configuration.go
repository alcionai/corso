package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppManagementConfiguration 
type AppManagementConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Collection of keyCredential restrictions settings to be applied to an application or service principal.
    keyCredentials []KeyCredentialConfigurationable
    // The OdataType property
    odataType *string
    // Collection of password restrictions settings to be applied to an application or service principal.
    passwordCredentials []PasswordCredentialConfigurationable
}
// NewAppManagementConfiguration instantiates a new appManagementConfiguration and sets the default values.
func NewAppManagementConfiguration()(*AppManagementConfiguration) {
    m := &AppManagementConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAppManagementConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppManagementConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppManagementConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppManagementConfiguration) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppManagementConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["keyCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyCredentialConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyCredentialConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(KeyCredentialConfigurationable)
            }
            m.SetKeyCredentials(res)
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
    res["passwordCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePasswordCredentialConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PasswordCredentialConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(PasswordCredentialConfigurationable)
            }
            m.SetPasswordCredentials(res)
        }
        return nil
    }
    return res
}
// GetKeyCredentials gets the keyCredentials property value. Collection of keyCredential restrictions settings to be applied to an application or service principal.
func (m *AppManagementConfiguration) GetKeyCredentials()([]KeyCredentialConfigurationable) {
    return m.keyCredentials
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AppManagementConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetPasswordCredentials gets the passwordCredentials property value. Collection of password restrictions settings to be applied to an application or service principal.
func (m *AppManagementConfiguration) GetPasswordCredentials()([]PasswordCredentialConfigurationable) {
    return m.passwordCredentials
}
// Serialize serializes information the current object
func (m *AppManagementConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetKeyCredentials() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetKeyCredentials()))
        for i, v := range m.GetKeyCredentials() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("keyCredentials", cast)
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
    if m.GetPasswordCredentials() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPasswordCredentials()))
        for i, v := range m.GetPasswordCredentials() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("passwordCredentials", cast)
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
func (m *AppManagementConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetKeyCredentials sets the keyCredentials property value. Collection of keyCredential restrictions settings to be applied to an application or service principal.
func (m *AppManagementConfiguration) SetKeyCredentials(value []KeyCredentialConfigurationable)() {
    m.keyCredentials = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AppManagementConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPasswordCredentials sets the passwordCredentials property value. Collection of password restrictions settings to be applied to an application or service principal.
func (m *AppManagementConfiguration) SetPasswordCredentials(value []PasswordCredentialConfigurationable)() {
    m.passwordCredentials = value
}
