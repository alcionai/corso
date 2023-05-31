package observe

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/alcionai/clues"
	"github.com/dustin/go-humanize"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"

	"github.com/alcionai/corso/src/pkg/logger"
)

const (
	hideProgressBarsFN   = "hide-progress"
	retainProgressBarsFN = "retain-progress"
	progressBarWidth     = 32
)

var (
	wg sync.WaitGroup
	// TODO: Revisit this being a global and make it a parameter to the progress methods
	// so that each bar can be initialized with different contexts if needed.
	contxt   context.Context
	writer   io.Writer
	progress *mpb.Progress
	cfg      *config
)

func init() {
	cfg = &config{}

	makeSpinFrames(progressBarWidth)
}

// adds the persistent boolean flag --hide-progress to the provided command.
// This is a hack for help displays.  Due to seeding the context, we also
// need to parse the configuration before we execute the command.
func AddProgressBarFlags(cmd *cobra.Command) {
	fs := cmd.PersistentFlags()
	fs.Bool(hideProgressBarsFN, false, "turn off the progress bar displays")
	fs.Bool(retainProgressBarsFN, false, "retain the progress bar displays after completion")
}

// Due to races between the lazy evaluation of flags in cobra and the need to init observer
// behavior in a ctx, these options get pre-processed manually here using pflags.  The canonical
// AddProgressBarFlag() ensures the flags are displayed as part of the help/usage output.
func PreloadFlags() *config {
	fs := pflag.NewFlagSet("seed-observer", pflag.ContinueOnError)
	fs.ParseErrorsWhitelist.UnknownFlags = true
	fs.Bool(hideProgressBarsFN, false, "turn off the progress bar displays")
	fs.Bool(retainProgressBarsFN, false, "retain the progress bar displays after completion")
	// prevents overriding the corso/cobra help processor
	fs.BoolP("help", "h", false, "")

	// parse the os args list to find the observer display flags
	if err := fs.Parse(os.Args[1:]); err != nil {
		return nil
	}

	// retrieve the user's preferred display
	// automatically defaults to "info"
	shouldHide, err := fs.GetBool(hideProgressBarsFN)
	if err != nil {
		return nil
	}

	// retrieve the user's preferred display
	// automatically defaults to "info"
	shouldAlwaysShow, err := fs.GetBool(retainProgressBarsFN)
	if err != nil {
		return nil
	}

	return &config{
		doNotDisplay:          shouldHide,
		keepBarsAfterComplete: shouldAlwaysShow,
	}
}

// ---------------------------------------------------------------------------
// configuration
// ---------------------------------------------------------------------------

// config handles observer configuration
type config struct {
	doNotDisplay          bool
	keepBarsAfterComplete bool
}

func (c config) hidden() bool {
	return c.doNotDisplay || writer == nil
}

// SeedWriter adds default writer to the observe package.
// Uses a noop writer until seeded.
func SeedWriter(ctx context.Context, w io.Writer, c *config) {
	writer = w
	contxt = ctx

	if contxt == nil {
		contxt = context.Background()
	}

	if c != nil {
		cfg = c
	}

	progress = mpb.NewWithContext(
		contxt,
		mpb.WithWidth(progressBarWidth),
		mpb.WithWaitGroup(&wg),
		mpb.WithOutput(writer))
}

// Complete blocks until the progress finishes writing out all data.
// Afterwards, the progress instance is reset.
func Complete() {
	if progress != nil {
		progress.Wait()
	}

	SeedWriter(contxt, writer, cfg)
}

const (
	ItemBackupMsg  = "Backing up item"
	ItemRestoreMsg = "Restoring item"
	ItemQueueMsg   = "Queuing items"
)

// Progress Updates

// Message is used to display a progress message
func Message(ctx context.Context, msgs ...any) {
	plainSl := make([]string, 0, len(msgs))
	loggableSl := make([]string, 0, len(msgs))

	for _, m := range msgs {
		plainSl = append(plainSl, plainString(m))
		loggableSl = append(loggableSl, fmt.Sprintf("%v", m))
	}

	plain := strings.Join(plainSl, " ")
	loggable := strings.Join(loggableSl, " ")

	logger.Ctx(ctx).Info(loggable)

	if cfg.hidden() {
		return
	}

	wg.Add(1)

	bar := progress.New(
		-1,
		mpb.NopStyle(),
		mpb.PrependDecorators(decor.Name(
			plain,
			decor.WC{
				W: len(plain) + 1,
				C: decor.DidentRight,
			})))

	// Complete the bar immediately
	bar.SetTotal(-1, true)

	waitAndCloseBar(bar, func() {})()
}

