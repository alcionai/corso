package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsPhone81StoreApp 
type WindowsPhone81StoreApp struct {
    MobileApp
    // The Windows Phone 8.1 app store URL.
    appStoreUrl *string
}
// NewWindowsPhone81StoreApp instantiates a new WindowsPhone81StoreApp and sets the default values.
func NewWindowsPhone81StoreApp()(*WindowsPhone81StoreApp) {
    m := &WindowsPhone81StoreApp{
        MobileApp: *NewMobileApp(),
    }
    odataTypeValue := "#microsoft.graph.windowsPhone81StoreApp";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsPhone81StoreAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsPhone81StoreAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsPhone81StoreApp(), nil
}
// GetAppStoreUrl gets the appStoreUrl property value. The Windows Phone 8.1 app store URL.
func (m *WindowsPhone81StoreApp) GetAppStoreUrl()(*string) {
    return m.appStoreUrl
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsPhone81StoreApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileApp.GetFieldDeserializers()
    res["appStoreUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppStoreUrl(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *WindowsPhone81StoreApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appStoreUrl", m.GetAppStoreUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppStoreUrl sets the appStoreUrl property value. The Windows Phone 8.1 app store URL.
func (m *WindowsPhone81StoreApp) SetAppStoreUrl(value *string)() {
    m.appStoreUrl = value
}
