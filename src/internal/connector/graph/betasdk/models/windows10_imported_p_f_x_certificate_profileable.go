package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10ImportedPFXCertificateProfileable 
type Windows10ImportedPFXCertificateProfileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WindowsCertificateProfileBaseable
    GetIntendedPurpose()(*IntendedPurpose)
    GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable)
    SetIntendedPurpose(value *IntendedPurpose)()
    SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)()
}
