package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CompanyInformation provides operations to manage the collection of site entities.
type CompanyInformation struct {
    Entity
    // The address property
    address PostalAddressTypeable
    // The currencyCode property
    currencyCode *string
    // The currentFiscalYearStartDate property
    currentFiscalYearStartDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The displayName property
    displayName *string
    // The email property
    email *string
    // The faxNumber property
    faxNumber *string
    // The industry property
    industry *string
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The phoneNumber property
    phoneNumber *string
    // The picture property
    picture []byte
    // The taxRegistrationNumber property
    taxRegistrationNumber *string
    // The website property
    website *string
}
// NewCompanyInformation instantiates a new companyInformation and sets the default values.
func NewCompanyInformation()(*CompanyInformation) {
    m := &CompanyInformation{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCompanyInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCompanyInformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCompanyInformation(), nil
}
// GetAddress gets the address property value. The address property
func (m *CompanyInformation) GetAddress()(PostalAddressTypeable) {
    return m.address
}
// GetCurrencyCode gets the currencyCode property value. The currencyCode property
func (m *CompanyInformation) GetCurrencyCode()(*string) {
    return m.currencyCode
}
// GetCurrentFiscalYearStartDate gets the currentFiscalYearStartDate property value. The currentFiscalYearStartDate property
func (m *CompanyInformation) GetCurrentFiscalYearStartDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.currentFiscalYearStartDate
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *CompanyInformation) GetDisplayName()(*string) {
    return m.displayName
}
// GetEmail gets the email property value. The email property
func (m *CompanyInformation) GetEmail()(*string) {
    return m.email
}
// GetFaxNumber gets the faxNumber property value. The faxNumber property
func (m *CompanyInformation) GetFaxNumber()(*string) {
    return m.faxNumber
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CompanyInformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePostalAddressTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val.(PostalAddressTypeable))
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
    res["currentFiscalYearStartDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentFiscalYearStartDate(val)
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
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["faxNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFaxNumber(val)
        }
        return nil
    }
    res["industry"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndustry(val)
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
    res["phoneNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneNumber(val)
        }
        return nil
    }
    res["picture"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPicture(val)
        }
        return nil
    }
    res["taxRegistrationNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTaxRegistrationNumber(val)
        }
        return nil
    }
    res["website"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebsite(val)
        }
        return nil
    }
    return res
}
// GetIndustry gets the industry property value. The industry property
func (m *CompanyInformation) GetIndustry()(*string) {
    return m.industry
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *CompanyInformation) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetPhoneNumber gets the phoneNumber property value. The phoneNumber property
func (m *CompanyInformation) GetPhoneNumber()(*string) {
    return m.phoneNumber
}
// GetPicture gets the picture property value. The picture property
func (m *CompanyInformation) GetPicture()([]byte) {
    return m.picture
}
// GetTaxRegistrationNumber gets the taxRegistrationNumber property value. The taxRegistrationNumber property
func (m *CompanyInformation) GetTaxRegistrationNumber()(*string) {
    return m.taxRegistrationNumber
}
// GetWebsite gets the website property value. The website property
func (m *CompanyInformation) GetWebsite()(*string) {
    return m.website
}
// Serialize serializes information the current object
func (m *CompanyInformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("address", m.GetAddress())
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
        err = writer.WriteDateOnlyValue("currentFiscalYearStartDate", m.GetCurrentFiscalYearStartDate())
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
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("faxNumber", m.GetFaxNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("industry", m.GetIndustry())
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
        err = writer.WriteStringValue("phoneNumber", m.GetPhoneNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("picture", m.GetPicture())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("taxRegistrationNumber", m.GetTaxRegistrationNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("website", m.GetWebsite())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddress sets the address property value. The address property
func (m *CompanyInformation) SetAddress(value PostalAddressTypeable)() {
    m.address = value
}
// SetCurrencyCode sets the currencyCode property value. The currencyCode property
func (m *CompanyInformation) SetCurrencyCode(value *string)() {
    m.currencyCode = value
}
// SetCurrentFiscalYearStartDate sets the currentFiscalYearStartDate property value. The currentFiscalYearStartDate property
func (m *CompanyInformation) SetCurrentFiscalYearStartDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.currentFiscalYearStartDate = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *CompanyInformation) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEmail sets the email property value. The email property
func (m *CompanyInformation) SetEmail(value *string)() {
    m.email = value
}
// SetFaxNumber sets the faxNumber property value. The faxNumber property
func (m *CompanyInformation) SetFaxNumber(value *string)() {
    m.faxNumber = value
}
// SetIndustry sets the industry property value. The industry property
func (m *CompanyInformation) SetIndustry(value *string)() {
    m.industry = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *CompanyInformation) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetPhoneNumber sets the phoneNumber property value. The phoneNumber property
func (m *CompanyInformation) SetPhoneNumber(value *string)() {
    m.phoneNumber = value
}
// SetPicture sets the picture property value. The picture property
func (m *CompanyInformation) SetPicture(value []byte)() {
    m.picture = value
}
// SetTaxRegistrationNumber sets the taxRegistrationNumber property value. The taxRegistrationNumber property
func (m *CompanyInformation) SetTaxRegistrationNumber(value *string)() {
    m.taxRegistrationNumber = value
}
// SetWebsite sets the website property value. The website property
func (m *CompanyInformation) SetWebsite(value *string)() {
    m.website = value
}
