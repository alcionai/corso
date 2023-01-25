package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemActivityOLD provides operations to call the delta method.
type ItemActivityOLD struct {
    Entity
    // The action property
    action ItemActionSetable
    // The actor property
    actor IdentitySetable
    // The driveItem property
    driveItem DriveItemable
    // The listItem property
    listItem ListItemable
    // The times property
    times ItemActivityTimeSetable
}
// NewItemActivityOLD instantiates a new itemActivityOLD and sets the default values.
func NewItemActivityOLD()(*ItemActivityOLD) {
    m := &ItemActivityOLD{
        Entity: *NewEntity(),
    }
    return m
}
// CreateItemActivityOLDFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemActivityOLDFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActivityOLD(), nil
}
// GetAction gets the action property value. The action property
func (m *ItemActivityOLD) GetAction()(ItemActionSetable) {
    return m.action
}
// GetActor gets the actor property value. The actor property
func (m *ItemActivityOLD) GetActor()(IdentitySetable) {
    return m.actor
}
// GetDriveItem gets the driveItem property value. The driveItem property
func (m *ItemActivityOLD) GetDriveItem()(DriveItemable) {
    return m.driveItem
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemActivityOLD) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemActionSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(ItemActionSetable))
        }
        return nil
    }
    res["actor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActor(val.(IdentitySetable))
        }
        return nil
    }
    res["driveItem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDriveItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDriveItem(val.(DriveItemable))
        }
        return nil
    }
    res["listItem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetListItem(val.(ListItemable))
        }
        return nil
    }
    res["times"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemActivityTimeSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimes(val.(ItemActivityTimeSetable))
        }
        return nil
    }
    return res
}
// GetListItem gets the listItem property value. The listItem property
func (m *ItemActivityOLD) GetListItem()(ListItemable) {
    return m.listItem
}
// GetTimes gets the times property value. The times property
func (m *ItemActivityOLD) GetTimes()(ItemActivityTimeSetable) {
    return m.times
}
// Serialize serializes information the current object
func (m *ItemActivityOLD) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("action", m.GetAction())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("actor", m.GetActor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("driveItem", m.GetDriveItem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("listItem", m.GetListItem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("times", m.GetTimes())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAction sets the action property value. The action property
func (m *ItemActivityOLD) SetAction(value ItemActionSetable)() {
    m.action = value
}
// SetActor sets the actor property value. The actor property
func (m *ItemActivityOLD) SetActor(value IdentitySetable)() {
    m.actor = value
}
// SetDriveItem sets the driveItem property value. The driveItem property
func (m *ItemActivityOLD) SetDriveItem(value DriveItemable)() {
    m.driveItem = value
}
// SetListItem sets the listItem property value. The listItem property
func (m *ItemActivityOLD) SetListItem(value ListItemable)() {
    m.listItem = value
}
// SetTimes sets the times property value. The times property
func (m *ItemActivityOLD) SetTimes(value ItemActivityTimeSetable)() {
    m.times = value
}
