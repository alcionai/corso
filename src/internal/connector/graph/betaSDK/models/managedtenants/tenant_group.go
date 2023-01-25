package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// TenantGroup provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type TenantGroup struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // A flag indicating whether all managed tenant are included in the tenant group. Required. Read-only.
    allTenantsIncluded *bool
    // The display name for the tenant group. Optional. Read-only.
    displayName *string
    // The collection of management action associated with the tenant group. Optional. Read-only.
    managementActions []ManagementActionInfoable
    // The collection of management intents associated with the tenant group. Optional. Read-only.
    managementIntents []ManagementIntentInfoable
    // The collection of managed tenant identifiers include in the tenant group. Optional. Read-only.
    tenantIds []string
}
// NewTenantGroup instantiates a new tenantGroup and sets the default values.
func NewTenantGroup()(*TenantGroup) {
    m := &TenantGroup{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateTenantGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenantGroup(), nil
}
// GetAllTenantsIncluded gets the allTenantsIncluded property value. A flag indicating whether all managed tenant are included in the tenant group. Required. Read-only.
func (m *TenantGroup) GetAllTenantsIncluded()(*bool) {
    return m.allTenantsIncluded
}
// GetDisplayName gets the displayName property value. The display name for the tenant group. Optional. Read-only.
func (m *TenantGroup) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allTenantsIncluded"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllTenantsIncluded(val)
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
    res["managementActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementActionInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementActionInfoable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementActionInfoable)
            }
            m.SetManagementActions(res)
        }
        return nil
    }
    res["managementIntents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementIntentInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementIntentInfoable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementIntentInfoable)
            }
            m.SetManagementIntents(res)
        }
        return nil
    }
    res["tenantIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTenantIds(res)
        }
        return nil
    }
    return res
}
// GetManagementActions gets the managementActions property value. The collection of management action associated with the tenant group. Optional. Read-only.
func (m *TenantGroup) GetManagementActions()([]ManagementActionInfoable) {
    return m.managementActions
}
// GetManagementIntents gets the managementIntents property value. The collection of management intents associated with the tenant group. Optional. Read-only.
func (m *TenantGroup) GetManagementIntents()([]ManagementIntentInfoable) {
    return m.managementIntents
}
// GetTenantIds gets the tenantIds property value. The collection of managed tenant identifiers include in the tenant group. Optional. Read-only.
func (m *TenantGroup) GetTenantIds()([]string) {
    return m.tenantIds
}
// Serialize serializes information the current object
func (m *TenantGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allTenantsIncluded", m.GetAllTenantsIncluded())
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
    if m.GetManagementActions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagementActions()))
        for i, v := range m.GetManagementActions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managementActions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementIntents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagementIntents()))
        for i, v := range m.GetManagementIntents() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managementIntents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenantIds() != nil {
        err = writer.WriteCollectionOfStringValues("tenantIds", m.GetTenantIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllTenantsIncluded sets the allTenantsIncluded property value. A flag indicating whether all managed tenant are included in the tenant group. Required. Read-only.
func (m *TenantGroup) SetAllTenantsIncluded(value *bool)() {
    m.allTenantsIncluded = value
}
// SetDisplayName sets the displayName property value. The display name for the tenant group. Optional. Read-only.
func (m *TenantGroup) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetManagementActions sets the managementActions property value. The collection of management action associated with the tenant group. Optional. Read-only.
func (m *TenantGroup) SetManagementActions(value []ManagementActionInfoable)() {
    m.managementActions = value
}
// SetManagementIntents sets the managementIntents property value. The collection of management intents associated with the tenant group. Optional. Read-only.
func (m *TenantGroup) SetManagementIntents(value []ManagementIntentInfoable)() {
    m.managementIntents = value
}
// SetTenantIds sets the tenantIds property value. The collection of managed tenant identifiers include in the tenant group. Optional. Read-only.
func (m *TenantGroup) SetTenantIds(value []string)() {
    m.tenantIds = value
}
