package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosBookmark iOS URL bookmark
type IosBookmark struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The folder into which the bookmark should be added in Safari
    bookmarkFolder *string
    // The display name of the bookmark
    displayName *string
    // The OdataType property
    odataType *string
    // URL allowed to access
    url *string
}
// NewIosBookmark instantiates a new iosBookmark and sets the default values.
func NewIosBookmark()(*IosBookmark) {
    m := &IosBookmark{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIosBookmarkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosBookmarkFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosBookmark(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IosBookmark) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetBookmarkFolder gets the bookmarkFolder property value. The folder into which the bookmark should be added in Safari
func (m *IosBookmark) GetBookmarkFolder()(*string) {
    return m.bookmarkFolder
}
// GetDisplayName gets the displayName property value. The display name of the bookmark
func (m *IosBookmark) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosBookmark) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bookmarkFolder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBookmarkFolder(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *IosBookmark) GetOdataType()(*string) {
    return m.odataType
}
// GetUrl gets the url property value. URL allowed to access
func (m *IosBookmark) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *IosBookmark) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("bookmarkFolder", m.GetBookmarkFolder())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
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
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *IosBookmark) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetBookmarkFolder sets the bookmarkFolder property value. The folder into which the bookmark should be added in Safari
func (m *IosBookmark) SetBookmarkFolder(value *string)() {
    m.bookmarkFolder = value
}
// SetDisplayName sets the displayName property value. The display name of the bookmark
func (m *IosBookmark) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *IosBookmark) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUrl sets the url property value. URL allowed to access
func (m *IosBookmark) SetUrl(value *string)() {
    m.url = value
}
