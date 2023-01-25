package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerGlobalProxyDirect 
type AndroidDeviceOwnerGlobalProxyDirect struct {
    AndroidDeviceOwnerGlobalProxy
    // The excluded hosts
    excludedHosts []string
    // The host name
    host *string
    // The port
    port *int32
}
// NewAndroidDeviceOwnerGlobalProxyDirect instantiates a new AndroidDeviceOwnerGlobalProxyDirect and sets the default values.
func NewAndroidDeviceOwnerGlobalProxyDirect()(*AndroidDeviceOwnerGlobalProxyDirect) {
    m := &AndroidDeviceOwnerGlobalProxyDirect{
        AndroidDeviceOwnerGlobalProxy: *NewAndroidDeviceOwnerGlobalProxy(),
    }
    odataTypeValue := "#microsoft.graph.androidDeviceOwnerGlobalProxyDirect";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidDeviceOwnerGlobalProxyDirectFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerGlobalProxyDirectFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerGlobalProxyDirect(), nil
}
// GetExcludedHosts gets the excludedHosts property value. The excluded hosts
func (m *AndroidDeviceOwnerGlobalProxyDirect) GetExcludedHosts()([]string) {
    return m.excludedHosts
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerGlobalProxyDirect) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AndroidDeviceOwnerGlobalProxy.GetFieldDeserializers()
    res["excludedHosts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExcludedHosts(res)
        }
        return nil
    }
    res["host"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHost(val)
        }
        return nil
    }
    res["port"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPort(val)
        }
        return nil
    }
    return res
}
// GetHost gets the host property value. The host name
func (m *AndroidDeviceOwnerGlobalProxyDirect) GetHost()(*string) {
    return m.host
}
// GetPort gets the port property value. The port
func (m *AndroidDeviceOwnerGlobalProxyDirect) GetPort()(*int32) {
    return m.port
}
// Serialize serializes information the current object
func (m *AndroidDeviceOwnerGlobalProxyDirect) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AndroidDeviceOwnerGlobalProxy.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetExcludedHosts() != nil {
        err = writer.WriteCollectionOfStringValues("excludedHosts", m.GetExcludedHosts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("host", m.GetHost())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("port", m.GetPort())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExcludedHosts sets the excludedHosts property value. The excluded hosts
func (m *AndroidDeviceOwnerGlobalProxyDirect) SetExcludedHosts(value []string)() {
    m.excludedHosts = value
}
// SetHost sets the host property value. The host name
func (m *AndroidDeviceOwnerGlobalProxyDirect) SetHost(value *string)() {
    m.host = value
}
// SetPort sets the port property value. The port
func (m *AndroidDeviceOwnerGlobalProxyDirect) SetPort(value *int32)() {
    m.port = value
}
