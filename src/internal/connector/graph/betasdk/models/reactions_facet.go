package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReactionsFacet 
type ReactionsFacet struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Count of comments.
    commentCount *int32
    // Count of likes.
    likeCount *int32
    // The OdataType property
    odataType *string
    // Count of shares.
    shareCount *int32
}
// NewReactionsFacet instantiates a new reactionsFacet and sets the default values.
func NewReactionsFacet()(*ReactionsFacet) {
    m := &ReactionsFacet{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateReactionsFacetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReactionsFacetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReactionsFacet(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReactionsFacet) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCommentCount gets the commentCount property value. Count of comments.
func (m *ReactionsFacet) GetCommentCount()(*int32) {
    return m.commentCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReactionsFacet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["commentCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommentCount(val)
        }
        return nil
    }
    res["likeCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLikeCount(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["shareCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShareCount(val)
        }
        return nil
    }
    return res
}
// GetLikeCount gets the likeCount property value. Count of likes.
func (m *ReactionsFacet) GetLikeCount()(*int32) {
    return m.likeCount
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ReactionsFacet) GetOdataType()(*string) {
    return m.odataType
}
// GetShareCount gets the shareCount property value. Count of shares.
func (m *ReactionsFacet) GetShareCount()(*int32) {
    return m.shareCount
}
// Serialize serializes information the current object
func (m *ReactionsFacet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("commentCount", m.GetCommentCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("likeCount", m.GetLikeCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("shareCount", m.GetShareCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReactionsFacet) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCommentCount sets the commentCount property value. Count of comments.
func (m *ReactionsFacet) SetCommentCount(value *int32)() {
    m.commentCount = value
}
// SetLikeCount sets the likeCount property value. Count of likes.
func (m *ReactionsFacet) SetLikeCount(value *int32)() {
    m.likeCount = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ReactionsFacet) SetOdataType(value *string)() {
    m.odataType = value
}
// SetShareCount sets the shareCount property value. Count of shares.
func (m *ReactionsFacet) SetShareCount(value *int32)() {
    m.shareCount = value
}
