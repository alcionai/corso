package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MembershipOutlierInsightable 
type MembershipOutlierInsightable interface {
    GovernanceInsightable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContainer()(DirectoryObjectable)
    GetContainerId()(*string)
    GetMember()(DirectoryObjectable)
    GetMemberId()(*string)
    GetOutlierContainerType()(*OutlierContainerType)
    GetOutlierMemberType()(*OutlierMemberType)
    SetContainer(value DirectoryObjectable)()
    SetContainerId(value *string)()
    SetMember(value DirectoryObjectable)()
    SetMemberId(value *string)()
    SetOutlierContainerType(value *OutlierContainerType)()
    SetOutlierMemberType(value *OutlierMemberType)()
}
