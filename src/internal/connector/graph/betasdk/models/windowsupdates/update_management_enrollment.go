package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UpdateManagementEnrollment 
type UpdateManagementEnrollment struct {
    UpdatableAssetEnrollment
    // The updateCategory property
    updateCategory *UpdateCategory
}
// NewUpdateManagementEnrollment instantiates a new UpdateManagementEnrollment and sets the default values.
func NewUpdateManagementEnrollment()(*UpdateManagementEnrollment) {
    m := &UpdateManagementEnrollment{
        UpdatableAssetEnrollment: *NewUpdatableAssetEnrollment(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdates.updateManagementEnrollment";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUpdateManagementEnrollmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdateManagementEnrollmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdateManagementEnrollment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdateManagementEnrollment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UpdatableAssetEnrollment.GetFieldDeserializers()
    res["updateCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUpdateCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateCategory(val.(*UpdateCategory))
        }
        return nil
    }
    return res
}
// GetUpdateCategory gets the updateCategory property value. The updateCategory property
func (m *UpdateManagementEnrollment) GetUpdateCategory()(*UpdateCategory) {
    return m.updateCategory
}
// Serialize serializes information the current object
func (m *UpdateManagementEnrollment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UpdatableAssetEnrollment.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetUpdateCategory() != nil {
        cast := (*m.GetUpdateCategory()).String()
        err = writer.WriteStringValue("updateCategory", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUpdateCategory sets the updateCategory property value. The updateCategory property
func (m *UpdateManagementEnrollment) SetUpdateCategory(value *UpdateCategory)() {
    m.updateCategory = value
}
