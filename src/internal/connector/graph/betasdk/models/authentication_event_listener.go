package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationEventListener 
type AuthenticationEventListener struct {
    Entity
    // The authenticationEventsFlowId property
    authenticationEventsFlowId *string
    // The conditions property
    conditions AuthenticationConditionsable
    // The priority property
    priority *int32
}
// NewAuthenticationEventListener instantiates a new AuthenticationEventListener and sets the default values.
func NewAuthenticationEventListener()(*AuthenticationEventListener) {
    m := &AuthenticationEventListener{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAuthenticationEventListenerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationEventListenerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.onTokenIssuanceStartListener":
                        return NewOnTokenIssuanceStartListener(), nil
                }
            }
        }
    }
    return NewAuthenticationEventListener(), nil
}
// GetAuthenticationEventsFlowId gets the authenticationEventsFlowId property value. The authenticationEventsFlowId property
func (m *AuthenticationEventListener) GetAuthenticationEventsFlowId()(*string) {
    return m.authenticationEventsFlowId
}
// GetConditions gets the conditions property value. The conditions property
func (m *AuthenticationEventListener) GetConditions()(AuthenticationConditionsable) {
    return m.conditions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationEventListener) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authenticationEventsFlowId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationEventsFlowId(val)
        }
        return nil
    }
    res["conditions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationConditionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditions(val.(AuthenticationConditionsable))
        }
        return nil
    }
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    return res
}
// GetPriority gets the priority property value. The priority property
func (m *AuthenticationEventListener) GetPriority()(*int32) {
    return m.priority
}
// Serialize serializes information the current object
func (m *AuthenticationEventListener) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("authenticationEventsFlowId", m.GetAuthenticationEventsFlowId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("conditions", m.GetConditions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthenticationEventsFlowId sets the authenticationEventsFlowId property value. The authenticationEventsFlowId property
func (m *AuthenticationEventListener) SetAuthenticationEventsFlowId(value *string)() {
    m.authenticationEventsFlowId = value
}
// SetConditions sets the conditions property value. The conditions property
func (m *AuthenticationEventListener) SetConditions(value AuthenticationConditionsable)() {
    m.conditions = value
}
// SetPriority sets the priority property value. The priority property
func (m *AuthenticationEventListener) SetPriority(value *int32)() {
    m.priority = value
}
