package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttackSimulationInfo 
type AttackSimulationInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The date and time of the attack simulation.
    attackSimDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The duration (in time) for the attack simulation.
    attackSimDurationTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The activity ID for the attack simulation.
    attackSimId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The unique identifier for the user who got the attack simulation email.
    attackSimUserId *string
    // The OdataType property
    odataType *string
}
// NewAttackSimulationInfo instantiates a new attackSimulationInfo and sets the default values.
func NewAttackSimulationInfo()(*AttackSimulationInfo) {
    m := &AttackSimulationInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAttackSimulationInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttackSimulationInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttackSimulationInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttackSimulationInfo) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAttackSimDateTime gets the attackSimDateTime property value. The date and time of the attack simulation.
func (m *AttackSimulationInfo) GetAttackSimDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.attackSimDateTime
}
// GetAttackSimDurationTime gets the attackSimDurationTime property value. The duration (in time) for the attack simulation.
func (m *AttackSimulationInfo) GetAttackSimDurationTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.attackSimDurationTime
}
// GetAttackSimId gets the attackSimId property value. The activity ID for the attack simulation.
func (m *AttackSimulationInfo) GetAttackSimId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.attackSimId
}
// GetAttackSimUserId gets the attackSimUserId property value. The unique identifier for the user who got the attack simulation email.
func (m *AttackSimulationInfo) GetAttackSimUserId()(*string) {
    return m.attackSimUserId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttackSimulationInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attackSimDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttackSimDateTime(val)
        }
        return nil
    }
    res["attackSimDurationTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttackSimDurationTime(val)
        }
        return nil
    }
    res["attackSimId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttackSimId(val)
        }
        return nil
    }
    res["attackSimUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttackSimUserId(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AttackSimulationInfo) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AttackSimulationInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("attackSimDateTime", m.GetAttackSimDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("attackSimDurationTime", m.GetAttackSimDurationTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("attackSimId", m.GetAttackSimId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("attackSimUserId", m.GetAttackSimUserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttackSimulationInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAttackSimDateTime sets the attackSimDateTime property value. The date and time of the attack simulation.
func (m *AttackSimulationInfo) SetAttackSimDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.attackSimDateTime = value
}
// SetAttackSimDurationTime sets the attackSimDurationTime property value. The duration (in time) for the attack simulation.
func (m *AttackSimulationInfo) SetAttackSimDurationTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.attackSimDurationTime = value
}
// SetAttackSimId sets the attackSimId property value. The activity ID for the attack simulation.
func (m *AttackSimulationInfo) SetAttackSimId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.attackSimId = value
}
// SetAttackSimUserId sets the attackSimUserId property value. The unique identifier for the user who got the attack simulation email.
func (m *AttackSimulationInfo) SetAttackSimUserId(value *string)() {
    m.attackSimUserId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AttackSimulationInfo) SetOdataType(value *string)() {
    m.odataType = value
}
