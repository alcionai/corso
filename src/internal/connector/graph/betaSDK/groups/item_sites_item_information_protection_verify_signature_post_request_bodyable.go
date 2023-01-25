package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSitesItemInformationProtectionVerifySignaturePostRequestBodyable 
type ItemSitesItemInformationProtectionVerifySignaturePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    IBackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(BackingStore)
    GetDigest()([]byte)
    GetSignature()([]byte)
    GetSigningKeyId()(*string)
    SetBackingStore(value BackingStore)()
    SetDigest(value []byte)()
    SetSignature(value []byte)()
    SetSigningKeyId(value *string)()
}
