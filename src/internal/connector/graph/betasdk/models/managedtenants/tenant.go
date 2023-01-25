package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// Tenant provides operations to manage the collection of site entities.
type Tenant struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // The relationship details for the tenant with the managing entity.
    contract TenantContractable
    // The date and time the tenant was created in the multi-tenant management platform. Optional. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The display name for the tenant. Required. Read-only.
    displayName *string
    // The date and time the tenant was last updated within the multi-tenant management platform. Optional. Read-only.
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
    tenantId *string
    // The onboarding status information for the tenant. Optional. Read-only.
    tenantStatusInformation TenantStatusInformationable
}
// NewTenant instantiates a new tenant and sets the default values.
func NewTenant()(*Tenant) {
    m := &Tenant{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateTenantFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenant(), nil
}
// GetContract gets the contract property value. The relationship details for the tenant with the managing entity.
func (m *Tenant) GetContract()(TenantContractable) {
    return m.contract
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the tenant was created in the multi-tenant management platform. Optional. Read-only.
func (m *Tenant) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDisplayName gets the displayName property value. The display name for the tenant. Required. Read-only.
func (m *Tenant) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Tenant) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["contract"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTenantContractFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContract(val.(TenantContractable))
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
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["tenantStatusInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTenantStatusInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantStatusInformation(val.(TenantStatusInformationable))
        }
        return nil
    }
    return res
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. The date and time the tenant was last updated within the multi-tenant management platform. Optional. Read-only.
func (m *Tenant) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdatedDateTime
}
// GetTenantId gets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *Tenant) GetTenantId()(*string) {
    return m.tenantId
}
// GetTenantStatusInformation gets the tenantStatusInformation property value. The onboarding status information for the tenant. Optional. Read-only.
func (m *Tenant) GetTenantStatusInformation()(TenantStatusInformationable) {
    return m.tenantStatusInformation
}
// Serialize serializes information the current object
func (m *Tenant) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("contract", m.GetContract())
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
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("tenantStatusInformation", m.GetTenantStatusInformation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContract sets the contract property value. The relationship details for the tenant with the managing entity.
func (m *Tenant) SetContract(value TenantContractable)() {
    m.contract = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the tenant was created in the multi-tenant management platform. Optional. Read-only.
func (m *Tenant) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDisplayName sets the displayName property value. The display name for the tenant. Required. Read-only.
func (m *Tenant) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. The date and time the tenant was last updated within the multi-tenant management platform. Optional. Read-only.
func (m *Tenant) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
// SetTenantId sets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *Tenant) SetTenantId(value *string)() {
    m.tenantId = value
}
// SetTenantStatusInformation sets the tenantStatusInformation property value. The onboarding status information for the tenant. Optional. Read-only.
func (m *Tenant) SetTenantStatusInformation(value TenantStatusInformationable)() {
    m.tenantStatusInformation = value
}
