package models
import (
    "errors"
)
// Provides operations to call the add method.
type DeviceComplianceScriptRulesValidationError int

const (
    // None error.
    NONE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR DeviceComplianceScriptRulesValidationError = iota
    // Json file invalid error.
    JSONFILEINVALID_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // Json file missing error.
    JSONFILEMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // Json file too large error.
    JSONFILETOOLARGE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // Rules missing error.
    RULESMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // Duplicate rules error.
    DUPLICATERULES_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // Too many rules specified error.
    TOOMANYRULESSPECIFIED_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // Operator missing error.
    OPERATORMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // Operator not supported error.
    OPERATORNOTSUPPORTED_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // Data type missing error.
    DATATYPEMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // Data type not supported error.
    DATATYPENOTSUPPORTED_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // Operator data type combination not supported error.
    OPERATORDATATYPECOMBINATIONNOTSUPPORTED_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // More info urlmissing error.
    MOREINFOURIMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // More info url invalid error.
    MOREINFOURIINVALID_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // More info ur ltoo large error.
    MOREINFOURITOOLARGE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // Description missing error.
    DESCRIPTIONMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // Description invalid error.
    DESCRIPTIONINVALID_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // Description too large error.
    DESCRIPTIONTOOLARGE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // Title missing error.
    TITLEMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // Title invalid error.
    TITLEINVALID_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // Title too large error.
    TITLETOOLARGE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // Operand missing error.
    OPERANDMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // Operand invalid error.
    OPERANDINVALID_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // Operand too large error.
    OPERANDTOOLARGE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // Setting name missing error.
    SETTINGNAMEMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // Setting name invalid error.
    SETTINGNAMEINVALID_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // Setting name too large error.
    SETTINGNAMETOOLARGE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // English locale missing error.
    ENGLISHLOCALEMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // Duplicate locales error.
    DUPLICATELOCALES_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // Unrecognized locale error.
    UNRECOGNIZEDLOCALE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // Unknown error.
    UNKNOWN_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    // Remediation strings missing error.
    REMEDIATIONSTRINGSMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
)

func (i DeviceComplianceScriptRulesValidationError) String() string {
    return []string{"none", "jsonFileInvalid", "jsonFileMissing", "jsonFileTooLarge", "rulesMissing", "duplicateRules", "tooManyRulesSpecified", "operatorMissing", "operatorNotSupported", "datatypeMissing", "datatypeNotSupported", "operatorDataTypeCombinationNotSupported", "moreInfoUriMissing", "moreInfoUriInvalid", "moreInfoUriTooLarge", "descriptionMissing", "descriptionInvalid", "descriptionTooLarge", "titleMissing", "titleInvalid", "titleTooLarge", "operandMissing", "operandInvalid", "operandTooLarge", "settingNameMissing", "settingNameInvalid", "settingNameTooLarge", "englishLocaleMissing", "duplicateLocales", "unrecognizedLocale", "unknown", "remediationStringsMissing"}[i]
}
func ParseDeviceComplianceScriptRulesValidationError(v string) (interface{}, error) {
    result := NONE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
    switch v {
        case "none":
            result = NONE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "jsonFileInvalid":
            result = JSONFILEINVALID_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "jsonFileMissing":
            result = JSONFILEMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "jsonFileTooLarge":
            result = JSONFILETOOLARGE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "rulesMissing":
            result = RULESMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "duplicateRules":
            result = DUPLICATERULES_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "tooManyRulesSpecified":
            result = TOOMANYRULESSPECIFIED_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "operatorMissing":
            result = OPERATORMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "operatorNotSupported":
            result = OPERATORNOTSUPPORTED_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "datatypeMissing":
            result = DATATYPEMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "datatypeNotSupported":
            result = DATATYPENOTSUPPORTED_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "operatorDataTypeCombinationNotSupported":
            result = OPERATORDATATYPECOMBINATIONNOTSUPPORTED_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "moreInfoUriMissing":
            result = MOREINFOURIMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "moreInfoUriInvalid":
            result = MOREINFOURIINVALID_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "moreInfoUriTooLarge":
            result = MOREINFOURITOOLARGE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "descriptionMissing":
            result = DESCRIPTIONMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "descriptionInvalid":
            result = DESCRIPTIONINVALID_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "descriptionTooLarge":
            result = DESCRIPTIONTOOLARGE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "titleMissing":
            result = TITLEMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "titleInvalid":
            result = TITLEINVALID_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "titleTooLarge":
            result = TITLETOOLARGE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "operandMissing":
            result = OPERANDMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "operandInvalid":
            result = OPERANDINVALID_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "operandTooLarge":
            result = OPERANDTOOLARGE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "settingNameMissing":
            result = SETTINGNAMEMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "settingNameInvalid":
            result = SETTINGNAMEINVALID_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "settingNameTooLarge":
            result = SETTINGNAMETOOLARGE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "englishLocaleMissing":
            result = ENGLISHLOCALEMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "duplicateLocales":
            result = DUPLICATELOCALES_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "unrecognizedLocale":
            result = UNRECOGNIZEDLOCALE_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "unknown":
            result = UNKNOWN_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        case "remediationStringsMissing":
            result = REMEDIATIONSTRINGSMISSING_DEVICECOMPLIANCESCRIPTRULESVALIDATIONERROR
        default:
            return 0, errors.New("Unknown DeviceComplianceScriptRulesValidationError value: " + v)
    }
    return &result, nil
}
func SerializeDeviceComplianceScriptRulesValidationError(values []DeviceComplianceScriptRulesValidationError) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
