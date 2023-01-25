package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// Updates 
type Updates struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // Catalog of content that can be approved for deployment by the deployment service. Read-only.
    catalog Catalogable
    // Deployments created using the deployment service. Read-only.
    deployments []Deploymentable
    // Service connections to external resources such as analytics workspaces.
    resourceConnections []ResourceConnectionable
    // Assets registered with the deployment service that can receive updates. Read-only.
    updatableAssets []UpdatableAssetable
}
// NewUpdates instantiates a new updates and sets the default values.
func NewUpdates()(*Updates) {
    m := &Updates{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateUpdatesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdatesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdates(), nil
}
// GetCatalog gets the catalog property value. Catalog of content that can be approved for deployment by the deployment service. Read-only.
func (m *Updates) GetCatalog()(Catalogable) {
    return m.catalog
}
// GetDeployments gets the deployments property value. Deployments created using the deployment service. Read-only.
func (m *Updates) GetDeployments()([]Deploymentable) {
    return m.deployments
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Updates) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["catalog"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCatalogFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalog(val.(Catalogable))
        }
        return nil
    }
    res["deployments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeploymentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Deploymentable, len(val))
            for i, v := range val {
                res[i] = v.(Deploymentable)
            }
            m.SetDeployments(res)
        }
        return nil
    }
    res["resourceConnections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateResourceConnectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ResourceConnectionable, len(val))
            for i, v := range val {
                res[i] = v.(ResourceConnectionable)
            }
            m.SetResourceConnections(res)
        }
        return nil
    }
    res["updatableAssets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUpdatableAssetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UpdatableAssetable, len(val))
            for i, v := range val {
                res[i] = v.(UpdatableAssetable)
            }
            m.SetUpdatableAssets(res)
        }
        return nil
    }
    return res
}
// GetResourceConnections gets the resourceConnections property value. Service connections to external resources such as analytics workspaces.
func (m *Updates) GetResourceConnections()([]ResourceConnectionable) {
    return m.resourceConnections
}
// GetUpdatableAssets gets the updatableAssets property value. Assets registered with the deployment service that can receive updates. Read-only.
func (m *Updates) GetUpdatableAssets()([]UpdatableAssetable) {
    return m.updatableAssets
}
// Serialize serializes information the current object
func (m *Updates) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("catalog", m.GetCatalog())
        if err != nil {
            return err
        }
    }
    if m.GetDeployments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeployments()))
        for i, v := range m.GetDeployments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deployments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetResourceConnections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResourceConnections()))
        for i, v := range m.GetResourceConnections() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("resourceConnections", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUpdatableAssets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUpdatableAssets()))
        for i, v := range m.GetUpdatableAssets() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("updatableAssets", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCatalog sets the catalog property value. Catalog of content that can be approved for deployment by the deployment service. Read-only.
func (m *Updates) SetCatalog(value Catalogable)() {
    m.catalog = value
}
// SetDeployments sets the deployments property value. Deployments created using the deployment service. Read-only.
func (m *Updates) SetDeployments(value []Deploymentable)() {
    m.deployments = value
}
// SetResourceConnections sets the resourceConnections property value. Service connections to external resources such as analytics workspaces.
func (m *Updates) SetResourceConnections(value []ResourceConnectionable)() {
    m.resourceConnections = value
}
// SetUpdatableAssets sets the updatableAssets property value. Assets registered with the deployment service that can receive updates. Read-only.
func (m *Updates) SetUpdatableAssets(value []UpdatableAssetable)() {
    m.updatableAssets = value
}
