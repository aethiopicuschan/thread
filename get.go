package thread

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// 引数の例: https://greta.5ch.net/poverty/
func Get(board string) (threads []Thread, err error) {
	url, err := url.Parse(board)
	if err != nil {
		return
	}
	url.Path = path.Join(url.Path, "subject.txt")
	res, err := http.Get(url.String())
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("StatusCode was %s.", res.Status))
	}
	raw, err := ioutil.ReadAll(transform.NewReader(res.Body, japanese.ShiftJIS.NewDecoder()))
	if err != nil {
		return
	}
	lines := strings.Split(decodeNCR(string(raw)), "\n")
	for _, line := range lines {
		thread, err := newThread(line)
		if err == nil {
			threads = append(threads, thread)
		}
	}
	return
}
