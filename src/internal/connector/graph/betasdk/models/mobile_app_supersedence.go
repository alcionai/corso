package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppSupersedence 
type MobileAppSupersedence struct {
    MobileAppRelationship
    // The total number of apps directly or indirectly superseded by the child app.
    supersededAppCount *int32
    // Indicates the supersedence type associated with a relationship between two mobile apps.
    supersedenceType *MobileAppSupersedenceType
    // The total number of apps directly or indirectly superseding the parent app.
    supersedingAppCount *int32
}
// NewMobileAppSupersedence instantiates a new MobileAppSupersedence and sets the default values.
func NewMobileAppSupersedence()(*MobileAppSupersedence) {
    m := &MobileAppSupersedence{
        MobileAppRelationship: *NewMobileAppRelationship(),
    }
    odataTypeValue := "#microsoft.graph.mobileAppSupersedence";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMobileAppSupersedenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppSupersedenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppSupersedence(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppSupersedence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileAppRelationship.GetFieldDeserializers()
    res["supersededAppCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupersededAppCount(val)
        }
        return nil
    }
    res["supersedenceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMobileAppSupersedenceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupersedenceType(val.(*MobileAppSupersedenceType))
        }
        return nil
    }
    res["supersedingAppCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupersedingAppCount(val)
        }
        return nil
    }
    return res
}
// GetSupersededAppCount gets the supersededAppCount property value. The total number of apps directly or indirectly superseded by the child app.
func (m *MobileAppSupersedence) GetSupersededAppCount()(*int32) {
    return m.supersededAppCount
}
// GetSupersedenceType gets the supersedenceType property value. Indicates the supersedence type associated with a relationship between two mobile apps.
func (m *MobileAppSupersedence) GetSupersedenceType()(*MobileAppSupersedenceType) {
    return m.supersedenceType
}
// GetSupersedingAppCount gets the supersedingAppCount property value. The total number of apps directly or indirectly superseding the parent app.
func (m *MobileAppSupersedence) GetSupersedingAppCount()(*int32) {
    return m.supersedingAppCount
}
// Serialize serializes information the current object
func (m *MobileAppSupersedence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileAppRelationship.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("supersededAppCount", m.GetSupersededAppCount())
        if err != nil {
            return err
        }
    }
    if m.GetSupersedenceType() != nil {
        cast := (*m.GetSupersedenceType()).String()
        err = writer.WriteStringValue("supersedenceType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("supersedingAppCount", m.GetSupersedingAppCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSupersededAppCount sets the supersededAppCount property value. The total number of apps directly or indirectly superseded by the child app.
func (m *MobileAppSupersedence) SetSupersededAppCount(value *int32)() {
    m.supersededAppCount = value
}
// SetSupersedenceType sets the supersedenceType property value. Indicates the supersedence type associated with a relationship between two mobile apps.
func (m *MobileAppSupersedence) SetSupersedenceType(value *MobileAppSupersedenceType)() {
    m.supersedenceType = value
}
// SetSupersedingAppCount sets the supersedingAppCount property value. The total number of apps directly or indirectly superseding the parent app.
func (m *MobileAppSupersedence) SetSupersedingAppCount(value *int32)() {
    m.supersedingAppCount = value
}
