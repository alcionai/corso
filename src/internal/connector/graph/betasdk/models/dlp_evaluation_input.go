package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DlpEvaluationInput 
type DlpEvaluationInput struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The currentLabel property
    currentLabel CurrentLabelable
    // The discoveredSensitiveTypes property
    discoveredSensitiveTypes []DiscoveredSensitiveTypeable
    // The OdataType property
    odataType *string
}
// NewDlpEvaluationInput instantiates a new dlpEvaluationInput and sets the default values.
func NewDlpEvaluationInput()(*DlpEvaluationInput) {
    m := &DlpEvaluationInput{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDlpEvaluationInputFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDlpEvaluationInputFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.dlpEvaluationWindowsDevicesInput":
                        return NewDlpEvaluationWindowsDevicesInput(), nil
                }
            }
        }
    }
    return NewDlpEvaluationInput(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DlpEvaluationInput) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCurrentLabel gets the currentLabel property value. The currentLabel property
func (m *DlpEvaluationInput) GetCurrentLabel()(CurrentLabelable) {
    return m.currentLabel
}
// GetDiscoveredSensitiveTypes gets the discoveredSensitiveTypes property value. The discoveredSensitiveTypes property
func (m *DlpEvaluationInput) GetDiscoveredSensitiveTypes()([]DiscoveredSensitiveTypeable) {
    return m.discoveredSensitiveTypes
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DlpEvaluationInput) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["currentLabel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCurrentLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentLabel(val.(CurrentLabelable))
        }
        return nil
    }
    res["discoveredSensitiveTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDiscoveredSensitiveTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DiscoveredSensitiveTypeable, len(val))
            for i, v := range val {
                res[i] = v.(DiscoveredSensitiveTypeable)
            }
            m.SetDiscoveredSensitiveTypes(res)
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
func (m *DlpEvaluationInput) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *DlpEvaluationInput) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("currentLabel", m.GetCurrentLabel())
        if err != nil {
            return err
        }
    }
    if m.GetDiscoveredSensitiveTypes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDiscoveredSensitiveTypes()))
        for i, v := range m.GetDiscoveredSensitiveTypes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("discoveredSensitiveTypes", cast)
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
func (m *DlpEvaluationInput) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCurrentLabel sets the currentLabel property value. The currentLabel property
func (m *DlpEvaluationInput) SetCurrentLabel(value CurrentLabelable)() {
    m.currentLabel = value
}
// SetDiscoveredSensitiveTypes sets the discoveredSensitiveTypes property value. The discoveredSensitiveTypes property
func (m *DlpEvaluationInput) SetDiscoveredSensitiveTypes(value []DiscoveredSensitiveTypeable)() {
    m.discoveredSensitiveTypes = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DlpEvaluationInput) SetOdataType(value *string)() {
    m.odataType = value
}
