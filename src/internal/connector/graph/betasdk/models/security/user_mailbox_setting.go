package security
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type UserMailboxSetting int

const (
    NONE_USERMAILBOXSETTING UserMailboxSetting = iota
    JUNKMAILDELETION_USERMAILBOXSETTING
    ISFROMADDRESSINADDRESSBOOK_USERMAILBOXSETTING
    ISFROMADDRESSINADDRESSSAFELIST_USERMAILBOXSETTING
    ISFROMADDRESSINADDRESSBLOCKLIST_USERMAILBOXSETTING
    ISFROMADDRESSINADDRESSIMPLICITSAFELIST_USERMAILBOXSETTING
    ISFROMADDRESSINADDRESSIMPLICITJUNKLIST_USERMAILBOXSETTING
    ISFROMDOMAININDOMAINSAFELIST_USERMAILBOXSETTING
    ISFROMDOMAININDOMAINBLOCKLIST_USERMAILBOXSETTING
    ISRECIPIENTINRECIPIENTSAFELIST_USERMAILBOXSETTING
    CUSTOMRULE_USERMAILBOXSETTING
    JUNKMAILRULE_USERMAILBOXSETTING
    SENDERPRAPRESENT_USERMAILBOXSETTING
    FROMFIRSTTIMESENDER_USERMAILBOXSETTING
    EXCLUSIVE_USERMAILBOXSETTING
    PRIORSEENPASS_USERMAILBOXSETTING
    SENDERAUTHENTICATIONSUCCEEDED_USERMAILBOXSETTING
    ISJUNKMAILRULEENABLED_USERMAILBOXSETTING
    UNKNOWNFUTUREVALUE_USERMAILBOXSETTING
)

func (i UserMailboxSetting) String() string {
    return []string{"none", "junkMailDeletion", "isFromAddressInAddressBook", "isFromAddressInAddressSafeList", "isFromAddressInAddressBlockList", "isFromAddressInAddressImplicitSafeList", "isFromAddressInAddressImplicitJunkList", "isFromDomainInDomainSafeList", "isFromDomainInDomainBlockList", "isRecipientInRecipientSafeList", "customRule", "junkMailRule", "senderPraPresent", "fromFirstTimeSender", "exclusive", "priorSeenPass", "senderAuthenticationSucceeded", "isJunkMailRuleEnabled", "unknownFutureValue"}[i]
}
func ParseUserMailboxSetting(v string) (interface{}, error) {
    result := NONE_USERMAILBOXSETTING
    switch v {
        case "none":
            result = NONE_USERMAILBOXSETTING
        case "junkMailDeletion":
            result = JUNKMAILDELETION_USERMAILBOXSETTING
        case "isFromAddressInAddressBook":
            result = ISFROMADDRESSINADDRESSBOOK_USERMAILBOXSETTING
        case "isFromAddressInAddressSafeList":
            result = ISFROMADDRESSINADDRESSSAFELIST_USERMAILBOXSETTING
        case "isFromAddressInAddressBlockList":
            result = ISFROMADDRESSINADDRESSBLOCKLIST_USERMAILBOXSETTING
        case "isFromAddressInAddressImplicitSafeList":
            result = ISFROMADDRESSINADDRESSIMPLICITSAFELIST_USERMAILBOXSETTING
        case "isFromAddressInAddressImplicitJunkList":
            result = ISFROMADDRESSINADDRESSIMPLICITJUNKLIST_USERMAILBOXSETTING
        case "isFromDomainInDomainSafeList":
            result = ISFROMDOMAININDOMAINSAFELIST_USERMAILBOXSETTING
        case "isFromDomainInDomainBlockList":
            result = ISFROMDOMAININDOMAINBLOCKLIST_USERMAILBOXSETTING
        case "isRecipientInRecipientSafeList":
            result = ISRECIPIENTINRECIPIENTSAFELIST_USERMAILBOXSETTING
        case "customRule":
            result = CUSTOMRULE_USERMAILBOXSETTING
        case "junkMailRule":
            result = JUNKMAILRULE_USERMAILBOXSETTING
        case "senderPraPresent":
            result = SENDERPRAPRESENT_USERMAILBOXSETTING
        case "fromFirstTimeSender":
            result = FROMFIRSTTIMESENDER_USERMAILBOXSETTING
        case "exclusive":
            result = EXCLUSIVE_USERMAILBOXSETTING
        case "priorSeenPass":
            result = PRIORSEENPASS_USERMAILBOXSETTING
        case "senderAuthenticationSucceeded":
            result = SENDERAUTHENTICATIONSUCCEEDED_USERMAILBOXSETTING
        case "isJunkMailRuleEnabled":
            result = ISJUNKMAILRULEENABLED_USERMAILBOXSETTING
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_USERMAILBOXSETTING
        default:
            return 0, errors.New("Unknown UserMailboxSetting value: " + v)
    }
    return &result, nil
}
func SerializeUserMailboxSetting(values []UserMailboxSetting) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
