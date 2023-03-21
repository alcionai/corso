package observe

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

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

// styling
const bullet = "∙"

const Bullet = Safe(bullet)

var (
	wg sync.WaitGroup
	// TODO: Revisit this being a global nd make it a parameter to the progress methods
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
		mpb.WithOutput(writer),
	)
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
func Message(ctx context.Context, msgs ...cleanable) {
	var (
		cleaned = make([]string, len(msgs))
		msg     = make([]string, len(msgs))
	)

	for i := range msgs {
		cleaned[i] = msgs[i].clean()
		msg[i] = msgs[i].String()
	}

	logger.Ctx(ctx).Info(strings.Join(cleaned, " "))
	message := strings.Join(msg, " ")

	if cfg.hidden() {
		return
	}

	wg.Add(1)

	bar := progress.New(
		-1,
		mpb.NopStyle(),
		mpb.PrependDecorators(
			decor.Name(message, decor.WC{W: len(message) + 1, C: decor.DidentRight}),
		),
	)

	// Complete the bar immediately
	bar.SetTotal(-1, true)

	waitAndCloseBar(bar, func() {})()
}

// MessageWithCompletion is used to display progress with a spinner
// that switches to "done" when the completion channel is signalled
func MessageWithCompletion(
	ctx context.Context,
	msg cleanable,
) (chan<- struct{}, func()) {
	var (
		clean   = msg.clean()
		message = msg.String()
		log     = logger.Ctx(ctx)
		ch      = make(chan struct{}, 1)
	)

	log.Info(clean)

	if cfg.hidden() {
		return ch, func() { log.Info("done - " + clean) }
	}

	wg.Add(1)

	frames := []string{"∙∙∙", "●∙∙", "∙●∙", "∙∙●", "∙∙∙"}

	bar := progress.New(
		-1,
		mpb.SpinnerStyle(frames...).PositionLeft(),
		mpb.PrependDecorators(
			decor.Name(message+":"),
			decor.Elapsed(decor.ET_STYLE_GO, decor.WC{W: 8}),
		),
		mpb.BarFillerOnComplete("done"),
	)

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
		log.Info("done - " + clean)
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
	iname cleanable,
	totalBytes int64,
) (io.ReadCloser, func()) {
	log := logger.Ctx(ctx).With(
		"item", iname.clean(),
		"size", humanize.Bytes(uint64(totalBytes)))
	log.Debug(header)

	if cfg.hidden() || rc == nil || totalBytes == 0 {
		return rc, func() { log.Debug("done - " + header) }
	}

	wg.Add(1)

	barOpts := []mpb.BarOption{
		mpb.PrependDecorators(
			decor.Name(header, decor.WCSyncSpaceR),
			decor.Name(iname.String(), decor.WCSyncSpaceR),
			decor.CountersKibiByte(" %.1f/%.1f ", decor.WC{W: 8}),
			decor.NewPercentage("%d ", decor.WC{W: 4}),
		),
	}

	if !cfg.keepBarsAfterComplete {
		barOpts = append(barOpts, mpb.BarRemoveOnComplete())
	}

	bar := progress.New(totalBytes, mpb.NopStyle(), barOpts...)

	wacb := waitAndCloseBar(bar, func() {
		// might be overly chatty, we can remove if needed.
		log.Debug("done - " + header)
	})

	return bar.ProxyReader(rc), wacb
}

// ProgressWithCount tracks the display of a bar that tracks the completion
// of the specified count.
// Each write to the provided channel counts as a single increment.
// The caller is expected to close the channel.
func ProgressWithCount(
	ctx context.Context,
	header string,
	message cleanable,
	count int64,
) (chan<- struct{}, func()) {
	var (
		log  = logger.Ctx(ctx)
		lmsg = fmt.Sprintf("%s %s - %d", header, message.clean(), count)
		ch   = make(chan struct{})
	)

	log.Info(lmsg)

	if cfg.hidden() {
		go listen(ctx, ch, nop, nop)
		return ch, func() { log.Info("done - " + lmsg) }
	}

	wg.Add(1)

	barOpts := []mpb.BarOption{
		mpb.PrependDecorators(
			decor.Name(header, decor.WCSyncSpaceR),
			decor.Name(message.String()),
			decor.Counters(0, " %d/%d "),
		),
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
		log.Info("done - " + lmsg)
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
	dirName cleanable,
) (chan<- struct{}, func()) {
	var (
		counted int
		ch      = make(chan struct{})
		log     = logger.Ctx(ctx).With(
			"category", category,
			"dir", dirName.clean())
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

	if cfg.hidden() || len(dirName.String()) == 0 {
		go listen(ctx, ch, nop, incCount)
		return ch, func() { log.Infow("done - "+message, "count", counted) }
	}

	wg.Add(1)

	barOpts := []mpb.BarOption{
		mpb.PrependDecorators(decor.Name(string(category))),
		mpb.AppendDecorators(
			decor.CurrentNoUnit("%d - ", decor.WCSyncSpace),
			decor.Name(dirName.String()),
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
		log()
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
// PII redaction
// ---------------------------------------------------------------------------

type cleanable interface {
	clean() string
	String() string
}

type PII string

func (p PII) clean() string {
	return "***"
}

func (p PII) String() string {
	return string(p)
}

type Safe string

func (s Safe) clean() string {
	return string(s)
}

func (s Safe) String() string {
	return string(s)
}

type bulletPII struct {
	tmpl string
	vars []cleanable
}

func Bulletf(template string, vs ...cleanable) bulletPII {
	return bulletPII{
		tmpl: "∙ " + template,
		vars: vs,
	}
}

func (b bulletPII) clean() string {
	vs := make([]any, 0, len(b.vars))

	for _, v := range b.vars {
		vs = append(vs, v.clean())
	}

	return fmt.Sprintf(b.tmpl, vs...)
}

func (b bulletPII) String() string {
	vs := make([]any, 0, len(b.vars))

	for _, v := range b.vars {
		vs = append(vs, v.String())
	}

	return fmt.Sprintf(b.tmpl, vs...)
}
