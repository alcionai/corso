package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BusinessScenarioGroupTarget 
type BusinessScenarioGroupTarget struct {
    BusinessScenarioTaskTargetBase
    // The unique identifier for the group.
    groupId *string
}
// NewBusinessScenarioGroupTarget instantiates a new BusinessScenarioGroupTarget and sets the default values.
func NewBusinessScenarioGroupTarget()(*BusinessScenarioGroupTarget) {
    m := &BusinessScenarioGroupTarget{
        BusinessScenarioTaskTargetBase: *NewBusinessScenarioTaskTargetBase(),
    }
    odataTypeValue := "#microsoft.graph.businessScenarioGroupTarget";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateBusinessScenarioGroupTargetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBusinessScenarioGroupTargetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBusinessScenarioGroupTarget(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BusinessScenarioGroupTarget) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BusinessScenarioTaskTargetBase.GetFieldDeserializers()
    res["groupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupId(val)
        }
        return nil
    }
    return res
}
// GetGroupId gets the groupId property value. The unique identifier for the group.
func (m *BusinessScenarioGroupTarget) GetGroupId()(*string) {
    return m.groupId
}
// Serialize serializes information the current object
func (m *BusinessScenarioGroupTarget) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BusinessScenarioTaskTargetBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("groupId", m.GetGroupId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroupId sets the groupId property value. The unique identifier for the group.
func (m *BusinessScenarioGroupTarget) SetGroupId(value *string)() {
    m.groupId = value
}
