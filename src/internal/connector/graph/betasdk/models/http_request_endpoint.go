package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// HttpRequestEndpoint 
type HttpRequestEndpoint struct {
    CustomExtensionEndpointConfiguration
    // The targetUrl property
    targetUrl *string
}
// NewHttpRequestEndpoint instantiates a new HttpRequestEndpoint and sets the default values.
func NewHttpRequestEndpoint()(*HttpRequestEndpoint) {
    m := &HttpRequestEndpoint{
        CustomExtensionEndpointConfiguration: *NewCustomExtensionEndpointConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.httpRequestEndpoint";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateHttpRequestEndpointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateHttpRequestEndpointFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHttpRequestEndpoint(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *HttpRequestEndpoint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomExtensionEndpointConfiguration.GetFieldDeserializers()
    res["targetUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetUrl(val)
        }
        return nil
    }
    return res
}
// GetTargetUrl gets the targetUrl property value. The targetUrl property
func (m *HttpRequestEndpoint) GetTargetUrl()(*string) {
    return m.targetUrl
}
// Serialize serializes information the current object
func (m *HttpRequestEndpoint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomExtensionEndpointConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("targetUrl", m.GetTargetUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTargetUrl sets the targetUrl property value. The targetUrl property
func (m *HttpRequestEndpoint) SetTargetUrl(value *string)() {
    m.targetUrl = value
}
