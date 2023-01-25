package models
import (
    "errors"
)
// Provides operations to manage the collection of site entities.
type TranslationBehavior int

const (
    ASK_TRANSLATIONBEHAVIOR TranslationBehavior = iota
    YES_TRANSLATIONBEHAVIOR
    NO_TRANSLATIONBEHAVIOR
)

func (i TranslationBehavior) String() string {
    return []string{"Ask", "Yes", "No"}[i]
}
func ParseTranslationBehavior(v string) (interface{}, error) {
    result := ASK_TRANSLATIONBEHAVIOR
    switch v {
        case "Ask":
            result = ASK_TRANSLATIONBEHAVIOR
        case "Yes":
            result = YES_TRANSLATIONBEHAVIOR
        case "No":
            result = NO_TRANSLATIONBEHAVIOR
        default:
            return 0, errors.New("Unknown TranslationBehavior value: " + v)
    }
    return &result, nil
}
func SerializeTranslationBehavior(values []TranslationBehavior) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
