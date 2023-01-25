package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// DeviceCompliancePolicySettingStateSummaryable 
type DeviceCompliancePolicySettingStateSummaryable interface {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConflictDeviceCount()(*int32)
    GetErrorDeviceCount()(*int32)
    GetFailedDeviceCount()(*int32)
    GetIntuneAccountId()(*string)
    GetIntuneSettingId()(*string)
    GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNotApplicableDeviceCount()(*int32)
    GetPendingDeviceCount()(*int32)
    GetPolicyType()(*string)
    GetSettingName()(*string)
    GetSucceededDeviceCount()(*int32)
    GetTenantDisplayName()(*string)
    GetTenantId()(*string)
    SetConflictDeviceCount(value *int32)()
    SetErrorDeviceCount(value *int32)()
    SetFailedDeviceCount(value *int32)()
    SetIntuneAccountId(value *string)()
    SetIntuneSettingId(value *string)()
    SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNotApplicableDeviceCount(value *int32)()
    SetPendingDeviceCount(value *int32)()
    SetPolicyType(value *string)()
    SetSettingName(value *string)()
    SetSucceededDeviceCount(value *int32)()
    SetTenantDisplayName(value *string)()
    SetTenantId(value *string)()
}
