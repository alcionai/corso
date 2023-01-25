package security
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type SubmissionResultDetail int

const (
    NONE_SUBMISSIONRESULTDETAIL SubmissionResultDetail = iota
    UNDERINVESTIGATION_SUBMISSIONRESULTDETAIL
    SIMULATEDTHREAT_SUBMISSIONRESULTDETAIL
    ALLOWEDBYSECOPS_SUBMISSIONRESULTDETAIL
    ALLOWEDBYTHIRDPARTYFILTERS_SUBMISSIONRESULTDETAIL
    MESSAGENOTFOUND_SUBMISSIONRESULTDETAIL
    URLFILESHOULDNOTBEBLOCKED_SUBMISSIONRESULTDETAIL
    URLFILESHOULDBEBLOCKED_SUBMISSIONRESULTDETAIL
    URLFILECANNOTMAKEDECISION_SUBMISSIONRESULTDETAIL
    DOMAINIMPERSONATION_SUBMISSIONRESULTDETAIL
    USERIMPERSONATION_SUBMISSIONRESULTDETAIL
    BRANDIMPERSONATION_SUBMISSIONRESULTDETAIL
    OUTBOUNDSHOULDNOTBEBLOCKED_SUBMISSIONRESULTDETAIL
    OUTBOUNDSHOULDBEBLOCKED_SUBMISSIONRESULTDETAIL
    OUTBOUNDBULK_SUBMISSIONRESULTDETAIL
    OUTBOUNDCANNOTMAKEDECISION_SUBMISSIONRESULTDETAIL
    OUTBOUNDNOTRESCANNED_SUBMISSIONRESULTDETAIL
    ZEROHOURAUTOPURGEALLOWED_SUBMISSIONRESULTDETAIL
    ZEROHOURAUTOPURGEBLOCKED_SUBMISSIONRESULTDETAIL
    ZEROHOURAUTOPURGEQUARANTINERELEASED_SUBMISSIONRESULTDETAIL
    ONPREMISESSKIP_SUBMISSIONRESULTDETAIL
    ALLOWEDBYTENANTALLOWBLOCKLIST_SUBMISSIONRESULTDETAIL
    BLOCKEDBYTENANTALLOWBLOCKLIST_SUBMISSIONRESULTDETAIL
    ALLOWEDURLBYTENANTALLOWBLOCKLIST_SUBMISSIONRESULTDETAIL
    ALLOWEDFILEBYTENANTALLOWBLOCKLIST_SUBMISSIONRESULTDETAIL
    ALLOWEDSENDERBYTENANTALLOWBLOCKLIST_SUBMISSIONRESULTDETAIL
    ALLOWEDRECIPIENTBYTENANTALLOWBLOCKLIST_SUBMISSIONRESULTDETAIL
    BLOCKEDURLBYTENANTALLOWBLOCKLIST_SUBMISSIONRESULTDETAIL
    BLOCKEDFILEBYTENANTALLOWBLOCKLIST_SUBMISSIONRESULTDETAIL
    BLOCKEDSENDERBYTENANTALLOWBLOCKLIST_SUBMISSIONRESULTDETAIL
    BLOCKEDRECIPIENTBYTENANTALLOWBLOCKLIST_SUBMISSIONRESULTDETAIL
    ALLOWEDBYCONNECTION_SUBMISSIONRESULTDETAIL
    BLOCKEDBYCONNECTION_SUBMISSIONRESULTDETAIL
    ALLOWEDBYEXCHANGETRANSPORTRULE_SUBMISSIONRESULTDETAIL
    BLOCKEDBYEXCHANGETRANSPORTRULE_SUBMISSIONRESULTDETAIL
    QUARANTINERELEASED_SUBMISSIONRESULTDETAIL
    QUARANTINERELEASEDTHENBLOCKED_SUBMISSIONRESULTDETAIL
    JUNKMAILRULEDISABLED_SUBMISSIONRESULTDETAIL
    ALLOWEDBYUSERSETTING_SUBMISSIONRESULTDETAIL
    BLOCKEDBYUSERSETTING_SUBMISSIONRESULTDETAIL
    ALLOWEDBYTENANT_SUBMISSIONRESULTDETAIL
    BLOCKEDBYTENANT_SUBMISSIONRESULTDETAIL
    INVALIDFALSEPOSITIVE_SUBMISSIONRESULTDETAIL
    INVALIDFALSENEGATIVE_SUBMISSIONRESULTDETAIL
    SPOOFBLOCKED_SUBMISSIONRESULTDETAIL
    GOODRECLASSIFIEDASBAD_SUBMISSIONRESULTDETAIL
    GOODRECLASSIFIEDASBULK_SUBMISSIONRESULTDETAIL
    GOODRECLASSIFIEDASGOOD_SUBMISSIONRESULTDETAIL
    GOODRECLASSIFIEDASCANNOTMAKEDECISION_SUBMISSIONRESULTDETAIL
    BADRECLASSIFIEDASGOOD_SUBMISSIONRESULTDETAIL
    BADRECLASSIFIEDASBULK_SUBMISSIONRESULTDETAIL
    BADRECLASSIFIEDASBAD_SUBMISSIONRESULTDETAIL
    BADRECLASSIFIEDASCANNOTMAKEDECISION_SUBMISSIONRESULTDETAIL
    UNKNOWNFUTUREVALUE_SUBMISSIONRESULTDETAIL
)

