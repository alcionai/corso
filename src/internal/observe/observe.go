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

func init() {
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
func PreloadFlags() config {
	fs := pflag.NewFlagSet("seed-observer", pflag.ContinueOnError)
	fs.ParseErrorsWhitelist.UnknownFlags = true
	fs.Bool(hideProgressBarsFN, false, "turn off the progress bar displays")
	fs.Bool(retainProgressBarsFN, false, "retain the progress bar displays after completion")
	// prevents overriding the corso/cobra help processor
	fs.BoolP("help", "h", false, "")

	// parse the os args list to find the observer display flags
	if err := fs.Parse(os.Args[1:]); err != nil {
		return config{}
	}

	// retrieve the user's preferred display
	// automatically defaults to "info"
	shouldHide, err := fs.GetBool(hideProgressBarsFN)
	if err != nil {
		return config{}
	}

	// retrieve the user's preferred display
	// automatically defaults to "info"
	shouldAlwaysShow, err := fs.GetBool(retainProgressBarsFN)
	if err != nil {
		return config{}
	}

	return config{
		doNotDisplay:          shouldHide,
		displayIsTerminal:     true,
		keepBarsAfterComplete: shouldAlwaysShow,
	}
}

// ---------------------------------------------------------------------------
// configuration
// ---------------------------------------------------------------------------

// config handles observer configuration
type config struct {
	// under certain conditions (ex: testing) we aren't outputting
	// to a terminal.  When this happens the observe bars need to be
	// given a specific optional value or they'll never flush the
	// writer.
	displayIsTerminal     bool
	doNotDisplay          bool
	keepBarsAfterComplete bool
}

type observerKey string

const ctxKey observerKey = "corsoObserver"

type observer struct {
	cfg config
	mp  *mpb.Progress
	w   io.Writer
	wg  *sync.WaitGroup
}

func (o observer) hidden() bool {
	return o.cfg.doNotDisplay || o.w == nil
}

func (o *observer) resetWriter(ctx context.Context) {
	opts := []mpb.ContainerOption{
		mpb.WithWidth(progressBarWidth),
		mpb.WithWaitGroup(o.wg),
		mpb.WithOutput(o.w),
	}

	// if o.cfg.displayIsTerminal {
	// 	opts = append(opts, mpb.WithManualRefresh())
	// }

	o.mp = mpb.NewWithContext(ctx, opts...)
}

// SeedObserver adds an observer to the context.  Any calls to observe
// funcs will retrieve the observer from the context.  If no observer
// is found in the context, the call no-ops.
func SeedObserver(ctx context.Context, w io.Writer, cfg config) context.Context {
	obs := &observer{
		w:   w,
		cfg: cfg,
		wg:  &sync.WaitGroup{},
	}

	obs.resetWriter(ctx)

	return setObserver(ctx, obs)
}

func setObserver(ctx context.Context, obs *observer) context.Context {
	return context.WithValue(ctx, ctxKey, obs)
}

func getObserver(ctx context.Context) *observer {
	o := ctx.Value(ctxKey)
	if o == nil {
		return &observer{cfg: config{doNotDisplay: true}}
	}

	return o.(*observer)
}

// Flush blocks until the progress finishes writing out all data.
// Afterwards, the progress instance is reset.
func Flush(ctx context.Context) {
	obs := getObserver(ctx)

	if obs.mp != nil {
		obs.mp.Wait()
	}

	obs.resetWriter(ctx)
}

const (
	ItemBackupMsg  = "Backing up item"
	ItemRestoreMsg = "Restoring item"
	ItemExportMsg  = "Exporting item"
	ItemQueueMsg   = "Queuing items"
)

// ---------------------------------------------------------------------------
// Progress Updates
// ---------------------------------------------------------------------------

// Message is used to display a progress message
func Message(ctx context.Context, msgs ...any) {
	var (
		obs        = getObserver(ctx)
		plainSl    = make([]string, 0, len(msgs))
		loggableSl = make([]string, 0, len(msgs))
	)

	for _, m := range msgs {
		plainSl = append(plainSl, plainString(m))
		loggableSl = append(loggableSl, fmt.Sprintf("%v", m))
	}

	plain := strings.Join(plainSl, " ")
	loggable := strings.Join(loggableSl, " ")

	logger.Ctx(ctx).Info(loggable)

	if obs.hidden() {
		return
	}

	obs.wg.Add(1)

	bar := obs.mp.New(
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
	waitAndCloseBar(ctx, bar, obs.wg, func() {})()
}

// MessageWithCompletion is used to display progress with a spinner
// that switches to "done" when the completion channel is signalled
func MessageWithCompletion(
	ctx context.Context,
	msg any,
) chan<- struct{} {
	var (
		obs      = getObserver(ctx)
		plain    = plainString(msg)
		loggable = fmt.Sprintf("%v", msg)
		log      = logger.Ctx(ctx)
		ch       = make(chan struct{}, 1)
	)

	log.Info(loggable)

	if obs.hidden() {
		defer log.Info("done - " + loggable)
		return ch
	}

	obs.wg.Add(1)

	frames := []string{"∙∙∙", "●∙∙", "∙●∙", "∙∙●", "∙∙∙"}

	bar := obs.mp.New(
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
		// callers should close the channel
		func() {})

	go waitAndCloseBar(ctx, bar, obs.wg, func() {
		log.Info("done - " + loggable)
	})()

	return ch
}

// ---------------------------------------------------------------------------
// Progress for Known Quantities
// ---------------------------------------------------------------------------

type autoCloser struct {
	rc     io.ReadCloser
	close  func()
	closed bool
}

func (ac *autoCloser) Read(p []byte) (n int, err error) {
	return ac.rc.Read(p)
}

func (ac *autoCloser) Close() error {
	if !ac.closed {
		ac.closed = true
		ac.close()
	}

	return ac.rc.Close()
}

// ItemProgress tracks the display of an item in a folder by counting the bytes
// read through the provided readcloser, up until the byte count matches
// the totalBytes.
// The progress bar will close automatically when the reader closes.  If an early
// close is needed due to abort or other issue, the returned func can be used.
func ItemProgress(
	ctx context.Context,
	rc io.ReadCloser,
	header string,
	iname any,
	totalBytes int64,
) (io.ReadCloser, func()) {
	var (
		obs   = getObserver(ctx)
		plain = plainString(iname)
		log   = logger.Ctx(ctx).With(
			"item", iname,
			"size", humanize.Bytes(uint64(totalBytes)))
	)

	log.Debug(header)

	if obs.hidden() || rc == nil {
		defer log.Debug("done - " + header)
		return rc, func() {}
	}

	obs.wg.Add(1)

	barOpts := []mpb.BarOption{
		mpb.PrependDecorators(
			decor.Name(header, decor.WCSyncSpaceR),
			decor.Name(plain, decor.WCSyncSpaceR),
			decor.CountersKibiByte(" %.1f/%.1f ", decor.WC{W: 8}),
			decor.NewPercentage("%d ", decor.WC{W: 4})),
	}

	if !obs.cfg.keepBarsAfterComplete {
		barOpts = append(barOpts, mpb.BarRemoveOnComplete())
	}

	bar := obs.mp.New(totalBytes, mpb.NopStyle(), barOpts...)

	go waitAndCloseBar(ctx, bar, obs.wg, func() {
		// might be overly chatty, we can remove if needed.
		log.Debug("done - " + header)
	})()

	closer := &autoCloser{rc: bar.ProxyReader(rc)}

	closer.close = func() {
		closer.closed = true
		bar.SetTotal(-1, true)
		bar.Abort(true)
	}

	return closer, closer.close
}

// ItemSpinner is similar to ItemProgress, but for use in cases where
// we don't know the file size but want to show progress.
func ItemSpinner(
	ctx context.Context,
	rc io.ReadCloser,
	header string,
	iname any,
) (io.ReadCloser, func()) {
	var (
		obs   = getObserver(ctx)
		plain = plainString(iname)
		log   = logger.Ctx(ctx).With("item", iname)
	)

	log.Debug(header)

	if obs.hidden() || rc == nil {
		defer log.Debug("done - " + header)
		return rc, func() {}
	}

	obs.wg.Add(1)

	barOpts := []mpb.BarOption{
		mpb.PrependDecorators(
			decor.Name(header, decor.WCSyncSpaceR),
			decor.Name(plain, decor.WCSyncSpaceR),
			decor.CurrentKibiByte(" %.1f", decor.WC{W: 8})),
	}

	if !obs.cfg.keepBarsAfterComplete {
		barOpts = append(barOpts, mpb.BarRemoveOnComplete())
	}

	bar := obs.mp.New(-1, mpb.NopStyle(), barOpts...)

	go waitAndCloseBar(ctx, bar, obs.wg, func() {
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
) chan<- struct{} {
	var (
		obs      = getObserver(ctx)
		plain    = plainString(msg)
		loggable = fmt.Sprintf("%s %v - %d", header, msg, count)
		log      = logger.Ctx(ctx)
		ch       = make(chan struct{})
	)

	log.Info(loggable)

	if obs.hidden() {
		go listen(ctx, ch, nop, nop)

		defer log.Info("done - " + loggable)

		return ch
	}

	obs.wg.Add(1)

	barOpts := []mpb.BarOption{
		mpb.PrependDecorators(
			decor.Name(header, decor.WCSyncSpaceR),
			decor.Name(plain),
			decor.Counters(0, " %d/%d ")),
	}

	if !obs.cfg.keepBarsAfterComplete {
		barOpts = append(barOpts, mpb.BarRemoveOnComplete())
	}

	bar := obs.mp.New(count, mpb.NopStyle(), barOpts...)

	go listen(
		ctx,
		ch,
		func() {
			bar.Abort(true)
		},
		bar.Increment)

	go waitAndCloseBar(ctx, bar, obs.wg, func() {
		log.Info("done - " + loggable)
	})()

	return ch
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
) chan<- struct{} {
	var (
		obs     = getObserver(ctx)
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

	if obs.hidden() || len(plain) == 0 {
		go listen(ctx, ch, nop, incCount)

		defer log.Infow("done - "+message, "count", counted)

		return ch
	}

	obs.wg.Add(1)

	barOpts := []mpb.BarOption{
		mpb.PrependDecorators(decor.Name(string(category))),
		mpb.AppendDecorators(
			decor.CurrentNoUnit("%d - ", decor.WCSyncSpace),
			decor.Name(plain),
		),
		mpb.BarFillerOnComplete(spinFrames[0]),
	}

	if !obs.cfg.keepBarsAfterComplete {
		barOpts = append(barOpts, mpb.BarRemoveOnComplete())
	}

	bar := obs.mp.New(
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

	go waitAndCloseBar(ctx, bar, obs.wg, func() {
		log.Infow("done - "+message, "count", counted)
	})()

	return ch
}

func waitAndCloseBar(ctx context.Context, bar *mpb.Bar, wg *sync.WaitGroup, log func()) func() {
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
	if c, ok := v.(clues.Concealer); ok {
		return c.PlainString()
	}

	return fmt.Sprintf("%v", v)
}
