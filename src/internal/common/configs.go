package common

type StringMapper[T any] interface {
	~map[string]T
}

type Configurer[V any, T StringMapper[V]] interface {
	Config() (T, error)
}

type (
	Config[T any] map[string]T
)

// UnionConfigs unions all provided configurers into a single
// map[string]string matching type.
func UnionConfigs[
	V any,
	T StringMapper[V],
](
	cfgs ...Configurer[V, T],
) (T, error) {
	union := T{}
	for _, cfg := range cfgs {
		c, err := cfg.Config()
		if err != nil {
			return nil, err
		}
		for k, v := range c {
			union[k] = v
		}
	}
	return union, nil
}
