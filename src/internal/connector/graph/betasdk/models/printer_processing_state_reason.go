package models
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type PrinterProcessingStateReason int

const (
    PAUSED_PRINTERPROCESSINGSTATEREASON PrinterProcessingStateReason = iota
    MEDIAJAM_PRINTERPROCESSINGSTATEREASON
    MEDIANEEDED_PRINTERPROCESSINGSTATEREASON
    MEDIALOW_PRINTERPROCESSINGSTATEREASON
    MEDIAEMPTY_PRINTERPROCESSINGSTATEREASON
    COVEROPEN_PRINTERPROCESSINGSTATEREASON
    INTERLOCKOPEN_PRINTERPROCESSINGSTATEREASON
    OUTPUTTRAYMISSING_PRINTERPROCESSINGSTATEREASON
    OUTPUTAREAFULL_PRINTERPROCESSINGSTATEREASON
    MARKERSUPPLYLOW_PRINTERPROCESSINGSTATEREASON
    MARKERSUPPLYEMPTY_PRINTERPROCESSINGSTATEREASON
    INPUTTRAYMISSING_PRINTERPROCESSINGSTATEREASON
    OUTPUTAREAALMOSTFULL_PRINTERPROCESSINGSTATEREASON
    MARKERWASTEALMOSTFULL_PRINTERPROCESSINGSTATEREASON
    MARKERWASTEFULL_PRINTERPROCESSINGSTATEREASON
    FUSEROVERTEMP_PRINTERPROCESSINGSTATEREASON
    FUSERUNDERTEMP_PRINTERPROCESSINGSTATEREASON
    OTHER_PRINTERPROCESSINGSTATEREASON
    NONE_PRINTERPROCESSINGSTATEREASON
    MOVINGTOPAUSED_PRINTERPROCESSINGSTATEREASON
    SHUTDOWN_PRINTERPROCESSINGSTATEREASON
    CONNECTINGTODEVICE_PRINTERPROCESSINGSTATEREASON
    TIMEDOUT_PRINTERPROCESSINGSTATEREASON
    STOPPING_PRINTERPROCESSINGSTATEREASON
    STOPPEDPARTIALLY_PRINTERPROCESSINGSTATEREASON
    TONERLOW_PRINTERPROCESSINGSTATEREASON
    TONEREMPTY_PRINTERPROCESSINGSTATEREASON
    SPOOLAREAFULL_PRINTERPROCESSINGSTATEREASON
    DOOROPEN_PRINTERPROCESSINGSTATEREASON
    OPTICALPHOTOCONDUCTORNEARENDOFLIFE_PRINTERPROCESSINGSTATEREASON
    OPTICALPHOTOCONDUCTORLIFEOVER_PRINTERPROCESSINGSTATEREASON
    DEVELOPERLOW_PRINTERPROCESSINGSTATEREASON
    DEVELOPEREMPTY_PRINTERPROCESSINGSTATEREASON
    INTERPRETERRESOURCEUNAVAILABLE_PRINTERPROCESSINGSTATEREASON
    UNKNOWNFUTUREVALUE_PRINTERPROCESSINGSTATEREASON
)

