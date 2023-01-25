package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnPremisesDirectorySynchronizationConfigurationable 
type OnPremisesDirectorySynchronizationConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccidentalDeletionPrevention()(OnPremisesAccidentalDeletionPreventionable)
    GetCustomerRequestedSynchronizationInterval()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetOdataType()(*string)
    GetSynchronizationInterval()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    SetAccidentalDeletionPrevention(value OnPremisesAccidentalDeletionPreventionable)()
    SetCustomerRequestedSynchronizationInterval(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetOdataType(value *string)()
    SetSynchronizationInterval(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
}
