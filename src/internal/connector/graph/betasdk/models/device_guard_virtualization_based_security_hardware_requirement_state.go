package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type DeviceGuardVirtualizationBasedSecurityHardwareRequirementState int

const (
    // System meets hardware configuration requirements
    MEETHARDWAREREQUIREMENTS_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYHARDWAREREQUIREMENTSTATE DeviceGuardVirtualizationBasedSecurityHardwareRequirementState = iota
    // Secure boot required
    SECUREBOOTREQUIRED_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYHARDWAREREQUIREMENTSTATE
    // DMA protection required
    DMAPROTECTIONREQUIRED_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYHARDWAREREQUIREMENTSTATE
    // HyperV not supported for Guest VM
    HYPERVNOTSUPPORTEDFORGUESTVM_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYHARDWAREREQUIREMENTSTATE
    // HyperV feature is not available
    HYPERVNOTAVAILABLE_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYHARDWAREREQUIREMENTSTATE
)

func (i DeviceGuardVirtualizationBasedSecurityHardwareRequirementState) String() string {
    return []string{"meetHardwareRequirements", "secureBootRequired", "dmaProtectionRequired", "hyperVNotSupportedForGuestVM", "hyperVNotAvailable"}[i]
}
func ParseDeviceGuardVirtualizationBasedSecurityHardwareRequirementState(v string) (interface{}, error) {
    result := MEETHARDWAREREQUIREMENTS_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYHARDWAREREQUIREMENTSTATE
    switch v {
        case "meetHardwareRequirements":
            result = MEETHARDWAREREQUIREMENTS_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYHARDWAREREQUIREMENTSTATE
        case "secureBootRequired":
            result = SECUREBOOTREQUIRED_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYHARDWAREREQUIREMENTSTATE
        case "dmaProtectionRequired":
            result = DMAPROTECTIONREQUIRED_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYHARDWAREREQUIREMENTSTATE
        case "hyperVNotSupportedForGuestVM":
            result = HYPERVNOTSUPPORTEDFORGUESTVM_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYHARDWAREREQUIREMENTSTATE
        case "hyperVNotAvailable":
            result = HYPERVNOTAVAILABLE_DEVICEGUARDVIRTUALIZATIONBASEDSECURITYHARDWAREREQUIREMENTSTATE
        default:
            return 0, errors.New("Unknown DeviceGuardVirtualizationBasedSecurityHardwareRequirementState value: " + v)
    }
    return &result, nil
}
func SerializeDeviceGuardVirtualizationBasedSecurityHardwareRequirementState(values []DeviceGuardVirtualizationBasedSecurityHardwareRequirementState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
