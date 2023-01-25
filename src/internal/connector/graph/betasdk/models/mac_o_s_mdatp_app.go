package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSMdatpApp 
type MacOSMdatpApp struct {
    MobileApp
}
// NewMacOSMdatpApp instantiates a new MacOSMdatpApp and sets the default values.
func NewMacOSMdatpApp()(*MacOSMdatpApp) {
    m := &MacOSMdatpApp{
        MobileApp: *NewMobileApp(),
    }
    odataTypeValue := "#microsoft.graph.macOSMdatpApp";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMacOSMdatpAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSMdatpAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSMdatpApp(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSMdatpApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileApp.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *MacOSMdatpApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileApp.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
