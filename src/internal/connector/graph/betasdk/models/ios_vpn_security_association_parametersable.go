package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosVpnSecurityAssociationParametersable 
type IosVpnSecurityAssociationParametersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLifetimeInMinutes()(*int32)
    GetOdataType()(*string)
    GetSecurityDiffieHellmanGroup()(*int32)
    GetSecurityEncryptionAlgorithm()(*VpnEncryptionAlgorithmType)
    GetSecurityIntegrityAlgorithm()(*VpnIntegrityAlgorithmType)
    SetLifetimeInMinutes(value *int32)()
    SetOdataType(value *string)()
    SetSecurityDiffieHellmanGroup(value *int32)()
    SetSecurityEncryptionAlgorithm(value *VpnEncryptionAlgorithmType)()
    SetSecurityIntegrityAlgorithm(value *VpnIntegrityAlgorithmType)()
}
