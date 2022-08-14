package thread

import (
	"testing"
)

func TestDecodeNCR(t *testing.T) {
	str := decodeNCR("家康の関東移封&#9876;左遷どころか大当たりだった&#127919;")
	if str != "家康の関東移封⚔左遷どころか大当たりだった🎯" {
		t.Errorf("want \"家康の関東移封⚔左遷どころか大当たりだった🎯\", got \"%s\"", str)
	}
}
