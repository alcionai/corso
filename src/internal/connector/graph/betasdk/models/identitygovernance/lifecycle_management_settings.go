package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// LifecycleManagementSettings provides operations to manage the collection of site entities.
type LifecycleManagementSettings struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entity
    // The interval in hours at which all workflows running in the tenant should be scheduled for execution. This interval has a minimum value of 1 and a maximum value of 24. The default value is 3 hours.
    workflowScheduleIntervalInHours *int32
}
// NewLifecycleManagementSettings instantiates a new lifecycleManagementSettings and sets the default values.
func NewLifecycleManagementSettings()(*LifecycleManagementSettings) {
    m := &LifecycleManagementSettings{
        Entity: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewEntity(),
    }
    return m
}
// CreateLifecycleManagementSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLifecycleManagementSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLifecycleManagementSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LifecycleManagementSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["workflowScheduleIntervalInHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkflowScheduleIntervalInHours(val)
        }
        return nil
    }
    return res
}
// GetWorkflowScheduleIntervalInHours gets the workflowScheduleIntervalInHours property value. The interval in hours at which all workflows running in the tenant should be scheduled for execution. This interval has a minimum value of 1 and a maximum value of 24. The default value is 3 hours.
func (m *LifecycleManagementSettings) GetWorkflowScheduleIntervalInHours()(*int32) {
    return m.workflowScheduleIntervalInHours
}
// Serialize serializes information the current object
func (m *LifecycleManagementSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("workflowScheduleIntervalInHours", m.GetWorkflowScheduleIntervalInHours())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetWorkflowScheduleIntervalInHours sets the workflowScheduleIntervalInHours property value. The interval in hours at which all workflows running in the tenant should be scheduled for execution. This interval has a minimum value of 1 and a maximum value of 24. The default value is 3 hours.
func (m *LifecycleManagementSettings) SetWorkflowScheduleIntervalInHours(value *int32)() {
    m.workflowScheduleIntervalInHours = value
}
