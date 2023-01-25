package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceComplianceLocalActionBase 
type AndroidDeviceComplianceLocalActionBase struct {
    Entity
    // Number of minutes to wait till a local action is enforced. Valid values 0 to 2147483647
    gracePeriodInMinutes *int32
}
// NewAndroidDeviceComplianceLocalActionBase instantiates a new AndroidDeviceComplianceLocalActionBase and sets the default values.
func NewAndroidDeviceComplianceLocalActionBase()(*AndroidDeviceComplianceLocalActionBase) {
    m := &AndroidDeviceComplianceLocalActionBase{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAndroidDeviceComplianceLocalActionBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceComplianceLocalActionBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.androidDeviceComplianceLocalActionLockDevice":
                        return NewAndroidDeviceComplianceLocalActionLockDevice(), nil
                    case "#microsoft.graph.androidDeviceComplianceLocalActionLockDeviceWithPasscode":
                        return NewAndroidDeviceComplianceLocalActionLockDeviceWithPasscode(), nil
                }
            }
        }
    }
    return NewAndroidDeviceComplianceLocalActionBase(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceComplianceLocalActionBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["gracePeriodInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGracePeriodInMinutes(val)
        }
        return nil
    }
    return res
}
// GetGracePeriodInMinutes gets the gracePeriodInMinutes property value. Number of minutes to wait till a local action is enforced. Valid values 0 to 2147483647
func (m *AndroidDeviceComplianceLocalActionBase) GetGracePeriodInMinutes()(*int32) {
    return m.gracePeriodInMinutes
}
// Serialize serializes information the current object
func (m *AndroidDeviceComplianceLocalActionBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("gracePeriodInMinutes", m.GetGracePeriodInMinutes())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGracePeriodInMinutes sets the gracePeriodInMinutes property value. Number of minutes to wait till a local action is enforced. Valid values 0 to 2147483647
func (m *AndroidDeviceComplianceLocalActionBase) SetGracePeriodInMinutes(value *int32)() {
    m.gracePeriodInMinutes = value
}
