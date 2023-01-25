package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkOnPremisesCalendarSyncConfiguration 
type TeamworkOnPremisesCalendarSyncConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The fully qualified domain name (FQDN) of the Skype for Business Server. Use the Exchange domain if the Skype for Business SIP domain is different from the Exchange domain of the user.
    domain *string
    // The domain and username of the console device, for example, Seattle/RanierConf.
    domainUserName *string
    // The OdataType property
    odataType *string
    // The Simple Mail Transfer Protocol (SMTP) address of the user account. This is only required if a different user principal name (UPN) is used to sign in to Exchange other than Microsoft Teams and Skype for Business. This is a common scenario in a hybrid environment where an on-premises Exchange server is used.
    smtpAddress *string
}
// NewTeamworkOnPremisesCalendarSyncConfiguration instantiates a new teamworkOnPremisesCalendarSyncConfiguration and sets the default values.
func NewTeamworkOnPremisesCalendarSyncConfiguration()(*TeamworkOnPremisesCalendarSyncConfiguration) {
    m := &TeamworkOnPremisesCalendarSyncConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamworkOnPremisesCalendarSyncConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkOnPremisesCalendarSyncConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkOnPremisesCalendarSyncConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkOnPremisesCalendarSyncConfiguration) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDomain gets the domain property value. The fully qualified domain name (FQDN) of the Skype for Business Server. Use the Exchange domain if the Skype for Business SIP domain is different from the Exchange domain of the user.
func (m *TeamworkOnPremisesCalendarSyncConfiguration) GetDomain()(*string) {
    return m.domain
}
// GetDomainUserName gets the domainUserName property value. The domain and username of the console device, for example, Seattle/RanierConf.
func (m *TeamworkOnPremisesCalendarSyncConfiguration) GetDomainUserName()(*string) {
    return m.domainUserName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkOnPremisesCalendarSyncConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["domain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomain(val)
        }
        return nil
    }
    res["domainUserName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainUserName(val)
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
    res["smtpAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmtpAddress(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkOnPremisesCalendarSyncConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetSmtpAddress gets the smtpAddress property value. The Simple Mail Transfer Protocol (SMTP) address of the user account. This is only required if a different user principal name (UPN) is used to sign in to Exchange other than Microsoft Teams and Skype for Business. This is a common scenario in a hybrid environment where an on-premises Exchange server is used.
func (m *TeamworkOnPremisesCalendarSyncConfiguration) GetSmtpAddress()(*string) {
    return m.smtpAddress
}
// Serialize serializes information the current object
func (m *TeamworkOnPremisesCalendarSyncConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("domain", m.GetDomain())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("domainUserName", m.GetDomainUserName())
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
        err := writer.WriteStringValue("smtpAddress", m.GetSmtpAddress())
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
func (m *TeamworkOnPremisesCalendarSyncConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDomain sets the domain property value. The fully qualified domain name (FQDN) of the Skype for Business Server. Use the Exchange domain if the Skype for Business SIP domain is different from the Exchange domain of the user.
func (m *TeamworkOnPremisesCalendarSyncConfiguration) SetDomain(value *string)() {
    m.domain = value
}
// SetDomainUserName sets the domainUserName property value. The domain and username of the console device, for example, Seattle/RanierConf.
func (m *TeamworkOnPremisesCalendarSyncConfiguration) SetDomainUserName(value *string)() {
    m.domainUserName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkOnPremisesCalendarSyncConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSmtpAddress sets the smtpAddress property value. The Simple Mail Transfer Protocol (SMTP) address of the user account. This is only required if a different user principal name (UPN) is used to sign in to Exchange other than Microsoft Teams and Skype for Business. This is a common scenario in a hybrid environment where an on-premises Exchange server is used.
func (m *TeamworkOnPremisesCalendarSyncConfiguration) SetSmtpAddress(value *string)() {
    m.smtpAddress = value
}
