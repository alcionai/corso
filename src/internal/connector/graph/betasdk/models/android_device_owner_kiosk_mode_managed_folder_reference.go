package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerKioskModeManagedFolderReference 
type AndroidDeviceOwnerKioskModeManagedFolderReference struct {
    AndroidDeviceOwnerKioskModeHomeScreenItem
    // Unique identifier for the folder
    folderIdentifier *string
    // Name of the folder
    folderName *string
}
// NewAndroidDeviceOwnerKioskModeManagedFolderReference instantiates a new AndroidDeviceOwnerKioskModeManagedFolderReference and sets the default values.
func NewAndroidDeviceOwnerKioskModeManagedFolderReference()(*AndroidDeviceOwnerKioskModeManagedFolderReference) {
    m := &AndroidDeviceOwnerKioskModeManagedFolderReference{
        AndroidDeviceOwnerKioskModeHomeScreenItem: *NewAndroidDeviceOwnerKioskModeHomeScreenItem(),
    }
    odataTypeValue := "#microsoft.graph.androidDeviceOwnerKioskModeManagedFolderReference";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidDeviceOwnerKioskModeManagedFolderReferenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerKioskModeManagedFolderReferenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerKioskModeManagedFolderReference(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerKioskModeManagedFolderReference) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AndroidDeviceOwnerKioskModeHomeScreenItem.GetFieldDeserializers()
    res["folderIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFolderIdentifier(val)
        }
        return nil
    }
    res["folderName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFolderName(val)
        }
        return nil
    }
    return res
}
// GetFolderIdentifier gets the folderIdentifier property value. Unique identifier for the folder
func (m *AndroidDeviceOwnerKioskModeManagedFolderReference) GetFolderIdentifier()(*string) {
    return m.folderIdentifier
}
// GetFolderName gets the folderName property value. Name of the folder
func (m *AndroidDeviceOwnerKioskModeManagedFolderReference) GetFolderName()(*string) {
    return m.folderName
}
// Serialize serializes information the current object
func (m *AndroidDeviceOwnerKioskModeManagedFolderReference) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AndroidDeviceOwnerKioskModeHomeScreenItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("folderIdentifier", m.GetFolderIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("folderName", m.GetFolderName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFolderIdentifier sets the folderIdentifier property value. Unique identifier for the folder
func (m *AndroidDeviceOwnerKioskModeManagedFolderReference) SetFolderIdentifier(value *string)() {
    m.folderIdentifier = value
}
// SetFolderName sets the folderName property value. Name of the folder
func (m *AndroidDeviceOwnerKioskModeManagedFolderReference) SetFolderName(value *string)() {
    m.folderName = value
}
