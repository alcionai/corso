package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerImportedPFXCertificateProfileable 
type AndroidDeviceOwnerImportedPFXCertificateProfileable interface {
    AndroidDeviceOwnerCertificateProfileBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificateAccessType()(*AndroidDeviceOwnerCertificateAccessType)
    GetIntendedPurpose()(*IntendedPurpose)
    GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable)
    GetSilentCertificateAccessDetails()([]AndroidDeviceOwnerSilentCertificateAccessable)
    SetCertificateAccessType(value *AndroidDeviceOwnerCertificateAccessType)()
    SetIntendedPurpose(value *IntendedPurpose)()
    SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)()
    SetSilentCertificateAccessDetails(value []AndroidDeviceOwnerSilentCertificateAccessable)()
}
