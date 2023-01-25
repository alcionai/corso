package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsProtectionState 
type WindowsProtectionState struct {
    Entity
    // Current anti malware version
    antiMalwareVersion *string
    // Device malware list
    detectedMalwareState []WindowsDeviceMalwareStateable
    // Computer's state (like clean or pending full scan or pending reboot etc). Possible values are: clean, fullScanPending, rebootPending, manualStepsPending, offlineScanPending, critical.
    deviceState *WindowsDeviceHealthState
    // Current endpoint protection engine's version
    engineVersion *string
    // Full scan overdue or not?
    fullScanOverdue *bool
    // Full scan required or not?
    fullScanRequired *bool
    // Indicates whether the device is a virtual machine.
    isVirtualMachine *bool
    // Last quick scan datetime
    lastFullScanDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Last full scan signature version
    lastFullScanSignatureVersion *string
    // Last quick scan datetime
    lastQuickScanDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Last quick scan signature version
    lastQuickScanSignatureVersion *string
    // Last device health status reported time
    lastReportedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Anti malware is enabled or not
    malwareProtectionEnabled *bool
    // Network inspection system enabled or not?
    networkInspectionSystemEnabled *bool
    // Product Status of Windows Defender Antivirus. Possible values are: noStatus, serviceNotRunning, serviceStartedWithoutMalwareProtection, pendingFullScanDueToThreatAction, pendingRebootDueToThreatAction, pendingManualStepsDueToThreatAction, avSignaturesOutOfDate, asSignaturesOutOfDate, noQuickScanHappenedForSpecifiedPeriod, noFullScanHappenedForSpecifiedPeriod, systemInitiatedScanInProgress, systemInitiatedCleanInProgress, samplesPendingSubmission, productRunningInEvaluationMode, productRunningInNonGenuineMode, productExpired, offlineScanRequired, serviceShutdownAsPartOfSystemShutdown, threatRemediationFailedCritically, threatRemediationFailedNonCritically, noStatusFlagsSet, platformOutOfDate, platformUpdateInProgress, platformAboutToBeOutdated, signatureOrPlatformEndOfLifeIsPastOrIsImpending, windowsSModeSignaturesInUseOnNonWin10SInstall.
    productStatus *WindowsDefenderProductStatus
    // Quick scan overdue or not?
    quickScanOverdue *bool
    // Real time protection is enabled or not?
    realTimeProtectionEnabled *bool
    // Reboot required or not?
    rebootRequired *bool
    // Signature out of date or not?
    signatureUpdateOverdue *bool
    // Current malware definitions version
    signatureVersion *string
    // Indicates whether the Windows Defender tamper protection feature is enabled.
    tamperProtectionEnabled *bool
}
// NewWindowsProtectionState instantiates a new windowsProtectionState and sets the default values.
func NewWindowsProtectionState()(*WindowsProtectionState) {
    m := &WindowsProtectionState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsProtectionStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsProtectionStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsProtectionState(), nil
}
// GetAntiMalwareVersion gets the antiMalwareVersion property value. Current anti malware version
func (m *WindowsProtectionState) GetAntiMalwareVersion()(*string) {
    return m.antiMalwareVersion
}
// GetDetectedMalwareState gets the detectedMalwareState property value. Device malware list
func (m *WindowsProtectionState) GetDetectedMalwareState()([]WindowsDeviceMalwareStateable) {
    return m.detectedMalwareState
}
// GetDeviceState gets the deviceState property value. Computer's state (like clean or pending full scan or pending reboot etc). Possible values are: clean, fullScanPending, rebootPending, manualStepsPending, offlineScanPending, critical.
func (m *WindowsProtectionState) GetDeviceState()(*WindowsDeviceHealthState) {
    return m.deviceState
}
// GetEngineVersion gets the engineVersion property value. Current endpoint protection engine's version
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
    res["detectedMalwareState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsDeviceMalwareStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsDeviceMalwareStateable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsDeviceMalwareStateable)
            }
            m.SetDetectedMalwareState(res)
        }
        return nil
    }
    res["deviceState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsDeviceHealthState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceState(val.(*WindowsDeviceHealthState))
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
    res["isVirtualMachine"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsVirtualMachine(val)
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
    res["productStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsDefenderProductStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductStatus(val.(*WindowsDefenderProductStatus))
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
    res["tamperProtectionEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTamperProtectionEnabled(val)
        }
        return nil
    }
    return res
}
// GetFullScanOverdue gets the fullScanOverdue property value. Full scan overdue or not?
func (m *WindowsProtectionState) GetFullScanOverdue()(*bool) {
    return m.fullScanOverdue
}
// GetFullScanRequired gets the fullScanRequired property value. Full scan required or not?
func (m *WindowsProtectionState) GetFullScanRequired()(*bool) {
    return m.fullScanRequired
}
// GetIsVirtualMachine gets the isVirtualMachine property value. Indicates whether the device is a virtual machine.
func (m *WindowsProtectionState) GetIsVirtualMachine()(*bool) {
    return m.isVirtualMachine
}
// GetLastFullScanDateTime gets the lastFullScanDateTime property value. Last quick scan datetime
func (m *WindowsProtectionState) GetLastFullScanDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastFullScanDateTime
}
// GetLastFullScanSignatureVersion gets the lastFullScanSignatureVersion property value. Last full scan signature version
func (m *WindowsProtectionState) GetLastFullScanSignatureVersion()(*string) {
    return m.lastFullScanSignatureVersion
}
// GetLastQuickScanDateTime gets the lastQuickScanDateTime property value. Last quick scan datetime
func (m *WindowsProtectionState) GetLastQuickScanDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastQuickScanDateTime
}
// GetLastQuickScanSignatureVersion gets the lastQuickScanSignatureVersion property value. Last quick scan signature version
func (m *WindowsProtectionState) GetLastQuickScanSignatureVersion()(*string) {
    return m.lastQuickScanSignatureVersion
}
// GetLastReportedDateTime gets the lastReportedDateTime property value. Last device health status reported time
func (m *WindowsProtectionState) GetLastReportedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastReportedDateTime
}
// GetMalwareProtectionEnabled gets the malwareProtectionEnabled property value. Anti malware is enabled or not
func (m *WindowsProtectionState) GetMalwareProtectionEnabled()(*bool) {
    return m.malwareProtectionEnabled
}
// GetNetworkInspectionSystemEnabled gets the networkInspectionSystemEnabled property value. Network inspection system enabled or not?
func (m *WindowsProtectionState) GetNetworkInspectionSystemEnabled()(*bool) {
    return m.networkInspectionSystemEnabled
}
// GetProductStatus gets the productStatus property value. Product Status of Windows Defender Antivirus. Possible values are: noStatus, serviceNotRunning, serviceStartedWithoutMalwareProtection, pendingFullScanDueToThreatAction, pendingRebootDueToThreatAction, pendingManualStepsDueToThreatAction, avSignaturesOutOfDate, asSignaturesOutOfDate, noQuickScanHappenedForSpecifiedPeriod, noFullScanHappenedForSpecifiedPeriod, systemInitiatedScanInProgress, systemInitiatedCleanInProgress, samplesPendingSubmission, productRunningInEvaluationMode, productRunningInNonGenuineMode, productExpired, offlineScanRequired, serviceShutdownAsPartOfSystemShutdown, threatRemediationFailedCritically, threatRemediationFailedNonCritically, noStatusFlagsSet, platformOutOfDate, platformUpdateInProgress, platformAboutToBeOutdated, signatureOrPlatformEndOfLifeIsPastOrIsImpending, windowsSModeSignaturesInUseOnNonWin10SInstall.
func (m *WindowsProtectionState) GetProductStatus()(*WindowsDefenderProductStatus) {
    return m.productStatus
}
// GetQuickScanOverdue gets the quickScanOverdue property value. Quick scan overdue or not?
func (m *WindowsProtectionState) GetQuickScanOverdue()(*bool) {
    return m.quickScanOverdue
}
// GetRealTimeProtectionEnabled gets the realTimeProtectionEnabled property value. Real time protection is enabled or not?
func (m *WindowsProtectionState) GetRealTimeProtectionEnabled()(*bool) {
    return m.realTimeProtectionEnabled
}
// GetRebootRequired gets the rebootRequired property value. Reboot required or not?
func (m *WindowsProtectionState) GetRebootRequired()(*bool) {
    return m.rebootRequired
}
// GetSignatureUpdateOverdue gets the signatureUpdateOverdue property value. Signature out of date or not?
func (m *WindowsProtectionState) GetSignatureUpdateOverdue()(*bool) {
    return m.signatureUpdateOverdue
}
// GetSignatureVersion gets the signatureVersion property value. Current malware definitions version
func (m *WindowsProtectionState) GetSignatureVersion()(*string) {
    return m.signatureVersion
}
// GetTamperProtectionEnabled gets the tamperProtectionEnabled property value. Indicates whether the Windows Defender tamper protection feature is enabled.
func (m *WindowsProtectionState) GetTamperProtectionEnabled()(*bool) {
    return m.tamperProtectionEnabled
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
    if m.GetDetectedMalwareState() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDetectedMalwareState()))
        for i, v := range m.GetDetectedMalwareState() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("detectedMalwareState", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceState() != nil {
        cast := (*m.GetDeviceState()).String()
        err = writer.WriteStringValue("deviceState", &cast)
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
        err = writer.WriteBoolValue("isVirtualMachine", m.GetIsVirtualMachine())
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
        err = writer.WriteBoolValue("networkInspectionSystemEnabled", m.GetNetworkInspectionSystemEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetProductStatus() != nil {
        cast := (*m.GetProductStatus()).String()
        err = writer.WriteStringValue("productStatus", &cast)
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
        err = writer.WriteBoolValue("tamperProtectionEnabled", m.GetTamperProtectionEnabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAntiMalwareVersion sets the antiMalwareVersion property value. Current anti malware version
func (m *WindowsProtectionState) SetAntiMalwareVersion(value *string)() {
    m.antiMalwareVersion = value
}
// SetDetectedMalwareState sets the detectedMalwareState property value. Device malware list
func (m *WindowsProtectionState) SetDetectedMalwareState(value []WindowsDeviceMalwareStateable)() {
    m.detectedMalwareState = value
}
// SetDeviceState sets the deviceState property value. Computer's state (like clean or pending full scan or pending reboot etc). Possible values are: clean, fullScanPending, rebootPending, manualStepsPending, offlineScanPending, critical.
func (m *WindowsProtectionState) SetDeviceState(value *WindowsDeviceHealthState)() {
    m.deviceState = value
}
// SetEngineVersion sets the engineVersion property value. Current endpoint protection engine's version
func (m *WindowsProtectionState) SetEngineVersion(value *string)() {
    m.engineVersion = value
}
// SetFullScanOverdue sets the fullScanOverdue property value. Full scan overdue or not?
func (m *WindowsProtectionState) SetFullScanOverdue(value *bool)() {
    m.fullScanOverdue = value
}
// SetFullScanRequired sets the fullScanRequired property value. Full scan required or not?
func (m *WindowsProtectionState) SetFullScanRequired(value *bool)() {
    m.fullScanRequired = value
}
// SetIsVirtualMachine sets the isVirtualMachine property value. Indicates whether the device is a virtual machine.
func (m *WindowsProtectionState) SetIsVirtualMachine(value *bool)() {
    m.isVirtualMachine = value
}
// SetLastFullScanDateTime sets the lastFullScanDateTime property value. Last quick scan datetime
func (m *WindowsProtectionState) SetLastFullScanDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastFullScanDateTime = value
}
// SetLastFullScanSignatureVersion sets the lastFullScanSignatureVersion property value. Last full scan signature version
func (m *WindowsProtectionState) SetLastFullScanSignatureVersion(value *string)() {
    m.lastFullScanSignatureVersion = value
}
// SetLastQuickScanDateTime sets the lastQuickScanDateTime property value. Last quick scan datetime
func (m *WindowsProtectionState) SetLastQuickScanDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastQuickScanDateTime = value
}
// SetLastQuickScanSignatureVersion sets the lastQuickScanSignatureVersion property value. Last quick scan signature version
func (m *WindowsProtectionState) SetLastQuickScanSignatureVersion(value *string)() {
    m.lastQuickScanSignatureVersion = value
}
// SetLastReportedDateTime sets the lastReportedDateTime property value. Last device health status reported time
func (m *WindowsProtectionState) SetLastReportedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastReportedDateTime = value
}
// SetMalwareProtectionEnabled sets the malwareProtectionEnabled property value. Anti malware is enabled or not
func (m *WindowsProtectionState) SetMalwareProtectionEnabled(value *bool)() {
    m.malwareProtectionEnabled = value
}
// SetNetworkInspectionSystemEnabled sets the networkInspectionSystemEnabled property value. Network inspection system enabled or not?
func (m *WindowsProtectionState) SetNetworkInspectionSystemEnabled(value *bool)() {
    m.networkInspectionSystemEnabled = value
}
// SetProductStatus sets the productStatus property value. Product Status of Windows Defender Antivirus. Possible values are: noStatus, serviceNotRunning, serviceStartedWithoutMalwareProtection, pendingFullScanDueToThreatAction, pendingRebootDueToThreatAction, pendingManualStepsDueToThreatAction, avSignaturesOutOfDate, asSignaturesOutOfDate, noQuickScanHappenedForSpecifiedPeriod, noFullScanHappenedForSpecifiedPeriod, systemInitiatedScanInProgress, systemInitiatedCleanInProgress, samplesPendingSubmission, productRunningInEvaluationMode, productRunningInNonGenuineMode, productExpired, offlineScanRequired, serviceShutdownAsPartOfSystemShutdown, threatRemediationFailedCritically, threatRemediationFailedNonCritically, noStatusFlagsSet, platformOutOfDate, platformUpdateInProgress, platformAboutToBeOutdated, signatureOrPlatformEndOfLifeIsPastOrIsImpending, windowsSModeSignaturesInUseOnNonWin10SInstall.
func (m *WindowsProtectionState) SetProductStatus(value *WindowsDefenderProductStatus)() {
    m.productStatus = value
}
// SetQuickScanOverdue sets the quickScanOverdue property value. Quick scan overdue or not?
func (m *WindowsProtectionState) SetQuickScanOverdue(value *bool)() {
    m.quickScanOverdue = value
}
// SetRealTimeProtectionEnabled sets the realTimeProtectionEnabled property value. Real time protection is enabled or not?
func (m *WindowsProtectionState) SetRealTimeProtectionEnabled(value *bool)() {
    m.realTimeProtectionEnabled = value
}
// SetRebootRequired sets the rebootRequired property value. Reboot required or not?
func (m *WindowsProtectionState) SetRebootRequired(value *bool)() {
    m.rebootRequired = value
}
// SetSignatureUpdateOverdue sets the signatureUpdateOverdue property value. Signature out of date or not?
func (m *WindowsProtectionState) SetSignatureUpdateOverdue(value *bool)() {
    m.signatureUpdateOverdue = value
}
// SetSignatureVersion sets the signatureVersion property value. Current malware definitions version
func (m *WindowsProtectionState) SetSignatureVersion(value *string)() {
    m.signatureVersion = value
}
// SetTamperProtectionEnabled sets the tamperProtectionEnabled property value. Indicates whether the Windows Defender tamper protection feature is enabled.
func (m *WindowsProtectionState) SetTamperProtectionEnabled(value *bool)() {
    m.tamperProtectionEnabled = value
}
