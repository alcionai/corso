package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// CredentialUserRegistrationsSummaryable 
type CredentialUserRegistrationsSummaryable interface {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMfaAndSsprCapableUserCount()(*int32)
    GetMfaConditionalAccessPolicyState()(*string)
    GetMfaExcludedUserCount()(*int32)
    GetMfaRegisteredUserCount()(*int32)
    GetSecurityDefaultsEnabled()(*bool)
    GetSsprEnabledUserCount()(*int32)
    GetSsprRegisteredUserCount()(*int32)
    GetTenantDisplayName()(*string)
    GetTenantId()(*string)
    GetTotalUserCount()(*int32)
    SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMfaAndSsprCapableUserCount(value *int32)()
    SetMfaConditionalAccessPolicyState(value *string)()
    SetMfaExcludedUserCount(value *int32)()
    SetMfaRegisteredUserCount(value *int32)()
    SetSecurityDefaultsEnabled(value *bool)()
    SetSsprEnabledUserCount(value *int32)()
    SetSsprRegisteredUserCount(value *int32)()
    SetTenantDisplayName(value *string)()
    SetTenantId(value *string)()
    SetTotalUserCount(value *int32)()
}
