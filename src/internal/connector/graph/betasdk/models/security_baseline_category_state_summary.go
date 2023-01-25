package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityBaselineCategoryStateSummary 
type SecurityBaselineCategoryStateSummary struct {
    SecurityBaselineStateSummary
    // The category name
    displayName *string
}
// NewSecurityBaselineCategoryStateSummary instantiates a new SecurityBaselineCategoryStateSummary and sets the default values.
func NewSecurityBaselineCategoryStateSummary()(*SecurityBaselineCategoryStateSummary) {
    m := &SecurityBaselineCategoryStateSummary{
        SecurityBaselineStateSummary: *NewSecurityBaselineStateSummary(),
    }
    odataTypeValue := "#microsoft.graph.securityBaselineCategoryStateSummary";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSecurityBaselineCategoryStateSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityBaselineCategoryStateSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityBaselineCategoryStateSummary(), nil
}
// GetDisplayName gets the displayName property value. The category name
func (m *SecurityBaselineCategoryStateSummary) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecurityBaselineCategoryStateSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SecurityBaselineStateSummary.GetFieldDeserializers()
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
    return res
}
// Serialize serializes information the current object
func (m *SecurityBaselineCategoryStateSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SecurityBaselineStateSummary.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The category name
func (m *SecurityBaselineCategoryStateSummary) SetDisplayName(value *string)() {
    m.displayName = value
}
