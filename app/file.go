package app

import (
	"path/filepath"
	"strings"

	"github.com/h2non/filetype"
)

var (
	mdType = filetype.NewType("md", "text/plain")
)

func init() {
	filetype.AddMatcher(mdType, fileMatcher_md)
}

func fileGetMIME(filename string) (string, bool) {

	mime := "application/octet-stream"
	extArr := strings.Split(filepath.Ext(filename), ".")
	if len(extArr) < 1 {
		return mime, false
	}
	ext := extArr[1]

	if !filetype.IsSupported(ext) {
		return mime, false
	}
	tp := filetype.GetType(ext)

	return tp.MIME.Value, true
}

func fileMatcher_md(buf []byte) bool {
	// return len(buf) > 1 && buf[0] == 0x23 && buf[1] == 0x20
	return true
}
