package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FocusActivityStatistics 
type FocusActivityStatistics struct {
    ActivityStatistics
}
// NewFocusActivityStatistics instantiates a new FocusActivityStatistics and sets the default values.
func NewFocusActivityStatistics()(*FocusActivityStatistics) {
    m := &FocusActivityStatistics{
        ActivityStatistics: *NewActivityStatistics(),
    }
    odataTypeValue := "#microsoft.graph.focusActivityStatistics";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateFocusActivityStatisticsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFocusActivityStatisticsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFocusActivityStatistics(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FocusActivityStatistics) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ActivityStatistics.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *FocusActivityStatistics) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ActivityStatistics.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
