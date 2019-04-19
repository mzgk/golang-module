package lib

import (
	"golang.org/x/text/width"
)

func Convert(s string) string {
	return width.Narrow.String(s)
}
