package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LogicAppTriggerEndpointConfiguration 
type LogicAppTriggerEndpointConfiguration struct {
    CustomExtensionEndpointConfiguration
    // The name of the logic app.
    logicAppWorkflowName *string
    // The Azure resource group name for the logic app.
    resourceGroupName *string
    // Identifier of the Azure subscription for the logic app.
    subscriptionId *string
}
// NewLogicAppTriggerEndpointConfiguration instantiates a new LogicAppTriggerEndpointConfiguration and sets the default values.
func NewLogicAppTriggerEndpointConfiguration()(*LogicAppTriggerEndpointConfiguration) {
    m := &LogicAppTriggerEndpointConfiguration{
        CustomExtensionEndpointConfiguration: *NewCustomExtensionEndpointConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.logicAppTriggerEndpointConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateLogicAppTriggerEndpointConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLogicAppTriggerEndpointConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLogicAppTriggerEndpointConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LogicAppTriggerEndpointConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomExtensionEndpointConfiguration.GetFieldDeserializers()
    res["logicAppWorkflowName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogicAppWorkflowName(val)
        }
        return nil
    }
    res["resourceGroupName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceGroupName(val)
        }
        return nil
    }
    res["subscriptionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptionId(val)
        }
        return nil
    }
    return res
}
// GetLogicAppWorkflowName gets the logicAppWorkflowName property value. The name of the logic app.
func (m *LogicAppTriggerEndpointConfiguration) GetLogicAppWorkflowName()(*string) {
    return m.logicAppWorkflowName
}
// GetResourceGroupName gets the resourceGroupName property value. The Azure resource group name for the logic app.
func (m *LogicAppTriggerEndpointConfiguration) GetResourceGroupName()(*string) {
    return m.resourceGroupName
}
// GetSubscriptionId gets the subscriptionId property value. Identifier of the Azure subscription for the logic app.
func (m *LogicAppTriggerEndpointConfiguration) GetSubscriptionId()(*string) {
    return m.subscriptionId
}
// Serialize serializes information the current object
func (m *LogicAppTriggerEndpointConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomExtensionEndpointConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("logicAppWorkflowName", m.GetLogicAppWorkflowName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceGroupName", m.GetResourceGroupName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subscriptionId", m.GetSubscriptionId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLogicAppWorkflowName sets the logicAppWorkflowName property value. The name of the logic app.
func (m *LogicAppTriggerEndpointConfiguration) SetLogicAppWorkflowName(value *string)() {
    m.logicAppWorkflowName = value
}
// SetResourceGroupName sets the resourceGroupName property value. The Azure resource group name for the logic app.
func (m *LogicAppTriggerEndpointConfiguration) SetResourceGroupName(value *string)() {
    m.resourceGroupName = value
}
// SetSubscriptionId sets the subscriptionId property value. Identifier of the Azure subscription for the logic app.
func (m *LogicAppTriggerEndpointConfiguration) SetSubscriptionId(value *string)() {
    m.subscriptionId = value
}
