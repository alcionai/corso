package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ClassificationError 
type ClassificationError struct {
    ClassifcationErrorBase
    // The details property
    details []ClassifcationErrorBaseable
}
// NewClassificationError instantiates a new ClassificationError and sets the default values.
func NewClassificationError()(*ClassificationError) {
    m := &ClassificationError{
        ClassifcationErrorBase: *NewClassifcationErrorBase(),
    }
    return m
}
// CreateClassificationErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateClassificationErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClassificationError(), nil
}
// GetDetails gets the details property value. The details property
func (m *ClassificationError) GetDetails()([]ClassifcationErrorBaseable) {
    return m.details
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClassificationError) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ClassifcationErrorBase.GetFieldDeserializers()
    res["details"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateClassifcationErrorBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ClassifcationErrorBaseable, len(val))
            for i, v := range val {
                res[i] = v.(ClassifcationErrorBaseable)
            }
            m.SetDetails(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ClassificationError) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ClassifcationErrorBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDetails()))
        for i, v := range m.GetDetails() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("details", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDetails sets the details property value. The details property
func (m *ClassificationError) SetDetails(value []ClassifcationErrorBaseable)() {
    m.details = value
}
