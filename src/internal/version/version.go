package version

import (
	"os/exec"
	"strings"
)

var Version = "dev"

func CurrentVersion() string {
	if len(Version) == 0 || Version == "dev" {
		c, b := exec.Command("git", "describe", "--tag"), new(strings.Builder)
		c.Stdout = b

		if err := c.Run(); err != nil {
			return "dev"
		}

		s := strings.TrimRight(b.String(), "\n")

		if len(s) != 0 {
			return "dev-" + s
		}

		return "dev"
	}

	return Version
}
