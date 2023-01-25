package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintUsageable 
type PrintUsageable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBlackAndWhitePageCount()(*int64)
    GetColorPageCount()(*int64)
    GetCompletedBlackAndWhiteJobCount()(*int64)
    GetCompletedColorJobCount()(*int64)
    GetCompletedJobCount()(*int64)
    GetDoubleSidedSheetCount()(*int64)
    GetIncompleteJobCount()(*int64)
    GetMediaSheetCount()(*int64)
    GetPageCount()(*int64)
    GetSingleSidedSheetCount()(*int64)
    GetUsageDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    SetBlackAndWhitePageCount(value *int64)()
    SetColorPageCount(value *int64)()
    SetCompletedBlackAndWhiteJobCount(value *int64)()
    SetCompletedColorJobCount(value *int64)()
    SetCompletedJobCount(value *int64)()
    SetDoubleSidedSheetCount(value *int64)()
    SetIncompleteJobCount(value *int64)()
    SetMediaSheetCount(value *int64)()
    SetPageCount(value *int64)()
    SetSingleSidedSheetCount(value *int64)()
    SetUsageDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
}
