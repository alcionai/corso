package security
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type FileProcessingStatus int

const (
    SUCCESS_FILEPROCESSINGSTATUS FileProcessingStatus = iota
    INTERNALERROR_FILEPROCESSINGSTATUS
    UNKNOWNERROR_FILEPROCESSINGSTATUS
    PROCESSINGTIMEOUT_FILEPROCESSINGSTATUS
    INVALIDFILEID_FILEPROCESSINGSTATUS
    FILESIZEISZERO_FILEPROCESSINGSTATUS
    FILESIZEISTOOLARGE_FILEPROCESSINGSTATUS
    FILEDEPTHLIMITEXCEEDED_FILEPROCESSINGSTATUS
    FILEBODYISTOOLONG_FILEPROCESSINGSTATUS
    FILETYPEISUNKNOWN_FILEPROCESSINGSTATUS
    FILETYPEISNOTSUPPORTED_FILEPROCESSINGSTATUS
    MALFORMEDFILE_FILEPROCESSINGSTATUS
    PROTECTEDFILE_FILEPROCESSINGSTATUS
    POISONFILE_FILEPROCESSINGSTATUS
    NOREVIEWSETSUMMARYGENERATED_FILEPROCESSINGSTATUS
    EXTRACTIONEXCEPTION_FILEPROCESSINGSTATUS
    OCRPROCESSINGTIMEOUT_FILEPROCESSINGSTATUS
    OCRFILESIZEEXCEEDSLIMIT_FILEPROCESSINGSTATUS
    UNKNOWNFUTUREVALUE_FILEPROCESSINGSTATUS
)

func (i FileProcessingStatus) String() string {
    return []string{"success", "internalError", "unknownError", "processingTimeout", "invalidFileId", "fileSizeIsZero", "fileSizeIsTooLarge", "fileDepthLimitExceeded", "fileBodyIsTooLong", "fileTypeIsUnknown", "fileTypeIsNotSupported", "malformedFile", "protectedFile", "poisonFile", "noReviewSetSummaryGenerated", "extractionException", "ocrProcessingTimeout", "ocrFileSizeExceedsLimit", "unknownFutureValue"}[i]
}
func ParseFileProcessingStatus(v string) (interface{}, error) {
    result := SUCCESS_FILEPROCESSINGSTATUS
    switch v {
        case "success":
            result = SUCCESS_FILEPROCESSINGSTATUS
        case "internalError":
            result = INTERNALERROR_FILEPROCESSINGSTATUS
        case "unknownError":
            result = UNKNOWNERROR_FILEPROCESSINGSTATUS
        case "processingTimeout":
            result = PROCESSINGTIMEOUT_FILEPROCESSINGSTATUS
        case "invalidFileId":
            result = INVALIDFILEID_FILEPROCESSINGSTATUS
        case "fileSizeIsZero":
            result = FILESIZEISZERO_FILEPROCESSINGSTATUS
        case "fileSizeIsTooLarge":
            result = FILESIZEISTOOLARGE_FILEPROCESSINGSTATUS
        case "fileDepthLimitExceeded":
            result = FILEDEPTHLIMITEXCEEDED_FILEPROCESSINGSTATUS
        case "fileBodyIsTooLong":
            result = FILEBODYISTOOLONG_FILEPROCESSINGSTATUS
        case "fileTypeIsUnknown":
            result = FILETYPEISUNKNOWN_FILEPROCESSINGSTATUS
        case "fileTypeIsNotSupported":
            result = FILETYPEISNOTSUPPORTED_FILEPROCESSINGSTATUS
        case "malformedFile":
            result = MALFORMEDFILE_FILEPROCESSINGSTATUS
        case "protectedFile":
            result = PROTECTEDFILE_FILEPROCESSINGSTATUS
        case "poisonFile":
            result = POISONFILE_FILEPROCESSINGSTATUS
        case "noReviewSetSummaryGenerated":
            result = NOREVIEWSETSUMMARYGENERATED_FILEPROCESSINGSTATUS
        case "extractionException":
            result = EXTRACTIONEXCEPTION_FILEPROCESSINGSTATUS
        case "ocrProcessingTimeout":
            result = OCRPROCESSINGTIMEOUT_FILEPROCESSINGSTATUS
        case "ocrFileSizeExceedsLimit":
            result = OCRFILESIZEEXCEEDSLIMIT_FILEPROCESSINGSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_FILEPROCESSINGSTATUS
        default:
            return 0, errors.New("Unknown FileProcessingStatus value: " + v)
    }
    return &result, nil
}
func SerializeFileProcessingStatus(values []FileProcessingStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
