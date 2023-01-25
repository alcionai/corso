package security
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ActionAfterRetentionPeriod int

const (
    NONE_ACTIONAFTERRETENTIONPERIOD ActionAfterRetentionPeriod = iota
    DELETE_ACTIONAFTERRETENTIONPERIOD
    STARTDISPOSITIONREVIEW_ACTIONAFTERRETENTIONPERIOD
    RELABEL_ACTIONAFTERRETENTIONPERIOD
    UNKNOWNFUTUREVALUE_ACTIONAFTERRETENTIONPERIOD
)

func (i ActionAfterRetentionPeriod) String() string {
    return []string{"none", "delete", "startDispositionReview", "relabel", "unknownFutureValue"}[i]
}
func ParseActionAfterRetentionPeriod(v string) (interface{}, error) {
    result := NONE_ACTIONAFTERRETENTIONPERIOD
    switch v {
        case "none":
            result = NONE_ACTIONAFTERRETENTIONPERIOD
        case "delete":
            result = DELETE_ACTIONAFTERRETENTIONPERIOD
        case "startDispositionReview":
            result = STARTDISPOSITIONREVIEW_ACTIONAFTERRETENTIONPERIOD
        case "relabel":
            result = RELABEL_ACTIONAFTERRETENTIONPERIOD
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ACTIONAFTERRETENTIONPERIOD
        default:
            return 0, errors.New("Unknown ActionAfterRetentionPeriod value: " + v)
    }
    return &result, nil
}
func SerializeActionAfterRetentionPeriod(values []ActionAfterRetentionPeriod) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
