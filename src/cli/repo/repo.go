package repo

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/alcionai/clues"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	repo "github.com/alcionai/corso/src/pkg/repository"
)

const (
	initCommand        = "init"
	connectCommand     = "connect"
	maintenanceCommand = "maintenance"
)

var (
	ErrConnectingRepo   = clues.New("connecting repository")
	ErrInitializingRepo = clues.New("initializing repository")
)

var repoCommands = []func(cmd *cobra.Command) *cobra.Command{
	addS3Commands,
	addFilesystemCommands,
}

// AddCommands attaches all `corso repo * *` commands to the parent.
func AddCommands(cmd *cobra.Command) {
	var (
		// Get new instances so that setting the context during tests works
		// properly.
		repoCmd        = repoCmd()
		initCmd        = initCmd()
		connectCmd     = connectCmd()
		maintenanceCmd = maintenanceCmd()
	)

	cmd.AddCommand(repoCmd)
	repoCmd.AddCommand(initCmd)
	repoCmd.AddCommand(connectCmd)
	repoCmd.AddCommand(maintenanceCmd)

	flags.AddMaintenanceModeFlag(maintenanceCmd)
	flags.AddForceMaintenanceFlag(maintenanceCmd)
	flags.AddMaintenanceUserFlag(maintenanceCmd)
	flags.AddMaintenanceHostnameFlag(maintenanceCmd)

	for _, addRepoTo := range repoCommands {
		addRepoTo(initCmd)
		addRepoTo(connectCmd)
	}
}

// The repo category of commands.
// `corso repo [<subcommand>] [<flag>...]`
func repoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "repo",
		Short: "Manage your repositories",
		Long:  `Initialize, configure, and connect to your account backup repositories.`,
		RunE:  handleRepoCmd,
		Args:  cobra.NoArgs,
	}
}

// Handler for flat calls to `corso repo`.
// Produces the same output as `corso repo --help`.
func handleRepoCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

// The repo init subcommand.
// `corso repo init <repository> [<flag>...]`
func initCmd() *cobra.Command {
	return &cobra.Command{
		Use:   initCommand,
		Short: "Initialize a repository.",
		Long:  `Create a new repository to store your backups.`,
		RunE:  handleInitCmd,
		Args:  cobra.NoArgs,
	}
}

// Handler for calls to `corso repo init`.
func handleInitCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

// The repo connect subcommand.
// `corso repo connect <repository> [<flag>...]`
func connectCmd() *cobra.Command {
	return &cobra.Command{
		Use:   connectCommand,
		Short: "Connect to a repository.",
		Long:  `Connect to an existing repository.`,
		RunE:  handleConnectCmd,
		Args:  cobra.NoArgs,
	}
}

// Handler for calls to `corso repo connect`.
func handleConnectCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

func maintenanceCmd() *cobra.Command {
	return &cobra.Command{
		Use:   maintenanceCommand,
		Short: "Run maintenance on an existing repository",
		Long:  `Run maintenance on an existing repository to optimize performance and storage use`,
		RunE:  handleMaintenanceCmd,
		Args:  cobra.NoArgs,
	}
}

func handleMaintenanceCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	t, err := getMaintenanceType(flags.MaintenanceModeFV)
	if err != nil {
		return err
	}

	r, _, err := utils.AccountConnectAndWriteRepoConfig(
		ctx,
		cmd,
		// Need to give it a valid service so it won't error out on us even though
		// we don't need the graph client.
		path.OneDriveService)
	if err != nil {
		return print.Only(ctx, err)
	}

	defer utils.CloseRepo(ctx, r)

	m, err := r.NewMaintenance(
		ctx,
		repository.Maintenance{
			Type:   t,
			Safety: repository.FullMaintenanceSafety,
			Force:  flags.ForceMaintenanceFV,
		})
	if err != nil {
		return print.Only(ctx, err)
	}

	err = m.Run(ctx)
	if err != nil {
		return print.Only(ctx, err)
	}

	return nil
}

func getMaintenanceType(t string) (repository.MaintenanceType, error) {
	res, ok := repository.StringToMaintenanceType[t]
	if !ok {
		modes := maps.Keys(repository.StringToMaintenanceType)
		allButLast := []string{}

		for i := 0; i < len(modes)-1; i++ {
			allButLast = append(allButLast, string(modes[i]))
		}

		valuesStr := strings.Join(allButLast, ", ") + " or " + string(modes[len(modes)-1])

		return res, clues.New(t + " is an unrecognized maintenance mode; must be one of " + valuesStr)
	}

	return res, nil
}

func printTree(root *repo.BackupNode, ident int) {
	if root == nil {
		return
	}

	fmt.Printf(strings.Repeat("\t", ident)+"%+v\n", root)

	for _, child := range root.Children {
		printTree(child.BackupNode, ident+1)
	}
}

