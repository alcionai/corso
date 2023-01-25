package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsDeliveryOptimizationConfigurationable 
type WindowsDeliveryOptimizationConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackgroundDownloadFromHttpDelayInSeconds()(*int64)
    GetBandwidthMode()(DeliveryOptimizationBandwidthable)
    GetCacheServerBackgroundDownloadFallbackToHttpDelayInSeconds()(*int32)
    GetCacheServerForegroundDownloadFallbackToHttpDelayInSeconds()(*int32)
    GetCacheServerHostNames()([]string)
    GetDeliveryOptimizationMode()(*WindowsDeliveryOptimizationMode)
    GetForegroundDownloadFromHttpDelayInSeconds()(*int64)
    GetGroupIdSource()(DeliveryOptimizationGroupIdSourceable)
    GetMaximumCacheAgeInDays()(*int32)
    GetMaximumCacheSize()(DeliveryOptimizationMaxCacheSizeable)
    GetMinimumBatteryPercentageAllowedToUpload()(*int32)
    GetMinimumDiskSizeAllowedToPeerInGigabytes()(*int32)
    GetMinimumFileSizeToCacheInMegabytes()(*int32)
    GetMinimumRamAllowedToPeerInGigabytes()(*int32)
    GetModifyCacheLocation()(*string)
    GetRestrictPeerSelectionBy()(*DeliveryOptimizationRestrictPeerSelectionByOptions)
    GetVpnPeerCaching()(*Enablement)
    SetBackgroundDownloadFromHttpDelayInSeconds(value *int64)()
    SetBandwidthMode(value DeliveryOptimizationBandwidthable)()
    SetCacheServerBackgroundDownloadFallbackToHttpDelayInSeconds(value *int32)()
    SetCacheServerForegroundDownloadFallbackToHttpDelayInSeconds(value *int32)()
    SetCacheServerHostNames(value []string)()
    SetDeliveryOptimizationMode(value *WindowsDeliveryOptimizationMode)()
    SetForegroundDownloadFromHttpDelayInSeconds(value *int64)()
    SetGroupIdSource(value DeliveryOptimizationGroupIdSourceable)()
    SetMaximumCacheAgeInDays(value *int32)()
    SetMaximumCacheSize(value DeliveryOptimizationMaxCacheSizeable)()
    SetMinimumBatteryPercentageAllowedToUpload(value *int32)()
    SetMinimumDiskSizeAllowedToPeerInGigabytes(value *int32)()
    SetMinimumFileSizeToCacheInMegabytes(value *int32)()
    SetMinimumRamAllowedToPeerInGigabytes(value *int32)()
    SetModifyCacheLocation(value *string)()
    SetRestrictPeerSelectionBy(value *DeliveryOptimizationRestrictPeerSelectionByOptions)()
    SetVpnPeerCaching(value *Enablement)()
}
