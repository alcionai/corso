package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceEnrollmentPlatformRestrictionConfiguration 
type DeviceEnrollmentPlatformRestrictionConfiguration struct {
    DeviceEnrollmentConfiguration
    // Restrictions based on platform, platform operating system version, and device ownership
    platformRestriction DeviceEnrollmentPlatformRestrictionable
    // This enum indicates the platform type for which the enrollment restriction applies.
    platformType *EnrollmentRestrictionPlatformType
}
// NewDeviceEnrollmentPlatformRestrictionConfiguration instantiates a new DeviceEnrollmentPlatformRestrictionConfiguration and sets the default values.
func NewDeviceEnrollmentPlatformRestrictionConfiguration()(*DeviceEnrollmentPlatformRestrictionConfiguration) {
    m := &DeviceEnrollmentPlatformRestrictionConfiguration{
        DeviceEnrollmentConfiguration: *NewDeviceEnrollmentConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.deviceEnrollmentPlatformRestrictionConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceEnrollmentPlatformRestrictionConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceEnrollmentPlatformRestrictionConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceEnrollmentPlatformRestrictionConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceEnrollmentPlatformRestrictionConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceEnrollmentConfiguration.GetFieldDeserializers()
    res["platformRestriction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceEnrollmentPlatformRestrictionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatformRestriction(val.(DeviceEnrollmentPlatformRestrictionable))
        }
        return nil
    }
    res["platformType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnrollmentRestrictionPlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatformType(val.(*EnrollmentRestrictionPlatformType))
        }
        return nil
    }
    return res
}
// GetPlatformRestriction gets the platformRestriction property value. Restrictions based on platform, platform operating system version, and device ownership
func (m *DeviceEnrollmentPlatformRestrictionConfiguration) GetPlatformRestriction()(DeviceEnrollmentPlatformRestrictionable) {
    return m.platformRestriction
}
// GetPlatformType gets the platformType property value. This enum indicates the platform type for which the enrollment restriction applies.
func (m *DeviceEnrollmentPlatformRestrictionConfiguration) GetPlatformType()(*EnrollmentRestrictionPlatformType) {
    return m.platformType
}
// Serialize serializes information the current object
func (m *DeviceEnrollmentPlatformRestrictionConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceEnrollmentConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("platformRestriction", m.GetPlatformRestriction())
        if err != nil {
            return err
        }
    }
    if m.GetPlatformType() != nil {
        cast := (*m.GetPlatformType()).String()
        err = writer.WriteStringValue("platformType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPlatformRestriction sets the platformRestriction property value. Restrictions based on platform, platform operating system version, and device ownership
func (m *DeviceEnrollmentPlatformRestrictionConfiguration) SetPlatformRestriction(value DeviceEnrollmentPlatformRestrictionable)() {
    m.platformRestriction = value
}
// SetPlatformType sets the platformType property value. This enum indicates the platform type for which the enrollment restriction applies.
func (m *DeviceEnrollmentPlatformRestrictionConfiguration) SetPlatformType(value *EnrollmentRestrictionPlatformType)() {
    m.platformType = value
}
