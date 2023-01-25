package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10XCustomSubjectAlternativeNameCollectionResponse 
type Windows10XCustomSubjectAlternativeNameCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []Windows10XCustomSubjectAlternativeNameable
}
// NewWindows10XCustomSubjectAlternativeNameCollectionResponse instantiates a new Windows10XCustomSubjectAlternativeNameCollectionResponse and sets the default values.
func NewWindows10XCustomSubjectAlternativeNameCollectionResponse()(*Windows10XCustomSubjectAlternativeNameCollectionResponse) {
    m := &Windows10XCustomSubjectAlternativeNameCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateWindows10XCustomSubjectAlternativeNameCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10XCustomSubjectAlternativeNameCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10XCustomSubjectAlternativeNameCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10XCustomSubjectAlternativeNameCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindows10XCustomSubjectAlternativeNameFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Windows10XCustomSubjectAlternativeNameable, len(val))
            for i, v := range val {
                res[i] = v.(Windows10XCustomSubjectAlternativeNameable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *Windows10XCustomSubjectAlternativeNameCollectionResponse) GetValue()([]Windows10XCustomSubjectAlternativeNameable) {
    return m.value
}
// Serialize serializes information the current object
func (m *Windows10XCustomSubjectAlternativeNameCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *Windows10XCustomSubjectAlternativeNameCollectionResponse) SetValue(value []Windows10XCustomSubjectAlternativeNameable)() {
    m.value = value
}
