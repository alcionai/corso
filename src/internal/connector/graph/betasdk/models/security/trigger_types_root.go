package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// TriggerTypesRoot 
type TriggerTypesRoot struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // The retentionEventTypes property
    retentionEventTypes []RetentionEventTypeable
}
// NewTriggerTypesRoot instantiates a new triggerTypesRoot and sets the default values.
func NewTriggerTypesRoot()(*TriggerTypesRoot) {
    m := &TriggerTypesRoot{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateTriggerTypesRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTriggerTypesRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTriggerTypesRoot(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TriggerTypesRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["retentionEventTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRetentionEventTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RetentionEventTypeable, len(val))
            for i, v := range val {
                res[i] = v.(RetentionEventTypeable)
            }
            m.SetRetentionEventTypes(res)
        }
        return nil
    }
    return res
}
// GetRetentionEventTypes gets the retentionEventTypes property value. The retentionEventTypes property
func (m *TriggerTypesRoot) GetRetentionEventTypes()([]RetentionEventTypeable) {
    return m.retentionEventTypes
}
// Serialize serializes information the current object
func (m *TriggerTypesRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetRetentionEventTypes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRetentionEventTypes()))
        for i, v := range m.GetRetentionEventTypes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("retentionEventTypes", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRetentionEventTypes sets the retentionEventTypes property value. The retentionEventTypes property
func (m *TriggerTypesRoot) SetRetentionEventTypes(value []RetentionEventTypeable)() {
    m.retentionEventTypes = value
}
