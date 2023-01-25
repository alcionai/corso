package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosWebContentFilterAutoFilter 
type IosWebContentFilterAutoFilter struct {
    IosWebContentFilterBase
    // Additional URLs allowed for access
    allowedUrls []string
    // Additional URLs blocked for access
    blockedUrls []string
}
// NewIosWebContentFilterAutoFilter instantiates a new IosWebContentFilterAutoFilter and sets the default values.
func NewIosWebContentFilterAutoFilter()(*IosWebContentFilterAutoFilter) {
    m := &IosWebContentFilterAutoFilter{
        IosWebContentFilterBase: *NewIosWebContentFilterBase(),
    }
    odataTypeValue := "#microsoft.graph.iosWebContentFilterAutoFilter";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateIosWebContentFilterAutoFilterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosWebContentFilterAutoFilterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosWebContentFilterAutoFilter(), nil
}
// GetAllowedUrls gets the allowedUrls property value. Additional URLs allowed for access
func (m *IosWebContentFilterAutoFilter) GetAllowedUrls()([]string) {
    return m.allowedUrls
}
// GetBlockedUrls gets the blockedUrls property value. Additional URLs blocked for access
func (m *IosWebContentFilterAutoFilter) GetBlockedUrls()([]string) {
    return m.blockedUrls
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosWebContentFilterAutoFilter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IosWebContentFilterBase.GetFieldDeserializers()
    res["allowedUrls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAllowedUrls(res)
        }
        return nil
    }
    res["blockedUrls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetBlockedUrls(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *IosWebContentFilterAutoFilter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IosWebContentFilterBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedUrls() != nil {
        err = writer.WriteCollectionOfStringValues("allowedUrls", m.GetAllowedUrls())
        if err != nil {
            return err
        }
    }
    if m.GetBlockedUrls() != nil {
        err = writer.WriteCollectionOfStringValues("blockedUrls", m.GetBlockedUrls())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedUrls sets the allowedUrls property value. Additional URLs allowed for access
func (m *IosWebContentFilterAutoFilter) SetAllowedUrls(value []string)() {
    m.allowedUrls = value
}
// SetBlockedUrls sets the blockedUrls property value. Additional URLs blocked for access
func (m *IosWebContentFilterAutoFilter) SetBlockedUrls(value []string)() {
    m.blockedUrls = value
}
