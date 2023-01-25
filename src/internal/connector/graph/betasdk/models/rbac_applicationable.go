package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RbacApplicationable 
type RbacApplicationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetResourceNamespaces()([]UnifiedRbacResourceNamespaceable)
    GetRoleAssignmentApprovals()([]Approvalable)
    GetRoleAssignments()([]UnifiedRoleAssignmentable)
    GetRoleAssignmentScheduleInstances()([]UnifiedRoleAssignmentScheduleInstanceable)
    GetRoleAssignmentScheduleRequests()([]UnifiedRoleAssignmentScheduleRequestable)
    GetRoleAssignmentSchedules()([]UnifiedRoleAssignmentScheduleable)
    GetRoleDefinitions()([]UnifiedRoleDefinitionable)
    GetRoleEligibilityScheduleInstances()([]UnifiedRoleEligibilityScheduleInstanceable)
    GetRoleEligibilityScheduleRequests()([]UnifiedRoleEligibilityScheduleRequestable)
    GetRoleEligibilitySchedules()([]UnifiedRoleEligibilityScheduleable)
    GetTransitiveRoleAssignments()([]UnifiedRoleAssignmentable)
    SetResourceNamespaces(value []UnifiedRbacResourceNamespaceable)()
    SetRoleAssignmentApprovals(value []Approvalable)()
    SetRoleAssignments(value []UnifiedRoleAssignmentable)()
    SetRoleAssignmentScheduleInstances(value []UnifiedRoleAssignmentScheduleInstanceable)()
    SetRoleAssignmentScheduleRequests(value []UnifiedRoleAssignmentScheduleRequestable)()
    SetRoleAssignmentSchedules(value []UnifiedRoleAssignmentScheduleable)()
    SetRoleDefinitions(value []UnifiedRoleDefinitionable)()
    SetRoleEligibilityScheduleInstances(value []UnifiedRoleEligibilityScheduleInstanceable)()
    SetRoleEligibilityScheduleRequests(value []UnifiedRoleEligibilityScheduleRequestable)()
    SetRoleEligibilitySchedules(value []UnifiedRoleEligibilityScheduleable)()
    SetTransitiveRoleAssignments(value []UnifiedRoleAssignmentable)()
}
