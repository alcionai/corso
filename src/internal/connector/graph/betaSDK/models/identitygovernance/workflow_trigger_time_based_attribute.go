package identitygovernance
import (
    "errors"
)
// Provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type WorkflowTriggerTimeBasedAttribute int

const (
    EMPLOYEEHIREDATE_WORKFLOWTRIGGERTIMEBASEDATTRIBUTE WorkflowTriggerTimeBasedAttribute = iota
    EMPLOYEELEAVEDATETIME_WORKFLOWTRIGGERTIMEBASEDATTRIBUTE
    UNKNOWNFUTUREVALUE_WORKFLOWTRIGGERTIMEBASEDATTRIBUTE
)

func (i WorkflowTriggerTimeBasedAttribute) String() string {
    return []string{"employeeHireDate", "employeeLeaveDateTime", "unknownFutureValue"}[i]
}
func ParseWorkflowTriggerTimeBasedAttribute(v string) (interface{}, error) {
    result := EMPLOYEEHIREDATE_WORKFLOWTRIGGERTIMEBASEDATTRIBUTE
    switch v {
        case "employeeHireDate":
            result = EMPLOYEEHIREDATE_WORKFLOWTRIGGERTIMEBASEDATTRIBUTE
        case "employeeLeaveDateTime":
            result = EMPLOYEELEAVEDATETIME_WORKFLOWTRIGGERTIMEBASEDATTRIBUTE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_WORKFLOWTRIGGERTIMEBASEDATTRIBUTE
        default:
            return 0, errors.New("Unknown WorkflowTriggerTimeBasedAttribute value: " + v)
    }
    return &result, nil
}
func SerializeWorkflowTriggerTimeBasedAttribute(values []WorkflowTriggerTimeBasedAttribute) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
