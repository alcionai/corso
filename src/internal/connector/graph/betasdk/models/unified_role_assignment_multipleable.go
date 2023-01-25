package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleAssignmentMultipleable 
type UnifiedRoleAssignmentMultipleable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppScopeIds()([]string)
    GetAppScopes()([]AppScopeable)
    GetCondition()(*string)
    GetDescription()(*string)
    GetDirectoryScopeIds()([]string)
    GetDirectoryScopes()([]DirectoryObjectable)
    GetDisplayName()(*string)
    GetPrincipalIds()([]string)
    GetPrincipals()([]DirectoryObjectable)
    GetRoleDefinition()(UnifiedRoleDefinitionable)
    GetRoleDefinitionId()(*string)
    SetAppScopeIds(value []string)()
    SetAppScopes(value []AppScopeable)()
    SetCondition(value *string)()
    SetDescription(value *string)()
    SetDirectoryScopeIds(value []string)()
    SetDirectoryScopes(value []DirectoryObjectable)()
    SetDisplayName(value *string)()
    SetPrincipalIds(value []string)()
    SetPrincipals(value []DirectoryObjectable)()
    SetRoleDefinition(value UnifiedRoleDefinitionable)()
    SetRoleDefinitionId(value *string)()
}
