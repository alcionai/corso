package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Workspace provides operations to manage the collection of site entities.
type Workspace struct {
    Place
    // Specifies the building name or building number that the workspace is in.
    building *string
    // Specifies the capacity of the workspace.
    capacity *int32
    // Email address of the workspace.
    emailAddress *string
    // Specifies a descriptive label for the floor, for example, P.
    floorLabel *string
    // Specifies the floor number that the workspace is on.
    floorNumber *int32
    // Specifies whether the workspace is wheelchair accessible.
    isWheelChairAccessible *bool
    // Specifies a descriptive label for the workspace, for example, a number or name.
    label *string
    // Specifies a nickname for the workspace, for example, 'quiet workspace'.
    nickname *string
    // Specifies additional features of the workspace, for example, details like the type of view or furniture type.
    tags []string
}
// NewWorkspace instantiates a new workspace and sets the default values.
func NewWorkspace()(*Workspace) {
    m := &Workspace{
        Place: *NewPlace(),
    }
    odataTypeValue := "#microsoft.graph.workspace";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWorkspaceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkspaceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkspace(), nil
}
// GetBuilding gets the building property value. Specifies the building name or building number that the workspace is in.
func (m *Workspace) GetBuilding()(*string) {
    return m.building
}
// GetCapacity gets the capacity property value. Specifies the capacity of the workspace.
func (m *Workspace) GetCapacity()(*int32) {
    return m.capacity
}
// GetEmailAddress gets the emailAddress property value. Email address of the workspace.
func (m *Workspace) GetEmailAddress()(*string) {
    return m.emailAddress
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Workspace) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Place.GetFieldDeserializers()
    res["building"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBuilding(val)
        }
        return nil
    }
    res["capacity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCapacity(val)
        }
        return nil
    }
    res["emailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddress(val)
        }
        return nil
    }
    res["floorLabel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFloorLabel(val)
        }
        return nil
    }
    res["floorNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFloorNumber(val)
        }
        return nil
    }
    res["isWheelChairAccessible"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsWheelChairAccessible(val)
        }
        return nil
    }
    res["label"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabel(val)
        }
        return nil
    }
    res["nickname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNickname(val)
        }
        return nil
    }
    res["tags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTags(res)
        }
        return nil
    }
    return res
}
// GetFloorLabel gets the floorLabel property value. Specifies a descriptive label for the floor, for example, P.
func (m *Workspace) GetFloorLabel()(*string) {
    return m.floorLabel
}
// GetFloorNumber gets the floorNumber property value. Specifies the floor number that the workspace is on.
func (m *Workspace) GetFloorNumber()(*int32) {
    return m.floorNumber
}
// GetIsWheelChairAccessible gets the isWheelChairAccessible property value. Specifies whether the workspace is wheelchair accessible.
func (m *Workspace) GetIsWheelChairAccessible()(*bool) {
    return m.isWheelChairAccessible
}
// GetLabel gets the label property value. Specifies a descriptive label for the workspace, for example, a number or name.
func (m *Workspace) GetLabel()(*string) {
    return m.label
}
// GetNickname gets the nickname property value. Specifies a nickname for the workspace, for example, 'quiet workspace'.
func (m *Workspace) GetNickname()(*string) {
    return m.nickname
}
// GetTags gets the tags property value. Specifies additional features of the workspace, for example, details like the type of view or furniture type.
func (m *Workspace) GetTags()([]string) {
    return m.tags
}
// Serialize serializes information the current object
func (m *Workspace) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Place.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("building", m.GetBuilding())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("capacity", m.GetCapacity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("emailAddress", m.GetEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("floorLabel", m.GetFloorLabel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("floorNumber", m.GetFloorNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isWheelChairAccessible", m.GetIsWheelChairAccessible())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("label", m.GetLabel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("nickname", m.GetNickname())
        if err != nil {
            return err
        }
    }
    if m.GetTags() != nil {
        err = writer.WriteCollectionOfStringValues("tags", m.GetTags())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBuilding sets the building property value. Specifies the building name or building number that the workspace is in.
func (m *Workspace) SetBuilding(value *string)() {
    m.building = value
}
// SetCapacity sets the capacity property value. Specifies the capacity of the workspace.
func (m *Workspace) SetCapacity(value *int32)() {
    m.capacity = value
}
// SetEmailAddress sets the emailAddress property value. Email address of the workspace.
func (m *Workspace) SetEmailAddress(value *string)() {
    m.emailAddress = value
}
// SetFloorLabel sets the floorLabel property value. Specifies a descriptive label for the floor, for example, P.
func (m *Workspace) SetFloorLabel(value *string)() {
    m.floorLabel = value
}
// SetFloorNumber sets the floorNumber property value. Specifies the floor number that the workspace is on.
func (m *Workspace) SetFloorNumber(value *int32)() {
    m.floorNumber = value
}
// SetIsWheelChairAccessible sets the isWheelChairAccessible property value. Specifies whether the workspace is wheelchair accessible.
func (m *Workspace) SetIsWheelChairAccessible(value *bool)() {
    m.isWheelChairAccessible = value
}
// SetLabel sets the label property value. Specifies a descriptive label for the workspace, for example, a number or name.
func (m *Workspace) SetLabel(value *string)() {
    m.label = value
}
// SetNickname sets the nickname property value. Specifies a nickname for the workspace, for example, 'quiet workspace'.
func (m *Workspace) SetNickname(value *string)() {
    m.nickname = value
}
// SetTags sets the tags property value. Specifies additional features of the workspace, for example, details like the type of view or furniture type.
func (m *Workspace) SetTags(value []string)() {
    m.tags = value
}
