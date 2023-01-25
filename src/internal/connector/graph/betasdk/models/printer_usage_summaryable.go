package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrinterUsageSummaryable 
type PrinterUsageSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompletedJobCount()(*int32)
    GetIncompleteJobCount()(*int32)
    GetOdataType()(*string)
    GetPrinter()(DirectoryObjectable)
    GetPrinterDisplayName()(*string)
    GetPrinterId()(*string)
    GetPrinterManufacturer()(*string)
    GetPrinterModel()(*string)
    SetCompletedJobCount(value *int32)()
    SetIncompleteJobCount(value *int32)()
    SetOdataType(value *string)()
    SetPrinter(value DirectoryObjectable)()
    SetPrinterDisplayName(value *string)()
    SetPrinterId(value *string)()
    SetPrinterManufacturer(value *string)()
    SetPrinterModel(value *string)()
}
