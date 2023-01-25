package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSSystemExtension represents a specific macOS system extension.
type MacOSSystemExtension struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Gets or sets the bundle identifier of the system extension.
    bundleId *string
    // The OdataType property
    odataType *string
    // Gets or sets the team identifier that was used to sign the system extension.
    teamIdentifier *string
}
// NewMacOSSystemExtension instantiates a new macOSSystemExtension and sets the default values.
func NewMacOSSystemExtension()(*MacOSSystemExtension) {
    m := &MacOSSystemExtension{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMacOSSystemExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSSystemExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSSystemExtension(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MacOSSystemExtension) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetBundleId gets the bundleId property value. Gets or sets the bundle identifier of the system extension.
func (m *MacOSSystemExtension) GetBundleId()(*string) {
    return m.bundleId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSSystemExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *MacOSSystemExtension) GetOdataType()(*string) {
    return m.odataType
}
// GetTeamIdentifier gets the teamIdentifier property value. Gets or sets the team identifier that was used to sign the system extension.
func (m *MacOSSystemExtension) GetTeamIdentifier()(*string) {
    return m.teamIdentifier
}
// Serialize serializes information the current object
func (m *MacOSSystemExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *MacOSSystemExtension) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetBundleId sets the bundleId property value. Gets or sets the bundle identifier of the system extension.
func (m *MacOSSystemExtension) SetBundleId(value *string)() {
    m.bundleId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MacOSSystemExtension) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTeamIdentifier sets the teamIdentifier property value. Gets or sets the team identifier that was used to sign the system extension.
func (m *MacOSSystemExtension) SetTeamIdentifier(value *string)() {
    m.teamIdentifier = value
}
