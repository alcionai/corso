package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/alcionai/clues"
	"github.com/minio/minio-go/v7"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/cmd/s3checker/pkg/s3"
	"github.com/alcionai/corso/src/internal/common/crash"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/storage"
)

// Matches other definitions of this const.
const defaultS3Endpoint = "s3.amazonaws.com"

type flags struct {
	bucket                string
	bucketPrefix          string
	prefixes              []string
	withDeleted           bool
	liveRetentionDuration time.Duration
	deadRetentionDuration time.Duration
	retentionMode         string
}

func checkerCommand() (*cobra.Command, error) {
	f := flags{}
	cmd := &cobra.Command{
		Use:   "check",
		Short: "Check S3 objects' retention properties",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handleCheckerCommand(cmd, args, f)
		},
	}

	fs := cmd.Flags()

	// AWS/global config.
	fs.StringVar(
		&f.bucket,
		"bucket",
		"",
		"Name of bucket to check")
	fs.StringVar(
		&f.bucketPrefix,
		"bucket-prefix",
		"",
		"Prefix to add to all object lookups")
	fs.StringSliceVar(
		&f.prefixes,
		"prefix",
		nil,
		"Set of object prefixes to check. Pass multiple times for multiple prefixes")

	// Live object config.
	fs.StringVar(
		&f.retentionMode,
		"retention-mode",
		"",
		"Retention mode to check for on live objects")
	fs.DurationVar(
		&f.liveRetentionDuration,
		"live-retention-duration",
		0,
		"Minimum amount of time from now that live objects should be locked for")

	// Dead object config.
	fs.BoolVar(
		&f.withDeleted,
		"with-non-current",
		false,
		"Whether to check non-current objects")
	fs.DurationVar(
		&f.deadRetentionDuration,
		"dead-retention-duration",
		0,
		"Maximum amount of time from now that dead objects should be locked for. "+
			"Can be negative if dead object lockss should have expired already")

	required := []string{
		"bucket",
		"prefix",
		"retention-mode",
		"live-retention-duration",
	}

	for _, req := range required {
		if err := cmd.MarkFlagRequired(req); err != nil {
			return nil, clues.Wrap(err, "setting flag "+req+" as required")
		}
	}

	return cmd, nil
}

func main() {
	cmd, err := checkerCommand()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	ls := logger.Settings{
		Level:  logger.LLDebug,
		Format: logger.LFText,
	}

	ctx, _ := logger.CtxOrSeed(context.Background(), ls)

	defer func() {
		if err := crash.Recovery(ctx, recover(), "s3Checker"); err != nil {
			logger.CtxErr(ctx, err).Error("panic in s3 checker")
		}

		logger.Flush(ctx)
	}()

	if err := cmd.ExecuteContext(ctx); err != nil {
		logger.Flush(ctx)
		os.Exit(1)
	}
}

func validateFlags(f flags) error {
	if f.liveRetentionDuration <= 0 {
		return clues.New("live object retention duration must be > 0")
	}

	if f.retentionMode != "GOVERNANCE" && f.retentionMode != "COMPLIANCE" {
		return clues.New("invalid retention mode")
	}

	return nil
}

func reportMissingPrefixes(
	objDescriptor string,
	wanted []string,
	got map[string]s3.ObjInfo,
) error {
	var err error

	for _, want := range wanted {
		if _, ok := got[want]; !ok {
			fmt.Printf("missing %s object for prefix %q\n", objDescriptor, want)

			err = clues.Stack(
				err,
				clues.New("missing "+objDescriptor+" object prefix \""+want+"\""))
		}
	}

	return err
}

