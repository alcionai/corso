package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsUpdateCatalogItem windows update catalog item entity
type WindowsUpdateCatalogItem struct {
    Entity
    // The display name for the catalog item.
    displayName *string
    // The last supported date for a catalog item
    endOfSupportDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date the catalog item was released
    releaseDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewWindowsUpdateCatalogItem instantiates a new windowsUpdateCatalogItem and sets the default values.
func NewWindowsUpdateCatalogItem()(*WindowsUpdateCatalogItem) {
    m := &WindowsUpdateCatalogItem{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsUpdateCatalogItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsUpdateCatalogItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.windowsFeatureUpdateCatalogItem":
                        return NewWindowsFeatureUpdateCatalogItem(), nil
                    case "#microsoft.graph.windowsQualityUpdateCatalogItem":
                        return NewWindowsQualityUpdateCatalogItem(), nil
                }
            }
        }
    }
    return NewWindowsUpdateCatalogItem(), nil
}
// GetDisplayName gets the displayName property value. The display name for the catalog item.
func (m *WindowsUpdateCatalogItem) GetDisplayName()(*string) {
    return m.displayName
}
// GetEndOfSupportDate gets the endOfSupportDate property value. The last supported date for a catalog item
func (m *WindowsUpdateCatalogItem) GetEndOfSupportDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endOfSupportDate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsUpdateCatalogItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["endOfSupportDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndOfSupportDate(val)
        }
        return nil
    }
    res["releaseDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReleaseDateTime(val)
        }
        return nil
    }
    return res
}
// GetReleaseDateTime gets the releaseDateTime property value. The date the catalog item was released
func (m *WindowsUpdateCatalogItem) GetReleaseDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.releaseDateTime
}
// Serialize serializes information the current object
func (m *WindowsUpdateCatalogItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteTimeValue("endOfSupportDate", m.GetEndOfSupportDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("releaseDateTime", m.GetReleaseDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The display name for the catalog item.
func (m *WindowsUpdateCatalogItem) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEndOfSupportDate sets the endOfSupportDate property value. The last supported date for a catalog item
func (m *WindowsUpdateCatalogItem) SetEndOfSupportDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endOfSupportDate = value
}
// SetReleaseDateTime sets the releaseDateTime property value. The date the catalog item was released
func (m *WindowsUpdateCatalogItem) SetReleaseDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.releaseDateTime = value
}
