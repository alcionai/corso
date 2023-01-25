package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VirtualAppointmentUser 
type VirtualAppointmentUser struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The display name of the user who participates in a virtual appointment. Optional.
    displayName *string
    // The email address of the user who participates in a virtual appointment. Optional.
    emailAddress *string
    // The OdataType property
    odataType *string
    // The phone number for sending SMS texts for the user who participates in a virtual appointment. Optional.
    smsCapablePhoneNumber *string
}
// NewVirtualAppointmentUser instantiates a new virtualAppointmentUser and sets the default values.
func NewVirtualAppointmentUser()(*VirtualAppointmentUser) {
    m := &VirtualAppointmentUser{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateVirtualAppointmentUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVirtualAppointmentUserFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualAppointmentUser(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VirtualAppointmentUser) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. The display name of the user who participates in a virtual appointment. Optional.
func (m *VirtualAppointmentUser) GetDisplayName()(*string) {
    return m.displayName
}
// GetEmailAddress gets the emailAddress property value. The email address of the user who participates in a virtual appointment. Optional.
func (m *VirtualAppointmentUser) GetEmailAddress()(*string) {
    return m.emailAddress
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VirtualAppointmentUser) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["emailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddress(val)
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
    res["smsCapablePhoneNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmsCapablePhoneNumber(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *VirtualAppointmentUser) GetOdataType()(*string) {
    return m.odataType
}
// GetSmsCapablePhoneNumber gets the smsCapablePhoneNumber property value. The phone number for sending SMS texts for the user who participates in a virtual appointment. Optional.
func (m *VirtualAppointmentUser) GetSmsCapablePhoneNumber()(*string) {
    return m.smsCapablePhoneNumber
}
// Serialize serializes information the current object
func (m *VirtualAppointmentUser) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("emailAddress", m.GetEmailAddress())
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
        err := writer.WriteStringValue("smsCapablePhoneNumber", m.GetSmsCapablePhoneNumber())
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
func (m *VirtualAppointmentUser) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. The display name of the user who participates in a virtual appointment. Optional.
func (m *VirtualAppointmentUser) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEmailAddress sets the emailAddress property value. The email address of the user who participates in a virtual appointment. Optional.
func (m *VirtualAppointmentUser) SetEmailAddress(value *string)() {
    m.emailAddress = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *VirtualAppointmentUser) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSmsCapablePhoneNumber sets the smsCapablePhoneNumber property value. The phone number for sending SMS texts for the user who participates in a virtual appointment. Optional.
func (m *VirtualAppointmentUser) SetSmsCapablePhoneNumber(value *string)() {
    m.smsCapablePhoneNumber = value
}
