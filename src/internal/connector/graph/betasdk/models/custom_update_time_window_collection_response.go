package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomUpdateTimeWindowCollectionResponse 
type CustomUpdateTimeWindowCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []CustomUpdateTimeWindowable
}
// NewCustomUpdateTimeWindowCollectionResponse instantiates a new CustomUpdateTimeWindowCollectionResponse and sets the default values.
func NewCustomUpdateTimeWindowCollectionResponse()(*CustomUpdateTimeWindowCollectionResponse) {
    m := &CustomUpdateTimeWindowCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateCustomUpdateTimeWindowCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomUpdateTimeWindowCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomUpdateTimeWindowCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomUpdateTimeWindowCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomUpdateTimeWindowFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomUpdateTimeWindowable, len(val))
            for i, v := range val {
                res[i] = v.(CustomUpdateTimeWindowable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *CustomUpdateTimeWindowCollectionResponse) GetValue()([]CustomUpdateTimeWindowable) {
    return m.value
}
// Serialize serializes information the current object
func (m *CustomUpdateTimeWindowCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *CustomUpdateTimeWindowCollectionResponse) SetValue(value []CustomUpdateTimeWindowable)() {
    m.value = value
}
