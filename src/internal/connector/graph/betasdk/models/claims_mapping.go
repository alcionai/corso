package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ClaimsMapping 
type ClaimsMapping struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The claim that provides the display name or full name for the user. It is a required propoerty.
    displayName *string
    // The claim that provides the email address of the user.
    email *string
    // The claim that provides the first name of the user.
    givenName *string
    // The OdataType property
    odataType *string
    // The claim that provides the last name of the user.
    surname *string
    // The claim that provides the unique identifier for the signed-in user. It is a required propoerty.
    userId *string
}
// NewClaimsMapping instantiates a new claimsMapping and sets the default values.
func NewClaimsMapping()(*ClaimsMapping) {
    m := &ClaimsMapping{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateClaimsMappingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateClaimsMappingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClaimsMapping(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClaimsMapping) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. The claim that provides the display name or full name for the user. It is a required propoerty.
func (m *ClaimsMapping) GetDisplayName()(*string) {
    return m.displayName
}
// GetEmail gets the email property value. The claim that provides the email address of the user.
func (m *ClaimsMapping) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClaimsMapping) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["givenName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGivenName(val)
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
    res["surname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSurname(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetGivenName gets the givenName property value. The claim that provides the first name of the user.
func (m *ClaimsMapping) GetGivenName()(*string) {
    return m.givenName
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ClaimsMapping) GetOdataType()(*string) {
    return m.odataType
}
// GetSurname gets the surname property value. The claim that provides the last name of the user.
func (m *ClaimsMapping) GetSurname()(*string) {
    return m.surname
}
// GetUserId gets the userId property value. The claim that provides the unique identifier for the signed-in user. It is a required propoerty.
func (m *ClaimsMapping) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *ClaimsMapping) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("givenName", m.GetGivenName())
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
        err := writer.WriteStringValue("surname", m.GetSurname())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userId", m.GetUserId())
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
func (m *ClaimsMapping) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. The claim that provides the display name or full name for the user. It is a required propoerty.
func (m *ClaimsMapping) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEmail sets the email property value. The claim that provides the email address of the user.
func (m *ClaimsMapping) SetEmail(value *string)() {
    m.email = value
}
// SetGivenName sets the givenName property value. The claim that provides the first name of the user.
func (m *ClaimsMapping) SetGivenName(value *string)() {
    m.givenName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ClaimsMapping) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSurname sets the surname property value. The claim that provides the last name of the user.
func (m *ClaimsMapping) SetSurname(value *string)() {
    m.surname = value
}
// SetUserId sets the userId property value. The claim that provides the unique identifier for the signed-in user. It is a required propoerty.
func (m *ClaimsMapping) SetUserId(value *string)() {
    m.userId = value
}
