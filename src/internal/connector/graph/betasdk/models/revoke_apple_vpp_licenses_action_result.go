package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RevokeAppleVppLicensesActionResult 
type RevokeAppleVppLicensesActionResult struct {
    DeviceActionResult
    // Total number of Apple Vpp licenses that failed to revoke
    failedLicensesCount *int32
    // Total number of Apple Vpp licenses associated
    totalLicensesCount *int32
}
// NewRevokeAppleVppLicensesActionResult instantiates a new RevokeAppleVppLicensesActionResult and sets the default values.
func NewRevokeAppleVppLicensesActionResult()(*RevokeAppleVppLicensesActionResult) {
    m := &RevokeAppleVppLicensesActionResult{
        DeviceActionResult: *NewDeviceActionResult(),
    }
    return m
}
// CreateRevokeAppleVppLicensesActionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRevokeAppleVppLicensesActionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRevokeAppleVppLicensesActionResult(), nil
}
// GetFailedLicensesCount gets the failedLicensesCount property value. Total number of Apple Vpp licenses that failed to revoke
func (m *RevokeAppleVppLicensesActionResult) GetFailedLicensesCount()(*int32) {
    return m.failedLicensesCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RevokeAppleVppLicensesActionResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceActionResult.GetFieldDeserializers()
    res["failedLicensesCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedLicensesCount(val)
        }
        return nil
    }
    res["totalLicensesCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalLicensesCount(val)
        }
        return nil
    }
    return res
}
// GetTotalLicensesCount gets the totalLicensesCount property value. Total number of Apple Vpp licenses associated
func (m *RevokeAppleVppLicensesActionResult) GetTotalLicensesCount()(*int32) {
    return m.totalLicensesCount
}
// Serialize serializes information the current object
func (m *RevokeAppleVppLicensesActionResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceActionResult.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("failedLicensesCount", m.GetFailedLicensesCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalLicensesCount", m.GetTotalLicensesCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFailedLicensesCount sets the failedLicensesCount property value. Total number of Apple Vpp licenses that failed to revoke
func (m *RevokeAppleVppLicensesActionResult) SetFailedLicensesCount(value *int32)() {
    m.failedLicensesCount = value
}
// SetTotalLicensesCount sets the totalLicensesCount property value. Total number of Apple Vpp licenses associated
func (m *RevokeAppleVppLicensesActionResult) SetTotalLicensesCount(value *int32)() {
    m.totalLicensesCount = value
}
