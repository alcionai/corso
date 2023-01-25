package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Settings 
type Settings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Specifies if the user's primary mailbox is hosted in the cloud and is enabled for Microsoft Graph.
    hasGraphMailbox *bool
    // Specifies if the user has a MyAnalytics license assigned.
    hasLicense *bool
    // Specifies if the user opted out of MyAnalytics.
    hasOptedOut *bool
    // The OdataType property
    odataType *string
}
// NewSettings instantiates a new settings and sets the default values.
func NewSettings()(*Settings) {
    m := &Settings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Settings) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Settings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["hasGraphMailbox"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasGraphMailbox(val)
        }
        return nil
    }
    res["hasLicense"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasLicense(val)
        }
        return nil
    }
    res["hasOptedOut"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasOptedOut(val)
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
// GetHasGraphMailbox gets the hasGraphMailbox property value. Specifies if the user's primary mailbox is hosted in the cloud and is enabled for Microsoft Graph.
func (m *Settings) GetHasGraphMailbox()(*bool) {
    return m.hasGraphMailbox
}
// GetHasLicense gets the hasLicense property value. Specifies if the user has a MyAnalytics license assigned.
func (m *Settings) GetHasLicense()(*bool) {
    return m.hasLicense
}
// GetHasOptedOut gets the hasOptedOut property value. Specifies if the user opted out of MyAnalytics.
func (m *Settings) GetHasOptedOut()(*bool) {
    return m.hasOptedOut
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Settings) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *Settings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("hasGraphMailbox", m.GetHasGraphMailbox())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hasLicense", m.GetHasLicense())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hasOptedOut", m.GetHasOptedOut())
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
func (m *Settings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetHasGraphMailbox sets the hasGraphMailbox property value. Specifies if the user's primary mailbox is hosted in the cloud and is enabled for Microsoft Graph.
func (m *Settings) SetHasGraphMailbox(value *bool)() {
    m.hasGraphMailbox = value
}
// SetHasLicense sets the hasLicense property value. Specifies if the user has a MyAnalytics license assigned.
func (m *Settings) SetHasLicense(value *bool)() {
    m.hasLicense = value
}
// SetHasOptedOut sets the hasOptedOut property value. Specifies if the user opted out of MyAnalytics.
func (m *Settings) SetHasOptedOut(value *bool)() {
    m.hasOptedOut = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Settings) SetOdataType(value *string)() {
    m.odataType = value
}
