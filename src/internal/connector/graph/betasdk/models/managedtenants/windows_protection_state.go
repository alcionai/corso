package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// WindowsProtectionState provides operations to call the add method.
type WindowsProtectionState struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // The anti-malware version for the managed device. Optional. Read-only.
    antiMalwareVersion *string
    // A flag indicating whether attention is required for the managed device. Optional. Read-only.
    attentionRequired *bool
    // A flag indicating whether the managed device has been deleted. Optional. Read-only.
    deviceDeleted *bool
    // The date and time the device property has been refreshed. Optional. Read-only.
    devicePropertyRefreshDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The anti-virus engine version for the managed device. Optional. Read-only.
    engineVersion *string
    // A flag indicating whether quick scan is overdue for the managed device. Optional. Read-only.
    fullScanOverdue *bool
    // A flag indicating whether full scan is overdue for the managed device. Optional. Read-only.
    fullScanRequired *bool
    // The date and time a full scan was completed. Optional. Read-only.
    lastFullScanDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The version anti-malware version used to perform the last full scan. Optional. Read-only.
    lastFullScanSignatureVersion *string
    // The date and time a quick scan was completed. Optional. Read-only.
    lastQuickScanDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The version anti-malware version used to perform the last full scan. Optional. Read-only.
    lastQuickScanSignatureVersion *string
    // Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
    lastRefreshedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date and time the protection state was last reported for the managed device. Optional. Read-only.
    lastReportedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A flag indicating whether malware protection is enabled for the managed device. Optional. Read-only.
    malwareProtectionEnabled *bool
    // The health state for the managed device. Optional. Read-only.
    managedDeviceHealthState *string
    // The unique identifier for the managed device. Optional. Read-only.
    managedDeviceId *string
    // The display name for the managed device. Optional. Read-only.
    managedDeviceName *string
    // A flag indicating whether the network inspection system is enabled. Optional. Read-only.
    networkInspectionSystemEnabled *bool
    // A flag indicating weather a quick scan is overdue. Optional. Read-only.
    quickScanOverdue *bool
    // A flag indicating whether real time protection is enabled. Optional. Read-only.
    realTimeProtectionEnabled *bool
    // A flag indicating whether a reboot is required. Optional. Read-only.
    rebootRequired *bool
    // A flag indicating whether an signature update is overdue. Optional. Read-only.
    signatureUpdateOverdue *bool
    // The signature version for the managed device. Optional. Read-only.
    signatureVersion *string
    // The display name for the managed tenant. Optional. Read-only.
    tenantDisplayName *string
    // The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
    tenantId *string
}
// NewWindowsProtectionState instantiates a new windowsProtectionState and sets the default values.
func NewWindowsProtectionState()(*WindowsProtectionState) {
    m := &WindowsProtectionState{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateWindowsProtectionStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsProtectionStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsProtectionState(), nil
}
// GetAntiMalwareVersion gets the antiMalwareVersion property value. The anti-malware version for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) GetAntiMalwareVersion()(*string) {
    return m.antiMalwareVersion
}
// GetAttentionRequired gets the attentionRequired property value. A flag indicating whether attention is required for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) GetAttentionRequired()(*bool) {
    return m.attentionRequired
}
// GetDeviceDeleted gets the deviceDeleted property value. A flag indicating whether the managed device has been deleted. Optional. Read-only.
func (m *WindowsProtectionState) GetDeviceDeleted()(*bool) {
    return m.deviceDeleted
}
// GetDevicePropertyRefreshDateTime gets the devicePropertyRefreshDateTime property value. The date and time the device property has been refreshed. Optional. Read-only.
func (m *WindowsProtectionState) GetDevicePropertyRefreshDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.devicePropertyRefreshDateTime
}
// GetEngineVersion gets the engineVersion property value. The anti-virus engine version for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) GetEngineVersion()(*string) {
    return m.engineVersion
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsProtectionState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["antiMalwareVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAntiMalwareVersion(val)
        }
        return nil
    }
    res["attentionRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttentionRequired(val)
        }
        return nil
    }
    res["deviceDeleted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceDeleted(val)
        }
        return nil
    }
    res["devicePropertyRefreshDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevicePropertyRefreshDateTime(val)
        }
        return nil
    }
    res["engineVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEngineVersion(val)
        }
        return nil
    }
    res["fullScanOverdue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFullScanOverdue(val)
        }
        return nil
    }
    res["fullScanRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFullScanRequired(val)
        }
        return nil
    }
    res["lastFullScanDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastFullScanDateTime(val)
        }
        return nil
    }
    res["lastFullScanSignatureVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastFullScanSignatureVersion(val)
        }
        return nil
    }
    res["lastQuickScanDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastQuickScanDateTime(val)
        }
        return nil
    }
    res["lastQuickScanSignatureVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastQuickScanSignatureVersion(val)
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
    res["lastReportedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastReportedDateTime(val)
        }
        return nil
    }
    res["malwareProtectionEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMalwareProtectionEnabled(val)
        }
        return nil
    }
    res["managedDeviceHealthState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceHealthState(val)
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
    res["networkInspectionSystemEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkInspectionSystemEnabled(val)
        }
        return nil
    }
    res["quickScanOverdue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuickScanOverdue(val)
        }
        return nil
    }
    res["realTimeProtectionEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRealTimeProtectionEnabled(val)
        }
        return nil
    }
    res["rebootRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRebootRequired(val)
        }
        return nil
    }
    res["signatureUpdateOverdue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignatureUpdateOverdue(val)
        }
        return nil
    }
    res["signatureVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignatureVersion(val)
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
// GetFullScanOverdue gets the fullScanOverdue property value. A flag indicating whether quick scan is overdue for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) GetFullScanOverdue()(*bool) {
    return m.fullScanOverdue
}
// GetFullScanRequired gets the fullScanRequired property value. A flag indicating whether full scan is overdue for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) GetFullScanRequired()(*bool) {
    return m.fullScanRequired
}
// GetLastFullScanDateTime gets the lastFullScanDateTime property value. The date and time a full scan was completed. Optional. Read-only.
func (m *WindowsProtectionState) GetLastFullScanDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastFullScanDateTime
}
// GetLastFullScanSignatureVersion gets the lastFullScanSignatureVersion property value. The version anti-malware version used to perform the last full scan. Optional. Read-only.
func (m *WindowsProtectionState) GetLastFullScanSignatureVersion()(*string) {
    return m.lastFullScanSignatureVersion
}
// GetLastQuickScanDateTime gets the lastQuickScanDateTime property value. The date and time a quick scan was completed. Optional. Read-only.
func (m *WindowsProtectionState) GetLastQuickScanDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastQuickScanDateTime
}
// GetLastQuickScanSignatureVersion gets the lastQuickScanSignatureVersion property value. The version anti-malware version used to perform the last full scan. Optional. Read-only.
func (m *WindowsProtectionState) GetLastQuickScanSignatureVersion()(*string) {
    return m.lastQuickScanSignatureVersion
}
// GetLastRefreshedDateTime gets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
func (m *WindowsProtectionState) GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastRefreshedDateTime
}
// GetLastReportedDateTime gets the lastReportedDateTime property value. The date and time the protection state was last reported for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) GetLastReportedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastReportedDateTime
}
// GetMalwareProtectionEnabled gets the malwareProtectionEnabled property value. A flag indicating whether malware protection is enabled for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) GetMalwareProtectionEnabled()(*bool) {
    return m.malwareProtectionEnabled
}
// GetManagedDeviceHealthState gets the managedDeviceHealthState property value. The health state for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) GetManagedDeviceHealthState()(*string) {
    return m.managedDeviceHealthState
}
// GetManagedDeviceId gets the managedDeviceId property value. The unique identifier for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) GetManagedDeviceId()(*string) {
    return m.managedDeviceId
}
// GetManagedDeviceName gets the managedDeviceName property value. The display name for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) GetManagedDeviceName()(*string) {
    return m.managedDeviceName
}
// GetNetworkInspectionSystemEnabled gets the networkInspectionSystemEnabled property value. A flag indicating whether the network inspection system is enabled. Optional. Read-only.
func (m *WindowsProtectionState) GetNetworkInspectionSystemEnabled()(*bool) {
    return m.networkInspectionSystemEnabled
}
// GetQuickScanOverdue gets the quickScanOverdue property value. A flag indicating weather a quick scan is overdue. Optional. Read-only.
func (m *WindowsProtectionState) GetQuickScanOverdue()(*bool) {
    return m.quickScanOverdue
}
// GetRealTimeProtectionEnabled gets the realTimeProtectionEnabled property value. A flag indicating whether real time protection is enabled. Optional. Read-only.
func (m *WindowsProtectionState) GetRealTimeProtectionEnabled()(*bool) {
    return m.realTimeProtectionEnabled
}
// GetRebootRequired gets the rebootRequired property value. A flag indicating whether a reboot is required. Optional. Read-only.
func (m *WindowsProtectionState) GetRebootRequired()(*bool) {
    return m.rebootRequired
}
// GetSignatureUpdateOverdue gets the signatureUpdateOverdue property value. A flag indicating whether an signature update is overdue. Optional. Read-only.
func (m *WindowsProtectionState) GetSignatureUpdateOverdue()(*bool) {
    return m.signatureUpdateOverdue
}
// GetSignatureVersion gets the signatureVersion property value. The signature version for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) GetSignatureVersion()(*string) {
    return m.signatureVersion
}
// GetTenantDisplayName gets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *WindowsProtectionState) GetTenantDisplayName()(*string) {
    return m.tenantDisplayName
}
// GetTenantId gets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *WindowsProtectionState) GetTenantId()(*string) {
    return m.tenantId
}
// Serialize serializes information the current object
func (m *WindowsProtectionState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("antiMalwareVersion", m.GetAntiMalwareVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("attentionRequired", m.GetAttentionRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deviceDeleted", m.GetDeviceDeleted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("devicePropertyRefreshDateTime", m.GetDevicePropertyRefreshDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("engineVersion", m.GetEngineVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("fullScanOverdue", m.GetFullScanOverdue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("fullScanRequired", m.GetFullScanRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastFullScanDateTime", m.GetLastFullScanDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastFullScanSignatureVersion", m.GetLastFullScanSignatureVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastQuickScanDateTime", m.GetLastQuickScanDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastQuickScanSignatureVersion", m.GetLastQuickScanSignatureVersion())
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
        err = writer.WriteTimeValue("lastReportedDateTime", m.GetLastReportedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("malwareProtectionEnabled", m.GetMalwareProtectionEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceHealthState", m.GetManagedDeviceHealthState())
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
        err = writer.WriteBoolValue("networkInspectionSystemEnabled", m.GetNetworkInspectionSystemEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("quickScanOverdue", m.GetQuickScanOverdue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("realTimeProtectionEnabled", m.GetRealTimeProtectionEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("rebootRequired", m.GetRebootRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("signatureUpdateOverdue", m.GetSignatureUpdateOverdue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("signatureVersion", m.GetSignatureVersion())
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
// SetAntiMalwareVersion sets the antiMalwareVersion property value. The anti-malware version for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) SetAntiMalwareVersion(value *string)() {
    m.antiMalwareVersion = value
}
// SetAttentionRequired sets the attentionRequired property value. A flag indicating whether attention is required for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) SetAttentionRequired(value *bool)() {
    m.attentionRequired = value
}
// SetDeviceDeleted sets the deviceDeleted property value. A flag indicating whether the managed device has been deleted. Optional. Read-only.
func (m *WindowsProtectionState) SetDeviceDeleted(value *bool)() {
    m.deviceDeleted = value
}
// SetDevicePropertyRefreshDateTime sets the devicePropertyRefreshDateTime property value. The date and time the device property has been refreshed. Optional. Read-only.
func (m *WindowsProtectionState) SetDevicePropertyRefreshDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.devicePropertyRefreshDateTime = value
}
// SetEngineVersion sets the engineVersion property value. The anti-virus engine version for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) SetEngineVersion(value *string)() {
    m.engineVersion = value
}
// SetFullScanOverdue sets the fullScanOverdue property value. A flag indicating whether quick scan is overdue for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) SetFullScanOverdue(value *bool)() {
    m.fullScanOverdue = value
}
// SetFullScanRequired sets the fullScanRequired property value. A flag indicating whether full scan is overdue for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) SetFullScanRequired(value *bool)() {
    m.fullScanRequired = value
}
// SetLastFullScanDateTime sets the lastFullScanDateTime property value. The date and time a full scan was completed. Optional. Read-only.
func (m *WindowsProtectionState) SetLastFullScanDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastFullScanDateTime = value
}
// SetLastFullScanSignatureVersion sets the lastFullScanSignatureVersion property value. The version anti-malware version used to perform the last full scan. Optional. Read-only.
func (m *WindowsProtectionState) SetLastFullScanSignatureVersion(value *string)() {
    m.lastFullScanSignatureVersion = value
}
// SetLastQuickScanDateTime sets the lastQuickScanDateTime property value. The date and time a quick scan was completed. Optional. Read-only.
func (m *WindowsProtectionState) SetLastQuickScanDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastQuickScanDateTime = value
}
// SetLastQuickScanSignatureVersion sets the lastQuickScanSignatureVersion property value. The version anti-malware version used to perform the last full scan. Optional. Read-only.
func (m *WindowsProtectionState) SetLastQuickScanSignatureVersion(value *string)() {
    m.lastQuickScanSignatureVersion = value
}
// SetLastRefreshedDateTime sets the lastRefreshedDateTime property value. Date and time the entity was last updated in the multi-tenant management platform. Optional. Read-only.
func (m *WindowsProtectionState) SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastRefreshedDateTime = value
}
// SetLastReportedDateTime sets the lastReportedDateTime property value. The date and time the protection state was last reported for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) SetLastReportedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastReportedDateTime = value
}
// SetMalwareProtectionEnabled sets the malwareProtectionEnabled property value. A flag indicating whether malware protection is enabled for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) SetMalwareProtectionEnabled(value *bool)() {
    m.malwareProtectionEnabled = value
}
// SetManagedDeviceHealthState sets the managedDeviceHealthState property value. The health state for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) SetManagedDeviceHealthState(value *string)() {
    m.managedDeviceHealthState = value
}
// SetManagedDeviceId sets the managedDeviceId property value. The unique identifier for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
// SetManagedDeviceName sets the managedDeviceName property value. The display name for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) SetManagedDeviceName(value *string)() {
    m.managedDeviceName = value
}
// SetNetworkInspectionSystemEnabled sets the networkInspectionSystemEnabled property value. A flag indicating whether the network inspection system is enabled. Optional. Read-only.
func (m *WindowsProtectionState) SetNetworkInspectionSystemEnabled(value *bool)() {
    m.networkInspectionSystemEnabled = value
}
// SetQuickScanOverdue sets the quickScanOverdue property value. A flag indicating weather a quick scan is overdue. Optional. Read-only.
func (m *WindowsProtectionState) SetQuickScanOverdue(value *bool)() {
    m.quickScanOverdue = value
}
// SetRealTimeProtectionEnabled sets the realTimeProtectionEnabled property value. A flag indicating whether real time protection is enabled. Optional. Read-only.
func (m *WindowsProtectionState) SetRealTimeProtectionEnabled(value *bool)() {
    m.realTimeProtectionEnabled = value
}
// SetRebootRequired sets the rebootRequired property value. A flag indicating whether a reboot is required. Optional. Read-only.
func (m *WindowsProtectionState) SetRebootRequired(value *bool)() {
    m.rebootRequired = value
}
// SetSignatureUpdateOverdue sets the signatureUpdateOverdue property value. A flag indicating whether an signature update is overdue. Optional. Read-only.
func (m *WindowsProtectionState) SetSignatureUpdateOverdue(value *bool)() {
    m.signatureUpdateOverdue = value
}
// SetSignatureVersion sets the signatureVersion property value. The signature version for the managed device. Optional. Read-only.
func (m *WindowsProtectionState) SetSignatureVersion(value *string)() {
    m.signatureVersion = value
}
// SetTenantDisplayName sets the tenantDisplayName property value. The display name for the managed tenant. Optional. Read-only.
func (m *WindowsProtectionState) SetTenantDisplayName(value *string)() {
    m.tenantDisplayName = value
}
// SetTenantId sets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *WindowsProtectionState) SetTenantId(value *string)() {
    m.tenantId = value
}
