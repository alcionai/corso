package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PaymentTerm 
type PaymentTerm struct {
    Entity
    // The calculateDiscountOnCreditMemos property
    calculateDiscountOnCreditMemos *bool
    // The code property
    code *string
    // The discountDateCalculation property
    discountDateCalculation *string
    // The discountPercent property
    discountPercent *float64
    // The displayName property
    displayName *string
    // The dueDateCalculation property
    dueDateCalculation *string
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewPaymentTerm instantiates a new paymentTerm and sets the default values.
func NewPaymentTerm()(*PaymentTerm) {
    m := &PaymentTerm{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePaymentTermFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePaymentTermFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPaymentTerm(), nil
}
// GetCalculateDiscountOnCreditMemos gets the calculateDiscountOnCreditMemos property value. The calculateDiscountOnCreditMemos property
func (m *PaymentTerm) GetCalculateDiscountOnCreditMemos()(*bool) {
    return m.calculateDiscountOnCreditMemos
}
// GetCode gets the code property value. The code property
func (m *PaymentTerm) GetCode()(*string) {
    return m.code
}
// GetDiscountDateCalculation gets the discountDateCalculation property value. The discountDateCalculation property
func (m *PaymentTerm) GetDiscountDateCalculation()(*string) {
    return m.discountDateCalculation
}
// GetDiscountPercent gets the discountPercent property value. The discountPercent property
func (m *PaymentTerm) GetDiscountPercent()(*float64) {
    return m.discountPercent
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *PaymentTerm) GetDisplayName()(*string) {
    return m.displayName
}
// GetDueDateCalculation gets the dueDateCalculation property value. The dueDateCalculation property
func (m *PaymentTerm) GetDueDateCalculation()(*string) {
    return m.dueDateCalculation
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PaymentTerm) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["calculateDiscountOnCreditMemos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalculateDiscountOnCreditMemos(val)
        }
        return nil
    }
    res["code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
        }
        return nil
    }
    res["discountDateCalculation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscountDateCalculation(val)
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
    res["dueDateCalculation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueDateCalculation(val)
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
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *PaymentTerm) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// Serialize serializes information the current object
func (m *PaymentTerm) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("calculateDiscountOnCreditMemos", m.GetCalculateDiscountOnCreditMemos())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("discountDateCalculation", m.GetDiscountDateCalculation())
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("dueDateCalculation", m.GetDueDateCalculation())
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
    return nil
}
// SetCalculateDiscountOnCreditMemos sets the calculateDiscountOnCreditMemos property value. The calculateDiscountOnCreditMemos property
func (m *PaymentTerm) SetCalculateDiscountOnCreditMemos(value *bool)() {
    m.calculateDiscountOnCreditMemos = value
}
// SetCode sets the code property value. The code property
func (m *PaymentTerm) SetCode(value *string)() {
    m.code = value
}
// SetDiscountDateCalculation sets the discountDateCalculation property value. The discountDateCalculation property
func (m *PaymentTerm) SetDiscountDateCalculation(value *string)() {
    m.discountDateCalculation = value
}
// SetDiscountPercent sets the discountPercent property value. The discountPercent property
func (m *PaymentTerm) SetDiscountPercent(value *float64)() {
    m.discountPercent = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *PaymentTerm) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDueDateCalculation sets the dueDateCalculation property value. The dueDateCalculation property
func (m *PaymentTerm) SetDueDateCalculation(value *string)() {
    m.dueDateCalculation = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *PaymentTerm) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
