package storage

type CommonConfig struct {
	CorsoPassword string
}

// envvar consts
const (
	CORSO_PASSWORD = "CORSO_PASSWORD"
)

// config key consts
const (
	keyCommonCorsoPassword = "common_corsoPassword"
)

func (c CommonConfig) Config() config {
	return config{
		keyCommonCorsoPassword: c.CorsoPassword,
	}
}

// CommonConfig retrieves the CommonConfig details from the Storage config.
func (s Storage) CommonConfig() CommonConfig {
	c := CommonConfig{}
	if len(s.Config) > 0 {
		c.CorsoPassword = orEmptyString(s.Config[keyCommonCorsoPassword])
	}
	return c
}
