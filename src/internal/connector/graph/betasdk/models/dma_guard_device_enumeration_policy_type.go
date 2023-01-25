package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type DmaGuardDeviceEnumerationPolicyType int

const (
    // Default value. Devices with DMA remapping incompatible drivers will only be enumerated after the user unlocks the screen.
    DEVICEDEFAULT_DMAGUARDDEVICEENUMERATIONPOLICYTYPE DmaGuardDeviceEnumerationPolicyType = iota
    // Devices with DMA remapping incompatible drivers will never be allowed to start and perform DMA at any time.
    BLOCKALL_DMAGUARDDEVICEENUMERATIONPOLICYTYPE
    // All external DMA capable PCIe devices will be enumerated at any time.
    ALLOWALL_DMAGUARDDEVICEENUMERATIONPOLICYTYPE
)

func (i DmaGuardDeviceEnumerationPolicyType) String() string {
    return []string{"deviceDefault", "blockAll", "allowAll"}[i]
}
func ParseDmaGuardDeviceEnumerationPolicyType(v string) (interface{}, error) {
    result := DEVICEDEFAULT_DMAGUARDDEVICEENUMERATIONPOLICYTYPE
    switch v {
        case "deviceDefault":
            result = DEVICEDEFAULT_DMAGUARDDEVICEENUMERATIONPOLICYTYPE
        case "blockAll":
            result = BLOCKALL_DMAGUARDDEVICEENUMERATIONPOLICYTYPE
        case "allowAll":
            result = ALLOWALL_DMAGUARDDEVICEENUMERATIONPOLICYTYPE
        default:
            return 0, errors.New("Unknown DmaGuardDeviceEnumerationPolicyType value: " + v)
    }
    return &result, nil
}
func SerializeDmaGuardDeviceEnumerationPolicyType(values []DmaGuardDeviceEnumerationPolicyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