func drawTree(ctx context.Context, roots []*repo.BackupNode) error {
	const port = ":6060"

	g := graphviz.New()

	graph, err := g.Graph()
	if err != nil {
		return clues.Wrap(err, "getting graph")
	}

	defer func() {
		graph.Close()
		g.Close()
	}()

	graph.SetRankDir(cgraph.LRRank)

	for _, root := range roots {
		if err := buildGraph(ctx, graph, root); err != nil {
			return clues.Wrap(err, "building graph")
		}
	}

	fmt.Printf("starting http server on port %s", port)

	// Start an http server that has the redered image.
	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "image/svg+xml")

			if err := g.Render(graph, graphviz.SVG, w); err != nil {
				logger.CtxErr(ctx, err).Info("sending svg to server")
			}
		},
	)

	return clues.Stack(http.ListenAndServe(port, nil)).OrNil()
}

func buildGraph(
	ctx context.Context,
	graph *cgraph.Graph,
	root *repo.BackupNode,
) error {
	// Add all nodes to the map and track them by ID. The root is just mapped as
	// "root" and has the resource ID.
	allNodes := map[string]*cgraph.Node{}

	if err := addNodes(graph, root, allNodes); err != nil {
		return clues.Stack(err)
	}

	// To keep from adding edges multiple times, track which nodes we've already
	// processed. This is required because there can be multiple paths to a node.
	visitedNodes := map[string]struct{}{}

	if err := addEdges(graph, root, allNodes, visitedNodes); err != nil {
		return clues.Stack(err)
	}

	// Go through and add edges between all nodes. The edge info will be based
	// on the Reason contained in the edge struct.
	return nil
}

func addNodes(
	graph *cgraph.Graph,
	node *repo.BackupNode,
	allNodes map[string]*cgraph.Node,
) error {
	if node == nil {
		return nil
	}

	if _, ok := allNodes[node.Label]; ok {
		return nil
	}

	// Need unique keys for nodes so use the backupID.
	n, err := graph.CreateNode(node.Label)
	if err != nil {
		return clues.Wrap(err, "creating node").With("backup_id", node.Label)
	}

	// Set tooltip info to have Reasons for backup and backup type.
	var toolTip string

	if node.Deleted {
		toolTip += "This backup was deleted, Reasons are a best guess!\n"
	}

	toolTip += "BackupID: " + node.Label + "\n"

	switch node.Type {
	case repo.MergeNode:
		toolTip += "Base Type: merge\n"

	case repo.AssistNode:
		toolTip += "Base Type: assist\n"
	}

	toolTip += fmt.Sprintf("Created At: %v\n", node.Created)

	var reasonStrings []string

	for _, reason := range node.Reasons {
		reasonStrings = append(
			reasonStrings,
			fmt.Sprintf("%s/%s", reason.Service(), reason.Category()),
		)
	}

	n.
		SetLabel(strings.Join(reasonStrings, "\n")).
		SetTooltip(toolTip).
		SetStyle(cgraph.FilledNodeStyle).
		SetFillColor("white")

	if node.Deleted {
		n.SetFillColor("indianred")
	}

	if node.Type == repo.AssistNode {
		n.SetFillColor("grey")
	}

	allNodes[node.Label] = n

	for _, child := range node.Children {
		if err := addNodes(graph, child.BackupNode, allNodes); err != nil {
			return clues.Stack(err)
		}
	}

	return nil
}

func addEdges(
	graph *cgraph.Graph,
	node *repo.BackupNode,
	allNodes map[string]*cgraph.Node,
	visitedNodes map[string]struct{},
) error {
	if node == nil {
		return nil
	}

	if _, ok := visitedNodes[node.Label]; ok {
		return nil
	}

	visitedNodes[node.Label] = struct{}{}

	n := allNodes[node.Label]

	for _, child := range node.Children {
		var edgeReasons []string

		for _, reason := range child.Reasons {
			edgeReasons = append(
				edgeReasons,
				fmt.Sprintf("%s/%s", reason.Service(), reason.Category()),
			)
		}

		edgeLabel := strings.Join(edgeReasons, ",\n")

		e, err := graph.CreateEdge(edgeLabel, n, allNodes[child.Label])
		if err != nil {
			return clues.Wrap(err, "adding edge").With(
				"parent", node.Label,
				"child", child.Label,
			)
		}

		e.SetDir(cgraph.ForwardDir).SetLabel(edgeLabel).SetTooltip(" ")

		if err := addEdges(graph, child.BackupNode, allNodes, visitedNodes); err != nil {
			return clues.Stack(err)
		}
	}

	return nil
}