func handleCheckerCommand(cmd *cobra.Command, args []string, f flags) error {
	if len(f.prefixes) == 0 {
		return nil
	}

	if err := validateFlags(f); err != nil {
		return clues.Stack(err)
	}

	cmd.SilenceUsage = true

	fmt.Printf("Checking objects with prefix(es) %v\n", f.prefixes)

	if err := config.InitFunc(cmd, args); err != nil {
		return clues.Wrap(err, "setting viper")
	}

	ctx := cmd.Context()

	// Scavenged from src/internal/kopia/s3/s3.go.
	overrides := map[string]string{
		storage.Bucket: f.bucket,
		storage.Prefix: f.bucketPrefix,
	}

	repoDetails, err := config.GetConfigRepoDetails(ctx, false, false, overrides)
	if err != nil {
		return clues.Wrap(err, "getting storage config")
	}

	cfg, err := repoDetails.Storage.S3Config()
	if err != nil {
		return clues.Wrap(err, "getting S3 config")
	}

	endpoint := defaultS3Endpoint
	if len(cfg.Endpoint) > 0 {
		endpoint = cfg.Endpoint
	}

	opts := &s3.Options{
		BucketName:      cfg.Bucket,
		Endpoint:        endpoint,
		Prefix:          cfg.Prefix,
		DoNotUseTLS:     cfg.DoNotUseTLS,
		DoNotVerifyTLS:  cfg.DoNotVerifyTLS,
		AccessKeyID:     cfg.AccessKey,
		SecretAccessKey: cfg.SecretKey,
		SessionToken:    cfg.SessionToken,
	}

	client, err := s3.New(opts)
	if err != nil {
		return clues.Wrap(err, "initializing S3 client")
	}

	live, dead, err := client.ListUntilAllFound(ctx, f.prefixes, f.withDeleted)
	if err != nil {
		return clues.Wrap(err, "getting objects to check")
	}

	// Reset error so we can return something at the end.
	err = nil

	if err2 := reportMissingPrefixes("live", f.prefixes, live); err2 != nil {
		err = clues.Stack(err, clues.New("some live objects missing"))
	}

	if f.withDeleted {
		// Only print here because it's possible there aren't dead objects for the
		// given prefix.
		//nolint:errcheck
		reportMissingPrefixes("dead", f.prefixes, dead)
	}

	now := time.Now()
	retentionMode := minio.RetentionMode(f.retentionMode)
	lowerBound := now.Add(f.liveRetentionDuration)
	upperBound := now.Add(f.deadRetentionDuration)

	liveErrs := checkObjsWithRetention(
		ctx,
		client,
		maps.Values(live),
		hasAtLeastRetention(retentionMode, lowerBound))

	deadErrs := checkObjsWithRetention(
		ctx,
		client,
		maps.Values(dead),
		hasAtMostRetention(upperBound))

	if len(liveErrs) > 0 {
		fmt.Printf("%d error(s) checking live object retention\n", len(liveErrs))

		for i, err := range liveErrs {
			fmt.Printf("\t%d: %s\n", i, err.Error())
		}

		err = clues.Stack(err, clues.New("live objects"))
	} else {
		fmt.Println("no errors for live objects")
	}

	if len(deadErrs) > 0 {
		fmt.Printf("%d error(s) checking dead object retention\n", len(deadErrs))

		for i, err := range deadErrs {
			fmt.Printf("\t%d: %s\n", i, err.Error())
		}

		err = clues.Stack(err, clues.New("dead objects"))
	} else {
		fmt.Println("no errors for dead objects")
	}

	return err
}

// hasAtLeastRetention takes a mode and time and returns a function that checks
// that objects have that retention mode and will expire at or after the given
// time.
func hasAtLeastRetention(
	wantedMode minio.RetentionMode,
	lowerBound time.Time,
) func(*minio.RetentionMode, *time.Time) error {
	return func(mode *minio.RetentionMode, expiry *time.Time) error {
		if mode == nil {
			return clues.New("nil retention mode")
		}

		if expiry == nil {
			return clues.New("nil retention expiry")
		}

		if wantedMode != *mode {
			return clues.New("unexpected retention mode " + string(*mode))
		}

		if expiry.Before(lowerBound) {
			return clues.New("unexpected retention expiry " + expiry.String())
		}

		return nil
	}
}

// hasAtMostRetention takes a time and returns a function that checks that
// objects have either no retention set or retention that expires at or before
// the given time.
func hasAtMostRetention(
	upperBound time.Time,
) func(*minio.RetentionMode, *time.Time) error {
	return func(mode *minio.RetentionMode, expiry *time.Time) error {
		// We turn expired retention into (nil, nil) for ease of writing checks.
		if mode == nil && expiry == nil {
			return nil
		}

		if mode == nil {
			return clues.New("nil retention mode")
		}

		if expiry == nil {
			return clues.New("nil retention expiry")
		}

		if upperBound.Before(*expiry) {
			return clues.New("unexpected retention expiry " + expiry.String())
		}

		return nil
	}
}

// checkObjsWithRetention takes a set of objs to check retention for and a
// function to use to check retention and returns a slice of errors, one for
// every item that either could not be fetched or that failed the retention
// check.
func checkObjsWithRetention(
	ctx context.Context,
	client *s3.Client,
	objs []s3.ObjInfo,
	checkFunc func(*minio.RetentionMode, *time.Time) error,
) []error {
	var errs []error

	for _, obj := range objs {
		mode, expiry, err := client.ObjectRetention(ctx, obj)
		// Locks that have expired start returning errors instead. Turn those
		// specific errors into (nil, nil) so that writing checks is a bit easier.
		if errors.Is(err, s3.ErrNoRetention) {
			mode = nil
			expiry = nil
			err = nil
		}

		if err != nil {
			errs = append(errs, clues.Stack(err))
			continue
		}

		if err := checkFunc(mode, expiry); err != nil {
			errs = append(errs, clues.Wrap(err, fmt.Sprintf(
				"checking object (key) %q (versionID) %q",
				obj.Key,
				obj.Version)))
		}
	}

	return errs
}
