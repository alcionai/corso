package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcPartnerAgentInstallResultable 
type CloudPcPartnerAgentInstallResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInstallStatus()(*CloudPcPartnerAgentInstallStatus)
    GetIsThirdPartyPartner()(*bool)
    GetOdataType()(*string)
    GetPartnerAgentName()(*CloudPcPartnerAgentName)
    GetRetriable()(*bool)
    SetInstallStatus(value *CloudPcPartnerAgentInstallStatus)()
    SetIsThirdPartyPartner(value *bool)()
    SetOdataType(value *string)()
    SetPartnerAgentName(value *CloudPcPartnerAgentName)()
    SetRetriable(value *bool)()
}
