package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// ExternalActivityResult 
type ExternalActivityResult struct {
    ExternalActivity
    // Error information explaining failure to process external activity.
    error ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.PublicErrorable
}
// NewExternalActivityResult instantiates a new ExternalActivityResult and sets the default values.
func NewExternalActivityResult()(*ExternalActivityResult) {
    m := &ExternalActivityResult{
        ExternalActivity: *NewExternalActivity(),
    }
    return m
}
// CreateExternalActivityResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExternalActivityResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalActivityResult(), nil
}
// GetError gets the error property value. Error information explaining failure to process external activity.
func (m *ExternalActivityResult) GetError()(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.PublicErrorable) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExternalActivityResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ExternalActivity.GetFieldDeserializers()
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreatePublicErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.PublicErrorable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ExternalActivityResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ExternalActivity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetError sets the error property value. Error information explaining failure to process external activity.
func (m *ExternalActivityResult) SetError(value ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.PublicErrorable)() {
    m.error = value
}
