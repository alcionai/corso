package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosLobAppProvisioningConfiguration this topic provides descriptions of the declared methods, properties and relationships exposed by the iOS Lob App Provisioning Configuration resource.
type IosLobAppProvisioningConfiguration struct {
    Entity
    // The associated group assignments for IosLobAppProvisioningConfiguration.
    assignments []IosLobAppProvisioningConfigurationAssignmentable
    // DateTime the object was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Admin provided description of the Device Configuration.
    description *string
    // The list of device installation states for this mobile app configuration.
    deviceStatuses []ManagedDeviceMobileAppConfigurationDeviceStatusable
    // Admin provided name of the device configuration.
    displayName *string
    // Optional profile expiration date and time.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The associated group assignments.
    groupAssignments []MobileAppProvisioningConfigGroupAssignmentable
    // DateTime the object was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Payload. (UTF8 encoded byte array)
    payload []byte
    // Payload file name (.mobileprovision
    payloadFileName *string
    // List of Scope Tags for this iOS LOB app provisioning configuration entity.
    roleScopeTagIds []string
    // The list of user installation states for this mobile app configuration.
    userStatuses []ManagedDeviceMobileAppConfigurationUserStatusable
    // Version of the device configuration.
    version *int32
}
// NewIosLobAppProvisioningConfiguration instantiates a new iosLobAppProvisioningConfiguration and sets the default values.
func NewIosLobAppProvisioningConfiguration()(*IosLobAppProvisioningConfiguration) {
    m := &IosLobAppProvisioningConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// CreateIosLobAppProvisioningConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosLobAppProvisioningConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosLobAppProvisioningConfiguration(), nil
}
// GetAssignments gets the assignments property value. The associated group assignments for IosLobAppProvisioningConfiguration.
func (m *IosLobAppProvisioningConfiguration) GetAssignments()([]IosLobAppProvisioningConfigurationAssignmentable) {
    return m.assignments
}
// GetCreatedDateTime gets the createdDateTime property value. DateTime the object was created.
func (m *IosLobAppProvisioningConfiguration) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. Admin provided description of the Device Configuration.
func (m *IosLobAppProvisioningConfiguration) GetDescription()(*string) {
    return m.description
}
// GetDeviceStatuses gets the deviceStatuses property value. The list of device installation states for this mobile app configuration.
func (m *IosLobAppProvisioningConfiguration) GetDeviceStatuses()([]ManagedDeviceMobileAppConfigurationDeviceStatusable) {
    return m.deviceStatuses
}
// GetDisplayName gets the displayName property value. Admin provided name of the device configuration.
func (m *IosLobAppProvisioningConfiguration) GetDisplayName()(*string) {
    return m.displayName
}
// GetExpirationDateTime gets the expirationDateTime property value. Optional profile expiration date and time.
func (m *IosLobAppProvisioningConfiguration) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosLobAppProvisioningConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIosLobAppProvisioningConfigurationAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosLobAppProvisioningConfigurationAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(IosLobAppProvisioningConfigurationAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["deviceStatuses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceMobileAppConfigurationDeviceStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceMobileAppConfigurationDeviceStatusable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedDeviceMobileAppConfigurationDeviceStatusable)
            }
            m.SetDeviceStatuses(res)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["groupAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileAppProvisioningConfigGroupAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppProvisioningConfigGroupAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(MobileAppProvisioningConfigGroupAssignmentable)
            }
            m.SetGroupAssignments(res)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["payload"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayload(val)
        }
        return nil
    }
    res["payloadFileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadFileName(val)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["userStatuses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceMobileAppConfigurationUserStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceMobileAppConfigurationUserStatusable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedDeviceMobileAppConfigurationUserStatusable)
            }
            m.SetUserStatuses(res)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
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
// GetGroupAssignments gets the groupAssignments property value. The associated group assignments.
func (m *IosLobAppProvisioningConfiguration) GetGroupAssignments()([]MobileAppProvisioningConfigGroupAssignmentable) {
    return m.groupAssignments
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *IosLobAppProvisioningConfiguration) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetPayload gets the payload property value. Payload. (UTF8 encoded byte array)
func (m *IosLobAppProvisioningConfiguration) GetPayload()([]byte) {
    return m.payload
}
// GetPayloadFileName gets the payloadFileName property value. Payload file name (.mobileprovision
func (m *IosLobAppProvisioningConfiguration) GetPayloadFileName()(*string) {
    return m.payloadFileName
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this iOS LOB app provisioning configuration entity.
func (m *IosLobAppProvisioningConfiguration) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// GetUserStatuses gets the userStatuses property value. The list of user installation states for this mobile app configuration.
func (m *IosLobAppProvisioningConfiguration) GetUserStatuses()([]ManagedDeviceMobileAppConfigurationUserStatusable) {
    return m.userStatuses
}
// GetVersion gets the version property value. Version of the device configuration.
func (m *IosLobAppProvisioningConfiguration) GetVersion()(*int32) {
    return m.version
}
// Serialize serializes information the current object
func (m *IosLobAppProvisioningConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceStatuses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceStatuses()))
        for i, v := range m.GetDeviceStatuses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetGroupAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroupAssignments()))
        for i, v := range m.GetGroupAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("groupAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("payload", m.GetPayload())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("payloadFileName", m.GetPayloadFileName())
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    if m.GetUserStatuses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserStatuses()))
        for i, v := range m.GetUserStatuses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The associated group assignments for IosLobAppProvisioningConfiguration.
