package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GovernanceRoleAssignmentRequestStatus 
type GovernanceRoleAssignmentRequestStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // The status of the role assignment request. The value can be InProgress or Closed.
    status *string
    // The details of the status of the role assignment request. It represents the evaluation results of different rules.
    statusDetails []KeyValueable
    // The sub status of the role assignment request. The values can be Accepted, PendingEvaluation, Granted, Denied, PendingProvisioning, Provisioned, PendingRevocation, Revoked, Canceled, Failed, PendingApprovalProvisioning, PendingApproval, FailedAsResourceIsLocked, PendingAdminDecision, AdminApproved, AdminDenied, TimedOut, and ProvisioningStarted.
    subStatus *string
}
// NewGovernanceRoleAssignmentRequestStatus instantiates a new governanceRoleAssignmentRequestStatus and sets the default values.
func NewGovernanceRoleAssignmentRequestStatus()(*GovernanceRoleAssignmentRequestStatus) {
    m := &GovernanceRoleAssignmentRequestStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGovernanceRoleAssignmentRequestStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGovernanceRoleAssignmentRequestStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGovernanceRoleAssignmentRequestStatus(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GovernanceRoleAssignmentRequestStatus) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GovernanceRoleAssignmentRequestStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["statusDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValueable, len(val))
            for i, v := range val {
                res[i] = v.(KeyValueable)
            }
            m.SetStatusDetails(res)
        }
        return nil
    }
    res["subStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubStatus(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *GovernanceRoleAssignmentRequestStatus) GetOdataType()(*string) {
    return m.odataType
}
// GetStatus gets the status property value. The status of the role assignment request. The value can be InProgress or Closed.
func (m *GovernanceRoleAssignmentRequestStatus) GetStatus()(*string) {
    return m.status
}
// GetStatusDetails gets the statusDetails property value. The details of the status of the role assignment request. It represents the evaluation results of different rules.
func (m *GovernanceRoleAssignmentRequestStatus) GetStatusDetails()([]KeyValueable) {
    return m.statusDetails
}
// GetSubStatus gets the subStatus property value. The sub status of the role assignment request. The values can be Accepted, PendingEvaluation, Granted, Denied, PendingProvisioning, Provisioned, PendingRevocation, Revoked, Canceled, Failed, PendingApprovalProvisioning, PendingApproval, FailedAsResourceIsLocked, PendingAdminDecision, AdminApproved, AdminDenied, TimedOut, and ProvisioningStarted.
func (m *GovernanceRoleAssignmentRequestStatus) GetSubStatus()(*string) {
    return m.subStatus
}
// Serialize serializes information the current object
func (m *GovernanceRoleAssignmentRequestStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    if m.GetStatusDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetStatusDetails()))
        for i, v := range m.GetStatusDetails() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("statusDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subStatus", m.GetSubStatus())
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
func (m *GovernanceRoleAssignmentRequestStatus) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *GovernanceRoleAssignmentRequestStatus) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStatus sets the status property value. The status of the role assignment request. The value can be InProgress or Closed.
func (m *GovernanceRoleAssignmentRequestStatus) SetStatus(value *string)() {
    m.status = value
}
// SetStatusDetails sets the statusDetails property value. The details of the status of the role assignment request. It represents the evaluation results of different rules.
func (m *GovernanceRoleAssignmentRequestStatus) SetStatusDetails(value []KeyValueable)() {
    m.statusDetails = value
}
// SetSubStatus sets the subStatus property value. The sub status of the role assignment request. The values can be Accepted, PendingEvaluation, Granted, Denied, PendingProvisioning, Provisioned, PendingRevocation, Revoked, Canceled, Failed, PendingApprovalProvisioning, PendingApproval, FailedAsResourceIsLocked, PendingAdminDecision, AdminApproved, AdminDenied, TimedOut, and ProvisioningStarted.
func (m *GovernanceRoleAssignmentRequestStatus) SetSubStatus(value *string)() {
    m.subStatus = value
}
