package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApprovalWorkflowProvider 
type ApprovalWorkflowProvider struct {
    Entity
    // The businessFlows property
    businessFlows []BusinessFlowable
    // The businessFlowsWithRequestsAwaitingMyDecision property
    businessFlowsWithRequestsAwaitingMyDecision []BusinessFlowable
    // The displayName property
    displayName *string
    // The policyTemplates property
    policyTemplates []GovernancePolicyTemplateable
}
// NewApprovalWorkflowProvider instantiates a new ApprovalWorkflowProvider and sets the default values.
func NewApprovalWorkflowProvider()(*ApprovalWorkflowProvider) {
    m := &ApprovalWorkflowProvider{
        Entity: *NewEntity(),
    }
    return m
}
// CreateApprovalWorkflowProviderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApprovalWorkflowProviderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApprovalWorkflowProvider(), nil
}
// GetBusinessFlows gets the businessFlows property value. The businessFlows property
func (m *ApprovalWorkflowProvider) GetBusinessFlows()([]BusinessFlowable) {
    return m.businessFlows
}
// GetBusinessFlowsWithRequestsAwaitingMyDecision gets the businessFlowsWithRequestsAwaitingMyDecision property value. The businessFlowsWithRequestsAwaitingMyDecision property
func (m *ApprovalWorkflowProvider) GetBusinessFlowsWithRequestsAwaitingMyDecision()([]BusinessFlowable) {
    return m.businessFlowsWithRequestsAwaitingMyDecision
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *ApprovalWorkflowProvider) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApprovalWorkflowProvider) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["businessFlows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBusinessFlowFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BusinessFlowable, len(val))
            for i, v := range val {
                res[i] = v.(BusinessFlowable)
            }
            m.SetBusinessFlows(res)
        }
        return nil
    }
    res["businessFlowsWithRequestsAwaitingMyDecision"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBusinessFlowFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BusinessFlowable, len(val))
            for i, v := range val {
                res[i] = v.(BusinessFlowable)
            }
            m.SetBusinessFlowsWithRequestsAwaitingMyDecision(res)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["policyTemplates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGovernancePolicyTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GovernancePolicyTemplateable, len(val))
            for i, v := range val {
                res[i] = v.(GovernancePolicyTemplateable)
            }
            m.SetPolicyTemplates(res)
        }
        return nil
    }
    return res
}
// GetPolicyTemplates gets the policyTemplates property value. The policyTemplates property
func (m *ApprovalWorkflowProvider) GetPolicyTemplates()([]GovernancePolicyTemplateable) {
    return m.policyTemplates
}
// Serialize serializes information the current object
func (m *ApprovalWorkflowProvider) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBusinessFlows() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBusinessFlows()))
        for i, v := range m.GetBusinessFlows() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("businessFlows", cast)
        if err != nil {
            return err
        }
    }
    if m.GetBusinessFlowsWithRequestsAwaitingMyDecision() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBusinessFlowsWithRequestsAwaitingMyDecision()))
        for i, v := range m.GetBusinessFlowsWithRequestsAwaitingMyDecision() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("businessFlowsWithRequestsAwaitingMyDecision", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetPolicyTemplates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPolicyTemplates()))
        for i, v := range m.GetPolicyTemplates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("policyTemplates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBusinessFlows sets the businessFlows property value. The businessFlows property
func (m *ApprovalWorkflowProvider) SetBusinessFlows(value []BusinessFlowable)() {
    m.businessFlows = value
}
// SetBusinessFlowsWithRequestsAwaitingMyDecision sets the businessFlowsWithRequestsAwaitingMyDecision property value. The businessFlowsWithRequestsAwaitingMyDecision property
func (m *ApprovalWorkflowProvider) SetBusinessFlowsWithRequestsAwaitingMyDecision(value []BusinessFlowable)() {
    m.businessFlowsWithRequestsAwaitingMyDecision = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *ApprovalWorkflowProvider) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetPolicyTemplates sets the policyTemplates property value. The policyTemplates property
func (m *ApprovalWorkflowProvider) SetPolicyTemplates(value []GovernancePolicyTemplateable)() {
    m.policyTemplates = value
}
