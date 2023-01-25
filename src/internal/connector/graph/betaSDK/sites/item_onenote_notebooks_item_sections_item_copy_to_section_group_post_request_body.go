package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody provides operations to call the copyToSectionGroup method.
type ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody struct {
    // Stores model information.
    backingStore BackingStore
}
// NewItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody instantiates a new ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody and sets the default values.
func NewItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody()(*ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) {
    m := &ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody{
    }
    m._backingStore = BackingStoreFactorySingleton.Instance.CreateBackingStore();
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    map[string]interface{} value = m._backingStore.Get("additionalData")
    if value == nil {
        value = make(map[string]interface{});
        m.SetAdditionalData(value);
    }
    return value;
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) GetBackingStore()(BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) GetGroupId()(*string) {
    return m.GetBackingStore().Get("groupId");
}
// GetId gets the id property value. The id property
func (m *ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) GetId()(*string) {
    return m.GetBackingStore().Get("id");
}
// GetRenameAs gets the renameAs property value. The renameAs property
func (m *ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) GetRenameAs()(*string) {
    return m.GetBackingStore().Get("renameAs");
}
// GetSiteCollectionId gets the siteCollectionId property value. The siteCollectionId property
func (m *ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) GetSiteCollectionId()(*string) {
    return m.GetBackingStore().Get("siteCollectionId");
}
// GetSiteId gets the siteId property value. The siteId property
func (m *ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) GetSiteId()(*string) {
    return m.GetBackingStore().Get("siteId");
}
// Serialize serializes information the current object
func (m *ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.GetBackingStore().Set("additionalData", value)
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) SetBackingStore(value BackingStore)() {
    m.GetBackingStore().Set("backingStore", value)
}
// SetGroupId sets the groupId property value. The groupId property
func (m *ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) SetGroupId(value *string)() {
    m.GetBackingStore().Set("groupId", value)
}
// SetId sets the id property value. The id property
func (m *ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) SetId(value *string)() {
    m.GetBackingStore().Set("id", value)
}
// SetRenameAs sets the renameAs property value. The renameAs property
func (m *ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) SetRenameAs(value *string)() {
    m.GetBackingStore().Set("renameAs", value)
}
// SetSiteCollectionId sets the siteCollectionId property value. The siteCollectionId property
func (m *ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) SetSiteCollectionId(value *string)() {
    m.GetBackingStore().Set("siteCollectionId", value)
}
// SetSiteId sets the siteId property value. The siteId property
func (m *ItemOnenoteNotebooksItemSectionsItemCopyToSectionGroupPostRequestBody) SetSiteId(value *string)() {
    m.GetBackingStore().Set("siteId", value)
}
