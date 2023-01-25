package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommunicationsApplicationInstanceIdentityable 
type CommunicationsApplicationInstanceIdentityable interface {
    Identityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHidden()(*bool)
    GetTenantId()(*string)
    SetHidden(value *bool)()
    SetTenantId(value *string)()
}
