package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10EnrollmentCompletionPageConfigurationCollectionResponse 
type Windows10EnrollmentCompletionPageConfigurationCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []Windows10EnrollmentCompletionPageConfigurationable
}
// NewWindows10EnrollmentCompletionPageConfigurationCollectionResponse instantiates a new Windows10EnrollmentCompletionPageConfigurationCollectionResponse and sets the default values.
func NewWindows10EnrollmentCompletionPageConfigurationCollectionResponse()(*Windows10EnrollmentCompletionPageConfigurationCollectionResponse) {
    m := &Windows10EnrollmentCompletionPageConfigurationCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateWindows10EnrollmentCompletionPageConfigurationCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10EnrollmentCompletionPageConfigurationCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10EnrollmentCompletionPageConfigurationCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10EnrollmentCompletionPageConfigurationCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindows10EnrollmentCompletionPageConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Windows10EnrollmentCompletionPageConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(Windows10EnrollmentCompletionPageConfigurationable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *Windows10EnrollmentCompletionPageConfigurationCollectionResponse) GetValue()([]Windows10EnrollmentCompletionPageConfigurationable) {
    return m.value
}
// Serialize serializes information the current object
func (m *Windows10EnrollmentCompletionPageConfigurationCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *Windows10EnrollmentCompletionPageConfigurationCollectionResponse) SetValue(value []Windows10EnrollmentCompletionPageConfigurationable)() {
    m.value = value
}
