package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationSynchronizationProfile 
type EducationSynchronizationProfile struct {
    Entity
    // The dataProvider property
    dataProvider EducationSynchronizationDataProviderable
    // Name of the configuration profile for syncing identities.
    displayName *string
    // All errors associated with this synchronization profile.
    errors []EducationSynchronizationErrorable
    // The date the profile should be considered expired and cease syncing. Provide the date in YYYY-MM-DD format, following ISO 8601. Maximum value is 18 months from profile creation.  (optional)
    expirationDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // Determines if School Data Sync should automatically replace unsupported special characters while syncing from source.
    handleSpecialCharacterConstraint *bool
    // The identitySynchronizationConfiguration property
    identitySynchronizationConfiguration EducationIdentitySynchronizationConfigurationable
    // License setup configuration.
    licensesToAssign []EducationSynchronizationLicenseAssignmentable
    // The synchronization status.
    profileStatus EducationSynchronizationProfileStatusable
    // The state of the profile. Possible values are: provisioning, provisioned, provisioningFailed, deleting, deletionFailed.
    state *EducationSynchronizationProfileState
}
// NewEducationSynchronizationProfile instantiates a new EducationSynchronizationProfile and sets the default values.
func NewEducationSynchronizationProfile()(*EducationSynchronizationProfile) {
    m := &EducationSynchronizationProfile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEducationSynchronizationProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationSynchronizationProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationSynchronizationProfile(), nil
}
// GetDataProvider gets the dataProvider property value. The dataProvider property
func (m *EducationSynchronizationProfile) GetDataProvider()(EducationSynchronizationDataProviderable) {
    return m.dataProvider
}
// GetDisplayName gets the displayName property value. Name of the configuration profile for syncing identities.
func (m *EducationSynchronizationProfile) GetDisplayName()(*string) {
    return m.displayName
}
// GetErrors gets the errors property value. All errors associated with this synchronization profile.
func (m *EducationSynchronizationProfile) GetErrors()([]EducationSynchronizationErrorable) {
    return m.errors
}
// GetExpirationDate gets the expirationDate property value. The date the profile should be considered expired and cease syncing. Provide the date in YYYY-MM-DD format, following ISO 8601. Maximum value is 18 months from profile creation.  (optional)
func (m *EducationSynchronizationProfile) GetExpirationDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.expirationDate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationSynchronizationProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["dataProvider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationSynchronizationDataProviderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataProvider(val.(EducationSynchronizationDataProviderable))
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
    res["errors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationSynchronizationErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationSynchronizationErrorable, len(val))
            for i, v := range val {
                res[i] = v.(EducationSynchronizationErrorable)
            }
            m.SetErrors(res)
        }
        return nil
    }
    res["expirationDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDate(val)
        }
        return nil
    }
    res["handleSpecialCharacterConstraint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHandleSpecialCharacterConstraint(val)
        }
        return nil
    }
    res["identitySynchronizationConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationIdentitySynchronizationConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentitySynchronizationConfiguration(val.(EducationIdentitySynchronizationConfigurationable))
        }
        return nil
    }
    res["licensesToAssign"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationSynchronizationLicenseAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationSynchronizationLicenseAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(EducationSynchronizationLicenseAssignmentable)
            }
            m.SetLicensesToAssign(res)
        }
        return nil
    }
    res["profileStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationSynchronizationProfileStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfileStatus(val.(EducationSynchronizationProfileStatusable))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationSynchronizationProfileState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*EducationSynchronizationProfileState))
        }
        return nil
    }
    return res
}
// GetHandleSpecialCharacterConstraint gets the handleSpecialCharacterConstraint property value. Determines if School Data Sync should automatically replace unsupported special characters while syncing from source.
func (m *EducationSynchronizationProfile) GetHandleSpecialCharacterConstraint()(*bool) {
    return m.handleSpecialCharacterConstraint
}
// GetIdentitySynchronizationConfiguration gets the identitySynchronizationConfiguration property value. The identitySynchronizationConfiguration property
func (m *EducationSynchronizationProfile) GetIdentitySynchronizationConfiguration()(EducationIdentitySynchronizationConfigurationable) {
    return m.identitySynchronizationConfiguration
}
// GetLicensesToAssign gets the licensesToAssign property value. License setup configuration.
func (m *EducationSynchronizationProfile) GetLicensesToAssign()([]EducationSynchronizationLicenseAssignmentable) {
    return m.licensesToAssign
}
// GetProfileStatus gets the profileStatus property value. The synchronization status.
func (m *EducationSynchronizationProfile) GetProfileStatus()(EducationSynchronizationProfileStatusable) {
    return m.profileStatus
}
// GetState gets the state property value. The state of the profile. Possible values are: provisioning, provisioned, provisioningFailed, deleting, deletionFailed.
func (m *EducationSynchronizationProfile) GetState()(*EducationSynchronizationProfileState) {
    return m.state
}
// Serialize serializes information the current object
func (m *EducationSynchronizationProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("dataProvider", m.GetDataProvider())
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
    if m.GetErrors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetErrors()))
        for i, v := range m.GetErrors() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("errors", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("expirationDate", m.GetExpirationDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("handleSpecialCharacterConstraint", m.GetHandleSpecialCharacterConstraint())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("identitySynchronizationConfiguration", m.GetIdentitySynchronizationConfiguration())
        if err != nil {
            return err
        }
    }
    if m.GetLicensesToAssign() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLicensesToAssign()))
        for i, v := range m.GetLicensesToAssign() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("licensesToAssign", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("profileStatus", m.GetProfileStatus())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDataProvider sets the dataProvider property value. The dataProvider property
func (m *EducationSynchronizationProfile) SetDataProvider(value EducationSynchronizationDataProviderable)() {
    m.dataProvider = value
}
// SetDisplayName sets the displayName property value. Name of the configuration profile for syncing identities.
func (m *EducationSynchronizationProfile) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetErrors sets the errors property value. All errors associated with this synchronization profile.
func (m *EducationSynchronizationProfile) SetErrors(value []EducationSynchronizationErrorable)() {
    m.errors = value
}
// SetExpirationDate sets the expirationDate property value. The date the profile should be considered expired and cease syncing. Provide the date in YYYY-MM-DD format, following ISO 8601. Maximum value is 18 months from profile creation.  (optional)
func (m *EducationSynchronizationProfile) SetExpirationDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.expirationDate = value
}
// SetHandleSpecialCharacterConstraint sets the handleSpecialCharacterConstraint property value. Determines if School Data Sync should automatically replace unsupported special characters while syncing from source.
func (m *EducationSynchronizationProfile) SetHandleSpecialCharacterConstraint(value *bool)() {
    m.handleSpecialCharacterConstraint = value
}
// SetIdentitySynchronizationConfiguration sets the identitySynchronizationConfiguration property value. The identitySynchronizationConfiguration property
func (m *EducationSynchronizationProfile) SetIdentitySynchronizationConfiguration(value EducationIdentitySynchronizationConfigurationable)() {
    m.identitySynchronizationConfiguration = value
}
// SetLicensesToAssign sets the licensesToAssign property value. License setup configuration.
func (m *EducationSynchronizationProfile) SetLicensesToAssign(value []EducationSynchronizationLicenseAssignmentable)() {
    m.licensesToAssign = value
}
// SetProfileStatus sets the profileStatus property value. The synchronization status.
func (m *EducationSynchronizationProfile) SetProfileStatus(value EducationSynchronizationProfileStatusable)() {
    m.profileStatus = value
}
// SetState sets the state property value. The state of the profile. Possible values are: provisioning, provisioned, provisioningFailed, deleting, deletionFailed.
func (m *EducationSynchronizationProfile) SetState(value *EducationSynchronizationProfileState)() {
    m.state = value
}
