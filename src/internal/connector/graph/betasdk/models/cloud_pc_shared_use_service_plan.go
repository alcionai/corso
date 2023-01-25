package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcSharedUseServicePlan 
type CloudPcSharedUseServicePlan struct {
    Entity
    // The displayName property
    displayName *string
    // The totalCount property
    totalCount *int32
    // The usedCount property
    usedCount *int32
}
// NewCloudPcSharedUseServicePlan instantiates a new CloudPcSharedUseServicePlan and sets the default values.
func NewCloudPcSharedUseServicePlan()(*CloudPcSharedUseServicePlan) {
    m := &CloudPcSharedUseServicePlan{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPcSharedUseServicePlanFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcSharedUseServicePlanFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcSharedUseServicePlan(), nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *CloudPcSharedUseServicePlan) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcSharedUseServicePlan) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["totalCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCount(val)
        }
        return nil
    }
    res["usedCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsedCount(val)
        }
        return nil
    }
    return res
}
// GetTotalCount gets the totalCount property value. The totalCount property
func (m *CloudPcSharedUseServicePlan) GetTotalCount()(*int32) {
    return m.totalCount
}
// GetUsedCount gets the usedCount property value. The usedCount property
func (m *CloudPcSharedUseServicePlan) GetUsedCount()(*int32) {
    return m.usedCount
}
// Serialize serializes information the current object
func (m *CloudPcSharedUseServicePlan) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalCount", m.GetTotalCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("usedCount", m.GetUsedCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *CloudPcSharedUseServicePlan) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetTotalCount sets the totalCount property value. The totalCount property
func (m *CloudPcSharedUseServicePlan) SetTotalCount(value *int32)() {
    m.totalCount = value
}
// SetUsedCount sets the usedCount property value. The usedCount property
func (m *CloudPcSharedUseServicePlan) SetUsedCount(value *int32)() {
    m.usedCount = value
}
