package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommunicationsIdentitySet 
type CommunicationsIdentitySet struct {
    IdentitySet
    // The application instance associated with this action.
    applicationInstance Identityable
    // An identity the participant would like to present itself as to the other participants in the call.
    assertedIdentity Identityable
    // The Azure Communication Services user associated with this action.
    azureCommunicationServicesUser Identityable
    // The encrypted user associated with this action.
    encrypted Identityable
    // Type of endpoint the participant is using. Possible values are: default, voicemail, skypeForBusiness, skypeForBusinessVoipPhone and unknownFutureValue.
    endpointType *EndpointType
    // The guest user associated with this action.
    guest Identityable
    // The Skype for Business On-Premises user associated with this action.
    onPremises Identityable
    // Inherited from identitySet. The phone user associated with this action.
    phone Identityable
}
// NewCommunicationsIdentitySet instantiates a new CommunicationsIdentitySet and sets the default values.
func NewCommunicationsIdentitySet()(*CommunicationsIdentitySet) {
    m := &CommunicationsIdentitySet{
        IdentitySet: *NewIdentitySet(),
    }
    odataTypeValue := "#microsoft.graph.communicationsIdentitySet";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCommunicationsIdentitySetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCommunicationsIdentitySetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommunicationsIdentitySet(), nil
}
// GetApplicationInstance gets the applicationInstance property value. The application instance associated with this action.
func (m *CommunicationsIdentitySet) GetApplicationInstance()(Identityable) {
    return m.applicationInstance
}
// GetAssertedIdentity gets the assertedIdentity property value. An identity the participant would like to present itself as to the other participants in the call.
func (m *CommunicationsIdentitySet) GetAssertedIdentity()(Identityable) {
    return m.assertedIdentity
}
// GetAzureCommunicationServicesUser gets the azureCommunicationServicesUser property value. The Azure Communication Services user associated with this action.
func (m *CommunicationsIdentitySet) GetAzureCommunicationServicesUser()(Identityable) {
    return m.azureCommunicationServicesUser
}
// GetEncrypted gets the encrypted property value. The encrypted user associated with this action.
func (m *CommunicationsIdentitySet) GetEncrypted()(Identityable) {
    return m.encrypted
}
// GetEndpointType gets the endpointType property value. Type of endpoint the participant is using. Possible values are: default, voicemail, skypeForBusiness, skypeForBusinessVoipPhone and unknownFutureValue.
func (m *CommunicationsIdentitySet) GetEndpointType()(*EndpointType) {
    return m.endpointType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CommunicationsIdentitySet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentitySet.GetFieldDeserializers()
    res["applicationInstance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationInstance(val.(Identityable))
        }
        return nil
    }
    res["assertedIdentity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssertedIdentity(val.(Identityable))
        }
        return nil
    }
    res["azureCommunicationServicesUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureCommunicationServicesUser(val.(Identityable))
        }
        return nil
    }
    res["encrypted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncrypted(val.(Identityable))
        }
        return nil
    }
    res["endpointType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEndpointType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndpointType(val.(*EndpointType))
        }
        return nil
    }
    res["guest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGuest(val.(Identityable))
        }
        return nil
    }
    res["onPremises"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremises(val.(Identityable))
        }
        return nil
    }
    res["phone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhone(val.(Identityable))
        }
        return nil
    }
    return res
}
// GetGuest gets the guest property value. The guest user associated with this action.
func (m *CommunicationsIdentitySet) GetGuest()(Identityable) {
    return m.guest
}
// GetOnPremises gets the onPremises property value. The Skype for Business On-Premises user associated with this action.
func (m *CommunicationsIdentitySet) GetOnPremises()(Identityable) {
    return m.onPremises
}
// GetPhone gets the phone property value. Inherited from identitySet. The phone user associated with this action.
func (m *CommunicationsIdentitySet) GetPhone()(Identityable) {
    return m.phone
}
// Serialize serializes information the current object
func (m *CommunicationsIdentitySet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentitySet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("applicationInstance", m.GetApplicationInstance())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("assertedIdentity", m.GetAssertedIdentity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("azureCommunicationServicesUser", m.GetAzureCommunicationServicesUser())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("encrypted", m.GetEncrypted())
        if err != nil {
            return err
        }
    }
    if m.GetEndpointType() != nil {
        cast := (*m.GetEndpointType()).String()
        err = writer.WriteStringValue("endpointType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("guest", m.GetGuest())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("onPremises", m.GetOnPremises())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("phone", m.GetPhone())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicationInstance sets the applicationInstance property value. The application instance associated with this action.
func (m *CommunicationsIdentitySet) SetApplicationInstance(value Identityable)() {
    m.applicationInstance = value
}
// SetAssertedIdentity sets the assertedIdentity property value. An identity the participant would like to present itself as to the other participants in the call.
func (m *CommunicationsIdentitySet) SetAssertedIdentity(value Identityable)() {
    m.assertedIdentity = value
}
// SetAzureCommunicationServicesUser sets the azureCommunicationServicesUser property value. The Azure Communication Services user associated with this action.
func (m *CommunicationsIdentitySet) SetAzureCommunicationServicesUser(value Identityable)() {
    m.azureCommunicationServicesUser = value
}
// SetEncrypted sets the encrypted property value. The encrypted user associated with this action.
func (m *CommunicationsIdentitySet) SetEncrypted(value Identityable)() {
    m.encrypted = value
}
// SetEndpointType sets the endpointType property value. Type of endpoint the participant is using. Possible values are: default, voicemail, skypeForBusiness, skypeForBusinessVoipPhone and unknownFutureValue.
func (m *CommunicationsIdentitySet) SetEndpointType(value *EndpointType)() {
    m.endpointType = value
}
// SetGuest sets the guest property value. The guest user associated with this action.
func (m *CommunicationsIdentitySet) SetGuest(value Identityable)() {
    m.guest = value
}
// SetOnPremises sets the onPremises property value. The Skype for Business On-Premises user associated with this action.
func (m *CommunicationsIdentitySet) SetOnPremises(value Identityable)() {
    m.onPremises = value
}
// SetPhone sets the phone property value. Inherited from identitySet. The phone user associated with this action.
func (m *CommunicationsIdentitySet) SetPhone(value Identityable)() {
    m.phone = value
}
