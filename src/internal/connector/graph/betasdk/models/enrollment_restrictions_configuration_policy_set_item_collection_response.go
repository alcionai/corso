package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EnrollmentRestrictionsConfigurationPolicySetItemCollectionResponse 
type EnrollmentRestrictionsConfigurationPolicySetItemCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []EnrollmentRestrictionsConfigurationPolicySetItemable
}
// NewEnrollmentRestrictionsConfigurationPolicySetItemCollectionResponse instantiates a new EnrollmentRestrictionsConfigurationPolicySetItemCollectionResponse and sets the default values.
func NewEnrollmentRestrictionsConfigurationPolicySetItemCollectionResponse()(*EnrollmentRestrictionsConfigurationPolicySetItemCollectionResponse) {
    m := &EnrollmentRestrictionsConfigurationPolicySetItemCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateEnrollmentRestrictionsConfigurationPolicySetItemCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEnrollmentRestrictionsConfigurationPolicySetItemCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnrollmentRestrictionsConfigurationPolicySetItemCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EnrollmentRestrictionsConfigurationPolicySetItemCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEnrollmentRestrictionsConfigurationPolicySetItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EnrollmentRestrictionsConfigurationPolicySetItemable, len(val))
            for i, v := range val {
                res[i] = v.(EnrollmentRestrictionsConfigurationPolicySetItemable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *EnrollmentRestrictionsConfigurationPolicySetItemCollectionResponse) GetValue()([]EnrollmentRestrictionsConfigurationPolicySetItemable) {
    return m.value
}
// Serialize serializes information the current object
func (m *EnrollmentRestrictionsConfigurationPolicySetItemCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *EnrollmentRestrictionsConfigurationPolicySetItemCollectionResponse) SetValue(value []EnrollmentRestrictionsConfigurationPolicySetItemable)() {
    m.value = value
}
