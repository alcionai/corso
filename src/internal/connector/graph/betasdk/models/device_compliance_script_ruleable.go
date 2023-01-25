package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceComplianceScriptRuleable 
type DeviceComplianceScriptRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDataType()(*DataType)
    GetDeviceComplianceScriptRuleDataType()(*DeviceComplianceScriptRuleDataType)
    GetDeviceComplianceScriptRulOperator()(*DeviceComplianceScriptRulOperator)
    GetOdataType()(*string)
    GetOperand()(*string)
    GetOperator()(*Operator)
    GetSettingName()(*string)
    SetDataType(value *DataType)()
    SetDeviceComplianceScriptRuleDataType(value *DeviceComplianceScriptRuleDataType)()
    SetDeviceComplianceScriptRulOperator(value *DeviceComplianceScriptRulOperator)()
    SetOdataType(value *string)()
    SetOperand(value *string)()
    SetOperator(value *Operator)()
    SetSettingName(value *string)()
}
