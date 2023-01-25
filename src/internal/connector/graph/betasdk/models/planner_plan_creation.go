package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerPlanCreation 
type PlannerPlanCreation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Specifies what kind of creation source the plan is created with. The possible values are: external, publication and unknownFutureValue.
    creationSourceKind *PlannerCreationSourceKind
    // The OdataType property
    odataType *string
}
// NewPlannerPlanCreation instantiates a new plannerPlanCreation and sets the default values.
func NewPlannerPlanCreation()(*PlannerPlanCreation) {
    m := &PlannerPlanCreation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePlannerPlanCreationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerPlanCreationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.plannerExternalPlanSource":
                        return NewPlannerExternalPlanSource(), nil
                }
            }
        }
    }
    return NewPlannerPlanCreation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerPlanCreation) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCreationSourceKind gets the creationSourceKind property value. Specifies what kind of creation source the plan is created with. The possible values are: external, publication and unknownFutureValue.
func (m *PlannerPlanCreation) GetCreationSourceKind()(*PlannerCreationSourceKind) {
    return m.creationSourceKind
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerPlanCreation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["creationSourceKind"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePlannerCreationSourceKind)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreationSourceKind(val.(*PlannerCreationSourceKind))
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
func (m *PlannerPlanCreation) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *PlannerPlanCreation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCreationSourceKind() != nil {
        cast := (*m.GetCreationSourceKind()).String()
        err := writer.WriteStringValue("creationSourceKind", &cast)
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
func (m *PlannerPlanCreation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCreationSourceKind sets the creationSourceKind property value. Specifies what kind of creation source the plan is created with. The possible values are: external, publication and unknownFutureValue.
func (m *PlannerPlanCreation) SetCreationSourceKind(value *PlannerCreationSourceKind)() {
    m.creationSourceKind = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PlannerPlanCreation) SetOdataType(value *string)() {
    m.odataType = value
}
