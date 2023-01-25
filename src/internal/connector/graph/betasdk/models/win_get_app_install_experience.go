package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WinGetAppInstallExperience represents the install experience settings associated with WinGet apps. This is used to ensure the desired install experiences on the target device are taken into account. Required at creation time.
type WinGetAppInstallExperience struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // Indicates the type of execution context the app runs in.
    runAsAccount *RunAsAccountType
}
// NewWinGetAppInstallExperience instantiates a new winGetAppInstallExperience and sets the default values.
func NewWinGetAppInstallExperience()(*WinGetAppInstallExperience) {
    m := &WinGetAppInstallExperience{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWinGetAppInstallExperienceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWinGetAppInstallExperienceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWinGetAppInstallExperience(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WinGetAppInstallExperience) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WinGetAppInstallExperience) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["runAsAccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRunAsAccountType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunAsAccount(val.(*RunAsAccountType))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WinGetAppInstallExperience) GetOdataType()(*string) {
    return m.odataType
}
// GetRunAsAccount gets the runAsAccount property value. Indicates the type of execution context the app runs in.
func (m *WinGetAppInstallExperience) GetRunAsAccount()(*RunAsAccountType) {
    return m.runAsAccount
}
// Serialize serializes information the current object
func (m *WinGetAppInstallExperience) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetRunAsAccount() != nil {
        cast := (*m.GetRunAsAccount()).String()
        err := writer.WriteStringValue("runAsAccount", &cast)
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
func (m *WinGetAppInstallExperience) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WinGetAppInstallExperience) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRunAsAccount sets the runAsAccount property value. Indicates the type of execution context the app runs in.
func (m *WinGetAppInstallExperience) SetRunAsAccount(value *RunAsAccountType)() {
    m.runAsAccount = value
}
