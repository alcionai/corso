package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationLinkedObjects 
type SynchronizationLinkedObjects struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The manager property
    manager SynchronizationJobSubjectable
    // All group members that you would like to provision.
    members []SynchronizationJobSubjectable
    // The OdataType property
    odataType *string
    // The owners property
    owners []SynchronizationJobSubjectable
}
// NewSynchronizationLinkedObjects instantiates a new synchronizationLinkedObjects and sets the default values.
func NewSynchronizationLinkedObjects()(*SynchronizationLinkedObjects) {
    m := &SynchronizationLinkedObjects{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSynchronizationLinkedObjectsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationLinkedObjectsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationLinkedObjects(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SynchronizationLinkedObjects) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationLinkedObjects) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["manager"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSynchronizationJobSubjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManager(val.(SynchronizationJobSubjectable))
        }
        return nil
    }
    res["members"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSynchronizationJobSubjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SynchronizationJobSubjectable, len(val))
            for i, v := range val {
                res[i] = v.(SynchronizationJobSubjectable)
            }
            m.SetMembers(res)
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
    res["owners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSynchronizationJobSubjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SynchronizationJobSubjectable, len(val))
            for i, v := range val {
                res[i] = v.(SynchronizationJobSubjectable)
            }
            m.SetOwners(res)
        }
        return nil
    }
    return res
}
// GetManager gets the manager property value. The manager property
func (m *SynchronizationLinkedObjects) GetManager()(SynchronizationJobSubjectable) {
    return m.manager
}
// GetMembers gets the members property value. All group members that you would like to provision.
func (m *SynchronizationLinkedObjects) GetMembers()([]SynchronizationJobSubjectable) {
    return m.members
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SynchronizationLinkedObjects) GetOdataType()(*string) {
    return m.odataType
}
// GetOwners gets the owners property value. The owners property
func (m *SynchronizationLinkedObjects) GetOwners()([]SynchronizationJobSubjectable) {
    return m.owners
}
// Serialize serializes information the current object
func (m *SynchronizationLinkedObjects) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("manager", m.GetManager())
        if err != nil {
            return err
        }
    }
    if m.GetMembers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("members", cast)
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
    if m.GetOwners() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOwners()))
        for i, v := range m.GetOwners() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("owners", cast)
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
func (m *SynchronizationLinkedObjects) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetManager sets the manager property value. The manager property
func (m *SynchronizationLinkedObjects) SetManager(value SynchronizationJobSubjectable)() {
    m.manager = value
}
// SetMembers sets the members property value. All group members that you would like to provision.
func (m *SynchronizationLinkedObjects) SetMembers(value []SynchronizationJobSubjectable)() {
    m.members = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SynchronizationLinkedObjects) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOwners sets the owners property value. The owners property
func (m *SynchronizationLinkedObjects) SetOwners(value []SynchronizationJobSubjectable)() {
    m.owners = value
}
