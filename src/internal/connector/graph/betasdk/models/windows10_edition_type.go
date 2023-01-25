package models
import (
    "errors"
)
// Provides operations to call the add method.
type Windows10EditionType int

const (
    // Windows 10 Enterprise
    WINDOWS10ENTERPRISE_WINDOWS10EDITIONTYPE Windows10EditionType = iota
    // Windows 10 EnterpriseN
    WINDOWS10ENTERPRISEN_WINDOWS10EDITIONTYPE
    // Windows 10 Education
    WINDOWS10EDUCATION_WINDOWS10EDITIONTYPE
    // Windows 10 EducationN
    WINDOWS10EDUCATIONN_WINDOWS10EDITIONTYPE
    // Windows 10 Mobile Enterprise
    WINDOWS10MOBILEENTERPRISE_WINDOWS10EDITIONTYPE
    // Windows 10 Holographic Enterprise
    WINDOWS10HOLOGRAPHICENTERPRISE_WINDOWS10EDITIONTYPE
    // Windows 10 Professional
    WINDOWS10PROFESSIONAL_WINDOWS10EDITIONTYPE
    // Windows 10 ProfessionalN
    WINDOWS10PROFESSIONALN_WINDOWS10EDITIONTYPE
    // Windows 10 Professional Education
    WINDOWS10PROFESSIONALEDUCATION_WINDOWS10EDITIONTYPE
    // Windows 10 Professional EducationN
    WINDOWS10PROFESSIONALEDUCATIONN_WINDOWS10EDITIONTYPE
    // Windows 10 Professional for Workstations
    WINDOWS10PROFESSIONALWORKSTATION_WINDOWS10EDITIONTYPE
    // Windows 10 Professional for Workstations N
    WINDOWS10PROFESSIONALWORKSTATIONN_WINDOWS10EDITIONTYPE
    // NotConfigured
    NOTCONFIGURED_WINDOWS10EDITIONTYPE
    // Windows 10 Home
    WINDOWS10HOME_WINDOWS10EDITIONTYPE
    // Windows 10 Home China
    WINDOWS10HOMECHINA_WINDOWS10EDITIONTYPE
    // Windows 10 Home N
    WINDOWS10HOMEN_WINDOWS10EDITIONTYPE
    // Windows 10 Home Single Language
    WINDOWS10HOMESINGLELANGUAGE_WINDOWS10EDITIONTYPE
    // Windows 10 Mobile
    WINDOWS10MOBILE_WINDOWS10EDITIONTYPE
    // Windows 10 IoT Core
    WINDOWS10IOTCORE_WINDOWS10EDITIONTYPE
    // Windows 10 IoT Core Commercial
    WINDOWS10IOTCORECOMMERCIAL_WINDOWS10EDITIONTYPE
)

func (i Windows10EditionType) String() string {
    return []string{"windows10Enterprise", "windows10EnterpriseN", "windows10Education", "windows10EducationN", "windows10MobileEnterprise", "windows10HolographicEnterprise", "windows10Professional", "windows10ProfessionalN", "windows10ProfessionalEducation", "windows10ProfessionalEducationN", "windows10ProfessionalWorkstation", "windows10ProfessionalWorkstationN", "notConfigured", "windows10Home", "windows10HomeChina", "windows10HomeN", "windows10HomeSingleLanguage", "windows10Mobile", "windows10IoTCore", "windows10IoTCoreCommercial"}[i]
}
func ParseWindows10EditionType(v string) (interface{}, error) {
    result := WINDOWS10ENTERPRISE_WINDOWS10EDITIONTYPE
    switch v {
        case "windows10Enterprise":
            result = WINDOWS10ENTERPRISE_WINDOWS10EDITIONTYPE
        case "windows10EnterpriseN":
            result = WINDOWS10ENTERPRISEN_WINDOWS10EDITIONTYPE
        case "windows10Education":
            result = WINDOWS10EDUCATION_WINDOWS10EDITIONTYPE
        case "windows10EducationN":
            result = WINDOWS10EDUCATIONN_WINDOWS10EDITIONTYPE
        case "windows10MobileEnterprise":
            result = WINDOWS10MOBILEENTERPRISE_WINDOWS10EDITIONTYPE
        case "windows10HolographicEnterprise":
            result = WINDOWS10HOLOGRAPHICENTERPRISE_WINDOWS10EDITIONTYPE
        case "windows10Professional":
            result = WINDOWS10PROFESSIONAL_WINDOWS10EDITIONTYPE
        case "windows10ProfessionalN":
            result = WINDOWS10PROFESSIONALN_WINDOWS10EDITIONTYPE
        case "windows10ProfessionalEducation":
            result = WINDOWS10PROFESSIONALEDUCATION_WINDOWS10EDITIONTYPE
        case "windows10ProfessionalEducationN":
            result = WINDOWS10PROFESSIONALEDUCATIONN_WINDOWS10EDITIONTYPE
        case "windows10ProfessionalWorkstation":
            result = WINDOWS10PROFESSIONALWORKSTATION_WINDOWS10EDITIONTYPE
        case "windows10ProfessionalWorkstationN":
            result = WINDOWS10PROFESSIONALWORKSTATIONN_WINDOWS10EDITIONTYPE
        case "notConfigured":
            result = NOTCONFIGURED_WINDOWS10EDITIONTYPE
        case "windows10Home":
            result = WINDOWS10HOME_WINDOWS10EDITIONTYPE
        case "windows10HomeChina":
            result = WINDOWS10HOMECHINA_WINDOWS10EDITIONTYPE
        case "windows10HomeN":
            result = WINDOWS10HOMEN_WINDOWS10EDITIONTYPE
        case "windows10HomeSingleLanguage":
            result = WINDOWS10HOMESINGLELANGUAGE_WINDOWS10EDITIONTYPE
        case "windows10Mobile":
            result = WINDOWS10MOBILE_WINDOWS10EDITIONTYPE
        case "windows10IoTCore":
            result = WINDOWS10IOTCORE_WINDOWS10EDITIONTYPE
        case "windows10IoTCoreCommercial":
            result = WINDOWS10IOTCORECOMMERCIAL_WINDOWS10EDITIONTYPE
        default:
            return 0, errors.New("Unknown Windows10EditionType value: " + v)
    }
    return &result, nil
}
func SerializeWindows10EditionType(values []Windows10EditionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
