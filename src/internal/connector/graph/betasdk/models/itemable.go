package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Itemable 
type Itemable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBaseUnitOfMeasureId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetBlocked()(*bool)
    GetDisplayName()(*string)
    GetGtin()(*string)
    GetInventory()(*float64)
    GetItemCategory()(ItemCategoryable)
    GetItemCategoryCode()(*string)
    GetItemCategoryId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNumber()(*string)
    GetPicture()([]Pictureable)
    GetPriceIncludesTax()(*bool)
    GetTaxGroupCode()(*string)
    GetTaxGroupId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetType()(*string)
    GetUnitCost()(*float64)
    GetUnitPrice()(*float64)
    SetBaseUnitOfMeasureId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetBlocked(value *bool)()
    SetDisplayName(value *string)()
    SetGtin(value *string)()
    SetInventory(value *float64)()
    SetItemCategory(value ItemCategoryable)()
    SetItemCategoryCode(value *string)()
    SetItemCategoryId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNumber(value *string)()
    SetPicture(value []Pictureable)()
    SetPriceIncludesTax(value *bool)()
    SetTaxGroupCode(value *string)()
    SetTaxGroupId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetType(value *string)()
    SetUnitCost(value *float64)()
    SetUnitPrice(value *float64)()
}
