package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedAccessGroupable 
type PrivilegedAccessGroupable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignmentScheduleInstances()([]PrivilegedAccessGroupAssignmentScheduleInstanceable)
    GetAssignmentScheduleRequests()([]PrivilegedAccessGroupAssignmentScheduleRequestable)
    GetAssignmentSchedules()([]PrivilegedAccessGroupAssignmentScheduleable)
    GetEligibilityScheduleInstances()([]PrivilegedAccessGroupEligibilityScheduleInstanceable)
    GetEligibilityScheduleRequests()([]PrivilegedAccessGroupEligibilityScheduleRequestable)
    GetEligibilitySchedules()([]PrivilegedAccessGroupEligibilityScheduleable)
    SetAssignmentScheduleInstances(value []PrivilegedAccessGroupAssignmentScheduleInstanceable)()
    SetAssignmentScheduleRequests(value []PrivilegedAccessGroupAssignmentScheduleRequestable)()
    SetAssignmentSchedules(value []PrivilegedAccessGroupAssignmentScheduleable)()
    SetEligibilityScheduleInstances(value []PrivilegedAccessGroupEligibilityScheduleInstanceable)()
    SetEligibilityScheduleRequests(value []PrivilegedAccessGroupEligibilityScheduleRequestable)()
    SetEligibilitySchedules(value []PrivilegedAccessGroupEligibilityScheduleable)()
}
