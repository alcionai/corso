package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// ItemSitesItemPermissionsItemRevokeGrantsPostRequestBody provides operations to call the revokeGrants method.
type ItemSitesItemPermissionsItemRevokeGrantsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The grantees property
    grantees []ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.DriveRecipientable
}
// NewItemSitesItemPermissionsItemRevokeGrantsPostRequestBody instantiates a new ItemSitesItemPermissionsItemRevokeGrantsPostRequestBody and sets the default values.
func NewItemSitesItemPermissionsItemRevokeGrantsPostRequestBody()(*ItemSitesItemPermissionsItemRevokeGrantsPostRequestBody) {
    m := &ItemSitesItemPermissionsItemRevokeGrantsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemSitesItemPermissionsItemRevokeGrantsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSitesItemPermissionsItemRevokeGrantsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemPermissionsItemRevokeGrantsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSitesItemPermissionsItemRevokeGrantsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSitesItemPermissionsItemRevokeGrantsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["grantees"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateDriveRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.DriveRecipientable, len(val))
            for i, v := range val {
                res[i] = v.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.DriveRecipientable)
            }
            m.SetGrantees(res)
        }
        return nil
    }
    return res
}
// GetGrantees gets the grantees property value. The grantees property
func (m *ItemSitesItemPermissionsItemRevokeGrantsPostRequestBody) GetGrantees()([]ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.DriveRecipientable) {
    return m.grantees
}
// Serialize serializes information the current object
func (m *ItemSitesItemPermissionsItemRevokeGrantsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetGrantees() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGrantees()))
        for i, v := range m.GetGrantees() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("grantees", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSitesItemPermissionsItemRevokeGrantsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetGrantees sets the grantees property value. The grantees property
func (m *ItemSitesItemPermissionsItemRevokeGrantsPostRequestBody) SetGrantees(value []ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.DriveRecipientable)() {
    m.grantees = value
}
