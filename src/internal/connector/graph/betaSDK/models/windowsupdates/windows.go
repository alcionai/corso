package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// Windows 
type Windows struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // Entity that acts as a container for the functionality of the Windows Update for Business deployment service. Read-only.
    updates Updatesable
}
// NewWindows instantiates a new Windows and sets the default values.
func NewWindows()(*Windows) {
    m := &Windows{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateWindowsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["updates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUpdatesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdates(val.(Updatesable))
        }
        return nil
    }
    return res
}
// GetUpdates gets the updates property value. Entity that acts as a container for the functionality of the Windows Update for Business deployment service. Read-only.
func (m *Windows) GetUpdates()(Updatesable) {
    return m.updates
}
// Serialize serializes information the current object
func (m *Windows) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("updates", m.GetUpdates())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUpdates sets the updates property value. Entity that acts as a container for the functionality of the Windows Update for Business deployment service. Read-only.
func (m *Windows) SetUpdates(value Updatesable)() {
    m.updates = value
}
