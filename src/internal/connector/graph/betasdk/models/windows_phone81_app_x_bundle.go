package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsPhone81AppXBundle 
type WindowsPhone81AppXBundle struct {
    WindowsPhone81AppX
    // The list of AppX Package Information.
    appXPackageInformationList []WindowsPackageInformationable
}
// NewWindowsPhone81AppXBundle instantiates a new WindowsPhone81AppXBundle and sets the default values.
func NewWindowsPhone81AppXBundle()(*WindowsPhone81AppXBundle) {
    m := &WindowsPhone81AppXBundle{
        WindowsPhone81AppX: *NewWindowsPhone81AppX(),
    }
    odataTypeValue := "#microsoft.graph.windowsPhone81AppXBundle";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsPhone81AppXBundleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsPhone81AppXBundleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsPhone81AppXBundle(), nil
}
// GetAppXPackageInformationList gets the appXPackageInformationList property value. The list of AppX Package Information.
func (m *WindowsPhone81AppXBundle) GetAppXPackageInformationList()([]WindowsPackageInformationable) {
    return m.appXPackageInformationList
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsPhone81AppXBundle) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsPhone81AppX.GetFieldDeserializers()
    res["appXPackageInformationList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsPackageInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsPackageInformationable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsPackageInformationable)
            }
            m.SetAppXPackageInformationList(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *WindowsPhone81AppXBundle) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsPhone81AppX.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAppXPackageInformationList() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppXPackageInformationList()))
        for i, v := range m.GetAppXPackageInformationList() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("appXPackageInformationList", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppXPackageInformationList sets the appXPackageInformationList property value. The list of AppX Package Information.
func (m *WindowsPhone81AppXBundle) SetAppXPackageInformationList(value []WindowsPackageInformationable)() {
    m.appXPackageInformationList = value
}
