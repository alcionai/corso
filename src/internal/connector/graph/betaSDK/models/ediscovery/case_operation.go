package ediscovery

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// CaseOperation 
type CaseOperation struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // The type of action the operation represents. Possible values are: addToReviewSet,applyTags,contentExport,convertToPdf,estimateStatistics, purgeData
    action *CaseAction
    // The date and time the operation was completed.
    completedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The user that created the operation.
    createdBy ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable
    // The date and time the operation was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The progress of the operation.
    percentProgress *int32
    // Contains success and failure-specific result information.
    resultInfo ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.ResultInfoable
    // The status of the case operation. Possible values are: notStarted, submissionFailed, running, succeeded, partiallySucceeded, failed.
    status *CaseOperationStatus
}
// NewCaseOperation instantiates a new caseOperation and sets the default values.
func NewCaseOperation()(*CaseOperation) {
    m := &CaseOperation{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateCaseOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCaseOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.ediscovery.addToReviewSetOperation":
                        return NewAddToReviewSetOperation(), nil
                    case "#microsoft.graph.ediscovery.caseExportOperation":
                        return NewCaseExportOperation(), nil
                    case "#microsoft.graph.ediscovery.caseHoldOperation":
                        return NewCaseHoldOperation(), nil
                    case "#microsoft.graph.ediscovery.caseIndexOperation":
                        return NewCaseIndexOperation(), nil
                    case "#microsoft.graph.ediscovery.estimateStatisticsOperation":
                        return NewEstimateStatisticsOperation(), nil
                    case "#microsoft.graph.ediscovery.purgeDataOperation":
                        return NewPurgeDataOperation(), nil
                    case "#microsoft.graph.ediscovery.tagOperation":
                        return NewTagOperation(), nil
                }
            }
        }
    }
    return NewCaseOperation(), nil
}
// GetAction gets the action property value. The type of action the operation represents. Possible values are: addToReviewSet,applyTags,contentExport,convertToPdf,estimateStatistics, purgeData
func (m *CaseOperation) GetAction()(*CaseAction) {
    return m.action
}
// GetCompletedDateTime gets the completedDateTime property value. The date and time the operation was completed.
func (m *CaseOperation) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.completedDateTime
}
// GetCreatedBy gets the createdBy property value. The user that created the operation.
func (m *CaseOperation) GetCreatedBy()(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the operation was created.
func (m *CaseOperation) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CaseOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCaseAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*CaseAction))
        }
        return nil
    }
    res["completedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedDateTime(val)
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["percentProgress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPercentProgress(val)
        }
        return nil
    }
    res["resultInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateResultInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResultInfo(val.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.ResultInfoable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCaseOperationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*CaseOperationStatus))
        }
        return nil
    }
    return res
}
// GetPercentProgress gets the percentProgress property value. The progress of the operation.
func (m *CaseOperation) GetPercentProgress()(*int32) {
    return m.percentProgress
}
// GetResultInfo gets the resultInfo property value. Contains success and failure-specific result information.
func (m *CaseOperation) GetResultInfo()(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.ResultInfoable) {
    return m.resultInfo
}
// GetStatus gets the status property value. The status of the case operation. Possible values are: notStarted, submissionFailed, running, succeeded, partiallySucceeded, failed.
func (m *CaseOperation) GetStatus()(*CaseOperationStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *CaseOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAction() != nil {
        cast := (*m.GetAction()).String()
        err = writer.WriteStringValue("action", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("completedDateTime", m.GetCompletedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("percentProgress", m.GetPercentProgress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resultInfo", m.GetResultInfo())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAction sets the action property value. The type of action the operation represents. Possible values are: addToReviewSet,applyTags,contentExport,convertToPdf,estimateStatistics, purgeData
func (m *CaseOperation) SetAction(value *CaseAction)() {
    m.action = value
}
// SetCompletedDateTime sets the completedDateTime property value. The date and time the operation was completed.
func (m *CaseOperation) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completedDateTime = value
}
// SetCreatedBy sets the createdBy property value. The user that created the operation.
func (m *CaseOperation) SetCreatedBy(value ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.IdentitySetable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the operation was created.
func (m *CaseOperation) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetPercentProgress sets the percentProgress property value. The progress of the operation.
func (m *CaseOperation) SetPercentProgress(value *int32)() {
    m.percentProgress = value
}
// SetResultInfo sets the resultInfo property value. Contains success and failure-specific result information.
func (m *CaseOperation) SetResultInfo(value ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.ResultInfoable)() {
    m.resultInfo = value
}
// SetStatus sets the status property value. The status of the case operation. Possible values are: notStarted, submissionFailed, running, succeeded, partiallySucceeded, failed.
func (m *CaseOperation) SetStatus(value *CaseOperationStatus)() {
    m.status = value
}
