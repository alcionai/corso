package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidManagedStoreWebApp 
type AndroidManagedStoreWebApp struct {
    AndroidManagedStoreApp
}
// NewAndroidManagedStoreWebApp instantiates a new AndroidManagedStoreWebApp and sets the default values.
func NewAndroidManagedStoreWebApp()(*AndroidManagedStoreWebApp) {
    m := &AndroidManagedStoreWebApp{
        AndroidManagedStoreApp: *NewAndroidManagedStoreApp(),
    }
    odataTypeValue := "#microsoft.graph.androidManagedStoreWebApp";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidManagedStoreWebAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidManagedStoreWebAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidManagedStoreWebApp(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidManagedStoreWebApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AndroidManagedStoreApp.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AndroidManagedStoreWebApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AndroidManagedStoreApp.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
