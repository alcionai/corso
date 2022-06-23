package common

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
		for k, v := range c {
			union[k] = v
		}
	}
	return union, nil
}
