package storage

// config exported name consts
const (
	FilesystemPath = "filesystem_path"
)

type FsConfig struct {
	FilesystemPath string
}

func (s Storage) FsConfig() (FsConfig, error) {
	c := FsConfig{}

	if len(s.Config) > 0 {
		c.FilesystemPath = orEmptyString(s.Config[FilesystemPath])
	}

	return c, c.validate()
}

func (c FsConfig) validate() error {
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
