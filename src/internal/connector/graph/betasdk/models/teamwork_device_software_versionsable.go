package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkDeviceSoftwareVersionsable 
type TeamworkDeviceSoftwareVersionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdminAgentSoftwareVersion()(*string)
    GetFirmwareSoftwareVersion()(*string)
    GetOdataType()(*string)
    GetOperatingSystemSoftwareVersion()(*string)
    GetPartnerAgentSoftwareVersion()(*string)
    GetTeamsClientSoftwareVersion()(*string)
    SetAdminAgentSoftwareVersion(value *string)()
    SetFirmwareSoftwareVersion(value *string)()
    SetOdataType(value *string)()
    SetOperatingSystemSoftwareVersion(value *string)()
    SetPartnerAgentSoftwareVersion(value *string)()
    SetTeamsClientSoftwareVersion(value *string)()
}
