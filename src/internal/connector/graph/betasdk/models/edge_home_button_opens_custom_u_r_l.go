package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdgeHomeButtonOpensCustomURL 
type EdgeHomeButtonOpensCustomURL struct {
    EdgeHomeButtonConfiguration
    // The specific URL to load.
    homeButtonCustomURL *string
}
// NewEdgeHomeButtonOpensCustomURL instantiates a new EdgeHomeButtonOpensCustomURL and sets the default values.
func NewEdgeHomeButtonOpensCustomURL()(*EdgeHomeButtonOpensCustomURL) {
    m := &EdgeHomeButtonOpensCustomURL{
        EdgeHomeButtonConfiguration: *NewEdgeHomeButtonConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.edgeHomeButtonOpensCustomURL";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateEdgeHomeButtonOpensCustomURLFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdgeHomeButtonOpensCustomURLFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdgeHomeButtonOpensCustomURL(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdgeHomeButtonOpensCustomURL) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EdgeHomeButtonConfiguration.GetFieldDeserializers()
    res["homeButtonCustomURL"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomeButtonCustomURL(val)
        }
        return nil
    }
    return res
}
// GetHomeButtonCustomURL gets the homeButtonCustomURL property value. The specific URL to load.
func (m *EdgeHomeButtonOpensCustomURL) GetHomeButtonCustomURL()(*string) {
    return m.homeButtonCustomURL
}
// Serialize serializes information the current object
func (m *EdgeHomeButtonOpensCustomURL) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EdgeHomeButtonConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("homeButtonCustomURL", m.GetHomeButtonCustomURL())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetHomeButtonCustomURL sets the homeButtonCustomURL property value. The specific URL to load.
func (m *EdgeHomeButtonOpensCustomURL) SetHomeButtonCustomURL(value *string)() {
    m.homeButtonCustomURL = value
}
