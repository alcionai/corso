package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// WorkloadActionDeploymentStatus 
type WorkloadActionDeploymentStatus struct {
    // The unique identifier for the workload action. Required. Read-only.
    actionId *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The identifier of any policy that was created by applying the workload action. Optional. Read-only.
    deployedPolicyId *string
    // The detailed information for exceptions that occur when deploying the workload action. Optional. Required.
    error ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.GenericErrorable
    // The excludeGroups property
    excludeGroups []string
    // The includeAllUsers property
    includeAllUsers *bool
    // The includeGroups property
    includeGroups []string
    // The date and time the workload action was last deployed. Optional.
    lastDeploymentDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // The status property
    status *WorkloadActionStatus
}
// NewWorkloadActionDeploymentStatus instantiates a new workloadActionDeploymentStatus and sets the default values.
func NewWorkloadActionDeploymentStatus()(*WorkloadActionDeploymentStatus) {
    m := &WorkloadActionDeploymentStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWorkloadActionDeploymentStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkloadActionDeploymentStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkloadActionDeploymentStatus(), nil
}
// GetActionId gets the actionId property value. The unique identifier for the workload action. Required. Read-only.
func (m *WorkloadActionDeploymentStatus) GetActionId()(*string) {
    return m.actionId
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkloadActionDeploymentStatus) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeployedPolicyId gets the deployedPolicyId property value. The identifier of any policy that was created by applying the workload action. Optional. Read-only.
func (m *WorkloadActionDeploymentStatus) GetDeployedPolicyId()(*string) {
    return m.deployedPolicyId
}
// GetError gets the error property value. The detailed information for exceptions that occur when deploying the workload action. Optional. Required.
func (m *WorkloadActionDeploymentStatus) GetError()(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.GenericErrorable) {
    return m.error
}
// GetExcludeGroups gets the excludeGroups property value. The excludeGroups property
func (m *WorkloadActionDeploymentStatus) GetExcludeGroups()([]string) {
    return m.excludeGroups
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkloadActionDeploymentStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionId(val)
        }
        return nil
    }
    res["deployedPolicyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeployedPolicyId(val)
        }
        return nil
    }
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateGenericErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.GenericErrorable))
        }
        return nil
    }
    res["excludeGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExcludeGroups(res)
        }
        return nil
    }
    res["includeAllUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncludeAllUsers(val)
        }
        return nil
    }
    res["includeGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIncludeGroups(res)
        }
        return nil
    }
    res["lastDeploymentDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastDeploymentDateTime(val)
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWorkloadActionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*WorkloadActionStatus))
        }
        return nil
    }
    return res
}
// GetIncludeAllUsers gets the includeAllUsers property value. The includeAllUsers property
func (m *WorkloadActionDeploymentStatus) GetIncludeAllUsers()(*bool) {
    return m.includeAllUsers
}
// GetIncludeGroups gets the includeGroups property value. The includeGroups property
func (m *WorkloadActionDeploymentStatus) GetIncludeGroups()([]string) {
    return m.includeGroups
}
// GetLastDeploymentDateTime gets the lastDeploymentDateTime property value. The date and time the workload action was last deployed. Optional.
func (m *WorkloadActionDeploymentStatus) GetLastDeploymentDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastDeploymentDateTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WorkloadActionDeploymentStatus) GetOdataType()(*string) {
    return m.odataType
}
// GetStatus gets the status property value. The status property
func (m *WorkloadActionDeploymentStatus) GetStatus()(*WorkloadActionStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *WorkloadActionDeploymentStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("actionId", m.GetActionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deployedPolicyId", m.GetDeployedPolicyId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    if m.GetExcludeGroups() != nil {
        err := writer.WriteCollectionOfStringValues("excludeGroups", m.GetExcludeGroups())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("includeAllUsers", m.GetIncludeAllUsers())
        if err != nil {
            return err
        }
    }
    if m.GetIncludeGroups() != nil {
        err := writer.WriteCollectionOfStringValues("includeGroups", m.GetIncludeGroups())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastDeploymentDateTime", m.GetLastDeploymentDateTime())
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
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
// SetActionId sets the actionId property value. The unique identifier for the workload action. Required. Read-only.
func (m *WorkloadActionDeploymentStatus) SetActionId(value *string)() {
    m.actionId = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkloadActionDeploymentStatus) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeployedPolicyId sets the deployedPolicyId property value. The identifier of any policy that was created by applying the workload action. Optional. Read-only.
func (m *WorkloadActionDeploymentStatus) SetDeployedPolicyId(value *string)() {
    m.deployedPolicyId = value
}
// SetError sets the error property value. The detailed information for exceptions that occur when deploying the workload action. Optional. Required.
func (m *WorkloadActionDeploymentStatus) SetError(value ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.GenericErrorable)() {
    m.error = value
}
// SetExcludeGroups sets the excludeGroups property value. The excludeGroups property
func (m *WorkloadActionDeploymentStatus) SetExcludeGroups(value []string)() {
    m.excludeGroups = value
}
// SetIncludeAllUsers sets the includeAllUsers property value. The includeAllUsers property
func (m *WorkloadActionDeploymentStatus) SetIncludeAllUsers(value *bool)() {
    m.includeAllUsers = value
}
// SetIncludeGroups sets the includeGroups property value. The includeGroups property
func (m *WorkloadActionDeploymentStatus) SetIncludeGroups(value []string)() {
    m.includeGroups = value
}
// SetLastDeploymentDateTime sets the lastDeploymentDateTime property value. The date and time the workload action was last deployed. Optional.
func (m *WorkloadActionDeploymentStatus) SetLastDeploymentDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastDeploymentDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WorkloadActionDeploymentStatus) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStatus sets the status property value. The status property
func (m *WorkloadActionDeploymentStatus) SetStatus(value *WorkloadActionStatus)() {
    m.status = value
}
