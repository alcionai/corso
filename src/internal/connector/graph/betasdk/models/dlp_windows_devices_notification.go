package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DlpWindowsDevicesNotification 
type DlpWindowsDevicesNotification struct {
    DlpNotification
    // The contentName property
    contentName *string
    // The lastModfiedBy property
    lastModfiedBy *string
}
// NewDlpWindowsDevicesNotification instantiates a new DlpWindowsDevicesNotification and sets the default values.
func NewDlpWindowsDevicesNotification()(*DlpWindowsDevicesNotification) {
    m := &DlpWindowsDevicesNotification{
        DlpNotification: *NewDlpNotification(),
    }
    odataTypeValue := "#microsoft.graph.dlpWindowsDevicesNotification";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDlpWindowsDevicesNotificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDlpWindowsDevicesNotificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDlpWindowsDevicesNotification(), nil
}
// GetContentName gets the contentName property value. The contentName property
func (m *DlpWindowsDevicesNotification) GetContentName()(*string) {
    return m.contentName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DlpWindowsDevicesNotification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DlpNotification.GetFieldDeserializers()
    res["contentName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentName(val)
        }
        return nil
    }
    res["lastModfiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModfiedBy(val)
        }
        return nil
    }
    return res
}
// GetLastModfiedBy gets the lastModfiedBy property value. The lastModfiedBy property
func (m *DlpWindowsDevicesNotification) GetLastModfiedBy()(*string) {
    return m.lastModfiedBy
}
// Serialize serializes information the current object
func (m *DlpWindowsDevicesNotification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DlpNotification.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("contentName", m.GetContentName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastModfiedBy", m.GetLastModfiedBy())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContentName sets the contentName property value. The contentName property
func (m *DlpWindowsDevicesNotification) SetContentName(value *string)() {
    m.contentName = value
}
// SetLastModfiedBy sets the lastModfiedBy property value. The lastModfiedBy property
func (m *DlpWindowsDevicesNotification) SetLastModfiedBy(value *string)() {
    m.lastModfiedBy = value
}
