package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChromeOSOnboardingSettings entity that represents a Chromebook tenant settings
type ChromeOSOnboardingSettings struct {
    Entity
    // The ChromebookTenant's LastDirectorySyncDateTime
    lastDirectorySyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The ChromebookTenant's LastModifiedDateTime
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The onboarding status of the tenant.
    onboardingStatus *OnboardingStatus
    // The ChromebookTenant's OwnerUserPrincipalName
    ownerUserPrincipalName *string
}
// NewChromeOSOnboardingSettings instantiates a new chromeOSOnboardingSettings and sets the default values.
func NewChromeOSOnboardingSettings()(*ChromeOSOnboardingSettings) {
    m := &ChromeOSOnboardingSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateChromeOSOnboardingSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChromeOSOnboardingSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChromeOSOnboardingSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChromeOSOnboardingSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lastDirectorySyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastDirectorySyncDateTime(val)
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
    res["onboardingStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnboardingStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnboardingStatus(val.(*OnboardingStatus))
        }
        return nil
    }
    res["ownerUserPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetLastDirectorySyncDateTime gets the lastDirectorySyncDateTime property value. The ChromebookTenant's LastDirectorySyncDateTime
func (m *ChromeOSOnboardingSettings) GetLastDirectorySyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastDirectorySyncDateTime
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The ChromebookTenant's LastModifiedDateTime
func (m *ChromeOSOnboardingSettings) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetOnboardingStatus gets the onboardingStatus property value. The onboarding status of the tenant.
func (m *ChromeOSOnboardingSettings) GetOnboardingStatus()(*OnboardingStatus) {
    return m.onboardingStatus
}
// GetOwnerUserPrincipalName gets the ownerUserPrincipalName property value. The ChromebookTenant's OwnerUserPrincipalName
func (m *ChromeOSOnboardingSettings) GetOwnerUserPrincipalName()(*string) {
    return m.ownerUserPrincipalName
}
// Serialize serializes information the current object
func (m *ChromeOSOnboardingSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("lastDirectorySyncDateTime", m.GetLastDirectorySyncDateTime())
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
    if m.GetOnboardingStatus() != nil {
        cast := (*m.GetOnboardingStatus()).String()
        err = writer.WriteStringValue("onboardingStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ownerUserPrincipalName", m.GetOwnerUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLastDirectorySyncDateTime sets the lastDirectorySyncDateTime property value. The ChromebookTenant's LastDirectorySyncDateTime
func (m *ChromeOSOnboardingSettings) SetLastDirectorySyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastDirectorySyncDateTime = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The ChromebookTenant's LastModifiedDateTime
func (m *ChromeOSOnboardingSettings) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetOnboardingStatus sets the onboardingStatus property value. The onboarding status of the tenant.
func (m *ChromeOSOnboardingSettings) SetOnboardingStatus(value *OnboardingStatus)() {
    m.onboardingStatus = value
}
// SetOwnerUserPrincipalName sets the ownerUserPrincipalName property value. The ChromebookTenant's OwnerUserPrincipalName
func (m *ChromeOSOnboardingSettings) SetOwnerUserPrincipalName(value *string)() {
    m.ownerUserPrincipalName = value
}
