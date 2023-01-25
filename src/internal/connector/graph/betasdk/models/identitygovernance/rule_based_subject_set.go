package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// RuleBasedSubjectSet 
type RuleBasedSubjectSet struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.SubjectSet
    // The rule for the subject set. Lifecycle Workflows supports a rich set of user properties for configuring the rules using $filter query expressions. For more information, see supported user and query parameters.
    rule *string
}
// NewRuleBasedSubjectSet instantiates a new RuleBasedSubjectSet and sets the default values.
func NewRuleBasedSubjectSet()(*RuleBasedSubjectSet) {
    m := &RuleBasedSubjectSet{
        SubjectSet: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewSubjectSet(),
    }
    odataTypeValue := "#microsoft.graph.identityGovernance.ruleBasedSubjectSet";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateRuleBasedSubjectSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRuleBasedSubjectSetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRuleBasedSubjectSet(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RuleBasedSubjectSet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SubjectSet.GetFieldDeserializers()
    res["rule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRule(val)
        }
        return nil
    }
    return res
}
// GetRule gets the rule property value. The rule for the subject set. Lifecycle Workflows supports a rich set of user properties for configuring the rules using $filter query expressions. For more information, see supported user and query parameters.
func (m *RuleBasedSubjectSet) GetRule()(*string) {
    return m.rule
}
// Serialize serializes information the current object
func (m *RuleBasedSubjectSet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SubjectSet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("rule", m.GetRule())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRule sets the rule property value. The rule for the subject set. Lifecycle Workflows supports a rich set of user properties for configuring the rules using $filter query expressions. For more information, see supported user and query parameters.
func (m *RuleBasedSubjectSet) SetRule(value *string)() {
    m.rule = value
}
