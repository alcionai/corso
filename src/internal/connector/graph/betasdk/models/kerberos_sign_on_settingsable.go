package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// KerberosSignOnSettingsable 
type KerberosSignOnSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKerberosServicePrincipalName()(*string)
    GetKerberosSignOnMappingAttributeType()(*KerberosSignOnMappingAttributeType)
    GetOdataType()(*string)
    SetKerberosServicePrincipalName(value *string)()
    SetKerberosSignOnMappingAttributeType(value *KerberosSignOnMappingAttributeType)()
    SetOdataType(value *string)()
}
