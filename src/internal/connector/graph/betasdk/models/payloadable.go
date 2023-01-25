package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Payloadable 
type Payloadable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBrand()(*PayloadBrand)
    GetComplexity()(*PayloadComplexity)
    GetCreatedBy()(EmailIdentityable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDetail()(PayloadDetailable)
    GetDisplayName()(*string)
    GetIndustry()(*PayloadIndustry)
    GetIsAutomated()(*bool)
    GetIsControversial()(*bool)
    GetIsCurrentEvent()(*bool)
    GetLanguage()(*string)
    GetLastModifiedBy()(EmailIdentityable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPayloadTags()([]string)
    GetPlatform()(*PayloadDeliveryPlatform)
    GetPredictedCompromiseRate()(*float64)
    GetSimulationAttackType()(*SimulationAttackType)
    GetSource()(*SimulationContentSource)
    GetStatus()(*SimulationContentStatus)
    GetTechnique()(*SimulationAttackTechnique)
    GetTheme()(*PayloadTheme)
    SetBrand(value *PayloadBrand)()
    SetComplexity(value *PayloadComplexity)()
    SetCreatedBy(value EmailIdentityable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDetail(value PayloadDetailable)()
    SetDisplayName(value *string)()
    SetIndustry(value *PayloadIndustry)()
    SetIsAutomated(value *bool)()
    SetIsControversial(value *bool)()
    SetIsCurrentEvent(value *bool)()
    SetLanguage(value *string)()
    SetLastModifiedBy(value EmailIdentityable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPayloadTags(value []string)()
    SetPlatform(value *PayloadDeliveryPlatform)()
    SetPredictedCompromiseRate(value *float64)()
    SetSimulationAttackType(value *SimulationAttackType)()
    SetSource(value *SimulationContentSource)()
    SetStatus(value *SimulationContentStatus)()
    SetTechnique(value *SimulationAttackTechnique)()
    SetTheme(value *PayloadTheme)()
}
