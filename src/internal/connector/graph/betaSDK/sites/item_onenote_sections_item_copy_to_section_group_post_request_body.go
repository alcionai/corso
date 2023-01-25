package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemOnenoteSectionsItemCopyToSectionGroupPostRequestBody provides operations to call the copyToSectionGroup method.
type ItemOnenoteSectionsItemCopyToSectionGroupPostRequestBody struct {
    // Stores model information.
    backingStore BackingStore
}
// NewItemOnenoteSectionsItemCopyToSectionGroupPostRequestBody instantiates a new ItemOnenoteSectionsItemCopyToSectionGroupPostRequestBody and sets the default values.
func NewItemOnenoteSectionsItemCopyToSectionGroupPostRequestBody()(*ItemOnenoteSectionsItemCopyToSectionGroupPostRequestBody) {
    m := &ItemOnenoteSectionsItemCopyToSectionGroupPostRequestBody{
    }
    m._backingStore = BackingStoreFactorySingleton.Instance.CreateBackingStore();
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemOnenoteSectionsItemCopyToSectionGroupPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemOnenoteSectionsItemCopyToSectionGroupPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOnenoteSectionsItemCopyToSectionGroupPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemOnenoteSectionsItemCopyToSectionGroupPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    map[string]interface{} value = m._backingStore.Get("additionalData")
    if value == nil {
        value = make(map[string]interface{});
        m.SetAdditionalData(value);
    }
    return value;
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ItemOnenoteSectionsItemCopyToSectionGroupPostRequestBody) GetBackingStore()(BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemOnenoteSectionsItemCopyToSectionGroupPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["renameAs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRenameAs(val)
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
func (m *ItemOnenoteSectionsItemCopyToSectionGroupPostRequestBody) GetGroupId()(*string) {
    return m.GetBackingStore().Get("groupId");
}
// GetId gets the id property value. The id property
func (m *ItemOnenoteSectionsItemCopyToSectionGroupPostRequestBody) GetId()(*string) {
    return m.GetBackingStore().Get("id");
}
// GetRenameAs gets the renameAs property value. The renameAs property
func (m *ItemOnenoteSectionsItemCopyToSectionGroupPostRequestBody) GetRenameAs()(*string) {
    return m.GetBackingStore().Get("renameAs");
}
// GetSiteCollectionId gets the siteCollectionId property value. The siteCollectionId property
func (m *ItemOnenoteSectionsItemCopyToSectionGroupPostRequestBody) GetSiteCollectionId()(*string) {
    return m.GetBackingStore().Get("siteCollectionId");
}
// GetSiteId gets the siteId property value. The siteId property
func (m *ItemOnenoteSectionsItemCopyToSectionGroupPostRequestBody) GetSiteId()(*string) {
    return m.GetBackingStore().Get("siteId");
}
// Serialize serializes information the current object
func (m *ItemOnenoteSectionsItemCopyToSectionGroupPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("renameAs", m.GetRenameAs())
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
func (m *ItemOnenoteSectionsItemCopyToSectionGroupPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.GetBackingStore().Set("additionalData", value)
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ItemOnenoteSectionsItemCopyToSectionGroupPostRequestBody) SetBackingStore(value BackingStore)() {
    m.GetBackingStore().Set("backingStore", value)
}
// SetGroupId sets the groupId property value. The groupId property
func (m *ItemOnenoteSectionsItemCopyToSectionGroupPostRequestBody) SetGroupId(value *string)() {
    m.GetBackingStore().Set("groupId", value)
}
// SetId sets the id property value. The id property
func (m *ItemOnenoteSectionsItemCopyToSectionGroupPostRequestBody) SetId(value *string)() {
    m.GetBackingStore().Set("id", value)
}
// SetRenameAs sets the renameAs property value. The renameAs property
func (m *ItemOnenoteSectionsItemCopyToSectionGroupPostRequestBody) SetRenameAs(value *string)() {
    m.GetBackingStore().Set("renameAs", value)
}
// SetSiteCollectionId sets the siteCollectionId property value. The siteCollectionId property
func (m *ItemOnenoteSectionsItemCopyToSectionGroupPostRequestBody) SetSiteCollectionId(value *string)() {
    m.GetBackingStore().Set("siteCollectionId", value)
}
// SetSiteId sets the siteId property value. The siteId property
func (m *ItemOnenoteSectionsItemCopyToSectionGroupPostRequestBody) SetSiteId(value *string)() {
    m.GetBackingStore().Set("siteId", value)
}