// MessageWithCompletion is used to display progress with a spinner
// that switches to "done" when the completion channel is signalled
func MessageWithCompletion(
	ctx context.Context,
	msg any,
) (chan<- struct{}, func()) {
	var (
		plain    = plainString(msg)
		loggable = fmt.Sprintf("%v", msg)
		log      = logger.Ctx(ctx)
		ch       = make(chan struct{}, 1)
	)

	log.Info(loggable)

	if cfg.hidden() {
		return ch, func() { log.Info("done - " + loggable) }
	}

	wg.Add(1)

	frames := []string{"∙∙∙", "●∙∙", "∙●∙", "∙∙●", "∙∙∙"}

	bar := progress.New(
		-1,
		mpb.SpinnerStyle(frames...).PositionLeft(),
		mpb.PrependDecorators(
			decor.Name(plain+":"),
			decor.Elapsed(decor.ET_STYLE_GO, decor.WC{W: 8})),
		mpb.BarFillerOnComplete("done"))

	go listen(
		ctx,
		ch,
		func() {
			bar.SetTotal(-1, true)
			bar.Abort(true)
		},
		func() {
			// We don't care whether the channel was signalled or closed
			// Use either one as an indication that the bar is done
			bar.SetTotal(-1, true)
		})

	wacb := waitAndCloseBar(bar, func() {
		log.Info("done - " + loggable)
	})

	return ch, wacb
}

// ---------------------------------------------------------------------------
// Progress for Known Quantities
// ---------------------------------------------------------------------------

// ItemProgress tracks the display of an item in a folder by counting the bytes
// read through the provided readcloser, up until the byte count matches
// the totalBytes.
func ItemProgress(
	ctx context.Context,
	rc io.ReadCloser,
	header string,
	iname any,
	totalBytes int64,
) (io.ReadCloser, func()) {
	plain := plainString(iname)
	log := logger.Ctx(ctx).With(
		"item", iname,
		"size", humanize.Bytes(uint64(totalBytes)))
	log.Debug(header)

	if cfg.hidden() || rc == nil || totalBytes == 0 {
		defer log.Debug("done - " + header)
		return rc, func() {}
	}

	wg.Add(1)

	barOpts := []mpb.BarOption{
		mpb.PrependDecorators(
			decor.Name(header, decor.WCSyncSpaceR),
			decor.Name(plain, decor.WCSyncSpaceR),
			decor.CountersKibiByte(" %.1f/%.1f ", decor.WC{W: 8}),
			decor.NewPercentage("%d ", decor.WC{W: 4})),
	}

	if !cfg.keepBarsAfterComplete {
		barOpts = append(barOpts, mpb.BarRemoveOnComplete())
	}

	bar := progress.New(totalBytes, mpb.NopStyle(), barOpts...)

	go waitAndCloseBar(bar, func() {
		// might be overly chatty, we can remove if needed.
		log.Debug("done - " + header)
	})()

	abort := func() {
		bar.SetTotal(-1, true)
		bar.Abort(true)
	}

	return bar.ProxyReader(rc), abort
}

