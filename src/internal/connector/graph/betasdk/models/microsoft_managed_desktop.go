package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftManagedDesktop 
type MicrosoftManagedDesktop struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // The name of the Microsoft Managed Desktop profile that the Windows 365 Cloud PC is associated with.
    profile *string
    // Indicates whether the provisioning policy enables Microsoft Managed Desktop. It indicates the type of plan under which the device is managed if the provisioning policy is enabled. Possible values are: notManaged, premiumManaged, standardManaged, starterManaged, unknownFutureValue.
    type_escaped *MicrosoftManagedDesktopType
}
// NewMicrosoftManagedDesktop instantiates a new microsoftManagedDesktop and sets the default values.
func NewMicrosoftManagedDesktop()(*MicrosoftManagedDesktop) {
    m := &MicrosoftManagedDesktop{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMicrosoftManagedDesktopFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftManagedDesktopFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftManagedDesktop(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MicrosoftManagedDesktop) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftManagedDesktop) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["profile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfile(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMicrosoftManagedDesktopType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*MicrosoftManagedDesktopType))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MicrosoftManagedDesktop) GetOdataType()(*string) {
    return m.odataType
}
// GetProfile gets the profile property value. The name of the Microsoft Managed Desktop profile that the Windows 365 Cloud PC is associated with.
func (m *MicrosoftManagedDesktop) GetProfile()(*string) {
    return m.profile
}
// GetType gets the type property value. Indicates whether the provisioning policy enables Microsoft Managed Desktop. It indicates the type of plan under which the device is managed if the provisioning policy is enabled. Possible values are: notManaged, premiumManaged, standardManaged, starterManaged, unknownFutureValue.
func (m *MicrosoftManagedDesktop) GetType()(*MicrosoftManagedDesktopType) {
    return m.type_escaped
}
// Serialize serializes information the current object
func (m *MicrosoftManagedDesktop) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("profile", m.GetProfile())
        if err != nil {
            return err
        }
    }
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *MicrosoftManagedDesktop) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MicrosoftManagedDesktop) SetOdataType(value *string)() {
    m.odataType = value
}
// SetProfile sets the profile property value. The name of the Microsoft Managed Desktop profile that the Windows 365 Cloud PC is associated with.
func (m *MicrosoftManagedDesktop) SetProfile(value *string)() {
    m.profile = value
}
// SetType sets the type property value. Indicates whether the provisioning policy enables Microsoft Managed Desktop. It indicates the type of plan under which the device is managed if the provisioning policy is enabled. Possible values are: notManaged, premiumManaged, standardManaged, starterManaged, unknownFutureValue.
func (m *MicrosoftManagedDesktop) SetType(value *MicrosoftManagedDesktopType)() {
    m.type_escaped = value
}
