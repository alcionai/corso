package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InformationProtectionLabelCollectionResponse provides operations to manage the labels property of the microsoft.graph.informationProtectionPolicy entity.
type InformationProtectionLabelCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []InformationProtectionLabelable
}
// NewInformationProtectionLabelCollectionResponse instantiates a new InformationProtectionLabelCollectionResponse and sets the default values.
func NewInformationProtectionLabelCollectionResponse()(*InformationProtectionLabelCollectionResponse) {
    m := &InformationProtectionLabelCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateInformationProtectionLabelCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInformationProtectionLabelCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInformationProtectionLabelCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InformationProtectionLabelCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateInformationProtectionLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]InformationProtectionLabelable, len(val))
            for i, v := range val {
                res[i] = v.(InformationProtectionLabelable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *InformationProtectionLabelCollectionResponse) GetValue()([]InformationProtectionLabelable) {
    return m.value
}
// Serialize serializes information the current object
func (m *InformationProtectionLabelCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *InformationProtectionLabelCollectionResponse) SetValue(value []InformationProtectionLabelable)() {
    m.value = value
}
