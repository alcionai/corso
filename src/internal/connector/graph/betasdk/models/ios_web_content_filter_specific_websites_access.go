package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosWebContentFilterSpecificWebsitesAccess 
type IosWebContentFilterSpecificWebsitesAccess struct {
    IosWebContentFilterBase
    // URL bookmarks which will be installed into built-in browser and user is only allowed to access websites through bookmarks. This collection can contain a maximum of 500 elements.
    specificWebsitesOnly []IosBookmarkable
    // URL bookmarks which will be installed into built-in browser and user is only allowed to access websites through bookmarks. This collection can contain a maximum of 500 elements.
    websiteList []IosBookmarkable
}
// NewIosWebContentFilterSpecificWebsitesAccess instantiates a new IosWebContentFilterSpecificWebsitesAccess and sets the default values.
func NewIosWebContentFilterSpecificWebsitesAccess()(*IosWebContentFilterSpecificWebsitesAccess) {
    m := &IosWebContentFilterSpecificWebsitesAccess{
        IosWebContentFilterBase: *NewIosWebContentFilterBase(),
    }
    odataTypeValue := "#microsoft.graph.iosWebContentFilterSpecificWebsitesAccess";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateIosWebContentFilterSpecificWebsitesAccessFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosWebContentFilterSpecificWebsitesAccessFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosWebContentFilterSpecificWebsitesAccess(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosWebContentFilterSpecificWebsitesAccess) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IosWebContentFilterBase.GetFieldDeserializers()
    res["specificWebsitesOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIosBookmarkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosBookmarkable, len(val))
            for i, v := range val {
                res[i] = v.(IosBookmarkable)
            }
            m.SetSpecificWebsitesOnly(res)
        }
        return nil
    }
    res["websiteList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIosBookmarkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosBookmarkable, len(val))
            for i, v := range val {
                res[i] = v.(IosBookmarkable)
            }
            m.SetWebsiteList(res)
        }
        return nil
    }
    return res
}
// GetSpecificWebsitesOnly gets the specificWebsitesOnly property value. URL bookmarks which will be installed into built-in browser and user is only allowed to access websites through bookmarks. This collection can contain a maximum of 500 elements.
func (m *IosWebContentFilterSpecificWebsitesAccess) GetSpecificWebsitesOnly()([]IosBookmarkable) {
    return m.specificWebsitesOnly
}
// GetWebsiteList gets the websiteList property value. URL bookmarks which will be installed into built-in browser and user is only allowed to access websites through bookmarks. This collection can contain a maximum of 500 elements.
func (m *IosWebContentFilterSpecificWebsitesAccess) GetWebsiteList()([]IosBookmarkable) {
    return m.websiteList
}
// Serialize serializes information the current object
func (m *IosWebContentFilterSpecificWebsitesAccess) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IosWebContentFilterBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSpecificWebsitesOnly() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSpecificWebsitesOnly()))
        for i, v := range m.GetSpecificWebsitesOnly() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("specificWebsitesOnly", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWebsiteList() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWebsiteList()))
        for i, v := range m.GetWebsiteList() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("websiteList", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSpecificWebsitesOnly sets the specificWebsitesOnly property value. URL bookmarks which will be installed into built-in browser and user is only allowed to access websites through bookmarks. This collection can contain a maximum of 500 elements.
func (m *IosWebContentFilterSpecificWebsitesAccess) SetSpecificWebsitesOnly(value []IosBookmarkable)() {
    m.specificWebsitesOnly = value
}
// SetWebsiteList sets the websiteList property value. URL bookmarks which will be installed into built-in browser and user is only allowed to access websites through bookmarks. This collection can contain a maximum of 500 elements.
func (m *IosWebContentFilterSpecificWebsitesAccess) SetWebsiteList(value []IosBookmarkable)() {
    m.websiteList = value
}
