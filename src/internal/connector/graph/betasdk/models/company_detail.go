package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CompanyDetail 
type CompanyDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Address of the company.
    address PhysicalAddressable
    // Department Name within a company.
    department *string
    // Company name.
    displayName *string
    // The OdataType property
    odataType *string
    // Office Location of the person referred to.
    officeLocation *string
    // Pronunciation guide for the company name.
    pronunciation *string
    // Link to the company home page.
    webUrl *string
}
// NewCompanyDetail instantiates a new companyDetail and sets the default values.
func NewCompanyDetail()(*CompanyDetail) {
    m := &CompanyDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCompanyDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCompanyDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCompanyDetail(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CompanyDetail) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAddress gets the address property value. Address of the company.
func (m *CompanyDetail) GetAddress()(PhysicalAddressable) {
    return m.address
}
// GetDepartment gets the department property value. Department Name within a company.
func (m *CompanyDetail) GetDepartment()(*string) {
    return m.department
}
// GetDisplayName gets the displayName property value. Company name.
func (m *CompanyDetail) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CompanyDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePhysicalAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val.(PhysicalAddressable))
        }
        return nil
    }
    res["department"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDepartment(val)
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
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["officeLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfficeLocation(val)
        }
        return nil
    }
    res["pronunciation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPronunciation(val)
        }
        return nil
    }
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CompanyDetail) GetOdataType()(*string) {
    return m.odataType
}
// GetOfficeLocation gets the officeLocation property value. Office Location of the person referred to.
func (m *CompanyDetail) GetOfficeLocation()(*string) {
    return m.officeLocation
}
// GetPronunciation gets the pronunciation property value. Pronunciation guide for the company name.
func (m *CompanyDetail) GetPronunciation()(*string) {
    return m.pronunciation
}
// GetWebUrl gets the webUrl property value. Link to the company home page.
func (m *CompanyDetail) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *CompanyDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("department", m.GetDepartment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("officeLocation", m.GetOfficeLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("pronunciation", m.GetPronunciation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("webUrl", m.GetWebUrl())
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
func (m *CompanyDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAddress sets the address property value. Address of the company.
func (m *CompanyDetail) SetAddress(value PhysicalAddressable)() {
    m.address = value
}
// SetDepartment sets the department property value. Department Name within a company.
func (m *CompanyDetail) SetDepartment(value *string)() {
    m.department = value
}
// SetDisplayName sets the displayName property value. Company name.
func (m *CompanyDetail) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CompanyDetail) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOfficeLocation sets the officeLocation property value. Office Location of the person referred to.
func (m *CompanyDetail) SetOfficeLocation(value *string)() {
    m.officeLocation = value
}
// SetPronunciation sets the pronunciation property value. Pronunciation guide for the company name.
func (m *CompanyDetail) SetPronunciation(value *string)() {
    m.pronunciation = value
}
// SetWebUrl sets the webUrl property value. Link to the company home page.
func (m *CompanyDetail) SetWebUrl(value *string)() {
    m.webUrl = value
}
