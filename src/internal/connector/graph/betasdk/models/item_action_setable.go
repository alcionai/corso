package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemActionSetable 
type ItemActionSetable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComment()(CommentActionable)
    GetCreate()(CreateActionable)
    GetDelete()(DeleteActionable)
    GetEdit()(EditActionable)
    GetMention()(MentionActionable)
    GetMove()(MoveActionable)
    GetOdataType()(*string)
    GetRename()(RenameActionable)
    GetRestore()(RestoreActionable)
    GetShare()(ShareActionable)
    GetVersion()(VersionActionable)
    SetComment(value CommentActionable)()
    SetCreate(value CreateActionable)()
    SetDelete(value DeleteActionable)()
    SetEdit(value EditActionable)()
    SetMention(value MentionActionable)()
    SetMove(value MoveActionable)()
    SetOdataType(value *string)()
    SetRename(value RenameActionable)()
    SetRestore(value RestoreActionable)()
    SetShare(value ShareActionable)()
    SetVersion(value VersionActionable)()
}
