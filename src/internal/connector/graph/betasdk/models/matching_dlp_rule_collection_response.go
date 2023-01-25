package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MatchingDlpRuleCollectionResponse 
type MatchingDlpRuleCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []MatchingDlpRuleable
}
// NewMatchingDlpRuleCollectionResponse instantiates a new MatchingDlpRuleCollectionResponse and sets the default values.
func NewMatchingDlpRuleCollectionResponse()(*MatchingDlpRuleCollectionResponse) {
    m := &MatchingDlpRuleCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateMatchingDlpRuleCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMatchingDlpRuleCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMatchingDlpRuleCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MatchingDlpRuleCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMatchingDlpRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MatchingDlpRuleable, len(val))
            for i, v := range val {
                res[i] = v.(MatchingDlpRuleable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *MatchingDlpRuleCollectionResponse) GetValue()([]MatchingDlpRuleable) {
    return m.value
}
// Serialize serializes information the current object
func (m *MatchingDlpRuleCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *MatchingDlpRuleCollectionResponse) SetValue(value []MatchingDlpRuleable)() {
    m.value = value
}
