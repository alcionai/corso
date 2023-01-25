package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LabelDetails 
type LabelDetails struct {
    ParentLabelDetails
}
// NewLabelDetails instantiates a new LabelDetails and sets the default values.
func NewLabelDetails()(*LabelDetails) {
    m := &LabelDetails{
        ParentLabelDetails: *NewParentLabelDetails(),
    }
    odataTypeValue := "#microsoft.graph.labelDetails";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateLabelDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLabelDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLabelDetails(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LabelDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ParentLabelDetails.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *LabelDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ParentLabelDetails.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
