package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExcludedAppsable 
type ExcludedAppsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccess()(*bool)
    GetBing()(*bool)
    GetExcel()(*bool)
    GetGroove()(*bool)
    GetInfoPath()(*bool)
    GetLync()(*bool)
    GetOdataType()(*string)
    GetOneDrive()(*bool)
    GetOneNote()(*bool)
    GetOutlook()(*bool)
    GetPowerPoint()(*bool)
    GetPublisher()(*bool)
    GetSharePointDesigner()(*bool)
    GetTeams()(*bool)
    GetVisio()(*bool)
    GetWord()(*bool)
    SetAccess(value *bool)()
    SetBing(value *bool)()
    SetExcel(value *bool)()
    SetGroove(value *bool)()
    SetInfoPath(value *bool)()
    SetLync(value *bool)()
    SetOdataType(value *string)()
    SetOneDrive(value *bool)()
    SetOneNote(value *bool)()
    SetOutlook(value *bool)()
    SetPowerPoint(value *bool)()
    SetPublisher(value *bool)()
    SetSharePointDesigner(value *bool)()
    SetTeams(value *bool)()
    SetVisio(value *bool)()
    SetWord(value *bool)()
}
