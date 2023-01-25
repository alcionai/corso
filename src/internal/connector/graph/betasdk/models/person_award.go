package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PersonAward 
type PersonAward struct {
    ItemFacet
    // Descpription of the award or honor.
    description *string
    // Name of the award or honor.
    displayName *string
    // The date that the award or honor was granted.
    issuedDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // Authority which granted the award or honor.
    issuingAuthority *string
    // URL referencing a thumbnail of the award or honor.
    thumbnailUrl *string
    // URL referencing the award or honor.
    webUrl *string
}
// NewPersonAward instantiates a new PersonAward and sets the default values.
func NewPersonAward()(*PersonAward) {
    m := &PersonAward{
        ItemFacet: *NewItemFacet(),
    }
    odataTypeValue := "#microsoft.graph.personAward";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePersonAwardFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePersonAwardFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPersonAward(), nil
}
// GetDescription gets the description property value. Descpription of the award or honor.
func (m *PersonAward) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Name of the award or honor.
func (m *PersonAward) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PersonAward) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetIssuedDate gets the issuedDate property value. The date that the award or honor was granted.
func (m *PersonAward) GetIssuedDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.issuedDate
}
// GetIssuingAuthority gets the issuingAuthority property value. Authority which granted the award or honor.
func (m *PersonAward) GetIssuingAuthority()(*string) {
    return m.issuingAuthority
}
// GetThumbnailUrl gets the thumbnailUrl property value. URL referencing a thumbnail of the award or honor.
func (m *PersonAward) GetThumbnailUrl()(*string) {
    return m.thumbnailUrl
}
// GetWebUrl gets the webUrl property value. URL referencing the award or honor.
func (m *PersonAward) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *PersonAward) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetDescription sets the description property value. Descpription of the award or honor.
func (m *PersonAward) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Name of the award or honor.
func (m *PersonAward) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIssuedDate sets the issuedDate property value. The date that the award or honor was granted.
func (m *PersonAward) SetIssuedDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.issuedDate = value
}
// SetIssuingAuthority sets the issuingAuthority property value. Authority which granted the award or honor.
func (m *PersonAward) SetIssuingAuthority(value *string)() {
    m.issuingAuthority = value
}
// SetThumbnailUrl sets the thumbnailUrl property value. URL referencing a thumbnail of the award or honor.
func (m *PersonAward) SetThumbnailUrl(value *string)() {
    m.thumbnailUrl = value
}
// SetWebUrl sets the webUrl property value. URL referencing the award or honor.
func (m *PersonAward) SetWebUrl(value *string)() {
    m.webUrl = value
}
