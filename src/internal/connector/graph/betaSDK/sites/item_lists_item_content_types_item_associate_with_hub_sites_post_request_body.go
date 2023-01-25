package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody provides operations to call the associateWithHubSites method.
type ItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody struct {
    // Stores model information.
    backingStore BackingStore
}
// NewItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody instantiates a new ItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody and sets the default values.
func NewItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody()(*ItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody) {
    m := &ItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody{
    }
    m._backingStore = BackingStoreFactorySingleton.Instance.CreateBackingStore();
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    map[string]interface{} value = m._backingStore.Get("additionalData")
    if value == nil {
        value = make(map[string]interface{});
        m.SetAdditionalData(value);
    }
    return value;
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody) GetBackingStore()(BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["hubSiteUrls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetHubSiteUrls(res)
        }
        return nil
    }
    res["propagateToExistingLists"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPropagateToExistingLists(val)
        }
        return nil
    }
    return res
}
// GetHubSiteUrls gets the hubSiteUrls property value. The hubSiteUrls property
func (m *ItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody) GetHubSiteUrls()([]string) {
    return m.GetBackingStore().Get("hubSiteUrls");
}
// GetPropagateToExistingLists gets the propagateToExistingLists property value. The propagateToExistingLists property
func (m *ItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody) GetPropagateToExistingLists()(*bool) {
    return m.GetBackingStore().Get("propagateToExistingLists");
}
// Serialize serializes information the current object
func (m *ItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetHubSiteUrls() != nil {
        err := writer.WriteCollectionOfStringValues("hubSiteUrls", m.GetHubSiteUrls())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("propagateToExistingLists", m.GetPropagateToExistingLists())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.GetBackingStore().Set("additionalData", value)
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody) SetBackingStore(value BackingStore)() {
    m.GetBackingStore().Set("backingStore", value)
}
// SetHubSiteUrls sets the hubSiteUrls property value. The hubSiteUrls property
func (m *ItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody) SetHubSiteUrls(value []string)() {
    m.GetBackingStore().Set("hubSiteUrls", value)
}
// SetPropagateToExistingLists sets the propagateToExistingLists property value. The propagateToExistingLists property
func (m *ItemListsItemContentTypesItemAssociateWithHubSitesPostRequestBody) SetPropagateToExistingLists(value *bool)() {
    m.GetBackingStore().Set("propagateToExistingLists", value)
}
