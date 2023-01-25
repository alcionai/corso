package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewSet 
type AccessReviewSet struct {
    Entity
    // Represents an Azure AD access review decision on an instance of a review.
    decisions []AccessReviewInstanceDecisionItemable
    // Represents the template and scheduling for an access review.
    definitions []AccessReviewScheduleDefinitionable
    // Represents a collection of access review history data and the scopes used to collect that data.
    historyDefinitions []AccessReviewHistoryDefinitionable
    // Resource that enables administrators to manage directory-level access review policies in their tenant.
    policy AccessReviewPolicyable
}
// NewAccessReviewSet instantiates a new AccessReviewSet and sets the default values.
func NewAccessReviewSet()(*AccessReviewSet) {
    m := &AccessReviewSet{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessReviewSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewSetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewSet(), nil
}
// GetDecisions gets the decisions property value. Represents an Azure AD access review decision on an instance of a review.
func (m *AccessReviewSet) GetDecisions()([]AccessReviewInstanceDecisionItemable) {
    return m.decisions
}
// GetDefinitions gets the definitions property value. Represents the template and scheduling for an access review.
func (m *AccessReviewSet) GetDefinitions()([]AccessReviewScheduleDefinitionable) {
    return m.definitions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewSet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["decisions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewInstanceDecisionItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewInstanceDecisionItemable, len(val))
            for i, v := range val {
                res[i] = v.(AccessReviewInstanceDecisionItemable)
            }
            m.SetDecisions(res)
        }
        return nil
    }
    res["definitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewScheduleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewScheduleDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(AccessReviewScheduleDefinitionable)
            }
            m.SetDefinitions(res)
        }
        return nil
    }
    res["historyDefinitions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewHistoryDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewHistoryDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(AccessReviewHistoryDefinitionable)
            }
            m.SetHistoryDefinitions(res)
        }
        return nil
    }
    res["policy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessReviewPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicy(val.(AccessReviewPolicyable))
        }
        return nil
    }
    return res
}
// GetHistoryDefinitions gets the historyDefinitions property value. Represents a collection of access review history data and the scopes used to collect that data.
func (m *AccessReviewSet) GetHistoryDefinitions()([]AccessReviewHistoryDefinitionable) {
    return m.historyDefinitions
}
// GetPolicy gets the policy property value. Resource that enables administrators to manage directory-level access review policies in their tenant.
func (m *AccessReviewSet) GetPolicy()(AccessReviewPolicyable) {
    return m.policy
}
// Serialize serializes information the current object
func (m *AccessReviewSet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDecisions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDecisions()))
        for i, v := range m.GetDecisions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("decisions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDefinitions()))
        for i, v := range m.GetDefinitions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("definitions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHistoryDefinitions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHistoryDefinitions()))
        for i, v := range m.GetHistoryDefinitions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("historyDefinitions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("policy", m.GetPolicy())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDecisions sets the decisions property value. Represents an Azure AD access review decision on an instance of a review.
func (m *AccessReviewSet) SetDecisions(value []AccessReviewInstanceDecisionItemable)() {
    m.decisions = value
}
// SetDefinitions sets the definitions property value. Represents the template and scheduling for an access review.
func (m *AccessReviewSet) SetDefinitions(value []AccessReviewScheduleDefinitionable)() {
    m.definitions = value
}
// SetHistoryDefinitions sets the historyDefinitions property value. Represents a collection of access review history data and the scopes used to collect that data.
func (m *AccessReviewSet) SetHistoryDefinitions(value []AccessReviewHistoryDefinitionable)() {
    m.historyDefinitions = value
}
// SetPolicy sets the policy property value. Resource that enables administrators to manage directory-level access review policies in their tenant.
func (m *AccessReviewSet) SetPolicy(value AccessReviewPolicyable)() {
    m.policy = value
}
