package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// CloudPcConnection provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type CloudPcConnection struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // The display name of the cloud PC connection. Required. Read-only.
    displayName *string
    // The health status of the cloud PC connection. Possible values are: pending, running, passed, failed, unknownFutureValue.  Required. Read-only.
    healthCheckStatus *string
    // Date and time the entity was last updated in the multi-tenant management platform. Required. Read-only.
    lastRefreshedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The display name for the managed tenant. Required. Read-only.
    tenantDisplayName *string
    // The Azure Active Directory tenant identifier for the managed tenant. Required. Read-only.
    tenantId *string
}
// NewCloudPcConnection instantiates a new cloudPcConnection and sets the default values.
func NewCloudPcConnection()(*CloudPcConnection) {
    m := &CloudPcConnection{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateCloudPcConnectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcConnectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcConnection(), nil
}
// GetDisplayName gets the displayName property value. The display name of the cloud PC connection. Required. Read-only.
func (m *CloudPcConnection) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcConnection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["healthCheckStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthCheckStatus(val)
        }
        return nil
    }
    res["lastRefreshedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRefreshedDateTime(val)
        }
        return nil
    }
    res["tenantDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantDisplayName(val)
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
    return res
}
// GetHealthCheckStatus gets the healthCheckStatus property value. The health status of the cloud PC connection. Possible values are: pending, running, passed, failed, unknownFutureValue.  Required. Read-only.
func (m *CloudPcConnection) GetHealthCheckStatus()(*string) {
    return m.healthCheckStatus
}
// GetLastRefreshedDateTime gets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Required. Read-only.
func (m *CloudPcConnection) GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastRefreshedDateTime
}
// GetTenantDisplayName gets the tenantDisplayName property value. The display name for the managed tenant. Required. Read-only.
func (m *CloudPcConnection) GetTenantDisplayName()(*string) {
    return m.tenantDisplayName
}
// GetTenantId gets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Required. Read-only.
func (m *CloudPcConnection) GetTenantId()(*string) {
    return m.tenantId
}
// Serialize serializes information the current object
func (m *CloudPcConnection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("healthCheckStatus", m.GetHealthCheckStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastRefreshedDateTime", m.GetLastRefreshedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantDisplayName", m.GetTenantDisplayName())
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
    return nil
}
// SetDisplayName sets the displayName property value. The display name of the cloud PC connection. Required. Read-only.
func (m *CloudPcConnection) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetHealthCheckStatus sets the healthCheckStatus property value. The health status of the cloud PC connection. Possible values are: pending, running, passed, failed, unknownFutureValue.  Required. Read-only.
func (m *CloudPcConnection) SetHealthCheckStatus(value *string)() {
    m.healthCheckStatus = value
}
// SetLastRefreshedDateTime sets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Required. Read-only.
func (m *CloudPcConnection) SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastRefreshedDateTime = value
}
// SetTenantDisplayName sets the tenantDisplayName property value. The display name for the managed tenant. Required. Read-only.
func (m *CloudPcConnection) SetTenantDisplayName(value *string)() {
    m.tenantDisplayName = value
}
// SetTenantId sets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Required. Read-only.
func (m *CloudPcConnection) SetTenantId(value *string)() {
    m.tenantId = value
}
