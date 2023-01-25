package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsAutopilotDeploymentProfilePolicySetItemCollectionResponse 
type WindowsAutopilotDeploymentProfilePolicySetItemCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []WindowsAutopilotDeploymentProfilePolicySetItemable
}
// NewWindowsAutopilotDeploymentProfilePolicySetItemCollectionResponse instantiates a new WindowsAutopilotDeploymentProfilePolicySetItemCollectionResponse and sets the default values.
func NewWindowsAutopilotDeploymentProfilePolicySetItemCollectionResponse()(*WindowsAutopilotDeploymentProfilePolicySetItemCollectionResponse) {
    m := &WindowsAutopilotDeploymentProfilePolicySetItemCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateWindowsAutopilotDeploymentProfilePolicySetItemCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsAutopilotDeploymentProfilePolicySetItemCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsAutopilotDeploymentProfilePolicySetItemCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsAutopilotDeploymentProfilePolicySetItemCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsAutopilotDeploymentProfilePolicySetItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsAutopilotDeploymentProfilePolicySetItemable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsAutopilotDeploymentProfilePolicySetItemable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *WindowsAutopilotDeploymentProfilePolicySetItemCollectionResponse) GetValue()([]WindowsAutopilotDeploymentProfilePolicySetItemable) {
    return m.value
}
// Serialize serializes information the current object
func (m *WindowsAutopilotDeploymentProfilePolicySetItemCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *WindowsAutopilotDeploymentProfilePolicySetItemCollectionResponse) SetValue(value []WindowsAutopilotDeploymentProfilePolicySetItemable)() {
    m.value = value
}
