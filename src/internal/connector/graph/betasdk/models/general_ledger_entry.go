package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GeneralLedgerEntry provides operations to call the add method.
type GeneralLedgerEntry struct {
    Entity
    // The account property
    account Accountable
    // The accountId property
    accountId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The accountNumber property
    accountNumber *string
    // The creditAmount property
    creditAmount *float64
    // The debitAmount property
    debitAmount *float64
    // The description property
    description *string
    // The documentNumber property
    documentNumber *string
    // The documentType property
    documentType *string
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The postingDate property
    postingDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
}
// NewGeneralLedgerEntry instantiates a new generalLedgerEntry and sets the default values.
func NewGeneralLedgerEntry()(*GeneralLedgerEntry) {
    m := &GeneralLedgerEntry{
        Entity: *NewEntity(),
    }
    return m
}
// CreateGeneralLedgerEntryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGeneralLedgerEntryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGeneralLedgerEntry(), nil
}
// GetAccount gets the account property value. The account property
func (m *GeneralLedgerEntry) GetAccount()(Accountable) {
    return m.account
}
// GetAccountId gets the accountId property value. The accountId property
func (m *GeneralLedgerEntry) GetAccountId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.accountId
}
// GetAccountNumber gets the accountNumber property value. The accountNumber property
func (m *GeneralLedgerEntry) GetAccountNumber()(*string) {
    return m.accountNumber
}
// GetCreditAmount gets the creditAmount property value. The creditAmount property
func (m *GeneralLedgerEntry) GetCreditAmount()(*float64) {
    return m.creditAmount
}
// GetDebitAmount gets the debitAmount property value. The debitAmount property
func (m *GeneralLedgerEntry) GetDebitAmount()(*float64) {
    return m.debitAmount
}
// GetDescription gets the description property value. The description property
func (m *GeneralLedgerEntry) GetDescription()(*string) {
    return m.description
}
// GetDocumentNumber gets the documentNumber property value. The documentNumber property
func (m *GeneralLedgerEntry) GetDocumentNumber()(*string) {
    return m.documentNumber
}
// GetDocumentType gets the documentType property value. The documentType property
func (m *GeneralLedgerEntry) GetDocumentType()(*string) {
    return m.documentType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GeneralLedgerEntry) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["accountNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountNumber(val)
        }
        return nil
    }
    res["creditAmount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreditAmount(val)
        }
        return nil
    }
    res["debitAmount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDebitAmount(val)
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
    res["documentNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentNumber(val)
        }
        return nil
    }
    res["documentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentType(val)
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
    res["postingDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostingDate(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *GeneralLedgerEntry) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetPostingDate gets the postingDate property value. The postingDate property
func (m *GeneralLedgerEntry) GetPostingDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.postingDate
}
// Serialize serializes information the current object
func (m *GeneralLedgerEntry) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("accountNumber", m.GetAccountNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("creditAmount", m.GetCreditAmount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("debitAmount", m.GetDebitAmount())
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
        err = writer.WriteStringValue("documentNumber", m.GetDocumentNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("documentType", m.GetDocumentType())
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
        err = writer.WriteDateOnlyValue("postingDate", m.GetPostingDate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccount sets the account property value. The account property
func (m *GeneralLedgerEntry) SetAccount(value Accountable)() {
    m.account = value
}
// SetAccountId sets the accountId property value. The accountId property
func (m *GeneralLedgerEntry) SetAccountId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.accountId = value
}
// SetAccountNumber sets the accountNumber property value. The accountNumber property
func (m *GeneralLedgerEntry) SetAccountNumber(value *string)() {
    m.accountNumber = value
}
// SetCreditAmount sets the creditAmount property value. The creditAmount property
func (m *GeneralLedgerEntry) SetCreditAmount(value *float64)() {
    m.creditAmount = value
}
// SetDebitAmount sets the debitAmount property value. The debitAmount property
func (m *GeneralLedgerEntry) SetDebitAmount(value *float64)() {
    m.debitAmount = value
}
// SetDescription sets the description property value. The description property
func (m *GeneralLedgerEntry) SetDescription(value *string)() {
    m.description = value
}
// SetDocumentNumber sets the documentNumber property value. The documentNumber property
func (m *GeneralLedgerEntry) SetDocumentNumber(value *string)() {
    m.documentNumber = value
}
// SetDocumentType sets the documentType property value. The documentType property
func (m *GeneralLedgerEntry) SetDocumentType(value *string)() {
    m.documentType = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *GeneralLedgerEntry) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetPostingDate sets the postingDate property value. The postingDate property
func (m *GeneralLedgerEntry) SetPostingDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.postingDate = value
}