// ProgressWithCount tracks the display of a bar that tracks the completion
// of the specified count.
// Each write to the provided channel counts as a single increment.
// The caller is expected to close the channel.
func ProgressWithCount(
	ctx context.Context,
	header string,
	msg any,
	count int64,
) (chan<- struct{}, func()) {
	var (
		plain    = plainString(msg)
		loggable = fmt.Sprintf("%s %v - %d", header, msg, count)
		log      = logger.Ctx(ctx)
		ch       = make(chan struct{})
	)

	log.Info(loggable)

	if cfg.hidden() {
		go listen(ctx, ch, nop, nop)
		return ch, func() { log.Info("done - " + loggable) }
	}

	wg.Add(1)

	barOpts := []mpb.BarOption{
		mpb.PrependDecorators(
			decor.Name(header, decor.WCSyncSpaceR),
			decor.Name(plain),
			decor.Counters(0, " %d/%d ")),
	}

	if !cfg.keepBarsAfterComplete {
		barOpts = append(barOpts, mpb.BarRemoveOnComplete())
	}

	bar := progress.New(count, mpb.NopStyle(), barOpts...)

	go listen(
		ctx,
		ch,
		func() { bar.Abort(true) },
		bar.Increment)

	wacb := waitAndCloseBar(bar, func() {
		log.Info("done - " + loggable)
	})

	return ch, wacb
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

// CollectionProgress tracks the display a spinner that idles while the collection
// incrementing the count of items handled.  Each write to the provided channel
// counts as a single increment.  The caller is expected to close the channel.
func CollectionProgress(
	ctx context.Context,
	category string,
	dirName any,
) (chan<- struct{}, func()) {
	var (
		counted int
		plain   = plainString(dirName)
		ch      = make(chan struct{})
		log     = logger.Ctx(ctx).With(
			"category", category,
			"dir", dirName)
		message = "Collecting Directory"
	)

	log.Info(message)

	incCount := func() {
		counted++
		// Log every 1000 items that are processed
		if counted%1000 == 0 {
			log.Infow("uploading", "count", counted)
		}
	}

	if cfg.hidden() || len(plain) == 0 {
		go listen(ctx, ch, nop, incCount)
		return ch, func() { log.Infow("done - "+message, "count", counted) }
	}

	wg.Add(1)

	barOpts := []mpb.BarOption{
		mpb.PrependDecorators(decor.Name(string(category))),
		mpb.AppendDecorators(
			decor.CurrentNoUnit("%d - ", decor.WCSyncSpace),
			decor.Name(plain),
		),
		mpb.BarFillerOnComplete(spinFrames[0]),
	}

	if !cfg.keepBarsAfterComplete {
		barOpts = append(barOpts, mpb.BarRemoveOnComplete())
	}

	bar := progress.New(
		-1, // -1 to indicate an unbounded count
		mpb.SpinnerStyle(spinFrames...),
		barOpts...)

	go listen(
		ctx,
		ch,
		func() { bar.SetTotal(-1, true) },
		func() {
			incCount()
			bar.Increment()
		})

	wacb := waitAndCloseBar(bar, func() {
		log.Infow("done - "+message, "count", counted)
	})

	return ch, wacb
}

func waitAndCloseBar(bar *mpb.Bar, log func()) func() {
	return func() {
		bar.Wait()
		wg.Done()

		if !bar.Aborted() {
			log()
		}
	}
}

// ---------------------------------------------------------------------------
// other funcs
// ---------------------------------------------------------------------------

var nop = func() {}

// listen handles reading, and exiting, from a channel.  It assumes the
// caller will run it inside a goroutine (ex: go listen(...)).
// On context timeout or channel close, the loop exits.
// onEnd() is called on both ctx.Done() and channel close.  onInc is
// called on every channel read except when closing.
func listen(ctx context.Context, ch <-chan struct{}, onEnd, onInc func()) {
	for {
		select {
		case <-ctx.Done():
			onEnd()
			return

		case _, ok := <-ch:
			if !ok {
				onEnd()
				return
			}

			onInc()
		}
	}
}

// ---------------------------------------------------------------------------
// Styling
// ---------------------------------------------------------------------------

const Bullet = "∙"

type bulletf struct {
	tmpl string
	vs   []any
}

func Bulletf(template string, vs ...any) bulletf {
	return bulletf{template, vs}
}

func (b bulletf) PlainString() string {
	ps := make([]any, 0, len(b.vs))
	for _, v := range b.vs {
		ps = append(ps, plainString(v))
	}

	return fmt.Sprintf("∙ "+b.tmpl, ps...)
}

func (b bulletf) String() string {
	return fmt.Sprintf("∙ "+b.tmpl, b.vs...)
}

// plainString attempts to cast v to a PlainStringer
// interface, and retrieve the un-altered value.  If
// v is not compliant with PlainStringer, returns the
// %v fmt of v.
//
// This should only be used to display the value in the
// observe progress bar.  Logged values should only use
// the fmt %v to ensure Concealers hide PII.
func plainString(v any) string {
	if ps, ok := v.(clues.PlainStringer); ok {
		return ps.PlainString()
	}

	return fmt.Sprintf("%v", v)
}
