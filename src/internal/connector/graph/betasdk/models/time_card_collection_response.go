package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimeCardCollectionResponse 
type TimeCardCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []TimeCardable
}
// NewTimeCardCollectionResponse instantiates a new TimeCardCollectionResponse and sets the default values.
func NewTimeCardCollectionResponse()(*TimeCardCollectionResponse) {
    m := &TimeCardCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateTimeCardCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTimeCardCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTimeCardCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeCardCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTimeCardFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TimeCardable, len(val))
            for i, v := range val {
                res[i] = v.(TimeCardable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *TimeCardCollectionResponse) GetValue()([]TimeCardable) {
    return m.value
}
// Serialize serializes information the current object
func (m *TimeCardCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *TimeCardCollectionResponse) SetValue(value []TimeCardable)() {
    m.value = value
}
