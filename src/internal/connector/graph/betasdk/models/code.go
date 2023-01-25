package models
import (
    "errors"
)
// Provides operations to call the add method.
type Code int

const (
    // None error.
    NONE_CODE Code = iota
    // Json file invalid error.
    JSONFILEINVALID_CODE
    // Json file missing error.
    JSONFILEMISSING_CODE
    // Json file too large error.
    JSONFILETOOLARGE_CODE
    // Rules missing error.
    RULESMISSING_CODE
    // Duplicate rules error.
    DUPLICATERULES_CODE
    // Too many rules specified error.
    TOOMANYRULESSPECIFIED_CODE
    // Operator missing error.
    OPERATORMISSING_CODE
    // Operator not supported error.
    OPERATORNOTSUPPORTED_CODE
    // Data type missing error.
    DATATYPEMISSING_CODE
    // Data type not supported error.
    DATATYPENOTSUPPORTED_CODE
    // Operator data type combination not supported error.
    OPERATORDATATYPECOMBINATIONNOTSUPPORTED_CODE
    // More info urlmissing error.
    MOREINFOURIMISSING_CODE
    // More info url invalid error.
    MOREINFOURIINVALID_CODE
    // More info ur ltoo large error.
    MOREINFOURITOOLARGE_CODE
    // Description missing error.
    DESCRIPTIONMISSING_CODE
    // Description invalid error.
    DESCRIPTIONINVALID_CODE
    // Description too large error.
    DESCRIPTIONTOOLARGE_CODE
    // Title missing error.
    TITLEMISSING_CODE
    // Title invalid error.
    TITLEINVALID_CODE
    // Title too large error.
    TITLETOOLARGE_CODE
    // Operand missing error.
    OPERANDMISSING_CODE
    // Operand invalid error.
    OPERANDINVALID_CODE
    // Operand too large error.
    OPERANDTOOLARGE_CODE
    // Setting name missing error.
    SETTINGNAMEMISSING_CODE
    // Setting name invalid error.
    SETTINGNAMEINVALID_CODE
    // Setting name too large error.
    SETTINGNAMETOOLARGE_CODE
    // English locale missing error.
    ENGLISHLOCALEMISSING_CODE
    // Duplicate locales error.
    DUPLICATELOCALES_CODE
    // Unrecognized locale error.
    UNRECOGNIZEDLOCALE_CODE
    // Unknown error.
    UNKNOWN_CODE
    // Remediation strings missing error.
    REMEDIATIONSTRINGSMISSING_CODE
)

func (i Code) String() string {
    return []string{"none", "jsonFileInvalid", "jsonFileMissing", "jsonFileTooLarge", "rulesMissing", "duplicateRules", "tooManyRulesSpecified", "operatorMissing", "operatorNotSupported", "datatypeMissing", "datatypeNotSupported", "operatorDataTypeCombinationNotSupported", "moreInfoUriMissing", "moreInfoUriInvalid", "moreInfoUriTooLarge", "descriptionMissing", "descriptionInvalid", "descriptionTooLarge", "titleMissing", "titleInvalid", "titleTooLarge", "operandMissing", "operandInvalid", "operandTooLarge", "settingNameMissing", "settingNameInvalid", "settingNameTooLarge", "englishLocaleMissing", "duplicateLocales", "unrecognizedLocale", "unknown", "remediationStringsMissing"}[i]
}
func ParseCode(v string) (interface{}, error) {
    result := NONE_CODE
    switch v {
        case "none":
            result = NONE_CODE
        case "jsonFileInvalid":
            result = JSONFILEINVALID_CODE
        case "jsonFileMissing":
            result = JSONFILEMISSING_CODE
        case "jsonFileTooLarge":
            result = JSONFILETOOLARGE_CODE
        case "rulesMissing":
            result = RULESMISSING_CODE
        case "duplicateRules":
            result = DUPLICATERULES_CODE
        case "tooManyRulesSpecified":
            result = TOOMANYRULESSPECIFIED_CODE
        case "operatorMissing":
            result = OPERATORMISSING_CODE
        case "operatorNotSupported":
            result = OPERATORNOTSUPPORTED_CODE
        case "datatypeMissing":
            result = DATATYPEMISSING_CODE
        case "datatypeNotSupported":
            result = DATATYPENOTSUPPORTED_CODE
        case "operatorDataTypeCombinationNotSupported":
            result = OPERATORDATATYPECOMBINATIONNOTSUPPORTED_CODE
        case "moreInfoUriMissing":
            result = MOREINFOURIMISSING_CODE
        case "moreInfoUriInvalid":
            result = MOREINFOURIINVALID_CODE
        case "moreInfoUriTooLarge":
            result = MOREINFOURITOOLARGE_CODE
        case "descriptionMissing":
            result = DESCRIPTIONMISSING_CODE
        case "descriptionInvalid":
            result = DESCRIPTIONINVALID_CODE
        case "descriptionTooLarge":
            result = DESCRIPTIONTOOLARGE_CODE
        case "titleMissing":
            result = TITLEMISSING_CODE
        case "titleInvalid":
            result = TITLEINVALID_CODE
        case "titleTooLarge":
            result = TITLETOOLARGE_CODE
        case "operandMissing":
            result = OPERANDMISSING_CODE
        case "operandInvalid":
            result = OPERANDINVALID_CODE
        case "operandTooLarge":
            result = OPERANDTOOLARGE_CODE
        case "settingNameMissing":
            result = SETTINGNAMEMISSING_CODE
        case "settingNameInvalid":
            result = SETTINGNAMEINVALID_CODE
        case "settingNameTooLarge":
            result = SETTINGNAMETOOLARGE_CODE
        case "englishLocaleMissing":
            result = ENGLISHLOCALEMISSING_CODE
        case "duplicateLocales":
            result = DUPLICATELOCALES_CODE
        case "unrecognizedLocale":
            result = UNRECOGNIZEDLOCALE_CODE
        case "unknown":
            result = UNKNOWN_CODE
        case "remediationStringsMissing":
            result = REMEDIATIONSTRINGSMISSING_CODE
        default:
            return 0, errors.New("Unknown Code value: " + v)
    }
    return &result, nil
}
func SerializeCode(values []Code) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
