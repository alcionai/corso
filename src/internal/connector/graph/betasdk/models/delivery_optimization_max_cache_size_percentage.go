package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeliveryOptimizationMaxCacheSizePercentage 
type DeliveryOptimizationMaxCacheSizePercentage struct {
    DeliveryOptimizationMaxCacheSize
    // Specifies the maximum cache size that Delivery Optimization can utilize, as a percentage of disk size (1-100). Valid values 1 to 100
    maximumCacheSizePercentage *int32
}
// NewDeliveryOptimizationMaxCacheSizePercentage instantiates a new DeliveryOptimizationMaxCacheSizePercentage and sets the default values.
func NewDeliveryOptimizationMaxCacheSizePercentage()(*DeliveryOptimizationMaxCacheSizePercentage) {
    m := &DeliveryOptimizationMaxCacheSizePercentage{
        DeliveryOptimizationMaxCacheSize: *NewDeliveryOptimizationMaxCacheSize(),
    }
    odataTypeValue := "#microsoft.graph.deliveryOptimizationMaxCacheSizePercentage";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeliveryOptimizationMaxCacheSizePercentageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeliveryOptimizationMaxCacheSizePercentageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeliveryOptimizationMaxCacheSizePercentage(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeliveryOptimizationMaxCacheSizePercentage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeliveryOptimizationMaxCacheSize.GetFieldDeserializers()
    res["maximumCacheSizePercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumCacheSizePercentage(val)
        }
        return nil
    }
    return res
}
// GetMaximumCacheSizePercentage gets the maximumCacheSizePercentage property value. Specifies the maximum cache size that Delivery Optimization can utilize, as a percentage of disk size (1-100). Valid values 1 to 100
func (m *DeliveryOptimizationMaxCacheSizePercentage) GetMaximumCacheSizePercentage()(*int32) {
    return m.maximumCacheSizePercentage
}
// Serialize serializes information the current object
func (m *DeliveryOptimizationMaxCacheSizePercentage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeliveryOptimizationMaxCacheSize.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("maximumCacheSizePercentage", m.GetMaximumCacheSizePercentage())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMaximumCacheSizePercentage sets the maximumCacheSizePercentage property value. Specifies the maximum cache size that Delivery Optimization can utilize, as a percentage of disk size (1-100). Valid values 1 to 100
func (m *DeliveryOptimizationMaxCacheSizePercentage) SetMaximumCacheSizePercentage(value *int32)() {
    m.maximumCacheSizePercentage = value
}
