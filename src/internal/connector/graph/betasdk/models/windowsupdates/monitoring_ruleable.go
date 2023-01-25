package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MonitoringRuleable 
type MonitoringRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAction()(*MonitoringAction)
    GetOdataType()(*string)
    GetSignal()(*MonitoringSignal)
    GetThreshold()(*int32)
    SetAction(value *MonitoringAction)()
    SetOdataType(value *string)()
    SetSignal(value *MonitoringSignal)()
    SetThreshold(value *int32)()
}
