package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeetingInfo 
type MeetingInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The allowConversationWithoutHost property
    allowConversationWithoutHost *bool
    // The OdataType property
    odataType *string
}
// NewMeetingInfo instantiates a new meetingInfo and sets the default values.
func NewMeetingInfo()(*MeetingInfo) {
    m := &MeetingInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeetingInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeetingInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.joinMeetingIdMeetingInfo":
                        return NewJoinMeetingIdMeetingInfo(), nil
                    case "#microsoft.graph.organizerMeetingInfo":
                        return NewOrganizerMeetingInfo(), nil
                    case "#microsoft.graph.tokenMeetingInfo":
                        return NewTokenMeetingInfo(), nil
                }
            }
        }
    }
    return NewMeetingInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeetingInfo) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAllowConversationWithoutHost gets the allowConversationWithoutHost property value. The allowConversationWithoutHost property
func (m *MeetingInfo) GetAllowConversationWithoutHost()(*bool) {
    return m.allowConversationWithoutHost
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowConversationWithoutHost"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowConversationWithoutHost(val)
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MeetingInfo) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *MeetingInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowConversationWithoutHost", m.GetAllowConversationWithoutHost())
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
func (m *MeetingInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAllowConversationWithoutHost sets the allowConversationWithoutHost property value. The allowConversationWithoutHost property
func (m *MeetingInfo) SetAllowConversationWithoutHost(value *bool)() {
    m.allowConversationWithoutHost = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MeetingInfo) SetOdataType(value *string)() {
    m.odataType = value
}
