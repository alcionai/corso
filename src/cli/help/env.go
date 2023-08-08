package help

import (
	"github.com/spf13/cobra"

	. "github.com/alcionai/corso/src/cli/print"
)

// AddCommands attaches all `corso env * *` commands to the parent.
func AddCommands(cmd *cobra.Command) {
	cmd.AddCommand(envCmd())
}

// The env command: purely a help display.
// `corso env [--help]`
func envCmd() *cobra.Command {
	envCmd := &cobra.Command{
		Use:   "env",
		Short: "env var guide",
		Long:  `A guide to using environment variables in Corso.`,
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
	category    string
	name        string
	description string
}

// interface compliance check
var _ Printable = &envVar{}

// no modifications needed, just passthrough.
func (ev envVar) MinimumPrintable() any {
	return ev
}

func (ev envVar) Headers() []string {
	return []string{ev.category, " "}
}

func (ev envVar) Values() []string {
	return []string{ev.name, ev.description}
}

// headers
const (
	corso = "Corso"
	azure = "Azure AD App Credentials"
	aws   = "AWS Credentials"
)

var (
	corsoEVs = []envVar{
		{corso, "CORSO_PASSPHRASE", "Passphrase to protect encrypted repository contents. " +
			"It is impossible to use the repository or recover any backups without this key."},
	}
	azureEVs = []envVar{
		{azure, "AZURE_CLIENT_ID", "Client ID for your Azure AD application used to access your M365 tenant."},
		{azure, "AZURE_TENANT_ID", "ID for the M365 tenant where the Azure AD application is registered."},
		{azure, "AZURE_CLIENT_SECRET", "Azure secret for your Azure AD application used to access your M365 tenant."},
	}
	awsEVs = []envVar{
		{aws, "AWS_ACCESS_KEY_ID", "Access key for an IAM user or role for accessing an S3 bucket."},
		{aws, "AWS_SECRET_ACCESS_KEY", "Secret key associated with the access key."},
		{aws, "AWS_SESSION_TOKEN", "Session token required when using temporary credentials."},
	}
)

func toPrintable(evs []envVar) []Printable {
	ps := []Printable{}
	for _, ev := range evs {
		ps = append(ps, ev)
	}

	return ps
}

// envGuide outputs a help menu for setting env vars in Corso.
func envGuide(cmd *cobra.Command, args []string) {
	ctx := cmd.Context()
	Info(ctx,
		"\n--- Environment Variable Guide ---\n",
		"As a best practice, Corso retrieves credentials and sensitive information from environment variables.\n ",
		"\n",
	)
	Table(ctx, toPrintable(corsoEVs))
	Info(ctx, "\n")
	Table(ctx, toPrintable(azureEVs))
	Info(ctx, "\n")
	Table(ctx, toPrintable(awsEVs))
}
