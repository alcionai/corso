package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerKioskModeManagedFolder a folder containing pages of apps and weblinks on the Managed Home Screen
type AndroidDeviceOwnerKioskModeManagedFolder struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Unique identifier for the folder
    folderIdentifier *string
    // Display name for the folder
    folderName *string
    // Items to be added to managed folder. This collection can contain a maximum of 500 elements.
    items []AndroidDeviceOwnerKioskModeFolderItemable
    // The OdataType property
    odataType *string
}
// NewAndroidDeviceOwnerKioskModeManagedFolder instantiates a new androidDeviceOwnerKioskModeManagedFolder and sets the default values.
func NewAndroidDeviceOwnerKioskModeManagedFolder()(*AndroidDeviceOwnerKioskModeManagedFolder) {
    m := &AndroidDeviceOwnerKioskModeManagedFolder{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAndroidDeviceOwnerKioskModeManagedFolderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerKioskModeManagedFolderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerKioskModeManagedFolder(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidDeviceOwnerKioskModeManagedFolder) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerKioskModeManagedFolder) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidDeviceOwnerKioskModeFolderItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidDeviceOwnerKioskModeFolderItemable, len(val))
            for i, v := range val {
                res[i] = v.(AndroidDeviceOwnerKioskModeFolderItemable)
            }
            m.SetItems(res)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetFolderIdentifier gets the folderIdentifier property value. Unique identifier for the folder
func (m *AndroidDeviceOwnerKioskModeManagedFolder) GetFolderIdentifier()(*string) {
    return m.folderIdentifier
}
// GetFolderName gets the folderName property value. Display name for the folder
func (m *AndroidDeviceOwnerKioskModeManagedFolder) GetFolderName()(*string) {
    return m.folderName
}
// GetItems gets the items property value. Items to be added to managed folder. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerKioskModeManagedFolder) GetItems()([]AndroidDeviceOwnerKioskModeFolderItemable) {
    return m.items
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AndroidDeviceOwnerKioskModeManagedFolder) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AndroidDeviceOwnerKioskModeManagedFolder) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("folderIdentifier", m.GetFolderIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("folderName", m.GetFolderName())
        if err != nil {
            return err
        }
    }
    if m.GetItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItems()))
        for i, v := range m.GetItems() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("items", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidDeviceOwnerKioskModeManagedFolder) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetFolderIdentifier sets the folderIdentifier property value. Unique identifier for the folder
func (m *AndroidDeviceOwnerKioskModeManagedFolder) SetFolderIdentifier(value *string)() {
    m.folderIdentifier = value
}
// SetFolderName sets the folderName property value. Display name for the folder
func (m *AndroidDeviceOwnerKioskModeManagedFolder) SetFolderName(value *string)() {
    m.folderName = value
}
// SetItems sets the items property value. Items to be added to managed folder. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerKioskModeManagedFolder) SetItems(value []AndroidDeviceOwnerKioskModeFolderItemable)() {
    m.items = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AndroidDeviceOwnerKioskModeManagedFolder) SetOdataType(value *string)() {
    m.odataType = value
}
