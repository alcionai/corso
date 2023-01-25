package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccountTargetContent 
type AccountTargetContent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // The type of account target content. Possible values are: unknown,includeAll, addressBook,  unknownFutureValue.
    type_escaped *AccountTargetContentType
}
// NewAccountTargetContent instantiates a new accountTargetContent and sets the default values.
func NewAccountTargetContent()(*AccountTargetContent) {
    m := &AccountTargetContent{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAccountTargetContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccountTargetContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.addressBookAccountTargetContent":
                        return NewAddressBookAccountTargetContent(), nil
                    case "#microsoft.graph.includeAllAccountTargetContent":
                        return NewIncludeAllAccountTargetContent(), nil
                }
            }
        }
    }
    return NewAccountTargetContent(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccountTargetContent) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccountTargetContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
        val, err := n.GetEnumValue(ParseAccountTargetContentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*AccountTargetContentType))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AccountTargetContent) GetOdataType()(*string) {
    return m.odataType
}
// GetType gets the type property value. The type of account target content. Possible values are: unknown,includeAll, addressBook,  unknownFutureValue.
func (m *AccountTargetContent) GetType()(*AccountTargetContentType) {
    return m.type_escaped
}
// Serialize serializes information the current object
func (m *AccountTargetContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *AccountTargetContent) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AccountTargetContent) SetOdataType(value *string)() {
    m.odataType = value
}
// SetType sets the type property value. The type of account target content. Possible values are: unknown,includeAll, addressBook,  unknownFutureValue.
func (m *AccountTargetContent) SetType(value *AccountTargetContentType)() {
    m.type_escaped = value
}
