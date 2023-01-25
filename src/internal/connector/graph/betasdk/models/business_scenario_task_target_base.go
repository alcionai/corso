package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BusinessScenarioTaskTargetBase 
type BusinessScenarioTaskTargetBase struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // The taskTargetKind property
    taskTargetKind *PlannerTaskTargetKind
}
// NewBusinessScenarioTaskTargetBase instantiates a new businessScenarioTaskTargetBase and sets the default values.
func NewBusinessScenarioTaskTargetBase()(*BusinessScenarioTaskTargetBase) {
    m := &BusinessScenarioTaskTargetBase{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateBusinessScenarioTaskTargetBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBusinessScenarioTaskTargetBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.businessScenarioGroupTarget":
                        return NewBusinessScenarioGroupTarget(), nil
                }
            }
        }
    }
    return NewBusinessScenarioTaskTargetBase(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BusinessScenarioTaskTargetBase) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BusinessScenarioTaskTargetBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["taskTargetKind"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePlannerTaskTargetKind)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaskTargetKind(val.(*PlannerTaskTargetKind))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *BusinessScenarioTaskTargetBase) GetOdataType()(*string) {
    return m.odataType
}
// GetTaskTargetKind gets the taskTargetKind property value. The taskTargetKind property
func (m *BusinessScenarioTaskTargetBase) GetTaskTargetKind()(*PlannerTaskTargetKind) {
    return m.taskTargetKind
}
// Serialize serializes information the current object
func (m *BusinessScenarioTaskTargetBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetTaskTargetKind() != nil {
        cast := (*m.GetTaskTargetKind()).String()
        err := writer.WriteStringValue("taskTargetKind", &cast)
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
func (m *BusinessScenarioTaskTargetBase) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BusinessScenarioTaskTargetBase) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTaskTargetKind sets the taskTargetKind property value. The taskTargetKind property
func (m *BusinessScenarioTaskTargetBase) SetTaskTargetKind(value *PlannerTaskTargetKind)() {
    m.taskTargetKind = value
}
