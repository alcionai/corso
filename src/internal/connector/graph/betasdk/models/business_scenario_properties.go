package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BusinessScenarioProperties 
type BusinessScenarioProperties struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The identifier for the bucketDefinition configured in the plannerPlanConfiguration for the scenario. The task will be placed in the corresponding plannerBucket in the target plan. Required.
    externalBucketId *string
    // The identifier for the context of the task. Context is an application controlled value, and tasks can be queried by their externalContextId. Optional.
    externalContextId *string
    // Application-specific identifier for the task. Every task for the same scenario must have a unique identifier specified for this property. Required.
    externalObjectId *string
    // Application-specific version of the task. Optional.
    externalObjectVersion *string
    // The OdataType property
    odataType *string
    // The URL to the application-specific experience for this task. Optional.
    webUrl *string
}
// NewBusinessScenarioProperties instantiates a new businessScenarioProperties and sets the default values.
func NewBusinessScenarioProperties()(*BusinessScenarioProperties) {
    m := &BusinessScenarioProperties{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateBusinessScenarioPropertiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBusinessScenarioPropertiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBusinessScenarioProperties(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BusinessScenarioProperties) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetExternalBucketId gets the externalBucketId property value. The identifier for the bucketDefinition configured in the plannerPlanConfiguration for the scenario. The task will be placed in the corresponding plannerBucket in the target plan. Required.
func (m *BusinessScenarioProperties) GetExternalBucketId()(*string) {
    return m.externalBucketId
}
// GetExternalContextId gets the externalContextId property value. The identifier for the context of the task. Context is an application controlled value, and tasks can be queried by their externalContextId. Optional.
func (m *BusinessScenarioProperties) GetExternalContextId()(*string) {
    return m.externalContextId
}
// GetExternalObjectId gets the externalObjectId property value. Application-specific identifier for the task. Every task for the same scenario must have a unique identifier specified for this property. Required.
func (m *BusinessScenarioProperties) GetExternalObjectId()(*string) {
    return m.externalObjectId
}
// GetExternalObjectVersion gets the externalObjectVersion property value. Application-specific version of the task. Optional.
func (m *BusinessScenarioProperties) GetExternalObjectVersion()(*string) {
    return m.externalObjectVersion
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BusinessScenarioProperties) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["externalBucketId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalBucketId(val)
        }
        return nil
    }
    res["externalContextId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalContextId(val)
        }
        return nil
    }
    res["externalObjectId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalObjectId(val)
        }
        return nil
    }
    res["externalObjectVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalObjectVersion(val)
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
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *BusinessScenarioProperties) GetOdataType()(*string) {
    return m.odataType
}
// GetWebUrl gets the webUrl property value. The URL to the application-specific experience for this task. Optional.
func (m *BusinessScenarioProperties) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *BusinessScenarioProperties) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("externalBucketId", m.GetExternalBucketId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalContextId", m.GetExternalContextId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalObjectId", m.GetExternalObjectId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalObjectVersion", m.GetExternalObjectVersion())
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
        err := writer.WriteStringValue("webUrl", m.GetWebUrl())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BusinessScenarioProperties) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetExternalBucketId sets the externalBucketId property value. The identifier for the bucketDefinition configured in the plannerPlanConfiguration for the scenario. The task will be placed in the corresponding plannerBucket in the target plan. Required.
func (m *BusinessScenarioProperties) SetExternalBucketId(value *string)() {
    m.externalBucketId = value
}
// SetExternalContextId sets the externalContextId property value. The identifier for the context of the task. Context is an application controlled value, and tasks can be queried by their externalContextId. Optional.
func (m *BusinessScenarioProperties) SetExternalContextId(value *string)() {
    m.externalContextId = value
}
// SetExternalObjectId sets the externalObjectId property value. Application-specific identifier for the task. Every task for the same scenario must have a unique identifier specified for this property. Required.
func (m *BusinessScenarioProperties) SetExternalObjectId(value *string)() {
    m.externalObjectId = value
}
// SetExternalObjectVersion sets the externalObjectVersion property value. Application-specific version of the task. Optional.
func (m *BusinessScenarioProperties) SetExternalObjectVersion(value *string)() {
    m.externalObjectVersion = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BusinessScenarioProperties) SetOdataType(value *string)() {
    m.odataType = value
}
// SetWebUrl sets the webUrl property value. The URL to the application-specific experience for this task. Optional.
func (m *BusinessScenarioProperties) SetWebUrl(value *string)() {
    m.webUrl = value
}
