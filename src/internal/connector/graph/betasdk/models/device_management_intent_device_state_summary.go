package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementIntentDeviceStateSummary 
type DeviceManagementIntentDeviceStateSummary struct {
    Entity
    // Number of devices in conflict
    conflictCount *int32
    // Number of error devices
    errorCount *int32
    // Number of failed devices
    failedCount *int32
    // Number of not applicable devices
    notApplicableCount *int32
    // Number of not applicable devices due to mismatch platform and policy
    notApplicablePlatformCount *int32
    // Number of succeeded devices
    successCount *int32
}
// NewDeviceManagementIntentDeviceStateSummary instantiates a new deviceManagementIntentDeviceStateSummary and sets the default values.
func NewDeviceManagementIntentDeviceStateSummary()(*DeviceManagementIntentDeviceStateSummary) {
    m := &DeviceManagementIntentDeviceStateSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementIntentDeviceStateSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementIntentDeviceStateSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementIntentDeviceStateSummary(), nil
}
// GetConflictCount gets the conflictCount property value. Number of devices in conflict
func (m *DeviceManagementIntentDeviceStateSummary) GetConflictCount()(*int32) {
    return m.conflictCount
}
// GetErrorCount gets the errorCount property value. Number of error devices
func (m *DeviceManagementIntentDeviceStateSummary) GetErrorCount()(*int32) {
    return m.errorCount
}
// GetFailedCount gets the failedCount property value. Number of failed devices
func (m *DeviceManagementIntentDeviceStateSummary) GetFailedCount()(*int32) {
    return m.failedCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementIntentDeviceStateSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["conflictCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConflictCount(val)
        }
        return nil
    }
    res["errorCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCount(val)
        }
        return nil
    }
    res["failedCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedCount(val)
        }
        return nil
    }
    res["notApplicableCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotApplicableCount(val)
        }
        return nil
    }
    res["notApplicablePlatformCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotApplicablePlatformCount(val)
        }
        return nil
    }
    res["successCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessCount(val)
        }
        return nil
    }
    return res
}
// GetNotApplicableCount gets the notApplicableCount property value. Number of not applicable devices
func (m *DeviceManagementIntentDeviceStateSummary) GetNotApplicableCount()(*int32) {
    return m.notApplicableCount
}
// GetNotApplicablePlatformCount gets the notApplicablePlatformCount property value. Number of not applicable devices due to mismatch platform and policy
func (m *DeviceManagementIntentDeviceStateSummary) GetNotApplicablePlatformCount()(*int32) {
    return m.notApplicablePlatformCount
}
// GetSuccessCount gets the successCount property value. Number of succeeded devices
func (m *DeviceManagementIntentDeviceStateSummary) GetSuccessCount()(*int32) {
    return m.successCount
}
// Serialize serializes information the current object
func (m *DeviceManagementIntentDeviceStateSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("conflictCount", m.GetConflictCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("errorCount", m.GetErrorCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("failedCount", m.GetFailedCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notApplicableCount", m.GetNotApplicableCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notApplicablePlatformCount", m.GetNotApplicablePlatformCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("successCount", m.GetSuccessCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConflictCount sets the conflictCount property value. Number of devices in conflict
func (m *DeviceManagementIntentDeviceStateSummary) SetConflictCount(value *int32)() {
    m.conflictCount = value
}
// SetErrorCount sets the errorCount property value. Number of error devices
func (m *DeviceManagementIntentDeviceStateSummary) SetErrorCount(value *int32)() {
    m.errorCount = value
}
// SetFailedCount sets the failedCount property value. Number of failed devices
func (m *DeviceManagementIntentDeviceStateSummary) SetFailedCount(value *int32)() {
    m.failedCount = value
}
// SetNotApplicableCount sets the notApplicableCount property value. Number of not applicable devices
func (m *DeviceManagementIntentDeviceStateSummary) SetNotApplicableCount(value *int32)() {
    m.notApplicableCount = value
}
// SetNotApplicablePlatformCount sets the notApplicablePlatformCount property value. Number of not applicable devices due to mismatch platform and policy
func (m *DeviceManagementIntentDeviceStateSummary) SetNotApplicablePlatformCount(value *int32)() {
    m.notApplicablePlatformCount = value
}
// SetSuccessCount sets the successCount property value. Number of succeeded devices
func (m *DeviceManagementIntentDeviceStateSummary) SetSuccessCount(value *int32)() {
    m.successCount = value
}
