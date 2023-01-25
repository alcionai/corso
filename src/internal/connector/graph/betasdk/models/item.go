package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Item provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type Item struct {
    Entity
    // The baseUnitOfMeasureId property
    baseUnitOfMeasureId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The blocked property
    blocked *bool
    // The displayName property
    displayName *string
    // The gtin property
    gtin *string
    // The inventory property
    inventory *float64
    // The itemCategory property
    itemCategory ItemCategoryable
    // The itemCategoryCode property
    itemCategoryCode *string
    // The itemCategoryId property
    itemCategoryId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number property
    number *string
    // The picture property
    picture []Pictureable
    // The priceIncludesTax property
    priceIncludesTax *bool
    // The taxGroupCode property
    taxGroupCode *string
    // The taxGroupId property
    taxGroupId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The type property
    type_escaped *string
    // The unitCost property
    unitCost *float64
    // The unitPrice property
    unitPrice *float64
}
// NewItem instantiates a new item and sets the default values.
func NewItem()(*Item) {
    m := &Item{
        Entity: *NewEntity(),
    }
    return m
}
// CreateItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItem(), nil
}
// GetBaseUnitOfMeasureId gets the baseUnitOfMeasureId property value. The baseUnitOfMeasureId property
func (m *Item) GetBaseUnitOfMeasureId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.baseUnitOfMeasureId
}
// GetBlocked gets the blocked property value. The blocked property
func (m *Item) GetBlocked()(*bool) {
    return m.blocked
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *Item) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Item) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["baseUnitOfMeasureId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBaseUnitOfMeasureId(val)
        }
        return nil
    }
    res["blocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlocked(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["gtin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGtin(val)
        }
        return nil
    }
    res["inventory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInventory(val)
        }
        return nil
    }
    res["itemCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemCategory(val.(ItemCategoryable))
        }
        return nil
    }
    res["itemCategoryCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemCategoryCode(val)
        }
        return nil
    }
    res["itemCategoryId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemCategoryId(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumber(val)
        }
        return nil
    }
    res["picture"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePictureFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Pictureable, len(val))
            for i, v := range val {
                res[i] = v.(Pictureable)
            }
            m.SetPicture(res)
        }
        return nil
    }
    res["priceIncludesTax"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriceIncludesTax(val)
        }
        return nil
    }
    res["taxGroupCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaxGroupCode(val)
        }
        return nil
    }
    res["taxGroupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaxGroupId(val)
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
    res["unitCost"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnitCost(val)
        }
        return nil
    }
    res["unitPrice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnitPrice(val)
        }
        return nil
    }
    return res
}
// GetGtin gets the gtin property value. The gtin property
func (m *Item) GetGtin()(*string) {
    return m.gtin
}
// GetInventory gets the inventory property value. The inventory property
func (m *Item) GetInventory()(*float64) {
    return m.inventory
}
// GetItemCategory gets the itemCategory property value. The itemCategory property
func (m *Item) GetItemCategory()(ItemCategoryable) {
    return m.itemCategory
}
// GetItemCategoryCode gets the itemCategoryCode property value. The itemCategoryCode property
func (m *Item) GetItemCategoryCode()(*string) {
    return m.itemCategoryCode
}
// GetItemCategoryId gets the itemCategoryId property value. The itemCategoryId property
func (m *Item) GetItemCategoryId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.itemCategoryId
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Item) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetNumber gets the number property value. The number property
func (m *Item) GetNumber()(*string) {
    return m.number
}
// GetPicture gets the picture property value. The picture property
func (m *Item) GetPicture()([]Pictureable) {
    return m.picture
}
// GetPriceIncludesTax gets the priceIncludesTax property value. The priceIncludesTax property
func (m *Item) GetPriceIncludesTax()(*bool) {
    return m.priceIncludesTax
}
// GetTaxGroupCode gets the taxGroupCode property value. The taxGroupCode property
func (m *Item) GetTaxGroupCode()(*string) {
    return m.taxGroupCode
}
// GetTaxGroupId gets the taxGroupId property value. The taxGroupId property
func (m *Item) GetTaxGroupId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.taxGroupId
}
// GetType gets the type property value. The type property
func (m *Item) GetType()(*string) {
    return m.type_escaped
}
// GetUnitCost gets the unitCost property value. The unitCost property
func (m *Item) GetUnitCost()(*float64) {
    return m.unitCost
}
// GetUnitPrice gets the unitPrice property value. The unitPrice property
func (m *Item) GetUnitPrice()(*float64) {
    return m.unitPrice
}
// Serialize serializes information the current object
func (m *Item) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteUUIDValue("baseUnitOfMeasureId", m.GetBaseUnitOfMeasureId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("blocked", m.GetBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("gtin", m.GetGtin())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("inventory", m.GetInventory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("itemCategory", m.GetItemCategory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("itemCategoryCode", m.GetItemCategoryCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("itemCategoryId", m.GetItemCategoryId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("number", m.GetNumber())
        if err != nil {
            return err
        }
    }
    if m.GetPicture() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPicture()))
        for i, v := range m.GetPicture() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("picture", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("priceIncludesTax", m.GetPriceIncludesTax())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("taxGroupCode", m.GetTaxGroupCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("taxGroupId", m.GetTaxGroupId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("unitCost", m.GetUnitCost())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("unitPrice", m.GetUnitPrice())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBaseUnitOfMeasureId sets the baseUnitOfMeasureId property value. The baseUnitOfMeasureId property
func (m *Item) SetBaseUnitOfMeasureId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.baseUnitOfMeasureId = value
}
// SetBlocked sets the blocked property value. The blocked property
func (m *Item) SetBlocked(value *bool)() {
    m.blocked = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *Item) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGtin sets the gtin property value. The gtin property
func (m *Item) SetGtin(value *string)() {
    m.gtin = value
}
// SetInventory sets the inventory property value. The inventory property
func (m *Item) SetInventory(value *float64)() {
    m.inventory = value
}
// SetItemCategory sets the itemCategory property value. The itemCategory property
func (m *Item) SetItemCategory(value ItemCategoryable)() {
    m.itemCategory = value
}
// SetItemCategoryCode sets the itemCategoryCode property value. The itemCategoryCode property
func (m *Item) SetItemCategoryCode(value *string)() {
    m.itemCategoryCode = value
}
// SetItemCategoryId sets the itemCategoryId property value. The itemCategoryId property
func (m *Item) SetItemCategoryId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.itemCategoryId = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Item) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetNumber sets the number property value. The number property
func (m *Item) SetNumber(value *string)() {
    m.number = value
}
// SetPicture sets the picture property value. The picture property
func (m *Item) SetPicture(value []Pictureable)() {
    m.picture = value
}
// SetPriceIncludesTax sets the priceIncludesTax property value. The priceIncludesTax property
func (m *Item) SetPriceIncludesTax(value *bool)() {
    m.priceIncludesTax = value
}
// SetTaxGroupCode sets the taxGroupCode property value. The taxGroupCode property
func (m *Item) SetTaxGroupCode(value *string)() {
    m.taxGroupCode = value
}
// SetTaxGroupId sets the taxGroupId property value. The taxGroupId property
func (m *Item) SetTaxGroupId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.taxGroupId = value
}
// SetType sets the type property value. The type property
func (m *Item) SetType(value *string)() {
    m.type_escaped = value
}
// SetUnitCost sets the unitCost property value. The unitCost property
func (m *Item) SetUnitCost(value *float64)() {
    m.unitCost = value
}
// SetUnitPrice sets the unitPrice property value. The unitPrice property
func (m *Item) SetUnitPrice(value *float64)() {
    m.unitPrice = value
}
