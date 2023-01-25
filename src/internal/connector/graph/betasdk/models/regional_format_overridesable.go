package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RegionalFormatOverridesable 
type RegionalFormatOverridesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCalendar()(*string)
    GetFirstDayOfWeek()(*string)
    GetLongDateFormat()(*string)
    GetLongTimeFormat()(*string)
    GetOdataType()(*string)
    GetShortDateFormat()(*string)
    GetShortTimeFormat()(*string)
    GetTimeZone()(*string)
    SetCalendar(value *string)()
    SetFirstDayOfWeek(value *string)()
    SetLongDateFormat(value *string)()
    SetLongTimeFormat(value *string)()
    SetOdataType(value *string)()
    SetShortDateFormat(value *string)()
    SetShortTimeFormat(value *string)()
    SetTimeZone(value *string)()
}
