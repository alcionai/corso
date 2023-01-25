package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// Catalog 
type Catalog struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // Lists the content that you can approve for deployment. Read-only.
    entries []CatalogEntryable
}
// NewCatalog instantiates a new catalog and sets the default values.
func NewCatalog()(*Catalog) {
    m := &Catalog{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateCatalogFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCatalogFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCatalog(), nil
}
// GetEntries gets the entries property value. Lists the content that you can approve for deployment. Read-only.
func (m *Catalog) GetEntries()([]CatalogEntryable) {
    return m.entries
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Catalog) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["entries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCatalogEntryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CatalogEntryable, len(val))
            for i, v := range val {
                res[i] = v.(CatalogEntryable)
            }
            m.SetEntries(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *Catalog) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetEntries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEntries()))
        for i, v := range m.GetEntries() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("entries", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEntries sets the entries property value. Lists the content that you can approve for deployment. Read-only.
func (m *Catalog) SetEntries(value []CatalogEntryable)() {
    m.entries = value
}
