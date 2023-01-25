package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommunicationsIdentitySetable 
type CommunicationsIdentitySetable interface {
    IdentitySetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationInstance()(Identityable)
    GetAssertedIdentity()(Identityable)
    GetAzureCommunicationServicesUser()(Identityable)
    GetEncrypted()(Identityable)
    GetEndpointType()(*EndpointType)
    GetGuest()(Identityable)
    GetOnPremises()(Identityable)
    GetPhone()(Identityable)
    SetApplicationInstance(value Identityable)()
    SetAssertedIdentity(value Identityable)()
    SetAzureCommunicationServicesUser(value Identityable)()
    SetEncrypted(value Identityable)()
    SetEndpointType(value *EndpointType)()
    SetGuest(value Identityable)()
    SetOnPremises(value Identityable)()
    SetPhone(value Identityable)()
}
