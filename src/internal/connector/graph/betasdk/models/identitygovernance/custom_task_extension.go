package identitygovernance

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// CustomTaskExtension 
type CustomTaskExtension struct {
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CustomCalloutExtension
    // The callback configuration for a custom extension.
    callbackConfiguration ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CustomExtensionCallbackConfigurationable
    // The unique identifier of the Azure AD user that created the custom task extension.Supports $filter(eq, ne) and $expand.
    createdBy ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable
    // When the custom task extension was created.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The unique identifier of the Azure AD user that modified the custom task extension last.Supports $filter(eq, ne) and $expand.
    lastModifiedBy ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable
    // When the custom extension was last modified.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewCustomTaskExtension instantiates a new CustomTaskExtension and sets the default values.
func NewCustomTaskExtension()(*CustomTaskExtension) {
    m := &CustomTaskExtension{
        CustomCalloutExtension: *ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.NewCustomCalloutExtension(),
    }
    odataTypeValue := "#microsoft.graph.identityGovernance.customTaskExtension";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCustomTaskExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomTaskExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomTaskExtension(), nil
}
// GetCallbackConfiguration gets the callbackConfiguration property value. The callback configuration for a custom extension.
func (m *CustomTaskExtension) GetCallbackConfiguration()(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CustomExtensionCallbackConfigurationable) {
    return m.callbackConfiguration
}
// GetCreatedBy gets the createdBy property value. The unique identifier of the Azure AD user that created the custom task extension.Supports $filter(eq, ne) and $expand.
func (m *CustomTaskExtension) GetCreatedBy()(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. When the custom task extension was created.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *CustomTaskExtension) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomTaskExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomCalloutExtension.GetFieldDeserializers()
    res["callbackConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateCustomExtensionCallbackConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallbackConfiguration(val.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CustomExtensionCallbackConfigurationable))
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable))
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
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable))
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
    return res
}
// GetLastModifiedBy gets the lastModifiedBy property value. The unique identifier of the Azure AD user that modified the custom task extension last.Supports $filter(eq, ne) and $expand.
func (m *CustomTaskExtension) GetLastModifiedBy()(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. When the custom extension was last modified.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *CustomTaskExtension) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// Serialize serializes information the current object
func (m *CustomTaskExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomCalloutExtension.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("callbackConfiguration", m.GetCallbackConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
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
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
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
    return nil
}
// SetCallbackConfiguration sets the callbackConfiguration property value. The callback configuration for a custom extension.
func (m *CustomTaskExtension) SetCallbackConfiguration(value ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CustomExtensionCallbackConfigurationable)() {
    m.callbackConfiguration = value
}
// SetCreatedBy sets the createdBy property value. The unique identifier of the Azure AD user that created the custom task extension.Supports $filter(eq, ne) and $expand.
func (m *CustomTaskExtension) SetCreatedBy(value ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. When the custom task extension was created.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *CustomTaskExtension) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The unique identifier of the Azure AD user that modified the custom task extension last.Supports $filter(eq, ne) and $expand.
func (m *CustomTaskExtension) SetLastModifiedBy(value ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.Userable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. When the custom extension was last modified.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *CustomTaskExtension) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
