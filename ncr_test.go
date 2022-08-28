package thread_test

import (
	"testing"

	"github.com/aethiopicuschan/thread"
)

func TestDecodeNCR(t *testing.T) {
	str := thread.DecodeNCR("家康の関東移封&#9876;左遷どころか大当たりだった&#127919;")
	if str != "家康の関東移封⚔左遷どころか大当たりだった🎯" {
		t.Errorf("want \"家康の関東移封⚔左遷どころか大当たりだった🎯\", got \"%s\"", str)
	}
}
