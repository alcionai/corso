package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageMultipleChoiceQuestion 
type AccessPackageMultipleChoiceQuestion struct {
    AccessPackageQuestion
    // Indicates whether requestor can select multiple choices as their answer.
    allowsMultipleSelection *bool
    // List of answer choices.
    choices []AccessPackageAnswerChoiceable
}
// NewAccessPackageMultipleChoiceQuestion instantiates a new AccessPackageMultipleChoiceQuestion and sets the default values.
func NewAccessPackageMultipleChoiceQuestion()(*AccessPackageMultipleChoiceQuestion) {
    m := &AccessPackageMultipleChoiceQuestion{
        AccessPackageQuestion: *NewAccessPackageQuestion(),
    }
    odataTypeValue := "#microsoft.graph.accessPackageMultipleChoiceQuestion";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAccessPackageMultipleChoiceQuestionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageMultipleChoiceQuestionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageMultipleChoiceQuestion(), nil
}
// GetAllowsMultipleSelection gets the allowsMultipleSelection property value. Indicates whether requestor can select multiple choices as their answer.
func (m *AccessPackageMultipleChoiceQuestion) GetAllowsMultipleSelection()(*bool) {
    return m.allowsMultipleSelection
}
// GetChoices gets the choices property value. List of answer choices.
func (m *AccessPackageMultipleChoiceQuestion) GetChoices()([]AccessPackageAnswerChoiceable) {
    return m.choices
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageMultipleChoiceQuestion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AccessPackageQuestion.GetFieldDeserializers()
    res["allowsMultipleSelection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowsMultipleSelection(val)
        }
        return nil
    }
    res["choices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageAnswerChoiceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageAnswerChoiceable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageAnswerChoiceable)
            }
            m.SetChoices(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AccessPackageMultipleChoiceQuestion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AccessPackageQuestion.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowsMultipleSelection", m.GetAllowsMultipleSelection())
        if err != nil {
            return err
        }
    }
    if m.GetChoices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChoices()))
        for i, v := range m.GetChoices() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("choices", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowsMultipleSelection sets the allowsMultipleSelection property value. Indicates whether requestor can select multiple choices as their answer.
func (m *AccessPackageMultipleChoiceQuestion) SetAllowsMultipleSelection(value *bool)() {
    m.allowsMultipleSelection = value
}
// SetChoices sets the choices property value. List of answer choices.
func (m *AccessPackageMultipleChoiceQuestion) SetChoices(value []AccessPackageAnswerChoiceable)() {
    m.choices = value
}
