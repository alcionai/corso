package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// UpdatableAsset provides operations to call the add method.
type UpdatableAsset struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
}
// NewUpdatableAsset instantiates a new updatableAsset and sets the default values.
func NewUpdatableAsset()(*UpdatableAsset) {
    m := &UpdatableAsset{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateUpdatableAssetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdatableAssetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.windowsUpdates.azureADDevice":
                        return NewAzureADDevice(), nil
                    case "#microsoft.graph.windowsUpdates.updatableAssetGroup":
                        return NewUpdatableAssetGroup(), nil
                }
            }
        }
    }
    return NewUpdatableAsset(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdatableAsset) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *UpdatableAsset) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
