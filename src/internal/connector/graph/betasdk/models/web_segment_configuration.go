package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WebSegmentConfiguration 
type WebSegmentConfiguration struct {
    SegmentConfiguration
    // The applicationSegments property
    applicationSegments []WebApplicationSegmentable
}
// NewWebSegmentConfiguration instantiates a new WebSegmentConfiguration and sets the default values.
func NewWebSegmentConfiguration()(*WebSegmentConfiguration) {
    m := &WebSegmentConfiguration{
        SegmentConfiguration: *NewSegmentConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.webSegmentConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWebSegmentConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWebSegmentConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWebSegmentConfiguration(), nil
}
// GetApplicationSegments gets the applicationSegments property value. The applicationSegments property
func (m *WebSegmentConfiguration) GetApplicationSegments()([]WebApplicationSegmentable) {
    return m.applicationSegments
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WebSegmentConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SegmentConfiguration.GetFieldDeserializers()
    res["applicationSegments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWebApplicationSegmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WebApplicationSegmentable, len(val))
            for i, v := range val {
                res[i] = v.(WebApplicationSegmentable)
            }
            m.SetApplicationSegments(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *WebSegmentConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SegmentConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApplicationSegments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetApplicationSegments()))
        for i, v := range m.GetApplicationSegments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("applicationSegments", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicationSegments sets the applicationSegments property value. The applicationSegments property
func (m *WebSegmentConfiguration) SetApplicationSegments(value []WebApplicationSegmentable)() {
    m.applicationSegments = value
}
