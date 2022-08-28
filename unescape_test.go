package thread_test

import (
	"testing"

	"github.com/aethiopicuschan/thread"
)

func TestUnescapeHtml(t *testing.T) {
	str := thread.UnescapeHtml("&quot;アレ&quot;みたくなる")
	if str != "\"アレ\"みたくなる" {
		t.Errorf("want \"\"アレ\"みたくなる\", got \"%s\"", str)
	}
}
