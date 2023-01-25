package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidForWorkCompliancePolicyCollectionResponse 
type AndroidForWorkCompliancePolicyCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []AndroidForWorkCompliancePolicyable
}
// NewAndroidForWorkCompliancePolicyCollectionResponse instantiates a new AndroidForWorkCompliancePolicyCollectionResponse and sets the default values.
func NewAndroidForWorkCompliancePolicyCollectionResponse()(*AndroidForWorkCompliancePolicyCollectionResponse) {
    m := &AndroidForWorkCompliancePolicyCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateAndroidForWorkCompliancePolicyCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidForWorkCompliancePolicyCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidForWorkCompliancePolicyCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidForWorkCompliancePolicyCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidForWorkCompliancePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidForWorkCompliancePolicyable, len(val))
            for i, v := range val {
                res[i] = v.(AndroidForWorkCompliancePolicyable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *AndroidForWorkCompliancePolicyCollectionResponse) GetValue()([]AndroidForWorkCompliancePolicyable) {
    return m.value
}
// Serialize serializes information the current object
func (m *AndroidForWorkCompliancePolicyCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *AndroidForWorkCompliancePolicyCollectionResponse) SetValue(value []AndroidForWorkCompliancePolicyable)() {
    m.value = value
}
