package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosiPadOSWebClip 
type IosiPadOSWebClip struct {
    MobileApp
    // The web app URL.
    appUrl *string
    // Whether or not to use managed browser. When true, the app will be required to be opened in an Intune-protected browser. When false, the app will not be required to be opened in an Intune-protected browser.
    useManagedBrowser *bool
}
// NewIosiPadOSWebClip instantiates a new IosiPadOSWebClip and sets the default values.
func NewIosiPadOSWebClip()(*IosiPadOSWebClip) {
    m := &IosiPadOSWebClip{
        MobileApp: *NewMobileApp(),
    }
    odataTypeValue := "#microsoft.graph.iosiPadOSWebClip";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateIosiPadOSWebClipFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosiPadOSWebClipFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosiPadOSWebClip(), nil
}
// GetAppUrl gets the appUrl property value. The web app URL.
func (m *IosiPadOSWebClip) GetAppUrl()(*string) {
    return m.appUrl
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosiPadOSWebClip) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileApp.GetFieldDeserializers()
    res["appUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppUrl(val)
        }
        return nil
    }
    res["useManagedBrowser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseManagedBrowser(val)
        }
        return nil
    }
    return res
}
// GetUseManagedBrowser gets the useManagedBrowser property value. Whether or not to use managed browser. When true, the app will be required to be opened in an Intune-protected browser. When false, the app will not be required to be opened in an Intune-protected browser.
func (m *IosiPadOSWebClip) GetUseManagedBrowser()(*bool) {
    return m.useManagedBrowser
}
// Serialize serializes information the current object
func (m *IosiPadOSWebClip) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appUrl", m.GetAppUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("useManagedBrowser", m.GetUseManagedBrowser())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppUrl sets the appUrl property value. The web app URL.
func (m *IosiPadOSWebClip) SetAppUrl(value *string)() {
    m.appUrl = value
}
// SetUseManagedBrowser sets the useManagedBrowser property value. Whether or not to use managed browser. When true, the app will be required to be opened in an Intune-protected browser. When false, the app will not be required to be opened in an Intune-protected browser.
func (m *IosiPadOSWebClip) SetUseManagedBrowser(value *bool)() {
    m.useManagedBrowser = value
}
