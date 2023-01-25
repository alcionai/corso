package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CrossTenantAccessPolicyTenantRestrictions 
type CrossTenantAccessPolicyTenantRestrictions struct {
    CrossTenantAccessPolicyB2BSetting
    // The devices property
    devices DevicesFilterable
}
// NewCrossTenantAccessPolicyTenantRestrictions instantiates a new CrossTenantAccessPolicyTenantRestrictions and sets the default values.
func NewCrossTenantAccessPolicyTenantRestrictions()(*CrossTenantAccessPolicyTenantRestrictions) {
    m := &CrossTenantAccessPolicyTenantRestrictions{
        CrossTenantAccessPolicyB2BSetting: *NewCrossTenantAccessPolicyB2BSetting(),
    }
    odataTypeValue := "#microsoft.graph.crossTenantAccessPolicyTenantRestrictions";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCrossTenantAccessPolicyTenantRestrictionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCrossTenantAccessPolicyTenantRestrictionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCrossTenantAccessPolicyTenantRestrictions(), nil
}
// GetDevices gets the devices property value. The devices property
func (m *CrossTenantAccessPolicyTenantRestrictions) GetDevices()(DevicesFilterable) {
    return m.devices
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CrossTenantAccessPolicyTenantRestrictions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CrossTenantAccessPolicyB2BSetting.GetFieldDeserializers()
    res["devices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDevicesFilterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevices(val.(DevicesFilterable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CrossTenantAccessPolicyTenantRestrictions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CrossTenantAccessPolicyB2BSetting.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("devices", m.GetDevices())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDevices sets the devices property value. The devices property
func (m *CrossTenantAccessPolicyTenantRestrictions) SetDevices(value DevicesFilterable)() {
    m.devices = value
}
