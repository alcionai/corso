package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AlternativeSecurity 
type AlternativeSecurity struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // For internal use only
    identityProvider *string
    // For internal use only
    key []byte
    // The OdataType property
    odataType *string
    // For internal use only
    type_escaped *int32
}
// NewAlternativeSecurity instantiates a new AlternativeSecurity and sets the default values.
func NewAlternativeSecurity()(*AlternativeSecurity) {
    m := &AlternativeSecurity{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAlternativeSecurityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAlternativeSecurityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAlternativeSecurity(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AlternativeSecurity) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AlternativeSecurity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["identityProvider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityProvider(val)
        }
        return nil
    }
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
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
        val, err := n.GetInt32Value()
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
// GetIdentityProvider gets the identityProvider property value. For internal use only
func (m *AlternativeSecurity) GetIdentityProvider()(*string) {
    return m.identityProvider
}
// GetKey gets the key property value. For internal use only
func (m *AlternativeSecurity) GetKey()([]byte) {
    return m.key
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AlternativeSecurity) GetOdataType()(*string) {
    return m.odataType
}
// GetType gets the type property value. For internal use only
func (m *AlternativeSecurity) GetType()(*int32) {
    return m.type_escaped
}
// Serialize serializes information the current object
func (m *AlternativeSecurity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("identityProvider", m.GetIdentityProvider())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteByteArrayValue("key", m.GetKey())
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
        err := writer.WriteInt32Value("type", m.GetType())
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
func (m *AlternativeSecurity) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetIdentityProvider sets the identityProvider property value. For internal use only
func (m *AlternativeSecurity) SetIdentityProvider(value *string)() {
    m.identityProvider = value
}
// SetKey sets the key property value. For internal use only
func (m *AlternativeSecurity) SetKey(value []byte)() {
    m.key = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AlternativeSecurity) SetOdataType(value *string)() {
    m.odataType = value
}
// SetType sets the type property value. For internal use only
func (m *AlternativeSecurity) SetType(value *int32)() {
    m.type_escaped = value
}
