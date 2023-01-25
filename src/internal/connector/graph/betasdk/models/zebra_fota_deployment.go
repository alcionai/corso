package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ZebraFotaDeployment the Zebra FOTA deployment entity that describes settings, deployment device groups required to create a FOTA deployment, and deployment status.
type ZebraFotaDeployment struct {
    Entity
    // Collection of Android FOTA Assignment
    deploymentAssignments []AndroidFotaDeploymentAssignmentable
    // The Zebra FOTA deployment complex type that describes the settings required to create a FOTA deployment.
    deploymentSettings ZebraFotaDeploymentSettingsable
    // Represents the deployment status from Zebra. The status is a high level status of the deployment as opposed being a detailed status per device.
    deploymentStatus ZebraFotaDeploymentStatusable
    // A human readable description of the deployment.
    description *string
    // A human readable name of the deployment.
    displayName *string
}
// NewZebraFotaDeployment instantiates a new zebraFotaDeployment and sets the default values.
func NewZebraFotaDeployment()(*ZebraFotaDeployment) {
    m := &ZebraFotaDeployment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateZebraFotaDeploymentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateZebraFotaDeploymentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewZebraFotaDeployment(), nil
}
// GetDeploymentAssignments gets the deploymentAssignments property value. Collection of Android FOTA Assignment
func (m *ZebraFotaDeployment) GetDeploymentAssignments()([]AndroidFotaDeploymentAssignmentable) {
    return m.deploymentAssignments
}
// GetDeploymentSettings gets the deploymentSettings property value. The Zebra FOTA deployment complex type that describes the settings required to create a FOTA deployment.
func (m *ZebraFotaDeployment) GetDeploymentSettings()(ZebraFotaDeploymentSettingsable) {
    return m.deploymentSettings
}
// GetDeploymentStatus gets the deploymentStatus property value. Represents the deployment status from Zebra. The status is a high level status of the deployment as opposed being a detailed status per device.
func (m *ZebraFotaDeployment) GetDeploymentStatus()(ZebraFotaDeploymentStatusable) {
    return m.deploymentStatus
}
// GetDescription gets the description property value. A human readable description of the deployment.
func (m *ZebraFotaDeployment) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. A human readable name of the deployment.
func (m *ZebraFotaDeployment) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ZebraFotaDeployment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deploymentAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidFotaDeploymentAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidFotaDeploymentAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(AndroidFotaDeploymentAssignmentable)
            }
            m.SetDeploymentAssignments(res)
        }
        return nil
    }
    res["deploymentSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateZebraFotaDeploymentSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentSettings(val.(ZebraFotaDeploymentSettingsable))
        }
        return nil
    }
    res["deploymentStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateZebraFotaDeploymentStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentStatus(val.(ZebraFotaDeploymentStatusable))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ZebraFotaDeployment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDeploymentAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeploymentAssignments()))
        for i, v := range m.GetDeploymentAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deploymentAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deploymentSettings", m.GetDeploymentSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deploymentStatus", m.GetDeploymentStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeploymentAssignments sets the deploymentAssignments property value. Collection of Android FOTA Assignment
func (m *ZebraFotaDeployment) SetDeploymentAssignments(value []AndroidFotaDeploymentAssignmentable)() {
    m.deploymentAssignments = value
}
// SetDeploymentSettings sets the deploymentSettings property value. The Zebra FOTA deployment complex type that describes the settings required to create a FOTA deployment.
func (m *ZebraFotaDeployment) SetDeploymentSettings(value ZebraFotaDeploymentSettingsable)() {
    m.deploymentSettings = value
}
// SetDeploymentStatus sets the deploymentStatus property value. Represents the deployment status from Zebra. The status is a high level status of the deployment as opposed being a detailed status per device.
func (m *ZebraFotaDeployment) SetDeploymentStatus(value ZebraFotaDeploymentStatusable)() {
    m.deploymentStatus = value
}
// SetDescription sets the description property value. A human readable description of the deployment.
func (m *ZebraFotaDeployment) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. A human readable name of the deployment.
func (m *ZebraFotaDeployment) SetDisplayName(value *string)() {
    m.displayName = value
}
