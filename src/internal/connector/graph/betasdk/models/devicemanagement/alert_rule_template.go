package devicemanagement
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type AlertRuleTemplate int

const (
    CLOUDPCPROVISIONSCENARIO_ALERTRULETEMPLATE AlertRuleTemplate = iota
    CLOUDPCIMAGEUPLOADSCENARIO_ALERTRULETEMPLATE
    CLOUDPCONPREMISENETWORKCONNECTIONCHECKSCENARIO_ALERTRULETEMPLATE
    UNKNOWNFUTUREVALUE_ALERTRULETEMPLATE
)

func (i AlertRuleTemplate) String() string {
    return []string{"cloudPcProvisionScenario", "cloudPcImageUploadScenario", "cloudPcOnPremiseNetworkConnectionCheckScenario", "unknownFutureValue"}[i]
}
func ParseAlertRuleTemplate(v string) (interface{}, error) {
    result := CLOUDPCPROVISIONSCENARIO_ALERTRULETEMPLATE
    switch v {
        case "cloudPcProvisionScenario":
            result = CLOUDPCPROVISIONSCENARIO_ALERTRULETEMPLATE
        case "cloudPcImageUploadScenario":
            result = CLOUDPCIMAGEUPLOADSCENARIO_ALERTRULETEMPLATE
        case "cloudPcOnPremiseNetworkConnectionCheckScenario":
            result = CLOUDPCONPREMISENETWORKCONNECTIONCHECKSCENARIO_ALERTRULETEMPLATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ALERTRULETEMPLATE
        default:
            return 0, errors.New("Unknown AlertRuleTemplate value: " + v)
    }
    return &result, nil
}
func SerializeAlertRuleTemplate(values []AlertRuleTemplate) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
