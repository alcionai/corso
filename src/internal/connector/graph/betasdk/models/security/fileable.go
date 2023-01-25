package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// Fileable 
type Fileable interface {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()([]byte)
    GetDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExtension()(*string)
    GetExtractedTextContent()([]byte)
    GetMediaType()(*string)
    GetName()(*string)
    GetOtherProperties()(StringValueDictionaryable)
    GetProcessingStatus()(*FileProcessingStatus)
    GetSenderOrAuthors()([]string)
    GetSize()(*int64)
    GetSourceType()(*SourceType)
    GetSubjectTitle()(*string)
    SetContent(value []byte)()
    SetDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExtension(value *string)()
    SetExtractedTextContent(value []byte)()
    SetMediaType(value *string)()
    SetName(value *string)()
    SetOtherProperties(value StringValueDictionaryable)()
    SetProcessingStatus(value *FileProcessingStatus)()
    SetSenderOrAuthors(value []string)()
    SetSize(value *int64)()
    SetSourceType(value *SourceType)()
    SetSubjectTitle(value *string)()
}
