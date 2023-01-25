package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceProtectionOverview hardware information of a given device.
type DeviceProtectionOverview struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Clean device count.
    cleanDeviceCount *int32
    // Critical failures device count.
    criticalFailuresDeviceCount *int32
    // Device with inactive threat agent count
    inactiveThreatAgentDeviceCount *int32
    // The OdataType property
    odataType *string
    // Pending full scan device count.
    pendingFullScanDeviceCount *int32
    // Pending manual steps device count.
    pendingManualStepsDeviceCount *int32
    // Pending offline scan device count.
    pendingOfflineScanDeviceCount *int32
    // Pending quick scan device count. Valid values -2147483648 to 2147483647
    pendingQuickScanDeviceCount *int32
    // Pending restart device count.
    pendingRestartDeviceCount *int32
    // Device with old signature count.
    pendingSignatureUpdateDeviceCount *int32
    // Total device count.
    totalReportedDeviceCount *int32
    // Device with threat agent state as unknown count.
    unknownStateThreatAgentDeviceCount *int32
}
// NewDeviceProtectionOverview instantiates a new deviceProtectionOverview and sets the default values.
func NewDeviceProtectionOverview()(*DeviceProtectionOverview) {
    m := &DeviceProtectionOverview{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceProtectionOverviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceProtectionOverviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceProtectionOverview(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceProtectionOverview) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCleanDeviceCount gets the cleanDeviceCount property value. Clean device count.
func (m *DeviceProtectionOverview) GetCleanDeviceCount()(*int32) {
    return m.cleanDeviceCount
}
// GetCriticalFailuresDeviceCount gets the criticalFailuresDeviceCount property value. Critical failures device count.
func (m *DeviceProtectionOverview) GetCriticalFailuresDeviceCount()(*int32) {
    return m.criticalFailuresDeviceCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceProtectionOverview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cleanDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCleanDeviceCount(val)
        }
        return nil
    }
    res["criticalFailuresDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCriticalFailuresDeviceCount(val)
        }
        return nil
    }
    res["inactiveThreatAgentDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInactiveThreatAgentDeviceCount(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["pendingFullScanDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingFullScanDeviceCount(val)
        }
        return nil
    }
    res["pendingManualStepsDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingManualStepsDeviceCount(val)
        }
        return nil
    }
    res["pendingOfflineScanDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingOfflineScanDeviceCount(val)
        }
        return nil
    }
    res["pendingQuickScanDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingQuickScanDeviceCount(val)
        }
        return nil
    }
    res["pendingRestartDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingRestartDeviceCount(val)
        }
        return nil
    }
    res["pendingSignatureUpdateDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingSignatureUpdateDeviceCount(val)
        }
        return nil
    }
    res["totalReportedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalReportedDeviceCount(val)
        }
        return nil
    }
    res["unknownStateThreatAgentDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnknownStateThreatAgentDeviceCount(val)
        }
        return nil
    }
    return res
}
// GetInactiveThreatAgentDeviceCount gets the inactiveThreatAgentDeviceCount property value. Device with inactive threat agent count
func (m *DeviceProtectionOverview) GetInactiveThreatAgentDeviceCount()(*int32) {
    return m.inactiveThreatAgentDeviceCount
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceProtectionOverview) GetOdataType()(*string) {
    return m.odataType
}
// GetPendingFullScanDeviceCount gets the pendingFullScanDeviceCount property value. Pending full scan device count.
func (m *DeviceProtectionOverview) GetPendingFullScanDeviceCount()(*int32) {
    return m.pendingFullScanDeviceCount
}
// GetPendingManualStepsDeviceCount gets the pendingManualStepsDeviceCount property value. Pending manual steps device count.
func (m *DeviceProtectionOverview) GetPendingManualStepsDeviceCount()(*int32) {
    return m.pendingManualStepsDeviceCount
}
// GetPendingOfflineScanDeviceCount gets the pendingOfflineScanDeviceCount property value. Pending offline scan device count.
func (m *DeviceProtectionOverview) GetPendingOfflineScanDeviceCount()(*int32) {
    return m.pendingOfflineScanDeviceCount
}
// GetPendingQuickScanDeviceCount gets the pendingQuickScanDeviceCount property value. Pending quick scan device count. Valid values -2147483648 to 2147483647
func (m *DeviceProtectionOverview) GetPendingQuickScanDeviceCount()(*int32) {
    return m.pendingQuickScanDeviceCount
}
// GetPendingRestartDeviceCount gets the pendingRestartDeviceCount property value. Pending restart device count.
func (m *DeviceProtectionOverview) GetPendingRestartDeviceCount()(*int32) {
    return m.pendingRestartDeviceCount
}
// GetPendingSignatureUpdateDeviceCount gets the pendingSignatureUpdateDeviceCount property value. Device with old signature count.
func (m *DeviceProtectionOverview) GetPendingSignatureUpdateDeviceCount()(*int32) {
    return m.pendingSignatureUpdateDeviceCount
}
// GetTotalReportedDeviceCount gets the totalReportedDeviceCount property value. Total device count.
func (m *DeviceProtectionOverview) GetTotalReportedDeviceCount()(*int32) {
    return m.totalReportedDeviceCount
}
// GetUnknownStateThreatAgentDeviceCount gets the unknownStateThreatAgentDeviceCount property value. Device with threat agent state as unknown count.
func (m *DeviceProtectionOverview) GetUnknownStateThreatAgentDeviceCount()(*int32) {
    return m.unknownStateThreatAgentDeviceCount
}
// Serialize serializes information the current object
func (m *DeviceProtectionOverview) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("cleanDeviceCount", m.GetCleanDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("criticalFailuresDeviceCount", m.GetCriticalFailuresDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("inactiveThreatAgentDeviceCount", m.GetInactiveThreatAgentDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("pendingFullScanDeviceCount", m.GetPendingFullScanDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("pendingManualStepsDeviceCount", m.GetPendingManualStepsDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("pendingOfflineScanDeviceCount", m.GetPendingOfflineScanDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("pendingQuickScanDeviceCount", m.GetPendingQuickScanDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("pendingRestartDeviceCount", m.GetPendingRestartDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("pendingSignatureUpdateDeviceCount", m.GetPendingSignatureUpdateDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalReportedDeviceCount", m.GetTotalReportedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("unknownStateThreatAgentDeviceCount", m.GetUnknownStateThreatAgentDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceProtectionOverview) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCleanDeviceCount sets the cleanDeviceCount property value. Clean device count.
func (m *DeviceProtectionOverview) SetCleanDeviceCount(value *int32)() {
    m.cleanDeviceCount = value
}
// SetCriticalFailuresDeviceCount sets the criticalFailuresDeviceCount property value. Critical failures device count.
func (m *DeviceProtectionOverview) SetCriticalFailuresDeviceCount(value *int32)() {
    m.criticalFailuresDeviceCount = value
}
// SetInactiveThreatAgentDeviceCount sets the inactiveThreatAgentDeviceCount property value. Device with inactive threat agent count
func (m *DeviceProtectionOverview) SetInactiveThreatAgentDeviceCount(value *int32)() {
    m.inactiveThreatAgentDeviceCount = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceProtectionOverview) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPendingFullScanDeviceCount sets the pendingFullScanDeviceCount property value. Pending full scan device count.
func (m *DeviceProtectionOverview) SetPendingFullScanDeviceCount(value *int32)() {
    m.pendingFullScanDeviceCount = value
}
// SetPendingManualStepsDeviceCount sets the pendingManualStepsDeviceCount property value. Pending manual steps device count.
func (m *DeviceProtectionOverview) SetPendingManualStepsDeviceCount(value *int32)() {
    m.pendingManualStepsDeviceCount = value
}
// SetPendingOfflineScanDeviceCount sets the pendingOfflineScanDeviceCount property value. Pending offline scan device count.
func (m *DeviceProtectionOverview) SetPendingOfflineScanDeviceCount(value *int32)() {
    m.pendingOfflineScanDeviceCount = value
}
// SetPendingQuickScanDeviceCount sets the pendingQuickScanDeviceCount property value. Pending quick scan device count. Valid values -2147483648 to 2147483647
func (m *DeviceProtectionOverview) SetPendingQuickScanDeviceCount(value *int32)() {
    m.pendingQuickScanDeviceCount = value
}
// SetPendingRestartDeviceCount sets the pendingRestartDeviceCount property value. Pending restart device count.
func (m *DeviceProtectionOverview) SetPendingRestartDeviceCount(value *int32)() {
    m.pendingRestartDeviceCount = value
}
// SetPendingSignatureUpdateDeviceCount sets the pendingSignatureUpdateDeviceCount property value. Device with old signature count.
func (m *DeviceProtectionOverview) SetPendingSignatureUpdateDeviceCount(value *int32)() {
    m.pendingSignatureUpdateDeviceCount = value
}
// SetTotalReportedDeviceCount sets the totalReportedDeviceCount property value. Total device count.
func (m *DeviceProtectionOverview) SetTotalReportedDeviceCount(value *int32)() {
    m.totalReportedDeviceCount = value
}
// SetUnknownStateThreatAgentDeviceCount sets the unknownStateThreatAgentDeviceCount property value. Device with threat agent state as unknown count.
func (m *DeviceProtectionOverview) SetUnknownStateThreatAgentDeviceCount(value *int32)() {
    m.unknownStateThreatAgentDeviceCount = value
}
