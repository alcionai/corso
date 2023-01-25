package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSKernelExtension represents a specific macOS kernel extension. A macOS kernel extension can be described by its team identifier plus its bundle identifier.
type MacOSKernelExtension struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Bundle ID of the kernel extension.
    bundleId *string
    // The OdataType property
    odataType *string
    // The team identifier that was used to sign the kernel extension.
    teamIdentifier *string
}
// NewMacOSKernelExtension instantiates a new macOSKernelExtension and sets the default values.
func NewMacOSKernelExtension()(*MacOSKernelExtension) {
    m := &MacOSKernelExtension{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMacOSKernelExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSKernelExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSKernelExtension(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MacOSKernelExtension) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetBundleId gets the bundleId property value. Bundle ID of the kernel extension.
func (m *MacOSKernelExtension) GetBundleId()(*string) {
    return m.bundleId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSKernelExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bundleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBundleId(val)
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
    res["teamIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamIdentifier(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MacOSKernelExtension) GetOdataType()(*string) {
    return m.odataType
}
// GetTeamIdentifier gets the teamIdentifier property value. The team identifier that was used to sign the kernel extension.
func (m *MacOSKernelExtension) GetTeamIdentifier()(*string) {
    return m.teamIdentifier
}
// Serialize serializes information the current object
func (m *MacOSKernelExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("bundleId", m.GetBundleId())
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
        err := writer.WriteStringValue("teamIdentifier", m.GetTeamIdentifier())
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
func (m *MacOSKernelExtension) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetBundleId sets the bundleId property value. Bundle ID of the kernel extension.
func (m *MacOSKernelExtension) SetBundleId(value *string)() {
    m.bundleId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MacOSKernelExtension) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTeamIdentifier sets the teamIdentifier property value. The team identifier that was used to sign the kernel extension.
func (m *MacOSKernelExtension) SetTeamIdentifier(value *string)() {
    m.teamIdentifier = value
}
