package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsDriverUpdateInventory a new entity to represent driver inventories.
type WindowsDriverUpdateInventory struct {
    Entity
    // The number of devices for which this driver is applicable.
    applicableDeviceCount *int32
    // An enum type to represent approval status of a driver.
    approvalStatus *DriverApprovalStatus
    // An enum type to represent which category a driver belongs to.
    category *DriverCategory
    // The date time when a driver should be deployed if approvalStatus is approved.
    deployDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The class of the driver.
    driverClass *string
    // The manufacturer of the driver.
    manufacturer *string
    // The name of the driver.
    name *string
    // The release date time of the driver.
    releaseDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The version of the driver.
    version *string
}
// NewWindowsDriverUpdateInventory instantiates a new windowsDriverUpdateInventory and sets the default values.
func NewWindowsDriverUpdateInventory()(*WindowsDriverUpdateInventory) {
    m := &WindowsDriverUpdateInventory{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsDriverUpdateInventoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsDriverUpdateInventoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsDriverUpdateInventory(), nil
}
// GetApplicableDeviceCount gets the applicableDeviceCount property value. The number of devices for which this driver is applicable.
func (m *WindowsDriverUpdateInventory) GetApplicableDeviceCount()(*int32) {
    return m.applicableDeviceCount
}
// GetApprovalStatus gets the approvalStatus property value. An enum type to represent approval status of a driver.
func (m *WindowsDriverUpdateInventory) GetApprovalStatus()(*DriverApprovalStatus) {
    return m.approvalStatus
}
// GetCategory gets the category property value. An enum type to represent which category a driver belongs to.
func (m *WindowsDriverUpdateInventory) GetCategory()(*DriverCategory) {
    return m.category
}
// GetDeployDateTime gets the deployDateTime property value. The date time when a driver should be deployed if approvalStatus is approved.
func (m *WindowsDriverUpdateInventory) GetDeployDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.deployDateTime
}
// GetDriverClass gets the driverClass property value. The class of the driver.
func (m *WindowsDriverUpdateInventory) GetDriverClass()(*string) {
    return m.driverClass
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsDriverUpdateInventory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicableDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicableDeviceCount(val)
        }
        return nil
    }
    res["approvalStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDriverApprovalStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprovalStatus(val.(*DriverApprovalStatus))
        }
        return nil
    }
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDriverCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*DriverCategory))
        }
        return nil
    }
    res["deployDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeployDateTime(val)
        }
        return nil
    }
    res["driverClass"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDriverClass(val)
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
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["releaseDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReleaseDateTime(val)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetManufacturer gets the manufacturer property value. The manufacturer of the driver.
func (m *WindowsDriverUpdateInventory) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetName gets the name property value. The name of the driver.
func (m *WindowsDriverUpdateInventory) GetName()(*string) {
    return m.name
}
// GetReleaseDateTime gets the releaseDateTime property value. The release date time of the driver.
func (m *WindowsDriverUpdateInventory) GetReleaseDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.releaseDateTime
}
// GetVersion gets the version property value. The version of the driver.
func (m *WindowsDriverUpdateInventory) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *WindowsDriverUpdateInventory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("applicableDeviceCount", m.GetApplicableDeviceCount())
        if err != nil {
            return err
        }
    }
    if m.GetApprovalStatus() != nil {
        cast := (*m.GetApprovalStatus()).String()
        err = writer.WriteStringValue("approvalStatus", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
        err = writer.WriteStringValue("category", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("deployDateTime", m.GetDeployDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("driverClass", m.GetDriverClass())
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
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("releaseDateTime", m.GetReleaseDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicableDeviceCount sets the applicableDeviceCount property value. The number of devices for which this driver is applicable.
func (m *WindowsDriverUpdateInventory) SetApplicableDeviceCount(value *int32)() {
    m.applicableDeviceCount = value
}
// SetApprovalStatus sets the approvalStatus property value. An enum type to represent approval status of a driver.
func (m *WindowsDriverUpdateInventory) SetApprovalStatus(value *DriverApprovalStatus)() {
    m.approvalStatus = value
}
// SetCategory sets the category property value. An enum type to represent which category a driver belongs to.
func (m *WindowsDriverUpdateInventory) SetCategory(value *DriverCategory)() {
    m.category = value
}
// SetDeployDateTime sets the deployDateTime property value. The date time when a driver should be deployed if approvalStatus is approved.
func (m *WindowsDriverUpdateInventory) SetDeployDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.deployDateTime = value
}
// SetDriverClass sets the driverClass property value. The class of the driver.
func (m *WindowsDriverUpdateInventory) SetDriverClass(value *string)() {
    m.driverClass = value
}
// SetManufacturer sets the manufacturer property value. The manufacturer of the driver.
func (m *WindowsDriverUpdateInventory) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetName sets the name property value. The name of the driver.
func (m *WindowsDriverUpdateInventory) SetName(value *string)() {
    m.name = value
}
// SetReleaseDateTime sets the releaseDateTime property value. The release date time of the driver.
func (m *WindowsDriverUpdateInventory) SetReleaseDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.releaseDateTime = value
}
// SetVersion sets the version property value. The version of the driver.
func (m *WindowsDriverUpdateInventory) SetVersion(value *string)() {
    m.version = value
}
