package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConditionalAccessDeviceStates 
type ConditionalAccessDeviceStates struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // States excluded from the scope of the policy. Possible values: Compliant, DomainJoined.
    excludeStates []string
    // States in the scope of the policy. All is the only allowed value.
    includeStates []string
    // The OdataType property
    odataType *string
}
// NewConditionalAccessDeviceStates instantiates a new conditionalAccessDeviceStates and sets the default values.
func NewConditionalAccessDeviceStates()(*ConditionalAccessDeviceStates) {
    m := &ConditionalAccessDeviceStates{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateConditionalAccessDeviceStatesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConditionalAccessDeviceStatesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessDeviceStates(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessDeviceStates) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetExcludeStates gets the excludeStates property value. States excluded from the scope of the policy. Possible values: Compliant, DomainJoined.
func (m *ConditionalAccessDeviceStates) GetExcludeStates()([]string) {
    return m.excludeStates
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessDeviceStates) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["excludeStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExcludeStates(res)
        }
        return nil
    }
    res["includeStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIncludeStates(res)
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
    return res
}
// GetIncludeStates gets the includeStates property value. States in the scope of the policy. All is the only allowed value.
func (m *ConditionalAccessDeviceStates) GetIncludeStates()([]string) {
    return m.includeStates
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ConditionalAccessDeviceStates) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ConditionalAccessDeviceStates) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetExcludeStates() != nil {
        err := writer.WriteCollectionOfStringValues("excludeStates", m.GetExcludeStates())
        if err != nil {
            return err
        }
    }
    if m.GetIncludeStates() != nil {
        err := writer.WriteCollectionOfStringValues("includeStates", m.GetIncludeStates())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessDeviceStates) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetExcludeStates sets the excludeStates property value. States excluded from the scope of the policy. Possible values: Compliant, DomainJoined.
func (m *ConditionalAccessDeviceStates) SetExcludeStates(value []string)() {
    m.excludeStates = value
}
// SetIncludeStates sets the includeStates property value. States in the scope of the policy. All is the only allowed value.
func (m *ConditionalAccessDeviceStates) SetIncludeStates(value []string)() {
    m.includeStates = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ConditionalAccessDeviceStates) SetOdataType(value *string)() {
    m.odataType = value
}
