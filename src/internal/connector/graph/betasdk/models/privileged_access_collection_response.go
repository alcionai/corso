package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedAccessCollectionResponse 
type PrivilegedAccessCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []PrivilegedAccessable
}
// NewPrivilegedAccessCollectionResponse instantiates a new PrivilegedAccessCollectionResponse and sets the default values.
func NewPrivilegedAccessCollectionResponse()(*PrivilegedAccessCollectionResponse) {
    m := &PrivilegedAccessCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreatePrivilegedAccessCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegedAccessCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedAccessCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedAccessCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrivilegedAccessFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrivilegedAccessable, len(val))
            for i, v := range val {
                res[i] = v.(PrivilegedAccessable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *PrivilegedAccessCollectionResponse) GetValue()([]PrivilegedAccessable) {
    return m.value
}
// Serialize serializes information the current object
func (m *PrivilegedAccessCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *PrivilegedAccessCollectionResponse) SetValue(value []PrivilegedAccessable)() {
    m.value = value
}
