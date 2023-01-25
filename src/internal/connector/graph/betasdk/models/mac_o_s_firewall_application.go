package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSFirewallApplication represents an app in the list of macOS firewall applications
type MacOSFirewallApplication struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Whether or not incoming connections are allowed.
    allowsIncomingConnections *bool
    // BundleId of the application.
    bundleId *string
    // The OdataType property
    odataType *string
}
// NewMacOSFirewallApplication instantiates a new macOSFirewallApplication and sets the default values.
func NewMacOSFirewallApplication()(*MacOSFirewallApplication) {
    m := &MacOSFirewallApplication{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMacOSFirewallApplicationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSFirewallApplicationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSFirewallApplication(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MacOSFirewallApplication) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAllowsIncomingConnections gets the allowsIncomingConnections property value. Whether or not incoming connections are allowed.
func (m *MacOSFirewallApplication) GetAllowsIncomingConnections()(*bool) {
    return m.allowsIncomingConnections
}
// GetBundleId gets the bundleId property value. BundleId of the application.
func (m *MacOSFirewallApplication) GetBundleId()(*string) {
    return m.bundleId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSFirewallApplication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowsIncomingConnections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowsIncomingConnections(val)
        }
        return nil
    }
    res["bundleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBundleId(val)
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
func (m *MacOSFirewallApplication) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *MacOSFirewallApplication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowsIncomingConnections", m.GetAllowsIncomingConnections())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("bundleId", m.GetBundleId())
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
func (m *MacOSFirewallApplication) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAllowsIncomingConnections sets the allowsIncomingConnections property value. Whether or not incoming connections are allowed.
func (m *MacOSFirewallApplication) SetAllowsIncomingConnections(value *bool)() {
    m.allowsIncomingConnections = value
}
// SetBundleId sets the bundleId property value. BundleId of the application.
func (m *MacOSFirewallApplication) SetBundleId(value *string)() {
    m.bundleId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MacOSFirewallApplication) SetOdataType(value *string)() {
    m.odataType = value
}
