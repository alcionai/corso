package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MentionAction 
type MentionAction struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The identities of the users mentioned in this action.
    mentionees []IdentitySetable
    // The OdataType property
    odataType *string
}
// NewMentionAction instantiates a new mentionAction and sets the default values.
func NewMentionAction()(*MentionAction) {
    m := &MentionAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMentionActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMentionActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMentionAction(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MentionAction) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MentionAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["mentionees"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentitySetable, len(val))
            for i, v := range val {
                res[i] = v.(IdentitySetable)
            }
            m.SetMentionees(res)
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
// GetMentionees gets the mentionees property value. The identities of the users mentioned in this action.
func (m *MentionAction) GetMentionees()([]IdentitySetable) {
    return m.mentionees
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MentionAction) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *MentionAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetMentionees() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMentionees()))
        for i, v := range m.GetMentionees() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("mentionees", cast)
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
func (m *MentionAction) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetMentionees sets the mentionees property value. The identities of the users mentioned in this action.
func (m *MentionAction) SetMentionees(value []IdentitySetable)() {
    m.mentionees = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MentionAction) SetOdataType(value *string)() {
    m.odataType = value
}
