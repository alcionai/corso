package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComanagementEligibleDevice device Co-Management eligibility state
type ComanagementEligibleDevice struct {
    Entity
    // Device registration status.
    clientRegistrationStatus *DeviceRegistrationState
    // DeviceName
    deviceName *string
    // Device type.
    deviceType *DeviceType
    // EntitySource
    entitySource *int32
    // Management agent type.
    managementAgents *ManagementAgentType
    // Management state of device in Microsoft Intune.
    managementState *ManagementState
    // Manufacturer
    manufacturer *string
    // MDMStatus
    mdmStatus *string
    // Model
    model *string
    // OSDescription
    osDescription *string
    // OSVersion
    osVersion *string
    // Owner type of device.
    ownerType *OwnerType
    // ReferenceId
    referenceId *string
    // SerialNumber
    serialNumber *string
    // The status property
    status *ComanagementEligibleType
    // UPN
    upn *string
    // UserEmail
    userEmail *string
    // UserId
    userId *string
    // UserName
    userName *string
}
// NewComanagementEligibleDevice instantiates a new comanagementEligibleDevice and sets the default values.
func NewComanagementEligibleDevice()(*ComanagementEligibleDevice) {
    m := &ComanagementEligibleDevice{
        Entity: *NewEntity(),
    }
    return m
}
// CreateComanagementEligibleDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateComanagementEligibleDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComanagementEligibleDevice(), nil
}
// GetClientRegistrationStatus gets the clientRegistrationStatus property value. Device registration status.
func (m *ComanagementEligibleDevice) GetClientRegistrationStatus()(*DeviceRegistrationState) {
    return m.clientRegistrationStatus
}
// GetDeviceName gets the deviceName property value. DeviceName
func (m *ComanagementEligibleDevice) GetDeviceName()(*string) {
    return m.deviceName
}
// GetDeviceType gets the deviceType property value. Device type.
func (m *ComanagementEligibleDevice) GetDeviceType()(*DeviceType) {
    return m.deviceType
}
// GetEntitySource gets the entitySource property value. EntitySource
func (m *ComanagementEligibleDevice) GetEntitySource()(*int32) {
    return m.entitySource
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ComanagementEligibleDevice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["clientRegistrationStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceRegistrationState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientRegistrationStatus(val.(*DeviceRegistrationState))
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["deviceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceType(val.(*DeviceType))
        }
        return nil
    }
    res["entitySource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntitySource(val)
        }
        return nil
    }
    res["managementAgents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementAgentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementAgents(val.(*ManagementAgentType))
        }
        return nil
    }
    res["managementState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementState(val.(*ManagementState))
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
    res["mdmStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMdmStatus(val)
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
        val, err := n.GetEnumValue(ParseOwnerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerType(val.(*OwnerType))
        }
        return nil
    }
    res["referenceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferenceId(val)
        }
        return nil
    }
    res["serialNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSerialNumber(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseComanagementEligibleType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ComanagementEligibleType))
        }
        return nil
    }
    res["upn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpn(val)
        }
        return nil
    }
    res["userEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserEmail(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserName(val)
        }
        return nil
    }
    return res
}
// GetManagementAgents gets the managementAgents property value. Management agent type.
func (m *ComanagementEligibleDevice) GetManagementAgents()(*ManagementAgentType) {
    return m.managementAgents
}
// GetManagementState gets the managementState property value. Management state of device in Microsoft Intune.
func (m *ComanagementEligibleDevice) GetManagementState()(*ManagementState) {
    return m.managementState
}
// GetManufacturer gets the manufacturer property value. Manufacturer
func (m *ComanagementEligibleDevice) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetMdmStatus gets the mdmStatus property value. MDMStatus
func (m *ComanagementEligibleDevice) GetMdmStatus()(*string) {
    return m.mdmStatus
}
// GetModel gets the model property value. Model
func (m *ComanagementEligibleDevice) GetModel()(*string) {
    return m.model
}
// GetOsDescription gets the osDescription property value. OSDescription
func (m *ComanagementEligibleDevice) GetOsDescription()(*string) {
    return m.osDescription
}
// GetOsVersion gets the osVersion property value. OSVersion
func (m *ComanagementEligibleDevice) GetOsVersion()(*string) {
    return m.osVersion
}
// GetOwnerType gets the ownerType property value. Owner type of device.
func (m *ComanagementEligibleDevice) GetOwnerType()(*OwnerType) {
    return m.ownerType
}
// GetReferenceId gets the referenceId property value. ReferenceId
func (m *ComanagementEligibleDevice) GetReferenceId()(*string) {
    return m.referenceId
}
// GetSerialNumber gets the serialNumber property value. SerialNumber
func (m *ComanagementEligibleDevice) GetSerialNumber()(*string) {
    return m.serialNumber
}
// GetStatus gets the status property value. The status property
func (m *ComanagementEligibleDevice) GetStatus()(*ComanagementEligibleType) {
    return m.status
}
// GetUpn gets the upn property value. UPN
func (m *ComanagementEligibleDevice) GetUpn()(*string) {
    return m.upn
}
// GetUserEmail gets the userEmail property value. UserEmail
func (m *ComanagementEligibleDevice) GetUserEmail()(*string) {
    return m.userEmail
}
// GetUserId gets the userId property value. UserId
func (m *ComanagementEligibleDevice) GetUserId()(*string) {
    return m.userId
}
// GetUserName gets the userName property value. UserName
func (m *ComanagementEligibleDevice) GetUserName()(*string) {
    return m.userName
}
// Serialize serializes information the current object
func (m *ComanagementEligibleDevice) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClientRegistrationStatus() != nil {
        cast := (*m.GetClientRegistrationStatus()).String()
        err = writer.WriteStringValue("clientRegistrationStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceType() != nil {
        cast := (*m.GetDeviceType()).String()
        err = writer.WriteStringValue("deviceType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("entitySource", m.GetEntitySource())
        if err != nil {
            return err
        }
    }
    if m.GetManagementAgents() != nil {
        cast := (*m.GetManagementAgents()).String()
        err = writer.WriteStringValue("managementAgents", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementState() != nil {
        cast := (*m.GetManagementState()).String()
        err = writer.WriteStringValue("managementState", &cast)
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
        err = writer.WriteStringValue("mdmStatus", m.GetMdmStatus())
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
    if m.GetOwnerType() != nil {
        cast := (*m.GetOwnerType()).String()
        err = writer.WriteStringValue("ownerType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("referenceId", m.GetReferenceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serialNumber", m.GetSerialNumber())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("upn", m.GetUpn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userEmail", m.GetUserEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userName", m.GetUserName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClientRegistrationStatus sets the clientRegistrationStatus property value. Device registration status.
func (m *ComanagementEligibleDevice) SetClientRegistrationStatus(value *DeviceRegistrationState)() {
    m.clientRegistrationStatus = value
}
// SetDeviceName sets the deviceName property value. DeviceName
func (m *ComanagementEligibleDevice) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetDeviceType sets the deviceType property value. Device type.
func (m *ComanagementEligibleDevice) SetDeviceType(value *DeviceType)() {
    m.deviceType = value
}
// SetEntitySource sets the entitySource property value. EntitySource
func (m *ComanagementEligibleDevice) SetEntitySource(value *int32)() {
    m.entitySource = value
}
// SetManagementAgents sets the managementAgents property value. Management agent type.
func (m *ComanagementEligibleDevice) SetManagementAgents(value *ManagementAgentType)() {
    m.managementAgents = value
}
// SetManagementState sets the managementState property value. Management state of device in Microsoft Intune.
func (m *ComanagementEligibleDevice) SetManagementState(value *ManagementState)() {
    m.managementState = value
}
// SetManufacturer sets the manufacturer property value. Manufacturer
func (m *ComanagementEligibleDevice) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetMdmStatus sets the mdmStatus property value. MDMStatus
func (m *ComanagementEligibleDevice) SetMdmStatus(value *string)() {
    m.mdmStatus = value
}
// SetModel sets the model property value. Model
func (m *ComanagementEligibleDevice) SetModel(value *string)() {
    m.model = value
}
// SetOsDescription sets the osDescription property value. OSDescription
func (m *ComanagementEligibleDevice) SetOsDescription(value *string)() {
    m.osDescription = value
}
// SetOsVersion sets the osVersion property value. OSVersion
func (m *ComanagementEligibleDevice) SetOsVersion(value *string)() {
    m.osVersion = value
}
// SetOwnerType sets the ownerType property value. Owner type of device.
func (m *ComanagementEligibleDevice) SetOwnerType(value *OwnerType)() {
    m.ownerType = value
}
// SetReferenceId sets the referenceId property value. ReferenceId
func (m *ComanagementEligibleDevice) SetReferenceId(value *string)() {
    m.referenceId = value
}
// SetSerialNumber sets the serialNumber property value. SerialNumber
func (m *ComanagementEligibleDevice) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
// SetStatus sets the status property value. The status property
func (m *ComanagementEligibleDevice) SetStatus(value *ComanagementEligibleType)() {
    m.status = value
}
// SetUpn sets the upn property value. UPN
func (m *ComanagementEligibleDevice) SetUpn(value *string)() {
    m.upn = value
}
// SetUserEmail sets the userEmail property value. UserEmail
func (m *ComanagementEligibleDevice) SetUserEmail(value *string)() {
    m.userEmail = value
}
// SetUserId sets the userId property value. UserId
func (m *ComanagementEligibleDevice) SetUserId(value *string)() {
    m.userId = value
}
// SetUserName sets the userName property value. UserName
func (m *ComanagementEligibleDevice) SetUserName(value *string)() {
    m.userName = value
}
