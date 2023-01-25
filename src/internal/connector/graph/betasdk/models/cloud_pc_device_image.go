package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcDeviceImage 
type CloudPcDeviceImage struct {
    Entity
    // The image's display name.
    displayName *string
    // The date the image became unavailable.
    expirationDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The data and time that the image was last modified. The time is shown in ISO 8601 format and  Coordinated Universal Time (UTC) time. For example, midnight UTC on Jan 1, 2014 appears as '2014-01-01T00:00:00Z'.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The image's operating system. For example: Windows 10 Enterprise.
    operatingSystem *string
    // The image's OS build version. For example: 1909.
    osBuildNumber *string
    // The OS status of this image. Possible values are: supported, supportedWithWarning, unknownFutureValue.
    osStatus *CloudPcDeviceImageOsStatus
    // The ID of the source image resource on Azure. Required format: '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}'.
    sourceImageResourceId *string
    // The status of the image on Cloud PC. Possible values are: pending, ready, failed.
    status *CloudPcDeviceImageStatus
    // The details of the image's status, which indicates why the upload failed, if applicable. Possible values are: internalServerError, sourceImageNotFound, osVersionNotSupported, sourceImageInvalid, and sourceImageNotGeneralized.
    statusDetails *CloudPcDeviceImageStatusDetails
    // The image version. For example: 0.0.1, 1.5.13.
    version *string
}
// NewCloudPcDeviceImage instantiates a new CloudPcDeviceImage and sets the default values.
func NewCloudPcDeviceImage()(*CloudPcDeviceImage) {
    m := &CloudPcDeviceImage{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPcDeviceImageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcDeviceImageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcDeviceImage(), nil
}
// GetDisplayName gets the displayName property value. The image's display name.
func (m *CloudPcDeviceImage) GetDisplayName()(*string) {
    return m.displayName
}
// GetExpirationDate gets the expirationDate property value. The date the image became unavailable.
func (m *CloudPcDeviceImage) GetExpirationDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.expirationDate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcDeviceImage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["expirationDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDate(val)
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
    res["operatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystem(val)
        }
        return nil
    }
    res["osBuildNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsBuildNumber(val)
        }
        return nil
    }
    res["osStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcDeviceImageOsStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsStatus(val.(*CloudPcDeviceImageOsStatus))
        }
        return nil
    }
    res["sourceImageResourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceImageResourceId(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcDeviceImageStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*CloudPcDeviceImageStatus))
        }
        return nil
    }
    res["statusDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcDeviceImageStatusDetails)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusDetails(val.(*CloudPcDeviceImageStatusDetails))
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The data and time that the image was last modified. The time is shown in ISO 8601 format and  Coordinated Universal Time (UTC) time. For example, midnight UTC on Jan 1, 2014 appears as '2014-01-01T00:00:00Z'.
func (m *CloudPcDeviceImage) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetOperatingSystem gets the operatingSystem property value. The image's operating system. For example: Windows 10 Enterprise.
func (m *CloudPcDeviceImage) GetOperatingSystem()(*string) {
    return m.operatingSystem
}
// GetOsBuildNumber gets the osBuildNumber property value. The image's OS build version. For example: 1909.
func (m *CloudPcDeviceImage) GetOsBuildNumber()(*string) {
    return m.osBuildNumber
}
// GetOsStatus gets the osStatus property value. The OS status of this image. Possible values are: supported, supportedWithWarning, unknownFutureValue.
func (m *CloudPcDeviceImage) GetOsStatus()(*CloudPcDeviceImageOsStatus) {
    return m.osStatus
}
// GetSourceImageResourceId gets the sourceImageResourceId property value. The ID of the source image resource on Azure. Required format: '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}'.
func (m *CloudPcDeviceImage) GetSourceImageResourceId()(*string) {
    return m.sourceImageResourceId
}
// GetStatus gets the status property value. The status of the image on Cloud PC. Possible values are: pending, ready, failed.
func (m *CloudPcDeviceImage) GetStatus()(*CloudPcDeviceImageStatus) {
    return m.status
}
// GetStatusDetails gets the statusDetails property value. The details of the image's status, which indicates why the upload failed, if applicable. Possible values are: internalServerError, sourceImageNotFound, osVersionNotSupported, sourceImageInvalid, and sourceImageNotGeneralized.
func (m *CloudPcDeviceImage) GetStatusDetails()(*CloudPcDeviceImageStatusDetails) {
    return m.statusDetails
}
// GetVersion gets the version property value. The image version. For example: 0.0.1, 1.5.13.
func (m *CloudPcDeviceImage) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *CloudPcDeviceImage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("expirationDate", m.GetExpirationDate())
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
        err = writer.WriteStringValue("operatingSystem", m.GetOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osBuildNumber", m.GetOsBuildNumber())
        if err != nil {
            return err
        }
    }
    if m.GetOsStatus() != nil {
        cast := (*m.GetOsStatus()).String()
        err = writer.WriteStringValue("osStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sourceImageResourceId", m.GetSourceImageResourceId())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatusDetails() != nil {
        cast := (*m.GetStatusDetails()).String()
        err = writer.WriteStringValue("statusDetails", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The image's display name.
func (m *CloudPcDeviceImage) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExpirationDate sets the expirationDate property value. The date the image became unavailable.
func (m *CloudPcDeviceImage) SetExpirationDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.expirationDate = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The data and time that the image was last modified. The time is shown in ISO 8601 format and  Coordinated Universal Time (UTC) time. For example, midnight UTC on Jan 1, 2014 appears as '2014-01-01T00:00:00Z'.
func (m *CloudPcDeviceImage) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetOperatingSystem sets the operatingSystem property value. The image's operating system. For example: Windows 10 Enterprise.
func (m *CloudPcDeviceImage) SetOperatingSystem(value *string)() {
    m.operatingSystem = value
}
// SetOsBuildNumber sets the osBuildNumber property value. The image's OS build version. For example: 1909.
func (m *CloudPcDeviceImage) SetOsBuildNumber(value *string)() {
    m.osBuildNumber = value
}
// SetOsStatus sets the osStatus property value. The OS status of this image. Possible values are: supported, supportedWithWarning, unknownFutureValue.
func (m *CloudPcDeviceImage) SetOsStatus(value *CloudPcDeviceImageOsStatus)() {
    m.osStatus = value
}
// SetSourceImageResourceId sets the sourceImageResourceId property value. The ID of the source image resource on Azure. Required format: '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}'.
func (m *CloudPcDeviceImage) SetSourceImageResourceId(value *string)() {
    m.sourceImageResourceId = value
}
// SetStatus sets the status property value. The status of the image on Cloud PC. Possible values are: pending, ready, failed.
func (m *CloudPcDeviceImage) SetStatus(value *CloudPcDeviceImageStatus)() {
    m.status = value
}
// SetStatusDetails sets the statusDetails property value. The details of the image's status, which indicates why the upload failed, if applicable. Possible values are: internalServerError, sourceImageNotFound, osVersionNotSupported, sourceImageInvalid, and sourceImageNotGeneralized.
func (m *CloudPcDeviceImage) SetStatusDetails(value *CloudPcDeviceImageStatusDetails)() {
    m.statusDetails = value
}
// SetVersion sets the version property value. The image version. For example: 0.0.1, 1.5.13.
func (m *CloudPcDeviceImage) SetVersion(value *string)() {
    m.version = value
}
