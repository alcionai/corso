package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemPagesItemGetWebPartsByPositionPostRequestBody provides operations to call the getWebPartsByPosition method.
type ItemPagesItemGetWebPartsByPositionPostRequestBody struct {
    // Stores model information.
    backingStore BackingStore
}
// NewItemPagesItemGetWebPartsByPositionPostRequestBody instantiates a new ItemPagesItemGetWebPartsByPositionPostRequestBody and sets the default values.
func NewItemPagesItemGetWebPartsByPositionPostRequestBody()(*ItemPagesItemGetWebPartsByPositionPostRequestBody) {
    m := &ItemPagesItemGetWebPartsByPositionPostRequestBody{
    }
    m._backingStore = BackingStoreFactorySingleton.Instance.CreateBackingStore();
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemPagesItemGetWebPartsByPositionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemPagesItemGetWebPartsByPositionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPagesItemGetWebPartsByPositionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemPagesItemGetWebPartsByPositionPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    map[string]interface{} value = m._backingStore.Get("additionalData")
    if value == nil {
        value = make(map[string]interface{});
        m.SetAdditionalData(value);
    }
    return value;
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ItemPagesItemGetWebPartsByPositionPostRequestBody) GetBackingStore()(BackingStore) {
    return m.backingStore
}
// GetColumnId gets the columnId property value. The columnId property
func (m *ItemPagesItemGetWebPartsByPositionPostRequestBody) GetColumnId()(*float64) {
    return m.GetBackingStore().Get("columnId");
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemPagesItemGetWebPartsByPositionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["columnId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColumnId(val)
        }
        return nil
    }
    res["horizontalSectionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHorizontalSectionId(val)
        }
        return nil
    }
    res["isInVerticalSection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsInVerticalSection(val)
        }
        return nil
    }
    res["webPartIndex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebPartIndex(val)
        }
        return nil
    }
    return res
}
// GetHorizontalSectionId gets the horizontalSectionId property value. The horizontalSectionId property
func (m *ItemPagesItemGetWebPartsByPositionPostRequestBody) GetHorizontalSectionId()(*float64) {
    return m.GetBackingStore().Get("horizontalSectionId");
}
// GetIsInVerticalSection gets the isInVerticalSection property value. The isInVerticalSection property
func (m *ItemPagesItemGetWebPartsByPositionPostRequestBody) GetIsInVerticalSection()(*bool) {
    return m.GetBackingStore().Get("isInVerticalSection");
}
// GetWebPartIndex gets the webPartIndex property value. The webPartIndex property
func (m *ItemPagesItemGetWebPartsByPositionPostRequestBody) GetWebPartIndex()(*float64) {
    return m.GetBackingStore().Get("webPartIndex");
}
// Serialize serializes information the current object
func (m *ItemPagesItemGetWebPartsByPositionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("columnId", m.GetColumnId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("horizontalSectionId", m.GetHorizontalSectionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isInVerticalSection", m.GetIsInVerticalSection())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("webPartIndex", m.GetWebPartIndex())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemPagesItemGetWebPartsByPositionPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.GetBackingStore().Set("additionalData", value)
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ItemPagesItemGetWebPartsByPositionPostRequestBody) SetBackingStore(value BackingStore)() {
    m.GetBackingStore().Set("backingStore", value)
}
// SetColumnId sets the columnId property value. The columnId property
func (m *ItemPagesItemGetWebPartsByPositionPostRequestBody) SetColumnId(value *float64)() {
    m.GetBackingStore().Set("columnId", value)
}
// SetHorizontalSectionId sets the horizontalSectionId property value. The horizontalSectionId property
func (m *ItemPagesItemGetWebPartsByPositionPostRequestBody) SetHorizontalSectionId(value *float64)() {
    m.GetBackingStore().Set("horizontalSectionId", value)
}
// SetIsInVerticalSection sets the isInVerticalSection property value. The isInVerticalSection property
func (m *ItemPagesItemGetWebPartsByPositionPostRequestBody) SetIsInVerticalSection(value *bool)() {
    m.GetBackingStore().Set("isInVerticalSection", value)
}
// SetWebPartIndex sets the webPartIndex property value. The webPartIndex property
func (m *ItemPagesItemGetWebPartsByPositionPostRequestBody) SetWebPartIndex(value *float64)() {
    m.GetBackingStore().Set("webPartIndex", value)
}
