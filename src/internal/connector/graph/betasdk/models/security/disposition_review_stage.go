package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// DispositionReviewStage provides operations to manage the collection of site entities.
type DispositionReviewStage struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // Name representing each stage within a collection.
    name *string
    // A collection of reviewers at each stage.
    reviewersEmailAddresses []string
    // The sequence number for each stage of the disposition review.
    stageNumber *int32
}
// NewDispositionReviewStage instantiates a new dispositionReviewStage and sets the default values.
func NewDispositionReviewStage()(*DispositionReviewStage) {
    m := &DispositionReviewStage{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateDispositionReviewStageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDispositionReviewStageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDispositionReviewStage(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DispositionReviewStage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["reviewersEmailAddresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetReviewersEmailAddresses(res)
        }
        return nil
    }
    res["stageNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStageNumber(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Name representing each stage within a collection.
func (m *DispositionReviewStage) GetName()(*string) {
    return m.name
}
// GetReviewersEmailAddresses gets the reviewersEmailAddresses property value. A collection of reviewers at each stage.
func (m *DispositionReviewStage) GetReviewersEmailAddresses()([]string) {
    return m.reviewersEmailAddresses
}
// GetStageNumber gets the stageNumber property value. The sequence number for each stage of the disposition review.
func (m *DispositionReviewStage) GetStageNumber()(*int32) {
    return m.stageNumber
}
// Serialize serializes information the current object
func (m *DispositionReviewStage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetReviewersEmailAddresses() != nil {
        err = writer.WriteCollectionOfStringValues("reviewersEmailAddresses", m.GetReviewersEmailAddresses())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("stageNumber", m.GetStageNumber())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetName sets the name property value. Name representing each stage within a collection.
func (m *DispositionReviewStage) SetName(value *string)() {
    m.name = value
}
// SetReviewersEmailAddresses sets the reviewersEmailAddresses property value. A collection of reviewers at each stage.
func (m *DispositionReviewStage) SetReviewersEmailAddresses(value []string)() {
    m.reviewersEmailAddresses = value
}
// SetStageNumber sets the stageNumber property value. The sequence number for each stage of the disposition review.
func (m *DispositionReviewStage) SetStageNumber(value *int32)() {
    m.stageNumber = value
}
