package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SpecifiedCaptiveNetworkPlugins specifies all the Captive network plugins allowed during the IKEv2 AlwaysOn VPN connection
type SpecifiedCaptiveNetworkPlugins struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Address of the IKEv2 server. Must be a FQDN, UserFQDN, network address, or ASN1DN
    allowedBundleIdentifiers []string
    // The OdataType property
    odataType *string
}
// NewSpecifiedCaptiveNetworkPlugins instantiates a new specifiedCaptiveNetworkPlugins and sets the default values.
func NewSpecifiedCaptiveNetworkPlugins()(*SpecifiedCaptiveNetworkPlugins) {
    m := &SpecifiedCaptiveNetworkPlugins{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSpecifiedCaptiveNetworkPluginsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSpecifiedCaptiveNetworkPluginsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSpecifiedCaptiveNetworkPlugins(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SpecifiedCaptiveNetworkPlugins) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAllowedBundleIdentifiers gets the allowedBundleIdentifiers property value. Address of the IKEv2 server. Must be a FQDN, UserFQDN, network address, or ASN1DN
func (m *SpecifiedCaptiveNetworkPlugins) GetAllowedBundleIdentifiers()([]string) {
    return m.allowedBundleIdentifiers
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SpecifiedCaptiveNetworkPlugins) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowedBundleIdentifiers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAllowedBundleIdentifiers(res)
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
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SpecifiedCaptiveNetworkPlugins) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *SpecifiedCaptiveNetworkPlugins) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAllowedBundleIdentifiers() != nil {
        err := writer.WriteCollectionOfStringValues("allowedBundleIdentifiers", m.GetAllowedBundleIdentifiers())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SpecifiedCaptiveNetworkPlugins) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAllowedBundleIdentifiers sets the allowedBundleIdentifiers property value. Address of the IKEv2 server. Must be a FQDN, UserFQDN, network address, or ASN1DN
func (m *SpecifiedCaptiveNetworkPlugins) SetAllowedBundleIdentifiers(value []string)() {
    m.allowedBundleIdentifiers = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SpecifiedCaptiveNetworkPlugins) SetOdataType(value *string)() {
    m.odataType = value
}
