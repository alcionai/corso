package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyUploadedDefinitionFileable 
type GroupPolicyUploadedDefinitionFileable interface {
    GroupPolicyDefinitionFileable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()([]byte)
    GetDefaultLanguageCode()(*string)
    GetGroupPolicyOperations()([]GroupPolicyOperationable)
    GetGroupPolicyUploadedLanguageFiles()([]GroupPolicyUploadedLanguageFileable)
    GetStatus()(*GroupPolicyUploadedDefinitionFileStatus)
    GetUploadDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetContent(value []byte)()
    SetDefaultLanguageCode(value *string)()
    SetGroupPolicyOperations(value []GroupPolicyOperationable)()
    SetGroupPolicyUploadedLanguageFiles(value []GroupPolicyUploadedLanguageFileable)()
    SetStatus(value *GroupPolicyUploadedDefinitionFileStatus)()
    SetUploadDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
