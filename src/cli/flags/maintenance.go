package flags

import (
	"github.com/spf13/cobra"

	"github.com/alcionai/corso/src/pkg/control/repository"
)

const (
	MaintenanceModeFN     = "mode"
	ForceMaintenanceFN    = "force"
	UserMaintenanceFN     = "user"
	HostnameMaintenanceFN = "host"
)

var (
	MaintenanceModeFV     string
	ForceMaintenanceFV    bool
	UserMaintenanceFV     string
	HostnameMaintenanceFV string
)

func AddMaintenanceModeFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.StringVar(
		&MaintenanceModeFV,
		MaintenanceModeFN,
		repository.CompleteMaintenance.String(),
		"Type of maintenance operation to run. Pass '"+
			repository.MetadataMaintenance.String()+"' to run a faster maintenance "+
			"that does minimal clean-up and optimization. Pass '"+
			repository.CompleteMaintenance.String()+"' to fully compact existing "+
			"data and delete unused data.")
}

func AddForceMaintenanceFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.BoolVar(
		&ForceMaintenanceFV,
		ForceMaintenanceFN,
		false,
		"Force maintenance. Caution: user must ensure this is not run concurrently on a single repo")
	cobra.CheckErr(fs.MarkHidden(ForceMaintenanceFN))
}

func AddMaintenanceUserFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.StringVar(
		&UserMaintenanceFV,
		UserMaintenanceFN,
		"",
		"Attempt to run maintenance as the specified user for the repo owner user")
}

func AddMaintenanceHostnameFlag(cmd *cobra.Command) {
	fs := cmd.Flags()
	fs.StringVar(
		&HostnameMaintenanceFV,
		HostnameMaintenanceFN,
		"",
		"Attempt to run maintenance with the specified hostname for the repo owner hostname")
}
