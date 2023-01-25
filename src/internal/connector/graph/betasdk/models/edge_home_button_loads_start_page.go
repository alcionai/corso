package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdgeHomeButtonLoadsStartPage 
type EdgeHomeButtonLoadsStartPage struct {
    EdgeHomeButtonConfiguration
}
// NewEdgeHomeButtonLoadsStartPage instantiates a new EdgeHomeButtonLoadsStartPage and sets the default values.
func NewEdgeHomeButtonLoadsStartPage()(*EdgeHomeButtonLoadsStartPage) {
    m := &EdgeHomeButtonLoadsStartPage{
        EdgeHomeButtonConfiguration: *NewEdgeHomeButtonConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.edgeHomeButtonLoadsStartPage";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateEdgeHomeButtonLoadsStartPageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdgeHomeButtonLoadsStartPageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdgeHomeButtonLoadsStartPage(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdgeHomeButtonLoadsStartPage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EdgeHomeButtonConfiguration.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *EdgeHomeButtonLoadsStartPage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EdgeHomeButtonConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
