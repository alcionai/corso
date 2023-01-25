package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuditUserIdentityable 
type AuditUserIdentityable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UserIdentityable
    GetHomeTenantId()(*string)
    GetHomeTenantName()(*string)
    SetHomeTenantId(value *string)()
    SetHomeTenantName(value *string)()
}
