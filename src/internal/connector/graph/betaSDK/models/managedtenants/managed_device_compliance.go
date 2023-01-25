package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// ManagedDeviceCompliance provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ManagedDeviceCompliance struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // Compliance state of the device. This property is read-only. Possible values are: unknown, compliant, noncompliant, conflict, error, inGracePeriod, configManager. Optional. Read-only.
    complianceStatus *string
    // Platform of the device. This property is read-only. Possible values are: desktop, windowsRT, winMO6, nokia, windowsPhone, mac, winCE, winEmbedded, iPhone, iPad, iPod, android, iSocConsumer, unix, macMDM, holoLens, surfaceHub, androidForWork, androidEnterprise, windows10x, androidnGMS, chromeOS, linux, blackberry, palm, unknown, cloudPC.  Optional. Read-only.
    deviceType *string
    // The date and time when the grace period will expire. Optional. Read-only.
    inGracePeriodUntilDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
    lastRefreshedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date and time that the device last completed a successful sync with Microsoft Endpoint Manager. Optional. Read-only.
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The identifier for the managed device in Microsoft Endpoint Manager. Optional. Read-only.
    managedDeviceId *string
    // The display name for the managed device. Optional. Read-only.
    managedDeviceName *string
    // The manufacture for the device. Optional. Read-only.
    manufacturer *string
    // The model for the device. Optional. Read-only.
    model *string
    // The description of the operating system for the managed device. Optional. Read-only.
    osDescription *string
    // The version of the operating system for the managed device. Optional. Read-only.
    osVersion *string
    // The type of owner for the managed device. Optional. Read-only.
    ownerType *string
    // The display name for the managed tenant. Optional. Read-only.
    tenantDisplayName *string
    // The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
    tenantId *string
}
// NewManagedDeviceCompliance instantiates a new managedDeviceCompliance and sets the default values.
func NewManagedDeviceCompliance()(*ManagedDeviceCompliance) {
    m := &ManagedDeviceCompliance{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateManagedDeviceComplianceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDeviceComplianceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDeviceCompliance(), nil
}
// GetComplianceStatus gets the complianceStatus property value. Compliance state of the device. This property is read-only. Possible values are: unknown, compliant, noncompliant, conflict, error, inGracePeriod, configManager. Optional. Read-only.
func (m *ManagedDeviceCompliance) GetComplianceStatus()(*string) {
    return m.complianceStatus
}
// GetDeviceType gets the deviceType property value. Platform of the device. This property is read-only. Possible values are: desktop, windowsRT, winMO6, nokia, windowsPhone, mac, winCE, winEmbedded, iPhone, iPad, iPod, android, iSocConsumer, unix, macMDM, holoLens, surfaceHub, androidForWork, androidEnterprise, windows10x, androidnGMS, chromeOS, linux, blackberry, palm, unknown, cloudPC.  Optional. Read-only.
func (m *ManagedDeviceCompliance) GetDeviceType()(*string) {
    return m.deviceType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDeviceCompliance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["complianceStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceStatus(val)
        }
        return nil
    }
    res["deviceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceType(val)
        }
        return nil
    }
    res["inGracePeriodUntilDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInGracePeriodUntilDateTime(val)
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
    res["lastSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncDateTime(val)
        }
        return nil
    }
    res["managedDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceId(val)
        }
        return nil
    }
    res["managedDeviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceName(val)
        }
        return nil
    }
    res["manufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["osDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsDescription(val)
        }
        return nil
    }
    res["osVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    res["ownerType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerType(val)
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
// GetInGracePeriodUntilDateTime gets the inGracePeriodUntilDateTime property value. The date and time when the grace period will expire. Optional. Read-only.
func (m *ManagedDeviceCompliance) GetInGracePeriodUntilDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.inGracePeriodUntilDateTime
}
// GetLastRefreshedDateTime gets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
func (m *ManagedDeviceCompliance) GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastRefreshedDateTime
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. The date and time that the device last completed a successful sync with Microsoft Endpoint Manager. Optional. Read-only.
func (m *ManagedDeviceCompliance) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastSyncDateTime
}
// GetManagedDeviceId gets the managedDeviceId property value. The identifier for the managed device in Microsoft Endpoint Manager. Optional. Read-only.
func (m *ManagedDeviceCompliance) GetManagedDeviceId()(*string) {
    return m.managedDeviceId
}
// GetManagedDeviceName gets the managedDeviceName property value. The display name for the managed device. Optional. Read-only.
func (m *ManagedDeviceCompliance) GetManagedDeviceName()(*string) {
    return m.managedDeviceName
}
// GetManufacturer gets the manufacturer property value. The manufacture for the device. Optional. Read-only.
func (m *ManagedDeviceCompliance) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetModel gets the model property value. The model for the device. Optional. Read-only.
func (m *ManagedDeviceCompliance) GetModel()(*string) {
    return m.model
}
// GetOsDescription gets the osDescription property value. The description of the operating system for the managed device. Optional. Read-only.
func (m *ManagedDeviceCompliance) GetOsDescription()(*string) {
    return m.osDescription
}
// GetOsVersion gets the osVersion property value. The version of the operating system for the managed device. Optional. Read-only.
func (m *ManagedDeviceCompliance) GetOsVersion()(*string) {
    return m.osVersion
}
// GetOwnerType gets the ownerType property value. The type of owner for the managed device. Optional. Read-only.
func (m *ManagedDeviceCompliance) GetOwnerType()(*string) {
    return m.ownerType
}
// GetTenantDisplayName gets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *ManagedDeviceCompliance) GetTenantDisplayName()(*string) {
    return m.tenantDisplayName
}
// GetTenantId gets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *ManagedDeviceCompliance) GetTenantId()(*string) {
    return m.tenantId
}
// Serialize serializes information the current object
func (m *ManagedDeviceCompliance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("complianceStatus", m.GetComplianceStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceType", m.GetDeviceType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("inGracePeriodUntilDateTime", m.GetInGracePeriodUntilDateTime())
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
        err = writer.WriteTimeValue("lastSyncDateTime", m.GetLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceId", m.GetManagedDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceName", m.GetManagedDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osDescription", m.GetOsDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osVersion", m.GetOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ownerType", m.GetOwnerType())
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
// SetComplianceStatus sets the complianceStatus property value. Compliance state of the device. This property is read-only. Possible values are: unknown, compliant, noncompliant, conflict, error, inGracePeriod, configManager. Optional. Read-only.
func (m *ManagedDeviceCompliance) SetComplianceStatus(value *string)() {
    m.complianceStatus = value
}
// SetDeviceType sets the deviceType property value. Platform of the device. This property is read-only. Possible values are: desktop, windowsRT, winMO6, nokia, windowsPhone, mac, winCE, winEmbedded, iPhone, iPad, iPod, android, iSocConsumer, unix, macMDM, holoLens, surfaceHub, androidForWork, androidEnterprise, windows10x, androidnGMS, chromeOS, linux, blackberry, palm, unknown, cloudPC.  Optional. Read-only.
func (m *ManagedDeviceCompliance) SetDeviceType(value *string)() {
    m.deviceType = value
}
// SetInGracePeriodUntilDateTime sets the inGracePeriodUntilDateTime property value. The date and time when the grace period will expire. Optional. Read-only.
func (m *ManagedDeviceCompliance) SetInGracePeriodUntilDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.inGracePeriodUntilDateTime = value
}
// SetLastRefreshedDateTime sets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
func (m *ManagedDeviceCompliance) SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastRefreshedDateTime = value
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. The date and time that the device last completed a successful sync with Microsoft Endpoint Manager. Optional. Read-only.
func (m *ManagedDeviceCompliance) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
// SetManagedDeviceId sets the managedDeviceId property value. The identifier for the managed device in Microsoft Endpoint Manager. Optional. Read-only.
func (m *ManagedDeviceCompliance) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
// SetManagedDeviceName sets the managedDeviceName property value. The display name for the managed device. Optional. Read-only.
func (m *ManagedDeviceCompliance) SetManagedDeviceName(value *string)() {
    m.managedDeviceName = value
}
// SetManufacturer sets the manufacturer property value. The manufacture for the device. Optional. Read-only.
func (m *ManagedDeviceCompliance) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetModel sets the model property value. The model for the device. Optional. Read-only.
func (m *ManagedDeviceCompliance) SetModel(value *string)() {
    m.model = value
}
// SetOsDescription sets the osDescription property value. The description of the operating system for the managed device. Optional. Read-only.
func (m *ManagedDeviceCompliance) SetOsDescription(value *string)() {
    m.osDescription = value
}
// SetOsVersion sets the osVersion property value. The version of the operating system for the managed device. Optional. Read-only.
func (m *ManagedDeviceCompliance) SetOsVersion(value *string)() {
    m.osVersion = value
}
// SetOwnerType sets the ownerType property value. The type of owner for the managed device. Optional. Read-only.
func (m *ManagedDeviceCompliance) SetOwnerType(value *string)() {
    m.ownerType = value
}
// SetTenantDisplayName sets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *ManagedDeviceCompliance) SetTenantDisplayName(value *string)() {
    m.tenantDisplayName = value
}
// SetTenantId sets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *ManagedDeviceCompliance) SetTenantId(value *string)() {
    m.tenantId = value
}
