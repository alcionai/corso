package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type CertificateIssuanceStates int

const (
    UNKNOWN_CERTIFICATEISSUANCESTATES CertificateIssuanceStates = iota
    CHALLENGEISSUED_CERTIFICATEISSUANCESTATES
    CHALLENGEISSUEFAILED_CERTIFICATEISSUANCESTATES
    REQUESTCREATIONFAILED_CERTIFICATEISSUANCESTATES
    REQUESTSUBMITFAILED_CERTIFICATEISSUANCESTATES
    CHALLENGEVALIDATIONSUCCEEDED_CERTIFICATEISSUANCESTATES
    CHALLENGEVALIDATIONFAILED_CERTIFICATEISSUANCESTATES
    ISSUEFAILED_CERTIFICATEISSUANCESTATES
    ISSUEPENDING_CERTIFICATEISSUANCESTATES
    ISSUED_CERTIFICATEISSUANCESTATES
    RESPONSEPROCESSINGFAILED_CERTIFICATEISSUANCESTATES
    RESPONSEPENDING_CERTIFICATEISSUANCESTATES
    ENROLLMENTSUCCEEDED_CERTIFICATEISSUANCESTATES
    ENROLLMENTNOTNEEDED_CERTIFICATEISSUANCESTATES
    REVOKED_CERTIFICATEISSUANCESTATES
    REMOVEDFROMCOLLECTION_CERTIFICATEISSUANCESTATES
    RENEWVERIFIED_CERTIFICATEISSUANCESTATES
    INSTALLFAILED_CERTIFICATEISSUANCESTATES
    INSTALLED_CERTIFICATEISSUANCESTATES
    DELETEFAILED_CERTIFICATEISSUANCESTATES
    DELETED_CERTIFICATEISSUANCESTATES
    RENEWALREQUESTED_CERTIFICATEISSUANCESTATES
    REQUESTED_CERTIFICATEISSUANCESTATES
)

func (i CertificateIssuanceStates) String() string {
    return []string{"unknown", "challengeIssued", "challengeIssueFailed", "requestCreationFailed", "requestSubmitFailed", "challengeValidationSucceeded", "challengeValidationFailed", "issueFailed", "issuePending", "issued", "responseProcessingFailed", "responsePending", "enrollmentSucceeded", "enrollmentNotNeeded", "revoked", "removedFromCollection", "renewVerified", "installFailed", "installed", "deleteFailed", "deleted", "renewalRequested", "requested"}[i]
}
func ParseCertificateIssuanceStates(v string) (interface{}, error) {
    result := UNKNOWN_CERTIFICATEISSUANCESTATES
    switch v {
        case "unknown":
            result = UNKNOWN_CERTIFICATEISSUANCESTATES
        case "challengeIssued":
            result = CHALLENGEISSUED_CERTIFICATEISSUANCESTATES
        case "challengeIssueFailed":
            result = CHALLENGEISSUEFAILED_CERTIFICATEISSUANCESTATES
        case "requestCreationFailed":
            result = REQUESTCREATIONFAILED_CERTIFICATEISSUANCESTATES
        case "requestSubmitFailed":
            result = REQUESTSUBMITFAILED_CERTIFICATEISSUANCESTATES
        case "challengeValidationSucceeded":
            result = CHALLENGEVALIDATIONSUCCEEDED_CERTIFICATEISSUANCESTATES
        case "challengeValidationFailed":
            result = CHALLENGEVALIDATIONFAILED_CERTIFICATEISSUANCESTATES
        case "issueFailed":
            result = ISSUEFAILED_CERTIFICATEISSUANCESTATES
        case "issuePending":
            result = ISSUEPENDING_CERTIFICATEISSUANCESTATES
        case "issued":
            result = ISSUED_CERTIFICATEISSUANCESTATES
        case "responseProcessingFailed":
            result = RESPONSEPROCESSINGFAILED_CERTIFICATEISSUANCESTATES
        case "responsePending":
            result = RESPONSEPENDING_CERTIFICATEISSUANCESTATES
        case "enrollmentSucceeded":
            result = ENROLLMENTSUCCEEDED_CERTIFICATEISSUANCESTATES
        case "enrollmentNotNeeded":
            result = ENROLLMENTNOTNEEDED_CERTIFICATEISSUANCESTATES
        case "revoked":
            result = REVOKED_CERTIFICATEISSUANCESTATES
        case "removedFromCollection":
            result = REMOVEDFROMCOLLECTION_CERTIFICATEISSUANCESTATES
        case "renewVerified":
            result = RENEWVERIFIED_CERTIFICATEISSUANCESTATES
        case "installFailed":
            result = INSTALLFAILED_CERTIFICATEISSUANCESTATES
        case "installed":
            result = INSTALLED_CERTIFICATEISSUANCESTATES
        case "deleteFailed":
            result = DELETEFAILED_CERTIFICATEISSUANCESTATES
        case "deleted":
            result = DELETED_CERTIFICATEISSUANCESTATES
        case "renewalRequested":
            result = RENEWALREQUESTED_CERTIFICATEISSUANCESTATES
        case "requested":
            result = REQUESTED_CERTIFICATEISSUANCESTATES
        default:
            return 0, errors.New("Unknown CertificateIssuanceStates value: " + v)
    }
    return &result, nil
}
func SerializeCertificateIssuanceStates(values []CertificateIssuanceStates) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
