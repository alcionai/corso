package devicemanagement
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type AggregationType int

const (
    COUNT_AGGREGATIONTYPE AggregationType = iota
    PERCENTAGE_AGGREGATIONTYPE
    AFFECTEDCLOUDPCCOUNT_AGGREGATIONTYPE
    AFFECTEDCLOUDPCPERCENTAGE_AGGREGATIONTYPE
    UNKNOWNFUTUREVALUE_AGGREGATIONTYPE
)

func (i AggregationType) String() string {
    return []string{"count", "percentage", "affectedCloudPcCount", "affectedCloudPcPercentage", "unknownFutureValue"}[i]
}
func ParseAggregationType(v string) (interface{}, error) {
    result := COUNT_AGGREGATIONTYPE
    switch v {
        case "count":
            result = COUNT_AGGREGATIONTYPE
        case "percentage":
            result = PERCENTAGE_AGGREGATIONTYPE
        case "affectedCloudPcCount":
            result = AFFECTEDCLOUDPCCOUNT_AGGREGATIONTYPE
        case "affectedCloudPcPercentage":
            result = AFFECTEDCLOUDPCPERCENTAGE_AGGREGATIONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AGGREGATIONTYPE
        default:
            return 0, errors.New("Unknown AggregationType value: " + v)
    }
    return &result, nil
}
func SerializeAggregationType(values []AggregationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
