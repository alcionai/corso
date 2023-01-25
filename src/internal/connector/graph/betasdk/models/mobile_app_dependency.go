package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppDependency 
type MobileAppDependency struct {
    MobileAppRelationship
    // Indicates the dependency type associated with a relationship between two mobile apps.
    dependencyType *MobileAppDependencyType
    // The total number of apps that directly or indirectly depend on the parent app.
    dependentAppCount *int32
    // The total number of apps the child app directly or indirectly depends on.
    dependsOnAppCount *int32
}
// NewMobileAppDependency instantiates a new MobileAppDependency and sets the default values.
func NewMobileAppDependency()(*MobileAppDependency) {
    m := &MobileAppDependency{
        MobileAppRelationship: *NewMobileAppRelationship(),
    }
    odataTypeValue := "#microsoft.graph.mobileAppDependency";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMobileAppDependencyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppDependencyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppDependency(), nil
}
// GetDependencyType gets the dependencyType property value. Indicates the dependency type associated with a relationship between two mobile apps.
func (m *MobileAppDependency) GetDependencyType()(*MobileAppDependencyType) {
    return m.dependencyType
}
// GetDependentAppCount gets the dependentAppCount property value. The total number of apps that directly or indirectly depend on the parent app.
func (m *MobileAppDependency) GetDependentAppCount()(*int32) {
    return m.dependentAppCount
}
// GetDependsOnAppCount gets the dependsOnAppCount property value. The total number of apps the child app directly or indirectly depends on.
func (m *MobileAppDependency) GetDependsOnAppCount()(*int32) {
    return m.dependsOnAppCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppDependency) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileAppRelationship.GetFieldDeserializers()
    res["dependencyType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMobileAppDependencyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDependencyType(val.(*MobileAppDependencyType))
        }
        return nil
    }
    res["dependentAppCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDependentAppCount(val)
        }
        return nil
    }
    res["dependsOnAppCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDependsOnAppCount(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *MobileAppDependency) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileAppRelationship.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDependencyType() != nil {
        cast := (*m.GetDependencyType()).String()
        err = writer.WriteStringValue("dependencyType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("dependentAppCount", m.GetDependentAppCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("dependsOnAppCount", m.GetDependsOnAppCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDependencyType sets the dependencyType property value. Indicates the dependency type associated with a relationship between two mobile apps.
func (m *MobileAppDependency) SetDependencyType(value *MobileAppDependencyType)() {
    m.dependencyType = value
}
// SetDependentAppCount sets the dependentAppCount property value. The total number of apps that directly or indirectly depend on the parent app.
func (m *MobileAppDependency) SetDependentAppCount(value *int32)() {
    m.dependentAppCount = value
}
// SetDependsOnAppCount sets the dependsOnAppCount property value. The total number of apps the child app directly or indirectly depends on.
func (m *MobileAppDependency) SetDependsOnAppCount(value *int32)() {
    m.dependsOnAppCount = value
}
