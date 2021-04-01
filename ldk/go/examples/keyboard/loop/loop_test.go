package loop_test

import (
	"context"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/service"
	"github.com/open-olive/loop-development-kit/ldk/go/v2/whisper"
	"testing"

	"github.com/google/go-cmp/cmp"
	loop "github.com/open-olive/loop-development-kit/ldk/go/examples/keyboard/loop"
	ldk "github.com/open-olive/loop-development-kit/ldk/go/v2"
	ldktest "github.com/open-olive/loop-development-kit/ldk/go/v2/ldk-test"
)

func TestController(t *testing.T) {
	sidekick := &ldktest.Sidekick{
		KeyboardService: &ldktest.KeyboardService{
			ListenTextf: func(ctx context.Context, cb service.ListenTextHandler) error {
				cb("Some keyboard text", nil)

				return nil
			},
		},
		WhisperService: &ldktest.WhisperService{
			Markdownf: func(ctx context.Context, w *whisper.WhisperContentMarkdown) error {
				exp := "Text from the keyboard: Some keyboard text"
				if got := w.Markdown; !cmp.Equal(got, exp) {
					t.Errorf("unexpected markdown:\n%s\n", cmp.Diff(got, exp))
				}
				return nil
			},
		},
	}
	l := ldk.NewLogger("loop-example")
	c, err := loop.NewLoop(l)
	if err != nil {
		t.Fatal(err)
	}
	if err := c.LoopStart(sidekick); err != nil {
		t.Fatal(err)
	}
	if err := c.LoopStop(); err != nil {
		t.Fatal(err)
	}
}
