package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementSettingDefinitionable 
type DeviceManagementSettingDefinitionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConstraints()([]DeviceManagementConstraintable)
    GetDependencies()([]DeviceManagementSettingDependencyable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetDocumentationUrl()(*string)
    GetHeaderSubtitle()(*string)
    GetHeaderTitle()(*string)
    GetIsTopLevel()(*bool)
    GetKeywords()([]string)
    GetPlaceholderText()(*string)
    GetValueType()(*DeviceManangementIntentValueType)
    SetConstraints(value []DeviceManagementConstraintable)()
    SetDependencies(value []DeviceManagementSettingDependencyable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetDocumentationUrl(value *string)()
    SetHeaderSubtitle(value *string)()
    SetHeaderTitle(value *string)()
    SetIsTopLevel(value *bool)()
    SetKeywords(value []string)()
    SetPlaceholderText(value *string)()
    SetValueType(value *DeviceManangementIntentValueType)()
}
