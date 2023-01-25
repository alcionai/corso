package ediscovery
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type ExportOptions int

const (
    ORIGINALFILES_EXPORTOPTIONS ExportOptions = iota
    TEXT_EXPORTOPTIONS
    PDFREPLACEMENT_EXPORTOPTIONS
    FILEINFO_EXPORTOPTIONS
    TAGS_EXPORTOPTIONS
    UNKNOWNFUTUREVALUE_EXPORTOPTIONS
)

func (i ExportOptions) String() string {
    return []string{"originalFiles", "text", "pdfReplacement", "fileInfo", "tags", "unknownFutureValue"}[i]
}
func ParseExportOptions(v string) (interface{}, error) {
    result := ORIGINALFILES_EXPORTOPTIONS
    switch v {
        case "originalFiles":
            result = ORIGINALFILES_EXPORTOPTIONS
        case "text":
            result = TEXT_EXPORTOPTIONS
        case "pdfReplacement":
            result = PDFREPLACEMENT_EXPORTOPTIONS
        case "fileInfo":
            result = FILEINFO_EXPORTOPTIONS
        case "tags":
            result = TAGS_EXPORTOPTIONS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_EXPORTOPTIONS
        default:
            return 0, errors.New("Unknown ExportOptions value: " + v)
    }
    return &result, nil
}
func SerializeExportOptions(values []ExportOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
