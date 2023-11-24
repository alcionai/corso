package utils

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365"
	"github.com/alcionai/corso/src/pkg/store"
)

func GetBackups(ctx context.Context, cmd *cobra.Command, service path.ServiceType) ([]*backup.Backup, error) {
	r, _, err := GetAccountAndConnect(ctx, cmd, service)
	if err != nil {
		return nil, err
	}

	defer CloseRepo(ctx, r)

	return r.BackupsByTag(ctx, store.Service(service))
}

type completionFunc func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective)

func BackupIDCompletionFunc(service path.ServiceType) completionFunc {
	return func(
		cmd *cobra.Command,
		args []string,
		toComplete string,
	) ([]string, cobra.ShellCompDirective) {
		bs, err := GetBackups(cmd.Context(), cmd, service)
		if err != nil {
			return cobra.AppendActiveHelp(
				nil,
				fmt.Sprintf("Unable to fetch %s backups", service)), cobra.ShellCompDirectiveNoFileComp
		}

		if len(bs) == 0 {
			return cobra.AppendActiveHelp(
				nil,
				fmt.Sprintf("No %s backups found", service)), cobra.ShellCompDirectiveNoFileComp
		}

		backups := make([]string, len(bs))
		for _, b := range bs {
			backups = append(backups, fmt.Sprintf("%s\tCreated at %s", b.GetID(), b.CreationTime))
		}

		return cobra.AppendActiveHelp(backups, "Choose backup ID to use"), cobra.ShellCompDirectiveNoFileComp
	}
}

func UsersCompletionFunc(service path.ServiceType) completionFunc {
	return func(
		cmd *cobra.Command,
		args []string,
		toComplete string,
	) ([]string, cobra.ShellCompDirective) {
		ctx := cmd.Context()

		r, acct, err := AccountConnectAndWriteRepoConfig(ctx, cmd, service)
		if err != nil {
			return cobra.AppendActiveHelp(
				nil,
				fmt.Sprintf("Unable to fetch %s users", service)), cobra.ShellCompDirectiveNoFileComp
		}

		ins, err := UsersMap(ctx, *acct, Control(), r.Counter(), fault.New(true))
		if err != nil {
			return cobra.AppendActiveHelp(
				nil,
				fmt.Sprintf("Unable to fetch %s users", service)), cobra.ShellCompDirectiveNoFileComp
		}

		if len(ins.IDs()) == 0 {
			return cobra.AppendActiveHelp(
				nil,
				fmt.Sprintf("No %s users found", service)), cobra.ShellCompDirectiveNoFileComp
		}

		backups := make([]string, len(ins.IDs()))

		for _, u := range ins.IDs() {
			name, _ := ins.NameOf(u)
			backups = append(backups, fmt.Sprintf("%s\t%s", u, name))
		}

		return cobra.AppendActiveHelp(backups, "Choose user ID to use"), cobra.ShellCompDirectiveNoFileComp
	}
}

func MailboxCompletionFunc(service path.ServiceType) completionFunc {
	return func(
		cmd *cobra.Command,
		args []string,
		toComplete string,
	) ([]string, cobra.ShellCompDirective) {
		ctx := cmd.Context()

		r, acct, err := AccountConnectAndWriteRepoConfig(ctx, cmd, service)
		if err != nil {
			return cobra.AppendActiveHelp(
				nil,
				fmt.Sprintf("Unable to fetch %s mailboxes", service)), cobra.ShellCompDirectiveNoFileComp
		}

		ins, err := UsersMap(ctx, *acct, Control(), r.Counter(), fault.New(true))
		if err != nil {
			return cobra.AppendActiveHelp(
				nil,
				fmt.Sprintf("Unable to fetch %s mailboxes", service)), cobra.ShellCompDirectiveNoFileComp
		}

		if len(ins.IDs()) == 0 {
			return cobra.AppendActiveHelp(
				nil,
				fmt.Sprintf("No %s mailboxes found", service)), cobra.ShellCompDirectiveNoFileComp
		}

		backups := ins.Names()

		return cobra.AppendActiveHelp(backups, "Choose mailbox to use"), cobra.ShellCompDirectiveNoFileComp
	}
}

func GroupsCompletionFunc() completionFunc {
	return func(
		cmd *cobra.Command,
		args []string,
		toComplete string,
	) ([]string, cobra.ShellCompDirective) {
		ctx := cmd.Context()

		_, acct, err := AccountConnectAndWriteRepoConfig(ctx, cmd, path.GroupsService)
		if err != nil {
			return cobra.AppendActiveHelp(nil, "Unable to fetch Groups"), cobra.ShellCompDirectiveNoFileComp
		}

		ins, err := m365.GroupsMap(ctx, *acct, fault.New(true))
		if err != nil {
			return cobra.AppendActiveHelp(nil, "Unable to fetch Groups"), cobra.ShellCompDirectiveNoFileComp
		}

		if len(ins.IDs()) == 0 {
			return cobra.AppendActiveHelp(nil, "No Groups found"), cobra.ShellCompDirectiveNoFileComp
		}

		backups := make([]string, len(ins.IDs()))

		for _, u := range ins.IDs() {
			name, _ := ins.NameOf(u)
			backups = append(backups, fmt.Sprintf("%s\t%s", u, name))
		}

		return cobra.AppendActiveHelp(backups, "Choose Group to use"), cobra.ShellCompDirectiveNoFileComp
	}
}

func SitesCompletionFunc() completionFunc {
	return func(
		cmd *cobra.Command,
		args []string,
		toComplete string,
	) ([]string, cobra.ShellCompDirective) {
		ctx := cmd.Context()

		_, acct, err := AccountConnectAndWriteRepoConfig(ctx, cmd, path.SharePointService)
		if err != nil {
			return cobra.AppendActiveHelp(nil, "Unable to fetch Sites"), cobra.ShellCompDirectiveNoFileComp
		}

		ins, err := m365.SitesMap(ctx, *acct, fault.New(true))
		if err != nil {
			return cobra.AppendActiveHelp(nil, "Unable to fetch Sites"), cobra.ShellCompDirectiveNoFileComp
		}

		if len(ins.IDs()) == 0 {
			return cobra.AppendActiveHelp(nil, "No Sites found"), cobra.ShellCompDirectiveNoFileComp
		}

		backups := make([]string, len(ins.IDs()))

		for _, u := range ins.IDs() {
			name, _ := ins.NameOf(u)
			backups = append(backups, fmt.Sprintf("%s\t%s", u, name))
		}

		return cobra.AppendActiveHelp(backups, "Choose Site to use"), cobra.ShellCompDirectiveNoFileComp
	}
}
