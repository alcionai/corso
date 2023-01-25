package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VpnServer vPN Server definition.
type VpnServer struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Address (IP address, FQDN or URL)
    address *string
    // Description.
    description *string
    // Default server.
    isDefaultServer *bool
    // The OdataType property
    odataType *string
}
// NewVpnServer instantiates a new vpnServer and sets the default values.
func NewVpnServer()(*VpnServer) {
    m := &VpnServer{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateVpnServerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVpnServerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVpnServer(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VpnServer) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAddress gets the address property value. Address (IP address, FQDN or URL)
func (m *VpnServer) GetAddress()(*string) {
    return m.address
}
// GetDescription gets the description property value. Description.
func (m *VpnServer) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VpnServer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["isDefaultServer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefaultServer(val)
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
// GetIsDefaultServer gets the isDefaultServer property value. Default server.
func (m *VpnServer) GetIsDefaultServer()(*bool) {
    return m.isDefaultServer
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *VpnServer) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *VpnServer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isDefaultServer", m.GetIsDefaultServer())
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
func (m *VpnServer) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAddress sets the address property value. Address (IP address, FQDN or URL)
func (m *VpnServer) SetAddress(value *string)() {
    m.address = value
}
// SetDescription sets the description property value. Description.
func (m *VpnServer) SetDescription(value *string)() {
    m.description = value
}
// SetIsDefaultServer sets the isDefaultServer property value. Default server.
func (m *VpnServer) SetIsDefaultServer(value *bool)() {
    m.isDefaultServer = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *VpnServer) SetOdataType(value *string)() {
    m.odataType = value
}
