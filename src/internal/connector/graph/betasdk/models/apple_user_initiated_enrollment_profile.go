package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppleUserInitiatedEnrollmentProfile the enrollmentProfile resource represents a collection of configurations which must be provided pre-enrollment to enable enrolling certain devices whose identities have been pre-staged. Pre-staged device identities are assigned to this type of profile to apply the profile's configurations at enrollment of the corresponding device.
type AppleUserInitiatedEnrollmentProfile struct {
    Entity
    // The list of assignments for this profile.
    assignments []AppleEnrollmentProfileAssignmentable
    // List of available enrollment type options
    availableEnrollmentTypeOptions []AppleOwnerTypeEnrollmentTypeable
    // Profile creation time
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The defaultEnrollmentType property
    defaultEnrollmentType *AppleUserInitiatedEnrollmentType
    // Description of the profile
    description *string
    // Name of the profile
    displayName *string
    // Profile last modified time
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Supported platform types.
    platform *DevicePlatformType
    // Priority, 0 is highest
    priority *int32
}
// NewAppleUserInitiatedEnrollmentProfile instantiates a new appleUserInitiatedEnrollmentProfile and sets the default values.
func NewAppleUserInitiatedEnrollmentProfile()(*AppleUserInitiatedEnrollmentProfile) {
    m := &AppleUserInitiatedEnrollmentProfile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAppleUserInitiatedEnrollmentProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppleUserInitiatedEnrollmentProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppleUserInitiatedEnrollmentProfile(), nil
}
// GetAssignments gets the assignments property value. The list of assignments for this profile.
func (m *AppleUserInitiatedEnrollmentProfile) GetAssignments()([]AppleEnrollmentProfileAssignmentable) {
    return m.assignments
}
// GetAvailableEnrollmentTypeOptions gets the availableEnrollmentTypeOptions property value. List of available enrollment type options
func (m *AppleUserInitiatedEnrollmentProfile) GetAvailableEnrollmentTypeOptions()([]AppleOwnerTypeEnrollmentTypeable) {
    return m.availableEnrollmentTypeOptions
}
// GetCreatedDateTime gets the createdDateTime property value. Profile creation time
func (m *AppleUserInitiatedEnrollmentProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDefaultEnrollmentType gets the defaultEnrollmentType property value. The defaultEnrollmentType property
func (m *AppleUserInitiatedEnrollmentProfile) GetDefaultEnrollmentType()(*AppleUserInitiatedEnrollmentType) {
    return m.defaultEnrollmentType
}
// GetDescription gets the description property value. Description of the profile
func (m *AppleUserInitiatedEnrollmentProfile) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Name of the profile
func (m *AppleUserInitiatedEnrollmentProfile) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppleUserInitiatedEnrollmentProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppleEnrollmentProfileAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppleEnrollmentProfileAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(AppleEnrollmentProfileAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["availableEnrollmentTypeOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppleOwnerTypeEnrollmentTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppleOwnerTypeEnrollmentTypeable, len(val))
            for i, v := range val {
                res[i] = v.(AppleOwnerTypeEnrollmentTypeable)
            }
            m.SetAvailableEnrollmentTypeOptions(res)
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
    res["defaultEnrollmentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppleUserInitiatedEnrollmentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultEnrollmentType(val.(*AppleUserInitiatedEnrollmentType))
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
    res["platform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDevicePlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val.(*DevicePlatformType))
        }
        return nil
    }
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Profile last modified time
func (m *AppleUserInitiatedEnrollmentProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetPlatform gets the platform property value. Supported platform types.
func (m *AppleUserInitiatedEnrollmentProfile) GetPlatform()(*DevicePlatformType) {
    return m.platform
}
// GetPriority gets the priority property value. Priority, 0 is highest
func (m *AppleUserInitiatedEnrollmentProfile) GetPriority()(*int32) {
    return m.priority
}
// Serialize serializes information the current object
func (m *AppleUserInitiatedEnrollmentProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetAvailableEnrollmentTypeOptions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAvailableEnrollmentTypeOptions()))
        for i, v := range m.GetAvailableEnrollmentTypeOptions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("availableEnrollmentTypeOptions", cast)
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
    if m.GetDefaultEnrollmentType() != nil {
        cast := (*m.GetDefaultEnrollmentType()).String()
        err = writer.WriteStringValue("defaultEnrollmentType", &cast)
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
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
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
    if m.GetPlatform() != nil {
        cast := (*m.GetPlatform()).String()
        err = writer.WriteStringValue("platform", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The list of assignments for this profile.
func (m *AppleUserInitiatedEnrollmentProfile) SetAssignments(value []AppleEnrollmentProfileAssignmentable)() {
    m.assignments = value
}
// SetAvailableEnrollmentTypeOptions sets the availableEnrollmentTypeOptions property value. List of available enrollment type options
func (m *AppleUserInitiatedEnrollmentProfile) SetAvailableEnrollmentTypeOptions(value []AppleOwnerTypeEnrollmentTypeable)() {
    m.availableEnrollmentTypeOptions = value
}
// SetCreatedDateTime sets the createdDateTime property value. Profile creation time
func (m *AppleUserInitiatedEnrollmentProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDefaultEnrollmentType sets the defaultEnrollmentType property value. The defaultEnrollmentType property
func (m *AppleUserInitiatedEnrollmentProfile) SetDefaultEnrollmentType(value *AppleUserInitiatedEnrollmentType)() {
    m.defaultEnrollmentType = value
}
// SetDescription sets the description property value. Description of the profile
func (m *AppleUserInitiatedEnrollmentProfile) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Name of the profile
func (m *AppleUserInitiatedEnrollmentProfile) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Profile last modified time
func (m *AppleUserInitiatedEnrollmentProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetPlatform sets the platform property value. Supported platform types.
func (m *AppleUserInitiatedEnrollmentProfile) SetPlatform(value *DevicePlatformType)() {
    m.platform = value
}
// SetPriority sets the priority property value. Priority, 0 is highest
func (m *AppleUserInitiatedEnrollmentProfile) SetPriority(value *int32)() {
    m.priority = value
}
