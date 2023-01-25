package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsDeliveryOptimizationConfiguration 
type WindowsDeliveryOptimizationConfiguration struct {
    DeviceConfiguration
    // Specifies number of seconds to delay an HTTP source in a background download that is allowed to use peer-to-peer. Valid values 0 to 4294967295
    backgroundDownloadFromHttpDelayInSeconds *int64
    // Specifies foreground and background bandwidth usage using percentages, absolutes, or hours.
    bandwidthMode DeliveryOptimizationBandwidthable
    // Specifies number of seconds to delay a fall back from cache servers to an HTTP source for a background download. Valid values 0 to 2592000.
    cacheServerBackgroundDownloadFallbackToHttpDelayInSeconds *int32
    // Specifies number of seconds to delay a fall back from cache servers to an HTTP source for a foreground download. Valid values 0 to 2592000.​
    cacheServerForegroundDownloadFallbackToHttpDelayInSeconds *int32
    // Specifies cache servers host names.
    cacheServerHostNames []string
    // Delivery optimization mode for peer distribution
    deliveryOptimizationMode *WindowsDeliveryOptimizationMode
    // Specifies number of seconds to delay an HTTP source in a foreground download that is allowed to use peer-to-peer (0-86400). Valid values 0 to 86400
    foregroundDownloadFromHttpDelayInSeconds *int64
    // Specifies to restrict peer selection to a specfic source.
    groupIdSource DeliveryOptimizationGroupIdSourceable
    // Specifies the maximum time in days that each file is held in the Delivery Optimization cache after downloading successfully (0-3650). Valid values 0 to 3650
    maximumCacheAgeInDays *int32
    // Specifies the maximum cache size that Delivery Optimization either as a percentage or in GB.
    maximumCacheSize DeliveryOptimizationMaxCacheSizeable
    // Specifies the minimum battery percentage to allow the device to upload data (0-100). Valid values 0 to 100
    minimumBatteryPercentageAllowedToUpload *int32
    // Specifies the minimum disk size in GB to use Peer Caching (1-100000). Valid values 1 to 100000
    minimumDiskSizeAllowedToPeerInGigabytes *int32
    // Specifies the minimum content file size in MB enabled to use Peer Caching (1-100000). Valid values 1 to 100000
    minimumFileSizeToCacheInMegabytes *int32
    // Specifies the minimum RAM size in GB to use Peer Caching (1-100000). Valid values 1 to 100000
    minimumRamAllowedToPeerInGigabytes *int32
    // Specifies the drive that Delivery Optimization should use for its cache.
    modifyCacheLocation *string
    // Values to restrict peer selection by.
    restrictPeerSelectionBy *DeliveryOptimizationRestrictPeerSelectionByOptions
    // Possible values of a property
    vpnPeerCaching *Enablement
}
// NewWindowsDeliveryOptimizationConfiguration instantiates a new WindowsDeliveryOptimizationConfiguration and sets the default values.
func NewWindowsDeliveryOptimizationConfiguration()(*WindowsDeliveryOptimizationConfiguration) {
    m := &WindowsDeliveryOptimizationConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsDeliveryOptimizationConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsDeliveryOptimizationConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsDeliveryOptimizationConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsDeliveryOptimizationConfiguration(), nil
}
// GetBackgroundDownloadFromHttpDelayInSeconds gets the backgroundDownloadFromHttpDelayInSeconds property value. Specifies number of seconds to delay an HTTP source in a background download that is allowed to use peer-to-peer. Valid values 0 to 4294967295
func (m *WindowsDeliveryOptimizationConfiguration) GetBackgroundDownloadFromHttpDelayInSeconds()(*int64) {
    return m.backgroundDownloadFromHttpDelayInSeconds
}
// GetBandwidthMode gets the bandwidthMode property value. Specifies foreground and background bandwidth usage using percentages, absolutes, or hours.
func (m *WindowsDeliveryOptimizationConfiguration) GetBandwidthMode()(DeliveryOptimizationBandwidthable) {
    return m.bandwidthMode
}
// GetCacheServerBackgroundDownloadFallbackToHttpDelayInSeconds gets the cacheServerBackgroundDownloadFallbackToHttpDelayInSeconds property value. Specifies number of seconds to delay a fall back from cache servers to an HTTP source for a background download. Valid values 0 to 2592000.
func (m *WindowsDeliveryOptimizationConfiguration) GetCacheServerBackgroundDownloadFallbackToHttpDelayInSeconds()(*int32) {
    return m.cacheServerBackgroundDownloadFallbackToHttpDelayInSeconds
}
// GetCacheServerForegroundDownloadFallbackToHttpDelayInSeconds gets the cacheServerForegroundDownloadFallbackToHttpDelayInSeconds property value. Specifies number of seconds to delay a fall back from cache servers to an HTTP source for a foreground download. Valid values 0 to 2592000.​
func (m *WindowsDeliveryOptimizationConfiguration) GetCacheServerForegroundDownloadFallbackToHttpDelayInSeconds()(*int32) {
    return m.cacheServerForegroundDownloadFallbackToHttpDelayInSeconds
}
// GetCacheServerHostNames gets the cacheServerHostNames property value. Specifies cache servers host names.
func (m *WindowsDeliveryOptimizationConfiguration) GetCacheServerHostNames()([]string) {
    return m.cacheServerHostNames
}
// GetDeliveryOptimizationMode gets the deliveryOptimizationMode property value. Delivery optimization mode for peer distribution
func (m *WindowsDeliveryOptimizationConfiguration) GetDeliveryOptimizationMode()(*WindowsDeliveryOptimizationMode) {
    return m.deliveryOptimizationMode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsDeliveryOptimizationConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["backgroundDownloadFromHttpDelayInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBackgroundDownloadFromHttpDelayInSeconds(val)
        }
        return nil
    }
    res["bandwidthMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeliveryOptimizationBandwidthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBandwidthMode(val.(DeliveryOptimizationBandwidthable))
        }
        return nil
    }
    res["cacheServerBackgroundDownloadFallbackToHttpDelayInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCacheServerBackgroundDownloadFallbackToHttpDelayInSeconds(val)
        }
        return nil
    }
    res["cacheServerForegroundDownloadFallbackToHttpDelayInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCacheServerForegroundDownloadFallbackToHttpDelayInSeconds(val)
        }
        return nil
    }
    res["cacheServerHostNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetCacheServerHostNames(res)
        }
        return nil
    }
    res["deliveryOptimizationMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsDeliveryOptimizationMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeliveryOptimizationMode(val.(*WindowsDeliveryOptimizationMode))
        }
        return nil
    }
    res["foregroundDownloadFromHttpDelayInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForegroundDownloadFromHttpDelayInSeconds(val)
        }
        return nil
    }
    res["groupIdSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeliveryOptimizationGroupIdSourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupIdSource(val.(DeliveryOptimizationGroupIdSourceable))
        }
        return nil
    }
    res["maximumCacheAgeInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumCacheAgeInDays(val)
        }
        return nil
    }
    res["maximumCacheSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeliveryOptimizationMaxCacheSizeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumCacheSize(val.(DeliveryOptimizationMaxCacheSizeable))
        }
        return nil
    }
    res["minimumBatteryPercentageAllowedToUpload"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumBatteryPercentageAllowedToUpload(val)
        }
        return nil
    }
    res["minimumDiskSizeAllowedToPeerInGigabytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumDiskSizeAllowedToPeerInGigabytes(val)
        }
        return nil
    }
    res["minimumFileSizeToCacheInMegabytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumFileSizeToCacheInMegabytes(val)
        }
        return nil
    }
    res["minimumRamAllowedToPeerInGigabytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumRamAllowedToPeerInGigabytes(val)
        }
        return nil
    }
    res["modifyCacheLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModifyCacheLocation(val)
        }
        return nil
    }
    res["restrictPeerSelectionBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeliveryOptimizationRestrictPeerSelectionByOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestrictPeerSelectionBy(val.(*DeliveryOptimizationRestrictPeerSelectionByOptions))
        }
        return nil
    }
    res["vpnPeerCaching"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnablement)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVpnPeerCaching(val.(*Enablement))
        }
        return nil
    }
    return res
}
// GetForegroundDownloadFromHttpDelayInSeconds gets the foregroundDownloadFromHttpDelayInSeconds property value. Specifies number of seconds to delay an HTTP source in a foreground download that is allowed to use peer-to-peer (0-86400). Valid values 0 to 86400
func (m *WindowsDeliveryOptimizationConfiguration) GetForegroundDownloadFromHttpDelayInSeconds()(*int64) {
    return m.foregroundDownloadFromHttpDelayInSeconds
}
// GetGroupIdSource gets the groupIdSource property value. Specifies to restrict peer selection to a specfic source.
func (m *WindowsDeliveryOptimizationConfiguration) GetGroupIdSource()(DeliveryOptimizationGroupIdSourceable) {
    return m.groupIdSource
}
// GetMaximumCacheAgeInDays gets the maximumCacheAgeInDays property value. Specifies the maximum time in days that each file is held in the Delivery Optimization cache after downloading successfully (0-3650). Valid values 0 to 3650
func (m *WindowsDeliveryOptimizationConfiguration) GetMaximumCacheAgeInDays()(*int32) {
    return m.maximumCacheAgeInDays
}
// GetMaximumCacheSize gets the maximumCacheSize property value. Specifies the maximum cache size that Delivery Optimization either as a percentage or in GB.
func (m *WindowsDeliveryOptimizationConfiguration) GetMaximumCacheSize()(DeliveryOptimizationMaxCacheSizeable) {
    return m.maximumCacheSize
}
// GetMinimumBatteryPercentageAllowedToUpload gets the minimumBatteryPercentageAllowedToUpload property value. Specifies the minimum battery percentage to allow the device to upload data (0-100). Valid values 0 to 100
func (m *WindowsDeliveryOptimizationConfiguration) GetMinimumBatteryPercentageAllowedToUpload()(*int32) {
    return m.minimumBatteryPercentageAllowedToUpload
}
// GetMinimumDiskSizeAllowedToPeerInGigabytes gets the minimumDiskSizeAllowedToPeerInGigabytes property value. Specifies the minimum disk size in GB to use Peer Caching (1-100000). Valid values 1 to 100000
func (m *WindowsDeliveryOptimizationConfiguration) GetMinimumDiskSizeAllowedToPeerInGigabytes()(*int32) {
    return m.minimumDiskSizeAllowedToPeerInGigabytes
}
// GetMinimumFileSizeToCacheInMegabytes gets the minimumFileSizeToCacheInMegabytes property value. Specifies the minimum content file size in MB enabled to use Peer Caching (1-100000). Valid values 1 to 100000
func (m *WindowsDeliveryOptimizationConfiguration) GetMinimumFileSizeToCacheInMegabytes()(*int32) {
    return m.minimumFileSizeToCacheInMegabytes
}
// GetMinimumRamAllowedToPeerInGigabytes gets the minimumRamAllowedToPeerInGigabytes property value. Specifies the minimum RAM size in GB to use Peer Caching (1-100000). Valid values 1 to 100000
func (m *WindowsDeliveryOptimizationConfiguration) GetMinimumRamAllowedToPeerInGigabytes()(*int32) {
    return m.minimumRamAllowedToPeerInGigabytes
}
// GetModifyCacheLocation gets the modifyCacheLocation property value. Specifies the drive that Delivery Optimization should use for its cache.
func (m *WindowsDeliveryOptimizationConfiguration) GetModifyCacheLocation()(*string) {
    return m.modifyCacheLocation
}
// GetRestrictPeerSelectionBy gets the restrictPeerSelectionBy property value. Values to restrict peer selection by.
func (m *WindowsDeliveryOptimizationConfiguration) GetRestrictPeerSelectionBy()(*DeliveryOptimizationRestrictPeerSelectionByOptions) {
    return m.restrictPeerSelectionBy
}
// GetVpnPeerCaching gets the vpnPeerCaching property value. Possible values of a property
func (m *WindowsDeliveryOptimizationConfiguration) GetVpnPeerCaching()(*Enablement) {
    return m.vpnPeerCaching
}
// Serialize serializes information the current object
func (m *WindowsDeliveryOptimizationConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("backgroundDownloadFromHttpDelayInSeconds", m.GetBackgroundDownloadFromHttpDelayInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("bandwidthMode", m.GetBandwidthMode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("cacheServerBackgroundDownloadFallbackToHttpDelayInSeconds", m.GetCacheServerBackgroundDownloadFallbackToHttpDelayInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("cacheServerForegroundDownloadFallbackToHttpDelayInSeconds", m.GetCacheServerForegroundDownloadFallbackToHttpDelayInSeconds())
        if err != nil {
            return err
        }
    }
    if m.GetCacheServerHostNames() != nil {
        err = writer.WriteCollectionOfStringValues("cacheServerHostNames", m.GetCacheServerHostNames())
        if err != nil {
            return err
        }
    }
    if m.GetDeliveryOptimizationMode() != nil {
        cast := (*m.GetDeliveryOptimizationMode()).String()
        err = writer.WriteStringValue("deliveryOptimizationMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("foregroundDownloadFromHttpDelayInSeconds", m.GetForegroundDownloadFromHttpDelayInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("groupIdSource", m.GetGroupIdSource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("maximumCacheAgeInDays", m.GetMaximumCacheAgeInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("maximumCacheSize", m.GetMaximumCacheSize())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumBatteryPercentageAllowedToUpload", m.GetMinimumBatteryPercentageAllowedToUpload())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumDiskSizeAllowedToPeerInGigabytes", m.GetMinimumDiskSizeAllowedToPeerInGigabytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumFileSizeToCacheInMegabytes", m.GetMinimumFileSizeToCacheInMegabytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumRamAllowedToPeerInGigabytes", m.GetMinimumRamAllowedToPeerInGigabytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("modifyCacheLocation", m.GetModifyCacheLocation())
        if err != nil {
            return err
        }
    }
    if m.GetRestrictPeerSelectionBy() != nil {
        cast := (*m.GetRestrictPeerSelectionBy()).String()
        err = writer.WriteStringValue("restrictPeerSelectionBy", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetVpnPeerCaching() != nil {
        cast := (*m.GetVpnPeerCaching()).String()
        err = writer.WriteStringValue("vpnPeerCaching", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBackgroundDownloadFromHttpDelayInSeconds sets the backgroundDownloadFromHttpDelayInSeconds property value. Specifies number of seconds to delay an HTTP source in a background download that is allowed to use peer-to-peer. Valid values 0 to 4294967295
func (m *WindowsDeliveryOptimizationConfiguration) SetBackgroundDownloadFromHttpDelayInSeconds(value *int64)() {
    m.backgroundDownloadFromHttpDelayInSeconds = value
}
// SetBandwidthMode sets the bandwidthMode property value. Specifies foreground and background bandwidth usage using percentages, absolutes, or hours.
func (m *WindowsDeliveryOptimizationConfiguration) SetBandwidthMode(value DeliveryOptimizationBandwidthable)() {
    m.bandwidthMode = value
}
// SetCacheServerBackgroundDownloadFallbackToHttpDelayInSeconds sets the cacheServerBackgroundDownloadFallbackToHttpDelayInSeconds property value. Specifies number of seconds to delay a fall back from cache servers to an HTTP source for a background download. Valid values 0 to 2592000.
func (m *WindowsDeliveryOptimizationConfiguration) SetCacheServerBackgroundDownloadFallbackToHttpDelayInSeconds(value *int32)() {
    m.cacheServerBackgroundDownloadFallbackToHttpDelayInSeconds = value
}
// SetCacheServerForegroundDownloadFallbackToHttpDelayInSeconds sets the cacheServerForegroundDownloadFallbackToHttpDelayInSeconds property value. Specifies number of seconds to delay a fall back from cache servers to an HTTP source for a foreground download. Valid values 0 to 2592000.​
func (m *WindowsDeliveryOptimizationConfiguration) SetCacheServerForegroundDownloadFallbackToHttpDelayInSeconds(value *int32)() {
    m.cacheServerForegroundDownloadFallbackToHttpDelayInSeconds = value
}
// SetCacheServerHostNames sets the cacheServerHostNames property value. Specifies cache servers host names.
func (m *WindowsDeliveryOptimizationConfiguration) SetCacheServerHostNames(value []string)() {
    m.cacheServerHostNames = value
}
// SetDeliveryOptimizationMode sets the deliveryOptimizationMode property value. Delivery optimization mode for peer distribution
func (m *WindowsDeliveryOptimizationConfiguration) SetDeliveryOptimizationMode(value *WindowsDeliveryOptimizationMode)() {
    m.deliveryOptimizationMode = value
}
// SetForegroundDownloadFromHttpDelayInSeconds sets the foregroundDownloadFromHttpDelayInSeconds property value. Specifies number of seconds to delay an HTTP source in a foreground download that is allowed to use peer-to-peer (0-86400). Valid values 0 to 86400
func (m *WindowsDeliveryOptimizationConfiguration) SetForegroundDownloadFromHttpDelayInSeconds(value *int64)() {
    m.foregroundDownloadFromHttpDelayInSeconds = value
}
// SetGroupIdSource sets the groupIdSource property value. Specifies to restrict peer selection to a specfic source.
func (m *WindowsDeliveryOptimizationConfiguration) SetGroupIdSource(value DeliveryOptimizationGroupIdSourceable)() {
    m.groupIdSource = value
}
// SetMaximumCacheAgeInDays sets the maximumCacheAgeInDays property value. Specifies the maximum time in days that each file is held in the Delivery Optimization cache after downloading successfully (0-3650). Valid values 0 to 3650
func (m *WindowsDeliveryOptimizationConfiguration) SetMaximumCacheAgeInDays(value *int32)() {
    m.maximumCacheAgeInDays = value
}
// SetMaximumCacheSize sets the maximumCacheSize property value. Specifies the maximum cache size that Delivery Optimization either as a percentage or in GB.
func (m *WindowsDeliveryOptimizationConfiguration) SetMaximumCacheSize(value DeliveryOptimizationMaxCacheSizeable)() {
    m.maximumCacheSize = value
}
// SetMinimumBatteryPercentageAllowedToUpload sets the minimumBatteryPercentageAllowedToUpload property value. Specifies the minimum battery percentage to allow the device to upload data (0-100). Valid values 0 to 100
func (m *WindowsDeliveryOptimizationConfiguration) SetMinimumBatteryPercentageAllowedToUpload(value *int32)() {
    m.minimumBatteryPercentageAllowedToUpload = value
}
// SetMinimumDiskSizeAllowedToPeerInGigabytes sets the minimumDiskSizeAllowedToPeerInGigabytes property value. Specifies the minimum disk size in GB to use Peer Caching (1-100000). Valid values 1 to 100000
func (m *WindowsDeliveryOptimizationConfiguration) SetMinimumDiskSizeAllowedToPeerInGigabytes(value *int32)() {
    m.minimumDiskSizeAllowedToPeerInGigabytes = value
}
// SetMinimumFileSizeToCacheInMegabytes sets the minimumFileSizeToCacheInMegabytes property value. Specifies the minimum content file size in MB enabled to use Peer Caching (1-100000). Valid values 1 to 100000
func (m *WindowsDeliveryOptimizationConfiguration) SetMinimumFileSizeToCacheInMegabytes(value *int32)() {
    m.minimumFileSizeToCacheInMegabytes = value
}
// SetMinimumRamAllowedToPeerInGigabytes sets the minimumRamAllowedToPeerInGigabytes property value. Specifies the minimum RAM size in GB to use Peer Caching (1-100000). Valid values 1 to 100000
func (m *WindowsDeliveryOptimizationConfiguration) SetMinimumRamAllowedToPeerInGigabytes(value *int32)() {
    m.minimumRamAllowedToPeerInGigabytes = value
}
// SetModifyCacheLocation sets the modifyCacheLocation property value. Specifies the drive that Delivery Optimization should use for its cache.
func (m *WindowsDeliveryOptimizationConfiguration) SetModifyCacheLocation(value *string)() {
    m.modifyCacheLocation = value
}
// SetRestrictPeerSelectionBy sets the restrictPeerSelectionBy property value. Values to restrict peer selection by.
func (m *WindowsDeliveryOptimizationConfiguration) SetRestrictPeerSelectionBy(value *DeliveryOptimizationRestrictPeerSelectionByOptions)() {
    m.restrictPeerSelectionBy = value
}
// SetVpnPeerCaching sets the vpnPeerCaching property value. Possible values of a property
func (m *WindowsDeliveryOptimizationConfiguration) SetVpnPeerCaching(value *Enablement)() {
    m.vpnPeerCaching = value
}
