package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// ManagedDeviceComplianceTrend provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ManagedDeviceComplianceTrend struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // The number of devices with a compliant status. Required. Read-only.
    compliantDeviceCount *int32
    // The number of devices manged by Configuration Manager. Required. Read-only.
    configManagerDeviceCount *int32
    // The date and time compliance snapshot was performed. Required. Read-only.
    countDateTime *string
    // The number of devices with an error status. Required. Read-only.
    errorDeviceCount *int32
    // The number of devices that are in a grace period status. Required. Read-only.
    inGracePeriodDeviceCount *int32
    // The number of devices that are in a non-compliant status. Required. Read-only.
    noncompliantDeviceCount *int32
    // The display name for the managed tenant. Optional. Read-only.
    tenantDisplayName *string
    // The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
    tenantId *string
    // The number of devices in an unknown status. Required. Read-only.
    unknownDeviceCount *int32
}
// NewManagedDeviceComplianceTrend instantiates a new managedDeviceComplianceTrend and sets the default values.
func NewManagedDeviceComplianceTrend()(*ManagedDeviceComplianceTrend) {
    m := &ManagedDeviceComplianceTrend{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateManagedDeviceComplianceTrendFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDeviceComplianceTrendFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDeviceComplianceTrend(), nil
}
// GetCompliantDeviceCount gets the compliantDeviceCount property value. The number of devices with a compliant status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetCompliantDeviceCount()(*int32) {
    return m.compliantDeviceCount
}
// GetConfigManagerDeviceCount gets the configManagerDeviceCount property value. The number of devices manged by Configuration Manager. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetConfigManagerDeviceCount()(*int32) {
    return m.configManagerDeviceCount
}
// GetCountDateTime gets the countDateTime property value. The date and time compliance snapshot was performed. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetCountDateTime()(*string) {
    return m.countDateTime
}
// GetErrorDeviceCount gets the errorDeviceCount property value. The number of devices with an error status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetErrorDeviceCount()(*int32) {
    return m.errorDeviceCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDeviceComplianceTrend) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["compliantDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliantDeviceCount(val)
        }
        return nil
    }
    res["configManagerDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigManagerDeviceCount(val)
        }
        return nil
    }
    res["countDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountDateTime(val)
        }
        return nil
    }
    res["errorDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorDeviceCount(val)
        }
        return nil
    }
    res["inGracePeriodDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInGracePeriodDeviceCount(val)
        }
        return nil
    }
    res["noncompliantDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNoncompliantDeviceCount(val)
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
    res["unknownDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnknownDeviceCount(val)
        }
        return nil
    }
    return res
}
// GetInGracePeriodDeviceCount gets the inGracePeriodDeviceCount property value. The number of devices that are in a grace period status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetInGracePeriodDeviceCount()(*int32) {
    return m.inGracePeriodDeviceCount
}
// GetNoncompliantDeviceCount gets the noncompliantDeviceCount property value. The number of devices that are in a non-compliant status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetNoncompliantDeviceCount()(*int32) {
    return m.noncompliantDeviceCount
}
// GetTenantDisplayName gets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *ManagedDeviceComplianceTrend) GetTenantDisplayName()(*string) {
    return m.tenantDisplayName
}
// GetTenantId gets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *ManagedDeviceComplianceTrend) GetTenantId()(*string) {
    return m.tenantId
}
// GetUnknownDeviceCount gets the unknownDeviceCount property value. The number of devices in an unknown status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) GetUnknownDeviceCount()(*int32) {
    return m.unknownDeviceCount
}
// Serialize serializes information the current object
func (m *ManagedDeviceComplianceTrend) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("compliantDeviceCount", m.GetCompliantDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("configManagerDeviceCount", m.GetConfigManagerDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("countDateTime", m.GetCountDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("errorDeviceCount", m.GetErrorDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("inGracePeriodDeviceCount", m.GetInGracePeriodDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("noncompliantDeviceCount", m.GetNoncompliantDeviceCount())
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
        err = writer.WriteInt32Value("unknownDeviceCount", m.GetUnknownDeviceCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompliantDeviceCount sets the compliantDeviceCount property value. The number of devices with a compliant status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) SetCompliantDeviceCount(value *int32)() {
    m.compliantDeviceCount = value
}
// SetConfigManagerDeviceCount sets the configManagerDeviceCount property value. The number of devices manged by Configuration Manager. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) SetConfigManagerDeviceCount(value *int32)() {
    m.configManagerDeviceCount = value
}
// SetCountDateTime sets the countDateTime property value. The date and time compliance snapshot was performed. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) SetCountDateTime(value *string)() {
    m.countDateTime = value
}
// SetErrorDeviceCount sets the errorDeviceCount property value. The number of devices with an error status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) SetErrorDeviceCount(value *int32)() {
    m.errorDeviceCount = value
}
// SetInGracePeriodDeviceCount sets the inGracePeriodDeviceCount property value. The number of devices that are in a grace period status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) SetInGracePeriodDeviceCount(value *int32)() {
    m.inGracePeriodDeviceCount = value
}
// SetNoncompliantDeviceCount sets the noncompliantDeviceCount property value. The number of devices that are in a non-compliant status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) SetNoncompliantDeviceCount(value *int32)() {
    m.noncompliantDeviceCount = value
}
// SetTenantDisplayName sets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *ManagedDeviceComplianceTrend) SetTenantDisplayName(value *string)() {
    m.tenantDisplayName = value
}
// SetTenantId sets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *ManagedDeviceComplianceTrend) SetTenantId(value *string)() {
    m.tenantId = value
}
// SetUnknownDeviceCount sets the unknownDeviceCount property value. The number of devices in an unknown status. Required. Read-only.
func (m *ManagedDeviceComplianceTrend) SetUnknownDeviceCount(value *int32)() {
    m.unknownDeviceCount = value
}
