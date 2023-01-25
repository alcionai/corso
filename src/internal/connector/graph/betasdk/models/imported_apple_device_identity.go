package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ImportedAppleDeviceIdentity the importedAppleDeviceIdentity resource represents the imported device identity of an Apple device .
type ImportedAppleDeviceIdentity struct {
    Entity
    // Created Date Time of the device
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description of the device
    description *string
    // The discoverySource property
    discoverySource *DiscoverySource
    // The enrollmentState property
    enrollmentState *EnrollmentState
    // Indicates if the device is deleted from Apple Business Manager
    isDeleted *bool
    // Indicates if the Apple device is supervised. More information is at: https://support.apple.com/en-us/HT202837
    isSupervised *bool
    // Last Contacted Date Time of the device
    lastContactedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The platform property
    platform *Platform
    // The time enrollment profile was assigned to the device
    requestedEnrollmentProfileAssignmentDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Enrollment profile Id admin intends to apply to the device during next enrollment
    requestedEnrollmentProfileId *string
    // Device serial number
    serialNumber *string
}
// NewImportedAppleDeviceIdentity instantiates a new importedAppleDeviceIdentity and sets the default values.
func NewImportedAppleDeviceIdentity()(*ImportedAppleDeviceIdentity) {
    m := &ImportedAppleDeviceIdentity{
        Entity: *NewEntity(),
    }
    return m
}
// CreateImportedAppleDeviceIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateImportedAppleDeviceIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.importedAppleDeviceIdentityResult":
                        return NewImportedAppleDeviceIdentityResult(), nil
                }
            }
        }
    }
    return NewImportedAppleDeviceIdentity(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. Created Date Time of the device
func (m *ImportedAppleDeviceIdentity) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. The description of the device
func (m *ImportedAppleDeviceIdentity) GetDescription()(*string) {
    return m.description
}
// GetDiscoverySource gets the discoverySource property value. The discoverySource property
func (m *ImportedAppleDeviceIdentity) GetDiscoverySource()(*DiscoverySource) {
    return m.discoverySource
}
// GetEnrollmentState gets the enrollmentState property value. The enrollmentState property
func (m *ImportedAppleDeviceIdentity) GetEnrollmentState()(*EnrollmentState) {
    return m.enrollmentState
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ImportedAppleDeviceIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["discoverySource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDiscoverySource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscoverySource(val.(*DiscoverySource))
        }
        return nil
    }
    res["enrollmentState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnrollmentState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentState(val.(*EnrollmentState))
        }
        return nil
    }
    res["isDeleted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDeleted(val)
        }
        return nil
    }
    res["isSupervised"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSupervised(val)
        }
        return nil
    }
    res["lastContactedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastContactedDateTime(val)
        }
        return nil
    }
    res["platform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePlatform)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val.(*Platform))
        }
        return nil
    }
    res["requestedEnrollmentProfileAssignmentDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedEnrollmentProfileAssignmentDateTime(val)
        }
        return nil
    }
    res["requestedEnrollmentProfileId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedEnrollmentProfileId(val)
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
    return res
}
// GetIsDeleted gets the isDeleted property value. Indicates if the device is deleted from Apple Business Manager
func (m *ImportedAppleDeviceIdentity) GetIsDeleted()(*bool) {
    return m.isDeleted
}
// GetIsSupervised gets the isSupervised property value. Indicates if the Apple device is supervised. More information is at: https://support.apple.com/en-us/HT202837
func (m *ImportedAppleDeviceIdentity) GetIsSupervised()(*bool) {
    return m.isSupervised
}
// GetLastContactedDateTime gets the lastContactedDateTime property value. Last Contacted Date Time of the device
func (m *ImportedAppleDeviceIdentity) GetLastContactedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastContactedDateTime
}
// GetPlatform gets the platform property value. The platform property
func (m *ImportedAppleDeviceIdentity) GetPlatform()(*Platform) {
    return m.platform
}
// GetRequestedEnrollmentProfileAssignmentDateTime gets the requestedEnrollmentProfileAssignmentDateTime property value. The time enrollment profile was assigned to the device
func (m *ImportedAppleDeviceIdentity) GetRequestedEnrollmentProfileAssignmentDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.requestedEnrollmentProfileAssignmentDateTime
}
// GetRequestedEnrollmentProfileId gets the requestedEnrollmentProfileId property value. Enrollment profile Id admin intends to apply to the device during next enrollment
func (m *ImportedAppleDeviceIdentity) GetRequestedEnrollmentProfileId()(*string) {
    return m.requestedEnrollmentProfileId
}
// GetSerialNumber gets the serialNumber property value. Device serial number
func (m *ImportedAppleDeviceIdentity) GetSerialNumber()(*string) {
    return m.serialNumber
}
// Serialize serializes information the current object
func (m *ImportedAppleDeviceIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
    if m.GetDiscoverySource() != nil {
        cast := (*m.GetDiscoverySource()).String()
        err = writer.WriteStringValue("discoverySource", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEnrollmentState() != nil {
        cast := (*m.GetEnrollmentState()).String()
        err = writer.WriteStringValue("enrollmentState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDeleted", m.GetIsDeleted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSupervised", m.GetIsSupervised())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastContactedDateTime", m.GetLastContactedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPlatform() != nil {
        cast := (*m.GetPlatform()).String()
        err = writer.WriteStringValue("platform", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("requestedEnrollmentProfileAssignmentDateTime", m.GetRequestedEnrollmentProfileAssignmentDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestedEnrollmentProfileId", m.GetRequestedEnrollmentProfileId())
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
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. Created Date Time of the device
func (m *ImportedAppleDeviceIdentity) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. The description of the device
func (m *ImportedAppleDeviceIdentity) SetDescription(value *string)() {
    m.description = value
}
// SetDiscoverySource sets the discoverySource property value. The discoverySource property
func (m *ImportedAppleDeviceIdentity) SetDiscoverySource(value *DiscoverySource)() {
    m.discoverySource = value
}
// SetEnrollmentState sets the enrollmentState property value. The enrollmentState property
func (m *ImportedAppleDeviceIdentity) SetEnrollmentState(value *EnrollmentState)() {
    m.enrollmentState = value
}
// SetIsDeleted sets the isDeleted property value. Indicates if the device is deleted from Apple Business Manager
func (m *ImportedAppleDeviceIdentity) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
// SetIsSupervised sets the isSupervised property value. Indicates if the Apple device is supervised. More information is at: https://support.apple.com/en-us/HT202837
func (m *ImportedAppleDeviceIdentity) SetIsSupervised(value *bool)() {
    m.isSupervised = value
}
// SetLastContactedDateTime sets the lastContactedDateTime property value. Last Contacted Date Time of the device
func (m *ImportedAppleDeviceIdentity) SetLastContactedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastContactedDateTime = value
}
// SetPlatform sets the platform property value. The platform property
func (m *ImportedAppleDeviceIdentity) SetPlatform(value *Platform)() {
    m.platform = value
}
// SetRequestedEnrollmentProfileAssignmentDateTime sets the requestedEnrollmentProfileAssignmentDateTime property value. The time enrollment profile was assigned to the device
func (m *ImportedAppleDeviceIdentity) SetRequestedEnrollmentProfileAssignmentDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.requestedEnrollmentProfileAssignmentDateTime = value
}
// SetRequestedEnrollmentProfileId sets the requestedEnrollmentProfileId property value. Enrollment profile Id admin intends to apply to the device during next enrollment
func (m *ImportedAppleDeviceIdentity) SetRequestedEnrollmentProfileId(value *string)() {
    m.requestedEnrollmentProfileId = value
}
// SetSerialNumber sets the serialNumber property value. Device serial number
func (m *ImportedAppleDeviceIdentity) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
