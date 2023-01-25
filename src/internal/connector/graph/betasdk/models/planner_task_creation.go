package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerTaskCreation 
type PlannerTaskCreation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Specifies what kind of creation source the task is created with. The possible values are: external, publication and unknownFutureValue.
    creationSourceKind *PlannerCreationSourceKind
    // The OdataType property
    odataType *string
    // Information about the publication process that created this task. This field is deprecated and clients should move to using the new inheritance model.
    teamsPublicationInfo PlannerTeamsPublicationInfoable
}
// NewPlannerTaskCreation instantiates a new plannerTaskCreation and sets the default values.
func NewPlannerTaskCreation()(*PlannerTaskCreation) {
    m := &PlannerTaskCreation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePlannerTaskCreationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerTaskCreationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.plannerExternalTaskSource":
                        return NewPlannerExternalTaskSource(), nil
                    case "#microsoft.graph.plannerTeamsPublicationInfo":
                        return NewPlannerTeamsPublicationInfo(), nil
                }
            }
        }
    }
    return NewPlannerTaskCreation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerTaskCreation) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCreationSourceKind gets the creationSourceKind property value. Specifies what kind of creation source the task is created with. The possible values are: external, publication and unknownFutureValue.
func (m *PlannerTaskCreation) GetCreationSourceKind()(*PlannerCreationSourceKind) {
    return m.creationSourceKind
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerTaskCreation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["teamsPublicationInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerTeamsPublicationInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsPublicationInfo(val.(PlannerTeamsPublicationInfoable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PlannerTaskCreation) GetOdataType()(*string) {
    return m.odataType
}
// GetTeamsPublicationInfo gets the teamsPublicationInfo property value. Information about the publication process that created this task. This field is deprecated and clients should move to using the new inheritance model.
func (m *PlannerTaskCreation) GetTeamsPublicationInfo()(PlannerTeamsPublicationInfoable) {
    return m.teamsPublicationInfo
}
// Serialize serializes information the current object
func (m *PlannerTaskCreation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteObjectValue("teamsPublicationInfo", m.GetTeamsPublicationInfo())
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
func (m *PlannerTaskCreation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCreationSourceKind sets the creationSourceKind property value. Specifies what kind of creation source the task is created with. The possible values are: external, publication and unknownFutureValue.
func (m *PlannerTaskCreation) SetCreationSourceKind(value *PlannerCreationSourceKind)() {
    m.creationSourceKind = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PlannerTaskCreation) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTeamsPublicationInfo sets the teamsPublicationInfo property value. Information about the publication process that created this task. This field is deprecated and clients should move to using the new inheritance model.
func (m *PlannerTaskCreation) SetTeamsPublicationInfo(value PlannerTeamsPublicationInfoable)() {
    m.teamsPublicationInfo = value
}
