package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerKioskModeWeblink 
type AndroidDeviceOwnerKioskModeWeblink struct {
    AndroidDeviceOwnerKioskModeFolderItem
    // Display name for weblink
    label *string
    // Link for weblink
    link *string
}
// NewAndroidDeviceOwnerKioskModeWeblink instantiates a new AndroidDeviceOwnerKioskModeWeblink and sets the default values.
func NewAndroidDeviceOwnerKioskModeWeblink()(*AndroidDeviceOwnerKioskModeWeblink) {
    m := &AndroidDeviceOwnerKioskModeWeblink{
        AndroidDeviceOwnerKioskModeFolderItem: *NewAndroidDeviceOwnerKioskModeFolderItem(),
    }
    odataTypeValue := "#microsoft.graph.androidDeviceOwnerKioskModeWeblink";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidDeviceOwnerKioskModeWeblinkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerKioskModeWeblinkFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerKioskModeWeblink(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerKioskModeWeblink) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AndroidDeviceOwnerKioskModeFolderItem.GetFieldDeserializers()
    res["label"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabel(val)
        }
        return nil
    }
    res["link"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLink(val)
        }
        return nil
    }
    return res
}
// GetLabel gets the label property value. Display name for weblink
func (m *AndroidDeviceOwnerKioskModeWeblink) GetLabel()(*string) {
    return m.label
}
// GetLink gets the link property value. Link for weblink
func (m *AndroidDeviceOwnerKioskModeWeblink) GetLink()(*string) {
    return m.link
}
// Serialize serializes information the current object
func (m *AndroidDeviceOwnerKioskModeWeblink) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AndroidDeviceOwnerKioskModeFolderItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("label", m.GetLabel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("link", m.GetLink())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLabel sets the label property value. Display name for weblink
func (m *AndroidDeviceOwnerKioskModeWeblink) SetLabel(value *string)() {
    m.label = value
}
// SetLink sets the link property value. Link for weblink
func (m *AndroidDeviceOwnerKioskModeWeblink) SetLink(value *string)() {
    m.link = value
}
