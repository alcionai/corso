package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppleDeviceFeaturesConfigurationBase 
type AppleDeviceFeaturesConfigurationBase struct {
    DeviceConfiguration
    // An array of AirPrint printers that should always be shown. This collection can contain a maximum of 500 elements.
    airPrintDestinations []AirPrintDestinationable
}
// NewAppleDeviceFeaturesConfigurationBase instantiates a new AppleDeviceFeaturesConfigurationBase and sets the default values.
func NewAppleDeviceFeaturesConfigurationBase()(*AppleDeviceFeaturesConfigurationBase) {
    m := &AppleDeviceFeaturesConfigurationBase{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.appleDeviceFeaturesConfigurationBase";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAppleDeviceFeaturesConfigurationBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppleDeviceFeaturesConfigurationBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.iosDeviceFeaturesConfiguration":
                        return NewIosDeviceFeaturesConfiguration(), nil
                    case "#microsoft.graph.macOSDeviceFeaturesConfiguration":
                        return NewMacOSDeviceFeaturesConfiguration(), nil
                }
            }
        }
    }
    return NewAppleDeviceFeaturesConfigurationBase(), nil
}
// GetAirPrintDestinations gets the airPrintDestinations property value. An array of AirPrint printers that should always be shown. This collection can contain a maximum of 500 elements.
func (m *AppleDeviceFeaturesConfigurationBase) GetAirPrintDestinations()([]AirPrintDestinationable) {
    return m.airPrintDestinations
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppleDeviceFeaturesConfigurationBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["airPrintDestinations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAirPrintDestinationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AirPrintDestinationable, len(val))
            for i, v := range val {
                res[i] = v.(AirPrintDestinationable)
            }
            m.SetAirPrintDestinations(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AppleDeviceFeaturesConfigurationBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAirPrintDestinations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAirPrintDestinations()))
        for i, v := range m.GetAirPrintDestinations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("airPrintDestinations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAirPrintDestinations sets the airPrintDestinations property value. An array of AirPrint printers that should always be shown. This collection can contain a maximum of 500 elements.
func (m *AppleDeviceFeaturesConfigurationBase) SetAirPrintDestinations(value []AirPrintDestinationable)() {
    m.airPrintDestinations = value
}
