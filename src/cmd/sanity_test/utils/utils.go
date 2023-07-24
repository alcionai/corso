package utils

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/pkg/logger"
)

func Assert(
	ctx context.Context,
	passes func() bool,
	header string,
	expect, current any,
) {
	if passes() {
		return
	}

	header = "Error: " + header
	expected := fmt.Sprintf("* Expected: %+v", expect)
	got := fmt.Sprintf("* Current: %+v", current)

	logger.Ctx(ctx).Info(strings.Join([]string{header, expected, got}, " "))

	fmt.Println(header)
	fmt.Println(expected)
	fmt.Println(got)

	os.Exit(1)
}

func Fatal(ctx context.Context, msg string, err error) {
	logger.CtxErr(ctx, err).Error("test failure: " + msg)
	fmt.Println(msg+": ", err)
	os.Exit(1)
}

func MustGetTimeFromName(ctx context.Context, name string) (time.Time, bool) {
	t, err := dttm.ExtractTime(name)
	if err != nil && !errors.Is(err, dttm.ErrNoTimeString) {
		Fatal(ctx, "extracting time from name: "+name, err)
	}

	return t, !errors.Is(err, dttm.ErrNoTimeString)
}

func IsWithinTimeBound(ctx context.Context, bound, check time.Time, hasTime bool) bool {
	if hasTime {
		if bound.Before(check) {
			logger.Ctx(ctx).
				With("boundary_time", bound, "check_time", check).
				Info("skipping restore folder: not older than time bound")

			return false
		}
	}

	return true
}

func FilterSlice(sl []string, remove string) []string {
	r := []string{}

	for _, s := range sl {
		if !strings.EqualFold(s, remove) {
			r = append(r, s)
		}
	}

	return r
}

func LogAndPrint(ctx context.Context, tmpl string, vs ...any) {
	logger.Ctx(ctx).Infof(tmpl, vs...)
	fmt.Printf(tmpl+"\n", vs...)
}
