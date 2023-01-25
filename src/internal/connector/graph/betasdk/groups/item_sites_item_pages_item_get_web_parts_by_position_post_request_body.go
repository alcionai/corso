package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSitesItemPagesItemGetWebPartsByPositionPostRequestBody provides operations to call the getWebPartsByPosition method.
type ItemSitesItemPagesItemGetWebPartsByPositionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The columnId property
    columnId *float64
    // The horizontalSectionId property
    horizontalSectionId *float64
    // The isInVerticalSection property
    isInVerticalSection *bool
    // The webPartIndex property
    webPartIndex *float64
}
// NewItemSitesItemPagesItemGetWebPartsByPositionPostRequestBody instantiates a new ItemSitesItemPagesItemGetWebPartsByPositionPostRequestBody and sets the default values.
func NewItemSitesItemPagesItemGetWebPartsByPositionPostRequestBody()(*ItemSitesItemPagesItemGetWebPartsByPositionPostRequestBody) {
    m := &ItemSitesItemPagesItemGetWebPartsByPositionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemSitesItemPagesItemGetWebPartsByPositionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSitesItemPagesItemGetWebPartsByPositionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemPagesItemGetWebPartsByPositionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSitesItemPagesItemGetWebPartsByPositionPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetColumnId gets the columnId property value. The columnId property
func (m *ItemSitesItemPagesItemGetWebPartsByPositionPostRequestBody) GetColumnId()(*float64) {
    return m.columnId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSitesItemPagesItemGetWebPartsByPositionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ItemSitesItemPagesItemGetWebPartsByPositionPostRequestBody) GetHorizontalSectionId()(*float64) {
    return m.horizontalSectionId
}
// GetIsInVerticalSection gets the isInVerticalSection property value. The isInVerticalSection property
func (m *ItemSitesItemPagesItemGetWebPartsByPositionPostRequestBody) GetIsInVerticalSection()(*bool) {
    return m.isInVerticalSection
}
// GetWebPartIndex gets the webPartIndex property value. The webPartIndex property
func (m *ItemSitesItemPagesItemGetWebPartsByPositionPostRequestBody) GetWebPartIndex()(*float64) {
    return m.webPartIndex
}
// Serialize serializes information the current object
func (m *ItemSitesItemPagesItemGetWebPartsByPositionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSitesItemPagesItemGetWebPartsByPositionPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetColumnId sets the columnId property value. The columnId property
func (m *ItemSitesItemPagesItemGetWebPartsByPositionPostRequestBody) SetColumnId(value *float64)() {
    m.columnId = value
}
// SetHorizontalSectionId sets the horizontalSectionId property value. The horizontalSectionId property
func (m *ItemSitesItemPagesItemGetWebPartsByPositionPostRequestBody) SetHorizontalSectionId(value *float64)() {
    m.horizontalSectionId = value
}
// SetIsInVerticalSection sets the isInVerticalSection property value. The isInVerticalSection property
func (m *ItemSitesItemPagesItemGetWebPartsByPositionPostRequestBody) SetIsInVerticalSection(value *bool)() {
    m.isInVerticalSection = value
}
// SetWebPartIndex sets the webPartIndex property value. The webPartIndex property
func (m *ItemSitesItemPagesItemGetWebPartsByPositionPostRequestBody) SetWebPartIndex(value *float64)() {
    m.webPartIndex = value
}
