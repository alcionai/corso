package flags

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	UserFN              = "user"
	MailBoxFN           = "mailbox"
	AzureClientTenantFN = "azure-tenant-id"
	AzureClientIDFN     = "azure-client-id"
	AzureClientSecretFN = "azure-client-secret"
)

var (
	UserFV              []string
	AzureClientTenantFV string
	AzureClientIDFV     string
	AzureClientSecretFV string
)

// AddUserFlag adds the --user flag.
func AddUserFlag(
	cmd *cobra.Command,
	completionFunc func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective),
) {
	cmd.Flags().StringSliceVar(
		&UserFV,
		UserFN, nil,
		"Backup a specific user's data; accepts '"+Wildcard+"' to select all users.")
	cobra.CheckErr(cmd.MarkFlagRequired(UserFN))

	_ = cmd.RegisterFlagCompletionFunc(UserFN, completionFunc)
}

// AddMailBoxFlag adds the --user and --mailbox flag.
func AddMailBoxFlag(
	cmd *cobra.Command,
	completionFunc func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective),
) {
	flags := cmd.Flags()

	flags.StringSliceVar(
		&UserFV,
		UserFN, nil,
		"Backup a specific user's data; accepts '"+Wildcard+"' to select all users.")

	cobra.CheckErr(flags.MarkDeprecated(UserFN, fmt.Sprintf("use --%s instead", MailBoxFN)))

	_ = cmd.RegisterFlagCompletionFunc(UserFN,
		func(
			cmd *cobra.Command,
			args []string,
			toComplete string,
		) ([]string, cobra.ShellCompDirective) {
			message := fmt.Sprintf("This flag is deprecated, Use --%s instead", MailBoxFN)
			return cobra.AppendActiveHelp(nil, message), cobra.ShellCompDirectiveNoFileComp
		})

	flags.StringSliceVar(
		&UserFV,
		MailBoxFN, nil,
		"Backup a specific mailbox's data; accepts '"+Wildcard+"' to select all mailbox.")

	_ = cmd.RegisterFlagCompletionFunc(MailBoxFN, completionFunc)
}

// AddAzureCredsFlags adds M365 cred flags
func AddAzureCredsFlags(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.StringVar(&AzureClientTenantFV, AzureClientTenantFN, "", "Azure tenant ID")
	fs.StringVar(&AzureClientIDFV, AzureClientIDFN, "", "Azure app client ID")
	fs.StringVar(&AzureClientSecretFV, AzureClientSecretFN, "", "Azure app client secret")
}
