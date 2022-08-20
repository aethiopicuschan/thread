package thread

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Thread struct {
	Time   int64
	Title  string
	ResNum int
	Src    string
	BeID   string
	Ikioi  float32
}

// 勢い計算 テスト可能にするため切り出し
func ikioi(res int, time int64, now time.Time) float32 {
	// 「5ちゃんねるからのお知らせ」 などの特殊なスレッドで現在時刻を上回った値が設定されていることがある
	if time > now.Unix() {
		return 0
	}
	return float32(res) / (float32(now.Unix()-time) / 86400)
}

func newThread(src string) (thread Thread, err error) {
	r := regexp.MustCompile(`([0-9]+\.dat)<>(.+)\s+\(([0-9]+)\)`)
	if !r.MatchString(src) {
		err = errors.New("illegal source")
		return
	}
	a := r.FindStringSubmatch(src)
	// 元の文字列
	thread.Src = src
	// Unixtime
	thread.Time, err = strconv.ParseInt(strings.Split(a[1], ".dat")[0], 10, 64)
	// レス数
	thread.ResNum, err = strconv.Atoi(a[3])
	if err != nil {
		return
	}
	// BeIDとタイトル
	r2 := regexp.MustCompile(`\s\s\[([0-9]{9})\]$`)
	a2 := r2.FindStringSubmatch(a[2])
	if len(a2) > 0 {
		thread.BeID = a2[1]
		thread.Title = strings.Replace(a[2], fmt.Sprintf("  [%s]", a2[1]), "", -1)
	} else {
		thread.Title = a[2]
	}
	// 勢い
	thread.Ikioi = ikioi(thread.ResNum, thread.Time, time.Now())
	return
}
