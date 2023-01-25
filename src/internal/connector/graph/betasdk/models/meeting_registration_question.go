package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeetingRegistrationQuestion provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type MeetingRegistrationQuestion struct {
    Entity
    // Answer input type of the custom registration question.
    answerInputType *AnswerInputType
    // Answer options when answerInputType is radioButton.
    answerOptions []string
    // Display name of the custom registration question.
    displayName *string
    // Indicates whether the question is required. Default value is false.
    isRequired *bool
}
// NewMeetingRegistrationQuestion instantiates a new meetingRegistrationQuestion and sets the default values.
func NewMeetingRegistrationQuestion()(*MeetingRegistrationQuestion) {
    m := &MeetingRegistrationQuestion{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMeetingRegistrationQuestionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeetingRegistrationQuestionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeetingRegistrationQuestion(), nil
}
// GetAnswerInputType gets the answerInputType property value. Answer input type of the custom registration question.
func (m *MeetingRegistrationQuestion) GetAnswerInputType()(*AnswerInputType) {
    return m.answerInputType
}
// GetAnswerOptions gets the answerOptions property value. Answer options when answerInputType is radioButton.
func (m *MeetingRegistrationQuestion) GetAnswerOptions()([]string) {
    return m.answerOptions
}
// GetDisplayName gets the displayName property value. Display name of the custom registration question.
func (m *MeetingRegistrationQuestion) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingRegistrationQuestion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["answerInputType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAnswerInputType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnswerInputType(val.(*AnswerInputType))
        }
        return nil
    }
    res["answerOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAnswerOptions(res)
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
    res["isRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRequired(val)
        }
        return nil
    }
    return res
}
// GetIsRequired gets the isRequired property value. Indicates whether the question is required. Default value is false.
func (m *MeetingRegistrationQuestion) GetIsRequired()(*bool) {
    return m.isRequired
}
// Serialize serializes information the current object
func (m *MeetingRegistrationQuestion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAnswerInputType() != nil {
        cast := (*m.GetAnswerInputType()).String()
        err = writer.WriteStringValue("answerInputType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAnswerOptions() != nil {
        err = writer.WriteCollectionOfStringValues("answerOptions", m.GetAnswerOptions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isRequired", m.GetIsRequired())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAnswerInputType sets the answerInputType property value. Answer input type of the custom registration question.
func (m *MeetingRegistrationQuestion) SetAnswerInputType(value *AnswerInputType)() {
    m.answerInputType = value
}
// SetAnswerOptions sets the answerOptions property value. Answer options when answerInputType is radioButton.
func (m *MeetingRegistrationQuestion) SetAnswerOptions(value []string)() {
    m.answerOptions = value
}
// SetDisplayName sets the displayName property value. Display name of the custom registration question.
func (m *MeetingRegistrationQuestion) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsRequired sets the isRequired property value. Indicates whether the question is required. Default value is false.
func (m *MeetingRegistrationQuestion) SetIsRequired(value *bool)() {
    m.isRequired = value
}
