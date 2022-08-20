package thread_test

import (
	"testing"
	"time"

	"github.com/aethiopicuschan/thread"
)

func TestNewThread(t *testing.T) {
	// 普通のソース
	th, err := thread.NewThread("1660389180.dat<>(ヽ´ん`)「目を閉じて始めよう」  [697453962] (6)")
	if err != nil {
		t.Errorf("%s", err)
	}
	if th.BeID != "697453962" {
		t.Errorf("BeID is wrong, want 697453962, got %s", th.BeID)
	}
	if th.Title != "(ヽ´ん`)「目を閉じて始めよう」" {
		t.Errorf("Title is wrong, want \"(ヽ´ん`)「目を閉じて始めよう」\", got \"%s\"", th.Title)
	}
	if th.Time != 1660389180 {
		t.Errorf("Time is wrong, want 1660389180, got %d", th.Time)
	}
	if th.ResNum != 6 {
		t.Errorf("ResNum is wrong, want 6, got %d", th.ResNum)
	}
	// 怪しいソース
	th, err = thread.NewThread("9990000000.dat<>123.dat(81)[123456789]  [987654321] (100)")
	if err != nil {
		t.Errorf("%s", err)
	}
	if th.BeID != "987654321" {
		t.Errorf("BeID is wrong, want 987654321, got %s", th.BeID)
	}
	if th.Title != "123.dat(81)[123456789]" {
		t.Errorf("Title is wrong, want \"123.dat[123456789]\", got \"%s\"", th.Title)
	}
	if th.Time != 9990000000 {
		t.Errorf("Time is wrong, want 9990000000, got %d", th.Time)
	}
	if th.ResNum != 100 {
		t.Errorf("ResNum is wrong, want 100, got %d", th.ResNum)
	}
}

func TestIkioi(t *testing.T) {
	ikioi := thread.Ikioi(139, 1660986000, time.Unix(1660989360, 0))
	if ikioi != 3574.285645 {
		t.Errorf("Ikioi is wrong, want 3574.285645, got %f", ikioi)
	}
	ikioi = thread.Ikioi(6, 9245000000, time.Unix(1660989360, 0))
	if ikioi != 0 {
		t.Errorf("Ikioi is wrong, want 0, got %f", ikioi)
	}
}
