package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChannelModerationSettingsable 
type ChannelModerationSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowNewMessageFromBots()(*bool)
    GetAllowNewMessageFromConnectors()(*bool)
    GetOdataType()(*string)
    GetReplyRestriction()(*ReplyRestriction)
    GetUserNewMessageRestriction()(*UserNewMessageRestriction)
    SetAllowNewMessageFromBots(value *bool)()
    SetAllowNewMessageFromConnectors(value *bool)()
    SetOdataType(value *string)()
    SetReplyRestriction(value *ReplyRestriction)()
    SetUserNewMessageRestriction(value *UserNewMessageRestriction)()
}
