package ctxlog_test

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	"github.com/rs/zerolog"

	"github.com/akabos/ctxlog"
)

func ExampleWithNewLogger() {
	ctx := ctxlog.WithNewLogger(context.Background(), os.Stdout, zerolog.InfoLevel)

	ctxlog.Trace(ctx).Msg("trace")
	ctxlog.Debug(ctx).Msg("debug")
	ctxlog.Info(ctx).Msg("info")
	ctxlog.Warn(ctx).Msg("warn")
	ctxlog.Error(ctx).Msg("error")

	// Output:
	// {"level":"info","message":"info"}
	// {"level":"warn","message":"warn"}
	// {"level":"error","message":"error"}
}

func ExampleWithWriterWithLevel() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx = ctxlog.WithWriter(ctx, os.Stdout)
	ctx = ctxlog.WithLevel(ctx, zerolog.DebugLevel)

	ctxlog.Print(ctx, "message")

	// Output:
	// {"level":"debug","message":"message"}
}

func ExampleWithSampler() {
	ctx := ctxlog.WithNewLogger(context.Background(), os.Stdout, zerolog.InfoLevel)
	ctx = ctxlog.WithSampler(ctx, &zerolog.BasicSampler{N: 2})

	ctxlog.Log(ctx).Msg("message 1")
	ctxlog.Log(ctx).Msg("message 2")
	ctxlog.Log(ctx).Msg("message 3")

	// Output:
	// {"message":"message 1"}
	// {"message":"message 3"}
}

func BenchmarkZerolog(b *testing.B) {
	l := zerolog.New(ioutil.Discard).Level(zerolog.TraceLevel)
	b.RunParallel(func (b *testing.PB) {
		for b.Next() {
			l.Print("message")
		}
	})
}

func BenchmarkCtxlog(b *testing.B) {
	ctx := ctxlog.WithNewLogger(context.Background(), ioutil.Discard, zerolog.TraceLevel)
	b.RunParallel(func (b *testing.PB) {
		for b.Next() {
			ctxlog.Print(ctx, "message")
		}
	})
}
