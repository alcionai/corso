package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdiscoveryExportOperationable 
type EdiscoveryExportOperationable interface {
    CaseOperationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAzureBlobContainer()(*string)
    GetAzureBlobToken()(*string)
    GetDescription()(*string)
    GetExportOptions()(*ExportOptions)
    GetExportStructure()(*ExportFileStructure)
    GetOutputFolderId()(*string)
    GetOutputName()(*string)
    GetReviewSet()(EdiscoveryReviewSetable)
    GetReviewSetQuery()(EdiscoveryReviewSetQueryable)
    SetAzureBlobContainer(value *string)()
    SetAzureBlobToken(value *string)()
    SetDescription(value *string)()
    SetExportOptions(value *ExportOptions)()
    SetExportStructure(value *ExportFileStructure)()
    SetOutputFolderId(value *string)()
    SetOutputName(value *string)()
    SetReviewSet(value EdiscoveryReviewSetable)()
    SetReviewSetQuery(value EdiscoveryReviewSetQueryable)()
}
