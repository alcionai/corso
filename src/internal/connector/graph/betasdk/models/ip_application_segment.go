package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IpApplicationSegment 
type IpApplicationSegment struct {
    ApplicationSegment
    // The destinationHost property
    destinationHost *string
    // The port property
    port *int32
}
// NewIpApplicationSegment instantiates a new IpApplicationSegment and sets the default values.
func NewIpApplicationSegment()(*IpApplicationSegment) {
    m := &IpApplicationSegment{
        ApplicationSegment: *NewApplicationSegment(),
    }
    odataTypeValue := "#microsoft.graph.ipApplicationSegment";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateIpApplicationSegmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIpApplicationSegmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIpApplicationSegment(), nil
}
// GetDestinationHost gets the destinationHost property value. The destinationHost property
func (m *IpApplicationSegment) GetDestinationHost()(*string) {
    return m.destinationHost
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IpApplicationSegment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ApplicationSegment.GetFieldDeserializers()
    res["destinationHost"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationHost(val)
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
// GetPort gets the port property value. The port property
func (m *IpApplicationSegment) GetPort()(*int32) {
    return m.port
}
// Serialize serializes information the current object
func (m *IpApplicationSegment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ApplicationSegment.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("destinationHost", m.GetDestinationHost())
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
// SetDestinationHost sets the destinationHost property value. The destinationHost property
func (m *IpApplicationSegment) SetDestinationHost(value *string)() {
    m.destinationHost = value
}
// SetPort sets the port property value. The port property
func (m *IpApplicationSegment) SetPort(value *int32)() {
    m.port = value
}
