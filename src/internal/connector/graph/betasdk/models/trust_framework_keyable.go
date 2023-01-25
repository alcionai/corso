package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TrustFrameworkKeyable 
type TrustFrameworkKeyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetD()(*string)
    GetDp()(*string)
    GetDq()(*string)
    GetE()(*string)
    GetExp()(*int64)
    GetK()(*string)
    GetKid()(*string)
    GetKty()(*string)
    GetN()(*string)
    GetNbf()(*int64)
    GetOdataType()(*string)
    GetP()(*string)
    GetQ()(*string)
    GetQi()(*string)
    GetUse()(*string)
    GetX5c()([]string)
    GetX5t()(*string)
    SetD(value *string)()
    SetDp(value *string)()
    SetDq(value *string)()
    SetE(value *string)()
    SetExp(value *int64)()
    SetK(value *string)()
    SetKid(value *string)()
    SetKty(value *string)()
    SetN(value *string)()
    SetNbf(value *int64)()
    SetOdataType(value *string)()
    SetP(value *string)()
    SetQ(value *string)()
    SetQi(value *string)()
    SetUse(value *string)()
    SetX5c(value []string)()
    SetX5t(value *string)()
}
