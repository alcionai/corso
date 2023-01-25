package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PersonCertification 
type PersonCertification struct {
    ItemFacet
    // The referenceable identifier for the certification.
    certificationId *string
    // Description of the certification.
    description *string
    // Title of the certification.
    displayName *string
    // The date that the certification expires.
    endDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The date that the certification was issued.
    issuedDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // Authority which granted the certification.
    issuingAuthority *string
    // Company which granted the certification.
    issuingCompany *string
    // The date that the certification became valid.
    startDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // URL referencing a thumbnail of the certification.
    thumbnailUrl *string
    // URL referencing the certification.
    webUrl *string
}
// NewPersonCertification instantiates a new PersonCertification and sets the default values.
func NewPersonCertification()(*PersonCertification) {
    m := &PersonCertification{
        ItemFacet: *NewItemFacet(),
    }
    odataTypeValue := "#microsoft.graph.personCertification";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePersonCertificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePersonCertificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPersonCertification(), nil
}
// GetCertificationId gets the certificationId property value. The referenceable identifier for the certification.
func (m *PersonCertification) GetCertificationId()(*string) {
    return m.certificationId
}
// GetDescription gets the description property value. Description of the certification.
func (m *PersonCertification) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Title of the certification.
func (m *PersonCertification) GetDisplayName()(*string) {
    return m.displayName
}
// GetEndDate gets the endDate property value. The date that the certification expires.
func (m *PersonCertification) GetEndDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.endDate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PersonCertification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["certificationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificationId(val)
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
    res["endDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDate(val)
        }
        return nil
    }
    res["issuedDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuedDate(val)
        }
        return nil
    }
    res["issuingAuthority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuingAuthority(val)
        }
        return nil
    }
    res["issuingCompany"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuingCompany(val)
        }
        return nil
    }
    res["startDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDate(val)
        }
        return nil
    }
    res["thumbnailUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThumbnailUrl(val)
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
// GetIssuedDate gets the issuedDate property value. The date that the certification was issued.
func (m *PersonCertification) GetIssuedDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.issuedDate
}
// GetIssuingAuthority gets the issuingAuthority property value. Authority which granted the certification.
func (m *PersonCertification) GetIssuingAuthority()(*string) {
    return m.issuingAuthority
}
// GetIssuingCompany gets the issuingCompany property value. Company which granted the certification.
func (m *PersonCertification) GetIssuingCompany()(*string) {
    return m.issuingCompany
}
// GetStartDate gets the startDate property value. The date that the certification became valid.
func (m *PersonCertification) GetStartDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.startDate
}
// GetThumbnailUrl gets the thumbnailUrl property value. URL referencing a thumbnail of the certification.
func (m *PersonCertification) GetThumbnailUrl()(*string) {
    return m.thumbnailUrl
}
// GetWebUrl gets the webUrl property value. URL referencing the certification.
func (m *PersonCertification) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *PersonCertification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("certificationId", m.GetCertificationId())
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("endDate", m.GetEndDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("issuedDate", m.GetIssuedDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("issuingAuthority", m.GetIssuingAuthority())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("issuingCompany", m.GetIssuingCompany())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("startDate", m.GetStartDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("thumbnailUrl", m.GetThumbnailUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCertificationId sets the certificationId property value. The referenceable identifier for the certification.
func (m *PersonCertification) SetCertificationId(value *string)() {
    m.certificationId = value
}
// SetDescription sets the description property value. Description of the certification.
func (m *PersonCertification) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Title of the certification.
func (m *PersonCertification) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEndDate sets the endDate property value. The date that the certification expires.
func (m *PersonCertification) SetEndDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.endDate = value
}
// SetIssuedDate sets the issuedDate property value. The date that the certification was issued.
func (m *PersonCertification) SetIssuedDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.issuedDate = value
}
// SetIssuingAuthority sets the issuingAuthority property value. Authority which granted the certification.
func (m *PersonCertification) SetIssuingAuthority(value *string)() {
    m.issuingAuthority = value
}
// SetIssuingCompany sets the issuingCompany property value. Company which granted the certification.
func (m *PersonCertification) SetIssuingCompany(value *string)() {
    m.issuingCompany = value
}
// SetStartDate sets the startDate property value. The date that the certification became valid.
func (m *PersonCertification) SetStartDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.startDate = value
}
// SetThumbnailUrl sets the thumbnailUrl property value. URL referencing a thumbnail of the certification.
func (m *PersonCertification) SetThumbnailUrl(value *string)() {
    m.thumbnailUrl = value
}
// SetWebUrl sets the webUrl property value. URL referencing the certification.
func (m *PersonCertification) SetWebUrl(value *string)() {
    m.webUrl = value
}
