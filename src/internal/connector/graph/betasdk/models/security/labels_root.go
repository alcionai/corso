package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// LabelsRoot 
type LabelsRoot struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // The retentionLabels property
    retentionLabels []RetentionLabelable
}
// NewLabelsRoot instantiates a new labelsRoot and sets the default values.
func NewLabelsRoot()(*LabelsRoot) {
    m := &LabelsRoot{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateLabelsRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLabelsRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLabelsRoot(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LabelsRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["retentionLabels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRetentionLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RetentionLabelable, len(val))
            for i, v := range val {
                res[i] = v.(RetentionLabelable)
            }
            m.SetRetentionLabels(res)
        }
        return nil
    }
    return res
}
// GetRetentionLabels gets the retentionLabels property value. The retentionLabels property
func (m *LabelsRoot) GetRetentionLabels()([]RetentionLabelable) {
    return m.retentionLabels
}
// Serialize serializes information the current object
func (m *LabelsRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetRetentionLabels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRetentionLabels()))
        for i, v := range m.GetRetentionLabels() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("retentionLabels", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRetentionLabels sets the retentionLabels property value. The retentionLabels property
func (m *LabelsRoot) SetRetentionLabels(value []RetentionLabelable)() {
    m.retentionLabels = value
}