func (m *IosLobAppProvisioningConfiguration) SetAssignments(value []IosLobAppProvisioningConfigurationAssignmentable)() {
    m.assignments = value
}
// SetCreatedDateTime sets the createdDateTime property value. DateTime the object was created.
func (m *IosLobAppProvisioningConfiguration) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Admin provided description of the Device Configuration.
func (m *IosLobAppProvisioningConfiguration) SetDescription(value *string)() {
    m.description = value
}
// SetDeviceStatuses sets the deviceStatuses property value. The list of device installation states for this mobile app configuration.
func (m *IosLobAppProvisioningConfiguration) SetDeviceStatuses(value []ManagedDeviceMobileAppConfigurationDeviceStatusable)() {
    m.deviceStatuses = value
}
// SetDisplayName sets the displayName property value. Admin provided name of the device configuration.
func (m *IosLobAppProvisioningConfiguration) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExpirationDateTime sets the expirationDateTime property value. Optional profile expiration date and time.
func (m *IosLobAppProvisioningConfiguration) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetGroupAssignments sets the groupAssignments property value. The associated group assignments.
func (m *IosLobAppProvisioningConfiguration) SetGroupAssignments(value []MobileAppProvisioningConfigGroupAssignmentable)() {
    m.groupAssignments = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *IosLobAppProvisioningConfiguration) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetPayload sets the payload property value. Payload. (UTF8 encoded byte array)
func (m *IosLobAppProvisioningConfiguration) SetPayload(value []byte)() {
    m.payload = value
}
// SetPayloadFileName sets the payloadFileName property value. Payload file name (.mobileprovision
func (m *IosLobAppProvisioningConfiguration) SetPayloadFileName(value *string)() {
    m.payloadFileName = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this iOS LOB app provisioning configuration entity.
func (m *IosLobAppProvisioningConfiguration) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// SetUserStatuses sets the userStatuses property value. The list of user installation states for this mobile app configuration.
func (m *IosLobAppProvisioningConfiguration) SetUserStatuses(value []ManagedDeviceMobileAppConfigurationUserStatusable)() {
    m.userStatuses = value
}
// SetVersion sets the version property value. Version of the device configuration.
func (m *IosLobAppProvisioningConfiguration) SetVersion(value *int32)() {
    m.version = value
}
