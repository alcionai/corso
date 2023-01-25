package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProgramControlTypeCollectionResponse 
type ProgramControlTypeCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []ProgramControlTypeable
}
// NewProgramControlTypeCollectionResponse instantiates a new ProgramControlTypeCollectionResponse and sets the default values.
func NewProgramControlTypeCollectionResponse()(*ProgramControlTypeCollectionResponse) {
    m := &ProgramControlTypeCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateProgramControlTypeCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProgramControlTypeCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProgramControlTypeCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProgramControlTypeCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProgramControlTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProgramControlTypeable, len(val))
            for i, v := range val {
                res[i] = v.(ProgramControlTypeable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ProgramControlTypeCollectionResponse) GetValue()([]ProgramControlTypeable) {
    return m.value
}
// Serialize serializes information the current object
func (m *ProgramControlTypeCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ProgramControlTypeCollectionResponse) SetValue(value []ProgramControlTypeable)() {
    m.value = value
}
