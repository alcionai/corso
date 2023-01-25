package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewInstanceDecisionItemUserTarget 
type AccessReviewInstanceDecisionItemUserTarget struct {
    AccessReviewInstanceDecisionItemTarget
    // The name of user.
    userDisplayName *string
    // The identifier of user.
    userId *string
    // The user principal name.
    userPrincipalName *string
}
// NewAccessReviewInstanceDecisionItemUserTarget instantiates a new AccessReviewInstanceDecisionItemUserTarget and sets the default values.
func NewAccessReviewInstanceDecisionItemUserTarget()(*AccessReviewInstanceDecisionItemUserTarget) {
    m := &AccessReviewInstanceDecisionItemUserTarget{
        AccessReviewInstanceDecisionItemTarget: *NewAccessReviewInstanceDecisionItemTarget(),
    }
    odataTypeValue := "#microsoft.graph.accessReviewInstanceDecisionItemUserTarget";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAccessReviewInstanceDecisionItemUserTargetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewInstanceDecisionItemUserTargetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewInstanceDecisionItemUserTarget(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewInstanceDecisionItemUserTarget) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AccessReviewInstanceDecisionItemTarget.GetFieldDeserializers()
    res["userDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDisplayName(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetUserDisplayName gets the userDisplayName property value. The name of user.
func (m *AccessReviewInstanceDecisionItemUserTarget) GetUserDisplayName()(*string) {
    return m.userDisplayName
}
// GetUserId gets the userId property value. The identifier of user.
func (m *AccessReviewInstanceDecisionItemUserTarget) GetUserId()(*string) {
    return m.userId
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name.
func (m *AccessReviewInstanceDecisionItemUserTarget) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *AccessReviewInstanceDecisionItemUserTarget) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AccessReviewInstanceDecisionItemTarget.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("userDisplayName", m.GetUserDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUserDisplayName sets the userDisplayName property value. The name of user.
func (m *AccessReviewInstanceDecisionItemUserTarget) SetUserDisplayName(value *string)() {
    m.userDisplayName = value
}
// SetUserId sets the userId property value. The identifier of user.
func (m *AccessReviewInstanceDecisionItemUserTarget) SetUserId(value *string)() {
    m.userId = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name.
func (m *AccessReviewInstanceDecisionItemUserTarget) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
