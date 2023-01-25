package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppInstallSummary 
type MobileAppInstallSummary struct {
    Entity
    // Number of Devices that have failed to install this app.
    failedDeviceCount *int32
    // Number of Users that have 1 or more device that failed to install this app.
    failedUserCount *int32
    // Number of Devices that have successfully installed this app.
    installedDeviceCount *int32
    // Number of Users whose devices have all succeeded to install this app.
    installedUserCount *int32
    // Number of Devices that are not applicable for this app.
    notApplicableDeviceCount *int32
    // Number of Users whose devices were all not applicable for this app.
    notApplicableUserCount *int32
    // Number of Devices that does not have this app installed.
    notInstalledDeviceCount *int32
    // Number of Users that have 1 or more devices that did not install this app.
    notInstalledUserCount *int32
    // Number of Devices that have been notified to install this app.
    pendingInstallDeviceCount *int32
    // Number of Users that have 1 or more device that have been notified to install this app and have 0 devices with failures.
    pendingInstallUserCount *int32
}
// NewMobileAppInstallSummary instantiates a new mobileAppInstallSummary and sets the default values.
func NewMobileAppInstallSummary()(*MobileAppInstallSummary) {
    m := &MobileAppInstallSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMobileAppInstallSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppInstallSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppInstallSummary(), nil
}
// GetFailedDeviceCount gets the failedDeviceCount property value. Number of Devices that have failed to install this app.
func (m *MobileAppInstallSummary) GetFailedDeviceCount()(*int32) {
    return m.failedDeviceCount
}
// GetFailedUserCount gets the failedUserCount property value. Number of Users that have 1 or more device that failed to install this app.
func (m *MobileAppInstallSummary) GetFailedUserCount()(*int32) {
    return m.failedUserCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppInstallSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["failedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedDeviceCount(val)
        }
        return nil
    }
    res["failedUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedUserCount(val)
        }
        return nil
    }
    res["installedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstalledDeviceCount(val)
        }
        return nil
    }
    res["installedUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstalledUserCount(val)
        }
        return nil
    }
    res["notApplicableDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotApplicableDeviceCount(val)
        }
        return nil
    }
    res["notApplicableUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotApplicableUserCount(val)
        }
        return nil
    }
    res["notInstalledDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotInstalledDeviceCount(val)
        }
        return nil
    }
    res["notInstalledUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotInstalledUserCount(val)
        }
        return nil
    }
    res["pendingInstallDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingInstallDeviceCount(val)
        }
        return nil
    }
    res["pendingInstallUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingInstallUserCount(val)
        }
        return nil
    }
    return res
}
// GetInstalledDeviceCount gets the installedDeviceCount property value. Number of Devices that have successfully installed this app.
func (m *MobileAppInstallSummary) GetInstalledDeviceCount()(*int32) {
    return m.installedDeviceCount
}
// GetInstalledUserCount gets the installedUserCount property value. Number of Users whose devices have all succeeded to install this app.
func (m *MobileAppInstallSummary) GetInstalledUserCount()(*int32) {
    return m.installedUserCount
}
// GetNotApplicableDeviceCount gets the notApplicableDeviceCount property value. Number of Devices that are not applicable for this app.
func (m *MobileAppInstallSummary) GetNotApplicableDeviceCount()(*int32) {
    return m.notApplicableDeviceCount
}
// GetNotApplicableUserCount gets the notApplicableUserCount property value. Number of Users whose devices were all not applicable for this app.
func (m *MobileAppInstallSummary) GetNotApplicableUserCount()(*int32) {
    return m.notApplicableUserCount
}
// GetNotInstalledDeviceCount gets the notInstalledDeviceCount property value. Number of Devices that does not have this app installed.
func (m *MobileAppInstallSummary) GetNotInstalledDeviceCount()(*int32) {
    return m.notInstalledDeviceCount
}
// GetNotInstalledUserCount gets the notInstalledUserCount property value. Number of Users that have 1 or more devices that did not install this app.
func (m *MobileAppInstallSummary) GetNotInstalledUserCount()(*int32) {
    return m.notInstalledUserCount
}
// GetPendingInstallDeviceCount gets the pendingInstallDeviceCount property value. Number of Devices that have been notified to install this app.
func (m *MobileAppInstallSummary) GetPendingInstallDeviceCount()(*int32) {
    return m.pendingInstallDeviceCount
}
// GetPendingInstallUserCount gets the pendingInstallUserCount property value. Number of Users that have 1 or more device that have been notified to install this app and have 0 devices with failures.
func (m *MobileAppInstallSummary) GetPendingInstallUserCount()(*int32) {
    return m.pendingInstallUserCount
}
// Serialize serializes information the current object
func (m *MobileAppInstallSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("failedDeviceCount", m.GetFailedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("failedUserCount", m.GetFailedUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("installedDeviceCount", m.GetInstalledDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("installedUserCount", m.GetInstalledUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notApplicableDeviceCount", m.GetNotApplicableDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notApplicableUserCount", m.GetNotApplicableUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notInstalledDeviceCount", m.GetNotInstalledDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notInstalledUserCount", m.GetNotInstalledUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("pendingInstallDeviceCount", m.GetPendingInstallDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("pendingInstallUserCount", m.GetPendingInstallUserCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFailedDeviceCount sets the failedDeviceCount property value. Number of Devices that have failed to install this app.
func (m *MobileAppInstallSummary) SetFailedDeviceCount(value *int32)() {
    m.failedDeviceCount = value
}
// SetFailedUserCount sets the failedUserCount property value. Number of Users that have 1 or more device that failed to install this app.
func (m *MobileAppInstallSummary) SetFailedUserCount(value *int32)() {
    m.failedUserCount = value
}
// SetInstalledDeviceCount sets the installedDeviceCount property value. Number of Devices that have successfully installed this app.
func (m *MobileAppInstallSummary) SetInstalledDeviceCount(value *int32)() {
    m.installedDeviceCount = value
}
// SetInstalledUserCount sets the installedUserCount property value. Number of Users whose devices have all succeeded to install this app.
func (m *MobileAppInstallSummary) SetInstalledUserCount(value *int32)() {
    m.installedUserCount = value
}
// SetNotApplicableDeviceCount sets the notApplicableDeviceCount property value. Number of Devices that are not applicable for this app.
func (m *MobileAppInstallSummary) SetNotApplicableDeviceCount(value *int32)() {
    m.notApplicableDeviceCount = value
}
// SetNotApplicableUserCount sets the notApplicableUserCount property value. Number of Users whose devices were all not applicable for this app.
func (m *MobileAppInstallSummary) SetNotApplicableUserCount(value *int32)() {
    m.notApplicableUserCount = value
}
// SetNotInstalledDeviceCount sets the notInstalledDeviceCount property value. Number of Devices that does not have this app installed.
func (m *MobileAppInstallSummary) SetNotInstalledDeviceCount(value *int32)() {
    m.notInstalledDeviceCount = value
}
// SetNotInstalledUserCount sets the notInstalledUserCount property value. Number of Users that have 1 or more devices that did not install this app.
func (m *MobileAppInstallSummary) SetNotInstalledUserCount(value *int32)() {
    m.notInstalledUserCount = value
}
// SetPendingInstallDeviceCount sets the pendingInstallDeviceCount property value. Number of Devices that have been notified to install this app.
func (m *MobileAppInstallSummary) SetPendingInstallDeviceCount(value *int32)() {
    m.pendingInstallDeviceCount = value
}
// SetPendingInstallUserCount sets the pendingInstallUserCount property value. Number of Users that have 1 or more device that have been notified to install this app and have 0 devices with failures.
func (m *MobileAppInstallSummary) SetPendingInstallUserCount(value *int32)() {
    m.pendingInstallUserCount = value
}
