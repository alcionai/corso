package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// CloudPcOverview provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type CloudPcOverview struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
    lastRefreshedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number of cloud PC connections that have a status of failed. Optional. Read-only.
    numberOfCloudPcConnectionStatusFailed *int32
    // The number of cloud PC connections that have a status of passed. Optional. Read-only.
    numberOfCloudPcConnectionStatusPassed *int32
    // The number of cloud PC connections that have a status of pending. Optional. Read-only.
    numberOfCloudPcConnectionStatusPending *int32
    // The number of cloud PC connections that have a status of running. Optional. Read-only.
    numberOfCloudPcConnectionStatusRunning *int32
    // The number of cloud PC connections that have a status of unknownFutureValue. Optional. Read-only.
    numberOfCloudPcConnectionStatusUnkownFutureValue *int32
    // The number of cloud PCs that have a status of deprovisioning. Optional. Read-only.
    numberOfCloudPcStatusDeprovisioning *int32
    // The number of cloud PCs that have a status of failed. Optional. Read-only.
    numberOfCloudPcStatusFailed *int32
    // The number of cloud PCs that have a status of inGracePeriod. Optional. Read-only.
    numberOfCloudPcStatusInGracePeriod *int32
    // The number of cloud PCs that have a status of notProvisioned. Optional. Read-only.
    numberOfCloudPcStatusNotProvisioned *int32
    // The number of cloud PCs that have a status of provisioned. Optional. Read-only.
    numberOfCloudPcStatusProvisioned *int32
    // The number of cloud PCs that have a status of provisioning. Optional. Read-only.
    numberOfCloudPcStatusProvisioning *int32
    // The number of cloud PCs that have a status of unknown. Optional. Read-only.
    numberOfCloudPcStatusUnknown *int32
    // The number of cloud PCs that have a status of upgrading. Optional. Read-only.
    numberOfCloudPcStatusUpgrading *int32
    // The display name for the managed tenant. Optional. Read-only.
    tenantDisplayName *string
    // The tenantId property
    tenantId *string
    // The total number of cloud PC devices that have the Business SKU. Optional. Read-only.
    totalBusinessLicenses *int32
    // The total number of cloud PC connection statuses for the given managed tenant. Optional. Read-only.
    totalCloudPcConnectionStatus *int32
    // The total number of cloud PC statues for the given managed tenant. Optional. Read-only.
    totalCloudPcStatus *int32
    // The total number of cloud PC devices that have the Enterprise SKU. Optional. Read-only.
    totalEnterpriseLicenses *int32
}
// NewCloudPcOverview instantiates a new cloudPcOverview and sets the default values.
func NewCloudPcOverview()(*CloudPcOverview) {
    m := &CloudPcOverview{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateCloudPcOverviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcOverviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcOverview(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcOverview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["numberOfCloudPcConnectionStatusFailed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcConnectionStatusFailed(val)
        }
        return nil
    }
    res["numberOfCloudPcConnectionStatusPassed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcConnectionStatusPassed(val)
        }
        return nil
    }
    res["numberOfCloudPcConnectionStatusPending"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcConnectionStatusPending(val)
        }
        return nil
    }
    res["numberOfCloudPcConnectionStatusRunning"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcConnectionStatusRunning(val)
        }
        return nil
    }
    res["numberOfCloudPcConnectionStatusUnkownFutureValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcConnectionStatusUnkownFutureValue(val)
        }
        return nil
    }
    res["numberOfCloudPcStatusDeprovisioning"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcStatusDeprovisioning(val)
        }
        return nil
    }
    res["numberOfCloudPcStatusFailed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcStatusFailed(val)
        }
        return nil
    }
    res["numberOfCloudPcStatusInGracePeriod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcStatusInGracePeriod(val)
        }
        return nil
    }
    res["numberOfCloudPcStatusNotProvisioned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcStatusNotProvisioned(val)
        }
        return nil
    }
    res["numberOfCloudPcStatusProvisioned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcStatusProvisioned(val)
        }
        return nil
    }
    res["numberOfCloudPcStatusProvisioning"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcStatusProvisioning(val)
        }
        return nil
    }
    res["numberOfCloudPcStatusUnknown"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcStatusUnknown(val)
        }
        return nil
    }
    res["numberOfCloudPcStatusUpgrading"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfCloudPcStatusUpgrading(val)
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
    res["totalBusinessLicenses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalBusinessLicenses(val)
        }
        return nil
    }
    res["totalCloudPcConnectionStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCloudPcConnectionStatus(val)
        }
        return nil
    }
    res["totalCloudPcStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCloudPcStatus(val)
        }
        return nil
    }
    res["totalEnterpriseLicenses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalEnterpriseLicenses(val)
        }
        return nil
    }
    return res
}
// GetLastRefreshedDateTime gets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
func (m *CloudPcOverview) GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastRefreshedDateTime
}
// GetNumberOfCloudPcConnectionStatusFailed gets the numberOfCloudPcConnectionStatusFailed property value. The number of cloud PC connections that have a status of failed. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusFailed()(*int32) {
    return m.numberOfCloudPcConnectionStatusFailed
}
// GetNumberOfCloudPcConnectionStatusPassed gets the numberOfCloudPcConnectionStatusPassed property value. The number of cloud PC connections that have a status of passed. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusPassed()(*int32) {
    return m.numberOfCloudPcConnectionStatusPassed
}
// GetNumberOfCloudPcConnectionStatusPending gets the numberOfCloudPcConnectionStatusPending property value. The number of cloud PC connections that have a status of pending. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusPending()(*int32) {
    return m.numberOfCloudPcConnectionStatusPending
}
// GetNumberOfCloudPcConnectionStatusRunning gets the numberOfCloudPcConnectionStatusRunning property value. The number of cloud PC connections that have a status of running. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusRunning()(*int32) {
    return m.numberOfCloudPcConnectionStatusRunning
}
// GetNumberOfCloudPcConnectionStatusUnkownFutureValue gets the numberOfCloudPcConnectionStatusUnkownFutureValue property value. The number of cloud PC connections that have a status of unknownFutureValue. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcConnectionStatusUnkownFutureValue()(*int32) {
    return m.numberOfCloudPcConnectionStatusUnkownFutureValue
}
// GetNumberOfCloudPcStatusDeprovisioning gets the numberOfCloudPcStatusDeprovisioning property value. The number of cloud PCs that have a status of deprovisioning. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusDeprovisioning()(*int32) {
    return m.numberOfCloudPcStatusDeprovisioning
}
// GetNumberOfCloudPcStatusFailed gets the numberOfCloudPcStatusFailed property value. The number of cloud PCs that have a status of failed. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusFailed()(*int32) {
    return m.numberOfCloudPcStatusFailed
}
// GetNumberOfCloudPcStatusInGracePeriod gets the numberOfCloudPcStatusInGracePeriod property value. The number of cloud PCs that have a status of inGracePeriod. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusInGracePeriod()(*int32) {
    return m.numberOfCloudPcStatusInGracePeriod
}
// GetNumberOfCloudPcStatusNotProvisioned gets the numberOfCloudPcStatusNotProvisioned property value. The number of cloud PCs that have a status of notProvisioned. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusNotProvisioned()(*int32) {
    return m.numberOfCloudPcStatusNotProvisioned
}
// GetNumberOfCloudPcStatusProvisioned gets the numberOfCloudPcStatusProvisioned property value. The number of cloud PCs that have a status of provisioned. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusProvisioned()(*int32) {
    return m.numberOfCloudPcStatusProvisioned
}
// GetNumberOfCloudPcStatusProvisioning gets the numberOfCloudPcStatusProvisioning property value. The number of cloud PCs that have a status of provisioning. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusProvisioning()(*int32) {
    return m.numberOfCloudPcStatusProvisioning
}
// GetNumberOfCloudPcStatusUnknown gets the numberOfCloudPcStatusUnknown property value. The number of cloud PCs that have a status of unknown. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusUnknown()(*int32) {
    return m.numberOfCloudPcStatusUnknown
}
// GetNumberOfCloudPcStatusUpgrading gets the numberOfCloudPcStatusUpgrading property value. The number of cloud PCs that have a status of upgrading. Optional. Read-only.
func (m *CloudPcOverview) GetNumberOfCloudPcStatusUpgrading()(*int32) {
    return m.numberOfCloudPcStatusUpgrading
}
// GetTenantDisplayName gets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *CloudPcOverview) GetTenantDisplayName()(*string) {
    return m.tenantDisplayName
}
// GetTenantId gets the tenantId property value. The tenantId property
func (m *CloudPcOverview) GetTenantId()(*string) {
    return m.tenantId
}
// GetTotalBusinessLicenses gets the totalBusinessLicenses property value. The total number of cloud PC devices that have the Business SKU. Optional. Read-only.
func (m *CloudPcOverview) GetTotalBusinessLicenses()(*int32) {
    return m.totalBusinessLicenses
}
// GetTotalCloudPcConnectionStatus gets the totalCloudPcConnectionStatus property value. The total number of cloud PC connection statuses for the given managed tenant. Optional. Read-only.
func (m *CloudPcOverview) GetTotalCloudPcConnectionStatus()(*int32) {
    return m.totalCloudPcConnectionStatus
}
// GetTotalCloudPcStatus gets the totalCloudPcStatus property value. The total number of cloud PC statues for the given managed tenant. Optional. Read-only.
func (m *CloudPcOverview) GetTotalCloudPcStatus()(*int32) {
    return m.totalCloudPcStatus
}
// GetTotalEnterpriseLicenses gets the totalEnterpriseLicenses property value. The total number of cloud PC devices that have the Enterprise SKU. Optional. Read-only.
func (m *CloudPcOverview) GetTotalEnterpriseLicenses()(*int32) {
    return m.totalEnterpriseLicenses
}
// Serialize serializes information the current object
func (m *CloudPcOverview) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("lastRefreshedDateTime", m.GetLastRefreshedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberOfCloudPcConnectionStatusFailed", m.GetNumberOfCloudPcConnectionStatusFailed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberOfCloudPcConnectionStatusPassed", m.GetNumberOfCloudPcConnectionStatusPassed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberOfCloudPcConnectionStatusPending", m.GetNumberOfCloudPcConnectionStatusPending())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberOfCloudPcConnectionStatusRunning", m.GetNumberOfCloudPcConnectionStatusRunning())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberOfCloudPcConnectionStatusUnkownFutureValue", m.GetNumberOfCloudPcConnectionStatusUnkownFutureValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberOfCloudPcStatusDeprovisioning", m.GetNumberOfCloudPcStatusDeprovisioning())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberOfCloudPcStatusFailed", m.GetNumberOfCloudPcStatusFailed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberOfCloudPcStatusInGracePeriod", m.GetNumberOfCloudPcStatusInGracePeriod())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberOfCloudPcStatusNotProvisioned", m.GetNumberOfCloudPcStatusNotProvisioned())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberOfCloudPcStatusProvisioned", m.GetNumberOfCloudPcStatusProvisioned())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberOfCloudPcStatusProvisioning", m.GetNumberOfCloudPcStatusProvisioning())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberOfCloudPcStatusUnknown", m.GetNumberOfCloudPcStatusUnknown())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("numberOfCloudPcStatusUpgrading", m.GetNumberOfCloudPcStatusUpgrading())
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
    {
        err = writer.WriteInt32Value("totalBusinessLicenses", m.GetTotalBusinessLicenses())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalCloudPcConnectionStatus", m.GetTotalCloudPcConnectionStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalCloudPcStatus", m.GetTotalCloudPcStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalEnterpriseLicenses", m.GetTotalEnterpriseLicenses())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLastRefreshedDateTime sets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
func (m *CloudPcOverview) SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastRefreshedDateTime = value
}
// SetNumberOfCloudPcConnectionStatusFailed sets the numberOfCloudPcConnectionStatusFailed property value. The number of cloud PC connections that have a status of failed. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusFailed(value *int32)() {
    m.numberOfCloudPcConnectionStatusFailed = value
}
// SetNumberOfCloudPcConnectionStatusPassed sets the numberOfCloudPcConnectionStatusPassed property value. The number of cloud PC connections that have a status of passed. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusPassed(value *int32)() {
    m.numberOfCloudPcConnectionStatusPassed = value
}
// SetNumberOfCloudPcConnectionStatusPending sets the numberOfCloudPcConnectionStatusPending property value. The number of cloud PC connections that have a status of pending. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusPending(value *int32)() {
    m.numberOfCloudPcConnectionStatusPending = value
}
// SetNumberOfCloudPcConnectionStatusRunning sets the numberOfCloudPcConnectionStatusRunning property value. The number of cloud PC connections that have a status of running. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusRunning(value *int32)() {
    m.numberOfCloudPcConnectionStatusRunning = value
}
// SetNumberOfCloudPcConnectionStatusUnkownFutureValue sets the numberOfCloudPcConnectionStatusUnkownFutureValue property value. The number of cloud PC connections that have a status of unknownFutureValue. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcConnectionStatusUnkownFutureValue(value *int32)() {
    m.numberOfCloudPcConnectionStatusUnkownFutureValue = value
}
// SetNumberOfCloudPcStatusDeprovisioning sets the numberOfCloudPcStatusDeprovisioning property value. The number of cloud PCs that have a status of deprovisioning. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusDeprovisioning(value *int32)() {
    m.numberOfCloudPcStatusDeprovisioning = value
}
// SetNumberOfCloudPcStatusFailed sets the numberOfCloudPcStatusFailed property value. The number of cloud PCs that have a status of failed. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusFailed(value *int32)() {
    m.numberOfCloudPcStatusFailed = value
}
// SetNumberOfCloudPcStatusInGracePeriod sets the numberOfCloudPcStatusInGracePeriod property value. The number of cloud PCs that have a status of inGracePeriod. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusInGracePeriod(value *int32)() {
    m.numberOfCloudPcStatusInGracePeriod = value
}
// SetNumberOfCloudPcStatusNotProvisioned sets the numberOfCloudPcStatusNotProvisioned property value. The number of cloud PCs that have a status of notProvisioned. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusNotProvisioned(value *int32)() {
    m.numberOfCloudPcStatusNotProvisioned = value
}
// SetNumberOfCloudPcStatusProvisioned sets the numberOfCloudPcStatusProvisioned property value. The number of cloud PCs that have a status of provisioned. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusProvisioned(value *int32)() {
    m.numberOfCloudPcStatusProvisioned = value
}
// SetNumberOfCloudPcStatusProvisioning sets the numberOfCloudPcStatusProvisioning property value. The number of cloud PCs that have a status of provisioning. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusProvisioning(value *int32)() {
    m.numberOfCloudPcStatusProvisioning = value
}
// SetNumberOfCloudPcStatusUnknown sets the numberOfCloudPcStatusUnknown property value. The number of cloud PCs that have a status of unknown. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusUnknown(value *int32)() {
    m.numberOfCloudPcStatusUnknown = value
}
// SetNumberOfCloudPcStatusUpgrading sets the numberOfCloudPcStatusUpgrading property value. The number of cloud PCs that have a status of upgrading. Optional. Read-only.
func (m *CloudPcOverview) SetNumberOfCloudPcStatusUpgrading(value *int32)() {
    m.numberOfCloudPcStatusUpgrading = value
}
// SetTenantDisplayName sets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *CloudPcOverview) SetTenantDisplayName(value *string)() {
    m.tenantDisplayName = value
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *CloudPcOverview) SetTenantId(value *string)() {
    m.tenantId = value
}
// SetTotalBusinessLicenses sets the totalBusinessLicenses property value. The total number of cloud PC devices that have the Business SKU. Optional. Read-only.
func (m *CloudPcOverview) SetTotalBusinessLicenses(value *int32)() {
    m.totalBusinessLicenses = value
}
// SetTotalCloudPcConnectionStatus sets the totalCloudPcConnectionStatus property value. The total number of cloud PC connection statuses for the given managed tenant. Optional. Read-only.
func (m *CloudPcOverview) SetTotalCloudPcConnectionStatus(value *int32)() {
    m.totalCloudPcConnectionStatus = value
}
// SetTotalCloudPcStatus sets the totalCloudPcStatus property value. The total number of cloud PC statues for the given managed tenant. Optional. Read-only.
func (m *CloudPcOverview) SetTotalCloudPcStatus(value *int32)() {
    m.totalCloudPcStatus = value
}
// SetTotalEnterpriseLicenses sets the totalEnterpriseLicenses property value. The total number of cloud PC devices that have the Enterprise SKU. Optional. Read-only.
func (m *CloudPcOverview) SetTotalEnterpriseLicenses(value *int32)() {
    m.totalEnterpriseLicenses = value
}
