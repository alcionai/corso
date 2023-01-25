package ediscovery

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EstimateStatisticsOperationable 
type EstimateStatisticsOperationable interface {
    CaseOperationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIndexedItemCount()(*int64)
    GetIndexedItemsSize()(*int64)
    GetMailboxCount()(*int32)
    GetSiteCount()(*int32)
    GetSourceCollection()(SourceCollectionable)
    GetUnindexedItemCount()(*int64)
    GetUnindexedItemsSize()(*int64)
    SetIndexedItemCount(value *int64)()
    SetIndexedItemsSize(value *int64)()
    SetMailboxCount(value *int32)()
    SetSiteCount(value *int32)()
    SetSourceCollection(value SourceCollectionable)()
    SetUnindexedItemCount(value *int64)()
    SetUnindexedItemsSize(value *int64)()
}
