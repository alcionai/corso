package storage

type FilesystemConfig struct {
	FilesystemPath string
}

func (s Storage) FsConfig() (FilesystemConfig, error) {
	c := FilesystemConfig{}

	if len(s.Config) > 0 {
		c.FilesystemPath = orEmptyString(s.Config["path"])
	}

	return c, c.validate()
}

func (c FilesystemConfig) validate() error {
	// check := map[string]string{
	// 	Bucket: c.Bucket,
	// }
	// for k, v := range check {
	// 	if len(v) == 0 {
	// 		return clues.Stack(errMissingRequired, clues.New(k))
	// 	}
	// }
	return nil
}
