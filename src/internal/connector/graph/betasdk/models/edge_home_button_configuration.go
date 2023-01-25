package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdgeHomeButtonConfiguration the home button configuration base class used to identify the available options
type EdgeHomeButtonConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
}
// NewEdgeHomeButtonConfiguration instantiates a new edgeHomeButtonConfiguration and sets the default values.
func NewEdgeHomeButtonConfiguration()(*EdgeHomeButtonConfiguration) {
    m := &EdgeHomeButtonConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEdgeHomeButtonConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdgeHomeButtonConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.edgeHomeButtonHidden":
                        return NewEdgeHomeButtonHidden(), nil
                    case "#microsoft.graph.edgeHomeButtonLoadsStartPage":
                        return NewEdgeHomeButtonLoadsStartPage(), nil
                    case "#microsoft.graph.edgeHomeButtonOpensCustomURL":
                        return NewEdgeHomeButtonOpensCustomURL(), nil
                    case "#microsoft.graph.edgeHomeButtonOpensNewTab":
                        return NewEdgeHomeButtonOpensNewTab(), nil
                }
            }
        }
    }
    return NewEdgeHomeButtonConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EdgeHomeButtonConfiguration) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdgeHomeButtonConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EdgeHomeButtonConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *EdgeHomeButtonConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *EdgeHomeButtonConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EdgeHomeButtonConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
