package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// ManagedDeviceComplianceTrendCollectionResponse 
type ManagedDeviceComplianceTrendCollectionResponse struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.BaseCollectionPaginationCountResponse
    // The value property
    value []ManagedDeviceComplianceTrendable
}
// NewManagedDeviceComplianceTrendCollectionResponse instantiates a new ManagedDeviceComplianceTrendCollectionResponse and sets the default values.
func NewManagedDeviceComplianceTrendCollectionResponse()(*ManagedDeviceComplianceTrendCollectionResponse) {
    m := &ManagedDeviceComplianceTrendCollectionResponse{
        BaseCollectionPaginationCountResponse: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateManagedDeviceComplianceTrendCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDeviceComplianceTrendCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDeviceComplianceTrendCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDeviceComplianceTrendCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceComplianceTrendFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceComplianceTrendable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedDeviceComplianceTrendable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ManagedDeviceComplianceTrendCollectionResponse) GetValue()([]ManagedDeviceComplianceTrendable) {
    return m.value
}
// Serialize serializes information the current object
func (m *ManagedDeviceComplianceTrendCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ManagedDeviceComplianceTrendCollectionResponse) SetValue(value []ManagedDeviceComplianceTrendable)() {
    m.value = value
}
