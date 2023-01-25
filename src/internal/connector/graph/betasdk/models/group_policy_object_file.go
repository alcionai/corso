package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyObjectFile the Group Policy Object file uploaded by admin.
type GroupPolicyObjectFile struct {
    Entity
    // The Group Policy Object file content.
    content *string
    // The date and time at which the GroupPolicy was first uploaded.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The Group Policy Object GUID from GPO Xml content
    groupPolicyObjectId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The date and time at which the GroupPolicyObjectFile was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The distinguished name of the OU.
    ouDistinguishedName *string
    // The list of scope tags for the configuration.
    roleScopeTagIds []string
}
// NewGroupPolicyObjectFile instantiates a new groupPolicyObjectFile and sets the default values.
func NewGroupPolicyObjectFile()(*GroupPolicyObjectFile) {
    m := &GroupPolicyObjectFile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateGroupPolicyObjectFileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyObjectFileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyObjectFile(), nil
}
// GetContent gets the content property value. The Group Policy Object file content.
func (m *GroupPolicyObjectFile) GetContent()(*string) {
    return m.content
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time at which the GroupPolicy was first uploaded.
func (m *GroupPolicyObjectFile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyObjectFile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["groupPolicyObjectId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupPolicyObjectId(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["ouDistinguishedName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOuDistinguishedName(val)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    return res
}
// GetGroupPolicyObjectId gets the groupPolicyObjectId property value. The Group Policy Object GUID from GPO Xml content
func (m *GroupPolicyObjectFile) GetGroupPolicyObjectId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.groupPolicyObjectId
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time at which the GroupPolicyObjectFile was last modified.
func (m *GroupPolicyObjectFile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetOuDistinguishedName gets the ouDistinguishedName property value. The distinguished name of the OU.
func (m *GroupPolicyObjectFile) GetOuDistinguishedName()(*string) {
    return m.ouDistinguishedName
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. The list of scope tags for the configuration.
func (m *GroupPolicyObjectFile) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// Serialize serializes information the current object
func (m *GroupPolicyObjectFile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("groupPolicyObjectId", m.GetGroupPolicyObjectId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ouDistinguishedName", m.GetOuDistinguishedName())
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContent sets the content property value. The Group Policy Object file content.
func (m *GroupPolicyObjectFile) SetContent(value *string)() {
    m.content = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time at which the GroupPolicy was first uploaded.
func (m *GroupPolicyObjectFile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetGroupPolicyObjectId sets the groupPolicyObjectId property value. The Group Policy Object GUID from GPO Xml content
func (m *GroupPolicyObjectFile) SetGroupPolicyObjectId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.groupPolicyObjectId = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time at which the GroupPolicyObjectFile was last modified.
func (m *GroupPolicyObjectFile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetOuDistinguishedName sets the ouDistinguishedName property value. The distinguished name of the OU.
func (m *GroupPolicyObjectFile) SetOuDistinguishedName(value *string)() {
    m.ouDistinguishedName = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. The list of scope tags for the configuration.
func (m *GroupPolicyObjectFile) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
