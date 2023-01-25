package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConditionalAccessDevices 
type ConditionalAccessDevices struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Filter that defines the dynamic-device-syntax rule to include/exclude devices. A filter can use device properties (such as extension attributes) to include/exclude them. Cannot be set if includeDevices or excludeDevices is set.
    deviceFilter ConditionalAccessFilterable
    // States excluded from the scope of the policy. Possible values: Compliant, DomainJoined. Cannot be set if deviceFIlter is set.
    excludeDevices []string
    // The excludeDeviceStates property
    excludeDeviceStates []string
    // States in the scope of the policy. All is the only allowed value. Cannot be set if deviceFIlter is set.
    includeDevices []string
    // The includeDeviceStates property
    includeDeviceStates []string
    // The OdataType property
    odataType *string
}
// NewConditionalAccessDevices instantiates a new conditionalAccessDevices and sets the default values.
func NewConditionalAccessDevices()(*ConditionalAccessDevices) {
    m := &ConditionalAccessDevices{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateConditionalAccessDevicesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConditionalAccessDevicesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessDevices(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessDevices) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceFilter gets the deviceFilter property value. Filter that defines the dynamic-device-syntax rule to include/exclude devices. A filter can use device properties (such as extension attributes) to include/exclude them. Cannot be set if includeDevices or excludeDevices is set.
func (m *ConditionalAccessDevices) GetDeviceFilter()(ConditionalAccessFilterable) {
    return m.deviceFilter
}
// GetExcludeDevices gets the excludeDevices property value. States excluded from the scope of the policy. Possible values: Compliant, DomainJoined. Cannot be set if deviceFIlter is set.
func (m *ConditionalAccessDevices) GetExcludeDevices()([]string) {
    return m.excludeDevices
}
// GetExcludeDeviceStates gets the excludeDeviceStates property value. The excludeDeviceStates property
func (m *ConditionalAccessDevices) GetExcludeDeviceStates()([]string) {
    return m.excludeDeviceStates
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessDevices) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceFilter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConditionalAccessFilterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceFilter(val.(ConditionalAccessFilterable))
        }
        return nil
    }
    res["excludeDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExcludeDevices(res)
        }
        return nil
    }
    res["excludeDeviceStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExcludeDeviceStates(res)
        }
        return nil
    }
    res["includeDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIncludeDevices(res)
        }
        return nil
    }
    res["includeDeviceStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIncludeDeviceStates(res)
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
// GetIncludeDevices gets the includeDevices property value. States in the scope of the policy. All is the only allowed value. Cannot be set if deviceFIlter is set.
func (m *ConditionalAccessDevices) GetIncludeDevices()([]string) {
    return m.includeDevices
}
// GetIncludeDeviceStates gets the includeDeviceStates property value. The includeDeviceStates property
func (m *ConditionalAccessDevices) GetIncludeDeviceStates()([]string) {
    return m.includeDeviceStates
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ConditionalAccessDevices) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ConditionalAccessDevices) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("deviceFilter", m.GetDeviceFilter())
        if err != nil {
            return err
        }
    }
    if m.GetExcludeDevices() != nil {
        err := writer.WriteCollectionOfStringValues("excludeDevices", m.GetExcludeDevices())
        if err != nil {
            return err
        }
    }
    if m.GetExcludeDeviceStates() != nil {
        err := writer.WriteCollectionOfStringValues("excludeDeviceStates", m.GetExcludeDeviceStates())
        if err != nil {
            return err
        }
    }
    if m.GetIncludeDevices() != nil {
        err := writer.WriteCollectionOfStringValues("includeDevices", m.GetIncludeDevices())
        if err != nil {
            return err
        }
    }
    if m.GetIncludeDeviceStates() != nil {
        err := writer.WriteCollectionOfStringValues("includeDeviceStates", m.GetIncludeDeviceStates())
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
func (m *ConditionalAccessDevices) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeviceFilter sets the deviceFilter property value. Filter that defines the dynamic-device-syntax rule to include/exclude devices. A filter can use device properties (such as extension attributes) to include/exclude them. Cannot be set if includeDevices or excludeDevices is set.
func (m *ConditionalAccessDevices) SetDeviceFilter(value ConditionalAccessFilterable)() {
    m.deviceFilter = value
}
// SetExcludeDevices sets the excludeDevices property value. States excluded from the scope of the policy. Possible values: Compliant, DomainJoined. Cannot be set if deviceFIlter is set.
func (m *ConditionalAccessDevices) SetExcludeDevices(value []string)() {
    m.excludeDevices = value
}
// SetExcludeDeviceStates sets the excludeDeviceStates property value. The excludeDeviceStates property
func (m *ConditionalAccessDevices) SetExcludeDeviceStates(value []string)() {
    m.excludeDeviceStates = value
}
// SetIncludeDevices sets the includeDevices property value. States in the scope of the policy. All is the only allowed value. Cannot be set if deviceFIlter is set.
func (m *ConditionalAccessDevices) SetIncludeDevices(value []string)() {
    m.includeDevices = value
}
// SetIncludeDeviceStates sets the includeDeviceStates property value. The includeDeviceStates property
func (m *ConditionalAccessDevices) SetIncludeDeviceStates(value []string)() {
    m.includeDeviceStates = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ConditionalAccessDevices) SetOdataType(value *string)() {
    m.odataType = value
}
