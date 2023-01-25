package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSSoftwareUpdateStateSummary macOS software update state summary for a device and user
type MacOSSoftwareUpdateStateSummary struct {
    Entity
    // Human readable name of the software update
    displayName *string
    // Last date time the report for this device and product key was updated.
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Product key of the software update.
    productKey *string
    // MacOS Software Update State
    state *MacOSSoftwareUpdateState
    // MacOS Software Update Category
    updateCategory *MacOSSoftwareUpdateCategory
    // Version of the software update
    updateVersion *string
}
// NewMacOSSoftwareUpdateStateSummary instantiates a new macOSSoftwareUpdateStateSummary and sets the default values.
func NewMacOSSoftwareUpdateStateSummary()(*MacOSSoftwareUpdateStateSummary) {
    m := &MacOSSoftwareUpdateStateSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMacOSSoftwareUpdateStateSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSSoftwareUpdateStateSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSSoftwareUpdateStateSummary(), nil
}
// GetDisplayName gets the displayName property value. Human readable name of the software update
func (m *MacOSSoftwareUpdateStateSummary) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSSoftwareUpdateStateSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["lastUpdatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedDateTime(val)
        }
        return nil
    }
    res["productKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductKey(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMacOSSoftwareUpdateState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*MacOSSoftwareUpdateState))
        }
        return nil
    }
    res["updateCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMacOSSoftwareUpdateCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateCategory(val.(*MacOSSoftwareUpdateCategory))
        }
        return nil
    }
    res["updateVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateVersion(val)
        }
        return nil
    }
    return res
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. Last date time the report for this device and product key was updated.
func (m *MacOSSoftwareUpdateStateSummary) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdatedDateTime
}
// GetProductKey gets the productKey property value. Product key of the software update.
func (m *MacOSSoftwareUpdateStateSummary) GetProductKey()(*string) {
    return m.productKey
}
// GetState gets the state property value. MacOS Software Update State
func (m *MacOSSoftwareUpdateStateSummary) GetState()(*MacOSSoftwareUpdateState) {
    return m.state
}
// GetUpdateCategory gets the updateCategory property value. MacOS Software Update Category
func (m *MacOSSoftwareUpdateStateSummary) GetUpdateCategory()(*MacOSSoftwareUpdateCategory) {
    return m.updateCategory
}
// GetUpdateVersion gets the updateVersion property value. Version of the software update
func (m *MacOSSoftwareUpdateStateSummary) GetUpdateVersion()(*string) {
    return m.updateVersion
}
// Serialize serializes information the current object
func (m *MacOSSoftwareUpdateStateSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("productKey", m.GetProductKey())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUpdateCategory() != nil {
        cast := (*m.GetUpdateCategory()).String()
        err = writer.WriteStringValue("updateCategory", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("updateVersion", m.GetUpdateVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. Human readable name of the software update
func (m *MacOSSoftwareUpdateStateSummary) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. Last date time the report for this device and product key was updated.
func (m *MacOSSoftwareUpdateStateSummary) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
// SetProductKey sets the productKey property value. Product key of the software update.
func (m *MacOSSoftwareUpdateStateSummary) SetProductKey(value *string)() {
    m.productKey = value
}
// SetState sets the state property value. MacOS Software Update State
func (m *MacOSSoftwareUpdateStateSummary) SetState(value *MacOSSoftwareUpdateState)() {
    m.state = value
}
// SetUpdateCategory sets the updateCategory property value. MacOS Software Update Category
func (m *MacOSSoftwareUpdateStateSummary) SetUpdateCategory(value *MacOSSoftwareUpdateCategory)() {
    m.updateCategory = value
}
// SetUpdateVersion sets the updateVersion property value. Version of the software update
func (m *MacOSSoftwareUpdateStateSummary) SetUpdateVersion(value *string)() {
    m.updateVersion = value
}
