package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementTroubleshootingErrorDetailsable 
type DeviceManagementTroubleshootingErrorDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContext()(*string)
    GetFailure()(*string)
    GetFailureDetails()(*string)
    GetOdataType()(*string)
    GetRemediation()(*string)
    GetResources()([]DeviceManagementTroubleshootingErrorResourceable)
    SetContext(value *string)()
    SetFailure(value *string)()
    SetFailureDetails(value *string)()
    SetOdataType(value *string)()
    SetRemediation(value *string)()
    SetResources(value []DeviceManagementTroubleshootingErrorResourceable)()
}
