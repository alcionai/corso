package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AssignmentFilterEvaluationSummary represent result summary for assignment filter evaluation
type AssignmentFilterEvaluationSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The admin defined name for assignment filter.
    assignmentFilterDisplayName *string
    // Unique identifier for the assignment filter object
    assignmentFilterId *string
    // The time the assignment filter was last modified.
    assignmentFilterLastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Supported platform types.
    assignmentFilterPlatform *DevicePlatformType
    // Represents type of the assignment filter.
    assignmentFilterType *DeviceAndAppManagementAssignmentFilterType
    // A collection of filter types and their corresponding evaluation results.
    assignmentFilterTypeAndEvaluationResults []AssignmentFilterTypeAndEvaluationResultable
    // The time assignment filter was evaluated.
    evaluationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Supported evaluation results for filter.
    evaluationResult *AssignmentFilterEvaluationResult
    // The OdataType property
    odataType *string
}
// NewAssignmentFilterEvaluationSummary instantiates a new assignmentFilterEvaluationSummary and sets the default values.
func NewAssignmentFilterEvaluationSummary()(*AssignmentFilterEvaluationSummary) {
    m := &AssignmentFilterEvaluationSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAssignmentFilterEvaluationSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssignmentFilterEvaluationSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignmentFilterEvaluationSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentFilterEvaluationSummary) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAssignmentFilterDisplayName gets the assignmentFilterDisplayName property value. The admin defined name for assignment filter.
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterDisplayName()(*string) {
    return m.assignmentFilterDisplayName
}
// GetAssignmentFilterId gets the assignmentFilterId property value. Unique identifier for the assignment filter object
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterId()(*string) {
    return m.assignmentFilterId
}
// GetAssignmentFilterLastModifiedDateTime gets the assignmentFilterLastModifiedDateTime property value. The time the assignment filter was last modified.
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.assignmentFilterLastModifiedDateTime
}
// GetAssignmentFilterPlatform gets the assignmentFilterPlatform property value. Supported platform types.
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterPlatform()(*DevicePlatformType) {
    return m.assignmentFilterPlatform
}
// GetAssignmentFilterType gets the assignmentFilterType property value. Represents type of the assignment filter.
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterType()(*DeviceAndAppManagementAssignmentFilterType) {
    return m.assignmentFilterType
}
// GetAssignmentFilterTypeAndEvaluationResults gets the assignmentFilterTypeAndEvaluationResults property value. A collection of filter types and their corresponding evaluation results.
func (m *AssignmentFilterEvaluationSummary) GetAssignmentFilterTypeAndEvaluationResults()([]AssignmentFilterTypeAndEvaluationResultable) {
    return m.assignmentFilterTypeAndEvaluationResults
}
// GetEvaluationDateTime gets the evaluationDateTime property value. The time assignment filter was evaluated.
func (m *AssignmentFilterEvaluationSummary) GetEvaluationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.evaluationDateTime
}
// GetEvaluationResult gets the evaluationResult property value. Supported evaluation results for filter.
func (m *AssignmentFilterEvaluationSummary) GetEvaluationResult()(*AssignmentFilterEvaluationResult) {
    return m.evaluationResult
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignmentFilterEvaluationSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignmentFilterDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentFilterDisplayName(val)
        }
        return nil
    }
    res["assignmentFilterId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentFilterId(val)
        }
        return nil
    }
    res["assignmentFilterLastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentFilterLastModifiedDateTime(val)
        }
        return nil
    }
    res["assignmentFilterPlatform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDevicePlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentFilterPlatform(val.(*DevicePlatformType))
        }
        return nil
    }
    res["assignmentFilterType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAndAppManagementAssignmentFilterType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentFilterType(val.(*DeviceAndAppManagementAssignmentFilterType))
        }
        return nil
    }
    res["assignmentFilterTypeAndEvaluationResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAssignmentFilterTypeAndEvaluationResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignmentFilterTypeAndEvaluationResultable, len(val))
            for i, v := range val {
                res[i] = v.(AssignmentFilterTypeAndEvaluationResultable)
            }
            m.SetAssignmentFilterTypeAndEvaluationResults(res)
        }
        return nil
    }
    res["evaluationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvaluationDateTime(val)
        }
        return nil
    }
    res["evaluationResult"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAssignmentFilterEvaluationResult)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvaluationResult(val.(*AssignmentFilterEvaluationResult))
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
func (m *AssignmentFilterEvaluationSummary) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AssignmentFilterEvaluationSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("assignmentFilterDisplayName", m.GetAssignmentFilterDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("assignmentFilterId", m.GetAssignmentFilterId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("assignmentFilterLastModifiedDateTime", m.GetAssignmentFilterLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetAssignmentFilterPlatform() != nil {
        cast := (*m.GetAssignmentFilterPlatform()).String()
        err := writer.WriteStringValue("assignmentFilterPlatform", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignmentFilterType() != nil {
        cast := (*m.GetAssignmentFilterType()).String()
        err := writer.WriteStringValue("assignmentFilterType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignmentFilterTypeAndEvaluationResults() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignmentFilterTypeAndEvaluationResults()))
        for i, v := range m.GetAssignmentFilterTypeAndEvaluationResults() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("assignmentFilterTypeAndEvaluationResults", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("evaluationDateTime", m.GetEvaluationDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetEvaluationResult() != nil {
        cast := (*m.GetEvaluationResult()).String()
        err := writer.WriteStringValue("evaluationResult", &cast)
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
func (m *AssignmentFilterEvaluationSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAssignmentFilterDisplayName sets the assignmentFilterDisplayName property value. The admin defined name for assignment filter.
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterDisplayName(value *string)() {
    m.assignmentFilterDisplayName = value
}
// SetAssignmentFilterId sets the assignmentFilterId property value. Unique identifier for the assignment filter object
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterId(value *string)() {
    m.assignmentFilterId = value
}
// SetAssignmentFilterLastModifiedDateTime sets the assignmentFilterLastModifiedDateTime property value. The time the assignment filter was last modified.
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.assignmentFilterLastModifiedDateTime = value
}
// SetAssignmentFilterPlatform sets the assignmentFilterPlatform property value. Supported platform types.
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterPlatform(value *DevicePlatformType)() {
    m.assignmentFilterPlatform = value
}
// SetAssignmentFilterType sets the assignmentFilterType property value. Represents type of the assignment filter.
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterType(value *DeviceAndAppManagementAssignmentFilterType)() {
    m.assignmentFilterType = value
}
// SetAssignmentFilterTypeAndEvaluationResults sets the assignmentFilterTypeAndEvaluationResults property value. A collection of filter types and their corresponding evaluation results.
func (m *AssignmentFilterEvaluationSummary) SetAssignmentFilterTypeAndEvaluationResults(value []AssignmentFilterTypeAndEvaluationResultable)() {
    m.assignmentFilterTypeAndEvaluationResults = value
}
// SetEvaluationDateTime sets the evaluationDateTime property value. The time assignment filter was evaluated.
func (m *AssignmentFilterEvaluationSummary) SetEvaluationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.evaluationDateTime = value
}
// SetEvaluationResult sets the evaluationResult property value. Supported evaluation results for filter.
func (m *AssignmentFilterEvaluationSummary) SetEvaluationResult(value *AssignmentFilterEvaluationResult)() {
    m.evaluationResult = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AssignmentFilterEvaluationSummary) SetOdataType(value *string)() {
    m.odataType = value
}
