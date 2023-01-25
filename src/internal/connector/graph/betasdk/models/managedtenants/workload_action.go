package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkloadAction 
type WorkloadAction struct {
    // The unique identifier for the workload action. Required. Read-only.
    actionId *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The category for the workload action. Possible values are: automated, manual, unknownFutureValue. Optional. Read-only.
    category *WorkloadActionCategory
    // The description for the workload action. Optional. Read-only.
    description *string
    // The display name for the workload action. Optional. Read-only.
    displayName *string
    // The licenses property
    licenses []string
    // The OdataType property
    odataType *string
    // The service associated with workload action. Optional. Read-only.
    service *string
    // The collection of settings associated with the workload action. Optional. Read-only.
    settings []Settingable
}
// NewWorkloadAction instantiates a new workloadAction and sets the default values.
func NewWorkloadAction()(*WorkloadAction) {
    m := &WorkloadAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWorkloadActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkloadActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkloadAction(), nil
}
// GetActionId gets the actionId property value. The unique identifier for the workload action. Required. Read-only.
func (m *WorkloadAction) GetActionId()(*string) {
    return m.actionId
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkloadAction) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCategory gets the category property value. The category for the workload action. Possible values are: automated, manual, unknownFutureValue. Optional. Read-only.
func (m *WorkloadAction) GetCategory()(*WorkloadActionCategory) {
    return m.category
}
// GetDescription gets the description property value. The description for the workload action. Optional. Read-only.
func (m *WorkloadAction) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name for the workload action. Optional. Read-only.
func (m *WorkloadAction) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkloadAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionId(val)
        }
        return nil
    }
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWorkloadActionCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*WorkloadActionCategory))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["licenses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetLicenses(res)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["service"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetService(val)
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Settingable, len(val))
            for i, v := range val {
                res[i] = v.(Settingable)
            }
            m.SetSettings(res)
        }
        return nil
    }
    return res
}
// GetLicenses gets the licenses property value. The licenses property
func (m *WorkloadAction) GetLicenses()([]string) {
    return m.licenses
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WorkloadAction) GetOdataType()(*string) {
    return m.odataType
}
// GetService gets the service property value. The service associated with workload action. Optional. Read-only.
func (m *WorkloadAction) GetService()(*string) {
    return m.service
}
// GetSettings gets the settings property value. The collection of settings associated with the workload action. Optional. Read-only.
func (m *WorkloadAction) GetSettings()([]Settingable) {
    return m.settings
}
// Serialize serializes information the current object
func (m *WorkloadAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("actionId", m.GetActionId())
        if err != nil {
            return err
        }
    }
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
        err := writer.WriteStringValue("category", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetLicenses() != nil {
        err := writer.WriteCollectionOfStringValues("licenses", m.GetLicenses())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("service", m.GetService())
        if err != nil {
            return err
        }
    }
    if m.GetSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSettings()))
        for i, v := range m.GetSettings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("settings", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionId sets the actionId property value. The unique identifier for the workload action. Required. Read-only.
func (m *WorkloadAction) SetActionId(value *string)() {
    m.actionId = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkloadAction) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCategory sets the category property value. The category for the workload action. Possible values are: automated, manual, unknownFutureValue. Optional. Read-only.
func (m *WorkloadAction) SetCategory(value *WorkloadActionCategory)() {
    m.category = value
}
// SetDescription sets the description property value. The description for the workload action. Optional. Read-only.
func (m *WorkloadAction) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name for the workload action. Optional. Read-only.
func (m *WorkloadAction) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLicenses sets the licenses property value. The licenses property
func (m *WorkloadAction) SetLicenses(value []string)() {
    m.licenses = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WorkloadAction) SetOdataType(value *string)() {
    m.odataType = value
}
// SetService sets the service property value. The service associated with workload action. Optional. Read-only.
func (m *WorkloadAction) SetService(value *string)() {
    m.service = value
}
// SetSettings sets the settings property value. The collection of settings associated with the workload action. Optional. Read-only.
func (m *WorkloadAction) SetSettings(value []Settingable)() {
    m.settings = value
}
