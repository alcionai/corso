package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LabelActionBase 
type LabelActionBase struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The name property
    name *string
    // The OdataType property
    odataType *string
}
// NewLabelActionBase instantiates a new labelActionBase and sets the default values.
func NewLabelActionBase()(*LabelActionBase) {
    m := &LabelActionBase{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateLabelActionBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLabelActionBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.addFooter":
                        return NewAddFooter(), nil
                    case "#microsoft.graph.addHeader":
                        return NewAddHeader(), nil
                    case "#microsoft.graph.addWatermark":
                        return NewAddWatermark(), nil
                    case "#microsoft.graph.encryptContent":
                        return NewEncryptContent(), nil
                    case "#microsoft.graph.encryptWithTemplate":
                        return NewEncryptWithTemplate(), nil
                    case "#microsoft.graph.encryptWithUserDefinedRights":
                        return NewEncryptWithUserDefinedRights(), nil
                    case "#microsoft.graph.markContent":
                        return NewMarkContent(), nil
                    case "#microsoft.graph.protectGroup":
                        return NewProtectGroup(), nil
                    case "#microsoft.graph.protectOnlineMeetingAction":
                        return NewProtectOnlineMeetingAction(), nil
                    case "#microsoft.graph.protectSite":
                        return NewProtectSite(), nil
                }
            }
        }
    }
    return NewLabelActionBase(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LabelActionBase) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LabelActionBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
// GetName gets the name property value. The name property
func (m *LabelActionBase) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *LabelActionBase) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *LabelActionBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *LabelActionBase) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetName sets the name property value. The name property
func (m *LabelActionBase) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *LabelActionBase) SetOdataType(value *string)() {
    m.odataType = value
}
