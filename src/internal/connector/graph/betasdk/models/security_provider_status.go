package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityProviderStatus 
type SecurityProviderStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The enabled property
    enabled *bool
    // The endpoint property
    endpoint *string
    // The OdataType property
    odataType *string
    // The provider property
    provider *string
    // The region property
    region *string
    // The vendor property
    vendor_escaped *string
}
// NewSecurityProviderStatus instantiates a new securityProviderStatus and sets the default values.
func NewSecurityProviderStatus()(*SecurityProviderStatus) {
    m := &SecurityProviderStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSecurityProviderStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityProviderStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityProviderStatus(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecurityProviderStatus) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEnabled gets the enabled property value. The enabled property
func (m *SecurityProviderStatus) GetEnabled()(*bool) {
    return m.enabled
}
// GetEndpoint gets the endpoint property value. The endpoint property
func (m *SecurityProviderStatus) GetEndpoint()(*string) {
    return m.endpoint
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecurityProviderStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["endpoint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndpoint(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["provider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvider(val)
        }
        return nil
    }
    res["region"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegion(val)
        }
        return nil
    }
    res["vendor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendor(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SecurityProviderStatus) GetOdataType()(*string) {
    return m.odataType
}
// GetProvider gets the provider property value. The provider property
func (m *SecurityProviderStatus) GetProvider()(*string) {
    return m.provider
}
// GetRegion gets the region property value. The region property
func (m *SecurityProviderStatus) GetRegion()(*string) {
    return m.region
}
// GetVendor gets the vendor property value. The vendor property
func (m *SecurityProviderStatus) GetVendor()(*string) {
    return m.vendor_escaped
}
// Serialize serializes information the current object
func (m *SecurityProviderStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("endpoint", m.GetEndpoint())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("provider", m.GetProvider())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("region", m.GetRegion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("vendor", m.GetVendor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecurityProviderStatus) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEnabled sets the enabled property value. The enabled property
func (m *SecurityProviderStatus) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetEndpoint sets the endpoint property value. The endpoint property
func (m *SecurityProviderStatus) SetEndpoint(value *string)() {
    m.endpoint = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SecurityProviderStatus) SetOdataType(value *string)() {
    m.odataType = value
}
// SetProvider sets the provider property value. The provider property
func (m *SecurityProviderStatus) SetProvider(value *string)() {
    m.provider = value
}
// SetRegion sets the region property value. The region property
func (m *SecurityProviderStatus) SetRegion(value *string)() {
    m.region = value
}
// SetVendor sets the vendor property value. The vendor property
func (m *SecurityProviderStatus) SetVendor(value *string)() {
    m.vendor_escaped = value
}
