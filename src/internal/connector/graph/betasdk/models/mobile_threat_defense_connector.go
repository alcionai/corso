package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileThreatDefenseConnector entity which represents a connection to Mobile threat defense partner.
type MobileThreatDefenseConnector struct {
    Entity
    // For IOS devices, allows the admin to configure whether the data sync partner may also collect metadata about installed applications from Intune
    allowPartnerToCollectIOSApplicationMetadata *bool
    // For IOS devices, allows the admin to configure whether the data sync partner may also collect metadata about personally installed applications from Intune
    allowPartnerToCollectIOSPersonalApplicationMetadata *bool
    // For Android, set whether Intune must receive data from the data sync partner prior to marking a device compliant
    androidDeviceBlockedOnMissingPartnerData *bool
    // For Android, set whether data from the data sync partner should be used during compliance evaluations
    androidEnabled *bool
    // For Android, set whether data from the data sync partner should be used during Mobile Application Management (MAM) evaluations. Only one partner per platform may be enabled for Mobile Application Management (MAM) evaluation.
    androidMobileApplicationManagementEnabled *bool
    // For IOS, set whether Intune must receive data from the data sync partner prior to marking a device compliant
    iosDeviceBlockedOnMissingPartnerData *bool
    // For IOS, get or set whether data from the data sync partner should be used during compliance evaluations
    iosEnabled *bool
    // For IOS, get or set whether data from the data sync partner should be used during Mobile Application Management (MAM) evaluations. Only one partner per platform may be enabled for Mobile Application Management (MAM) evaluation.
    iosMobileApplicationManagementEnabled *bool
    // DateTime of last Heartbeat recieved from the Data Sync Partner
    lastHeartbeatDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // For Mac, get or set whether Intune must receive data from the data sync partner prior to marking a device compliant
    macDeviceBlockedOnMissingPartnerData *bool
    // For Mac, get or set whether data from the data sync partner should be used during compliance evaluations
    macEnabled *bool
    // When TRUE, configuration profile management via Microsoft Defender for Endpoint is enabled. When FALSE, configuration profile management via Microsoft Defender for Endpoint is disabled.
    microsoftDefenderForEndpointAttachEnabled *bool
    // Partner state of this tenant.
    partnerState *MobileThreatPartnerTenantState
    // Get or Set days the per tenant tolerance to unresponsiveness for this partner integration
    partnerUnresponsivenessThresholdInDays *int32
    // Get or set whether to block devices on the enabled platforms that do not meet the minimum version requirements of the Data Sync Partner
    partnerUnsupportedOsVersionBlocked *bool
    // For Windows, set whether Intune must receive data from the data sync partner prior to marking a device compliant
    windowsDeviceBlockedOnMissingPartnerData *bool
    // For Windows, get or set whether data from the data sync partner should be used during compliance evaluations
    windowsEnabled *bool
    // When TRUE, app protection policies using the Device Threat Level rule will evaluate devices including data from this connector for Windows. When FALSE, Intune will not use device risk details sent over this connector during app protection policies calculation for policies with a Device Threat Level configured. Existing devices that are not compliant due to risk levels obtained from this connector will also become compliant.
    windowsMobileApplicationManagementEnabled *bool
}
// NewMobileThreatDefenseConnector instantiates a new mobileThreatDefenseConnector and sets the default values.
func NewMobileThreatDefenseConnector()(*MobileThreatDefenseConnector) {
    m := &MobileThreatDefenseConnector{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMobileThreatDefenseConnectorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileThreatDefenseConnectorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileThreatDefenseConnector(), nil
}
// GetAllowPartnerToCollectIOSApplicationMetadata gets the allowPartnerToCollectIOSApplicationMetadata property value. For IOS devices, allows the admin to configure whether the data sync partner may also collect metadata about installed applications from Intune
func (m *MobileThreatDefenseConnector) GetAllowPartnerToCollectIOSApplicationMetadata()(*bool) {
    return m.allowPartnerToCollectIOSApplicationMetadata
}
// GetAllowPartnerToCollectIOSPersonalApplicationMetadata gets the allowPartnerToCollectIOSPersonalApplicationMetadata property value. For IOS devices, allows the admin to configure whether the data sync partner may also collect metadata about personally installed applications from Intune
func (m *MobileThreatDefenseConnector) GetAllowPartnerToCollectIOSPersonalApplicationMetadata()(*bool) {
    return m.allowPartnerToCollectIOSPersonalApplicationMetadata
}
// GetAndroidDeviceBlockedOnMissingPartnerData gets the androidDeviceBlockedOnMissingPartnerData property value. For Android, set whether Intune must receive data from the data sync partner prior to marking a device compliant
func (m *MobileThreatDefenseConnector) GetAndroidDeviceBlockedOnMissingPartnerData()(*bool) {
    return m.androidDeviceBlockedOnMissingPartnerData
}
// GetAndroidEnabled gets the androidEnabled property value. For Android, set whether data from the data sync partner should be used during compliance evaluations
func (m *MobileThreatDefenseConnector) GetAndroidEnabled()(*bool) {
    return m.androidEnabled
}
// GetAndroidMobileApplicationManagementEnabled gets the androidMobileApplicationManagementEnabled property value. For Android, set whether data from the data sync partner should be used during Mobile Application Management (MAM) evaluations. Only one partner per platform may be enabled for Mobile Application Management (MAM) evaluation.
func (m *MobileThreatDefenseConnector) GetAndroidMobileApplicationManagementEnabled()(*bool) {
    return m.androidMobileApplicationManagementEnabled
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileThreatDefenseConnector) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowPartnerToCollectIOSApplicationMetadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowPartnerToCollectIOSApplicationMetadata(val)
        }
        return nil
    }
    res["allowPartnerToCollectIOSPersonalApplicationMetadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowPartnerToCollectIOSPersonalApplicationMetadata(val)
        }
        return nil
    }
    res["androidDeviceBlockedOnMissingPartnerData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidDeviceBlockedOnMissingPartnerData(val)
        }
        return nil
    }
    res["androidEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidEnabled(val)
        }
        return nil
    }
    res["androidMobileApplicationManagementEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidMobileApplicationManagementEnabled(val)
        }
        return nil
    }
    res["iosDeviceBlockedOnMissingPartnerData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIosDeviceBlockedOnMissingPartnerData(val)
        }
        return nil
    }
    res["iosEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIosEnabled(val)
        }
        return nil
    }
    res["iosMobileApplicationManagementEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIosMobileApplicationManagementEnabled(val)
        }
        return nil
    }
    res["lastHeartbeatDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastHeartbeatDateTime(val)
        }
        return nil
    }
    res["macDeviceBlockedOnMissingPartnerData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacDeviceBlockedOnMissingPartnerData(val)
        }
        return nil
    }
    res["macEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacEnabled(val)
        }
        return nil
    }
    res["microsoftDefenderForEndpointAttachEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftDefenderForEndpointAttachEnabled(val)
        }
        return nil
    }
    res["partnerState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMobileThreatPartnerTenantState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerState(val.(*MobileThreatPartnerTenantState))
        }
        return nil
    }
    res["partnerUnresponsivenessThresholdInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerUnresponsivenessThresholdInDays(val)
        }
        return nil
    }
    res["partnerUnsupportedOsVersionBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerUnsupportedOsVersionBlocked(val)
        }
        return nil
    }
    res["windowsDeviceBlockedOnMissingPartnerData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsDeviceBlockedOnMissingPartnerData(val)
        }
        return nil
    }
    res["windowsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsEnabled(val)
        }
        return nil
    }
    res["windowsMobileApplicationManagementEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsMobileApplicationManagementEnabled(val)
        }
        return nil
    }
    return res
}
// GetIosDeviceBlockedOnMissingPartnerData gets the iosDeviceBlockedOnMissingPartnerData property value. For IOS, set whether Intune must receive data from the data sync partner prior to marking a device compliant
func (m *MobileThreatDefenseConnector) GetIosDeviceBlockedOnMissingPartnerData()(*bool) {
    return m.iosDeviceBlockedOnMissingPartnerData
}
// GetIosEnabled gets the iosEnabled property value. For IOS, get or set whether data from the data sync partner should be used during compliance evaluations
func (m *MobileThreatDefenseConnector) GetIosEnabled()(*bool) {
    return m.iosEnabled
}
// GetIosMobileApplicationManagementEnabled gets the iosMobileApplicationManagementEnabled property value. For IOS, get or set whether data from the data sync partner should be used during Mobile Application Management (MAM) evaluations. Only one partner per platform may be enabled for Mobile Application Management (MAM) evaluation.
func (m *MobileThreatDefenseConnector) GetIosMobileApplicationManagementEnabled()(*bool) {
    return m.iosMobileApplicationManagementEnabled
}
// GetLastHeartbeatDateTime gets the lastHeartbeatDateTime property value. DateTime of last Heartbeat recieved from the Data Sync Partner
func (m *MobileThreatDefenseConnector) GetLastHeartbeatDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastHeartbeatDateTime
}
// GetMacDeviceBlockedOnMissingPartnerData gets the macDeviceBlockedOnMissingPartnerData property value. For Mac, get or set whether Intune must receive data from the data sync partner prior to marking a device compliant
func (m *MobileThreatDefenseConnector) GetMacDeviceBlockedOnMissingPartnerData()(*bool) {
    return m.macDeviceBlockedOnMissingPartnerData
}
// GetMacEnabled gets the macEnabled property value. For Mac, get or set whether data from the data sync partner should be used during compliance evaluations
func (m *MobileThreatDefenseConnector) GetMacEnabled()(*bool) {
    return m.macEnabled
}
// GetMicrosoftDefenderForEndpointAttachEnabled gets the microsoftDefenderForEndpointAttachEnabled property value. When TRUE, configuration profile management via Microsoft Defender for Endpoint is enabled. When FALSE, configuration profile management via Microsoft Defender for Endpoint is disabled.
func (m *MobileThreatDefenseConnector) GetMicrosoftDefenderForEndpointAttachEnabled()(*bool) {
    return m.microsoftDefenderForEndpointAttachEnabled
}
// GetPartnerState gets the partnerState property value. Partner state of this tenant.
func (m *MobileThreatDefenseConnector) GetPartnerState()(*MobileThreatPartnerTenantState) {
    return m.partnerState
}
// GetPartnerUnresponsivenessThresholdInDays gets the partnerUnresponsivenessThresholdInDays property value. Get or Set days the per tenant tolerance to unresponsiveness for this partner integration
func (m *MobileThreatDefenseConnector) GetPartnerUnresponsivenessThresholdInDays()(*int32) {
    return m.partnerUnresponsivenessThresholdInDays
}
// GetPartnerUnsupportedOsVersionBlocked gets the partnerUnsupportedOsVersionBlocked property value. Get or set whether to block devices on the enabled platforms that do not meet the minimum version requirements of the Data Sync Partner
func (m *MobileThreatDefenseConnector) GetPartnerUnsupportedOsVersionBlocked()(*bool) {
    return m.partnerUnsupportedOsVersionBlocked
}
// GetWindowsDeviceBlockedOnMissingPartnerData gets the windowsDeviceBlockedOnMissingPartnerData property value. For Windows, set whether Intune must receive data from the data sync partner prior to marking a device compliant
func (m *MobileThreatDefenseConnector) GetWindowsDeviceBlockedOnMissingPartnerData()(*bool) {
    return m.windowsDeviceBlockedOnMissingPartnerData
}
// GetWindowsEnabled gets the windowsEnabled property value. For Windows, get or set whether data from the data sync partner should be used during compliance evaluations
func (m *MobileThreatDefenseConnector) GetWindowsEnabled()(*bool) {
    return m.windowsEnabled
}
// GetWindowsMobileApplicationManagementEnabled gets the windowsMobileApplicationManagementEnabled property value. When TRUE, app protection policies using the Device Threat Level rule will evaluate devices including data from this connector for Windows. When FALSE, Intune will not use device risk details sent over this connector during app protection policies calculation for policies with a Device Threat Level configured. Existing devices that are not compliant due to risk levels obtained from this connector will also become compliant.
func (m *MobileThreatDefenseConnector) GetWindowsMobileApplicationManagementEnabled()(*bool) {
    return m.windowsMobileApplicationManagementEnabled
}
// Serialize serializes information the current object
func (m *MobileThreatDefenseConnector) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowPartnerToCollectIOSApplicationMetadata", m.GetAllowPartnerToCollectIOSApplicationMetadata())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowPartnerToCollectIOSPersonalApplicationMetadata", m.GetAllowPartnerToCollectIOSPersonalApplicationMetadata())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("androidDeviceBlockedOnMissingPartnerData", m.GetAndroidDeviceBlockedOnMissingPartnerData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("androidEnabled", m.GetAndroidEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("androidMobileApplicationManagementEnabled", m.GetAndroidMobileApplicationManagementEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iosDeviceBlockedOnMissingPartnerData", m.GetIosDeviceBlockedOnMissingPartnerData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iosEnabled", m.GetIosEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iosMobileApplicationManagementEnabled", m.GetIosMobileApplicationManagementEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastHeartbeatDateTime", m.GetLastHeartbeatDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("macDeviceBlockedOnMissingPartnerData", m.GetMacDeviceBlockedOnMissingPartnerData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("macEnabled", m.GetMacEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("microsoftDefenderForEndpointAttachEnabled", m.GetMicrosoftDefenderForEndpointAttachEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetPartnerState() != nil {
        cast := (*m.GetPartnerState()).String()
        err = writer.WriteStringValue("partnerState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("partnerUnresponsivenessThresholdInDays", m.GetPartnerUnresponsivenessThresholdInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("partnerUnsupportedOsVersionBlocked", m.GetPartnerUnsupportedOsVersionBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("windowsDeviceBlockedOnMissingPartnerData", m.GetWindowsDeviceBlockedOnMissingPartnerData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("windowsEnabled", m.GetWindowsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("windowsMobileApplicationManagementEnabled", m.GetWindowsMobileApplicationManagementEnabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowPartnerToCollectIOSApplicationMetadata sets the allowPartnerToCollectIOSApplicationMetadata property value. For IOS devices, allows the admin to configure whether the data sync partner may also collect metadata about installed applications from Intune
func (m *MobileThreatDefenseConnector) SetAllowPartnerToCollectIOSApplicationMetadata(value *bool)() {
    m.allowPartnerToCollectIOSApplicationMetadata = value
}
// SetAllowPartnerToCollectIOSPersonalApplicationMetadata sets the allowPartnerToCollectIOSPersonalApplicationMetadata property value. For IOS devices, allows the admin to configure whether the data sync partner may also collect metadata about personally installed applications from Intune
func (m *MobileThreatDefenseConnector) SetAllowPartnerToCollectIOSPersonalApplicationMetadata(value *bool)() {
    m.allowPartnerToCollectIOSPersonalApplicationMetadata = value
}
// SetAndroidDeviceBlockedOnMissingPartnerData sets the androidDeviceBlockedOnMissingPartnerData property value. For Android, set whether Intune must receive data from the data sync partner prior to marking a device compliant
func (m *MobileThreatDefenseConnector) SetAndroidDeviceBlockedOnMissingPartnerData(value *bool)() {
    m.androidDeviceBlockedOnMissingPartnerData = value
}
// SetAndroidEnabled sets the androidEnabled property value. For Android, set whether data from the data sync partner should be used during compliance evaluations
func (m *MobileThreatDefenseConnector) SetAndroidEnabled(value *bool)() {
    m.androidEnabled = value
}
// SetAndroidMobileApplicationManagementEnabled sets the androidMobileApplicationManagementEnabled property value. For Android, set whether data from the data sync partner should be used during Mobile Application Management (MAM) evaluations. Only one partner per platform may be enabled for Mobile Application Management (MAM) evaluation.
func (m *MobileThreatDefenseConnector) SetAndroidMobileApplicationManagementEnabled(value *bool)() {
    m.androidMobileApplicationManagementEnabled = value
}
// SetIosDeviceBlockedOnMissingPartnerData sets the iosDeviceBlockedOnMissingPartnerData property value. For IOS, set whether Intune must receive data from the data sync partner prior to marking a device compliant
func (m *MobileThreatDefenseConnector) SetIosDeviceBlockedOnMissingPartnerData(value *bool)() {
    m.iosDeviceBlockedOnMissingPartnerData = value
}
// SetIosEnabled sets the iosEnabled property value. For IOS, get or set whether data from the data sync partner should be used during compliance evaluations
func (m *MobileThreatDefenseConnector) SetIosEnabled(value *bool)() {
    m.iosEnabled = value
}
// SetIosMobileApplicationManagementEnabled sets the iosMobileApplicationManagementEnabled property value. For IOS, get or set whether data from the data sync partner should be used during Mobile Application Management (MAM) evaluations. Only one partner per platform may be enabled for Mobile Application Management (MAM) evaluation.
func (m *MobileThreatDefenseConnector) SetIosMobileApplicationManagementEnabled(value *bool)() {
    m.iosMobileApplicationManagementEnabled = value
}
// SetLastHeartbeatDateTime sets the lastHeartbeatDateTime property value. DateTime of last Heartbeat recieved from the Data Sync Partner
func (m *MobileThreatDefenseConnector) SetLastHeartbeatDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastHeartbeatDateTime = value
}
// SetMacDeviceBlockedOnMissingPartnerData sets the macDeviceBlockedOnMissingPartnerData property value. For Mac, get or set whether Intune must receive data from the data sync partner prior to marking a device compliant
func (m *MobileThreatDefenseConnector) SetMacDeviceBlockedOnMissingPartnerData(value *bool)() {
    m.macDeviceBlockedOnMissingPartnerData = value
}
// SetMacEnabled sets the macEnabled property value. For Mac, get or set whether data from the data sync partner should be used during compliance evaluations
func (m *MobileThreatDefenseConnector) SetMacEnabled(value *bool)() {
    m.macEnabled = value
}
// SetMicrosoftDefenderForEndpointAttachEnabled sets the microsoftDefenderForEndpointAttachEnabled property value. When TRUE, configuration profile management via Microsoft Defender for Endpoint is enabled. When FALSE, configuration profile management via Microsoft Defender for Endpoint is disabled.
func (m *MobileThreatDefenseConnector) SetMicrosoftDefenderForEndpointAttachEnabled(value *bool)() {
    m.microsoftDefenderForEndpointAttachEnabled = value
}
// SetPartnerState sets the partnerState property value. Partner state of this tenant.
func (m *MobileThreatDefenseConnector) SetPartnerState(value *MobileThreatPartnerTenantState)() {
    m.partnerState = value
}
// SetPartnerUnresponsivenessThresholdInDays sets the partnerUnresponsivenessThresholdInDays property value. Get or Set days the per tenant tolerance to unresponsiveness for this partner integration
func (m *MobileThreatDefenseConnector) SetPartnerUnresponsivenessThresholdInDays(value *int32)() {
    m.partnerUnresponsivenessThresholdInDays = value
}
// SetPartnerUnsupportedOsVersionBlocked sets the partnerUnsupportedOsVersionBlocked property value. Get or set whether to block devices on the enabled platforms that do not meet the minimum version requirements of the Data Sync Partner
func (m *MobileThreatDefenseConnector) SetPartnerUnsupportedOsVersionBlocked(value *bool)() {
    m.partnerUnsupportedOsVersionBlocked = value
}
// SetWindowsDeviceBlockedOnMissingPartnerData sets the windowsDeviceBlockedOnMissingPartnerData property value. For Windows, set whether Intune must receive data from the data sync partner prior to marking a device compliant
func (m *MobileThreatDefenseConnector) SetWindowsDeviceBlockedOnMissingPartnerData(value *bool)() {
    m.windowsDeviceBlockedOnMissingPartnerData = value
}
// SetWindowsEnabled sets the windowsEnabled property value. For Windows, get or set whether data from the data sync partner should be used during compliance evaluations
func (m *MobileThreatDefenseConnector) SetWindowsEnabled(value *bool)() {
    m.windowsEnabled = value
}
// SetWindowsMobileApplicationManagementEnabled sets the windowsMobileApplicationManagementEnabled property value. When TRUE, app protection policies using the Device Threat Level rule will evaluate devices including data from this connector for Windows. When FALSE, Intune will not use device risk details sent over this connector during app protection policies calculation for policies with a Device Threat Level configured. Existing devices that are not compliant due to risk levels obtained from this connector will also become compliant.
func (m *MobileThreatDefenseConnector) SetWindowsMobileApplicationManagementEnabled(value *bool)() {
    m.windowsMobileApplicationManagementEnabled = value
}
