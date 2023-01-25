package models
import (
    "errors"
)
// Provides operations to call the add method.
type WindowsAutopilotEnrollmentType int

const (
    UNKNOWN_WINDOWSAUTOPILOTENROLLMENTTYPE WindowsAutopilotEnrollmentType = iota
    AZUREADJOINEDWITHAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE
    OFFLINEDOMAINJOINED_WINDOWSAUTOPILOTENROLLMENTTYPE
    AZUREADJOINEDUSINGDEVICEAUTHWITHAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE
    AZUREADJOINEDUSINGDEVICEAUTHWITHOUTAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE
    AZUREADJOINEDWITHOFFLINEAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE
    AZUREADJOINEDWITHWHITEGLOVE_WINDOWSAUTOPILOTENROLLMENTTYPE
    OFFLINEDOMAINJOINEDWITHWHITEGLOVE_WINDOWSAUTOPILOTENROLLMENTTYPE
    OFFLINEDOMAINJOINEDWITHOFFLINEAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE
)

func (i WindowsAutopilotEnrollmentType) String() string {
    return []string{"unknown", "azureADJoinedWithAutopilotProfile", "offlineDomainJoined", "azureADJoinedUsingDeviceAuthWithAutopilotProfile", "azureADJoinedUsingDeviceAuthWithoutAutopilotProfile", "azureADJoinedWithOfflineAutopilotProfile", "azureADJoinedWithWhiteGlove", "offlineDomainJoinedWithWhiteGlove", "offlineDomainJoinedWithOfflineAutopilotProfile"}[i]
}
func ParseWindowsAutopilotEnrollmentType(v string) (interface{}, error) {
    result := UNKNOWN_WINDOWSAUTOPILOTENROLLMENTTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_WINDOWSAUTOPILOTENROLLMENTTYPE
        case "azureADJoinedWithAutopilotProfile":
            result = AZUREADJOINEDWITHAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE
        case "offlineDomainJoined":
            result = OFFLINEDOMAINJOINED_WINDOWSAUTOPILOTENROLLMENTTYPE
        case "azureADJoinedUsingDeviceAuthWithAutopilotProfile":
            result = AZUREADJOINEDUSINGDEVICEAUTHWITHAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE
        case "azureADJoinedUsingDeviceAuthWithoutAutopilotProfile":
            result = AZUREADJOINEDUSINGDEVICEAUTHWITHOUTAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE
        case "azureADJoinedWithOfflineAutopilotProfile":
            result = AZUREADJOINEDWITHOFFLINEAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE
        case "azureADJoinedWithWhiteGlove":
            result = AZUREADJOINEDWITHWHITEGLOVE_WINDOWSAUTOPILOTENROLLMENTTYPE
        case "offlineDomainJoinedWithWhiteGlove":
            result = OFFLINEDOMAINJOINEDWITHWHITEGLOVE_WINDOWSAUTOPILOTENROLLMENTTYPE
        case "offlineDomainJoinedWithOfflineAutopilotProfile":
            result = OFFLINEDOMAINJOINEDWITHOFFLINEAUTOPILOTPROFILE_WINDOWSAUTOPILOTENROLLMENTTYPE
        default:
            return 0, errors.New("Unknown WindowsAutopilotEnrollmentType value: " + v)
    }
    return &result, nil
}
func SerializeWindowsAutopilotEnrollmentType(values []WindowsAutopilotEnrollmentType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
