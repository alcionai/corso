package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DocumentComment 
type DocumentComment struct {
    Entity
    // The content property
    content *string
    // The replies property
    replies []DocumentCommentReplyable
}
// NewDocumentComment instantiates a new DocumentComment and sets the default values.
func NewDocumentComment()(*DocumentComment) {
    m := &DocumentComment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDocumentCommentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDocumentCommentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDocumentComment(), nil
}
// GetContent gets the content property value. The content property
func (m *DocumentComment) GetContent()(*string) {
    return m.content
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DocumentComment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["replies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDocumentCommentReplyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DocumentCommentReplyable, len(val))
            for i, v := range val {
                res[i] = v.(DocumentCommentReplyable)
            }
            m.SetReplies(res)
        }
        return nil
    }
    return res
}
// GetReplies gets the replies property value. The replies property
func (m *DocumentComment) GetReplies()([]DocumentCommentReplyable) {
    return m.replies
}
// Serialize serializes information the current object
func (m *DocumentComment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    if m.GetReplies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReplies()))
        for i, v := range m.GetReplies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("replies", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContent sets the content property value. The content property
func (m *DocumentComment) SetContent(value *string)() {
    m.content = value
}
// SetReplies sets the replies property value. The replies property
func (m *DocumentComment) SetReplies(value []DocumentCommentReplyable)() {
    m.replies = value
}
