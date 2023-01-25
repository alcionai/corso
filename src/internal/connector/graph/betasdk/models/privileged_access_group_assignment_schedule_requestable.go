package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedAccessGroupAssignmentScheduleRequestable 
type PrivilegedAccessGroupAssignmentScheduleRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PrivilegedAccessScheduleRequestable
    GetAccessId()(*PrivilegedAccessGroupRelationships)
    GetActivatedUsing()(PrivilegedAccessGroupEligibilityScheduleable)
    GetGroup()(Groupable)
    GetGroupId()(*string)
    GetPrincipal()(DirectoryObjectable)
    GetPrincipalId()(*string)
    GetTargetSchedule()(PrivilegedAccessGroupEligibilityScheduleable)
    GetTargetScheduleId()(*string)
    SetAccessId(value *PrivilegedAccessGroupRelationships)()
    SetActivatedUsing(value PrivilegedAccessGroupEligibilityScheduleable)()
    SetGroup(value Groupable)()
    SetGroupId(value *string)()
    SetPrincipal(value DirectoryObjectable)()
    SetPrincipalId(value *string)()
    SetTargetSchedule(value PrivilegedAccessGroupEligibilityScheduleable)()
    SetTargetScheduleId(value *string)()
}
