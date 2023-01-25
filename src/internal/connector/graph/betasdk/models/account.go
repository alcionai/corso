package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Account 
type Account struct {
    Entity
    // The blocked property
    blocked *bool
    // The category property
    category *string
    // The displayName property
    displayName *string
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number property
    number *string
    // The subCategory property
    subCategory *string
}
// NewAccount instantiates a new Account and sets the default values.
func NewAccount()(*Account) {
    m := &Account{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccount(), nil
}
// GetBlocked gets the blocked property value. The blocked property
func (m *Account) GetBlocked()(*bool) {
    return m.blocked
}
// GetCategory gets the category property value. The category property
func (m *Account) GetCategory()(*string) {
    return m.category
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *Account) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Account) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["blocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlocked(val)
        }
        return nil
    }
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val)
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
    res["subCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubCategory(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Account) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetNumber gets the number property value. The number property
func (m *Account) GetNumber()(*string) {
    return m.number
}
// GetSubCategory gets the subCategory property value. The subCategory property
func (m *Account) GetSubCategory()(*string) {
    return m.subCategory
}
// Serialize serializes information the current object
func (m *Account) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("blocked", m.GetBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("category", m.GetCategory())
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
        err = writer.WriteStringValue("subCategory", m.GetSubCategory())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBlocked sets the blocked property value. The blocked property
func (m *Account) SetBlocked(value *bool)() {
    m.blocked = value
}
// SetCategory sets the category property value. The category property
func (m *Account) SetCategory(value *string)() {
    m.category = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *Account) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Account) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetNumber sets the number property value. The number property
func (m *Account) SetNumber(value *string)() {
    m.number = value
}
// SetSubCategory sets the subCategory property value. The subCategory property
func (m *Account) SetSubCategory(value *string)() {
    m.subCategory = value
}
