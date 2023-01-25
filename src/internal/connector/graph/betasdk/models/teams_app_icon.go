package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamsAppIcon 
type TeamsAppIcon struct {
    Entity
    // The contents of the app icon if the icon is hosted within the Teams infrastructure.
    hostedContent TeamworkHostedContentable
    // The web URL that can be used for downloading the image.
    webUrl *string
}
// NewTeamsAppIcon instantiates a new teamsAppIcon and sets the default values.
func NewTeamsAppIcon()(*TeamsAppIcon) {
    m := &TeamsAppIcon{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamsAppIconFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamsAppIconFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamsAppIcon(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsAppIcon) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["hostedContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkHostedContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostedContent(val.(TeamworkHostedContentable))
        }
        return nil
    }
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetHostedContent gets the hostedContent property value. The contents of the app icon if the icon is hosted within the Teams infrastructure.
func (m *TeamsAppIcon) GetHostedContent()(TeamworkHostedContentable) {
    return m.hostedContent
}
// GetWebUrl gets the webUrl property value. The web URL that can be used for downloading the image.
func (m *TeamsAppIcon) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *TeamsAppIcon) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("hostedContent", m.GetHostedContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetHostedContent sets the hostedContent property value. The contents of the app icon if the icon is hosted within the Teams infrastructure.
func (m *TeamsAppIcon) SetHostedContent(value TeamworkHostedContentable)() {
    m.hostedContent = value
}
// SetWebUrl sets the webUrl property value. The web URL that can be used for downloading the image.
func (m *TeamsAppIcon) SetWebUrl(value *string)() {
    m.webUrl = value
}
