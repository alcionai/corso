package repo

import (
	"strconv"
	"strings"

	"github.com/alcionai/clues"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cli/flags"
	. "github.com/alcionai/corso/src/cli/print"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/events"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/storage"
)

// s3 bucket info from flags
var (
	succeedIfExists bool
	bucket          string
	endpoint        string
	prefix          string
	doNotUseTLS     bool
	doNotVerifyTLS  bool
)

// s3 bucket flags
const (
	succeedIfExistsFN = "succeedIfExists"
	bucketFN          = "bucket"
	endpointFN        = "endpoint"
	prefixFN          = "prefix"
	doNotUseTLSFN     = "disable-tls"
	doNotVerifyTLSFN  = "disable-tls-verification"
)

// called by repo.go to map subcommands to provider-specific handling.
func addS3Commands(cmd *cobra.Command) *cobra.Command {
	var (
		c  *cobra.Command
		fs *pflag.FlagSet
	)

	switch cmd.Use {
	case initCommand:
		init := s3InitCmd()
		flags.AddRetentionConfigFlags(init)
		c, fs = utils.AddCommand(cmd, init)

	case connectCommand:
		c, fs = utils.AddCommand(cmd, s3ConnectCmd())
	}

	c.Use = c.Use + " " + s3ProviderCommandUseSuffix
	c.SetUsageTemplate(cmd.UsageTemplate())

	flags.AddAWSCredsFlags(c)
	flags.AddAzureCredsFlags(c)
	flags.AddCorsoPassphaseFlags(c)

	// Flags addition ordering should follow the order we want them to appear in help and docs:
	// More generic and more frequently used flags take precedence.
	fs.StringVar(&bucket, bucketFN, "", "Name of S3 bucket for repo. (required)")
	fs.StringVar(&prefix, prefixFN, "", "Repo prefix within bucket.")
	fs.StringVar(&endpoint, endpointFN, "", "S3 service endpoint.")
	fs.BoolVar(&doNotUseTLS, doNotUseTLSFN, false, "Disable TLS (HTTPS)")
	fs.BoolVar(&doNotVerifyTLS, doNotVerifyTLSFN, false, "Disable TLS (HTTPS) certificate verification.")

	// In general, we don't want to expose this flag to users and have them mistake it
	// for a broad-scale idempotency solution.  We can un-hide it later the need arises.
	fs.BoolVar(&succeedIfExists, "succeed-if-exists", false, "Exit with success if the repo has already been initialized.")
	cobra.CheckErr(fs.MarkHidden("succeed-if-exists"))

	return c
}

const (
	s3ProviderCommand          = "s3"
	s3ProviderCommandUseSuffix = "--bucket <bucket>"
)

const (
	s3ProviderCommandInitExamples = `# Create a new Corso repo in AWS S3 bucket named "my-bucket"
corso repo init s3 --bucket my-bucket

# Create a new Corso repo in AWS S3 bucket named "my-bucket" using a prefix
corso repo init s3 --bucket my-bucket --prefix my-prefix

# Create a new Corso repo in an S3 compliant storage provider
corso repo init s3 --bucket my-bucket --endpoint my-s3-server-endpoint`

	s3ProviderCommandConnectExamples = `# Connect to a Corso repo in AWS S3 bucket named "my-bucket"
corso repo connect s3 --bucket my-bucket

# Connect to a Corso repo in AWS S3 bucket named "my-bucket" using a prefix
corso repo connect s3 --bucket my-bucket --prefix my-prefix

# Connect to a Corso repo in an S3 compliant storage provider
corso repo connect s3 --bucket my-bucket --endpoint my-s3-server-endpoint`
)

// ---------------------------------------------------------------------------------------------------------
// Init
// ---------------------------------------------------------------------------------------------------------

// `corso repo init s3 [<flag>...]`
func s3InitCmd() *cobra.Command {
	return &cobra.Command{
		Use:     s3ProviderCommand,
		Short:   "Initialize a S3 repository",
		Long:    `Bootstraps a new S3 repository and connects it to your m365 account.`,
		RunE:    initS3Cmd,
		Args:    cobra.NoArgs,
		Example: s3ProviderCommandInitExamples,
	}
}

