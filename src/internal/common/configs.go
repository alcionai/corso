package common

import "golang.org/x/exp/maps"

type StringConfigurer interface {
	StringConfig() (map[string]string, error)
}

// UnionStringConfigs unions all provided configurers into a single
// map[string]string matching type.
func UnionStringConfigs(cfgs ...StringConfigurer) (map[string]string, error) {
	union := map[string]string{}

	for _, cfg := range cfgs {
		c, err := cfg.StringConfig()
		if err != nil {
			return nil, err
		}

		maps.Copy(union, c)
	}

	return union, nil
}
