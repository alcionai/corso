package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewInstanceDecisionItemServicePrincipalTarget 
type AccessReviewInstanceDecisionItemServicePrincipalTarget struct {
    AccessReviewInstanceDecisionItemTarget
    // The appId for the service principal entity being reviewed.
    appId *string
    // The display name of the service principal whose access is being reviewed.
    servicePrincipalDisplayName *string
    // The servicePrincipalId property
    servicePrincipalId *string
}
// NewAccessReviewInstanceDecisionItemServicePrincipalTarget instantiates a new AccessReviewInstanceDecisionItemServicePrincipalTarget and sets the default values.
func NewAccessReviewInstanceDecisionItemServicePrincipalTarget()(*AccessReviewInstanceDecisionItemServicePrincipalTarget) {
    m := &AccessReviewInstanceDecisionItemServicePrincipalTarget{
        AccessReviewInstanceDecisionItemTarget: *NewAccessReviewInstanceDecisionItemTarget(),
    }
    odataTypeValue := "#microsoft.graph.accessReviewInstanceDecisionItemServicePrincipalTarget";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAccessReviewInstanceDecisionItemServicePrincipalTargetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewInstanceDecisionItemServicePrincipalTargetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewInstanceDecisionItemServicePrincipalTarget(), nil
}
// GetAppId gets the appId property value. The appId for the service principal entity being reviewed.
func (m *AccessReviewInstanceDecisionItemServicePrincipalTarget) GetAppId()(*string) {
    return m.appId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewInstanceDecisionItemServicePrincipalTarget) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AccessReviewInstanceDecisionItemTarget.GetFieldDeserializers()
    res["appId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["servicePrincipalDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalDisplayName(val)
        }
        return nil
    }
    res["servicePrincipalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalId(val)
        }
        return nil
    }
    return res
}
// GetServicePrincipalDisplayName gets the servicePrincipalDisplayName property value. The display name of the service principal whose access is being reviewed.
func (m *AccessReviewInstanceDecisionItemServicePrincipalTarget) GetServicePrincipalDisplayName()(*string) {
    return m.servicePrincipalDisplayName
}
// GetServicePrincipalId gets the servicePrincipalId property value. The servicePrincipalId property
func (m *AccessReviewInstanceDecisionItemServicePrincipalTarget) GetServicePrincipalId()(*string) {
    return m.servicePrincipalId
}
// Serialize serializes information the current object
func (m *AccessReviewInstanceDecisionItemServicePrincipalTarget) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AccessReviewInstanceDecisionItemTarget.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePrincipalDisplayName", m.GetServicePrincipalDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePrincipalId", m.GetServicePrincipalId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppId sets the appId property value. The appId for the service principal entity being reviewed.
func (m *AccessReviewInstanceDecisionItemServicePrincipalTarget) SetAppId(value *string)() {
    m.appId = value
}
// SetServicePrincipalDisplayName sets the servicePrincipalDisplayName property value. The display name of the service principal whose access is being reviewed.
func (m *AccessReviewInstanceDecisionItemServicePrincipalTarget) SetServicePrincipalDisplayName(value *string)() {
    m.servicePrincipalDisplayName = value
}
// SetServicePrincipalId sets the servicePrincipalId property value. The servicePrincipalId property
func (m *AccessReviewInstanceDecisionItemServicePrincipalTarget) SetServicePrincipalId(value *string)() {
    m.servicePrincipalId = value
}
