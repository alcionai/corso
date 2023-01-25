package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ZebraFotaArtifact describes a single artifact for a specific device model.
type ZebraFotaArtifact struct {
    Entity
    // The version of the Board Support Package (BSP. E.g.: 01.18.02.00)
    boardSupportPackageVersion *string
    // Artifact description. (e.g.: `LifeGuard Update 98 (released 24-September-2021)
    description *string
    // Applicable device model (e.g.: TC8300)
    deviceModel *string
    // Artifact OS version (e.g.: 8.1.0)
    osVersion *string
    // Artifact patch version (e.g.: U00)
    patchVersion *string
    // Artifact release notes URL (e.g.: https://www.zebra.com/<filename.pdf>)
    releaseNotesUrl *string
}
// NewZebraFotaArtifact instantiates a new zebraFotaArtifact and sets the default values.
func NewZebraFotaArtifact()(*ZebraFotaArtifact) {
    m := &ZebraFotaArtifact{
        Entity: *NewEntity(),
    }
    return m
}
// CreateZebraFotaArtifactFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateZebraFotaArtifactFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewZebraFotaArtifact(), nil
}
// GetBoardSupportPackageVersion gets the boardSupportPackageVersion property value. The version of the Board Support Package (BSP. E.g.: 01.18.02.00)
func (m *ZebraFotaArtifact) GetBoardSupportPackageVersion()(*string) {
    return m.boardSupportPackageVersion
}
// GetDescription gets the description property value. Artifact description. (e.g.: `LifeGuard Update 98 (released 24-September-2021)
func (m *ZebraFotaArtifact) GetDescription()(*string) {
    return m.description
}
// GetDeviceModel gets the deviceModel property value. Applicable device model (e.g.: TC8300)
func (m *ZebraFotaArtifact) GetDeviceModel()(*string) {
    return m.deviceModel
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ZebraFotaArtifact) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["boardSupportPackageVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBoardSupportPackageVersion(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["deviceModel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceModel(val)
        }
        return nil
    }
    res["osVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    res["patchVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPatchVersion(val)
        }
        return nil
    }
    res["releaseNotesUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReleaseNotesUrl(val)
        }
        return nil
    }
    return res
}
// GetOsVersion gets the osVersion property value. Artifact OS version (e.g.: 8.1.0)
func (m *ZebraFotaArtifact) GetOsVersion()(*string) {
    return m.osVersion
}
// GetPatchVersion gets the patchVersion property value. Artifact patch version (e.g.: U00)
func (m *ZebraFotaArtifact) GetPatchVersion()(*string) {
    return m.patchVersion
}
// GetReleaseNotesUrl gets the releaseNotesUrl property value. Artifact release notes URL (e.g.: https://www.zebra.com/<filename.pdf>)
func (m *ZebraFotaArtifact) GetReleaseNotesUrl()(*string) {
    return m.releaseNotesUrl
}
// Serialize serializes information the current object
func (m *ZebraFotaArtifact) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("boardSupportPackageVersion", m.GetBoardSupportPackageVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceModel", m.GetDeviceModel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osVersion", m.GetOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("patchVersion", m.GetPatchVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("releaseNotesUrl", m.GetReleaseNotesUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBoardSupportPackageVersion sets the boardSupportPackageVersion property value. The version of the Board Support Package (BSP. E.g.: 01.18.02.00)
func (m *ZebraFotaArtifact) SetBoardSupportPackageVersion(value *string)() {
    m.boardSupportPackageVersion = value
}
// SetDescription sets the description property value. Artifact description. (e.g.: `LifeGuard Update 98 (released 24-September-2021)
func (m *ZebraFotaArtifact) SetDescription(value *string)() {
    m.description = value
}
// SetDeviceModel sets the deviceModel property value. Applicable device model (e.g.: TC8300)
func (m *ZebraFotaArtifact) SetDeviceModel(value *string)() {
    m.deviceModel = value
}
// SetOsVersion sets the osVersion property value. Artifact OS version (e.g.: 8.1.0)
func (m *ZebraFotaArtifact) SetOsVersion(value *string)() {
    m.osVersion = value
}
// SetPatchVersion sets the patchVersion property value. Artifact patch version (e.g.: U00)
func (m *ZebraFotaArtifact) SetPatchVersion(value *string)() {
    m.patchVersion = value
}
// SetReleaseNotesUrl sets the releaseNotesUrl property value. Artifact release notes URL (e.g.: https://www.zebra.com/<filename.pdf>)
func (m *ZebraFotaArtifact) SetReleaseNotesUrl(value *string)() {
    m.releaseNotesUrl = value
}
