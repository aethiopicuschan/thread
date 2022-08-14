package thread

import (
	"testing"
)

func TestNewThread(t *testing.T) {
	// 普通のソース
	thread, err := newThread("1660389180.dat<>(ヽ´ん`)「目を閉じて始めよう」  [697453962] (6)")
	if err != nil {
		t.Errorf("%s", err)
	}
	if thread.BeID != "697453962" {
		t.Errorf("BeID is wrong, want 697453962, got %s", thread.BeID)
	}
	if thread.Title != "(ヽ´ん`)「目を閉じて始めよう」" {
		t.Errorf("Title is wrong, want \"(ヽ´ん`)「目を閉じて始めよう」\", got \"%s\"", thread.Title)
	}
	if thread.Time != 1660389180 {
		t.Errorf("Time is wrong, want 1660389180, got %d", thread.Time)
	}
	if thread.ResNum != 6 {
		t.Errorf("ResNum is wrong, want 6, got %d", thread.ResNum)
	}
	// 怪しいソース
	thread, err = newThread("9990000000.dat<>123.dat(81)[123456789]  [987654321] (100)")
	if err != nil {
		t.Errorf("%s", err)
	}
	if thread.BeID != "987654321" {
		t.Errorf("BeID is wrong, want 987654321, got %s", thread.BeID)
	}
	if thread.Title != "123.dat(81)[123456789]" {
		t.Errorf("Title is wrong, want \"123.dat[123456789]\", got \"%s\"", thread.Title)
	}
	if thread.Time != 9990000000 {
		t.Errorf("Time is wrong, want 9990000000, got %d", thread.Time)
	}
	if thread.ResNum != 100 {
		t.Errorf("ResNum is wrong, want 100, got %d", thread.ResNum)
	}
}
