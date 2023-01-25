package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AgedAccountsPayable 
type AgedAccountsPayable struct {
    Entity
    // The agedAsOfDate property
    agedAsOfDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The balanceDue property
    balanceDue *float64
    // The currencyCode property
    currencyCode *string
    // The currentAmount property
    currentAmount *float64
    // The name property
    name *string
    // The period1Amount property
    period1Amount *float64
    // The period2Amount property
    period2Amount *float64
    // The period3Amount property
    period3Amount *float64
    // The periodLengthFilter property
    periodLengthFilter *string
    // The vendorNumber property
    vendorNumber *string
}
// NewAgedAccountsPayable instantiates a new AgedAccountsPayable and sets the default values.
func NewAgedAccountsPayable()(*AgedAccountsPayable) {
    m := &AgedAccountsPayable{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAgedAccountsPayableFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAgedAccountsPayableFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAgedAccountsPayable(), nil
}
// GetAgedAsOfDate gets the agedAsOfDate property value. The agedAsOfDate property
func (m *AgedAccountsPayable) GetAgedAsOfDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.agedAsOfDate
}
// GetBalanceDue gets the balanceDue property value. The balanceDue property
func (m *AgedAccountsPayable) GetBalanceDue()(*float64) {
    return m.balanceDue
}
// GetCurrencyCode gets the currencyCode property value. The currencyCode property
func (m *AgedAccountsPayable) GetCurrencyCode()(*string) {
    return m.currencyCode
}
// GetCurrentAmount gets the currentAmount property value. The currentAmount property
func (m *AgedAccountsPayable) GetCurrentAmount()(*float64) {
    return m.currentAmount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AgedAccountsPayable) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["agedAsOfDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAgedAsOfDate(val)
        }
        return nil
    }
    res["balanceDue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBalanceDue(val)
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
    res["currentAmount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentAmount(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["period1Amount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeriod1Amount(val)
        }
        return nil
    }
    res["period2Amount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeriod2Amount(val)
        }
        return nil
    }
    res["period3Amount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeriod3Amount(val)
        }
        return nil
    }
    res["periodLengthFilter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeriodLengthFilter(val)
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
// GetName gets the name property value. The name property
func (m *AgedAccountsPayable) GetName()(*string) {
    return m.name
}
// GetPeriod1Amount gets the period1Amount property value. The period1Amount property
func (m *AgedAccountsPayable) GetPeriod1Amount()(*float64) {
    return m.period1Amount
}
// GetPeriod2Amount gets the period2Amount property value. The period2Amount property
func (m *AgedAccountsPayable) GetPeriod2Amount()(*float64) {
    return m.period2Amount
}
// GetPeriod3Amount gets the period3Amount property value. The period3Amount property
func (m *AgedAccountsPayable) GetPeriod3Amount()(*float64) {
    return m.period3Amount
}
// GetPeriodLengthFilter gets the periodLengthFilter property value. The periodLengthFilter property
func (m *AgedAccountsPayable) GetPeriodLengthFilter()(*string) {
    return m.periodLengthFilter
}
// GetVendorNumber gets the vendorNumber property value. The vendorNumber property
func (m *AgedAccountsPayable) GetVendorNumber()(*string) {
    return m.vendorNumber
}
// Serialize serializes information the current object
func (m *AgedAccountsPayable) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteDateOnlyValue("agedAsOfDate", m.GetAgedAsOfDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("balanceDue", m.GetBalanceDue())
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
        err = writer.WriteFloat64Value("currentAmount", m.GetCurrentAmount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("period1Amount", m.GetPeriod1Amount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("period2Amount", m.GetPeriod2Amount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("period3Amount", m.GetPeriod3Amount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("periodLengthFilter", m.GetPeriodLengthFilter())
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
    return nil
}
// SetAgedAsOfDate sets the agedAsOfDate property value. The agedAsOfDate property
func (m *AgedAccountsPayable) SetAgedAsOfDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.agedAsOfDate = value
}
// SetBalanceDue sets the balanceDue property value. The balanceDue property
func (m *AgedAccountsPayable) SetBalanceDue(value *float64)() {
    m.balanceDue = value
}
// SetCurrencyCode sets the currencyCode property value. The currencyCode property
func (m *AgedAccountsPayable) SetCurrencyCode(value *string)() {
    m.currencyCode = value
}
// SetCurrentAmount sets the currentAmount property value. The currentAmount property
func (m *AgedAccountsPayable) SetCurrentAmount(value *float64)() {
    m.currentAmount = value
}
// SetName sets the name property value. The name property
func (m *AgedAccountsPayable) SetName(value *string)() {
    m.name = value
}
// SetPeriod1Amount sets the period1Amount property value. The period1Amount property
func (m *AgedAccountsPayable) SetPeriod1Amount(value *float64)() {
    m.period1Amount = value
}
// SetPeriod2Amount sets the period2Amount property value. The period2Amount property
func (m *AgedAccountsPayable) SetPeriod2Amount(value *float64)() {
    m.period2Amount = value
}
// SetPeriod3Amount sets the period3Amount property value. The period3Amount property
func (m *AgedAccountsPayable) SetPeriod3Amount(value *float64)() {
    m.period3Amount = value
}
// SetPeriodLengthFilter sets the periodLengthFilter property value. The periodLengthFilter property
func (m *AgedAccountsPayable) SetPeriodLengthFilter(value *string)() {
    m.periodLengthFilter = value
}
// SetVendorNumber sets the vendorNumber property value. The vendorNumber property
func (m *AgedAccountsPayable) SetVendorNumber(value *string)() {
    m.vendorNumber = value
}
