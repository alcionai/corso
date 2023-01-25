package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsResourcePerformanceable 
type UserExperienceAnalyticsResourcePerformanceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAverageSpikeTimeScore()(*int32)
    GetCpuSpikeTimePercentage()(*float64)
    GetCpuSpikeTimePercentageThreshold()(*float64)
    GetCpuSpikeTimeScore()(*int32)
    GetDeviceCount()(*int64)
    GetDeviceId()(*string)
    GetDeviceName()(*string)
    GetDeviceResourcePerformanceScore()(*int32)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetRamSpikeTimePercentage()(*float64)
    GetRamSpikeTimePercentageThreshold()(*float64)
    GetRamSpikeTimeScore()(*int32)
    SetAverageSpikeTimeScore(value *int32)()
    SetCpuSpikeTimePercentage(value *float64)()
    SetCpuSpikeTimePercentageThreshold(value *float64)()
    SetCpuSpikeTimeScore(value *int32)()
    SetDeviceCount(value *int64)()
    SetDeviceId(value *string)()
    SetDeviceName(value *string)()
    SetDeviceResourcePerformanceScore(value *int32)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetRamSpikeTimePercentage(value *float64)()
    SetRamSpikeTimePercentageThreshold(value *float64)()
    SetRamSpikeTimeScore(value *int32)()
}
