package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TenantContactInformation 
type TenantContactInformation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The email address for the contact. Optional
    email *string
    // The name for the contact. Required.
    name *string
    // The notes associated with the contact. Optional
    notes *string
    // The OdataType property
    odataType *string
    // The phone number for the contact. Optional.
    phone *string
    // The title for the contact. Required.
    title *string
}
// NewTenantContactInformation instantiates a new tenantContactInformation and sets the default values.
func NewTenantContactInformation()(*TenantContactInformation) {
    m := &TenantContactInformation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTenantContactInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantContactInformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenantContactInformation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TenantContactInformation) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEmail gets the email property value. The email address for the contact. Optional
func (m *TenantContactInformation) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantContactInformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val)
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
    res["phone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhone(val)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name for the contact. Required.
func (m *TenantContactInformation) GetName()(*string) {
    return m.name
}
// GetNotes gets the notes property value. The notes associated with the contact. Optional
func (m *TenantContactInformation) GetNotes()(*string) {
    return m.notes
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TenantContactInformation) GetOdataType()(*string) {
    return m.odataType
}
// GetPhone gets the phone property value. The phone number for the contact. Optional.
func (m *TenantContactInformation) GetPhone()(*string) {
    return m.phone
}
// GetTitle gets the title property value. The title for the contact. Required.
func (m *TenantContactInformation) GetTitle()(*string) {
    return m.title
}
// Serialize serializes information the current object
func (m *TenantContactInformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("notes", m.GetNotes())
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
        err := writer.WriteStringValue("phone", m.GetPhone())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
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
func (m *TenantContactInformation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEmail sets the email property value. The email address for the contact. Optional
func (m *TenantContactInformation) SetEmail(value *string)() {
    m.email = value
}
// SetName sets the name property value. The name for the contact. Required.
func (m *TenantContactInformation) SetName(value *string)() {
    m.name = value
}
// SetNotes sets the notes property value. The notes associated with the contact. Optional
func (m *TenantContactInformation) SetNotes(value *string)() {
    m.notes = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TenantContactInformation) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPhone sets the phone property value. The phone number for the contact. Optional.
func (m *TenantContactInformation) SetPhone(value *string)() {
    m.phone = value
}
// SetTitle sets the title property value. The title for the contact. Required.
func (m *TenantContactInformation) SetTitle(value *string)() {
    m.title = value
}
