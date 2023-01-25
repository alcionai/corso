package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PostalAddressType 
type PostalAddressType struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The city property
    city *string
    // The countryLetterCode property
    countryLetterCode *string
    // The OdataType property
    odataType *string
    // The postalCode property
    postalCode *string
    // The state property
    state *string
    // The street property
    street *string
}
// NewPostalAddressType instantiates a new postalAddressType and sets the default values.
func NewPostalAddressType()(*PostalAddressType) {
    m := &PostalAddressType{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePostalAddressTypeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePostalAddressTypeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPostalAddressType(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PostalAddressType) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCity gets the city property value. The city property
func (m *PostalAddressType) GetCity()(*string) {
    return m.city
}
// GetCountryLetterCode gets the countryLetterCode property value. The countryLetterCode property
func (m *PostalAddressType) GetCountryLetterCode()(*string) {
    return m.countryLetterCode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PostalAddressType) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["city"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCity(val)
        }
        return nil
    }
    res["countryLetterCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryLetterCode(val)
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
    res["postalCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostalCode(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    res["street"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStreet(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PostalAddressType) GetOdataType()(*string) {
    return m.odataType
}
// GetPostalCode gets the postalCode property value. The postalCode property
func (m *PostalAddressType) GetPostalCode()(*string) {
    return m.postalCode
}
// GetState gets the state property value. The state property
func (m *PostalAddressType) GetState()(*string) {
    return m.state
}
// GetStreet gets the street property value. The street property
func (m *PostalAddressType) GetStreet()(*string) {
    return m.street
}
// Serialize serializes information the current object
func (m *PostalAddressType) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("city", m.GetCity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("countryLetterCode", m.GetCountryLetterCode())
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
        err := writer.WriteStringValue("postalCode", m.GetPostalCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("street", m.GetStreet())
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
func (m *PostalAddressType) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCity sets the city property value. The city property
func (m *PostalAddressType) SetCity(value *string)() {
    m.city = value
}
// SetCountryLetterCode sets the countryLetterCode property value. The countryLetterCode property
func (m *PostalAddressType) SetCountryLetterCode(value *string)() {
    m.countryLetterCode = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PostalAddressType) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPostalCode sets the postalCode property value. The postalCode property
func (m *PostalAddressType) SetPostalCode(value *string)() {
    m.postalCode = value
}
// SetState sets the state property value. The state property
func (m *PostalAddressType) SetState(value *string)() {
    m.state = value
}
// SetStreet sets the street property value. The street property
func (m *PostalAddressType) SetStreet(value *string)() {
    m.street = value
}
