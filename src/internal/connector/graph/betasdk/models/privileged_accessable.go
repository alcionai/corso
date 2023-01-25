package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedAccessable 
type PrivilegedAccessable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetResources()([]GovernanceResourceable)
    GetRoleAssignmentRequests()([]GovernanceRoleAssignmentRequestable)
    GetRoleAssignments()([]GovernanceRoleAssignmentable)
    GetRoleDefinitions()([]GovernanceRoleDefinitionable)
    GetRoleSettings()([]GovernanceRoleSettingable)
    SetDisplayName(value *string)()
    SetResources(value []GovernanceResourceable)()
    SetRoleAssignmentRequests(value []GovernanceRoleAssignmentRequestable)()
    SetRoleAssignments(value []GovernanceRoleAssignmentable)()
    SetRoleDefinitions(value []GovernanceRoleDefinitionable)()
    SetRoleSettings(value []GovernanceRoleSettingable)()
}
