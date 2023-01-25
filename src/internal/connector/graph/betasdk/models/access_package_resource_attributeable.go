package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageResourceAttributeable 
type AccessPackageResourceAttributeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttributeDestination()(AccessPackageResourceAttributeDestinationable)
    GetAttributeName()(*string)
    GetAttributeSource()(AccessPackageResourceAttributeSourceable)
    GetId()(*string)
    GetIsEditable()(*bool)
    GetIsPersistedOnAssignmentRemoval()(*bool)
    GetOdataType()(*string)
    SetAttributeDestination(value AccessPackageResourceAttributeDestinationable)()
    SetAttributeName(value *string)()
    SetAttributeSource(value AccessPackageResourceAttributeSourceable)()
    SetId(value *string)()
    SetIsEditable(value *bool)()
    SetIsPersistedOnAssignmentRemoval(value *bool)()
    SetOdataType(value *string)()
}
