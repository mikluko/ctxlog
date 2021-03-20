package ctxlog

import (
	"context"
	"io"
	"io/ioutil"

	"github.com/rs/zerolog"
)

type ctxKey struct {}

// Logger returns logger associated with the context. Returns a disabled one if no logger in the context.
func Logger(ctx context.Context) zerolog.Logger {
	if l, ok := ctx.Value(ctxKey{}).(zerolog.Logger); ok {
		return l
	}
	return zerolog.New(ioutil.Discard).Level(zerolog.Disabled)
}

// WithLogger updates context with the provided logger
func WithLogger(ctx context.Context, l zerolog.Logger) context.Context {
	return context.WithValue(ctx, ctxKey{}, l)
}

// WithNewLogger updates context with the logger initialized with provided writer and level
func WithNewLogger(ctx context.Context, w io.Writer, lvl zerolog.Level) context.Context {
	return WithLogger(ctx, zerolog.New(w).Level(lvl))
}

func WithLevel(ctx context.Context, lvl zerolog.Level) context.Context {
	l := Logger(ctx).Level(lvl)
	return WithLogger(ctx, l)
}

func WithWriter(ctx context.Context, w io.Writer) context.Context {
	l := Logger(ctx).Output(w)
	return WithLogger(ctx, l)
}

func WithSampler(ctx context.Context, s zerolog.Sampler) context.Context {
	l := Logger(ctx).Sample(s)
	return WithLogger(ctx, l)
}

func Update(ctx context.Context, f func(zerolog.Context) zerolog.Context) context.Context {
	return WithLogger(ctx, f(Logger(ctx).With()).Logger())
}

func Trace(ctx context.Context) *zerolog.Event {
	l := Logger(ctx)
	return l.Trace()
}

func Debug(ctx context.Context) *zerolog.Event {
	l := Logger(ctx)
	return l.Debug()
}

func Info(ctx context.Context) *zerolog.Event {
	l := Logger(ctx)
	return l.Info()
}

func Warn(ctx context.Context) *zerolog.Event {
	l := Logger(ctx)
	return l.Warn()
}

func Error(ctx context.Context) *zerolog.Event {
	l := Logger(ctx)
	return l.Error()
}

func Fatal(ctx context.Context) *zerolog.Event {
	l := Logger(ctx)
	return l.Fatal()
}

func Log(ctx context.Context) *zerolog.Event {
	l := Logger(ctx)
	return l.Log()
}

func Level(ctx context.Context, lvl zerolog.Level) *zerolog.Event {
	l := Logger(ctx)
	return l.WithLevel(lvl)
}

func Print(ctx context.Context, v interface{}) {
	l := Logger(ctx)
	l.Print(v)
}

func Printf(ctx context.Context, format string, v interface{}) {
	l := Logger(ctx)
	l.Printf(format, v)
}