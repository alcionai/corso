package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ClassificationJobResponse 
type ClassificationJobResponse struct {
    JobResponseBase
    // The result property
    result DetectedSensitiveContentWrapperable
}
// NewClassificationJobResponse instantiates a new ClassificationJobResponse and sets the default values.
func NewClassificationJobResponse()(*ClassificationJobResponse) {
    m := &ClassificationJobResponse{
        JobResponseBase: *NewJobResponseBase(),
    }
    return m
}
// CreateClassificationJobResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateClassificationJobResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClassificationJobResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClassificationJobResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.JobResponseBase.GetFieldDeserializers()
    res["result"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDetectedSensitiveContentWrapperFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResult(val.(DetectedSensitiveContentWrapperable))
        }
        return nil
    }
    return res
}
// GetResult gets the result property value. The result property
func (m *ClassificationJobResponse) GetResult()(DetectedSensitiveContentWrapperable) {
    return m.result
}
// Serialize serializes information the current object
func (m *ClassificationJobResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.JobResponseBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("result", m.GetResult())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetResult sets the result property value. The result property
func (m *ClassificationJobResponse) SetResult(value DetectedSensitiveContentWrapperable)() {
    m.result = value
}
