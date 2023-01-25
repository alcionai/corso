package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ReplyRestriction int

const (
    EVERYONE_REPLYRESTRICTION ReplyRestriction = iota
    AUTHORANDMODERATORS_REPLYRESTRICTION
    UNKNOWNFUTUREVALUE_REPLYRESTRICTION
)

func (i ReplyRestriction) String() string {
    return []string{"everyone", "authorAndModerators", "unknownFutureValue"}[i]
}
func ParseReplyRestriction(v string) (interface{}, error) {
    result := EVERYONE_REPLYRESTRICTION
    switch v {
        case "everyone":
            result = EVERYONE_REPLYRESTRICTION
        case "authorAndModerators":
            result = AUTHORANDMODERATORS_REPLYRESTRICTION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_REPLYRESTRICTION
        default:
            return 0, errors.New("Unknown ReplyRestriction value: " + v)
    }
    return &result, nil
}
func SerializeReplyRestriction(values []ReplyRestriction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
