package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SalesInvoiceLine provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type SalesInvoiceLine struct {
    Entity
    // The account property
    account Accountable
    // The accountId property
    accountId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The amountExcludingTax property
    amountExcludingTax *float64
    // The amountIncludingTax property
    amountIncludingTax *float64
    // The description property
    description *string
    // The discountAmount property
    discountAmount *float64
    // The discountAppliedBeforeTax property
    discountAppliedBeforeTax *bool
    // The discountPercent property
    discountPercent *float64
    // The documentId property
    documentId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The invoiceDiscountAllocation property
    invoiceDiscountAllocation *float64
    // The item property
    item Itemable
    // The itemId property
    itemId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The lineType property
    lineType *string
    // The netAmount property
    netAmount *float64
    // The netAmountIncludingTax property
    netAmountIncludingTax *float64
    // The netTaxAmount property
    netTaxAmount *float64
    // The quantity property
    quantity *float64
    // The sequence property
    sequence *int32
    // The shipmentDate property
    shipmentDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The taxCode property
    taxCode *string
    // The taxPercent property
    taxPercent *float64
    // The totalTaxAmount property
    totalTaxAmount *float64
    // The unitOfMeasureId property
    unitOfMeasureId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The unitPrice property
    unitPrice *float64
}
// NewSalesInvoiceLine instantiates a new salesInvoiceLine and sets the default values.
func NewSalesInvoiceLine()(*SalesInvoiceLine) {
    m := &SalesInvoiceLine{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSalesInvoiceLineFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSalesInvoiceLineFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSalesInvoiceLine(), nil
}
// GetAccount gets the account property value. The account property
func (m *SalesInvoiceLine) GetAccount()(Accountable) {
    return m.account
}
// GetAccountId gets the accountId property value. The accountId property
func (m *SalesInvoiceLine) GetAccountId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.accountId
}
// GetAmountExcludingTax gets the amountExcludingTax property value. The amountExcludingTax property
func (m *SalesInvoiceLine) GetAmountExcludingTax()(*float64) {
    return m.amountExcludingTax
}
// GetAmountIncludingTax gets the amountIncludingTax property value. The amountIncludingTax property
func (m *SalesInvoiceLine) GetAmountIncludingTax()(*float64) {
    return m.amountIncludingTax
}
// GetDescription gets the description property value. The description property
func (m *SalesInvoiceLine) GetDescription()(*string) {
    return m.description
}
// GetDiscountAmount gets the discountAmount property value. The discountAmount property
func (m *SalesInvoiceLine) GetDiscountAmount()(*float64) {
    return m.discountAmount
}
// GetDiscountAppliedBeforeTax gets the discountAppliedBeforeTax property value. The discountAppliedBeforeTax property
func (m *SalesInvoiceLine) GetDiscountAppliedBeforeTax()(*bool) {
    return m.discountAppliedBeforeTax
}
// GetDiscountPercent gets the discountPercent property value. The discountPercent property
func (m *SalesInvoiceLine) GetDiscountPercent()(*float64) {
    return m.discountPercent
}
// GetDocumentId gets the documentId property value. The documentId property
func (m *SalesInvoiceLine) GetDocumentId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.documentId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SalesInvoiceLine) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["account"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccount(val.(Accountable))
        }
        return nil
    }
    res["accountId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountId(val)
        }
        return nil
    }
    res["amountExcludingTax"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAmountExcludingTax(val)
        }
        return nil
    }
    res["amountIncludingTax"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAmountIncludingTax(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["discountPercent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscountPercent(val)
        }
        return nil
    }
    res["documentId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentId(val)
        }
        return nil
    }
    res["invoiceDiscountAllocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvoiceDiscountAllocation(val)
        }
        return nil
    }
    res["item"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItem(val.(Itemable))
        }
        return nil
    }
    res["itemId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemId(val)
        }
        return nil
    }
    res["lineType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLineType(val)
        }
        return nil
    }
    res["netAmount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetAmount(val)
        }
        return nil
    }
    res["netAmountIncludingTax"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetAmountIncludingTax(val)
        }
        return nil
    }
    res["netTaxAmount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetTaxAmount(val)
        }
        return nil
    }
    res["quantity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuantity(val)
        }
        return nil
    }
    res["sequence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSequence(val)
        }
        return nil
    }
    res["shipmentDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShipmentDate(val)
        }
        return nil
    }
    res["taxCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaxCode(val)
        }
        return nil
    }
    res["taxPercent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaxPercent(val)
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
    res["unitOfMeasureId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnitOfMeasureId(val)
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
// GetInvoiceDiscountAllocation gets the invoiceDiscountAllocation property value. The invoiceDiscountAllocation property
func (m *SalesInvoiceLine) GetInvoiceDiscountAllocation()(*float64) {
    return m.invoiceDiscountAllocation
}
// GetItem gets the item property value. The item property
func (m *SalesInvoiceLine) GetItem()(Itemable) {
    return m.item
}
// GetItemId gets the itemId property value. The itemId property
func (m *SalesInvoiceLine) GetItemId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.itemId
}
// GetLineType gets the lineType property value. The lineType property
func (m *SalesInvoiceLine) GetLineType()(*string) {
    return m.lineType
}
// GetNetAmount gets the netAmount property value. The netAmount property
func (m *SalesInvoiceLine) GetNetAmount()(*float64) {
    return m.netAmount
}
// GetNetAmountIncludingTax gets the netAmountIncludingTax property value. The netAmountIncludingTax property
func (m *SalesInvoiceLine) GetNetAmountIncludingTax()(*float64) {
    return m.netAmountIncludingTax
}
// GetNetTaxAmount gets the netTaxAmount property value. The netTaxAmount property
func (m *SalesInvoiceLine) GetNetTaxAmount()(*float64) {
    return m.netTaxAmount
}
// GetQuantity gets the quantity property value. The quantity property
func (m *SalesInvoiceLine) GetQuantity()(*float64) {
    return m.quantity
}
// GetSequence gets the sequence property value. The sequence property
func (m *SalesInvoiceLine) GetSequence()(*int32) {
    return m.sequence
}
// GetShipmentDate gets the shipmentDate property value. The shipmentDate property
func (m *SalesInvoiceLine) GetShipmentDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.shipmentDate
}
// GetTaxCode gets the taxCode property value. The taxCode property
func (m *SalesInvoiceLine) GetTaxCode()(*string) {
    return m.taxCode
}
// GetTaxPercent gets the taxPercent property value. The taxPercent property
func (m *SalesInvoiceLine) GetTaxPercent()(*float64) {
    return m.taxPercent
}
// GetTotalTaxAmount gets the totalTaxAmount property value. The totalTaxAmount property
func (m *SalesInvoiceLine) GetTotalTaxAmount()(*float64) {
    return m.totalTaxAmount
}
// GetUnitOfMeasureId gets the unitOfMeasureId property value. The unitOfMeasureId property
func (m *SalesInvoiceLine) GetUnitOfMeasureId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.unitOfMeasureId
}
// GetUnitPrice gets the unitPrice property value. The unitPrice property
func (m *SalesInvoiceLine) GetUnitPrice()(*float64) {
    return m.unitPrice
}
// Serialize serializes information the current object
func (m *SalesInvoiceLine) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("account", m.GetAccount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("accountId", m.GetAccountId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("amountExcludingTax", m.GetAmountExcludingTax())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("amountIncludingTax", m.GetAmountIncludingTax())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
        err = writer.WriteFloat64Value("discountPercent", m.GetDiscountPercent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("documentId", m.GetDocumentId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("invoiceDiscountAllocation", m.GetInvoiceDiscountAllocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("item", m.GetItem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("itemId", m.GetItemId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lineType", m.GetLineType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("netAmount", m.GetNetAmount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("netAmountIncludingTax", m.GetNetAmountIncludingTax())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("netTaxAmount", m.GetNetTaxAmount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("quantity", m.GetQuantity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("sequence", m.GetSequence())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("shipmentDate", m.GetShipmentDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("taxCode", m.GetTaxCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("taxPercent", m.GetTaxPercent())
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
        err = writer.WriteUUIDValue("unitOfMeasureId", m.GetUnitOfMeasureId())
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
// SetAccount sets the account property value. The account property
func (m *SalesInvoiceLine) SetAccount(value Accountable)() {
    m.account = value
}
// SetAccountId sets the accountId property value. The accountId property
func (m *SalesInvoiceLine) SetAccountId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.accountId = value
}
// SetAmountExcludingTax sets the amountExcludingTax property value. The amountExcludingTax property
func (m *SalesInvoiceLine) SetAmountExcludingTax(value *float64)() {
    m.amountExcludingTax = value
}
// SetAmountIncludingTax sets the amountIncludingTax property value. The amountIncludingTax property
func (m *SalesInvoiceLine) SetAmountIncludingTax(value *float64)() {
    m.amountIncludingTax = value
}
// SetDescription sets the description property value. The description property
func (m *SalesInvoiceLine) SetDescription(value *string)() {
    m.description = value
}
// SetDiscountAmount sets the discountAmount property value. The discountAmount property
func (m *SalesInvoiceLine) SetDiscountAmount(value *float64)() {
    m.discountAmount = value
}
// SetDiscountAppliedBeforeTax sets the discountAppliedBeforeTax property value. The discountAppliedBeforeTax property
func (m *SalesInvoiceLine) SetDiscountAppliedBeforeTax(value *bool)() {
    m.discountAppliedBeforeTax = value
}
// SetDiscountPercent sets the discountPercent property value. The discountPercent property
func (m *SalesInvoiceLine) SetDiscountPercent(value *float64)() {
    m.discountPercent = value
}
// SetDocumentId sets the documentId property value. The documentId property
func (m *SalesInvoiceLine) SetDocumentId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.documentId = value
}
// SetInvoiceDiscountAllocation sets the invoiceDiscountAllocation property value. The invoiceDiscountAllocation property
func (m *SalesInvoiceLine) SetInvoiceDiscountAllocation(value *float64)() {
    m.invoiceDiscountAllocation = value
}
// SetItem sets the item property value. The item property
func (m *SalesInvoiceLine) SetItem(value Itemable)() {
    m.item = value
}
// SetItemId sets the itemId property value. The itemId property
func (m *SalesInvoiceLine) SetItemId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.itemId = value
}
// SetLineType sets the lineType property value. The lineType property
func (m *SalesInvoiceLine) SetLineType(value *string)() {
    m.lineType = value
}
// SetNetAmount sets the netAmount property value. The netAmount property
func (m *SalesInvoiceLine) SetNetAmount(value *float64)() {
    m.netAmount = value
}
// SetNetAmountIncludingTax sets the netAmountIncludingTax property value. The netAmountIncludingTax property
func (m *SalesInvoiceLine) SetNetAmountIncludingTax(value *float64)() {
    m.netAmountIncludingTax = value
}
// SetNetTaxAmount sets the netTaxAmount property value. The netTaxAmount property
func (m *SalesInvoiceLine) SetNetTaxAmount(value *float64)() {
    m.netTaxAmount = value
}
// SetQuantity sets the quantity property value. The quantity property
func (m *SalesInvoiceLine) SetQuantity(value *float64)() {
    m.quantity = value
}
// SetSequence sets the sequence property value. The sequence property
func (m *SalesInvoiceLine) SetSequence(value *int32)() {
    m.sequence = value
}
// SetShipmentDate sets the shipmentDate property value. The shipmentDate property
func (m *SalesInvoiceLine) SetShipmentDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.shipmentDate = value
}
// SetTaxCode sets the taxCode property value. The taxCode property
func (m *SalesInvoiceLine) SetTaxCode(value *string)() {
    m.taxCode = value
}
// SetTaxPercent sets the taxPercent property value. The taxPercent property
func (m *SalesInvoiceLine) SetTaxPercent(value *float64)() {
    m.taxPercent = value
}
// SetTotalTaxAmount sets the totalTaxAmount property value. The totalTaxAmount property
func (m *SalesInvoiceLine) SetTotalTaxAmount(value *float64)() {
    m.totalTaxAmount = value
}
// SetUnitOfMeasureId sets the unitOfMeasureId property value. The unitOfMeasureId property
func (m *SalesInvoiceLine) SetUnitOfMeasureId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.unitOfMeasureId = value
}
// SetUnitPrice sets the unitPrice property value. The unitPrice property
func (m *SalesInvoiceLine) SetUnitPrice(value *float64)() {
    m.unitPrice = value
}
