package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PayloadByFilter this entity represents a single payload with requested assignment filter Id
type PayloadByFilter struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Represents type of the assignment filter.
    assignmentFilterType *DeviceAndAppManagementAssignmentFilterType
    // The Azure AD security group ID
    groupId *string
    // The OdataType property
    odataType *string
    // The policy identifier
    payloadId *string
    // This enum represents associated assignment payload type
    payloadType *AssociatedAssignmentPayloadType
}
// NewPayloadByFilter instantiates a new payloadByFilter and sets the default values.
func NewPayloadByFilter()(*PayloadByFilter) {
    m := &PayloadByFilter{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePayloadByFilterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePayloadByFilterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPayloadByFilter(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PayloadByFilter) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAssignmentFilterType gets the assignmentFilterType property value. Represents type of the assignment filter.
func (m *PayloadByFilter) GetAssignmentFilterType()(*DeviceAndAppManagementAssignmentFilterType) {
    return m.assignmentFilterType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PayloadByFilter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignmentFilterType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAndAppManagementAssignmentFilterType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentFilterType(val.(*DeviceAndAppManagementAssignmentFilterType))
        }
        return nil
    }
    res["groupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupId(val)
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
    res["payloadId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadId(val)
        }
        return nil
    }
    res["payloadType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAssociatedAssignmentPayloadType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadType(val.(*AssociatedAssignmentPayloadType))
        }
        return nil
    }
    return res
}
// GetGroupId gets the groupId property value. The Azure AD security group ID
func (m *PayloadByFilter) GetGroupId()(*string) {
    return m.groupId
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PayloadByFilter) GetOdataType()(*string) {
    return m.odataType
}
// GetPayloadId gets the payloadId property value. The policy identifier
func (m *PayloadByFilter) GetPayloadId()(*string) {
    return m.payloadId
}
// GetPayloadType gets the payloadType property value. This enum represents associated assignment payload type
func (m *PayloadByFilter) GetPayloadType()(*AssociatedAssignmentPayloadType) {
    return m.payloadType
}
// Serialize serializes information the current object
func (m *PayloadByFilter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssignmentFilterType() != nil {
        cast := (*m.GetAssignmentFilterType()).String()
        err := writer.WriteStringValue("assignmentFilterType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("groupId", m.GetGroupId())
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
        err := writer.WriteStringValue("payloadId", m.GetPayloadId())
        if err != nil {
            return err
        }
    }
    if m.GetPayloadType() != nil {
        cast := (*m.GetPayloadType()).String()
        err := writer.WriteStringValue("payloadType", &cast)
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
func (m *PayloadByFilter) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAssignmentFilterType sets the assignmentFilterType property value. Represents type of the assignment filter.
func (m *PayloadByFilter) SetAssignmentFilterType(value *DeviceAndAppManagementAssignmentFilterType)() {
    m.assignmentFilterType = value
}
// SetGroupId sets the groupId property value. The Azure AD security group ID
func (m *PayloadByFilter) SetGroupId(value *string)() {
    m.groupId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PayloadByFilter) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPayloadId sets the payloadId property value. The policy identifier
func (m *PayloadByFilter) SetPayloadId(value *string)() {
    m.payloadId = value
}
// SetPayloadType sets the payloadType property value. This enum represents associated assignment payload type
func (m *PayloadByFilter) SetPayloadType(value *AssociatedAssignmentPayloadType)() {
    m.payloadType = value
}
