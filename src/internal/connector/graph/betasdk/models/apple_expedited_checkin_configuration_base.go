package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppleExpeditedCheckinConfigurationBase 
type AppleExpeditedCheckinConfigurationBase struct {
    DeviceConfiguration
    // Gets or sets whether to enable expedited device check-ins.
    enableExpeditedCheckin *bool
}
// NewAppleExpeditedCheckinConfigurationBase instantiates a new AppleExpeditedCheckinConfigurationBase and sets the default values.
func NewAppleExpeditedCheckinConfigurationBase()(*AppleExpeditedCheckinConfigurationBase) {
    m := &AppleExpeditedCheckinConfigurationBase{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.appleExpeditedCheckinConfigurationBase";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAppleExpeditedCheckinConfigurationBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppleExpeditedCheckinConfigurationBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.iosExpeditedCheckinConfiguration":
                        return NewIosExpeditedCheckinConfiguration(), nil
                }
            }
        }
    }
    return NewAppleExpeditedCheckinConfigurationBase(), nil
}
// GetEnableExpeditedCheckin gets the enableExpeditedCheckin property value. Gets or sets whether to enable expedited device check-ins.
func (m *AppleExpeditedCheckinConfigurationBase) GetEnableExpeditedCheckin()(*bool) {
    return m.enableExpeditedCheckin
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppleExpeditedCheckinConfigurationBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["enableExpeditedCheckin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableExpeditedCheckin(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AppleExpeditedCheckinConfigurationBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("enableExpeditedCheckin", m.GetEnableExpeditedCheckin())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEnableExpeditedCheckin sets the enableExpeditedCheckin property value. Gets or sets whether to enable expedited device check-ins.
func (m *AppleExpeditedCheckinConfigurationBase) SetEnableExpeditedCheckin(value *bool)() {
    m.enableExpeditedCheckin = value
}
