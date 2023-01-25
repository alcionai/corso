package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemPatent 
type ItemPatent struct {
    ItemFacet
    // Descpription of the patent or filing.
    description *string
    // Title of the patent or filing.
    displayName *string
    // Indicates the patent is pending.
    isPending *bool
    // The date that the patent was granted.
    issuedDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // Authority which granted the patent.
    issuingAuthority *string
    // The patent number.
    number *string
    // URL referencing the patent or filing.
    webUrl *string
}
// NewItemPatent instantiates a new ItemPatent and sets the default values.
func NewItemPatent()(*ItemPatent) {
    m := &ItemPatent{
        ItemFacet: *NewItemFacet(),
    }
    odataTypeValue := "#microsoft.graph.itemPatent";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateItemPatentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemPatentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPatent(), nil
}
// GetDescription gets the description property value. Descpription of the patent or filing.
func (m *ItemPatent) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Title of the patent or filing.
func (m *ItemPatent) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemPatent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
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
    res["isPending"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPending(val)
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
// GetIsPending gets the isPending property value. Indicates the patent is pending.
func (m *ItemPatent) GetIsPending()(*bool) {
    return m.isPending
}
// GetIssuedDate gets the issuedDate property value. The date that the patent was granted.
func (m *ItemPatent) GetIssuedDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.issuedDate
}
// GetIssuingAuthority gets the issuingAuthority property value. Authority which granted the patent.
func (m *ItemPatent) GetIssuingAuthority()(*string) {
    return m.issuingAuthority
}
// GetNumber gets the number property value. The patent number.
func (m *ItemPatent) GetNumber()(*string) {
    return m.number
}
// GetWebUrl gets the webUrl property value. URL referencing the patent or filing.
func (m *ItemPatent) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *ItemPatent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
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
        err = writer.WriteBoolValue("isPending", m.GetIsPending())
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
        err = writer.WriteStringValue("number", m.GetNumber())
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
// SetDescription sets the description property value. Descpription of the patent or filing.
func (m *ItemPatent) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Title of the patent or filing.
func (m *ItemPatent) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsPending sets the isPending property value. Indicates the patent is pending.
func (m *ItemPatent) SetIsPending(value *bool)() {
    m.isPending = value
}
// SetIssuedDate sets the issuedDate property value. The date that the patent was granted.
func (m *ItemPatent) SetIssuedDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.issuedDate = value
}
// SetIssuingAuthority sets the issuingAuthority property value. Authority which granted the patent.
func (m *ItemPatent) SetIssuingAuthority(value *string)() {
    m.issuingAuthority = value
}
// SetNumber sets the number property value. The patent number.
func (m *ItemPatent) SetNumber(value *string)() {
    m.number = value
}
// SetWebUrl sets the webUrl property value. URL referencing the patent or filing.
func (m *ItemPatent) SetWebUrl(value *string)() {
    m.webUrl = value
}
