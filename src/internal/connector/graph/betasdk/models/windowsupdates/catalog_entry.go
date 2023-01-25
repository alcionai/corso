package windowsupdates

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// CatalogEntry provides operations to call the add method.
type CatalogEntry struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // The date on which the content is no longer available to deploy using the service. Read-only.
    deployableUntilDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The display name of the content. Read-only.
    displayName *string
    // The release date for the content. Read-only.
    releaseDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewCatalogEntry instantiates a new catalogEntry and sets the default values.
func NewCatalogEntry()(*CatalogEntry) {
    m := &CatalogEntry{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateCatalogEntryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCatalogEntryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.windowsUpdates.featureUpdateCatalogEntry":
                        return NewFeatureUpdateCatalogEntry(), nil
                    case "#microsoft.graph.windowsUpdates.qualityUpdateCatalogEntry":
                        return NewQualityUpdateCatalogEntry(), nil
                    case "#microsoft.graph.windowsUpdates.softwareUpdateCatalogEntry":
                        return NewSoftwareUpdateCatalogEntry(), nil
                }
            }
        }
    }
    return NewCatalogEntry(), nil
}
// GetDeployableUntilDateTime gets the deployableUntilDateTime property value. The date on which the content is no longer available to deploy using the service. Read-only.
func (m *CatalogEntry) GetDeployableUntilDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.deployableUntilDateTime
}
// GetDisplayName gets the displayName property value. The display name of the content. Read-only.
func (m *CatalogEntry) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CatalogEntry) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deployableUntilDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeployableUntilDateTime(val)
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
// GetReleaseDateTime gets the releaseDateTime property value. The release date for the content. Read-only.
func (m *CatalogEntry) GetReleaseDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.releaseDateTime
}
// Serialize serializes information the current object
func (m *CatalogEntry) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("deployableUntilDateTime", m.GetDeployableUntilDateTime())
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
        err = writer.WriteTimeValue("releaseDateTime", m.GetReleaseDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeployableUntilDateTime sets the deployableUntilDateTime property value. The date on which the content is no longer available to deploy using the service. Read-only.
func (m *CatalogEntry) SetDeployableUntilDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.deployableUntilDateTime = value
}
// SetDisplayName sets the displayName property value. The display name of the content. Read-only.
func (m *CatalogEntry) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetReleaseDateTime sets the releaseDateTime property value. The release date for the content. Read-only.
func (m *CatalogEntry) SetReleaseDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.releaseDateTime = value
}
