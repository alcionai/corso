package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// TenantTag provides operations to manage the collection of site entities.
type TenantTag struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // The identifier for the account that created the tenant tag. Required. Read-only.
    createdByUserId *string
    // The date and time when the tenant tag was created. Required. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date and time when the tenant tag was deleted. Required. Read-only.
    deletedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description for the tenant tag. Optional. Read-only.
    description *string
    // The display name for the tenant tag. Required. Read-only.
    displayName *string
    // The identifier for the account that lasted on the tenant tag. Optional. Read-only.
    lastActionByUserId *string
    // The date and time the last action was performed against the tenant tag. Optional. Read-only.
    lastActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The collection of managed tenants associated with the tenant tag. Optional.
    tenants []TenantInfoable
}
// NewTenantTag instantiates a new tenantTag and sets the default values.
func NewTenantTag()(*TenantTag) {
    m := &TenantTag{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateTenantTagFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantTagFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenantTag(), nil
}
// GetCreatedByUserId gets the createdByUserId property value. The identifier for the account that created the tenant tag. Required. Read-only.
func (m *TenantTag) GetCreatedByUserId()(*string) {
    return m.createdByUserId
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time when the tenant tag was created. Required. Read-only.
func (m *TenantTag) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDeletedDateTime gets the deletedDateTime property value. The date and time when the tenant tag was deleted. Required. Read-only.
func (m *TenantTag) GetDeletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.deletedDateTime
}
// GetDescription gets the description property value. The description for the tenant tag. Optional. Read-only.
func (m *TenantTag) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name for the tenant tag. Required. Read-only.
func (m *TenantTag) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantTag) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdByUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedByUserId(val)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["deletedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeletedDateTime(val)
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
    res["lastActionByUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActionByUserId(val)
        }
        return nil
    }
    res["lastActionDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActionDateTime(val)
        }
        return nil
    }
    res["tenants"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTenantInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TenantInfoable, len(val))
            for i, v := range val {
                res[i] = v.(TenantInfoable)
            }
            m.SetTenants(res)
        }
        return nil
    }
    return res
}
// GetLastActionByUserId gets the lastActionByUserId property value. The identifier for the account that lasted on the tenant tag. Optional. Read-only.
func (m *TenantTag) GetLastActionByUserId()(*string) {
    return m.lastActionByUserId
}
// GetLastActionDateTime gets the lastActionDateTime property value. The date and time the last action was performed against the tenant tag. Optional. Read-only.
func (m *TenantTag) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastActionDateTime
}
// GetTenants gets the tenants property value. The collection of managed tenants associated with the tenant tag. Optional.
func (m *TenantTag) GetTenants()([]TenantInfoable) {
    return m.tenants
}
// Serialize serializes information the current object
func (m *TenantTag) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("createdByUserId", m.GetCreatedByUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("deletedDateTime", m.GetDeletedDateTime())
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
        err = writer.WriteStringValue("lastActionByUserId", m.GetLastActionByUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastActionDateTime", m.GetLastActionDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetTenants() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTenants()))
        for i, v := range m.GetTenants() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tenants", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedByUserId sets the createdByUserId property value. The identifier for the account that created the tenant tag. Required. Read-only.
func (m *TenantTag) SetCreatedByUserId(value *string)() {
    m.createdByUserId = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time when the tenant tag was created. Required. Read-only.
func (m *TenantTag) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDeletedDateTime sets the deletedDateTime property value. The date and time when the tenant tag was deleted. Required. Read-only.
func (m *TenantTag) SetDeletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.deletedDateTime = value
}
// SetDescription sets the description property value. The description for the tenant tag. Optional. Read-only.
func (m *TenantTag) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name for the tenant tag. Required. Read-only.
func (m *TenantTag) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastActionByUserId sets the lastActionByUserId property value. The identifier for the account that lasted on the tenant tag. Optional. Read-only.
func (m *TenantTag) SetLastActionByUserId(value *string)() {
    m.lastActionByUserId = value
}
// SetLastActionDateTime sets the lastActionDateTime property value. The date and time the last action was performed against the tenant tag. Optional. Read-only.
func (m *TenantTag) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastActionDateTime = value
}
// SetTenants sets the tenants property value. The collection of managed tenants associated with the tenant tag. Optional.
func (m *TenantTag) SetTenants(value []TenantInfoable)() {
    m.tenants = value
}
