package observe

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
)

const (
	noProgressBarsFN = "no-progress-bars"
	progressBarWidth = 32
)

var (
	wg       sync.WaitGroup
	con      context.Context
	writer   io.Writer
	progress *mpb.Progress
	cfg      *config
)

func init() {
	makeSpinFrames(progressBarWidth)
}

// adds the persistent boolean flag --no-progress-bars to the provided command.
// This is a hack for help displays.  Due to seeding the context, we also
// need to parse the configuration before we execute the command.
func AddProgressBarFlags(parent *cobra.Command) {
	fs := parent.PersistentFlags()
	fs.Bool(noProgressBarsFN, false, "turn off the progress bar displays")
}

// Due to races between the lazy evaluation of flags in cobra and the need to init observer
// behavior in a ctx, these options get pre-processed manually here using pflags.  The canonical
// AddProgressBarFlag() ensures the flags are displayed as part of the help/usage output.
func PreloadFlags() bool {
	fs := pflag.NewFlagSet("seed-observer", pflag.ContinueOnError)
	fs.ParseErrorsWhitelist.UnknownFlags = true
	fs.Bool(noProgressBarsFN, false, "turn off the progress bar displays")
	// prevents overriding the corso/cobra help processor
	fs.BoolP("help", "h", false, "")

	// parse the os args list to find the log level flag
	if err := fs.Parse(os.Args[1:]); err != nil {
		return false
	}

	// retrieve the user's preferred display
	// automatically defaults to "info"
	shouldHide, err := fs.GetBool(noProgressBarsFN)
	if err != nil {
		return false
	}

	return shouldHide
}

// ---------------------------------------------------------------------------
// configuration
// ---------------------------------------------------------------------------

// config handles observer configuration
type config struct {
	doNotDisplay bool
}

// SeedWriter adds default writer to the observe package.
// Uses a noop writer until seeded.
func SeedWriter(ctx context.Context, w io.Writer, hide bool) {
	writer = w
	con = ctx

	if con == nil {
		con = context.Background()
	}

	if cfg == nil {
		cfg = &config{
			doNotDisplay: hide,
		}
	}

	progress = mpb.NewWithContext(
		con,
		mpb.WithWidth(progressBarWidth),
		mpb.WithWaitGroup(&wg),
		mpb.WithOutput(writer),
	)
}

// Complete blocks until the progress finishes writing out all data.
// Afterwards, the progress instance is reset.
func Complete() {
	if progress != nil {
		progress.Wait()
	}

	SeedWriter(con, writer, false)
}

// ---------------------------------------------------------------------------
// Progress for Known Quantities
// ---------------------------------------------------------------------------

// ItemProgress tracks the display of an item by counting the bytes
// read through the provided readcloser, up until the byte count matches
// the totalBytes.
func ItemProgress(rc io.ReadCloser, iname string, totalBytes int64) (io.ReadCloser, func()) {
	if cfg.doNotDisplay || writer == nil || rc == nil || totalBytes == 0 {
		return rc, func() {}
	}

	wg.Add(1)

	bar := progress.AddBar(
		totalBytes,
		mpb.BarFillerOnComplete(""),
		mpb.BarRemoveOnComplete(),
		mpb.PrependDecorators(
			decor.OnComplete(decor.NewPercentage("%d", decor.WC{W: 4}), ""),
			decor.OnComplete(decor.TotalKiloByte("%.1f", decor.WCSyncSpace), ""),
		),
		mpb.AppendDecorators(
			decor.OnComplete(decor.Name(iname), ""),
		),
	)

	return bar.ProxyReader(rc), waitAndCloseBar(bar)
}

// ---------------------------------------------------------------------------
// Progress for Unknown Quantities
// ---------------------------------------------------------------------------

var spinFrames []string

// The bar width is set to a static 32 characters.  The default spinner is only
// one char wide, which puts a lot of white space between it and the useful text.
// This builds a custom spinner animation to fill up that whitespace for a cleaner
// display.
func makeSpinFrames(barWidth int) {
	s, l := rune('∙'), rune('●')

	line := []rune{}
	for i := 0; i < barWidth; i++ {
		line = append(line, s)
	}

	sl := make([]string, 0, barWidth+1)
	sl = append(sl, string(line))

	for i := 1; i < barWidth; i++ {
		l2 := make([]rune, len(line))
		copy(l2, line)
		l2[i] = l

		sl = append(sl, string(l2))
	}

	spinFrames = sl
}

// ItemProgress tracks the display a spinner that idles while the collection
// incrementing the count of items handled.  Each write to the provided channel
// counts as a single increment.  The caller is expected to close the channel.
func CollectionProgress(user, category, dirName string) (chan<- struct{}, func()) {
	if cfg.doNotDisplay || writer == nil || len(user) == 0 || len(dirName) == 0 {
		ch := make(chan struct{})

		go func(ci <-chan struct{}) {
			for {
				_, ok := <-ci
				if !ok {
					return
				}
			}
		}(ch)

		return ch, func() {}
	}

	wg.Add(1)

	bar := progress.New(
		-1, // -1 to indicate an unbounded count
		mpb.SpinnerStyle(spinFrames...),
		mpb.BarFillerOnComplete(""),
		mpb.BarRemoveOnComplete(),
		mpb.PrependDecorators(
			decor.OnComplete(decor.Name(category), ""),
		),
		mpb.AppendDecorators(
			decor.OnComplete(decor.CurrentNoUnit("%d - ", decor.WCSyncSpace), ""),
			decor.OnComplete(
				decor.Name(fmt.Sprintf("%s - %s", user, dirName)),
				""),
		),
	)

	ch := make(chan struct{})

	go func(ci <-chan struct{}) {
		for {
			select {
			case <-con.Done():
				bar.SetTotal(-1, true)
				return

			case _, ok := <-ci:
				if !ok {
					bar.SetTotal(-1, true)
					return
				}

				bar.Increment()
			}
		}
	}(ch)

	return ch, waitAndCloseBar(bar)
}

func waitAndCloseBar(bar *mpb.Bar) func() {
	return func() {
		bar.Wait()
		wg.Done()
	}
}