func (i SubmissionResultDetail) String() string {
    return []string{"none", "underInvestigation", "simulatedThreat", "allowedBySecOps", "allowedByThirdPartyFilters", "messageNotFound", "urlFileShouldNotBeBlocked", "urlFileShouldBeBlocked", "urlFileCannotMakeDecision", "domainImpersonation", "userImpersonation", "brandImpersonation", "outboundShouldNotBeBlocked", "outboundShouldBeBlocked", "outboundBulk", "outboundCannotMakeDecision", "outboundNotRescanned", "zeroHourAutoPurgeAllowed", "zeroHourAutoPurgeBlocked", "zeroHourAutoPurgeQuarantineReleased", "onPremisesSkip", "allowedByTenantAllowBlockList", "blockedByTenantAllowBlockList", "allowedUrlByTenantAllowBlockList", "allowedFileByTenantAllowBlockList", "allowedSenderByTenantAllowBlockList", "allowedRecipientByTenantAllowBlockList", "blockedUrlByTenantAllowBlockList", "blockedFileByTenantAllowBlockList", "blockedSenderByTenantAllowBlockList", "blockedRecipientByTenantAllowBlockList", "allowedByConnection", "blockedByConnection", "allowedByExchangeTransportRule", "blockedByExchangeTransportRule", "quarantineReleased", "quarantineReleasedThenBlocked", "junkMailRuleDisabled", "allowedByUserSetting", "blockedByUserSetting", "allowedByTenant", "blockedByTenant", "invalidFalsePositive", "invalidFalseNegative", "spoofBlocked", "goodReclassifiedAsBad", "goodReclassifiedAsBulk", "goodReclassifiedAsGood", "goodReclassifiedAsCannotMakeDecision", "badReclassifiedAsGood", "badReclassifiedAsBulk", "badReclassifiedAsBad", "badReclassifiedAsCannotMakeDecision", "unknownFutureValue"}[i]
}
func ParseSubmissionResultDetail(v string) (interface{}, error) {
    result := NONE_SUBMISSIONRESULTDETAIL
    switch v {
        case "none":
            result = NONE_SUBMISSIONRESULTDETAIL
        case "underInvestigation":
            result = UNDERINVESTIGATION_SUBMISSIONRESULTDETAIL
        case "simulatedThreat":
            result = SIMULATEDTHREAT_SUBMISSIONRESULTDETAIL
        case "allowedBySecOps":
            result = ALLOWEDBYSECOPS_SUBMISSIONRESULTDETAIL
        case "allowedByThirdPartyFilters":
            result = ALLOWEDBYTHIRDPARTYFILTERS_SUBMISSIONRESULTDETAIL
        case "messageNotFound":
            result = MESSAGENOTFOUND_SUBMISSIONRESULTDETAIL
        case "urlFileShouldNotBeBlocked":
            result = URLFILESHOULDNOTBEBLOCKED_SUBMISSIONRESULTDETAIL
        case "urlFileShouldBeBlocked":
            result = URLFILESHOULDBEBLOCKED_SUBMISSIONRESULTDETAIL
        case "urlFileCannotMakeDecision":
            result = URLFILECANNOTMAKEDECISION_SUBMISSIONRESULTDETAIL
        case "domainImpersonation":
            result = DOMAINIMPERSONATION_SUBMISSIONRESULTDETAIL
        case "userImpersonation":
            result = USERIMPERSONATION_SUBMISSIONRESULTDETAIL
        case "brandImpersonation":
            result = BRANDIMPERSONATION_SUBMISSIONRESULTDETAIL
        case "outboundShouldNotBeBlocked":
            result = OUTBOUNDSHOULDNOTBEBLOCKED_SUBMISSIONRESULTDETAIL
        case "outboundShouldBeBlocked":
            result = OUTBOUNDSHOULDBEBLOCKED_SUBMISSIONRESULTDETAIL
        case "outboundBulk":
            result = OUTBOUNDBULK_SUBMISSIONRESULTDETAIL
        case "outboundCannotMakeDecision":
            result = OUTBOUNDCANNOTMAKEDECISION_SUBMISSIONRESULTDETAIL
        case "outboundNotRescanned":
            result = OUTBOUNDNOTRESCANNED_SUBMISSIONRESULTDETAIL
        case "zeroHourAutoPurgeAllowed":
            result = ZEROHOURAUTOPURGEALLOWED_SUBMISSIONRESULTDETAIL
        case "zeroHourAutoPurgeBlocked":
            result = ZEROHOURAUTOPURGEBLOCKED_SUBMISSIONRESULTDETAIL
        case "zeroHourAutoPurgeQuarantineReleased":
            result = ZEROHOURAUTOPURGEQUARANTINERELEASED_SUBMISSIONRESULTDETAIL
        case "onPremisesSkip":
            result = ONPREMISESSKIP_SUBMISSIONRESULTDETAIL
        case "allowedByTenantAllowBlockList":
            result = ALLOWEDBYTENANTALLOWBLOCKLIST_SUBMISSIONRESULTDETAIL
        case "blockedByTenantAllowBlockList":
            result = BLOCKEDBYTENANTALLOWBLOCKLIST_SUBMISSIONRESULTDETAIL
        case "allowedUrlByTenantAllowBlockList":
            result = ALLOWEDURLBYTENANTALLOWBLOCKLIST_SUBMISSIONRESULTDETAIL
        case "allowedFileByTenantAllowBlockList":
            result = ALLOWEDFILEBYTENANTALLOWBLOCKLIST_SUBMISSIONRESULTDETAIL
        case "allowedSenderByTenantAllowBlockList":
            result = ALLOWEDSENDERBYTENANTALLOWBLOCKLIST_SUBMISSIONRESULTDETAIL
        case "allowedRecipientByTenantAllowBlockList":
            result = ALLOWEDRECIPIENTBYTENANTALLOWBLOCKLIST_SUBMISSIONRESULTDETAIL
        case "blockedUrlByTenantAllowBlockList":
            result = BLOCKEDURLBYTENANTALLOWBLOCKLIST_SUBMISSIONRESULTDETAIL
        case "blockedFileByTenantAllowBlockList":
            result = BLOCKEDFILEBYTENANTALLOWBLOCKLIST_SUBMISSIONRESULTDETAIL
        case "blockedSenderByTenantAllowBlockList":
            result = BLOCKEDSENDERBYTENANTALLOWBLOCKLIST_SUBMISSIONRESULTDETAIL
        case "blockedRecipientByTenantAllowBlockList":
            result = BLOCKEDRECIPIENTBYTENANTALLOWBLOCKLIST_SUBMISSIONRESULTDETAIL
        case "allowedByConnection":
            result = ALLOWEDBYCONNECTION_SUBMISSIONRESULTDETAIL
        case "blockedByConnection":
            result = BLOCKEDBYCONNECTION_SUBMISSIONRESULTDETAIL
        case "allowedByExchangeTransportRule":
            result = ALLOWEDBYEXCHANGETRANSPORTRULE_SUBMISSIONRESULTDETAIL
        case "blockedByExchangeTransportRule":
            result = BLOCKEDBYEXCHANGETRANSPORTRULE_SUBMISSIONRESULTDETAIL
        case "quarantineReleased":
            result = QUARANTINERELEASED_SUBMISSIONRESULTDETAIL
        case "quarantineReleasedThenBlocked":
            result = QUARANTINERELEASEDTHENBLOCKED_SUBMISSIONRESULTDETAIL
        case "junkMailRuleDisabled":
            result = JUNKMAILRULEDISABLED_SUBMISSIONRESULTDETAIL
        case "allowedByUserSetting":
            result = ALLOWEDBYUSERSETTING_SUBMISSIONRESULTDETAIL
        case "blockedByUserSetting":
            result = BLOCKEDBYUSERSETTING_SUBMISSIONRESULTDETAIL
        case "allowedByTenant":
            result = ALLOWEDBYTENANT_SUBMISSIONRESULTDETAIL
        case "blockedByTenant":
            result = BLOCKEDBYTENANT_SUBMISSIONRESULTDETAIL
        case "invalidFalsePositive":
            result = INVALIDFALSEPOSITIVE_SUBMISSIONRESULTDETAIL
        case "invalidFalseNegative":
            result = INVALIDFALSENEGATIVE_SUBMISSIONRESULTDETAIL
        case "spoofBlocked":
            result = SPOOFBLOCKED_SUBMISSIONRESULTDETAIL
        case "goodReclassifiedAsBad":
            result = GOODRECLASSIFIEDASBAD_SUBMISSIONRESULTDETAIL
        case "goodReclassifiedAsBulk":
            result = GOODRECLASSIFIEDASBULK_SUBMISSIONRESULTDETAIL
        case "goodReclassifiedAsGood":
            result = GOODRECLASSIFIEDASGOOD_SUBMISSIONRESULTDETAIL
        case "goodReclassifiedAsCannotMakeDecision":
            result = GOODRECLASSIFIEDASCANNOTMAKEDECISION_SUBMISSIONRESULTDETAIL
        case "badReclassifiedAsGood":
            result = BADRECLASSIFIEDASGOOD_SUBMISSIONRESULTDETAIL
        case "badReclassifiedAsBulk":
            result = BADRECLASSIFIEDASBULK_SUBMISSIONRESULTDETAIL
        case "badReclassifiedAsBad":
            result = BADRECLASSIFIEDASBAD_SUBMISSIONRESULTDETAIL
        case "badReclassifiedAsCannotMakeDecision":
            result = BADRECLASSIFIEDASCANNOTMAKEDECISION_SUBMISSIONRESULTDETAIL
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SUBMISSIONRESULTDETAIL
        default:
            return 0, errors.New("Unknown SubmissionResultDetail value: " + v)
    }
    return &result, nil
}
func SerializeSubmissionResultDetail(values []SubmissionResultDetail) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
