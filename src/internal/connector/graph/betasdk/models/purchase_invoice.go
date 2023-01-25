package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PurchaseInvoice provides operations to call the add method.
type PurchaseInvoice struct {
    Entity
    // The buyFromAddress property
    buyFromAddress PostalAddressTypeable
    // The currency property
    currency Currencyable
    // The currencyCode property
    currencyCode *string
    // The currencyId property
    currencyId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The discountAmount property
    discountAmount *float64
    // The discountAppliedBeforeTax property
    discountAppliedBeforeTax *bool
    // The dueDate property
    dueDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The invoiceDate property
    invoiceDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number property
    number *string
    // The payToAddress property
    payToAddress PostalAddressTypeable
    // The payToContact property
    payToContact *string
    // The payToName property
    payToName *string
    // The payToVendorId property
    payToVendorId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The payToVendorNumber property
    payToVendorNumber *string
    // The pricesIncludeTax property
    pricesIncludeTax *bool
    // The purchaseInvoiceLines property
    purchaseInvoiceLines []PurchaseInvoiceLineable
    // The shipToAddress property
    shipToAddress PostalAddressTypeable
    // The shipToContact property
    shipToContact *string
    // The shipToName property
    shipToName *string
    // The status property
    status *string
    // The totalAmountExcludingTax property
    totalAmountExcludingTax *float64
    // The totalAmountIncludingTax property
    totalAmountIncludingTax *float64
    // The totalTaxAmount property
    totalTaxAmount *float64
    // The vendor property
    vendor_escaped Vendor_escapedable
    // The vendorId property
    vendorId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The vendorInvoiceNumber property
    vendorInvoiceNumber *string
    // The vendorName property
    vendorName *string
    // The vendorNumber property
    vendorNumber *string
}
// NewPurchaseInvoice instantiates a new purchaseInvoice and sets the default values.
func NewPurchaseInvoice()(*PurchaseInvoice) {
    m := &PurchaseInvoice{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePurchaseInvoiceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePurchaseInvoiceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPurchaseInvoice(), nil
}
// GetBuyFromAddress gets the buyFromAddress property value. The buyFromAddress property
func (m *PurchaseInvoice) GetBuyFromAddress()(PostalAddressTypeable) {
    return m.buyFromAddress
}
// GetCurrency gets the currency property value. The currency property
func (m *PurchaseInvoice) GetCurrency()(Currencyable) {
    return m.currency
}
// GetCurrencyCode gets the currencyCode property value. The currencyCode property
func (m *PurchaseInvoice) GetCurrencyCode()(*string) {
    return m.currencyCode
}
// GetCurrencyId gets the currencyId property value. The currencyId property
func (m *PurchaseInvoice) GetCurrencyId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.currencyId
}
// GetDiscountAmount gets the discountAmount property value. The discountAmount property
func (m *PurchaseInvoice) GetDiscountAmount()(*float64) {
    return m.discountAmount
}
// GetDiscountAppliedBeforeTax gets the discountAppliedBeforeTax property value. The discountAppliedBeforeTax property
func (m *PurchaseInvoice) GetDiscountAppliedBeforeTax()(*bool) {
    return m.discountAppliedBeforeTax
}
// GetDueDate gets the dueDate property value. The dueDate property
func (m *PurchaseInvoice) GetDueDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.dueDate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PurchaseInvoice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["buyFromAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePostalAddressTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBuyFromAddress(val.(PostalAddressTypeable))
        }
        return nil
    }
    res["currency"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCurrencyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrency(val.(Currencyable))
        }
        return nil
    }
    res["currencyCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrencyCode(val)
        }
        return nil
    }
    res["currencyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrencyId(val)
        }
        return nil
    }
    res["discountAmount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscountAmount(val)
        }
        return nil
    }
    res["discountAppliedBeforeTax"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscountAppliedBeforeTax(val)
        }
        return nil
    }
    res["dueDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueDate(val)
        }
        return nil
    }
    res["invoiceDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvoiceDate(val)
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
    res["payToAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePostalAddressTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayToAddress(val.(PostalAddressTypeable))
        }
        return nil
    }
    res["payToContact"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayToContact(val)
        }
        return nil
    }
    res["payToName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayToName(val)
        }
        return nil
    }
    res["payToVendorId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayToVendorId(val)
        }
        return nil
    }
    res["payToVendorNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayToVendorNumber(val)
        }
        return nil
    }
    res["pricesIncludeTax"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPricesIncludeTax(val)
        }
        return nil
    }
    res["purchaseInvoiceLines"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePurchaseInvoiceLineFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PurchaseInvoiceLineable, len(val))
            for i, v := range val {
                res[i] = v.(PurchaseInvoiceLineable)
            }
            m.SetPurchaseInvoiceLines(res)
        }
        return nil
    }
    res["shipToAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePostalAddressTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShipToAddress(val.(PostalAddressTypeable))
        }
        return nil
    }
    res["shipToContact"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShipToContact(val)
        }
        return nil
    }
    res["shipToName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShipToName(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["totalAmountExcludingTax"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalAmountExcludingTax(val)
        }
        return nil
    }
    res["totalAmountIncludingTax"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalAmountIncludingTax(val)
        }
        return nil
    }
    res["totalTaxAmount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalTaxAmount(val)
        }
        return nil
    }
    res["vendor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVendor_escapedFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendor(val.(Vendor_escapedable))
        }
        return nil
    }
    res["vendorId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorId(val)
        }
        return nil
    }
    res["vendorInvoiceNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorInvoiceNumber(val)
        }
        return nil
    }
    res["vendorName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorName(val)
        }
        return nil
    }
    res["vendorNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorNumber(val)
        }
        return nil
    }
    return res
}
// GetInvoiceDate gets the invoiceDate property value. The invoiceDate property
func (m *PurchaseInvoice) GetInvoiceDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.invoiceDate
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *PurchaseInvoice) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetNumber gets the number property value. The number property
func (m *PurchaseInvoice) GetNumber()(*string) {
    return m.number
}
// GetPayToAddress gets the payToAddress property value. The payToAddress property
func (m *PurchaseInvoice) GetPayToAddress()(PostalAddressTypeable) {
    return m.payToAddress
}
// GetPayToContact gets the payToContact property value. The payToContact property
func (m *PurchaseInvoice) GetPayToContact()(*string) {
    return m.payToContact
}
// GetPayToName gets the payToName property value. The payToName property
func (m *PurchaseInvoice) GetPayToName()(*string) {
    return m.payToName
}
// GetPayToVendorId gets the payToVendorId property value. The payToVendorId property
func (m *PurchaseInvoice) GetPayToVendorId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.payToVendorId
}
// GetPayToVendorNumber gets the payToVendorNumber property value. The payToVendorNumber property
func (m *PurchaseInvoice) GetPayToVendorNumber()(*string) {
    return m.payToVendorNumber
}
// GetPricesIncludeTax gets the pricesIncludeTax property value. The pricesIncludeTax property
func (m *PurchaseInvoice) GetPricesIncludeTax()(*bool) {
    return m.pricesIncludeTax
}
// GetPurchaseInvoiceLines gets the purchaseInvoiceLines property value. The purchaseInvoiceLines property
func (m *PurchaseInvoice) GetPurchaseInvoiceLines()([]PurchaseInvoiceLineable) {
    return m.purchaseInvoiceLines
}
// GetShipToAddress gets the shipToAddress property value. The shipToAddress property
func (m *PurchaseInvoice) GetShipToAddress()(PostalAddressTypeable) {
    return m.shipToAddress
}
// GetShipToContact gets the shipToContact property value. The shipToContact property
func (m *PurchaseInvoice) GetShipToContact()(*string) {
    return m.shipToContact
}
// GetShipToName gets the shipToName property value. The shipToName property
func (m *PurchaseInvoice) GetShipToName()(*string) {
    return m.shipToName
}
// GetStatus gets the status property value. The status property
func (m *PurchaseInvoice) GetStatus()(*string) {
    return m.status
}
// GetTotalAmountExcludingTax gets the totalAmountExcludingTax property value. The totalAmountExcludingTax property
func (m *PurchaseInvoice) GetTotalAmountExcludingTax()(*float64) {
    return m.totalAmountExcludingTax
}
// GetTotalAmountIncludingTax gets the totalAmountIncludingTax property value. The totalAmountIncludingTax property
func (m *PurchaseInvoice) GetTotalAmountIncludingTax()(*float64) {
    return m.totalAmountIncludingTax
}
// GetTotalTaxAmount gets the totalTaxAmount property value. The totalTaxAmount property
func (m *PurchaseInvoice) GetTotalTaxAmount()(*float64) {
    return m.totalTaxAmount
}
// GetVendor gets the vendor property value. The vendor property
func (m *PurchaseInvoice) GetVendor()(Vendor_escapedable) {
    return m.vendor_escaped
}
// GetVendorId gets the vendorId property value. The vendorId property
func (m *PurchaseInvoice) GetVendorId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.vendorId
}
// GetVendorInvoiceNumber gets the vendorInvoiceNumber property value. The vendorInvoiceNumber property
func (m *PurchaseInvoice) GetVendorInvoiceNumber()(*string) {
    return m.vendorInvoiceNumber
}
// GetVendorName gets the vendorName property value. The vendorName property
func (m *PurchaseInvoice) GetVendorName()(*string) {
    return m.vendorName
}
// GetVendorNumber gets the vendorNumber property value. The vendorNumber property
func (m *PurchaseInvoice) GetVendorNumber()(*string) {
    return m.vendorNumber
}
// Serialize serializes information the current object
func (m *PurchaseInvoice) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("buyFromAddress", m.GetBuyFromAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("currency", m.GetCurrency())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("currencyCode", m.GetCurrencyCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("currencyId", m.GetCurrencyId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("discountAmount", m.GetDiscountAmount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("discountAppliedBeforeTax", m.GetDiscountAppliedBeforeTax())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("dueDate", m.GetDueDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("invoiceDate", m.GetInvoiceDate())
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
    {
        err = writer.WriteObjectValue("payToAddress", m.GetPayToAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("payToContact", m.GetPayToContact())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("payToName", m.GetPayToName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("payToVendorId", m.GetPayToVendorId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("payToVendorNumber", m.GetPayToVendorNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("pricesIncludeTax", m.GetPricesIncludeTax())
        if err != nil {
            return err
        }
    }
    if m.GetPurchaseInvoiceLines() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPurchaseInvoiceLines()))
        for i, v := range m.GetPurchaseInvoiceLines() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("purchaseInvoiceLines", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("shipToAddress", m.GetShipToAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("shipToContact", m.GetShipToContact())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("shipToName", m.GetShipToName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("totalAmountExcludingTax", m.GetTotalAmountExcludingTax())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("totalAmountIncludingTax", m.GetTotalAmountIncludingTax())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("totalTaxAmount", m.GetTotalTaxAmount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("vendorId", m.GetVendorId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("vendorInvoiceNumber", m.GetVendorInvoiceNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("vendorName", m.GetVendorName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("vendorNumber", m.GetVendorNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("vendor", m.GetVendor())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBuyFromAddress sets the buyFromAddress property value. The buyFromAddress property
func (m *PurchaseInvoice) SetBuyFromAddress(value PostalAddressTypeable)() {
    m.buyFromAddress = value
}
// SetCurrency sets the currency property value. The currency property
func (m *PurchaseInvoice) SetCurrency(value Currencyable)() {
    m.currency = value
}
// SetCurrencyCode sets the currencyCode property value. The currencyCode property
func (m *PurchaseInvoice) SetCurrencyCode(value *string)() {
    m.currencyCode = value
}
// SetCurrencyId sets the currencyId property value. The currencyId property
func (m *PurchaseInvoice) SetCurrencyId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.currencyId = value
}
// SetDiscountAmount sets the discountAmount property value. The discountAmount property
func (m *PurchaseInvoice) SetDiscountAmount(value *float64)() {
    m.discountAmount = value
}
// SetDiscountAppliedBeforeTax sets the discountAppliedBeforeTax property value. The discountAppliedBeforeTax property
func (m *PurchaseInvoice) SetDiscountAppliedBeforeTax(value *bool)() {
    m.discountAppliedBeforeTax = value
}
// SetDueDate sets the dueDate property value. The dueDate property
func (m *PurchaseInvoice) SetDueDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.dueDate = value
}
// SetInvoiceDate sets the invoiceDate property value. The invoiceDate property
func (m *PurchaseInvoice) SetInvoiceDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.invoiceDate = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *PurchaseInvoice) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetNumber sets the number property value. The number property
func (m *PurchaseInvoice) SetNumber(value *string)() {
    m.number = value
}
// SetPayToAddress sets the payToAddress property value. The payToAddress property
func (m *PurchaseInvoice) SetPayToAddress(value PostalAddressTypeable)() {
    m.payToAddress = value
}
// SetPayToContact sets the payToContact property value. The payToContact property
func (m *PurchaseInvoice) SetPayToContact(value *string)() {
    m.payToContact = value
}
// SetPayToName sets the payToName property value. The payToName property
func (m *PurchaseInvoice) SetPayToName(value *string)() {
    m.payToName = value
}
// SetPayToVendorId sets the payToVendorId property value. The payToVendorId property
func (m *PurchaseInvoice) SetPayToVendorId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.payToVendorId = value
}
// SetPayToVendorNumber sets the payToVendorNumber property value. The payToVendorNumber property
func (m *PurchaseInvoice) SetPayToVendorNumber(value *string)() {
    m.payToVendorNumber = value
}
// SetPricesIncludeTax sets the pricesIncludeTax property value. The pricesIncludeTax property
func (m *PurchaseInvoice) SetPricesIncludeTax(value *bool)() {
    m.pricesIncludeTax = value
}
// SetPurchaseInvoiceLines sets the purchaseInvoiceLines property value. The purchaseInvoiceLines property
func (m *PurchaseInvoice) SetPurchaseInvoiceLines(value []PurchaseInvoiceLineable)() {
    m.purchaseInvoiceLines = value
}
// SetShipToAddress sets the shipToAddress property value. The shipToAddress property
func (m *PurchaseInvoice) SetShipToAddress(value PostalAddressTypeable)() {
    m.shipToAddress = value
}
// SetShipToContact sets the shipToContact property value. The shipToContact property
func (m *PurchaseInvoice) SetShipToContact(value *string)() {
    m.shipToContact = value
}
// SetShipToName sets the shipToName property value. The shipToName property
func (m *PurchaseInvoice) SetShipToName(value *string)() {
    m.shipToName = value
}
// SetStatus sets the status property value. The status property
func (m *PurchaseInvoice) SetStatus(value *string)() {
    m.status = value
}
// SetTotalAmountExcludingTax sets the totalAmountExcludingTax property value. The totalAmountExcludingTax property
func (m *PurchaseInvoice) SetTotalAmountExcludingTax(value *float64)() {
    m.totalAmountExcludingTax = value
}
// SetTotalAmountIncludingTax sets the totalAmountIncludingTax property value. The totalAmountIncludingTax property
func (m *PurchaseInvoice) SetTotalAmountIncludingTax(value *float64)() {
    m.totalAmountIncludingTax = value
}
// SetTotalTaxAmount sets the totalTaxAmount property value. The totalTaxAmount property
func (m *PurchaseInvoice) SetTotalTaxAmount(value *float64)() {
    m.totalTaxAmount = value
}
// SetVendor sets the vendor property value. The vendor property
func (m *PurchaseInvoice) SetVendor(value Vendor_escapedable)() {
    m.vendor_escaped = value
}
// SetVendorId sets the vendorId property value. The vendorId property
func (m *PurchaseInvoice) SetVendorId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.vendorId = value
}
// SetVendorInvoiceNumber sets the vendorInvoiceNumber property value. The vendorInvoiceNumber property
func (m *PurchaseInvoice) SetVendorInvoiceNumber(value *string)() {
    m.vendorInvoiceNumber = value
}
// SetVendorName sets the vendorName property value. The vendorName property
func (m *PurchaseInvoice) SetVendorName(value *string)() {
    m.vendorName = value
}
// SetVendorNumber sets the vendorNumber property value. The vendorNumber property
func (m *PurchaseInvoice) SetVendorNumber(value *string)() {
    m.vendorNumber = value
}
