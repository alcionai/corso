package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationCsvDataProvider 
type EducationCsvDataProvider struct {
    EducationSynchronizationDataProvider
    // Optional customizations to be applied to the synchronization profile.
    customizations EducationSynchronizationCustomizationsable
}
// NewEducationCsvDataProvider instantiates a new EducationCsvDataProvider and sets the default values.
func NewEducationCsvDataProvider()(*EducationCsvDataProvider) {
    m := &EducationCsvDataProvider{
        EducationSynchronizationDataProvider: *NewEducationSynchronizationDataProvider(),
    }
    odataTypeValue := "#microsoft.graph.educationCsvDataProvider";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateEducationCsvDataProviderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationCsvDataProviderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationCsvDataProvider(), nil
}
// GetCustomizations gets the customizations property value. Optional customizations to be applied to the synchronization profile.
func (m *EducationCsvDataProvider) GetCustomizations()(EducationSynchronizationCustomizationsable) {
    return m.customizations
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationCsvDataProvider) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EducationSynchronizationDataProvider.GetFieldDeserializers()
    res["customizations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationSynchronizationCustomizationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomizations(val.(EducationSynchronizationCustomizationsable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *EducationCsvDataProvider) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EducationSynchronizationDataProvider.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("customizations", m.GetCustomizations())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustomizations sets the customizations property value. Optional customizations to be applied to the synchronization profile.
func (m *EducationCsvDataProvider) SetCustomizations(value EducationSynchronizationCustomizationsable)() {
    m.customizations = value
}
