package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type SkillProficiencyLevel int

const (
    ELEMENTARY_SKILLPROFICIENCYLEVEL SkillProficiencyLevel = iota
    LIMITEDWORKING_SKILLPROFICIENCYLEVEL
    GENERALPROFESSIONAL_SKILLPROFICIENCYLEVEL
    ADVANCEDPROFESSIONAL_SKILLPROFICIENCYLEVEL
    EXPERT_SKILLPROFICIENCYLEVEL
    UNKNOWNFUTUREVALUE_SKILLPROFICIENCYLEVEL
)

func (i SkillProficiencyLevel) String() string {
    return []string{"elementary", "limitedWorking", "generalProfessional", "advancedProfessional", "expert", "unknownFutureValue"}[i]
}
func ParseSkillProficiencyLevel(v string) (interface{}, error) {
    result := ELEMENTARY_SKILLPROFICIENCYLEVEL
    switch v {
        case "elementary":
            result = ELEMENTARY_SKILLPROFICIENCYLEVEL
        case "limitedWorking":
            result = LIMITEDWORKING_SKILLPROFICIENCYLEVEL
        case "generalProfessional":
            result = GENERALPROFESSIONAL_SKILLPROFICIENCYLEVEL
        case "advancedProfessional":
            result = ADVANCEDPROFESSIONAL_SKILLPROFICIENCYLEVEL
        case "expert":
            result = EXPERT_SKILLPROFICIENCYLEVEL
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SKILLPROFICIENCYLEVEL
        default:
            return 0, errors.New("Unknown SkillProficiencyLevel value: " + v)
    }
    return &result, nil
}
func SerializeSkillProficiencyLevel(values []SkillProficiencyLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
