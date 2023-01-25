package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Noteable 
type Noteable interface {
    OutlookItemable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttachments()([]Attachmentable)
    GetBody()(ItemBodyable)
    GetExtensions()([]Extensionable)
    GetHasAttachments()(*bool)
    GetIsDeleted()(*bool)
    GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable)
    GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable)
    GetSubject()(*string)
    SetAttachments(value []Attachmentable)()
    SetBody(value ItemBodyable)()
    SetExtensions(value []Extensionable)()
    SetHasAttachments(value *bool)()
    SetIsDeleted(value *bool)()
    SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)()
    SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)()
    SetSubject(value *string)()
}
