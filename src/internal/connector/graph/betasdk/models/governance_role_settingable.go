package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GovernanceRoleSettingable 
type GovernanceRoleSettingable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdminEligibleSettings()([]GovernanceRuleSettingable)
    GetAdminMemberSettings()([]GovernanceRuleSettingable)
    GetIsDefault()(*bool)
    GetLastUpdatedBy()(*string)
    GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetResource()(GovernanceResourceable)
    GetResourceId()(*string)
    GetRoleDefinition()(GovernanceRoleDefinitionable)
    GetRoleDefinitionId()(*string)
    GetUserEligibleSettings()([]GovernanceRuleSettingable)
    GetUserMemberSettings()([]GovernanceRuleSettingable)
    SetAdminEligibleSettings(value []GovernanceRuleSettingable)()
    SetAdminMemberSettings(value []GovernanceRuleSettingable)()
    SetIsDefault(value *bool)()
    SetLastUpdatedBy(value *string)()
    SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetResource(value GovernanceResourceable)()
    SetResourceId(value *string)()
    SetRoleDefinition(value GovernanceRoleDefinitionable)()
    SetRoleDefinitionId(value *string)()
    SetUserEligibleSettings(value []GovernanceRuleSettingable)()
    SetUserMemberSettings(value []GovernanceRuleSettingable)()
}
