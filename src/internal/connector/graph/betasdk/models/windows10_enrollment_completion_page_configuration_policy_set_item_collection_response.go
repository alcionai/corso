package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10EnrollmentCompletionPageConfigurationPolicySetItemCollectionResponse 
type Windows10EnrollmentCompletionPageConfigurationPolicySetItemCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []Windows10EnrollmentCompletionPageConfigurationPolicySetItemable
}
// NewWindows10EnrollmentCompletionPageConfigurationPolicySetItemCollectionResponse instantiates a new Windows10EnrollmentCompletionPageConfigurationPolicySetItemCollectionResponse and sets the default values.
func NewWindows10EnrollmentCompletionPageConfigurationPolicySetItemCollectionResponse()(*Windows10EnrollmentCompletionPageConfigurationPolicySetItemCollectionResponse) {
    m := &Windows10EnrollmentCompletionPageConfigurationPolicySetItemCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateWindows10EnrollmentCompletionPageConfigurationPolicySetItemCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10EnrollmentCompletionPageConfigurationPolicySetItemCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10EnrollmentCompletionPageConfigurationPolicySetItemCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10EnrollmentCompletionPageConfigurationPolicySetItemCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindows10EnrollmentCompletionPageConfigurationPolicySetItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Windows10EnrollmentCompletionPageConfigurationPolicySetItemable, len(val))
            for i, v := range val {
                res[i] = v.(Windows10EnrollmentCompletionPageConfigurationPolicySetItemable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *Windows10EnrollmentCompletionPageConfigurationPolicySetItemCollectionResponse) GetValue()([]Windows10EnrollmentCompletionPageConfigurationPolicySetItemable) {
    return m.value
}
// Serialize serializes information the current object
func (m *Windows10EnrollmentCompletionPageConfigurationPolicySetItemCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *Windows10EnrollmentCompletionPageConfigurationPolicySetItemCollectionResponse) SetValue(value []Windows10EnrollmentCompletionPageConfigurationPolicySetItemable)() {
    m.value = value
}
