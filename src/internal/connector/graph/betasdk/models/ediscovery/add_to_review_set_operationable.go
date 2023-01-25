package ediscovery

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AddToReviewSetOperationable 
type AddToReviewSetOperationable interface {
    CaseOperationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetReviewSet()(ReviewSetable)
    GetSourceCollection()(SourceCollectionable)
    SetReviewSet(value ReviewSetable)()
    SetSourceCollection(value SourceCollectionable)()
}