func (i PrinterProcessingStateReason) String() string {
    return []string{"paused", "mediaJam", "mediaNeeded", "mediaLow", "mediaEmpty", "coverOpen", "interlockOpen", "outputTrayMissing", "outputAreaFull", "markerSupplyLow", "markerSupplyEmpty", "inputTrayMissing", "outputAreaAlmostFull", "markerWasteAlmostFull", "markerWasteFull", "fuserOverTemp", "fuserUnderTemp", "other", "none", "movingToPaused", "shutdown", "connectingToDevice", "timedOut", "stopping", "stoppedPartially", "tonerLow", "tonerEmpty", "spoolAreaFull", "doorOpen", "opticalPhotoConductorNearEndOfLife", "opticalPhotoConductorLifeOver", "developerLow", "developerEmpty", "interpreterResourceUnavailable", "unknownFutureValue"}[i]
}
func ParsePrinterProcessingStateReason(v string) (interface{}, error) {
    result := PAUSED_PRINTERPROCESSINGSTATEREASON
    switch v {
        case "paused":
            result = PAUSED_PRINTERPROCESSINGSTATEREASON
        case "mediaJam":
            result = MEDIAJAM_PRINTERPROCESSINGSTATEREASON
        case "mediaNeeded":
            result = MEDIANEEDED_PRINTERPROCESSINGSTATEREASON
        case "mediaLow":
            result = MEDIALOW_PRINTERPROCESSINGSTATEREASON
        case "mediaEmpty":
            result = MEDIAEMPTY_PRINTERPROCESSINGSTATEREASON
        case "coverOpen":
            result = COVEROPEN_PRINTERPROCESSINGSTATEREASON
        case "interlockOpen":
            result = INTERLOCKOPEN_PRINTERPROCESSINGSTATEREASON
        case "outputTrayMissing":
            result = OUTPUTTRAYMISSING_PRINTERPROCESSINGSTATEREASON
        case "outputAreaFull":
            result = OUTPUTAREAFULL_PRINTERPROCESSINGSTATEREASON
        case "markerSupplyLow":
            result = MARKERSUPPLYLOW_PRINTERPROCESSINGSTATEREASON
        case "markerSupplyEmpty":
            result = MARKERSUPPLYEMPTY_PRINTERPROCESSINGSTATEREASON
        case "inputTrayMissing":
            result = INPUTTRAYMISSING_PRINTERPROCESSINGSTATEREASON
        case "outputAreaAlmostFull":
            result = OUTPUTAREAALMOSTFULL_PRINTERPROCESSINGSTATEREASON
        case "markerWasteAlmostFull":
            result = MARKERWASTEALMOSTFULL_PRINTERPROCESSINGSTATEREASON
        case "markerWasteFull":
            result = MARKERWASTEFULL_PRINTERPROCESSINGSTATEREASON
        case "fuserOverTemp":
            result = FUSEROVERTEMP_PRINTERPROCESSINGSTATEREASON
        case "fuserUnderTemp":
            result = FUSERUNDERTEMP_PRINTERPROCESSINGSTATEREASON
        case "other":
            result = OTHER_PRINTERPROCESSINGSTATEREASON
        case "none":
            result = NONE_PRINTERPROCESSINGSTATEREASON
        case "movingToPaused":
            result = MOVINGTOPAUSED_PRINTERPROCESSINGSTATEREASON
        case "shutdown":
            result = SHUTDOWN_PRINTERPROCESSINGSTATEREASON
        case "connectingToDevice":
            result = CONNECTINGTODEVICE_PRINTERPROCESSINGSTATEREASON
        case "timedOut":
            result = TIMEDOUT_PRINTERPROCESSINGSTATEREASON
        case "stopping":
            result = STOPPING_PRINTERPROCESSINGSTATEREASON
        case "stoppedPartially":
            result = STOPPEDPARTIALLY_PRINTERPROCESSINGSTATEREASON
        case "tonerLow":
            result = TONERLOW_PRINTERPROCESSINGSTATEREASON
        case "tonerEmpty":
            result = TONEREMPTY_PRINTERPROCESSINGSTATEREASON
        case "spoolAreaFull":
            result = SPOOLAREAFULL_PRINTERPROCESSINGSTATEREASON
        case "doorOpen":
            result = DOOROPEN_PRINTERPROCESSINGSTATEREASON
        case "opticalPhotoConductorNearEndOfLife":
            result = OPTICALPHOTOCONDUCTORNEARENDOFLIFE_PRINTERPROCESSINGSTATEREASON
        case "opticalPhotoConductorLifeOver":
            result = OPTICALPHOTOCONDUCTORLIFEOVER_PRINTERPROCESSINGSTATEREASON
        case "developerLow":
            result = DEVELOPERLOW_PRINTERPROCESSINGSTATEREASON
        case "developerEmpty":
            result = DEVELOPEREMPTY_PRINTERPROCESSINGSTATEREASON
        case "interpreterResourceUnavailable":
            result = INTERPRETERRESOURCEUNAVAILABLE_PRINTERPROCESSINGSTATEREASON
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PRINTERPROCESSINGSTATEREASON
        default:
            return 0, errors.New("Unknown PrinterProcessingStateReason value: " + v)
    }
    return &result, nil
}
func SerializePrinterProcessingStateReason(values []PrinterProcessingStateReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
