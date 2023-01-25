package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationStrength 
type AuthenticationStrength struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Identifier of the authentication strength.
    authenticationStrengthId *string
    // The name of the authentication strength.
    displayName *string
    // The OdataType property
    odataType *string
}
// NewAuthenticationStrength instantiates a new authenticationStrength and sets the default values.
func NewAuthenticationStrength()(*AuthenticationStrength) {
    m := &AuthenticationStrength{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAuthenticationStrengthFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationStrengthFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationStrength(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationStrength) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAuthenticationStrengthId gets the authenticationStrengthId property value. Identifier of the authentication strength.
func (m *AuthenticationStrength) GetAuthenticationStrengthId()(*string) {
    return m.authenticationStrengthId
}
// GetDisplayName gets the displayName property value. The name of the authentication strength.
func (m *AuthenticationStrength) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationStrength) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["authenticationStrengthId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationStrengthId(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
func (m *AuthenticationStrength) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AuthenticationStrength) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("authenticationStrengthId", m.GetAuthenticationStrengthId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
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
func (m *AuthenticationStrength) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAuthenticationStrengthId sets the authenticationStrengthId property value. Identifier of the authentication strength.
func (m *AuthenticationStrength) SetAuthenticationStrengthId(value *string)() {
    m.authenticationStrengthId = value
}
// SetDisplayName sets the displayName property value. The name of the authentication strength.
func (m *AuthenticationStrength) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AuthenticationStrength) SetOdataType(value *string)() {
    m.odataType = value
}
