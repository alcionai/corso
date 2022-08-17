package help

import (
	"fmt"

	"github.com/spf13/cobra"

	. "github.com/alcionai/corso/cli/print"
)

// AddCommands attaches all `corso env * *` commands to the parent.
func AddCommands(parent *cobra.Command) {
	parent.AddCommand(envCmd())
}

// The env command: purely a help display.
// `corso env [--help]`
func envCmd() *cobra.Command {
	envCmd := &cobra.Command{
		Use:   "env",
		Short: "env var guide",
		Long:  `A guide to using env variables in Corso.`,
		RunE:  handleEnvCmd,
		Args:  cobra.NoArgs,
	}
	envCmd.SetHelpFunc(envGuide)
	return envCmd
}

// Handler for flat calls to `corso env`.
// Produces the same output as `corso env --help`.
func handleEnvCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

type envVar struct {
	name        string
	description string
}

func (ev envVar) String() string {
	return fmt.Sprintf("%s\t%s", ev.name, ev.description)
}

type envVarSet struct {
	name string
	vars []envVar
}

func (evs envVarSet) Strings() []any {
	s := []any{"\n", "- " + evs.name + "\n"}
	for _, ev := range evs.vars {
		s = append(s, ev.String()+"\n")
	}
	return s
}

var (
	corsoEVs = envVarSet{
		"Corso",
		[]envVar{{"CORSO_PASSWORD", "Protects repository encryption keys."}},
	}
	azureEVs = envVarSet{
		"Azure",
		[]envVar{
			{"CLIENT_ID", "Azure client ID for your m365 Tenant."},
			{"CLIENT_SECRET", "Azure secret for your m365 Tenant."},
		},
	}
	awsEVs = envVarSet{
		"AWS",
		[]envVar{
			{"AWS_ACCESS_KEY_ID", "Access Key to communicate to your s3 bucket."},
			{"AWS_SECRET_ACCESS_KEY", "Secret Key to communicate to your s3 bucket."},
			{"AWS_SESSION_TOKEN", "Temporary session token for AWS communication."},
		},
	}
)

// envGuide outputs a help menu for setting env vars in Corso.
func envGuide(cmd *cobra.Command, args []string) {
	guide := []any{
		"\n--- Environment Variable Guide ---\n",
		"In order to keep your information secure, Corso retrieves" +
			"credentials and other secrets from environment variables.\n ",
	}
	guide = append(guide, corsoEVs.Strings()...)
	guide = append(guide, azureEVs.Strings()...)
	guide = append(guide, awsEVs.Strings()...)

	Info(cmd.Context(), guide...)
}
