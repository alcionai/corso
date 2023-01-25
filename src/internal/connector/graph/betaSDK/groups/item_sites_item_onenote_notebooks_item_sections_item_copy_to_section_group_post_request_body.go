package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSitesItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody provides operations to call the copyToSectionGroup method.
type ItemSitesItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody struct {
    // Stores model information.
    backingStore BackingStore
}
// NewItemSitesItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody instantiates a new ItemSitesItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody and sets the default values.
func NewItemSitesItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody()(*ItemSitesItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) {
    m := &ItemSitesItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody{
    }
    m._backingStore = BackingStoreFactorySingleton.Instance.CreateBackingStore();
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemSitesItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSitesItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    map[string]interface{} value = m._backingStore.Get("additionalData")
    if value == nil {
        value = make(map[string]interface{});
        m.SetAdditionalData(value);
    }
    return value;
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) GetBackingStore()(BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) GetGroupId()(*string) {
    return m.GetBackingStore().Get("groupId");
}
// GetId gets the id property value. The id property
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) GetId()(*string) {
    return m.GetBackingStore().Get("id");
}
// GetRenameAs gets the renameAs property value. The renameAs property
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) GetRenameAs()(*string) {
    return m.GetBackingStore().Get("renameAs");
}
// GetSiteCollectionId gets the siteCollectionId property value. The siteCollectionId property
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) GetSiteCollectionId()(*string) {
    return m.GetBackingStore().Get("siteCollectionId");
}
// GetSiteId gets the siteId property value. The siteId property
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) GetSiteId()(*string) {
    return m.GetBackingStore().Get("siteId");
}
// Serialize serializes information the current object
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.GetBackingStore().Set("additionalData", value)
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) SetBackingStore(value BackingStore)() {
    m.GetBackingStore().Set("backingStore", value)
}
// SetGroupId sets the groupId property value. The groupId property
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) SetGroupId(value *string)() {
    m.GetBackingStore().Set("groupId", value)
}
// SetId sets the id property value. The id property
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) SetId(value *string)() {
    m.GetBackingStore().Set("id", value)
}
// SetRenameAs sets the renameAs property value. The renameAs property
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) SetRenameAs(value *string)() {
    m.GetBackingStore().Set("renameAs", value)
}
// SetSiteCollectionId sets the siteCollectionId property value. The siteCollectionId property
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) SetSiteCollectionId(value *string)() {
    m.GetBackingStore().Set("siteCollectionId", value)
}
// SetSiteId sets the siteId property value. The siteId property
func (m *ItemSitesItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) SetSiteId(value *string)() {
    m.GetBackingStore().Set("siteId", value)
}
