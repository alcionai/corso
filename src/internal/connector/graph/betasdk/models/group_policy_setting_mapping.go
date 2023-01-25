package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicySettingMapping the Group Policy setting to MDM/Intune mapping.
type GroupPolicySettingMapping struct {
    Entity
    // Admx Group Policy Id
    admxSettingDefinitionId *string
    // List of Child Ids of the group policy setting.
    childIdList []string
    // The Intune Setting Definition Id
    intuneSettingDefinitionId *string
    // The list of Intune Setting URIs this group policy setting maps to
    intuneSettingUriList []string
    // Indicates if the setting is supported by Intune or not
    isMdmSupported *bool
    // The CSP name this group policy setting maps to.
    mdmCspName *string
    // The minimum OS version this mdm setting supports.
    mdmMinimumOSVersion *int32
    // The MDM CSP URI this group policy setting maps to.
    mdmSettingUri *string
    // Mdm Support Status of the setting.
    mdmSupportedState *MdmSupportedState
    // Parent Id of the group policy setting.
    parentId *string
    // The category the group policy setting is in.
    settingCategory *string
    // The display name of this group policy setting.
    settingDisplayName *string
    // The display value of this group policy setting.
    settingDisplayValue *string
    // The display value type of this group policy setting.
    settingDisplayValueType *string
    // The name of this group policy setting.
    settingName *string
    // Scope of the group policy setting.
    settingScope *GroupPolicySettingScope
    // Setting type of the group policy.
    settingType *GroupPolicySettingType
    // The value of this group policy setting.
    settingValue *string
    // The display units of this group policy setting value
    settingValueDisplayUnits *string
    // The value type of this group policy setting.
    settingValueType *string
}
// NewGroupPolicySettingMapping instantiates a new groupPolicySettingMapping and sets the default values.
func NewGroupPolicySettingMapping()(*GroupPolicySettingMapping) {
    m := &GroupPolicySettingMapping{
        Entity: *NewEntity(),
    }
    return m
}
// CreateGroupPolicySettingMappingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicySettingMappingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicySettingMapping(), nil
}
// GetAdmxSettingDefinitionId gets the admxSettingDefinitionId property value. Admx Group Policy Id
func (m *GroupPolicySettingMapping) GetAdmxSettingDefinitionId()(*string) {
    return m.admxSettingDefinitionId
}
// GetChildIdList gets the childIdList property value. List of Child Ids of the group policy setting.
func (m *GroupPolicySettingMapping) GetChildIdList()([]string) {
    return m.childIdList
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicySettingMapping) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["admxSettingDefinitionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdmxSettingDefinitionId(val)
        }
        return nil
    }
    res["childIdList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetChildIdList(res)
        }
        return nil
    }
    res["intuneSettingDefinitionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntuneSettingDefinitionId(val)
        }
        return nil
    }
    res["intuneSettingUriList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIntuneSettingUriList(res)
        }
        return nil
    }
    res["isMdmSupported"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMdmSupported(val)
        }
        return nil
    }
    res["mdmCspName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMdmCspName(val)
        }
        return nil
    }
    res["mdmMinimumOSVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMdmMinimumOSVersion(val)
        }
        return nil
    }
    res["mdmSettingUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMdmSettingUri(val)
        }
        return nil
    }
    res["mdmSupportedState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMdmSupportedState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMdmSupportedState(val.(*MdmSupportedState))
        }
        return nil
    }
    res["parentId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentId(val)
        }
        return nil
    }
    res["settingCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingCategory(val)
        }
        return nil
    }
    res["settingDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingDisplayName(val)
        }
        return nil
    }
    res["settingDisplayValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingDisplayValue(val)
        }
        return nil
    }
    res["settingDisplayValueType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingDisplayValueType(val)
        }
        return nil
    }
    res["settingName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingName(val)
        }
        return nil
    }
    res["settingScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicySettingScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingScope(val.(*GroupPolicySettingScope))
        }
        return nil
    }
    res["settingType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicySettingType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingType(val.(*GroupPolicySettingType))
        }
        return nil
    }
    res["settingValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingValue(val)
        }
        return nil
    }
    res["settingValueDisplayUnits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingValueDisplayUnits(val)
        }
        return nil
    }
    res["settingValueType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingValueType(val)
        }
        return nil
    }
    return res
}
// GetIntuneSettingDefinitionId gets the intuneSettingDefinitionId property value. The Intune Setting Definition Id
func (m *GroupPolicySettingMapping) GetIntuneSettingDefinitionId()(*string) {
    return m.intuneSettingDefinitionId
}
// GetIntuneSettingUriList gets the intuneSettingUriList property value. The list of Intune Setting URIs this group policy setting maps to
func (m *GroupPolicySettingMapping) GetIntuneSettingUriList()([]string) {
    return m.intuneSettingUriList
}
// GetIsMdmSupported gets the isMdmSupported property value. Indicates if the setting is supported by Intune or not
func (m *GroupPolicySettingMapping) GetIsMdmSupported()(*bool) {
    return m.isMdmSupported
}
// GetMdmCspName gets the mdmCspName property value. The CSP name this group policy setting maps to.
func (m *GroupPolicySettingMapping) GetMdmCspName()(*string) {
    return m.mdmCspName
}
// GetMdmMinimumOSVersion gets the mdmMinimumOSVersion property value. The minimum OS version this mdm setting supports.
func (m *GroupPolicySettingMapping) GetMdmMinimumOSVersion()(*int32) {
    return m.mdmMinimumOSVersion
}
// GetMdmSettingUri gets the mdmSettingUri property value. The MDM CSP URI this group policy setting maps to.
func (m *GroupPolicySettingMapping) GetMdmSettingUri()(*string) {
    return m.mdmSettingUri
}
// GetMdmSupportedState gets the mdmSupportedState property value. Mdm Support Status of the setting.
func (m *GroupPolicySettingMapping) GetMdmSupportedState()(*MdmSupportedState) {
    return m.mdmSupportedState
}
// GetParentId gets the parentId property value. Parent Id of the group policy setting.
func (m *GroupPolicySettingMapping) GetParentId()(*string) {
    return m.parentId
}
// GetSettingCategory gets the settingCategory property value. The category the group policy setting is in.
func (m *GroupPolicySettingMapping) GetSettingCategory()(*string) {
    return m.settingCategory
}
// GetSettingDisplayName gets the settingDisplayName property value. The display name of this group policy setting.
func (m *GroupPolicySettingMapping) GetSettingDisplayName()(*string) {
    return m.settingDisplayName
}
// GetSettingDisplayValue gets the settingDisplayValue property value. The display value of this group policy setting.
func (m *GroupPolicySettingMapping) GetSettingDisplayValue()(*string) {
    return m.settingDisplayValue
}
// GetSettingDisplayValueType gets the settingDisplayValueType property value. The display value type of this group policy setting.
func (m *GroupPolicySettingMapping) GetSettingDisplayValueType()(*string) {
    return m.settingDisplayValueType
}
// GetSettingName gets the settingName property value. The name of this group policy setting.
func (m *GroupPolicySettingMapping) GetSettingName()(*string) {
    return m.settingName
}
// GetSettingScope gets the settingScope property value. Scope of the group policy setting.
func (m *GroupPolicySettingMapping) GetSettingScope()(*GroupPolicySettingScope) {
    return m.settingScope
}
// GetSettingType gets the settingType property value. Setting type of the group policy.
func (m *GroupPolicySettingMapping) GetSettingType()(*GroupPolicySettingType) {
    return m.settingType
}
// GetSettingValue gets the settingValue property value. The value of this group policy setting.
func (m *GroupPolicySettingMapping) GetSettingValue()(*string) {
    return m.settingValue
}
// GetSettingValueDisplayUnits gets the settingValueDisplayUnits property value. The display units of this group policy setting value
func (m *GroupPolicySettingMapping) GetSettingValueDisplayUnits()(*string) {
    return m.settingValueDisplayUnits
}
// GetSettingValueType gets the settingValueType property value. The value type of this group policy setting.
func (m *GroupPolicySettingMapping) GetSettingValueType()(*string) {
    return m.settingValueType
}
// Serialize serializes information the current object
func (m *GroupPolicySettingMapping) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("admxSettingDefinitionId", m.GetAdmxSettingDefinitionId())
        if err != nil {
            return err
        }
    }
    if m.GetChildIdList() != nil {
        err = writer.WriteCollectionOfStringValues("childIdList", m.GetChildIdList())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("intuneSettingDefinitionId", m.GetIntuneSettingDefinitionId())
        if err != nil {
            return err
        }
    }
    if m.GetIntuneSettingUriList() != nil {
        err = writer.WriteCollectionOfStringValues("intuneSettingUriList", m.GetIntuneSettingUriList())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMdmSupported", m.GetIsMdmSupported())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mdmCspName", m.GetMdmCspName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("mdmMinimumOSVersion", m.GetMdmMinimumOSVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mdmSettingUri", m.GetMdmSettingUri())
        if err != nil {
            return err
        }
    }
    if m.GetMdmSupportedState() != nil {
        cast := (*m.GetMdmSupportedState()).String()
        err = writer.WriteStringValue("mdmSupportedState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parentId", m.GetParentId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingCategory", m.GetSettingCategory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingDisplayName", m.GetSettingDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingDisplayValue", m.GetSettingDisplayValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingDisplayValueType", m.GetSettingDisplayValueType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingName", m.GetSettingName())
        if err != nil {
            return err
        }
    }
    if m.GetSettingScope() != nil {
        cast := (*m.GetSettingScope()).String()
        err = writer.WriteStringValue("settingScope", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSettingType() != nil {
        cast := (*m.GetSettingType()).String()
        err = writer.WriteStringValue("settingType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingValue", m.GetSettingValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingValueDisplayUnits", m.GetSettingValueDisplayUnits())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingValueType", m.GetSettingValueType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdmxSettingDefinitionId sets the admxSettingDefinitionId property value. Admx Group Policy Id
func (m *GroupPolicySettingMapping) SetAdmxSettingDefinitionId(value *string)() {
    m.admxSettingDefinitionId = value
}
// SetChildIdList sets the childIdList property value. List of Child Ids of the group policy setting.
func (m *GroupPolicySettingMapping) SetChildIdList(value []string)() {
    m.childIdList = value
}
// SetIntuneSettingDefinitionId sets the intuneSettingDefinitionId property value. The Intune Setting Definition Id
func (m *GroupPolicySettingMapping) SetIntuneSettingDefinitionId(value *string)() {
    m.intuneSettingDefinitionId = value
}
// SetIntuneSettingUriList sets the intuneSettingUriList property value. The list of Intune Setting URIs this group policy setting maps to
func (m *GroupPolicySettingMapping) SetIntuneSettingUriList(value []string)() {
    m.intuneSettingUriList = value
}
// SetIsMdmSupported sets the isMdmSupported property value. Indicates if the setting is supported by Intune or not
func (m *GroupPolicySettingMapping) SetIsMdmSupported(value *bool)() {
    m.isMdmSupported = value
}
// SetMdmCspName sets the mdmCspName property value. The CSP name this group policy setting maps to.
func (m *GroupPolicySettingMapping) SetMdmCspName(value *string)() {
    m.mdmCspName = value
}
// SetMdmMinimumOSVersion sets the mdmMinimumOSVersion property value. The minimum OS version this mdm setting supports.
func (m *GroupPolicySettingMapping) SetMdmMinimumOSVersion(value *int32)() {
    m.mdmMinimumOSVersion = value
}
// SetMdmSettingUri sets the mdmSettingUri property value. The MDM CSP URI this group policy setting maps to.
func (m *GroupPolicySettingMapping) SetMdmSettingUri(value *string)() {
    m.mdmSettingUri = value
}
// SetMdmSupportedState sets the mdmSupportedState property value. Mdm Support Status of the setting.
func (m *GroupPolicySettingMapping) SetMdmSupportedState(value *MdmSupportedState)() {
    m.mdmSupportedState = value
}
// SetParentId sets the parentId property value. Parent Id of the group policy setting.
func (m *GroupPolicySettingMapping) SetParentId(value *string)() {
    m.parentId = value
}
// SetSettingCategory sets the settingCategory property value. The category the group policy setting is in.
func (m *GroupPolicySettingMapping) SetSettingCategory(value *string)() {
    m.settingCategory = value
}
// SetSettingDisplayName sets the settingDisplayName property value. The display name of this group policy setting.
func (m *GroupPolicySettingMapping) SetSettingDisplayName(value *string)() {
    m.settingDisplayName = value
}
// SetSettingDisplayValue sets the settingDisplayValue property value. The display value of this group policy setting.
func (m *GroupPolicySettingMapping) SetSettingDisplayValue(value *string)() {
    m.settingDisplayValue = value
}
// SetSettingDisplayValueType sets the settingDisplayValueType property value. The display value type of this group policy setting.
func (m *GroupPolicySettingMapping) SetSettingDisplayValueType(value *string)() {
    m.settingDisplayValueType = value
}
// SetSettingName sets the settingName property value. The name of this group policy setting.
func (m *GroupPolicySettingMapping) SetSettingName(value *string)() {
    m.settingName = value
}
// SetSettingScope sets the settingScope property value. Scope of the group policy setting.
func (m *GroupPolicySettingMapping) SetSettingScope(value *GroupPolicySettingScope)() {
    m.settingScope = value
}
// SetSettingType sets the settingType property value. Setting type of the group policy.
func (m *GroupPolicySettingMapping) SetSettingType(value *GroupPolicySettingType)() {
    m.settingType = value
}
// SetSettingValue sets the settingValue property value. The value of this group policy setting.
func (m *GroupPolicySettingMapping) SetSettingValue(value *string)() {
    m.settingValue = value
}
// SetSettingValueDisplayUnits sets the settingValueDisplayUnits property value. The display units of this group policy setting value
func (m *GroupPolicySettingMapping) SetSettingValueDisplayUnits(value *string)() {
    m.settingValueDisplayUnits = value
}
// SetSettingValueType sets the settingValueType property value. The value type of this group policy setting.
func (m *GroupPolicySettingMapping) SetSettingValueType(value *string)() {
    m.settingValueType = value
}