// initializes a s3 repo.
func initS3Cmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	// s3 values from flags
	s3Override := S3Overrides(cmd)

	// Need to send provider here
	cfg, err := config.GetConfigRepoDetails(ctx, true, false, s3Override)
	if err != nil {
		return Only(ctx, err)
	}

	opt := utils.ControlWithConfig(cfg)

	retentionOpts, err := utils.MakeRetentionOpts(cmd)
	if err != nil {
		return Only(ctx, err)
	}

	// SendStartCorsoEvent uses distict ID as tenant ID because repoID is still not generated
	utils.SendStartCorsoEvent(
		ctx,
		cfg.Storage,
		cfg.Account.ID(),
		map[string]any{"command": "init repo"},
		cfg.Account.ID(),
		opt)

	storageCfg, err := config.NewStorageConfigFrom(cfg.Storage)
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Retrieving s3 configuration"))
	}

	s3Cfg, ok := storageCfg.(config.S3Config)
	if !ok {
		return Only(ctx, clues.New("Casting storage config to S3Config"))
	}

	// TODO: move this to cfg validate
	if strings.HasPrefix(s3Cfg.Endpoint, "http://") || strings.HasPrefix(s3Cfg.Endpoint, "https://") {
		invalidEndpointErr := "endpoint doesn't support specifying protocol. " +
			"pass --disable-tls flag to use http:// instead of default https://"

		return Only(ctx, clues.New(invalidEndpointErr))
	}

	m365, err := cfg.Account.M365Config()
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to parse m365 account config"))
	}

	r, err := repository.Initialize(
		ctx,
		cfg.Account,
		cfg.Storage,
		opt,
		retentionOpts)
	if err != nil {
		if succeedIfExists && errors.Is(err, repository.ErrorRepoAlreadyExists) {
			return nil
		}

		return Only(ctx, clues.Wrap(err, "Failed to initialize a new S3 repository"))
	}

	defer utils.CloseRepo(ctx, r)

	Infof(ctx, "Initialized a S3 repository within bucket %s.", s3Cfg.Bucket)

	if err = config.WriteRepoConfig(ctx, s3Cfg, m365, opt.Repo, r.GetID()); err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to write repository configuration"))
	}

	return nil
}

// ---------------------------------------------------------------------------------------------------------
// Connect
// ---------------------------------------------------------------------------------------------------------

// `corso repo connect s3 [<flag>...]`
func s3ConnectCmd() *cobra.Command {
	return &cobra.Command{
		Use:     s3ProviderCommand,
		Short:   "Connect to a S3 repository",
		Long:    `Ensures a connection to an existing S3 repository.`,
		RunE:    connectS3Cmd,
		Args:    cobra.NoArgs,
		Example: s3ProviderCommandConnectExamples,
	}
}

// connects to an existing s3 repo.
func connectS3Cmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	// s3 values from flags
	s3Override := S3Overrides(cmd)

	// Send provider here
	cfg, err := config.GetConfigRepoDetails(ctx, true, true, s3Override)
	if err != nil {
		return Only(ctx, err)
	}

	repoID := cfg.RepoID
	if len(repoID) == 0 {
		repoID = events.RepoIDNotFound
	}

	storageCfg, err := config.NewStorageConfigFrom(cfg.Storage)
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Retrieving s3 configuration"))
	}

	s3Cfg, ok := storageCfg.(config.S3Config)
	if !ok {
		return Only(ctx, clues.New("Casting storage config to S3Config"))
	}

	m365, err := cfg.Account.M365Config()
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to parse m365 account config"))
	}

	if strings.HasPrefix(s3Cfg.Endpoint, "http://") || strings.HasPrefix(s3Cfg.Endpoint, "https://") {
		invalidEndpointErr := "endpoint doesn't support specifying protocol. " +
			"pass --disable-tls flag to use http:// instead of default https://"

		return Only(ctx, clues.New(invalidEndpointErr))
	}

	opts := utils.ControlWithConfig(cfg)

	r, err := repository.ConnectAndSendConnectEvent(
		ctx,
		cfg.Account,
		cfg.Storage,
		repoID,
		opts)
	if err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to connect to the S3 repository"))
	}

	defer utils.CloseRepo(ctx, r)

	Infof(ctx, "Connected to S3 bucket %s.", s3Cfg.Bucket)

	if err = config.WriteRepoConfig(ctx, s3Cfg, m365, opts.Repo, r.GetID()); err != nil {
		return Only(ctx, clues.Wrap(err, "Failed to write repository configuration"))
	}

	return nil
}

func S3Overrides(cmd *cobra.Command) map[string]string {
	fs := flags.GetPopulatedFlags(cmd)
	return PopulateS3Flags(fs)
}

func PopulateS3Flags(flagset flags.PopulatedFlags) map[string]string {
	s3Overrides := make(map[string]string)
	s3Overrides[config.AccountProviderTypeKey] = account.ProviderM365.String()
	s3Overrides[config.StorageProviderTypeKey] = storage.ProviderS3.String()

	if _, ok := flagset[flags.AWSAccessKeyFN]; ok {
		s3Overrides[credentials.AWSAccessKeyID] = flags.AWSAccessKeyFV
	}

	if _, ok := flagset[flags.AWSSecretAccessKeyFN]; ok {
		s3Overrides[credentials.AWSSecretAccessKey] = flags.AWSSecretAccessKeyFV
	}

	if _, ok := flagset[flags.AWSSessionTokenFN]; ok {
		s3Overrides[credentials.AWSSessionToken] = flags.AWSSessionTokenFV
	}

	if _, ok := flagset[bucketFN]; ok {
		s3Overrides[config.Bucket] = bucket
	}

	if _, ok := flagset[prefixFN]; ok {
		s3Overrides[config.Prefix] = prefix
	}

	if _, ok := flagset[doNotUseTLSFN]; ok {
		s3Overrides[config.DoNotUseTLS] = strconv.FormatBool(doNotUseTLS)
	}

	if _, ok := flagset[doNotVerifyTLSFN]; ok {
		s3Overrides[config.DoNotVerifyTLS] = strconv.FormatBool(doNotVerifyTLS)
	}

	if _, ok := flagset[endpointFN]; ok {
		s3Overrides[config.Endpoint] = endpoint
	}

	return s3Overrides
}
