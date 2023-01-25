package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SubmissionResultable 
type SubmissionResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategory()(*SubmissionResultCategory)
    GetDetail()(*SubmissionResultDetail)
    GetDetectedFiles()([]SubmissionDetectedFileable)
    GetDetectedUrls()([]string)
    GetOdataType()(*string)
    GetUserMailboxSetting()(*UserMailboxSetting)
    SetCategory(value *SubmissionResultCategory)()
    SetDetail(value *SubmissionResultDetail)()
    SetDetectedFiles(value []SubmissionDetectedFileable)()
    SetDetectedUrls(value []string)()
    SetOdataType(value *string)()
    SetUserMailboxSetting(value *UserMailboxSetting)()
}
