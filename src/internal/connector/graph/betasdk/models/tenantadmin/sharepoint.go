package tenantadmin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// Sharepoint 
type Sharepoint struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // Represents the tenant-level settings for SharePoint and OneDrive.
    settings Settingsable
}
// NewSharepoint instantiates a new Sharepoint and sets the default values.
func NewSharepoint()(*Sharepoint) {
    m := &Sharepoint{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateSharepointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSharepointFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSharepoint(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Sharepoint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(Settingsable))
        }
        return nil
    }
    return res
}
// GetSettings gets the settings property value. Represents the tenant-level settings for SharePoint and OneDrive.
func (m *Sharepoint) GetSettings()(Settingsable) {
    return m.settings
}
// Serialize serializes information the current object
func (m *Sharepoint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSettings sets the settings property value. Represents the tenant-level settings for SharePoint and OneDrive.
func (m *Sharepoint) SetSettings(value Settingsable)() {
    m.settings = value
}
