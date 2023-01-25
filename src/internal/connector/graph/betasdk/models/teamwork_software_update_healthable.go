package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkSoftwareUpdateHealthable 
type TeamworkSoftwareUpdateHealthable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdminAgentSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable)
    GetCompanyPortalSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable)
    GetFirmwareSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable)
    GetOdataType()(*string)
    GetOperatingSystemSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable)
    GetPartnerAgentSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable)
    GetTeamsClientSoftwareUpdateStatus()(TeamworkSoftwareUpdateStatusable)
    SetAdminAgentSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)()
    SetCompanyPortalSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)()
    SetFirmwareSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)()
    SetOdataType(value *string)()
    SetOperatingSystemSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)()
    SetPartnerAgentSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)()
    SetTeamsClientSoftwareUpdateStatus(value TeamworkSoftwareUpdateStatusable)()
}
