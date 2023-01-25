package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InformationProtectionContentLabel 
type InformationProtectionContentLabel struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The assignmentMethod property
    assignmentMethod *AssignmentMethod
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    creationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Details on the label that is currently applied to the file.
    label LabelDetailsable
    // The OdataType property
    odataType *string
}
// NewInformationProtectionContentLabel instantiates a new informationProtectionContentLabel and sets the default values.
func NewInformationProtectionContentLabel()(*InformationProtectionContentLabel) {
    m := &InformationProtectionContentLabel{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateInformationProtectionContentLabelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInformationProtectionContentLabelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInformationProtectionContentLabel(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InformationProtectionContentLabel) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAssignmentMethod gets the assignmentMethod property value. The assignmentMethod property
func (m *InformationProtectionContentLabel) GetAssignmentMethod()(*AssignmentMethod) {
    return m.assignmentMethod
}
// GetCreationDateTime gets the creationDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *InformationProtectionContentLabel) GetCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.creationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InformationProtectionContentLabel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignmentMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAssignmentMethod)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentMethod(val.(*AssignmentMethod))
        }
        return nil
    }
    res["creationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreationDateTime(val)
        }
        return nil
    }
    res["label"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLabelDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabel(val.(LabelDetailsable))
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
// GetLabel gets the label property value. Details on the label that is currently applied to the file.
func (m *InformationProtectionContentLabel) GetLabel()(LabelDetailsable) {
    return m.label
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *InformationProtectionContentLabel) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *InformationProtectionContentLabel) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssignmentMethod() != nil {
        cast := (*m.GetAssignmentMethod()).String()
        err := writer.WriteStringValue("assignmentMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("creationDateTime", m.GetCreationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("label", m.GetLabel())
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
func (m *InformationProtectionContentLabel) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAssignmentMethod sets the assignmentMethod property value. The assignmentMethod property
func (m *InformationProtectionContentLabel) SetAssignmentMethod(value *AssignmentMethod)() {
    m.assignmentMethod = value
}
// SetCreationDateTime sets the creationDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *InformationProtectionContentLabel) SetCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.creationDateTime = value
}
// SetLabel sets the label property value. Details on the label that is currently applied to the file.
func (m *InformationProtectionContentLabel) SetLabel(value LabelDetailsable)() {
    m.label = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *InformationProtectionContentLabel) SetOdataType(value *string)() {
    m.odataType = value
}
