package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicySettingMappingable 
type GroupPolicySettingMappingable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdmxSettingDefinitionId()(*string)
    GetChildIdList()([]string)
    GetIntuneSettingDefinitionId()(*string)
    GetIntuneSettingUriList()([]string)
    GetIsMdmSupported()(*bool)
    GetMdmCspName()(*string)
    GetMdmMinimumOSVersion()(*int32)
    GetMdmSettingUri()(*string)
    GetMdmSupportedState()(*MdmSupportedState)
    GetParentId()(*string)
    GetSettingCategory()(*string)
    GetSettingDisplayName()(*string)
    GetSettingDisplayValue()(*string)
    GetSettingDisplayValueType()(*string)
    GetSettingName()(*string)
    GetSettingScope()(*GroupPolicySettingScope)
    GetSettingType()(*GroupPolicySettingType)
    GetSettingValue()(*string)
    GetSettingValueDisplayUnits()(*string)
    GetSettingValueType()(*string)
    SetAdmxSettingDefinitionId(value *string)()
    SetChildIdList(value []string)()
    SetIntuneSettingDefinitionId(value *string)()
    SetIntuneSettingUriList(value []string)()
    SetIsMdmSupported(value *bool)()
    SetMdmCspName(value *string)()
    SetMdmMinimumOSVersion(value *int32)()
    SetMdmSettingUri(value *string)()
    SetMdmSupportedState(value *MdmSupportedState)()
    SetParentId(value *string)()
    SetSettingCategory(value *string)()
    SetSettingDisplayName(value *string)()
    SetSettingDisplayValue(value *string)()
    SetSettingDisplayValueType(value *string)()
    SetSettingName(value *string)()
    SetSettingScope(value *GroupPolicySettingScope)()
    SetSettingType(value *GroupPolicySettingType)()
    SetSettingValue(value *string)()
    SetSettingValueDisplayUnits(value *string)()
    SetSettingValueType(value *string)()
}
