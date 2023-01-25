package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MfaDetail 
type MfaDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Indicates the MFA auth detail for the corresponding Sign-in activity when the MFA Required is 'Yes'.
    authDetail *string
    // Indicates the MFA Auth methods (SMS, Phone, Authenticator App are some of the value) for the corresponding sign-in activity when the MFA Required field is 'Yes'.
    authMethod *string
    // The OdataType property
    odataType *string
}
// NewMfaDetail instantiates a new mfaDetail and sets the default values.
func NewMfaDetail()(*MfaDetail) {
    m := &MfaDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMfaDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMfaDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMfaDetail(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MfaDetail) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAuthDetail gets the authDetail property value. Indicates the MFA auth detail for the corresponding Sign-in activity when the MFA Required is 'Yes'.
func (m *MfaDetail) GetAuthDetail()(*string) {
    return m.authDetail
}
// GetAuthMethod gets the authMethod property value. Indicates the MFA Auth methods (SMS, Phone, Authenticator App are some of the value) for the corresponding sign-in activity when the MFA Required field is 'Yes'.
func (m *MfaDetail) GetAuthMethod()(*string) {
    return m.authMethod
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MfaDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["authDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthDetail(val)
        }
        return nil
    }
    res["authMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthMethod(val)
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MfaDetail) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *MfaDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("authDetail", m.GetAuthDetail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("authMethod", m.GetAuthMethod())
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
func (m *MfaDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAuthDetail sets the authDetail property value. Indicates the MFA auth detail for the corresponding Sign-in activity when the MFA Required is 'Yes'.
func (m *MfaDetail) SetAuthDetail(value *string)() {
    m.authDetail = value
}
// SetAuthMethod sets the authMethod property value. Indicates the MFA Auth methods (SMS, Phone, Authenticator App are some of the value) for the corresponding sign-in activity when the MFA Required field is 'Yes'.
func (m *MfaDetail) SetAuthMethod(value *string)() {
    m.authMethod = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MfaDetail) SetOdataType(value *string)() {
    m.odataType = value
}
