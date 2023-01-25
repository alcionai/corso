package sites

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4 "betasdk/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemListsItemItemsItemCreateLinkPostRequestBody provides operations to call the createLink method.
type ItemListsItemItemsItemCreateLinkPostRequestBody struct {
    // Stores model information.
    backingStore BackingStore
    // The type property
    Type_escaped *string
}
// NewItemListsItemItemsItemCreateLinkPostRequestBody instantiates a new ItemListsItemItemsItemCreateLinkPostRequestBody and sets the default values.
func NewItemListsItemItemsItemCreateLinkPostRequestBody()(*ItemListsItemItemsItemCreateLinkPostRequestBody) {
    m := &ItemListsItemItemsItemCreateLinkPostRequestBody{
    }
    m._backingStore = BackingStoreFactorySingleton.Instance.CreateBackingStore();
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemListsItemItemsItemCreateLinkPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemListsItemItemsItemCreateLinkPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemListsItemItemsItemCreateLinkPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemListsItemItemsItemCreateLinkPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    map[string]interface{} value = m._backingStore.Get("additionalData")
    if value == nil {
        value = make(map[string]interface{});
        m.SetAdditionalData(value);
    }
    return value;
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ItemListsItemItemsItemCreateLinkPostRequestBody) GetBackingStore()(BackingStore) {
    return m.backingStore
}
// GetExpirationDateTime gets the expirationDateTime property value. The expirationDateTime property
func (m *ItemListsItemItemsItemCreateLinkPostRequestBody) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.GetBackingStore().Get("expirationDateTime");
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemListsItemItemsItemCreateLinkPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    res["recipients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateDriveRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DriveRecipient, len(val))
            for i, v := range val {
                res[i] = *(v.(*ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DriveRecipient))
            }
            m.SetRecipients(res)
        }
        return nil
    }
    res["retainInheritedPermissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRetainInheritedPermissions(val)
        }
        return nil
    }
    res["scope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetPassword gets the password property value. The password property
func (m *ItemListsItemItemsItemCreateLinkPostRequestBody) GetPassword()(*string) {
    return m.GetBackingStore().Get("password");
}
// GetRecipients gets the recipients property value. The recipients property
func (m *ItemListsItemItemsItemCreateLinkPostRequestBody) GetRecipients()([]ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DriveRecipientable) {
    return m.GetBackingStore().Get("recipients");
}
// GetRetainInheritedPermissions gets the retainInheritedPermissions property value. The retainInheritedPermissions property
func (m *ItemListsItemItemsItemCreateLinkPostRequestBody) GetRetainInheritedPermissions()(*bool) {
    return m.GetBackingStore().Get("retainInheritedPermissions");
}
// GetScope gets the scope property value. The scope property
func (m *ItemListsItemItemsItemCreateLinkPostRequestBody) GetScope()(*string) {
    return m.GetBackingStore().Get("scope");
}
// GetType gets the type property value. The type property
func (m *ItemListsItemItemsItemCreateLinkPostRequestBody) GetType()(*string) {
    return m.GetBackingStore().Get("type_escaped");
}
// Serialize serializes information the current object
func (m *ItemListsItemItemsItemCreateLinkPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("password", m.GetPassword())
        if err != nil {
            return err
        }
    }
    if m.GetRecipients() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRecipients()))
        for i, v := range m.GetRecipients() {
            temp := v
            cast[i] = i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("recipients", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("retainInheritedPermissions", m.GetRetainInheritedPermissions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("scope", m.GetScope())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemListsItemItemsItemCreateLinkPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.GetBackingStore().Set("additionalData", value)
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ItemListsItemItemsItemCreateLinkPostRequestBody) SetBackingStore(value BackingStore)() {
    m.GetBackingStore().Set("backingStore", value)
}
// SetExpirationDateTime sets the expirationDateTime property value. The expirationDateTime property
func (m *ItemListsItemItemsItemCreateLinkPostRequestBody) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.GetBackingStore().Set("expirationDateTime", value)
}
// SetPassword sets the password property value. The password property
func (m *ItemListsItemItemsItemCreateLinkPostRequestBody) SetPassword(value *string)() {
    m.GetBackingStore().Set("password", value)
}
// SetRecipients sets the recipients property value. The recipients property
func (m *ItemListsItemItemsItemCreateLinkPostRequestBody) SetRecipients(value []ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DriveRecipientable)() {
    m.GetBackingStore().Set("recipients", value)
}
// SetRetainInheritedPermissions sets the retainInheritedPermissions property value. The retainInheritedPermissions property
func (m *ItemListsItemItemsItemCreateLinkPostRequestBody) SetRetainInheritedPermissions(value *bool)() {
    m.GetBackingStore().Set("retainInheritedPermissions", value)
}
// SetScope sets the scope property value. The scope property
func (m *ItemListsItemItemsItemCreateLinkPostRequestBody) SetScope(value *string)() {
    m.GetBackingStore().Set("scope", value)
}
// SetType sets the type property value. The type property
func (m *ItemListsItemItemsItemCreateLinkPostRequestBody) SetType(value *string)() {
    m.GetBackingStore().Set("type_escaped", value)
}
