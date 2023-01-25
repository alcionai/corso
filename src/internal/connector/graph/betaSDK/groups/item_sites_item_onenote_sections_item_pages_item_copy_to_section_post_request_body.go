package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSitesItemOnenoteSectionsItemPagesItemCopyToSectionPostRequestBody provides operations to call the copyToSection method.
type ItemSitesItemOnenoteSectionsItemPagesItemCopyToSectionPostRequestBody struct {
    // Stores model information.
    backingStore BackingStore
}
// NewItemSitesItemOnenoteSectionsItemPagesItemCopyToSectionPostRequestBody instantiates a new ItemSitesItemOnenoteSectionsItemPagesItemCopyToSectionPostRequestBody and sets the default values.
func NewItemSitesItemOnenoteSectionsItemPagesItemCopyToSectionPostRequestBody()(*ItemSitesItemOnenoteSectionsItemPagesItemCopyToSectionPostRequestBody) {
    m := &ItemSitesItemOnenoteSectionsItemPagesItemCopyToSectionPostRequestBody{
    }
    m._backingStore = BackingStoreFactorySingleton.Instance.CreateBackingStore();
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemSitesItemOnenoteSectionsItemPagesItemCopyToSectionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSitesItemOnenoteSectionsItemPagesItemCopyToSectionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemOnenoteSectionsItemPagesItemCopyToSectionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSitesItemOnenoteSectionsItemPagesItemCopyToSectionPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    map[string]interface{} value = m._backingStore.Get("additionalData")
    if value == nil {
        value = make(map[string]interface{});
        m.SetAdditionalData(value);
    }
    return value;
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ItemSitesItemOnenoteSectionsItemPagesItemCopyToSectionPostRequestBody) GetBackingStore()(BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSitesItemOnenoteSectionsItemPagesItemCopyToSectionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["groupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupId(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["siteCollectionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteCollectionId(val)
        }
        return nil
    }
    res["siteId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteId(val)
        }
        return nil
    }
    return res
}
// GetGroupId gets the groupId property value. The groupId property
func (m *ItemSitesItemOnenoteSectionsItemPagesItemCopyToSectionPostRequestBody) GetGroupId()(*string) {
    return m.GetBackingStore().Get("groupId");
}
// GetId gets the id property value. The id property
func (m *ItemSitesItemOnenoteSectionsItemPagesItemCopyToSectionPostRequestBody) GetId()(*string) {
    return m.GetBackingStore().Get("id");
}
// GetSiteCollectionId gets the siteCollectionId property value. The siteCollectionId property
func (m *ItemSitesItemOnenoteSectionsItemPagesItemCopyToSectionPostRequestBody) GetSiteCollectionId()(*string) {
    return m.GetBackingStore().Get("siteCollectionId");
}
// GetSiteId gets the siteId property value. The siteId property
func (m *ItemSitesItemOnenoteSectionsItemPagesItemCopyToSectionPostRequestBody) GetSiteId()(*string) {
    return m.GetBackingStore().Get("siteId");
}
// Serialize serializes information the current object
func (m *ItemSitesItemOnenoteSectionsItemPagesItemCopyToSectionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("groupId", m.GetGroupId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("siteCollectionId", m.GetSiteCollectionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("siteId", m.GetSiteId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSitesItemOnenoteSectionsItemPagesItemCopyToSectionPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.GetBackingStore().Set("additionalData", value)
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ItemSitesItemOnenoteSectionsItemPagesItemCopyToSectionPostRequestBody) SetBackingStore(value BackingStore)() {
    m.GetBackingStore().Set("backingStore", value)
}
// SetGroupId sets the groupId property value. The groupId property
func (m *ItemSitesItemOnenoteSectionsItemPagesItemCopyToSectionPostRequestBody) SetGroupId(value *string)() {
    m.GetBackingStore().Set("groupId", value)
}
// SetId sets the id property value. The id property
func (m *ItemSitesItemOnenoteSectionsItemPagesItemCopyToSectionPostRequestBody) SetId(value *string)() {
    m.GetBackingStore().Set("id", value)
}
// SetSiteCollectionId sets the siteCollectionId property value. The siteCollectionId property
func (m *ItemSitesItemOnenoteSectionsItemPagesItemCopyToSectionPostRequestBody) SetSiteCollectionId(value *string)() {
    m.GetBackingStore().Set("siteCollectionId", value)
}
// SetSiteId sets the siteId property value. The siteId property
func (m *ItemSitesItemOnenoteSectionsItemPagesItemCopyToSectionPostRequestBody) SetSiteId(value *string)() {
    m.GetBackingStore().Set("siteId", value)
}
