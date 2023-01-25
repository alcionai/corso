package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MyRoleable 
type MyRoleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignments()([]RoleAssignmentable)
    GetOdataType()(*string)
    GetTenantId()(*string)
    SetAssignments(value []RoleAssignmentable)()
    SetOdataType(value *string)()
    SetTenantId(value *string)()
}
