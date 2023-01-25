package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RbacApplicationMultipleable 
type RbacApplicationMultipleable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetResourceNamespaces()([]UnifiedRbacResourceNamespaceable)
    GetRoleAssignments()([]UnifiedRoleAssignmentMultipleable)
    GetRoleDefinitions()([]UnifiedRoleDefinitionable)
    SetResourceNamespaces(value []UnifiedRbacResourceNamespaceable)()
    SetRoleAssignments(value []UnifiedRoleAssignmentMultipleable)()
    SetRoleDefinitions(value []UnifiedRoleDefinitionable)()
}
